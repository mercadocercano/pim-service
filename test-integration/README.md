# 🧪 Tests de Integración - PIM Service

Este directorio contiene todos los tests de integración para el servicio PIM, organizados para facilitar su ejecución y mantenimiento.

## 📁 Estructura

```
test-integration/
├── run_integration_tests.sh          # Script maestro para ejecutar todos los tests
├── test-mongodb-endpoints.sh         # Tests básicos de endpoints MongoDB
├── test-mongodb-crud-endpoints.sh    # Tests de operaciones CRUD en MongoDB
├── test-marketplace-endpoints.sh     # Tests completos del flujo Marketplace
└── README.md                         # Esta documentación
```

## 🚀 Uso Rápido

### Ejecutar todos los tests
```bash
cd test-integration
./run_integration_tests.sh
```

### Ejecutar con opciones
```bash
# Modo verbose (ver detalles)
./run_integration_tests.sh --verbose

# Solo tests de MongoDB
./run_integration_tests.sh --mongodb-only

# Solo tests de Marketplace
./run_integration_tests.sh --marketplace-only

# Parar en el primer error
./run_integration_tests.sh --fail-fast

# Mostrar solo resumen
./run_integration_tests.sh --summary-only
```

### Ejecutar test específico
```bash
./run_integration_tests.sh test-marketplace-endpoints.sh
```

## 📋 Tests Disponibles

### 1. **test-mongodb-endpoints.sh**
- **Descripción**: Tests básicos de conectividad y endpoints MongoDB
- **Cobertura**: Conexión, health checks, operaciones básicas
- **Duración**: ~30 segundos

### 2. **test-mongodb-crud-endpoints.sh**
- **Descripción**: Tests de operaciones CRUD completas en MongoDB
- **Cobertura**: Create, Read, Update, Delete para todas las entidades
- **Duración**: ~60 segundos

### 3. **test-marketplace-endpoints.sh**
- **Descripción**: Tests completos del flujo Marketplace
- **Cobertura**: 
  - Creación de categorías marketplace
  - Validación de jerarquías
  - Mapeos de categorías tenant
  - Atributos personalizados
  - Sincronización de cambios
  - Validaciones de autorización
- **Duración**: ~90 segundos

## ⚙️ Configuración

### Variables de Entorno

```bash
# URL del servicio PIM (default: http://localhost:8001/pim/api/v1)
export PIM_BASE_URL="http://localhost:8001/pim/api/v1"

# ID del tenant para tests (default: test tenant)
export TENANT_ID="9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8"

# Token JWT para autenticación (opcional)
export AUTH_TOKEN="your-jwt-token-here"

# No limpiar datos de test (default: false)
export SKIP_CLEANUP="false"

# Timeout por test en segundos (default: 300)
export TEST_TIMEOUT="300"
```

### Prerequisitos

1. **Servicios Docker ejecutándose**:
   ```bash
   docker-compose up -d pim-service saas-mongodb
   ```

2. **Herramientas instaladas**:
   - `curl` - Para peticiones HTTP
   - `jq` - Para procesamiento JSON
   - `docker` - Para verificar servicios

## 📊 Interpretación de Resultados

### Códigos de Salida
- `0` - Todos los tests pasaron
- `1` - Algunos tests fallaron
- `2` - Error en prerequisitos

### Formato de Salida
```
🧪 Test Name... ✅ (exitoso) / ❌ (fallido)
📊 REPORTE FINAL:
  ⏱️  Tiempo total: 120s
  📈 Tests ejecutados: 3
  ✅ Tests exitosos: 3
  ❌ Tests fallidos: 0
  📊 Tasa de éxito: 100%
```

## 🔧 Desarrollo

### Agregar un Nuevo Test

1. **Crear el script de test**:
   ```bash
   touch test-integration/test-new-feature.sh
   chmod +x test-integration/test-new-feature.sh
   ```

2. **Actualizar el script maestro**:
   Agregar el nuevo test al array `INTEGRATION_TESTS` en `run_integration_tests.sh`:
   ```bash
   declare -a INTEGRATION_TESTS=(
       "test-mongodb-endpoints.sh:MongoDB Basic Endpoints"
       "test-mongodb-crud-endpoints.sh:MongoDB CRUD Operations"
       "test-marketplace-endpoints.sh:Marketplace Complete Flow"
       "test-new-feature.sh:New Feature Tests"  # ← Agregar aquí
   )
   ```

### Estructura de un Test

```bash
#!/bin/bash

# Configuración
set -e
source "$(dirname "$0")/common.sh"  # Si existe archivo común

# Variables
TEST_NAME="Feature Test"
BASE_URL="${PIM_BASE_URL:-http://localhost:8001/pim/api/v1}"

# Funciones de utilidad
print_success() { echo -e "\033[0;32m✅ $1\033[0m"; }
print_error() { echo -e "\033[0;31m❌ $1\033[0m"; }

# Tests
test_feature_endpoint() {
    local response=$(curl -s "$BASE_URL/feature")
    # Validaciones...
}

# Ejecución principal
main() {
    echo "🧪 Ejecutando $TEST_NAME..."
    test_feature_endpoint
    echo "✅ $TEST_NAME completado"
}

main "$@"
```

## 🐛 Troubleshooting

### Problemas Comunes

1. **Error de conexión**:
   ```
   ❌ No se puede conectar al servicio PIM
   ```
   **Solución**: Verificar que los servicios estén ejecutándose:
   ```bash
   docker-compose ps
   docker-compose up -d pim-service
   ```

2. **Tests fallando por datos existentes**:
   ```
   ❌ Error: duplicate key error
   ```
   **Solución**: Limpiar datos de test:
   ```bash
   docker exec saas-mongodb mongosh --eval "use pim_marketplace; db.dropDatabase();"
   docker-compose restart pim-service
   ```

3. **Timeout en tests**:
   ```
   ❌ Test timeout after 300s
   ```
   **Solución**: Aumentar timeout:
   ```bash
   export TEST_TIMEOUT=600
   ```

### Logs de Debug

Para obtener más información sobre fallos:

```bash
# Logs del servicio PIM
docker logs pim-service --tail 50

# Logs de MongoDB
docker logs saas-mongodb --tail 50

# Ejecutar test individual con verbose
./test-marketplace-endpoints.sh --verbose
```

## 📈 Métricas y Monitoreo

### Integración con CI/CD

```yaml
# .github/workflows/integration-tests.yml
name: Integration Tests
on: [push, pull_request]

jobs:
  integration-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Start services
        run: docker-compose up -d
      - name: Run integration tests
        run: |
          cd services/saas-mt-pim-service/test-integration
          ./run_integration_tests.sh --summary-only
```

### Reportes de Cobertura

Los tests generan métricas que pueden ser integradas con herramientas de monitoreo:

- Tiempo de ejecución por test
- Tasa de éxito/fallo
- Cobertura de endpoints
- Performance de respuesta

---

## 📞 Soporte

Para problemas o mejoras en los tests de integración:

1. Revisar logs de servicios
2. Verificar prerequisitos
3. Consultar esta documentación
4. Crear issue con detalles del problema

---

**Última actualización**: $(date +%Y-%m-%d)
**Versión**: 1.0.0 