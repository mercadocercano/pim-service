# Tests de Integración — PIM Service

Tests contra el servicio levantado (bash + curl). Para tests de integración con TestContainers (Go), ver `docs/testing.md`.

## Scripts disponibles

```
test-integration/
├── run_integration_tests.sh    # Orquestador: ejecuta todos los scripts
├── test-all-endpoints.sh       # Smoke test de todos los endpoints
├── test-ai-templates.sh        # Tests del wizard de quickstart / AI templates
└── README.md                   # Esta documentación
```

## Uso rápido

```bash
# Ejecutar todos los tests
cd test-integration
./run_integration_tests.sh

# Solo un script
./test-all-endpoints.sh
./test-ai-templates.sh
```

## Configuración

```bash
# URL del servicio (default: via Kong)
export PIM_BASE_URL="http://localhost:8001/pim/api/v1"
# o directo al servicio:
export PIM_BASE_URL="http://localhost:8090/api/v1"

# Tenant ID para los tests
export TENANT_ID="9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8"

# JWT (si los endpoints requieren auth)
export JWT_TOKEN="<token>"
```

## Prerrequisitos

El servicio debe estar corriendo y con la base de datos migrada:

```bash
# Lab completo
cd ~/Projects && make infra && make mc

# O directo
./scripts/migrate.sh && go run main.go
```

## Opciones del orquestador

```bash
./run_integration_tests.sh --verbose      # Ver detalles de cada request
./run_integration_tests.sh --fail-fast    # Parar en el primer error
./run_integration_tests.sh --summary-only # Solo mostrar resumen final
```
