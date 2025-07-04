#!/bin/bash

# Script de tests de integración para el Wizard de Quickstart
# Asume que el servicio PIM está ejecutándose en localhost:8090

set -e  # Exit on error

# Configuración
BASE_URL="http://localhost:8090/api/v1"
TENANT_ID="9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8"  # Tenant demo existente
BUSINESS_TYPE_ID=""
WIZARD_ID=""

# Colores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Función para imprimir con colores
print_test() {
    echo -e "${BLUE}🧪 TEST: $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ SUCCESS: $1${NC}"
}

print_error() {
    echo -e "${RED}❌ ERROR: $1${NC}"
}

print_info() {
    echo -e "${YELLOW}ℹ️  INFO: $1${NC}"
}

# Función para hacer requests HTTP con headers
make_request() {
    local method=$1
    local endpoint=$2
    local data=$3
    
    if [ -z "$data" ]; then
        curl -s -X "$method" \
            -H "Content-Type: application/json" \
            -H "X-Tenant-ID: $TENANT_ID" \
            "$BASE_URL$endpoint"
    else
        curl -s -X "$method" \
            -H "Content-Type: application/json" \
            -H "X-Tenant-ID: $TENANT_ID" \
            -d "$data" \
            "$BASE_URL$endpoint"
    fi
}

# Función para verificar que el servicio está disponible
check_service() {
    print_test "Verificando que el servicio PIM esté disponible"
    
    response=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/health" || echo "000")
    
    if [ "$response" = "200" ]; then
        print_success "Servicio PIM disponible"
    else
        print_error "Servicio PIM no disponible (HTTP $response)"
        print_info "Asegúrate de que el servicio esté ejecutándose en $BASE_URL"
        exit 1
    fi
}

# Test 1: Obtener business types disponibles
test_get_business_types() {
    print_test "Obteniendo business types disponibles"
    
    response=$(make_request "GET" "/business-types")
    http_code=$(echo "$response" | tail -n1)
    
    if echo "$response" | jq -e '.business_types' > /dev/null 2>&1; then
        count=$(echo "$response" | jq '.business_types | length')
        print_success "Business types obtenidos: $count encontrados"
        
        # Extraer el primer business type ID para usar en tests posteriores
        BUSINESS_TYPE_ID=$(echo "$response" | jq -r '.business_types[0].id')
        print_info "Usando business type ID: $BUSINESS_TYPE_ID"
    else
        print_error "Error obteniendo business types: $response"
        exit 1
    fi
}

# Test 2: Iniciar wizard
test_start_wizard() {
    print_test "Iniciando wizard de configuración"
    
    if [ -z "$BUSINESS_TYPE_ID" ]; then
        print_error "Business type ID no disponible"
        exit 1
    fi
    
    data="{\"business_type_id\": \"$BUSINESS_TYPE_ID\"}"
    response=$(make_request "POST" "/wizard/start" "$data")
    
    if echo "$response" | jq -e '.wizard_id' > /dev/null 2>&1; then
        WIZARD_ID=$(echo "$response" | jq -r '.wizard_id')
        print_success "Wizard iniciado exitosamente. ID: $WIZARD_ID"
        print_info "Tenant ID: $(echo "$response" | jq -r '.tenant_id')"
        print_info "Business Type ID: $(echo "$response" | jq -r '.business_type_id')"
    else
        print_error "Error iniciando wizard: $response"
        exit 1
    fi
}

# Test 3: Obtener estado del wizard
test_get_wizard_status() {
    print_test "Obteniendo estado actual del wizard"
    
    response=$(make_request "GET" "/wizard/status")
    
    if echo "$response" | jq -e '.wizard_id' > /dev/null 2>&1; then
        wizard_id=$(echo "$response" | jq -r '.wizard_id')
        setup_data=$(echo "$response" | jq -r '.setup_data')
        print_success "Estado del wizard obtenido"
        print_info "Wizard ID: $wizard_id"
        print_info "Setup data: $setup_data"
    else
        print_error "Error obteniendo estado del wizard: $response"
        exit 1
    fi
}

