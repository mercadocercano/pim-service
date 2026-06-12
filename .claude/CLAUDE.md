# CLAUDE.md - PIM Service

Servicio de Product Information Management del ecosistema SaaS multi-tenant.
Gestiona productos, categorías, atributos, marcas, catálogo global y sistema de quickstart.

**Puerto**: 8090 | **Stack**: Go + Gin + PostgreSQL | **Arquitectura**: Hexagonal modular

Hablame siempre en español.

## Comandos esenciales

```bash
go run main.go                        # Ejecutar (puerto 8090)
go test ./...                         # Tests unitarios
./run_tests.sh                        # Tests de integración
./scripts/migrate.sh                  # Migraciones
```

## Regla: MCP Go Generator primero

Antes de implementar código Go, consultar MCP:
```bash
analyzeUsecaseWorkflow --service_name="pim" --entity_name="product"
generateWorkflowRoadmap --service_name="pim"
generateComponentByStep --step_type="dto" --entity_name="product_variant"
```

## Contexto on-demand (cargar según necesidad)

| Archivo | Cuándo cargar |
|---------|---------------|
| `pim-service-management/api-endpoints.md` | Al trabajar con endpoints, quickstart, catálogo |
| `pim-service-management/architecture.md` | Al crear módulos, entidades, seeds |
| `pim-service-management/product-states.md` | Al trabajar con estados de productos o validación CSV |
| `pim-service-management/config.md` | Al configurar env vars, Docker, MongoDB |

## Reglas compartidas (cargar según contexto)

| Regla | Archivo |
|-------|---------|
| Arquitectura hexagonal | `ai-tools/rules/architecture.md` |
| Multi-tenancy | `ai-tools/rules/multi-tenant.md` |
| API Gateway / Kong | `ai-tools/rules/api-gateway.md` |
| Testing standards | `ai-tools/rules/testing-standards.md` |
| Formato respuesta API | `ai-tools/rules/api-response-format.md` |

## Memoria persistente (Engram)

Tenés acceso a memoria persistente entre sesiones vía las herramientas MCP de Engram (`mem_save`, `mem_search`, `mem_context`, etc.). Proyecto: **`mercado-cercano`** (memoria compartida con iam-service y el resto del ecosistema).

**Cuándo guardar** — sin esperar que te lo pidan:
- Al resolver un bug no trivial: síntoma, causa raíz, fix aplicado.
- Al tomar una decisión de diseño: qué se decidió y por qué.
- Al descubrir un patrón o convención del proyecto que no está documentada.
- Al completar una feature o refactor significativo: qué cambió y dónde.

**Cuándo buscar** — antes de empezar cualquier tarea:
- `mem_context` al inicio de sesión o tras una compaction para recuperar el estado anterior.
- `mem_search` cuando el usuario menciona algo que puede tener historial ("el bug de autenticación", "la migración de la semana pasada").

**Al cerrar sesión**: llamar `mem_session_summary` para dejar un resumen recuperable.

## Notas críticas

- Puerto **8090**, NO 8080
- Solo PostgreSQL — MongoDB fue removido completamente
- Migraciones evolutivas, nunca destructivas (tests de integración dependen de `migrations/`)
- pim-service está en el workspace `go.work` del monorepo (go 1.25.0)
