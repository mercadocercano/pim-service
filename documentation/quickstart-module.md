# 🚀 Módulo Quickstart - PIM

## Descripción General

El módulo **Quickstart** del PIM proporciona un sistema de configuración rápida y automatizada para nuevos tenants, permitiendo la inicialización de catálogos de productos predefinidos basados en el tipo de negocio del cliente.

## Objetivos

- **Acelerar el time-to-value** para nuevos clientes
- **Reducir la fricción** en el proceso de onboarding
- **Proporcionar configuraciones optimizadas** según el tipo de negocio
- **Facilitar la adopción** del sistema PIM

## Arquitectura del Módulo

### Estructura de Directorios

```
pim/src/quickstart/
├── domain/
│   ├── entity/
│   │   ├── business_type.go
│   │   ├── quickstart_template.go
│   │   └── tenant_quickstart_history.go
│   ├── port/
│   │   ├── quickstart_repository.go
│   │   └── quickstart_service.go
│   ├── service/
│   │   └── quickstart_service.go
│   ├── value_object/
│   │   └── setup_data.go
│   └── exception/
│       └── quickstart_exceptions.go
├── application/
│   ├── usecase/
│   │   ├── get_business_types_usecase.go
│   │   ├── get_categories_by_business_type_usecase.go
│   │   ├── get_attributes_by_business_type_usecase.go
│   │   ├── get_variants_by_business_type_usecase.go
│   │   ├── get_products_by_business_type_usecase.go
│   │   └── setup_tenant_usecase.go
│   ├── request/
│   │   └── setup_tenant_request.go
│   └── response/
│       ├── business_type_response.go
│       └── setup_response.go
├── infrastructure/
│   ├── controller/
│   │   └── quickstart_controller.go
│   └── loader/
│       └── yaml_data_loader.go
└── data/
    ├── business-types.yaml
    ├── categories/
    │   └── retail.yaml
    ├── attributes/
    │   └── retail.yaml
    ├── variants/
    │   └── retail.yaml
    ├── products/
    │   └── retail.yaml
    └── brands/
        └── retail.yaml
```

## Entidades del Dominio

### BusinessType
Representa un tipo de negocio con configuraciones predefinidas.

```go
type BusinessType struct {
    ID          string
    Name        string
    Description string
    Icon        string
}
```

### QuickstartTemplate
Contiene toda la configuración predefinida para un tipo de negocio específico.

```go
type QuickstartTemplate struct {
    ID           string
    BusinessType string
    Categories   []Category
    Attributes   []Attribute
    Variants     []Variant
    Products     []Product
    Brands       []Brand
}
```

### TenantQuickstartHistory
Registra el historial de configuraciones quickstart aplicadas a un tenant.

```go
type TenantQuickstartHistory struct {
    ID           string
    TenantID     string
    BusinessType string
    Status       SetupStatus
    SetupData    SetupData
    CreatedAt    time.Time
    CompletedAt  *time.Time
}
```

## Casos de Uso

### 1. GetBusinessTypesUseCase
Obtiene la lista de tipos de negocio disponibles.

**Endpoint:** `GET /api/quickstart/business-types`

**Respuesta:**
```json
{
  "business_types": [
    {
      "id": "retail",
      "name": "Comercio Minorista",
      "description": "Tiendas de venta al por menor, supermercados, minimarkets",
      "icon": "store"
    }
  ]
}
```

### 2. GetCategoriesByBusinessTypeUseCase
Obtiene las categorías predefinidas para un tipo de negocio.

**Endpoint:** `GET /api/quickstart/categories/{businessType}`

### 3. GetAttributesByBusinessTypeUseCase
Obtiene los atributos predefinidos para un tipo de negocio.

**Endpoint:** `GET /api/quickstart/attributes/{businessType}`

### 4. GetVariantsByBusinessTypeUseCase
Obtiene las configuraciones de variantes predefinidas.

**Endpoint:** `GET /api/quickstart/variants/{businessType}`

### 5. GetProductsByBusinessTypeUseCase
Obtiene los productos de ejemplo predefinidos.

**Endpoint:** `GET /api/quickstart/products/{businessType}`

### 6. SetupTenantUseCase
Aplica la configuración quickstart completa a un tenant.

**Endpoint:** `POST /api/quickstart/setup`

**Request:**
```json
{
  "business_type": "retail",
  "include_categories": true,
  "include_attributes": true,
  "include_variants": true,
  "include_products": false,
  "include_brands": true
}
```

## Tipos de Negocio Soportados

### 1. Retail (Comercio Minorista)
- **Categorías:** 5 categorías principales con 3 niveles de jerarquía
- **Atributos:** 20 atributos específicos para retail
- **Variantes:** 10 configuraciones de variantes
- **Productos:** 20 productos de ejemplo
- **Marcas:** 40+ marcas reconocidas

### 2. Food & Beverage (En desarrollo)
- Configuraciones específicas para restaurantes y distribuidores de alimentos

### 3. Fashion (En desarrollo)
- Configuraciones para tiendas de moda y boutiques

### 4. Electronics (En desarrollo)
- Configuraciones para tiendas de electrónicos

### 5. Automotive (En desarrollo)
- Configuraciones para repuestos automotrices

### 6. Sports & Fitness (En desarrollo)
- Configuraciones para equipamiento deportivo

### 7. Health & Pharmacy (En desarrollo)
- Configuraciones para farmacias y productos de salud

### 8. Books & Media (En desarrollo)
- Configuraciones para librerías y medios

### 9. Home & Construction (En desarrollo)
- Configuraciones para ferreterías y construcción