# Test 4: Obtener datos del template
test_get_template_data() {
    print_test "Obteniendo datos completos del template"
    
    if [ -z "$BUSINESS_TYPE_ID" ]; then
        print_error "Business type ID no disponible"
        exit 1
    fi
    
    response=$(make_request "GET" "/wizard/template/$BUSINESS_TYPE_ID")
    
    if echo "$response" | jq -e '.template_data' > /dev/null 2>&1; then
        categories_count=$(echo "$response" | jq '.template_data.categories | length')
        products_count=$(echo "$response" | jq '.template_data.products | length')
        brands_count=$(echo "$response" | jq '.template_data.brands | length')
        
        print_success "Datos del template obtenidos"
        print_info "Categorías: $categories_count"
        print_info "Productos: $products_count"
        print_info "Marcas: $brands_count"
    else
        print_error "Error obteniendo datos del template: $response"
        exit 1
    fi
}

# Test 5: Obtener sección específica del template (categorías)
test_get_template_categories() {
    print_test "Obteniendo categorías del template"
    
    if [ -z "$BUSINESS_TYPE_ID" ]; then
        print_error "Business type ID no disponible"
        exit 1
    fi
    
    response=$(make_request "GET" "/wizard/template/$BUSINESS_TYPE_ID/categories")
    
    if echo "$response" | jq -e '.data' > /dev/null 2>&1; then
        section=$(echo "$response" | jq -r '.section')
        count=$(echo "$response" | jq '.data | length')
        print_success "Categorías obtenidas: $count en sección '$section'"
    else
        print_error "Error obteniendo categorías: $response"
        exit 1
    fi
}

# Test 6: Actualizar step del wizard (selección de categorías)
test_update_wizard_step() {
    print_test "Actualizando step del wizard (selección de categorías)"
    
    data='{
        "current_step": "categories_selection",
        "step_data": {
            "selected_categories": ["electronics", "home"],
            "custom_categories": ["Mi Categoría Personalizada"],
            "timestamp": "'$(date -u +%Y-%m-%dT%H:%M:%SZ)'"
        }
    }'
    
    response=$(make_request "PUT" "/wizard/step" "$data")
    
    if echo "$response" | jq -e '.wizard_id' > /dev/null 2>&1; then
        updated_at=$(echo "$response" | jq -r '.updated_at')
        print_success "Step del wizard actualizado"
        print_info "Actualizado en: $updated_at"
    else
        print_error "Error actualizando step del wizard: $response"
        exit 1
    fi
}

# Test 7: Obtener sección específica del template (productos con paginación)
test_get_template_products_paginated() {
    print_test "Obteniendo productos del template con paginación"
    
    if [ -z "$BUSINESS_TYPE_ID" ]; then
        print_error "Business type ID no disponible"
        exit 1
    fi
    
    response=$(make_request "GET" "/wizard/template/$BUSINESS_TYPE_ID/products?page=1&page_size=3")
    
    if echo "$response" | jq -e '.data' > /dev/null 2>&1; then
        current_page=$(echo "$response" | jq '.pagination.page')
        total_items=$(echo "$response" | jq '.pagination.total_items')
        has_next=$(echo "$response" | jq '.pagination.has_next')
        
        print_success "Productos paginados obtenidos"
        print_info "Página: $current_page"
        print_info "Total items: $total_items"
        print_info "Tiene siguiente: $has_next"
    else
        print_error "Error obteniendo productos paginados: $response"
        exit 1
    fi
}

# Test 8: Actualizar múltiples steps del wizard
test_multiple_wizard_steps() {
    print_test "Actualizando múltiples steps del wizard"
    
    # Step 1: Marcas
    print_info "Actualizando selección de marcas..."
    data1='{
        "current_step": "brands_selection",
        "step_data": {
            "selected_brands": ["Samsung", "LG", "Sony"],
            "custom_brands": ["Mi Marca Local"]
        }
    }'
    
    response1=$(make_request "PUT" "/wizard/step" "$data1")
    if ! echo "$response1" | jq -e '.wizard_id' > /dev/null 2>&1; then
        print_error "Error en step de marcas: $response1"
        exit 1
    fi
    
    # Step 2: Productos
    print_info "Actualizando selección de productos..."
    data2='{
        "current_step": "products_selection",
        "step_data": {
            "selected_products": [1, 2, 3, 4, 5],
            "custom_products": ["Mi Producto Personalizado"]
        }
    }'
    
    response2=$(make_request "PUT" "/wizard/step" "$data2")
    if ! echo "$response2" | jq -e '.wizard_id' > /dev/null 2>&1; then
        print_error "Error en step de productos: $response2"
        exit 1
    fi
    
    print_success "Múltiples steps actualizados exitosamente"
}

