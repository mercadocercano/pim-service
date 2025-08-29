#!/bin/bash

# Script de pruebas de integración para AI Templates
# Este script prueba los endpoints de templates inteligentes con AI

set -e

# Colores para output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuración
BASE_URL="${BASE_URL:-http://localhost:8090}"
API_URL="${BASE_URL}/api/v1"
TENANT_ID="${TENANT_ID:-00000000-0000-0000-0000-000000000001}"
AUTH_TOKEN="${AUTH_TOKEN:-test-token}"

# Variables para almacenar IDs generados
TEMPLATE_ID=""
BUSINESS_TYPE_ID="4f4e9b9e-7b8a-4c6a-9c5a-3e5f7a8b9c1d" # Almacén/Kiosco

echo "🤖 Iniciando pruebas de AI Templates"
echo "=================================="
echo "URL Base: $BASE_URL"
echo "Tenant ID: $TENANT_ID"
echo ""

# Función para imprimir resultado
print_result() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✅ $2 exitoso${NC}"
    else
        echo -e "${RED}❌ $2 falló${NC}"
        exit 1
    fi
}

# Función para hacer requests con headers
make_request() {
    local method=$1
    local endpoint=$2
    local data=$3
    
    if [ -z "$data" ]; then
        curl -s -w "\n%{http_code}" -X $method \
            -H "Authorization: Bearer $AUTH_TOKEN" \
            -H "X-Tenant-ID: $TENANT_ID" \
            -H "Content-Type: application/json" \
            "$API_URL$endpoint"
    else
        curl -s -w "\n%{http_code}" -X $method \
            -H "Authorization: Bearer $AUTH_TOKEN" \
            -H "X-Tenant-ID: $TENANT_ID" \
            -H "Content-Type: application/json" \
            -d "$data" \
            "$API_URL$endpoint"
    fi
}

# Test 1: Generar Template Inteligente
echo -e "${YELLOW}Test 1: Generar Template Inteligente${NC}"
echo "------------------------------------"

GENERATE_DATA='{
  "business_type_id": "'$BUSINESS_TYPE_ID'",
  "name": "Mi Almacén Premium AI",
  "target_size": "medium",
  "preferences": {
    "price_range": "standard",
    "include_generics": true,
    "generic_percentage": 25,
    "categories_focus": ["bebidas", "snacks", "limpieza"],
    "exclude_brands": [],
    "regional_preferences": "buenos_aires"
  }
}'

response=$(make_request "POST" "/templates/generate" "$GENERATE_DATA")
http_code=$(echo "$response" | tail -n 1)
body=$(echo "$response" | sed '$d')

if [ "$http_code" = "200" ]; then
    TEMPLATE_ID=$(echo "$body" | grep -o '"template_id":"[^"]*' | cut -d'"' -f4)
    echo "Template ID generado: $TEMPLATE_ID"
    echo "Response: $body" | jq '.' 2>/dev/null || echo "$body"
    print_result 0 "Generación de template"
else
    echo "HTTP Code: $http_code"
    echo "Response: $body"
    print_result 1 "Generación de template"
fi

echo ""

# Test 2: Aplicar Template (si se generó exitosamente)
if [ ! -z "$TEMPLATE_ID" ]; then
    echo -e "${YELLOW}Test 2: Aplicar Template al Catálogo${NC}"
    echo "------------------------------------"
    
    APPLY_DATA='{
      "customizations": {
        "exclude_products": [],
        "price_multiplier": 1.1,
        "quantity_adjustments": {}
      },
      "apply_options": {
        "create_categories": true,
        "create_brands": true,
        "create_products": true,
        "initial_stock": false
      }
    }'
    
    response=$(make_request "POST" "/templates/$TEMPLATE_ID/apply" "$APPLY_DATA")
    http_code=$(echo "$response" | tail -n 1)
    body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" = "200" ]; then
        echo "Response: $body" | jq '.' 2>/dev/null || echo "$body"
        print_result 0 "Aplicación de template"
    else
        echo "HTTP Code: $http_code"
        echo "Response: $body"
        print_result 1 "Aplicación de template"
    fi
    
    echo ""
fi

# Test 3: Obtener Performance del Template
if [ ! -z "$TEMPLATE_ID" ]; then
    echo -e "${YELLOW}Test 3: Obtener Performance del Template${NC}"
    echo "----------------------------------------"
    
    response=$(make_request "GET" "/templates/$TEMPLATE_ID/performance" "")
    http_code=$(echo "$response" | tail -n 1)
    body=$(echo "$response" | sed '$d')
    
    if [ "$http_code" = "200" ]; then
        echo "Response: $body" | jq '.' 2>/dev/null || echo "$body"
        print_result 0 "Obtención de performance"
    else
        echo "HTTP Code: $http_code"
        echo "Response: $body"
        print_result 1 "Obtención de performance"
    fi
    
    echo ""
fi

# Test 4: Actualizar Template con Feedback
echo -e "${YELLOW}Test 4: Actualizar Template con Feedback${NC}"
echo "----------------------------------------"

UPDATE_DATA='{
  "template_id": "'$TEMPLATE_ID'",
  "feedback_items": [
    {
      "product_id": "00000000-0000-0000-0000-000000000001",
      "action": "removed",
      "reason": "No se vende bien en esta zona"
    },
    {
      "product_id": "00000000-0000-0000-0000-000000000002",
      "action": "quantity_changed",
      "new_quantity": 24,
      "reason": "Alta demanda"
    }
  ],
  "optimization_goal": "maximize_satisfaction"
}'

response=$(make_request "POST" "/templates/update-from-feedback" "$UPDATE_DATA")
http_code=$(echo "$response" | tail -n 1)
body=$(echo "$response" | sed '$d')

if [ "$http_code" = "200" ]; then
    echo "Response: $body" | jq '.' 2>/dev/null || echo "$body"
    print_result 0 "Actualización con feedback"
else
    echo "HTTP Code: $http_code"
    echo "Response: $body"
    print_result 1 "Actualización con feedback"
fi

echo ""

# Resumen
echo "=================================="
echo -e "${GREEN}🎉 Pruebas de AI Templates completadas${NC}"
echo ""
echo "Nota: Algunos tests pueden fallar si:"
echo "1. El servicio AI Gateway no está disponible"
echo "2. No hay productos en el catálogo global"
echo "3. Las migraciones no se han ejecutado"
echo ""
echo "Para ejecutar las migraciones:"
echo "  cd scripts && ./run_migration_034.sh"
echo ""