### 10. Beauty & Cosmetics (En desarrollo)
- Configuraciones para productos de belleza

### 11. Toys & Games (En desarrollo)
- Configuraciones para jugueterías

### 12. Pet Supplies (En desarrollo)
- Configuraciones para productos de mascotas

### 13. Office Supplies (En desarrollo)
- Configuraciones para papelerías y oficina

### 14. Jewelry & Accessories (En desarrollo)
- Configuraciones para joyerías

## Configuración de Datos

### Estructura de Archivos YAML

#### business-types.yaml
Define los tipos de negocio disponibles:

```yaml
business_types:
  - id: "retail"
    name: "Comercio Minorista"
    description: "Tiendas de venta al por menor, supermercados, minimarkets"
    icon: "store"
```

#### categories/{business_type}.yaml
Define la estructura de categorías:

```yaml
business_type: "retail"
categories:
  - id: "home-garden"
    name: "Hogar y Jardín"
    description: "Productos para el hogar y jardinería"
    parent_id: null
    subcategories: [...]
```

#### attributes/{business_type}.yaml
Define los atributos disponibles:

```yaml
business_type: "retail"
attributes:
  - id: "color"
    name: "Color"
    type: "select"
    required: false
    values: ["Negro", "Blanco", "Gris", ...]
```

#### variants/{business_type}.yaml
Define las configuraciones de variantes:

```yaml
business_type: "retail"
variants:
  - id: "color-size"
    name: "Color y Talla"
    attributes: ["color", "size"]
    combinations: [...]
```

#### products/{business_type}.yaml
Define productos de ejemplo:

```yaml
business_type: "retail"
products:
  - id: "sofa-3-plazas"
    name: "Sofá de 3 Plazas"
    category: "living-room-furniture"
    base_price: 89999
    currency: "ARS"
    attributes: {...}
```

#### brands/{business_type}.yaml
Define marcas reconocidas:

```yaml
business_type: "retail"
brands:
  - id: "ikea"
    name: "IKEA"
    description: "Muebles y artículos para el hogar"
    country_origin: "Suecia"
```

## Sistema de Caché

El módulo implementa un sistema de caché thread-safe para optimizar el rendimiento:

```go
type YAMLDataLoader struct {
    cache map[string]interface{}
    mutex sync.RWMutex
}
```

- **Lectura concurrente:** Múltiples goroutines pueden leer simultáneamente
- **Escritura exclusiva:** Solo una goroutine puede escribir a la vez
- **Invalidación automática:** El caché se invalida cuando se detectan cambios

## Manejo de Errores

### Excepciones del Dominio

- `BusinessTypeNotFoundError`
- `TemplateNotFoundError`
- `InvalidBusinessTypeError`
- `SetupAlreadyExistsError`
- `SetupNotFoundError`
- `InvalidSetupDataError`
- `SetupInProgressError`
- `SetupFailedError`
- `TenantNotFoundError`
- `DataLoadingError`

### Códigos de Estado HTTP

- `200 OK` - Operación exitosa
- `400 Bad Request` - Datos inválidos
- `404 Not Found` - Recurso no encontrado
- `409 Conflict` - Configuración ya existe
- `500 Internal Server Error` - Error interno

## Integración con Otros Módulos

### Dependencias

- **Category Module:** Para crear categorías
- **Attribute Module:** Para crear atributos
- **Product Module:** Para crear productos
- **Brand Module:** Para crear marcas
- **Tenant Module:** Para validar tenants

### Eventos de Dominio

- `QuickstartSetupStartedEvent`
- `QuickstartSetupCompletedEvent`
- `QuickstartSetupFailedEvent`
- `BusinessTypeSelectedEvent`

## Configuración y Deployment

### Variables de Entorno

```bash
QUICKSTART_DATA_PATH=/path/to/quickstart/data
QUICKSTART_CACHE_TTL=3600
QUICKSTART_ENABLE_CACHE=true
```

### Inicialización

```go
// Cargar datos al inicio de la aplicación
loader := NewYAMLDataLoader(dataPath)
service := NewQuickstartService(loader, repository)
```

## Testing

### Datos de Prueba

Se incluyen datos de prueba completos para el tipo "retail":
- 5 categorías principales
- 20 atributos
- 10 configuraciones de variantes
- 20 productos de ejemplo
- 40+ marcas

### Casos de Prueba

- Carga de tipos de negocio
- Obtención de configuraciones por tipo
- Aplicación de setup completo
- Manejo de errores
- Validación de datos

## Roadmap

### Fase 1 (Completada)
- ✅ Implementación del tipo "retail"
- ✅ API endpoints básicos
- ✅ Sistema de caché
- ✅ Documentación

### Fase 2 (En desarrollo)
- 🔄 Implementación de los 13 tipos restantes
- 🔄 Interfaz de administración
- 🔄 Métricas y analytics

### Fase 3 (Planificada)
- 📋 Configuraciones personalizables
- 📋 Templates dinámicos
- 📋 Importación/exportación de configuraciones
- 📋 Versionado de templates

## Contribución

Para agregar un nuevo tipo de negocio:

1. Crear archivos YAML en `data/{categories,attributes,variants,products,brands}/`
2. Agregar entrada en `business-types.yaml`
3. Crear tests correspondientes
4. Actualizar documentación

## Soporte

Para soporte técnico o consultas sobre el módulo quickstart:
- Documentación: `/pim/documentation/`
- Issues: Repositorio del proyecto
- Contacto: Equipo de desarrollo PIM 