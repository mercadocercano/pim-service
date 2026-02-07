#!/bin/bash

# Script para regenerar la colección de Postman desde el OpenAPI spec
# Fecha: 2026-02-05

set -e

echo "🚀 Generando colección de Postman desde OpenAPI..."

# Colores para output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Verificar si npx está disponible
if ! command -v npx &> /dev/null
then
    echo -e "${RED}❌ Error: npx no está instalado${NC}"
    echo "Por favor instala Node.js que incluye npx"
    exit 1
fi

echo -e "${GREEN}✅ Usando npx para ejecutar openapi-to-postmanv2${NC}"

# Directorio del script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Archivos
OPENAPI_FILE="$PROJECT_ROOT/api-docs/openapi.yaml"
POSTMAN_COLLECTION="$PROJECT_ROOT/postman_collection.json"
POSTMAN_ENVIRONMENT="$PROJECT_ROOT/postman_environment.json"
BACKUP_DIR="$PROJECT_ROOT/api-docs/legacy"

# Verificar que existe el OpenAPI
if [ ! -f "$OPENAPI_FILE" ]; then
    echo -e "${RED}❌ Error: No se encuentra el archivo OpenAPI: $OPENAPI_FILE${NC}"
    exit 1
fi

# Crear backup de la colección actual si existe
if [ -f "$POSTMAN_COLLECTION" ]; then
    echo -e "${YELLOW}📦 Creando backup de la colección actual...${NC}"
    mkdir -p "$BACKUP_DIR"
    BACKUP_FILE="$BACKUP_DIR/postman_collection_backup_$(date +%Y%m%d_%H%M%S).json"
    cp "$POSTMAN_COLLECTION" "$BACKUP_FILE"
    echo -e "${GREEN}✅ Backup creado: $BACKUP_FILE${NC}"
fi

# Generar nueva colección
echo -e "${YELLOW}🔄 Convirtiendo OpenAPI a Postman Collection...${NC}"

npx -y openapi-to-postmanv2 \
    -s "$OPENAPI_FILE" \
    -o "$POSTMAN_COLLECTION" \
    -p

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Colección de Postman generada exitosamente${NC}"
    echo -e "   Archivo: $POSTMAN_COLLECTION"
else
    echo -e "${RED}❌ Error generando la colección${NC}"
    exit 1
fi

# Generar environment si no existe
if [ ! -f "$POSTMAN_ENVIRONMENT" ]; then
    echo -e "${YELLOW}🌍 Generando archivo de environment...${NC}"
    cat > "$POSTMAN_ENVIRONMENT" <<EOF
{
  "name": "PIM Service Development",
  "values": [
    {
      "key": "base_url",
      "value": "http://localhost:8090",
      "enabled": true,
      "type": "default"
    },
    {
      "key": "kong_base_url",
      "value": "http://localhost:8001/pim",
      "enabled": true,
      "type": "default"
    },
    {
      "key": "tenant_id",
      "value": "",
      "enabled": true,
      "type": "default",
      "description": "UUID del tenant de prueba"
    },
    {
      "key": "auth_token",
      "value": "",
      "enabled": true,
      "type": "secret",
      "description": "JWT token para autenticación"
    },
    {
      "key": "user_role",
      "value": "tenant_admin",
      "enabled": true,
      "type": "default",
      "description": "Rol del usuario (super_admin, marketplace_admin, tenant_admin, tenant_user)"
    }
  ],
  "_postman_variable_scope": "environment"
}
EOF
    echo -e "${GREEN}✅ Environment file creado: $POSTMAN_ENVIRONMENT${NC}"
fi

# Resumen
echo ""
echo -e "${GREEN}╔════════════════════════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║  ✅ Generación completada exitosamente                      ║${NC}"
echo -e "${GREEN}╚════════════════════════════════════════════════════════════╝${NC}"
echo ""
echo -e "📁 Archivos generados:"
echo -e "   • Colección:   ${POSTMAN_COLLECTION}"
echo -e "   • Environment: ${POSTMAN_ENVIRONMENT}"
echo ""
echo -e "📖 Próximos pasos:"
echo -e "   1. Importar ${POSTMAN_COLLECTION} en Postman"
echo -e "   2. Importar ${POSTMAN_ENVIRONMENT} en Postman"
echo -e "   3. Configurar las variables de environment:"
echo -e "      - tenant_id: Obtener de IAM service"
echo -e "      - auth_token: Obtener token JWT"
echo -e "   4. Seleccionar environment 'PIM Service Development'"
echo ""
echo -e "🔗 URLs configuradas:"
echo -e "   • Directo:  http://localhost:8090/api/v1"
echo -e "   • Kong:     http://localhost:8001/pim/api/v1"
echo ""

# Estadísticas de la colección
ENDPOINT_COUNT=$(grep -o '"method":' "$POSTMAN_COLLECTION" | wc -l)
echo -e "📊 Estadísticas:"
echo -e "   • Total de endpoints: $ENDPOINT_COUNT"
echo ""

exit 0