# Test 9: Verificar estado final del wizard
test_final_wizard_status() {
    print_test "Verificando estado final del wizard"
    
    response=$(make_request "GET" "/wizard/status")
    
    if echo "$response" | jq -e '.wizard_id' > /dev/null 2>&1; then
        setup_data=$(echo "$response" | jq -r '.setup_data')
        current_step=$(echo "$setup_data" | jq -r '.step')
        completed_steps=$(echo "$setup_data" | jq -r '.completed_steps | length')
        
        print_success "Estado final verificado"
        print_info "Step actual: $current_step"
        print_info "Steps completados: $completed_steps"
        print_info "Setup data: $setup_data"
    else
        print_error "Error verificando estado final: $response"
        exit 1
    fi
}

# Test 10: Test de endpoints con errores (casos edge)
test_error_cases() {
    print_test "Probando casos de error"
    
    # Test con business type ID inválido
    print_info "Probando business type ID inválido..."
    response=$(make_request "GET" "/wizard/template/invalid-id")
    if echo "$response" | jq -e '.error' > /dev/null 2>&1; then
        print_success "Error manejado correctamente para business type inválido"
    else
        print_error "Error no manejado correctamente: $response"
    fi
    
    # Test con sección inválida
    print_info "Probando sección inválida..."
    response=$(make_request "GET" "/wizard/template/$BUSINESS_TYPE_ID/invalid_section")
    if echo "$response" | jq -e '.error' > /dev/null 2>&1; then
        print_success "Error manejado correctamente para sección inválida"
    else
        print_error "Error no manejado correctamente: $response"
    fi
    
    # Test con step data inválido
    print_info "Probando step data inválido..."
    invalid_data='{"current_step": "", "step_data": {}}'
    response=$(make_request "PUT" "/wizard/step" "$invalid_data")
    if echo "$response" | jq -e '.error' > /dev/null 2>&1; then
        print_success "Error manejado correctamente para step data inválido"
    else
        print_error "Error no manejado correctamente: $response"
    fi
}

# Función principal que ejecuta todos los tests
run_all_tests() {
    echo "==========================================="
    echo "🚀 INICIANDO TESTS DE INTEGRACIÓN WIZARD"
    echo "==========================================="
    echo ""
    
    check_service
    echo ""
    
    test_get_business_types
    echo ""
    
    test_start_wizard
    echo ""
    
    test_get_wizard_status
    echo ""
    
    test_get_template_data
    echo ""
    
    test_get_template_categories
    echo ""
    
    test_update_wizard_step
    echo ""
    
    test_get_template_products_paginated
    echo ""
    
    test_multiple_wizard_steps
    echo ""
    
    test_final_wizard_status
    echo ""
    
    test_error_cases
    echo ""
    
    echo "==========================================="
    print_success "TODOS LOS TESTS COMPLETADOS EXITOSAMENTE"
    echo "==========================================="
}

# Función para mostrar ayuda
show_help() {
    echo "Usage: $0 [OPTION]"
    echo ""
    echo "Opciones:"
    echo "  --help, -h          Mostrar esta ayuda"
    echo "  --service-check     Solo verificar que el servicio esté disponible"
    echo "  --basic             Ejecutar solo tests básicos"
    echo "  --all               Ejecutar todos los tests (por defecto)"
    echo ""
    echo "Variables de entorno:"
    echo "  BASE_URL           URL base del servicio (default: http://localhost:8090/api/v1)"
    echo "  TENANT_ID          ID del tenant para tests (default: 9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8)"
    echo ""
    echo "Ejemplos:"
    echo "  $0                 # Ejecutar todos los tests"
    echo "  $0 --basic         # Ejecutar solo tests básicos"
    echo "  BASE_URL=http://localhost:8090/api/v1 $0  # Usar URL personalizada"
}

# Función para tests básicos
run_basic_tests() {
    echo "==========================================="
    echo "🚀 EJECUTANDO TESTS BÁSICOS"
    echo "==========================================="
    echo ""
    
    check_service
    echo ""
    
    test_get_business_types
    echo ""
    
    test_start_wizard
    echo ""
    
    test_get_wizard_status
    echo ""
    
    print_success "TESTS BÁSICOS COMPLETADOS"
}

# Main script
case "${1:-}" in
    --help|-h)
        show_help
        exit 0
        ;;
    --service-check)
        check_service
        exit 0
        ;;
    --basic)
        run_basic_tests
        exit 0
        ;;
    --all|"")
        run_all_tests
        exit 0
        ;;
    *)
        echo "Opción no reconocida: $1"
        show_help
        exit 1
        ;;
esac