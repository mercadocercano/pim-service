#!/bin/bash
# =============================================================
# Ejecutar seeds de enriquecimiento de business type templates
# Seeds: 028, 029, 030, 031
# Requiere: columna generated_by en business_type_templates
# =============================================================

set -e

SEEDS_DIR="$(cd "$(dirname "$0")" && pwd)"

# Parámetros de conexión (defaults para local)
DB_HOST="${DB_HOST:-127.0.0.1}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-postgres}"
DB_NAME="${DB_NAME:-pim_db}"
PGPASSWORD="${PGPASSWORD:-postgres}"
export PGPASSWORD

PSQL="psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME"

echo "=== Conexión: $DB_USER@$DB_HOST:$DB_PORT/$DB_NAME ==="

# Paso 0: Agregar columna generated_by si no existe
echo "[0/5] Verificando columna generated_by..."
$PSQL -c "ALTER TABLE business_type_templates ADD COLUMN IF NOT EXISTS generated_by VARCHAR(100) DEFAULT 'manual';" 2>/dev/null

# Paso 1: Crear templates faltantes para business_types que no tienen
echo "[1/5] Creando templates base para business_types sin template..."
$PSQL -c "
INSERT INTO business_type_templates (business_type_id, name, description, version, region, is_default, is_active, categories)
SELECT bt.id, bt.name, 'Template de ' || bt.name, '1.0', 'AR', true, true, '[]'::jsonb
FROM business_types bt
WHERE NOT EXISTS (SELECT 1 FROM business_type_templates btt WHERE btt.business_type_id = bt.id);" 2>/dev/null

# Paso 2: Marcar como default los que no lo son (para que los UPDATE matcheen)
echo "[2/5] Marcando templates como default..."
$PSQL -c "
UPDATE business_type_templates SET is_default = true
WHERE is_default = false
AND business_type_id NOT IN (
  SELECT business_type_id FROM business_type_templates WHERE is_default = true
);" 2>/dev/null

# Paso 3: Ejecutar seeds
echo "[3/5] Ejecutando seed 028 (batch 1 - 11 rubros originales)..."
$PSQL -f "$SEEDS_DIR/028_templates_enriched_v4_hierarchical.sql" 2>&1 | grep -cE "UPDATE [1-9]|INSERT 0 [1-9]" | xargs -I{} echo "  {} operaciones exitosas"

echo "[4/5] Ejecutando seed 029 (batch 2 - 8 rubros)..."
$PSQL -f "$SEEDS_DIR/029_templates_enriched_v4_batch2.sql" 2>&1 | grep -cE "UPDATE [1-9]|INSERT 0 [1-9]" | xargs -I{} echo "  {} operaciones exitosas"

echo "[4/5] Ejecutando seed 030 (corralón + construcción en seco)..."
$PSQL -f "$SEEDS_DIR/030_new_business_types_construccion.sql" 2>&1 | grep -cE "UPDATE [1-9]|INSERT 0 [1-9]" | xargs -I{} echo "  {} operaciones exitosas"

echo "[5/5] Ejecutando seed 031 (6 rubros nuevos)..."
$PSQL -f "$SEEDS_DIR/031_new_business_types_batch2.sql" 2>&1 | grep -cE "UPDATE [1-9]|INSERT 0 [1-9]" | xargs -I{} echo "  {} operaciones exitosas"

# Verificación
echo ""
echo "=== Resultado final ==="
$PSQL -c "
SELECT bt.name, COALESCE(jsonb_array_length(btt.categories),0) as cats, COALESCE(btt.version,'—') as version
FROM business_types bt
LEFT JOIN business_type_templates btt ON bt.id = btt.business_type_id AND btt.is_default = true
WHERE COALESCE(jsonb_array_length(btt.categories),0) > 1
ORDER BY cats DESC;"

echo ""
TOTAL=$($PSQL -t -c "SELECT count(*) FROM business_type_templates btt WHERE jsonb_array_length(btt.categories) > 1;")
echo "Total templates enriquecidos: $TOTAL"
echo "=== Listo ==="
