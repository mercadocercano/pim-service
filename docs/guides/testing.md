# Testing — PIM Service

## Tests unitarios (Go)

```bash
go test ./...

# Con coverage real (incluye src/)
go test -coverprofile=coverage.out -coverpkg=./src/... ./...
go tool cover -html=coverage.out -o coverage.html
```

Los tests viven en `test/` y junto al código en `src/` (archivos `*_test.go`).

---

## Tests de integración (TestContainers)

Los tests de integración usan [TestContainers for Go](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/):
levantan un PostgreSQL real en Docker automáticamente, corren las migraciones y ejecutan los tests contra una BD aislada.

**Requisito**: Docker corriendo.

```bash
# Ejecutar todos los tests de integración
./run_tests.sh

# O directamente con Go
go test ./test-integration/... -v -timeout 120s
```

Los tests de integración están organizados por módulo:

| Módulo | Archivo test |
|--------|-------------|
| Global Catalog | `test-integration/global_catalog_*_test.go` |
| Attributes | `test-integration/attributes_*_test.go` |
| Quickstart | `test-integration/quickstart_*_test.go` |

---

## Tests de endpoints (bash)

Scripts en `test-integration/` para pruebas manuales o en CI contra un servicio levantado:

```bash
# Todos los endpoints
./test-integration/test-all-endpoints.sh

# Templates AI (quickstart wizard)
./test-integration/test-ai-templates.sh
```

**Variables de entorno para los scripts:**
```bash
export PIM_BASE_URL="http://localhost:8001/pim/api/v1"   # via Kong
# o directo:
export PIM_BASE_URL="http://localhost:8090/api/v1"

export TENANT_ID="9a4c3eb9-2471-4688-bfc8-973e5b3e4ce8"
export JWT_TOKEN="<token>"
```

---

## Ejecutar el stack para tests manuales

```bash
# Opción 1: con el lab completo
cd ~/Projects && make infra && make mc

# Opción 2: solo postgres + pim-service
docker compose up -d postgres
go run main.go
```

---

## Coverage actual

La medición correcta requiere `-coverpkg=./src/...` porque los tests de integración están en `test/` y `test-integration/`, fuera del paquete que instrumentan. Sin ese flag, Go reporta ~0.4% (artefacto de medición, no ausencia de tests).

```bash
go test -coverprofile=coverage.out -coverpkg=./src/... ./...
go tool cover -func=coverage.out | tail -1
```
