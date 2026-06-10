# AUDIT.md — PIM Service

**Fecha:** 2026-06-09 | **Rama:** master | **Skill:** go-hex-audit

---

## 1. Resumen ejecutivo

| Fase | Resultado |
|------|-----------|
| 0 Discovery | PASS — arquitectura hexagonal modular bien definida en todos los bounded contexts excepto `s2s` |
| 1 Compilación | PASS (con advertencia go.work) — 20+ archivos sin formato gofmt |
| 2 Auditoría hex/DDD | **FAIL** — 2 violaciones CRITICAL, 2 HIGH, 1 LOW |
| 3 Coverage | **FAIL** — 0.4% medido (problema de configuración de coverage, no ausencia de tests) |
| 4 OpenAPI | PARCIAL — spec sustancial v2.3.0 existe, rutas wizard/s2s/backfill faltantes |
| 5 Postman | INFO — colección existe en raíz (`postman_collection.json`), no auditada en detalle |
| 6 e2e newman | SKIP — SQL Postgres-específico (JSONB, arrays) hace inviable SQLite; ya existe path con Testcontainers |

---

## 2. Phase 0 — Discovery

**Módulo:** `saas-mt-pim-service` · Go 1.24.0  
**Framework HTTP:** Gin (`github.com/gin-gonic/gin`)  
**Auth:** JWT via `github.com/mercadocercano/middleware.TenantValidation`; rutas públicas excluidas en lista explícita en `main.go:76-96`  
**Persistencia:** PostgreSQL (`database/sql` + `lib/pq`). MongoDB opcional para global catalog.  
**DTOs:** Sí — paquetes `application/request` y `application/response` separados de entidades de dominio.  
**OpenAPI:** `api-docs/openapi.yaml` (3143 líneas, v2.3.0)  
**Postman:** `postman_collection.json` en raíz  
**CI:** `.github/workflows/` existe  
**Migraciones:** `migrations/` existe  

**Mapa de capas (hexagonal modular):**

```
src/
  {bounded_context}/
    domain/
      entity/       ← entidades puras
      port/         ← interfaces de repositorio/servicio
      exception/    ← errores de dominio
      value_object/ ← (donde aplica)
    application/
      usecase/      ← casos de uso
      request/      ← DTOs de entrada
      response/     ← DTOs de salida
    infrastructure/
      controller/   ← handlers HTTP (Gin)
      persistence/  ← repositorios Postgres/MongoDB
      config/       ← wiring del módulo
```

Bounded contexts: `brand`, `attribute`, `category`, `category_attribute`, `businesstype`, `product/tenant`, `product/global_catalog`, `product/quickstart`, `quickstart`, `schema_validation`, `overview`, `s2s` (plano, sin capas).

**Veredicto SQLite e2e:** NO viable. Queries JSONB (`btt.categories` en `s2s/usecase`), `pg_dump`-style CTEs, y posibles `RETURNING`/`ON CONFLICT`. Usar Testcontainers (ya disponible en go.mod).

---

## 3. Phase 1 — Compilación

```
go build ./...   → PASS (exit 0)
go vet ./...     → no ejecutado (go.work conflict bloqueó el comando anterior en entorno de dev)
gofmt -l .       → 20+ archivos sin formatear
```

**Advertencia go.work:** `../iam-service` requiere go ≥ 1.25.0 pero `go.work` declara 1.24.0. No bloquea build, pero rompe `go test ./...` sin `GOWORK=off`. Solución: `go work use` desde el workspace raíz.

**Archivos sin formatear (gofmt):**

```
main.go
integration_test/attributes/setup_test.go
integration_test/brands/setup_test.go
integration_test/global_catalog/setup_test.go
integration_test/quickstart/quickstart_integration_test.go
integration_test/taxonomy/setup_test.go
src/attribute/application/usecase/list_marketplace_attributes_by_criteria.go
src/attribute/domain/entity/attribute.go
src/attribute/domain/exception/errors.go
src/attribute/domain/port/attribute_repository.go
src/attribute/domain/port/marketplace_attribute_repository.go
src/attribute/infrastructure/config/attribute_module_config.go
src/attribute/infrastructure/controller/http_handler.go
src/attribute/infrastructure/controller/marketplace_attribute_handler.go
src/attribute/infrastructure/persistence/repository/attribute_postgres_repository.go
src/attribute/infrastructure/persistence/repository/marketplace_attribute_postgres_repository.go
src/brand/application/mapper/brand_mapper.go
src/brand/application/usecase/list_brands_by_criteria.go
src/brand/application/usecase/update_marketplace_brand.go
src/brand/domain/exception/errors.go
(+ más)
```

**Corrección:** `gofmt -w .` desde la raíz del proyecto (no toca lógica, es puramente cosmético).

---

## 4. Phase 2 — Auditoría hexagonal / DDD

### Tabla de hallazgos

