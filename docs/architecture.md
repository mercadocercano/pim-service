# Arquitectura — PIM Service

## Stack

| Componente | Tecnología |
|------------|------------|
| Lenguaje | Go 1.25 |
| HTTP framework | Gin |
| Base de datos | PostgreSQL 15 (pgvector:pg15 en lab) |
| Shared lib | `github.com/mercadocercano/go-shared` |
| Metrics | Prometheus (PrometheusRecorder via go-shared) |
| Compresión | gzip middleware (go-shared) |
| Auth | JWT via `github.com/mercadocercano/middleware` (Kong gateway) |

MongoDB fue eliminado. El servicio es PostgreSQL-only.

---

## Capas (hexagonal)

```
┌─────────────────────────────────────────────────┐
│                  Infrastructure                  │
│  Controllers (Gin) │ Repositories (sql) │ Config │
├─────────────────────────────────────────────────┤
│                  Application                     │
│       Use Cases │ Mappers │ DTOs                 │
├─────────────────────────────────────────────────┤
│                    Domain                        │
│      Entities │ Value Objects │ Ports            │
└─────────────────────────────────────────────────┘
```

**Dirección de dependencias**: Infrastructure → Application → Domain (nunca al revés).

El dominio define interfaces (ports). La infraestructura las implementa. La aplicación las usa sin conocer la implementación.

---

## Módulos activos

```
main.go
├── api          → health check, Swagger UI
├── category     → categorías del tenant + marketplace categories
├── category_attribute → relación categoría-atributo
├── brand        → marcas por tenant + marketplace brands
├── attribute    → atributos base + marketplace attributes
├── product/tenant      → productos, variantes, importación CSV
├── product/quickstart  → importación desde catálogo global
├── product/global_catalog → catálogo global de referencia
├── quickstart   → wizard de templates (apply_template)
├── businesstype → tipos de negocio + templates
├── schema_validation → validación de schemas
├── overview     → dashboard de resumen
└── s2s          → endpoints server-to-server (template status, refresh)
```

---

## Shared — go-shared vs local

### Vive en `github.com/mercadocercano/go-shared` (libs/go-shared)

| Tipo | Package |
|------|---------|
| `Executor` (SQL) | `domain/port` |
| `FileImporter[T]`, `ImportResult[T]` | `domain/port` |
| `MetricsRecorder`, `MetricEvent` | `domain/port` |
| `NotificationGateway` | `domain/port` |
| `BaseCSVFileImporter[T]`, `RowParser[T]` | `infrastructure/adapters` |
| `PrometheusRecorder` | `infrastructure/metrics` |
| `NoopRecorder` | `infrastructure/metrics` |
| `HTTPNotificationGateway` | `infrastructure/notifications` |
| `NoopNotificationGateway` | `infrastructure/notifications` |
| Middleware gzip / force_gzip / decompress | `infrastructure/middleware` |
| `SharedConfig`, `SetupSharedMiddleware` | `infrastructure/config` |

### Permanece local en `src/shared/`

| Tipo | Package | Por qué es local |
|------|---------|-----------------|
| `ImportJobRepository` | `shared/domain/port` | Depende de `entity.ImportJob` local |
| `NotificationService` | `shared/domain/port` | Contrato específico (NotifyImportJobComplete) |
| `ImportJob` entity | `shared/domain/entity` | Entidad de negocio del servicio |
| `PostgresImportJobRepository` | `shared/infrastructure/persistence` | Implementación local |
| CORS middleware | `api/infrastructure/middleware` | Configuración específica del servicio |
| `executor.go` (compile-time checks) | `shared/infrastructure/database` | Solo verifica `*sql.DB` y `*sql.Tx` satisfacen el port |

---

## MetricsRecorder (ADR-002)

Las métricas se inyectan como dependencia, nunca como variables globales.

```go
// Puerto (go-shared)
type MetricsRecorder interface {
    Record(event MetricEvent)
}

// Uso en un use case
uc.metrics.Record(sharedport.MetricEvent{
    Name:  port.MetricImportOperation,  // constante en domain/port/metrics.go
    Kind:  sharedport.MetricKindCounter,
    Value: 1.0,
    Labels: map[string]string{"tenant_id": tenantID, "result": "success"},
})

// Wiring en config/main
metricsRecorder := sharedmetrics.NewPrometheusRecorder()
cfg := productConfig.NewProductConfig(db, metricsRecorder)
```

Constantes de métricas viven en el bounded context: `src/product/tenant/domain/port/metrics.go`.

---

## Flujo de quickstart

```
Tenant nuevo
    │
    ▼
GET /quickstart/templates          → lista templates disponibles
    │
    ▼
POST /quickstart/apply-template    → crea categorías, marcas, productos y atributos
    │                                 desde el JSONB del template
    ▼
GET /quickstart/status             → progreso del wizard
    │
    ▼
POST /quickstart/wizard/start      → inicia wizard paso a paso
    │
    ▼
POST /quickstart/wizard/step       → avanza en el wizard
    │
    ▼
POST /quickstart/wizard/complete   → finaliza el wizard
```

---

## Flujo de importación CSV

```
POST /products/import/csv
    │
    ├── ProductController.ImportProductsFromCSV
    │       ├── RecordFileSize (MetricsRecorder)
    │       └── ImportProductsFromCSVUseCase.Execute
    │               ├── ProductCSVFileImporter.Import  ← BaseCSVFileImporter (go-shared)
    │               │       └── ParseRow (validación, value objects)
    │               ├── ProductRepository.Save (por cada producto)
    │               └── Record métricas (operación, duración, errores)
    └── Response: ImportProductsCSVResponse
```

---

## API Gateway (Kong)

El servicio corre detrás de Kong en el laboratorio:

```
Cliente → Kong :8000 → pim-service :8090
                ↑
         Kong Admin :8001
         (config declarativa en infra/kong/kong.yml)
```

Rutas públicas (sin JWT):
- `/api/v1/health`
- `/api/v1/marketplace/*`
- `/api/v1/business-types*`
- `/api/v1/global-catalog*`
- `/api/v1/internal*`
- `/api/v1/s2s*`
- `/api/v1/quickstart/backfill-*`

---

## Migraciones

Las migraciones viven en `migrations/` y son **evolutivas, nunca destructivas**.
Los tests de integración dependen de ellas — no eliminar archivos de migración.

```bash
./scripts/migrate.sh       # aplica migraciones pendientes
```

Convención de nombres: `NNN_descripcion.sql` (`.up.sql` / `.down.sql` para las que tienen rollback).
