#!/bin/bash
# apply-seeds-prod.sh
# Aplica seeds faltantes en producción (PostgreSQL en Kubernetes).
# Crea tabla de tracking pim_seeds_applied si no existe.
# Seeds son idempotentes — re-correr uno existente es seguro pero se evita.
#
# Uso:
#   ./apply-seeds-prod.sh                    # aplica todos los faltantes
#   ./apply-seeds-prod.sh --dry-run          # muestra qué correría sin ejecutar
#   ./apply-seeds-prod.sh --from 126         # aplica solo seeds >= NNN
#   ./apply-seeds-prod.sh --only 127         # aplica solo seed NNN exacto
#
# Requisitos:
#   - kubectl configurado apuntando al cluster de prod
#   - KUBECONFIG=/home/deploy/.kube/config si corrés desde droplet

set -euo pipefail

# ── Config ────────────────────────────────────────────────────────────────────
SEEDS_DIR="$(cd "$(dirname "$0")/../seeds" && pwd)"
POD="${PIM_PG_POD:-postgres-8b69b9-h44sr}"
PG_USER="${PIM_PG_USER:-postgres}"
PG_DB="${PIM_PG_DB:-pim_db}"
REMOTE_TMP="/tmp/pim_seeds"
DRY_RUN=false
FROM_NUM=0
ONLY_NUM=""

# ── Arg parsing ───────────────────────────────────────────────────────────────
while [[ $# -gt 0 ]]; do
  case "$1" in
    --dry-run)   DRY_RUN=true ;;
    --from)      FROM_NUM="$2"; shift ;;
    --only)      ONLY_NUM="$2"; shift ;;
    --pod)       POD="$2"; shift ;;
    --db)        PG_DB="$2"; shift ;;
    *) echo "Unknown arg: $1"; exit 1 ;;
  esac
  shift
done

