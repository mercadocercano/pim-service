#!/bin/bash

# =============================================================================
# 🧪 PIM Service - Test Completo de Todos los Endpoints
# =============================================================================
# Este script prueba todos los endpoints implementados en el PIM Service
# Basado en el análisis de controladores y OpenAPI actualizado
# =============================================================================

set -e

# Colores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m'

# Variables de configuración
PIM_BASE_URL="${PIM_BASE_URL:-http://localhost:8001/pim/api/v1}"
AUTH_TOKEN="${AUTH_TOKEN:-eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test-token}"
TENANT_ID="${TENANT_ID:-9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8}"

# Contadores
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Variables para IDs generados durante los tests
BUSINESS_TYPE_ID=""
BRAND_ID=""
CATEGORY_ID=""
ATTRIBUTE_ID=""
CATEGORY_ATTRIBUTE_ID=""
PRODUCT_ID=""

# Funciones de utilidad
print_header() {
    echo -e "${PURPLE}===============================================${NC}"
    echo -e "${PURPLE}$1${NC}"
    echo -e "${PURPLE}===============================================${NC}"
}

print_section() {
    echo -e "\n${CYAN}🔍 $1${NC}"
}

print_test() {
    echo -e "${BLUE}   Testing: $1${NC}"
}

print_success() {
    echo -e "${GREEN}   ✅ $1${NC}"
    ((PASSED_TESTS++))
}

print_error() {
    echo -e "${RED}   ❌ $1${NC}"
    ((FAILED_TESTS++))
}

print_warning() {
    echo -e "${YELLOW}   ⚠️  $1${NC}"
}

