# PIM Service

**Product Information Management** para el ecosistema SaaS Marketplace multitenant.
Gestiona productos, categorías, atributos, marcas, catálogo global y sistema de quickstart.

**Puerto**: 8090 | **Stack**: Go 1.25 + Gin + PostgreSQL 15 | **Arquitectura**: Hexagonal modular

---

## Inicio rápido

```bash
# Dependencias
go mod download

# Migraciones
./scripts/migrate.sh

# Ejecutar (puerto 8090)
go run main.go
```

**Variables de entorno:**
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=pim_db

PORT=8090
GIN_MODE=debug
PROMETHEUS_ENABLED=false   # true para exponer /metrics
JWT_SECRET=<secret>
```

---

## Módulos

| Módulo | Ruta base | Descripción |
|--------|-----------|-------------|
| Products (tenant) | `/api/v1/products` | CRUD de productos del tenant, variantes, CSV import |
| Categories | `/api/v1/categories` | Jerarquía de categorías multitenant |
| Attributes | `/api/v1/attributes` | Atributos base y marketplace |
| Category Attributes | `/api/v1/category-attributes` | Relación atributos-categorías |
| Brands | `/api/v1/brands` | Marcas por tenant |
| Business Types | `/api/v1/business-types` | Tipos de negocio para quickstart |
| Global Catalog | `/api/v1/global-catalog` | Catálogo de referencia compartido |
| Marketplace Categories | `/api/v1/marketplace/categories` | Categorías globales del marketplace |
| Quickstart | `/api/v1/quickstart` | Templates de configuración inicial |
| S2S (internal) | `/api/v1/s2s` | Endpoints server-to-server |
| Overview | `/api/v1/overview` | Dashboard de resumen |
| Health | `/api/v1/health` | Estado del servicio |

API completa documentada en `api-docs/openapi.yaml` (v2.4.0, 53 paths).
Swagger UI disponible en `/api-docs` cuando el servicio corre.

---

## Tests

```bash
# Unitarios
go test ./...

# Integración (requiere Docker)
./run_tests.sh

# Con coverage
go test -coverprofile=coverage.out -coverpkg=./src/... ./...
go tool cover -html=coverage.out -o coverage.html
```

Los tests de integración usan [TestContainers](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/) — levantan PostgreSQL automáticamente.

---

## Estructura del proyecto

```
pim-service/
├── main.go                     # Entry point — wiring de todos los módulos
├── go.mod                      # Módulo: saas-mt-pim-service
├── src/
│   ├── api/                    # Health check, Swagger UI, middleware CORS
│   ├── attribute/              # Atributos de marketplace
│   ├── brand/                  # Marcas
│   ├── businesstype/           # Tipos de negocio y templates
│   ├── category/               # Categorías del tenant y marketplace
│   ├── category_attribute/     # Relación categoría-atributo
│   ├── overview/               # Dashboard resumen por tenant
│   ├── product/
│   │   ├── tenant/             # Productos del tenant (domain + app + infra)
│   │   ├── quickstart/         # Importación desde catálogo global
│   │   └── global_catalog/     # Catálogo global de productos referencia
│   ├── quickstart/             # Wizard de configuración inicial (templates)
│   ├── s2s/                    # Endpoints server-to-server
│   ├── schema_validation/      # Validación de schemas JSON
│   └── shared/                 # Infraestructura compartida (DB, config, adapters)
├── migrations/                 # Migraciones SQL evolutivas (nunca destructivas)
├── scripts/                    # migrate.sh, update-openapi.sh, wait-for-db.sh
├── api-docs/                   # openapi.yaml (spec canónica)
├── test/                       # Tests unitarios Go
├── test-integration/           # Tests de integración (bash + httpie/curl)
├── docs/                       # Documentación técnica detallada
└── templates/                  # Plantillas HTML (error pages, etc.)
```

---

## Arquitectura hexagonal

Cada módulo sigue el patrón:
```
src/<módulo>/
├── domain/
│   ├── entity/          # Entidades y value objects
│   ├── port/            # Interfaces (repositorios, servicios)
│   └── service/         # Servicios de dominio
├── application/
│   ├── usecase/         # Casos de uso
│   ├── mapper/          # Transformaciones domain ↔ DTO
│   ├── request/         # DTOs de entrada
│   └── response/        # DTOs de salida
└── infrastructure/
    ├── controller/      # Handlers HTTP (Gin)
    ├── persistence/     # Repositorios PostgreSQL
    ├── config/          # Wiring de dependencias
    └── adapters/        # Adaptadores (CSV importer, etc.)
```

**Regla estricta**: domain no importa infrastructure. application no importa infrastructure.
Código compartido reutilizable vive en `github.com/mercadocercano/go-shared` (workspace `libs/go-shared`).

---

## Estados de producto

```
draft → pending, active, deleted
pending → draft, active, deleted
active → inactive, discontinued, deleted
inactive → active, discontinued, deleted
discontinued → deleted
```

- **draft**: solo nombre requerido
- **pending**: nombre + descripción + categoría
- **active**: todos los campos + al menos 1 variante activa
- **deleted**: soft delete (columna `deleted_at`)

---

## Deployment

```bash
# Build
go build -ldflags="-s -w" -o pim-service main.go

# Docker (desde monorepo)
docker compose up -d pim-service

# CI/CD: push a main → .github/workflows/deploy-pim.yml
```

---

## Documentación

Ver [`docs/`](docs/README.md) para arquitectura, ADRs, setup y guías operativas.

- `api-docs/openapi.yaml` — especificación OpenAPI 3.1.0 (fuente de verdad de la API)
- `postman/` — colección Postman + runner Newman (`./postman/run.sh`)
- `docs/archive/` — auditorías y reportes históricos (2026-02)
