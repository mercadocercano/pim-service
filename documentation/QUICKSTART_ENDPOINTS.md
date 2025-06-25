# Endpoints del Módulo Quickstart

## Descripción General
El módulo Quickstart proporciona funcionalidades para la configuración inicial rápida de productos en el tenant, incluyendo importación desde templates del catálogo global y seguimiento del progreso.

## Endpoints Implementados

### 1. Crear Producto desde Template
**POST** `/api/v1/quickstart/products/from-template`

Crea un producto del tenant basado en un template del catálogo global.

**Headers requeridos:**
- `X-Tenant-ID`: ID del tenant
- `Content-Type`: application/json

**Request Body:**
```json
{
  "template_id": "template-uuid",
  "tenant_id": "tenant-uuid",
  "initial_status": "draft",
  "custom_name": "Producto personalizado",
  "custom_description": "Descripción personalizada"
}
```

**Response 201:**
```json
{
  "product_id": "product-uuid",
  "template_id": "template-uuid",
  "tenant_id": "tenant-uuid",
  "product_name": "Producto personalizado",
  "status": "draft",
  "created_at": "2024-01-15T10:30:00Z"
}
```

---

### 2. Importación Masiva desde Tipo de Negocio
**POST** `/api/v1/quickstart/products/import-from-business-type`

Importa múltiples productos desde el catálogo global basado en el tipo de negocio del tenant.

**Headers requeridos:**
- `X-Tenant-ID`: ID del tenant
- `Content-Type`: application/json

**Request Body:**
```json
{
  "business_type_id": "business-type-uuid",
  "category_ids": ["cat-1", "cat-2"],
  "product_ids": ["prod-1", "prod-2"],
  "import_all": false,
  "initial_status": "draft"
}
```

**Response 201:**
```json
{
  "tenant_id": "tenant-uuid",
  "business_type_id": "business-type-uuid",
  "imported_products": [
    {
      "product_id": "prod-001",
      "template_id": "template-iphone-15",
      "product_name": "iPhone 15 Pro - tenant-uuid",
      "status": "draft",
      "category_name": "Smartphones"
    }
  ],
  "failed_imports": [
    {
      "template_id": "template-invalid",
      "error": "Template no encontrado",
      "reason": "El template especificado no existe en el catálogo global"
    }
  ],
  "summary": {
    "total_attempted": 3,
    "total_success": 2,
    "total_failed": 1,
    "success_rate_percentage": 66
  }
}
```

---

### 3. Obtener Progreso del Quickstart
**GET** `/api/v1/quickstart/progress`

Retorna el progreso del proceso de quickstart para un tenant.

**Headers requeridos:**
- `X-Tenant-ID`: ID del tenant

**Response 200:**
```json
{
  "tenant_id": "tenant-uuid",
  "quickstart_steps": {
    "business_type_selected": true,
    "categories_imported": true,
    "products_imported": true,
    "products_configured": true,
    "products_activated": false,
    "setup_completed": false
  },
  "completion_percentage": 66,
  "next_step": "activate_products",
  "recommendations": [
    "Tienes 2 productos en borrador. Configúralos para activarlos",
    "Activa al menos 3 productos para completar la configuración inicial"
  ],
  "products_stats": {
    "total_products": 5,
    "draft_products": 2,
    "pending_products": 1,
    "active_products": 2,
    "inactive_products": 0
  }
}
```

## Flujo de Negocio Recomendado

### 1. Configuración Inicial
```
1. Tenant registra → selecciona business_type
2. GET /global-catalog/business-types/{id}/products → Ve productos sugeridos
3. POST /quickstart/products/import-from-business-type → Importa productos masivamente
```

### 2. Configuración Individual
```
4. POST /quickstart/products/from-template → Crea productos específicos
5. PATCH /products/{id}/status → Transiciones: draft → pending → active
6. GET /quickstart/progress → Monitorea progreso
```

### 3. Finalización
```
7. GET /products → Ve productos activos listos para marketplace
```

## Estados de Productos en Quickstart

| Estado | Descripción | Siguiente Estado |
|--------|-------------|------------------|
| `draft` | Importado desde template, pendiente configuración | `pending`, `active`, `deleted` |
| `pending` | Configurado pero pendiente precios/stock | `draft`, `active`, `deleted` |
| `active` | Listo para vender | `inactive`, `discontinued`, `deleted` |
| `inactive` | Temporalmente deshabilitado | `active`, `discontinued`, `deleted` |
| `discontinued` | Descontinuado | `deleted` |
| `deleted` | Eliminado (soft delete) | - |

## Criterios de Progreso

### Pasos del Quickstart
1. **business_type_selected**: Tenant seleccionó tipo de negocio
2. **categories_imported**: Categorías importadas desde global catalog
3. **products_imported**: Al menos 1 producto importado (`total_products > 0`)
4. **products_configured**: Al menos 1 producto en `pending` o `active`
5. **products_activated**: Al menos 1 producto en `active`
6. **setup_completed**: Al menos 3 productos en `active`

### Recomendaciones Automáticas
- Si `draft_products > 0`: "Configura productos en borrador"
- Si `pending_products > 0`: "Agrega precios y stock"
- Si `active_products < 3`: "Activa al menos 3 productos"
- Si completado: "¡Configuración inicial finalizada!"

## Casos de Uso Implementados

### CreateFromTemplateUseCase
- **Propósito**: Crear producto individual desde template
- **Validaciones**: Template existe, tenant válido
- **Estado inicial**: Configurable (default: "draft")

### ImportFromBusinessTypeUseCase
- **Propósito**: Importación masiva basada en business_type
- **Filtros**: Por categorías, productos específicos, o todo
- **Reporte**: Éxitos, fallos y estadísticas

### GetQuickstartProgressUseCase
- **Propósito**: Evaluar progreso y generar recomendaciones
- **Métricas**: Porcentaje completitud, siguiente paso
- **Estadísticas**: Productos por estado

## Arquitectura

```
QuickstartController
    ├── CreateFromTemplateUseCase
    ├── ImportFromBusinessTypeUseCase
    └── GetQuickstartProgressUseCase
        └── ProductRepository (tenant)
```

## Integración con Otros Módulos

- **Global Catalog**: Fuente de templates y productos de referencia
- **Product Tenant**: Repositorio de productos del tenant
- **Business Types**: Clasificación para importación masiva
- **Categories**: Filtrado por categorías en importación 