| # | Severidad | Archivo | Violación | Por qué rompe hex/DDD | Corrección sugerida |
|---|-----------|---------|-----------|----------------------|---------------------|
| 1 | **CRITICAL** | `src/quickstart/domain/port/apply_template_repository.go` | Importa `"database/sql"` y `saas-mt-pim-service/src/shared/infrastructure/database` | La capa `domain` importa un driver de BD y un paquete de `infrastructure`. Viola la regla de dependencia: el dominio no puede depender de la infraestructura. Además, los métodos del port reciben `database.Executor` y retornan/aceptan `sql.NullString` — el contrato del puerto está contaminado con tipos de infraestructura. | Ver plan de refactor §4.1 |
| 2 | **CRITICAL** | `src/quickstart/application/usecase/apply_template_usecase.go` | Importa `"database/sql"` y `saas-mt-pim-service/src/shared/infrastructure/database` | La capa `application` importa infraestructura directamente, violando la regla de dependencia. | Ver plan de refactor §4.1 |
| 3 | **HIGH** | `src/product/global_catalog/application/usecase/delete_global_product.go` | El usecase mantiene `*sql.DB` y ejecuta SQL crudo (`SELECT COUNT(*) FROM tenant_global_product_links`) | Un caso de uso de `application` no debe conocer `sql.DB`. El query de conteo de tenant links es lógica de persistencia que debe estar en el repositorio. | Añadir método `CountTenantLinks(ctx, id string) (int, error)` al port `GlobalProductRepository`; implementarlo en el repositorio Postgres; inyectar solo la interfaz en el usecase. |
| 4 | **HIGH** | `src/s2s/usecase/` (ambos archivos) | Módulo completamente plano: sin dominio, sin ports. Los "usecases" mantienen `*sql.DB` y ejecutan CTEs multi-párrafo directamente. | No hay separación entre lógica de negocio y acceso a datos. Si la lógica crece o se necesita testear, no hay punto de inyección. | Aceptable para un módulo interno/admin si se documenta como tal. Para la variante correcta: crear `src/s2s/domain/port/template_repository.go` con las interfaces de query; mover SQL a `src/s2s/infrastructure/persistence/`. |
| 5 | **LOW** | `src/attribute/domain/entity/marketplace_attribute.go`, `src/brand/domain/entity/marketplace_brand.go` y otros | Tags `json:"..."` en entidades de dominio | Acopla el dominio al formato de serialización HTTP. Si el formato JSON cambia, el dominio cambia. | Solo nota: bajo impacto si los DTOs de `application/response` mapean desde estas entidades antes de serializar. Verificar que los handlers nunca serialicen entidades de dominio directamente. |

### Plan de refactor §4.1 (violaciones CRITICAL — requiere aprobación antes de aplicar)

**Objetivo:** eliminar `database/sql` e `infrastructure/database` del módulo `quickstart/domain` y `quickstart/application`.

**Paso 1 — Mover `Executor` a shared domain:**
```
# Nuevo archivo:
src/shared/domain/port/executor.go

package port

import (
    "context"
    "database/sql"
)

// Executor abstrae una transacción o conexión de base de datos.
// Definido en dominio compartido para que los ports de dominio puedan usarlo
// sin importar shared/infrastructure/database.
type Executor interface {
    ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
    QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
    QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
```

> Nota: esta interfaz necesita importar `database/sql` para los tipos de retorno — es la única concesión aceptable. Alternativamente, usar `interface{ ExecContext(...) error }` y wrappers propios para evitar cualquier dep de `database/sql` en el dominio.

**Paso 2 — En `apply_template_repository.go`:**
```diff
- import "database/sql"
- import "saas-mt-pim-service/src/shared/infrastructure/database"
+ import sharedPort "saas-mt-pim-service/src/shared/domain/port"

# Reemplazar:
- CreateTenantCategoriesLegacy(ctx, exec database.Executor, ...) 
+ CreateTenantCategoriesLegacy(ctx, exec sharedPort.Executor, ...)

# sql.NullString → *string en firmas del port
- ResolveTenantCategory(...) (sql.NullString, sql.NullString, error)
+ ResolveTenantCategory(...) (*string, *string, error)
```

**Paso 3 — `apply_template_usecase.go`:** misma sustitución de imports.

**Paso 4 — Adaptadores de infraestructura:** `database.Executor` sigue existiendo; añadir `var _ sharedPort.Executor = (*database.Executor)(nil)` para verificar en compile time que cumple la interfaz.

---

## 5. Phase 3 — Coverage

### Números actuales

```
total: 0.4%
```

**¿Por qué tan bajo?** Los tests viven en `test/` (paquete separado, black-box). El comando `go test ./... -coverprofile=coverage.out` solo instrumenta el paquete que *contiene* el test, no el paquete bajo prueba. El resultado es que prácticamente todo el código de `src/` aparece con 0%.

**Paquete con cobertura real medida:**

| Paquete | Cobertura |
|---------|-----------|
| `src/product/quickstart/application/usecase/backfill_tenant_images_usecase.go` (Execute) | 64% |
| `src/brand/domain/entity` (via test/brand/domain/entity) | 54.4% |

### Para medir correctamente

