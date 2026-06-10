---
name: audit-go-hex-2026-06-09
description: Estado del pipeline go-hex-audit + limpieza MongoDB + go.work — sesiones 2026-06-09
metadata:
  type: project
---

## Estado post-sesión 2026-06-09

**Why:** go-hex-audit completo + limpieza masiva de MongoDB/seeds/docs + integración go.work.

### Completado ✅

| Item | Estado |
|------|--------|
| gofmt -w . | ✅ 0 archivos sin formatear |
| Phase 4 OpenAPI v2.4.0 | ✅ 13 rutas nuevas agregadas |
| MongoDB removido | ✅ código, deps, scripts, init |
| seeds/ borrado | ✅ 127 SQL files eliminados |
| documentation/ borrado | ✅ 9 markdown guides eliminados |
| Scripts MongoDB/seeds borrados | ✅ 9 scripts eliminados |
| go.work actualizado | ✅ pim-service agregado, go 1.25.0 |
| go.mod limpiado | ✅ mongo-driver eliminado, go 1.25.0 |
| CLAUDE.md actualizado | ✅ stack correcto |

### Rutas OpenAPI v2.4.0 agregadas (13 nuevas)
- `/marketplace/categories` GET + POST
- `/marketplace/categories/tree` GET
- `/marketplace/categories/validate-hierarchy` POST
- `/marketplace/categories/{id}` GET + PUT + DELETE
- `/wizard/status` GET
- `/wizard/start` POST
- `/wizard/step` PUT
- `/wizard/template/{businessTypeId}` GET
- `/wizard/template/{businessTypeId}/{section}` GET
- `/wizard/complete` POST
- `/wizard/reset` DELETE
- `/quickstart/backfill-tenant-images` POST
- `/quickstart/backfill-all-tenant-images` POST

### Violaciones AUDIT Phase 2 — estado

| # | Severidad | Estado |
|---|-----------|--------|
| 1 | CRITICAL quickstart/domain importa infrastructure | ✅ Corregido (sesión anterior) |
| 2 | CRITICAL quickstart/application importa infrastructure | ✅ Corregido (sesión anterior) |
| 3 | HIGH global_catalog delete usecase con *sql.DB | ✅ Corregido (sesión anterior) |
| 4 | HIGH s2s módulo plano sin ports | ✅ Refactorizado (sesión anterior) |
| 5 | LOW json tags en entidades dominio | Nota — no crítico |

### Pendiente
- Commit de todos los cambios acumulados (sessions anteriores + esta)
- `-coverpkg=./src/...` en CI (Phase 3 coverage fix)
- `scripts/check-coverage.sh` (oferta — no creado)
- Validación redocly del spec (npx no disponible)

### Archivos clave
- `api-docs/openapi.yaml` → v2.4.0, 53 paths
- `go.work` → go 1.25.0, incluye iam-service + pim-service
- `migrations/` → solo DDL (tests de integración dependen de esta carpeta)

**How to apply:** pim-service ya está en workspace, usar `go build ./services/pim-service/...` desde raíz del monorepo.
