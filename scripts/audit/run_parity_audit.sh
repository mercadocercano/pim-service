#!/usr/bin/env bash
# run_parity_audit.sh — Corre pim_parity_audit.sql contra local y producción
# y genera un diff. 100% READ-ONLY (solo SELECT).
#
# Prerrequisito para producción: túnel SSH abierto en otra terminal:
#   ssh -N -L 5433:localhost:5432 root@<IP_POSTGRESQL_DROPLET>
#
# Variables de entorno:
#   LOCAL_PGPASSWORD   (default: postgres)
#   PROD_PGPASSWORD    (obligatoria para auditar prod; no se imprime ni se guarda)
#   PROD_PORT          (default: 5433 — el extremo local del túnel SSH)
#
# Uso:
#   ./run_parity_audit.sh local           # solo local
#   PROD_PGPASSWORD=… ./run_parity_audit.sh both   # local + prod + diff (default)

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SQL="$SCRIPT_DIR/pim_parity_audit.sql"
OUT_DIR="$SCRIPT_DIR/out"
mkdir -p "$OUT_DIR"

MODE="${1:-both}"
LOCAL_PGPASSWORD="${LOCAL_PGPASSWORD:-postgres}"
PROD_PORT="${PROD_PORT:-5433}"

run_local() {
  echo ">> Auditando LOCAL (localhost:5432/pim_db)…"
  PGPASSWORD="$LOCAL_PGPASSWORD" psql -h localhost -p 5432 -U postgres -d pim_db \
    -f "$SQL" > "$OUT_DIR/out_local.txt" 2>&1
  echo "   -> $OUT_DIR/out_local.txt"
}

run_prod() {
  if [ -z "${PROD_PGPASSWORD:-}" ]; then
    echo "!! PROD_PGPASSWORD no seteada — no se audita prod." >&2
    return 1
  fi
  echo ">> Auditando PROD (localhost:$PROD_PORT/pim_db via túnel SSH)…"
  PGPASSWORD="$PROD_PGPASSWORD" psql -h localhost -p "$PROD_PORT" -U postgres -d pim_db \
    -f "$SQL" > "$OUT_DIR/out_prod.txt" 2>&1
  echo "   -> $OUT_DIR/out_prod.txt"
}

case "$MODE" in
  local) run_local ;;
  prod)  run_prod ;;
  both)
    run_local
    run_prod
    echo ">> Diff (local <vs> prod)…"
    diff -u "$OUT_DIR/out_local.txt" "$OUT_DIR/out_prod.txt" > "$OUT_DIR/parity_diff.txt" || true
    echo "   -> $OUT_DIR/parity_diff.txt"
    echo ""
    echo "=== RESUMEN DEL DIFF ==="
    if [ -s "$OUT_DIR/parity_diff.txt" ]; then
      cat "$OUT_DIR/parity_diff.txt"
    else
      echo "Sin diferencias: local y prod están igualadas."
    fi
    ;;
  *) echo "Modo inválido: $MODE (usar: local | prod | both)" >&2; exit 1 ;;
esac