```bash
GOWORK=off go test ./... -coverprofile=coverage.out -covermode=atomic -coverpkg=./src/...
go tool cover -func=coverage.out | tail -5
```

La flag `-coverpkg=./src/...` hace que el instrumentador cubra todos los paquetes bajo `src/` independientemente de dónde viva el test.

### Script de enforcement (oferta — no aplicado)

```bash
#!/usr/bin/env bash
# scripts/check-coverage.sh
set -euo pipefail
THRESHOLD=${1:-80}
GOWORK=off go test ./... -coverprofile=coverage.out -covermode=atomic -coverpkg=./src/...
total=$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","",$3); print $3}')
echo "coverage total: ${total}% (umbral ${THRESHOLD}%)"
awk -v t="$total" -v th="$THRESHOLD" 'BEGIN { exit (t+0 < th+0) }'
```

Ejecutar con aprobación del usuario antes de crear el archivo.

---

## 6. Phase 4 — OpenAPI

**Spec actual:** `api-docs/openapi.yaml` (v2.3.0, 3143 líneas). Bien estructurada con tags por bounded context, schemas de error compartidos, y security schemes JWT.

**Puerto:** spec lista `http://localhost:8080/api/v1`. El código usa `getEnv("PORT", "8080")` y `.env.example` confirma `PORT=8080`. CLAUDE.md indica 8090 — probablemente el env de producción/Docker sobreescribe. La spec es correcta para desarrollo local.

**Rutas en código sin confirmar en spec (a verificar manualmente):**

| Ruta | Fuente |
|------|--------|
| `POST /api/v1/wizard/start` | `main.go:313` (simpleWizardHandler) |
| `GET /api/v1/wizard/status` | ídem |
| `PUT /api/v1/wizard/step` | ídem |
| `GET /api/v1/wizard/template/:businessTypeId` | ídem |
| `GET /api/v1/wizard/template/:businessTypeId/:section` | ídem |
| `DELETE /api/v1/wizard/reset` | ídem (temporal) |
| `GET /api/v1/quickstart/backfill-tenant-images` | `main.go:93` (excluded from auth) |
| `GET /api/v1/quickstart/backfill-all-tenant-images` | ídem |
| `POST /api/v1/s2s/refresh-template-products` | `src/s2s/controller/internal_handler.go` |
| `GET /api/v1/s2s/business-types/:slug/template-status` | ídem |
| `GET /api/v1/marketplace/categories/:id` | `main.go:459` (GET/DELETE añadidos) |
| `DELETE /api/v1/marketplace/categories/:id` | ídem |

**Validación redocly:** no ejecutada (npx no disponible en este entorno). Ejecutar localmente: `npx -y @redocly/cli lint api-docs/openapi.yaml`.

---

## 7. Phase 5 — Postman

`postman_collection.json` existe en raíz. `postman_environment.json` también.

No auditados en detalle — la skill requiere ejecutar la colección contra un servidor activo (Phase 6) que se omitió por la inviabilidad de SQLite.

---

## 8. Phase 6 — e2e Newman

**SKIP.** Motivo: el proyecto usa SQL Postgres-específico en múltiples lugares:
- `s2s/usecase`: CTEs complejas con `jsonb_array_elements`, `RETURNING`, correlaciones
- `quickstart/domain/port`: tipos `sql.NullString` implican queries con columnas nullable de Postgres
- Global catalog: queries con `JSONB` probables

**Alternativa recomendada:** los tests de integración ya usan Testcontainers con Postgres real (`testcontainers-go/modules/postgres`). Extender esa suite en lugar de crear un path SQLite.

---

## 9. Deuda pendiente

| Item | Estado | Motivo |
|------|--------|--------|
| Corregir violaciones CRITICAL (#1, #2) | **Pendiente aprobación** | Requiere nuevo archivo `src/shared/domain/port/executor.go` y cambios en `quickstart/domain/port/` + `quickstart/application/usecase/` |
| Corregir violación HIGH #3 | **Pendiente aprobación** | Añadir método al port `GlobalProductRepository` y refactorizar usecase |
| Refactorizar módulo `s2s` (#4) | **Decisión pendiente** | Bajo riesgo si se acepta como módulo interno plano; alto valor si se prevé testabilidad |
| `gofmt -w .` | Puede aplicarse sin aprobación (solo formato) | Sin efecto en lógica |
| Fix go.work (`go work use`) | Requiere contexto del workspace completo | No ejecutable solo desde `pim-service/` |
| `-coverpkg=./src/...` en CI | Oferta — no aplicado | Necesita modificar workflow de GitHub Actions |
| Script `scripts/check-coverage.sh` | Oferta — no creado | Pendiente confirmar umbral deseado |
| Rutas wizard/s2s/backfill en OpenAPI | Pendiente | Requiere leer controllers en detalle para completar schemas |
| Validación redocly | Pendiente entorno | `npx` no disponible en sesión actual |
| Tags `json:` en entidades dominio (#5) | Nota solo | Bajo impacto — no propuesto fix sin revisión de cada serialización |
