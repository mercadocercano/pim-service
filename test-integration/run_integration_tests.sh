#!/bin/bash

# =============================================================================
# 🧪 PIM Service - Integration Tests Runner
# =============================================================================
# Este script ejecuta todos los tests de integración del servicio PIM
# Autor: Sistema de Migración PIM
# Fecha: $(date +%Y-%m-%d)
# =============================================================================

set -e  # Salir en caso de error

# Colores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Función para imprimir con colores
print_header() {
    echo -e "${PURPLE}$1${NC}"
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

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_step() {
    echo -e "${CYAN}🧪 $1${NC}"
}

# Variables globales
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0
SKIPPED_TESTS=0

# Array de tests disponibles
declare -a INTEGRATION_TESTS=(
    "test-mongodb-endpoints.sh:MongoDB Basic Endpoints"
    "test-mongodb-crud-endpoints.sh:MongoDB CRUD Operations"
    "test-marketplace-endpoints.sh:Marketplace Complete Flow"
)

# Función para mostrar ayuda
show_help() {
    cat << EOF
🧪 PIM Service - Integration Tests Runner

USAGE:
    $0 [OPTIONS] [TEST_NAMES...]

OPTIONS:
    -h, --help              Mostrar esta ayuda
    -v, --verbose           Modo verbose (más detalles)
    -f, --fail-fast         Parar en el primer error
    -l, --list              Listar tests disponibles
    -s, --skip-setup        Saltar verificación de prerequisitos
    --mongodb-only          Ejecutar solo tests de MongoDB
    --marketplace-only      Ejecutar solo tests de Marketplace
    --summary-only          Mostrar solo el resumen final

EXAMPLES:
    $0                                    # Ejecutar todos los tests
    $0 --list                            # Listar tests disponibles
    $0 test-marketplace-endpoints.sh     # Ejecutar test específico
    $0 --mongodb-only                    # Solo tests de MongoDB
    $0 --fail-fast --verbose            # Modo verbose, parar en primer error

ENVIRONMENT VARIABLES:
    PIM_BASE_URL           URL base del servicio PIM (default: http://localhost:8001/pim/api/v1)
    TENANT_ID              ID del tenant para tests (default: test tenant)
    AUTH_TOKEN             Token JWT para autenticación
    SKIP_CLEANUP           No limpiar datos de test (default: false)
    TEST_TIMEOUT           Timeout por test en segundos (default: 300)

EOF
}

# Función para listar tests disponibles
list_tests() {
    print_header "📋 Tests de Integración Disponibles:"
    echo
    local index=1
    for test_info in "${INTEGRATION_TESTS[@]}"; do
        local test_file=$(echo "$test_info" | cut -d':' -f1)
        local test_desc=$(echo "$test_info" | cut -d':' -f2)
        echo -e "  ${index}. ${CYAN}${test_file}${NC} - ${test_desc}"
        ((index++))
    done
    echo
}

# Función para verificar prerequisitos
check_prerequisites() {
    print_step "Verificando prerequisitos..."
    
    local errors=0
    
    # Verificar que estamos en el directorio correcto
    if [[ ! -f "$PROJECT_ROOT/main.go" ]]; then
        print_error "No se encontró main.go. Ejecutar desde el directorio del proyecto PIM"
        ((errors++))
    fi
    
    # Verificar herramientas necesarias
    local tools=("curl" "jq" "docker")
    for tool in "${tools[@]}"; do
        if ! command -v "$tool" &> /dev/null; then
            print_error "Herramienta requerida no encontrada: $tool"
            ((errors++))
        fi
    done
    
    # Verificar servicios Docker
    if ! docker ps | grep -q "pim-service"; then
        print_error "Servicio pim-service no está ejecutándose"
        print_info "Ejecutar: docker-compose up -d pim-service"
        ((errors++))
    fi
    
    if ! docker ps | grep -q "saas-mongodb"; then
        print_error "Servicio MongoDB no está ejecutándose"
        print_info "Ejecutar: docker-compose up -d saas-mongodb"
        ((errors++))
    fi
    
    # Verificar conectividad
    local pim_url="${PIM_BASE_URL:-http://localhost:8001/pim/api/v1}"
    if ! curl -s "$pim_url/health" > /dev/null; then
        print_error "No se puede conectar al servicio PIM en: $pim_url"
        ((errors++))
    fi
    
    if [[ $errors -gt 0 ]]; then
        print_error "Se encontraron $errors errores en prerequisitos"
        return 1
    fi
    
    print_success "Prerequisitos verificados correctamente"
    return 0
}

# Función para ejecutar un test individual
run_single_test() {
    local test_file="$1"
    local test_desc="$2"
    local verbose="$3"
    
    print_step "Ejecutando: $test_desc"
    
    if [[ ! -f "$SCRIPT_DIR/$test_file" ]]; then
        print_error "Test no encontrado: $test_file"
        return 1
    fi
    
    # Hacer el script ejecutable
    chmod +x "$SCRIPT_DIR/$test_file"
    
    local start_time=$(date +%s)
    local test_output
    local exit_code
    
    # Ejecutar el test
    if [[ "$verbose" == "true" ]]; then
        echo "----------------------------------------"
        "$SCRIPT_DIR/$test_file"
        exit_code=$?
        echo "----------------------------------------"
    else
        test_output=$("$SCRIPT_DIR/$test_file" 2>&1)
        exit_code=$?
    fi
    
    local end_time=$(date +%s)
    local duration=$((end_time - start_time))
    
    if [[ $exit_code -eq 0 ]]; then
        print_success "$test_desc completado en ${duration}s"
        return 0
    else
        print_error "$test_desc falló en ${duration}s"
        if [[ "$verbose" != "true" && -n "$test_output" ]]; then
            echo "Output del test:"
            echo "$test_output"
        fi
        return 1
    fi
}

# Función para generar reporte final
generate_report() {
    local total_time="$1"
    
    echo
    print_header "📊 REPORTE FINAL DE TESTS DE INTEGRACIÓN"
    echo "========================================================"
    echo -e "⏱️  Tiempo total de ejecución: ${total_time}s"
    echo -e "📈 Tests ejecutados: $TOTAL_TESTS"
    echo -e "✅ Tests exitosos: $PASSED_TESTS"
    echo -e "❌ Tests fallidos: $FAILED_TESTS"
    echo -e "⏭️  Tests omitidos: $SKIPPED_TESTS"
    echo
    
    local success_rate=0
    if [[ $TOTAL_TESTS -gt 0 ]]; then
        success_rate=$((PASSED_TESTS * 100 / TOTAL_TESTS))
    fi
    
    echo -e "📊 Tasa de éxito: ${success_rate}%"
    
    if [[ $FAILED_TESTS -eq 0 ]]; then
        print_success "🎉 ¡TODOS LOS TESTS PASARON!"
        echo
        return 0
    else
        print_error "⚠️  ALGUNOS TESTS FALLARON"
        echo
        return 1
    fi
}

# Función principal
main() {
    local verbose=false
    local fail_fast=false
    local skip_setup=false
    local mongodb_only=false
    local marketplace_only=false
    local summary_only=false
    local specific_tests=()
    
    # Parsear argumentos
    while [[ $# -gt 0 ]]; do
        case $1 in
            -h|--help)
                show_help
                exit 0
                ;;
            -v|--verbose)
                verbose=true
                shift
                ;;
            -f|--fail-fast)
                fail_fast=true
                shift
                ;;
            -l|--list)
                list_tests
                exit 0
                ;;
            -s|--skip-setup)
                skip_setup=true
                shift
                ;;
            --mongodb-only)
                mongodb_only=true
                shift
                ;;
            --marketplace-only)
                marketplace_only=true
                shift
                ;;
            --summary-only)
                summary_only=true
                shift
                ;;
            test-*.sh)
                specific_tests+=("$1")
                shift
                ;;
            *)
                print_error "Opción desconocida: $1"
                show_help
                exit 1
                ;;
        esac
    done
    
    # Header
    print_header "🧪 PIM SERVICE - INTEGRATION TESTS RUNNER"
    print_header "=========================================="
    echo
    
    # Verificar prerequisitos
    if [[ "$skip_setup" != "true" ]]; then
        if ! check_prerequisites; then
            exit 1
        fi
        echo
    fi
    
    # Determinar qué tests ejecutar
    local tests_to_run=()
    
    if [[ ${#specific_tests[@]} -gt 0 ]]; then
        # Tests específicos
        for test in "${specific_tests[@]}"; do
            tests_to_run+=("$test:Test Específico")
        done
    elif [[ "$mongodb_only" == "true" ]]; then
        # Solo tests de MongoDB
        tests_to_run+=(
            "test-mongodb-endpoints.sh:MongoDB Basic Endpoints"
            "test-mongodb-crud-endpoints.sh:MongoDB CRUD Operations"
        )
    elif [[ "$marketplace_only" == "true" ]]; then
        # Solo tests de Marketplace
        tests_to_run+=("test-marketplace-endpoints.sh:Marketplace Complete Flow")
    else
        # Todos los tests
        tests_to_run=("${INTEGRATION_TESTS[@]}")
    fi
    
    print_info "Se ejecutarán ${#tests_to_run[@]} tests de integración"
    echo
    
    # Ejecutar tests
    local start_time=$(date +%s)
    
    for test_info in "${tests_to_run[@]}"; do
        local test_file=$(echo "$test_info" | cut -d':' -f1)
        local test_desc=$(echo "$test_info" | cut -d':' -f2)
        
        ((TOTAL_TESTS++))
        
        if [[ "$summary_only" == "true" ]]; then
            # Solo mostrar que se está ejecutando
            echo -n "🧪 $test_desc... "
        fi
        
        if run_single_test "$test_file" "$test_desc" "$verbose"; then
            ((PASSED_TESTS++))
            if [[ "$summary_only" == "true" ]]; then
                echo -e "${GREEN}✅${NC}"
            fi
        else
            ((FAILED_TESTS++))
            if [[ "$summary_only" == "true" ]]; then
                echo -e "${RED}❌${NC}"
            fi
            
            if [[ "$fail_fast" == "true" ]]; then
                print_error "Modo fail-fast activado. Deteniendo ejecución."
                break
            fi
        fi
        
        echo
    done
    
    local end_time=$(date +%s)
    local total_time=$((end_time - start_time))
    
    # Generar reporte
    generate_report "$total_time"
}

# Ejecutar función principal con todos los argumentos
main "$@" 