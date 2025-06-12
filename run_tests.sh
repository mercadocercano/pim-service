#!/bin/bash

# =============================================================================
# 🧪 PIM Service - Tests Runner (Convenience Script)
# =============================================================================
# Script de conveniencia para ejecutar tests de integración desde el directorio raíz
# =============================================================================

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEST_DIR="$SCRIPT_DIR/test-integration"

# Verificar que el directorio de tests existe
if [[ ! -d "$TEST_DIR" ]]; then
    echo "❌ Error: Directorio de tests no encontrado: $TEST_DIR"
    exit 1
fi

# Verificar que el script maestro existe
if [[ ! -f "$TEST_DIR/run_integration_tests.sh" ]]; then
    echo "❌ Error: Script maestro no encontrado: $TEST_DIR/run_integration_tests.sh"
    exit 1
fi

# Ejecutar el script maestro con todos los argumentos pasados
echo "🚀 Ejecutando tests de integración desde: $TEST_DIR"
echo

cd "$TEST_DIR"
exec "./run_integration_tests.sh" "$@" 