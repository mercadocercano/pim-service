---
adr: ADR-003
status: accepted
skills:
  implement:
    - dev/hexagonal-go
  verify:
    - dev/code-reviewer
    - dev/postgres-data-modeling
---
# ADR-003: Eliminación de MongoDB — PostgreSQL-only

**Estado**: Aceptado  
**Fecha**: 2026-02-01  
**Contexto**: La versión inicial del servicio usaba MongoDB para almacenar datos de productos y catálogo, y PostgreSQL para datos relacionales. Mantener dos bases de datos implicaba dos drivers, dos estrategias de migración y dos entornos en el laboratorio. Los datos de productos resultaron modelarse bien en tablas relacionales con columnas JSONB para atributos dinámicos.

## Decisión

Eliminamos MongoDB por completo. El servicio usa exclusivamente PostgreSQL 15 (con extensión `pgvector` disponible en el laboratorio). Los atributos dinámicos y templates se almacenan en columnas `JSONB`.

## Alternativas consideradas

| Opción | Por qué no |
|--------|-----------|
| Mantener MongoDB para catálogo/templates | Complejidad operativa sin beneficio real; PostgreSQL JSONB cubre los casos de uso de documentos |
| Migrar a un motor de búsqueda (Elasticsearch) | Fuera del scope actual; el catálogo no requiere full-text search avanzado |

## Consecuencias

**Positivas**: Stack simplificado — una sola BD, un solo driver (`database/sql`), migraciones SQL estándar.  
**Negativas / trade-offs**: Consultas sobre campos JSONB son menos expresivas que las de MongoDB; índices GIN necesarios para búsquedas en arrays.  
**Neutral**: Los seeds y fixtures pasaron de documentos BSON a scripts SQL.
