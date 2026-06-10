# Documentación — PIM Service

## Architecture Decision Records

| ADR | Título | Estado | Fecha |
|-----|--------|--------|-------|
| [ADR-001](adr/ADR-001-arquitectura-hexagonal-modular.md) | Arquitectura Hexagonal Modular | Aceptado | 2026-02-01 |
| [ADR-002](adr/ADR-002-metrics-recorder-inyectado.md) | MetricsRecorder como Dependencia Inyectada | Aceptado | 2026-06-09 |
| [ADR-003](adr/ADR-003-eliminacion-mongodb-postgresql-only.md) | Eliminación de MongoDB — PostgreSQL-only | Aceptado | 2026-02-01 |
| [ADR-004](adr/ADR-004-migracion-go-shared.md) | Extracción de Código Compartido a go-shared | Aceptado | 2026-06-09 |

## Arquitectura

- [Visión general](architecture/overview.md) — stack, capas hexagonales, módulos activos, flujos principales

## Setup

- [Deployment](setup/deployment.md) — variables de entorno, desarrollo local, laboratorio Docker, build de producción

## Guías

- [Testing](guides/testing.md) — tests unitarios, integración con TestContainers, scripts bash
- [Tests de integración](guides/integration-tests.md) — scripts bash contra el servicio levantado
- [Módulo category-attribute](guides/module-category-attribute.md) — endpoints, filtros avanzados, operadores

## Archivo histórico

- [Auditorías 2026-02](archive/README.md) — reportes de auditoría de endpoints y arquitectura