# ── Colores ───────────────────────────────────────────────────────────────────
RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'; CYAN='\033[0;36m'; NC='\033[0m'
info()    { echo -e "${CYAN}[INFO]${NC} $*"; }
ok()      { echo -e "${GREEN}[ OK ]${NC} $*"; }
warn()    { echo -e "${YELLOW}[WARN]${NC} $*"; }
error()   { echo -e "${RED}[ERR ]${NC} $*"; }

# ── Helper: psql en pod ───────────────────────────────────────────────────────
pg() {
  kubectl exec "$POD" -- psql -U "$PG_USER" -d "$PG_DB" -tAq -c "$1"
}

pg_file() {
  kubectl exec "$POD" -- psql -U "$PG_USER" -d "$PG_DB" -f "$1"
}

# ── Verificar kubectl ─────────────────────────────────────────────────────────
echo ""
info "Pod: $POD | DB: $PG_DB"

if ! kubectl get pod "$POD" &>/dev/null; then
  error "Pod '$POD' no encontrado. Verificá KUBECONFIG o --pod."
  exit 1
fi

# ── Crear tabla de tracking si no existe ─────────────────────────────────────
info "Verificando tabla de tracking pim_seeds_applied..."
pg "
CREATE TABLE IF NOT EXISTS pim_seeds_applied (
  seed_name   TEXT PRIMARY KEY,
  applied_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  duration_ms INT
);" > /dev/null
ok "Tabla de tracking lista."

# ── Obtener seeds ya aplicados ────────────────────────────────────────────────
APPLIED=$(pg "SELECT seed_name FROM pim_seeds_applied ORDER BY seed_name;")

# ── Listar seeds locales ──────────────────────────────────────────────────────
SEED_FILES=$(ls "$SEEDS_DIR"/*.sql 2>/dev/null | sort)

if [[ -z "$SEED_FILES" ]]; then
  error "No se encontraron archivos .sql en $SEEDS_DIR"
  exit 1
fi

# ── Calcular pendientes ───────────────────────────────────────────────────────
PENDING=()
for filepath in $SEED_FILES; do
  filename=$(basename "$filepath")
  num="${filename%%_*}"  # extrae NNN del nombre NNN_descripcion.sql

  # Filtros
  num_dec=$((10#$num))  # fuerza base 10 (evita octal con leading zeros)
  if [[ -n "$ONLY_NUM" && "$num_dec" -ne "$((10#$ONLY_NUM))" ]]; then continue; fi
  if [[ -n "$FROM_NUM" && "$num_dec" -lt "$((10#$FROM_NUM))" ]]; then continue; fi

  if echo "$APPLIED" | grep -qx "$filename"; then
    : # ya aplicado, skip
  else
    PENDING+=("$filepath")
  fi
done

# ── Resumen ───────────────────────────────────────────────────────────────────
TOTAL_LOCAL=$(echo "$SEED_FILES" | wc -l | xargs)
TOTAL_APPLIED=$(echo "$APPLIED" | grep -c . || true)
TOTAL_PENDING=${#PENDING[@]}

echo ""
echo "──────────────────────────────────────────────"
echo "  Seeds locales   : $TOTAL_LOCAL"
echo "  Ya aplicados    : $TOTAL_APPLIED"
echo "  Pendientes      : $TOTAL_PENDING"
echo "──────────────────────────────────────────────"

if [[ $TOTAL_PENDING -eq 0 ]]; then
  ok "Prod está al día. Nada que aplicar."
  exit 0
fi

echo ""
info "Seeds pendientes:"
for f in "${PENDING[@]}"; do
  echo "  → $(basename "$f")"
done

if $DRY_RUN; then
  echo ""
  warn "Modo dry-run. Nada fue ejecutado."
  exit 0
fi

# ── Confirmar ─────────────────────────────────────────────────────────────────
echo ""
read -rp "Aplicar $TOTAL_PENDING seed(s) en PROD? [s/N] " REPLY
if [[ ! "$REPLY" =~ ^[Ss]$ ]]; then
  warn "Cancelado."
  exit 0
fi

# ── Backup antes de aplicar ───────────────────────────────────────────────────
BACKUP_DIR="${BACKUP_DIR:-/tmp}"
BACKUP_FILE="$BACKUP_DIR/pim_db_backup_$(date +%Y%m%d_%H%M%S).sql"
info "Generando backup → $BACKUP_FILE ..."
kubectl exec "$POD" -- pg_dump -U "$PG_USER" "$PG_DB" --no-owner --no-acl > "$BACKUP_FILE"
ok "Backup guardado: $BACKUP_FILE ($(du -sh "$BACKUP_FILE" | cut -f1))"

# ── Crear directorio remoto ───────────────────────────────────────────────────
kubectl exec "$POD" -- mkdir -p "$REMOTE_TMP"

# ── Aplicar seeds ─────────────────────────────────────────────────────────────
echo ""
ERRORS=0
for filepath in "${PENDING[@]}"; do
  filename=$(basename "$filepath")
  remote_path="$REMOTE_TMP/$filename"

  info "Copiando $filename..."
  kubectl cp "$filepath" "$POD:$remote_path"

  info "Aplicando $filename..."
  START_MS=$(date +%s%3N)

  if pg_file "$remote_path"; then
    END_MS=$(date +%s%3N)
    DURATION=$(( END_MS - START_MS ))
    pg "INSERT INTO pim_seeds_applied (seed_name, duration_ms) VALUES ('$filename', $DURATION)
        ON CONFLICT (seed_name) DO UPDATE SET applied_at = NOW(), duration_ms = $DURATION;" > /dev/null
    ok "$filename (${DURATION}ms)"
  else
    error "FALLÓ: $filename"
    ERRORS=$(( ERRORS + 1 ))
    # No abort — continuar con el resto para ver el scope total del problema
  fi

  # Limpiar archivo remoto
  kubectl exec "$POD" -- rm -f "$remote_path"
done

# ── Limpiar directorio remoto ─────────────────────────────────────────────────
kubectl exec "$POD" -- rmdir "$REMOTE_TMP" 2>/dev/null || true

# ── Resultado final ───────────────────────────────────────────────────────────
echo ""
echo "──────────────────────────────────────────────"
if [[ $ERRORS -gt 0 ]]; then
  error "Completado con $ERRORS error(s). Revisá los logs arriba."
  exit 1
else
  ok "Todos los seeds aplicados exitosamente."

  # Mostrar conteo final en prod
  PROD_COUNT=$(pg "SELECT COUNT(*) FROM global_products;")
  PROD_WITH_IMG=$(pg "SELECT COUNT(*) FROM global_products WHERE image_url IS NOT NULL AND image_url != '';")
  echo ""
  echo "  global_products total      : $PROD_COUNT"
  echo "  global_products con imagen : $PROD_WITH_IMG"
  PCT=$(( PROD_WITH_IMG * 100 / PROD_COUNT ))
  echo "  Cobertura de imágenes      : ${PCT}%"
  echo "──────────────────────────────────────────────"
fi