# Función para hacer peticiones HTTP
test_endpoint() {
    local method="$1"
    local endpoint="$2"
    local expected_status="$3"
    local description="$4"
    
    ((TOTAL_TESTS++))
    
    local response
    response=$(curl -s -w "%{http_code}" -X "$method" \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $AUTH_TOKEN" \
        -H "X-Tenant-ID: $TENANT_ID" \
        "$PIM_BASE_URL$endpoint" 2>/dev/null)
    
    local http_code="${response: -3}"
    
    if [ "$http_code" = "$expected_status" ]; then
        print_success "$description (HTTP $http_code)"
        return 0
    else
        print_error "$description (Expected: $expected_status, Got: $http_code)"
        return 1
    fi
}

# Función para extraer ID del JSON response
extract_id() {
    local json="$1"
    echo "$json" | jq -r '.id // .data.id // empty'
}

# Verificar prerequisitos
check_prerequisites() {
    print_section "Verificando prerequisitos"
    
    if ! command -v curl &> /dev/null; then
        print_error "curl no está instalado"
        exit 1
    fi
    
    # jq es opcional para estos tests básicos
    if ! command -v jq &> /dev/null; then
        print_warning "jq no está instalado (funcionalidad reducida)"
    fi
    
    # Verificar conectividad
    if curl -s "$PIM_BASE_URL/health" > /dev/null; then
        print_success "PIM Service disponible"
    else
        print_error "PIM Service no disponible en $PIM_BASE_URL"
        exit 1
    fi
}

# Tests de endpoints
test_health_endpoints() {
    print_section "Health Endpoints"
    test_endpoint "GET" "/health" "200" "Health Check"
}

test_business_types_endpoints() {
    print_section "Business Types Endpoints"
    test_endpoint "GET" "/business-types" "200" "Listar business types"
    test_endpoint "GET" "/business-types?only_active=true" "200" "Listar business types activos"
}

test_brands_endpoints() {
    print_section "Brands Endpoints"
    test_endpoint "GET" "/brands" "200" "Listar brands"
    test_endpoint "GET" "/brands?page=1&page_size=5" "200" "Listar brands con paginación"
}

test_categories_endpoints() {
    print_section "Categories Endpoints"
    test_endpoint "GET" "/categories" "200" "Listar categories con criterios"
    test_endpoint "GET" "/categories/simple" "200" "Listar categories simple"
    test_endpoint "GET" "/categories/tree" "200" "Obtener árbol de categories"
}

test_attributes_endpoints() {
    print_section "Attributes Endpoints"
    test_endpoint "GET" "/attributes" "200" "Listar attributes"
}

test_category_attributes_endpoints() {
    print_section "Category-Attributes Endpoints"
    test_endpoint "GET" "/category-attributes" "200" "Listar category-attributes con criterios"
    test_endpoint "GET" "/category-attributes/simple" "200" "Listar category-attributes simple"
}

test_products_endpoints() {
    print_section "Products Endpoints"
    test_endpoint "GET" "/products" "200" "Listar products"
    test_endpoint "GET" "/products?page=1&page_size=5" "200" "Listar products con paginación"
}

test_global_catalog_endpoints() {
    print_section "Global Catalog Endpoints"
    test_endpoint "GET" "/public/global-catalog/health" "200" "Health check global catalog"
    test_endpoint "GET" "/public/global-catalog/search?ean=1234567890123" "404" "Búsqueda pública por EAN"
    test_endpoint "GET" "/public/global-catalog/suggestions?business_type=RESTAURANT" "200" "Sugerencias por tipo de negocio"
    test_endpoint "GET" "/global-catalog/products" "200" "Listar productos global catalog"
}

test_quickstart_endpoints() {
    print_section "Quickstart Endpoints"
    test_endpoint "GET" "/quickstart/business-types" "200" "Obtener business types para quickstart"
    test_endpoint "GET" "/quickstart/categories/RESTAURANT" "200" "Obtener categories por business type"
    test_endpoint "GET" "/quickstart/attributes/RESTAURANT" "200" "Obtener attributes por business type"
    test_endpoint "GET" "/quickstart/variants/RESTAURANT" "200" "Obtener variants por business type"
    test_endpoint "GET" "/quickstart/products/RESTAURANT" "200" "Obtener products por business type"
    test_endpoint "GET" "/quickstart/brands/RESTAURANT" "200" "Obtener brands por business type"
}

# Función para mostrar resumen final
show_summary() {
    print_header "RESUMEN FINAL"
    
    echo -e "${CYAN}📊 Estadísticas de Tests:${NC}"
    echo -e "   Total tests ejecutados: ${TOTAL_TESTS}"
    echo -e "   ${GREEN}✅ Tests exitosos: ${PASSED_TESTS}${NC}"
    echo -e "   ${RED}❌ Tests fallidos: ${FAILED_TESTS}${NC}"
    
    local success_rate=0
    if [ $TOTAL_TESTS -gt 0 ]; then
        success_rate=$((PASSED_TESTS * 100 / TOTAL_TESTS))
    fi
    
    echo -e "   📈 Tasa de éxito: ${success_rate}%"
    
    if [ $FAILED_TESTS -eq 0 ]; then
        echo -e "\n${GREEN}🎉 ¡Todos los tests pasaron exitosamente!${NC}"
        exit 0
    else
        echo -e "\n${RED}💥 Algunos tests fallaron. Revisar la salida anterior.${NC}"
        exit 1
    fi
}

# =============================================================================
# EJECUCIÓN PRINCIPAL
# =============================================================================

main() {
    print_header "PIM SERVICE - TEST COMPLETO DE ENDPOINTS"
    
    echo -e "${BLUE}🔧 Configuración:${NC}"
    echo -e "   PIM_BASE_URL: ${PIM_BASE_URL}"
    echo -e "   TENANT_ID: ${TENANT_ID}"
    echo -e "   AUTH_TOKEN: ${AUTH_TOKEN:0:20}..."
    echo ""
    
    check_prerequisites
    
    # Ejecutar todos los tests
    test_health_endpoints
    test_business_types_endpoints
    test_brands_endpoints
    test_categories_endpoints
    test_attributes_endpoints
    test_category_attributes_endpoints
    test_products_endpoints
    test_global_catalog_endpoints
    test_quickstart_endpoints
    
    show_summary
}

# Ejecutar si es llamado directamente
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi 