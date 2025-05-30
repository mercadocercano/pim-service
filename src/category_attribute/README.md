# Módulo Category Attribute

Este módulo administra la relación entre categorías y atributos, permitiendo definir qué atributos están disponibles para cada categoría y sus valores permitidos.

## Estructura

El módulo sigue la arquitectura hexagonal con las siguientes capas:

### Domain (Dominio)
- **Entity**: `CategoryAttribute` - Entidad principal que representa la relación categoría-atributo
- **Port**: `CategoryAttributeRepository` - Interfaz del repositorio

### Application (Aplicación)
- **UseCase**: Casos de uso para CRUD de atributos de categoría
- **Request/Response**: DTOs para las peticiones y respuestas HTTP

### Infrastructure (Infraestructura)
- **Controller**: Manejador HTTP REST
- **Repository**: Implementación PostgreSQL del repositorio
- **Model/Mapper**: Modelos de base de datos y conversores
- **Criteria**: Builder para filtros avanzados y paginación

## Endpoints REST

### GET /api/v1/category-attributes
Obtiene atributos de categoría con soporte completo para **filtros, ordenamiento y paginación**.

**Query Parameters:**
- `tenant_id` (requerido): ID del tenant
- `category_id` (opcional): ID de la categoría específica
- `attribute_id` (opcional): ID del atributo específico
- `status` (opcional): Estado (active/inactive)
- `active` (opcional): true para solo activos
- `allowed_value` (opcional): Filtrar por atributos que contengan este valor específico
- `page` (opcional, default: 1): Número de página
- `page_size` (opcional, default: 10, max: 100): Elementos por página
- `sort_by` (opcional, default: created_at): Campo de ordenamiento
- `sort_dir` (opcional, default: DESC): Dirección ASC/DESC

**Headers:**
- `X-Tenant-ID`: ID del tenant (alternativo al query param)

**Respuesta:**
```json
{
  "items": [
    {
      "id": "uuid",
      "category_id": "uuid",
      "attribute_id": "uuid", 
      "allowed_values": ["valor1", "valor2"],
      "active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "total_count": 25,
  "page": 1,
  "page_size": 10,
  "total_pages": 3
}
```

**Ejemplos de uso:**
```bash
# Listar todos los atributos de un tenant con paginación
GET /api/v1/category-attributes?tenant_id=123&page=1&page_size=20

# Filtrar por categoría específica
GET /api/v1/category-attributes?tenant_id=123&category_id=456

# Solo atributos activos, ordenados por fecha
GET /api/v1/category-attributes?tenant_id=123&active=true&sort_by=created_at&sort_dir=ASC

# Filtrar por múltiples criterios
GET /api/v1/category-attributes?tenant_id=123&category_id=456&status=active&page=2

# Filtrar por atributos que contengan un valor específico en allowed_values
GET /api/v1/category-attributes?tenant_id=123&allowed_value=talla&page=1
```

### GET /api/v1/category-attributes/simple
Endpoint simple sin paginación para casos básicos.

**Query Parameters:**
- `tenant_id` (requerido): ID del tenant
- `category_id` (opcional): ID de la categoría específica

**Respuesta:**
```json
{
  "data": [
    {
      "id": "uuid",
      "category_id": "uuid",
      "attribute_id": "uuid", 
      "allowed_values": ["valor1", "valor2"],
      "active": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### POST /api/v1/category-attributes
Crea una nueva relación categoría-atributo.

**Headers:**
- `X-Tenant-ID` (requerido): ID del tenant

**Body:**
```json
{
  "category_id": "uuid",
  "attribute_id": "uuid",
  "allowed_values": ["valor1", "valor2"]
}
```

### PUT /api/v1/category-attributes/{id}
Actualiza los valores permitidos de un atributo de categoría.

**Headers:**
- `X-Tenant-ID` (requerido): ID del tenant

**Body:**
```json
{
  "allowed_values": ["nuevo_valor1", "nuevo_valor2"]
}
```

### DELETE /api/v1/category-attributes/{id}
Elimina una relación categoría-atributo.

**Headers:**
- `X-Tenant-ID` (requerido): ID del tenant

## Filtros Avanzados

### Campos Permitidos para Filtrado
- `id`: UUID del registro
- `tenant_id`: ID del tenant
- `category_id`: ID de la categoría
- `attribute_id`: ID del atributo
- `status`: Estado (active/inactive)
- `created_at`: Fecha de creación
- `updated_at`: Fecha de actualización

### Campos Permitidos para Ordenamiento
Los mismos campos de filtrado están disponibles para ordenamiento usando `sort_by`.

### Operadores Soportados
- `=` (igual)
- `!=` (diferente)
- `>`, `>=`, `<`, `<=` (comparaciones)
- `LIKE` (búsqueda de texto)
- `IN` (valores múltiples)
- `NULL`, `NOT NULL` (valores nulos)
- `ARRAY_CONTAINS` (verificar si array contiene valor específico)