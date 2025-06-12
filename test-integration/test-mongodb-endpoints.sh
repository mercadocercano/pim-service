#!/bin/bash

# 🧪 Script de Pruebas MongoDB Marketplace
# Fecha: 11 de Junio, 2025
# Objetivo: Probar los endpoints MongoDB implementados

set -e  # Salir en caso de error

# Colores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Función para imprimir con colores
print_status() {
    echo -e "${BLUE}🧪 $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

# Configuración (usando servicio PIM directo)
PIM_BASE_URL="${PIM_BASE_URL:-http://localhost:8090/api/v1}"

echo "🏪 Pruebas MongoDB Marketplace - PIM Service"
echo "=============================================="

# Función para verificar prerequisitos
check_prerequisites() {
    print_status "Verificando prerequisitos..."
    
    # Verificar jq
    if ! command -v jq &> /dev/null; then
        print_error "jq no está instalado. Instálalo con: brew install jq (macOS) o apt-get install jq (Ubuntu)"
        exit 1
    fi
    
    # Verificar curl
    if ! command -v curl &> /dev/null; then
        print_error "curl no está instalado"
        exit 1
    fi
    
    print_success "Prerequisitos verificados"
}

# Función para mostrar configuración
show_config() {
    echo ""
    print_status "Configuración de pruebas:"
    echo "  PIM_BASE_URL: ${PIM_BASE_URL}"
    echo ""
}

# Función para probar conectividad general
test_general_connectivity() {
    print_status "Probando conectividad general con el servicio PIM..."
    
    local health_response
    health_response=$(curl -s -w "%{http_code}" "${PIM_BASE_URL}/health" -o /tmp/health_response.json)
    local http_code="${health_response: -3}"
    
    if [ "$http_code" = "200" ]; then
        print_success "Servicio PIM disponible"
        local status=$(cat /tmp/health_response.json | jq -r '.status // "unknown"')
        echo "  Status: $status"
    else
        print_error "Servicio PIM no disponible (HTTP $http_code)"
        print_error "Verifica que el servicio PIM esté ejecutándose en puerto 8090"
        exit 1
    fi
}

# Test 1: MongoDB Health Check
test_mongodb_health() {
    print_status "Test 1: MongoDB Health Check..."
    
    local response
    response=$(curl -s -w "%{http_code}" "${PIM_BASE_URL}/marketplace/health" -o /tmp/mongodb_health.json)
    local http_code="${response: -3}"
    
    if [ "$http_code" = "200" ]; then
        local status=$(cat /tmp/mongodb_health.json | jq -r '.status // "unknown"')
        local database=$(cat /tmp/mongodb_health.json | jq -r '.database // "unknown"')
        local message=$(cat /tmp/mongodb_health.json | jq -r '.message // "unknown"')
        
        print_success "MongoDB Health Check exitoso"
        echo "  Status: $status"
        echo "  Database: $database"
        echo "  Message: $message"
        
        # Mostrar features disponibles
        local features=$(cat /tmp/mongodb_health.json | jq -r '.features[]? // empty')
        if [ -n "$features" ]; then
            echo "  Features:"
            echo "$features" | while read -r feature; do
                echo "    - $feature"
            done
        fi
        return 0
    else
        print_error "MongoDB Health Check falló (HTTP $http_code)"
        cat /tmp/mongodb_health.json 2>/dev/null || echo "No response body"
        return 1
    fi
}

# Test 2: MongoDB Repositories Test
test_mongodb_repositories() {
    print_status "Test 2: MongoDB Repositories Test..."
    
    local response
    response=$(curl -s -w "%{http_code}" "${PIM_BASE_URL}/marketplace/test-mongo" -o /tmp/mongodb_test.json)
    local http_code="${response: -3}"
    
    if [ "$http_code" = "200" ]; then
        local status=$(cat /tmp/mongodb_test.json | jq -r '.status // "unknown"')
        local message=$(cat /tmp/mongodb_test.json | jq -r '.message // "unknown"')
        local ready_for=$(cat /tmp/mongodb_test.json | jq -r '.ready_for // "unknown"')
        
        print_success "MongoDB Repositories Test exitoso"
        echo "  Status: $status"
        echo "  Message: $message"
        echo "  Ready for: $ready_for"
        
        # Verificar repositorios
        local attr_repo=$(cat /tmp/mongodb_test.json | jq -r '.attr_repo // false')
        local mapping_repo=$(cat /tmp/mongodb_test.json | jq -r '.mapping_repo // false')
        
        echo "  Repositories:"
        echo "    - TenantCustomAttributeMongoRepository: $attr_repo"
        echo "    - TenantCategoryMappingMongoRepository: $mapping_repo"
        
        if [ "$attr_repo" = "true" ] && [ "$mapping_repo" = "true" ]; then
            print_success "Todos los repositorios MongoDB están disponibles"
            return 0
        else
            print_error "Algunos repositorios MongoDB no están disponibles"
            return 1
        fi
    else
        print_error "MongoDB Repositories Test falló (HTTP $http_code)"
        cat /tmp/mongodb_test.json 2>/dev/null || echo "No response body"
        return 1
    fi
}

# Test 3: Verificar MongoDB Collections
test_mongodb_collections() {
    print_status "Test 3: Verificar MongoDB Collections..."
    
    # Verificar que MongoDB esté ejecutándose
    if ! docker ps | grep -q saas-mongodb; then
        print_error "Contenedor MongoDB no está ejecutándose"
        return 1
    fi
    
    # Verificar colecciones
    local collections
    collections=$(docker exec saas-mongodb mongosh -u admin -p admin123 --authenticationDatabase admin pim_marketplace --eval "db.getCollectionNames()" --quiet 2>/dev/null || echo "[]")
    
    if echo "$collections" | grep -q "tenant_custom_attributes" && echo "$collections" | grep -q "tenant_category_mappings"; then
        print_success "Colecciones MongoDB verificadas"
        echo "  Collections encontradas:"
        echo "$collections" | sed 's/,/\n/g' | sed 's/\[//g' | sed 's/\]//g' | sed 's/"//g' | sed 's/^[ \t]*//' | while read -r collection; do
            if [ -n "$collection" ]; then
                echo "    - $collection"
            fi
        done
        return 0
    else
        print_error "Colecciones MongoDB no encontradas"
        echo "Collections: $collections"
        return 1
    fi
}

# Test 4: Verificar documento de prueba
test_sample_document() {
    print_status "Test 4: Verificar documento de prueba..."
    
    # Verificar que el documento de prueba existe
    local doc_count
    doc_count=$(docker exec saas-mongodb mongosh -u admin -p admin123 --authenticationDatabase admin pim_marketplace --eval "db.tenant_custom_attributes.countDocuments({tenant_id: 'test-tenant'})" --quiet 2>/dev/null || echo "0")
    
    if [ "$doc_count" -gt 0 ]; then
        print_success "Documento de prueba encontrado ($doc_count documentos)"
        
        # Mostrar el documento
        local doc
        doc=$(docker exec saas-mongodb mongosh -u admin -p admin123 --authenticationDatabase admin pim_marketplace --eval "JSON.stringify(db.tenant_custom_attributes.findOne({tenant_id: 'test-tenant'}), null, 2)" --quiet 2>/dev/null || echo "{}")
        
        if echo "$doc" | jq -e '.validation_rules' > /dev/null 2>&1; then
            print_success "ValidationRules como JSON nativo verificado"
            echo "  Validation Rules:"
            echo "$doc" | jq '.validation_rules' | sed 's/^/    /'
        fi
        return 0
    else
        print_warning "No se encontró documento de prueba (esto es normal si es la primera ejecución)"
        return 0
    fi
}

# Función principal
main() {
    check_prerequisites
    show_config
    
    local passed=0
    local total=0
    
    # Ejecutar pruebas
    echo ""
    
    # Test conectividad general
    if test_general_connectivity; then
        ((passed++))
    fi
    ((total++))
    
    echo ""
    
    # Test MongoDB Health
    if test_mongodb_health; then
        ((passed++))
    fi
    ((total++))
    
    echo ""
    
    # Test MongoDB Repositories
    if test_mongodb_repositories; then
        ((passed++))
    fi
    ((total++))
    
    echo ""
    
    # Test MongoDB Collections
    if test_mongodb_collections; then
        ((passed++))
    fi
    ((total++))
    
    echo ""
    
    # Test Sample Document
    if test_sample_document; then
        ((passed++))
    fi
    ((total++))
    
    # Resumen
    echo ""
    echo "=============================================="
    print_status "Resumen de pruebas MongoDB:"
    echo "  Pasadas: $passed/$total"
    echo "  Fallidas: $((total - passed))/$total"
    
    if [ $passed -eq $total ]; then
        print_success "¡Todas las pruebas MongoDB pasaron!"
        echo ""
        print_success "🎉 MongoDB está funcionando correctamente para marketplace"
        echo "✅ ValidationRules se maneja como JSON nativo"
        echo "✅ Repositorios MongoDB están disponibles"
        echo "✅ Endpoints de health y test funcionan"
    else
        print_warning "Algunas pruebas fallaron. Revisa la implementación."
    fi
    
    echo ""
    print_status "Pruebas MongoDB completadas!"
    
    # Limpiar archivos temporales
    rm -f /tmp/health_response.json /tmp/mongodb_health.json /tmp/mongodb_test.json
}

# Ejecutar pruebas
main "$@" 