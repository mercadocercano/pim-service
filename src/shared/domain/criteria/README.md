# Sistema de Criteria

Sistema compartido para filtrado, ordenamiento y paginación en los microservicios.

## Operadores Soportados

### Operadores Básicos
- `OpEqual` (`=`): Igualdad exacta
- `OpNotEqual` (`!=`): Diferente de
- `OpGreaterThan` (`>`): Mayor que
- `OpGreaterThanOrEqual` (`>=`): Mayor o igual que
- `OpLessThan` (`<`): Menor que
- `OpLessThanOrEqual` (`<=`): Menor o igual que

### Operadores de Texto
- `OpLike` (`LIKE`): Búsqueda de texto con wildcards

### Operadores de Arrays
- `OpIn` (`IN`): El valor está dentro de una lista
- `OpArrayContains` (`ARRAY_CONTAINS`): El array contiene un valor específico (PostgreSQL)

### Operadores de Nulidad
- `OpIsNull` (`NULL`): El campo es nulo
- `OpIsNotNull` (`NOT NULL`): El campo no es nulo

## Métodos Helper del CriteriaBuilder

```go
// Filtros básicos
AddEqualFilter(field, value)
AddNotEqualFilter(field, value)
AddLikeFilter(field, value)
AddGreaterThanFilter(field, value)
AddLessThanFilter(field, value)

// Filtros especializados
AddUUIDFilter(field, value)        // Valida formato UUID
AddBoolFilter(field, value)        // Convierte string/bool
AddInFilter(field, []values)       // Array de valores
AddArrayContainsFilter(field, value) // Nuevo: array contiene valor

// Ordenamiento y paginación
SetOrder(field, direction)
SetPagination(page, pageSize)
```

## Ejemplo de Uso

```go
// Usar el builder
criteriaBuilder := criteria.NewCriteriaBuilder()
crit := criteriaBuilder.
    AddUUIDFilter("tenant_id", "123").
    AddEqualFilter("status", "active").
    AddArrayContainsFilter("tags", "importante").
    SetOrder("created_at", "DESC").
    SetPagination(1, 20).
    Build()

// Ejecutar consulta
result, err := repository.ListByCriteria(ctx, crit)
```

## Implementación PostgreSQL

El operador `ARRAY_CONTAINS` se traduce a:
```sql
-- Para arrays de texto
WHERE tags @> ARRAY['importante']

-- Para arrays de números
WHERE ids @> ARRAY[123]
```

Esto permite filtrar registros donde un campo array contiene un valor específico. 