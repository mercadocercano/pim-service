# Arquitectura del Sistema PIM

## Visión General

El Sistema PIM está construido siguiendo los principios de **Arquitectura Hexagonal** (Ports & Adapters), **Domain-Driven Design (DDD)** y **principios SOLID**. Esta arquitectura garantiza:

- **Separación de responsabilidades**
- **Testabilidad**
- **Mantenibilidad**
- **Escalabilidad**
- **Independencia de frameworks**

## Arquitectura Hexagonal

### Capas de la Arquitectura

```
┌─────────────────────────────────────────────────────────────┐
│                    INFRASTRUCTURE                          │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │ Controllers │  │ Repositories│  │  Database   │        │
│  │   (HTTP)    │  │ (PostgreSQL)│  │ Migrations  │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────┬───────────────────────────────────────┘
                      │ Adapters
┌─────────────────────▼───────────────────────────────────────┐
│                   APPLICATION                              │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │ Use Cases   │  │    DTOs     │  │   Mappers   │        │
│  │ (Business   │  │ (Request/   │  │ (Domain <-> │        │
│  │ Workflows)  │  │ Response)   │  │    DTO)     │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────┬───────────────────────────────────────┘
                      │ Ports
┌─────────────────────▼───────────────────────────────────────┐
│                     DOMAIN                                 │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │  Entities   │  │ Value       │  │  Domain     │        │
│  │ (Aggregates)│  │ Objects     │  │  Services   │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
│                                                            │
│  ┌─────────────┐  ┌─────────────┐                         │
│  │   Ports     │  │ Exceptions  │                         │
│  │(Interfaces) │  │  (Domain)   │                         │
│  └─────────────┘  └─────────────┘                         │
└────────────────────────────────────────────────────────────┘
```

### 1. Capa de Dominio (Domain Layer)

**Responsabilidad**: Contiene la lógica de negocio pura, independiente de cualquier framework o tecnología externa.

#### Entidades (Entities)
```go
// Ejemplo: Product Entity
type Product struct {
    id          ProductID
    tenantID    TenantID
    name        string
    description *string
    sku         *ProductSKU
    status      ProductStatus
    category    *CategoryReference
    brand       *BrandReference
    variants    []*ProductVariant
    createdAt   time.Time
    updatedAt   time.Time
}

// Métodos de negocio
func (p *Product) Activate() error
func (p *Product) Deactivate() error
func (p *Product) AddVariant(...) (*ProductVariant, error)
func (p *Product) UpdateVariant(...) error
```

#### Value Objects
```go
// Ejemplo: ProductSKU Value Object
type ProductSKU struct {
    value string
}

func NewProductSKU(value string) (*ProductSKU, error) {
    if len(value) < 3 || len(value) > 50 {
        return nil, errors.New("SKU debe tener entre 3 y 50 caracteres")
    }
    // Validaciones adicionales...
    return &ProductSKU{value: value}, nil
}
```

#### Servicios de Dominio
```go
// Ejemplo: ProductDomainService
type ProductDomainService struct {
    productRepo ProductRepository
}

func (s *ProductDomainService) ValidateForCreation(product *Product) error {
    // Lógica de validación compleja que involucra múltiples entidades
}
```

#### Puertos (Ports)
```go
// Ejemplo: Repository Port
type ProductRepository interface {
    Save(ctx context.Context, product *Product) error
    FindByID(ctx context.Context, id ProductID, tenantID string) (*Product, error)
    FindBySKU(ctx context.Context, sku string, tenantID string) (*Product, error)
    Update(ctx context.Context, product *Product) error
    Delete(ctx context.Context, id ProductID, tenantID string) error
}
```

### 2. Capa de Aplicación (Application Layer)

**Responsabilidad**: Orquesta los casos de uso del sistema, coordina entre el dominio y la infraestructura.

#### Use Cases
```go
// Ejemplo: CreateProductUseCase
type CreateProductUseCase struct {
    productRepo    port.ProductRepository
    domainService  *domain.ProductDomainService
    mapper         *mapper.ProductMapper
}

func (uc *CreateProductUseCase) Execute(
    ctx context.Context,
    req *request.CreateProductRequest,
    tenantID string,
) (*response.ProductResponse, error) {
    // 1. Validar entrada
    // 2. Convertir DTO a entidad de dominio
    // 3. Aplicar reglas de negocio
    // 4. Persistir
    // 5. Convertir a DTO de respuesta
}
```

#### DTOs (Data Transfer Objects)
```go
// Request DTOs
type CreateProductRequest struct {
    Name        string  `json:"name" binding:"required,min=2,max=255"`
    Description *string `json:"description,omitempty" binding:"omitempty,max=1000"`
    SKU         *string `json:"sku,omitempty" binding:"omitempty,min=3,max=50"`
    CategoryID  *string `json:"category_id,omitempty" binding:"omitempty,uuid"`
    BrandID     *string `json:"brand_id,omitempty" binding:"omitempty,uuid"`
}

// Response DTOs
type ProductResponse struct {
    ID          string                     `json:"id"`
    Name        string                     `json:"name"`
    Description *string                    `json:"description,omitempty"`
    SKU         *string                    `json:"sku,omitempty"`
    Status      string                     `json:"status"`
    Category    *CategoryReferenceResponse `json:"category,omitempty"`
    Brand       *BrandReferenceResponse    `json:"brand,omitempty"`
    CreatedAt   time.Time                  `json:"created_at"`
    UpdatedAt   time.Time                  `json:"updated_at"`
}
```

#### Mappers
```go
// Ejemplo: ProductMapper
type ProductMapper struct{}

func (m *ProductMapper) ToResponse(product *domain.Product) *response.ProductResponse {
    // Conversión de entidad de dominio a DTO de respuesta
}

func (m *ProductMapper) RequestToDomain(req *request.CreateProductRequest, tenantID string) (*domain.Product, error) {
    // Conversión de DTO de request a entidad de dominio
}
```

### 3. Capa de Infraestructura (Infrastructure Layer)

**Responsabilidad**: Implementa los adaptadores para tecnologías externas (base de datos, HTTP, etc.).

#### Controllers (HTTP Adapters)
```go
// Ejemplo: ProductController
type ProductController struct {
    createProductUseCase *usecase.CreateProductUseCase
    getProductUseCase    *usecase.GetProductByIDUseCase
    // ... otros use cases
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
    // 1. Extraer datos de la petición HTTP
    // 2. Validar entrada
    // 3. Ejecutar use case
    // 4. Retornar respuesta HTTP
}
```

#### Repositories (Database Adapters)
```go
// Ejemplo: PostgresProductRepository
type PostgresProductRepository struct {
    db *sql.DB
}

func (r *PostgresProductRepository) Save(ctx context.Context, product *domain.Product) error {
    // Implementación específica para PostgreSQL
    query := `INSERT INTO products (id, tenant_id, name, description, sku, status, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
    // ...
}
```

## Domain-Driven Design (DDD)

### Agregados (Aggregates)

#### Product Aggregate
```
Product (Aggregate Root)
├── ProductID (Identity)
├── ProductSKU (Value Object)
├── ProductStatus (Value Object)
├── CategoryReference (Value Object)
├── BrandReference (Value Object)
└── ProductVariants (Entities)
    ├── VariantID (Identity)
    ├── VariantSKU (Value Object)
    ├── VariantStatus (Value Object)
    └── VariantAttributes (Value Object Collection)
```

**Reglas del Agregado**:
- Product es la raíz del agregado
- Las variantes solo se pueden acceder a través del Product
- Todas las modificaciones pasan por la raíz del agregado
- Consistencia transaccional dentro del agregado

### Bounded Contexts

```
┌─────────────────────────────────────────────────────────────┐
│                    PIM Bounded Context                     │
│                                                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │  Category    │  │    Brand     │  │   Product    │     │
│  │  Subdomain   │  │  Subdomain   │  │  Subdomain   │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
│                                                            │
│  ┌─────────────────────────────────────────────────────┐   │
│  │              Shared Kernel                          │   │
│  │  • TenantID                                         │   │
│  │  • Common Value Objects                             │   │
│  │  • Shared Exceptions                                │   │
│  │  • Criteria Pattern                                 │   │
│  └─────────────────────────────────────────────────────┘   │
└────────────────────────────────────────────────────────────┘
```

### Ubiquitous Language

**Términos del Dominio**:
- **Product**: Producto base con información general
- **Product Variant**: Variación específica de un producto con atributos únicos
- **Category**: Clasificación jerárquica de productos
- **Brand**: Marca comercial de productos
- **SKU**: Stock Keeping Unit, código único de identificación
- **Tenant**: Organización que usa el sistema (multi-tenancy)
- **Attribute**: Característica específica de una variante (color, tamaño, etc.)

## Patrones de Diseño Implementados

### 1. Repository Pattern
```go
// Abstracción en el dominio
type ProductRepository interface {
    Save(ctx context.Context, product *Product) error
    FindByID(ctx context.Context, id ProductID, tenantID string) (*Product, error)
}

// Implementación en infraestructura
type PostgresProductRepository struct {
    db *sql.DB
}
```

### 2. Criteria Pattern
```go
// Para consultas complejas y filtros dinámicos
type Criteria struct {
    filters   []Filter
    sorts     []Sort
    page      int
    pageSize  int
}

type ProductCriteriaBuilder struct {
    criteria *Criteria
}

func (b *ProductCriteriaBuilder) WithName(name string) *ProductCriteriaBuilder {
    b.criteria.AddFilter("name", "LIKE", "%"+name+"%")
    return b
}
```

### 3. Factory Pattern
```go
// Para creación de entidades complejas
type ProductFactory struct{}

func (f *ProductFactory) CreateProduct(
    name string,
    tenantID string,
    // ... otros parámetros
) (*Product, error) {
    // Lógica de creación compleja
    // Validaciones
    // Configuración de valores por defecto
}
```

### 4. Specification Pattern
```go
// Para reglas de negocio complejas
type ProductSpecification interface {
    IsSatisfiedBy(product *Product) bool
}

type ActiveProductSpecification struct{}

func (s *ActiveProductSpecification) IsSatisfiedBy(product *Product) bool {
    return product.Status() == ProductStatusActive
}
```

### 5. Command Query Responsibility Segregation (CQRS)
```go
// Commands (modifican estado)
type CreateProductCommand struct {
    Name        string
    Description string
    TenantID    string
}

// Queries (solo lectura)
type GetProductQuery struct {
    ProductID string
    TenantID  string
}
```

## Multi-Tenancy Architecture

### Estrategia: Shared Database, Shared Schema

```sql
-- Todas las tablas incluyen tenant_id
CREATE TABLE products (
    id UUID PRIMARY KEY,
    tenant_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    -- ... otros campos
    
    -- Índices para performance multi-tenant
    INDEX idx_products_tenant_id (tenant_id),
    INDEX idx_products_tenant_name (tenant_id, name),
    
    -- Constraints de unicidad por tenant
    UNIQUE(tenant_id, name),
    UNIQUE(tenant_id, sku)
);
```

### Filtrado Automático por Tenant
```go
// Todos los repositorios filtran automáticamente por tenant
func (r *PostgresProductRepository) FindByID(ctx context.Context, id ProductID, tenantID string) (*Product, error) {
    query := `SELECT * FROM products WHERE id = $1 AND tenant_id = $2`
    // ...
}
```

### Validación de Tenant en Controllers
```go
func (ctrl *ProductController) CreateProduct(c *gin.Context) {
    tenantID := c.GetHeader("X-Tenant-ID")
    if tenantID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
        return
    }
    // ...
}
```

## Principios SOLID Aplicados

### Single Responsibility Principle (SRP)
- Cada clase tiene una única responsabilidad
- Use cases específicos para cada operación
- Separación clara entre capas

### Open/Closed Principle (OCP)
- Extensible mediante interfaces
- Nuevos adaptadores sin modificar código existente
- Criterios dinámicos para filtros

### Liskov Substitution Principle (LSP)
- Implementaciones de repositorios intercambiables
- Polimorfismo en value objects

### Interface Segregation Principle (ISP)
- Interfaces específicas y pequeñas
- Puertos dedicados por funcionalidad

### Dependency Inversion Principle (DIP)
- Dependencias hacia abstracciones
- Inyección de dependencias
- Inversión de control

## Gestión de Errores

### Jerarquía de Excepciones
```go
// Errores de dominio
type DomainError struct {
    Code    string
    Message string
    Details map[string]interface{}
}

// Errores específicos
type ProductNotFoundError struct {
    ProductID string
    TenantID  string
}

type ProductAlreadyExistsError struct {
    Name     string
    TenantID string
}
```

### Manejo en Capas
```go
// En Use Cases
func (uc *CreateProductUseCase) Execute(...) (*response.ProductResponse, error) {
    if err := uc.domainService.ValidateForCreation(product); err != nil {
        return nil, fmt.Errorf("validation failed: %w", err)
    }
    // ...
}

// En Controllers
func (ctrl *ProductController) CreateProduct(c *gin.Context) {
    product, err := ctrl.createProductUseCase.Execute(...)
    if err != nil {
        switch {
        case errors.Is(err, domain.ProductAlreadyExistsError{}):
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        default:
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        }
        return
    }
    // ...
}
```

## Testing Strategy

### Unit Tests
- Tests de entidades de dominio
- Tests de value objects
- Tests de servicios de dominio
- Tests de use cases

### Integration Tests
- Tests de repositorios
- Tests de controllers
- Tests end-to-end

### Test Doubles
```go
// Mock Repository para tests
type MockProductRepository struct {
    products map[string]*domain.Product
}

func (m *MockProductRepository) Save(ctx context.Context, product *domain.Product) error {
    m.products[product.ID().String()] = product
    return nil
}
```

## Performance Considerations

### Database Optimization
- Índices optimizados para multi-tenancy
- Prepared statements para prevenir SQL injection
- Connection pooling

### Caching Strategy
- Cache de entidades frecuentemente accedidas
- Cache de resultados de consultas complejas
- Invalidación de cache por eventos de dominio

### Pagination
- Cursor-based pagination para grandes datasets
- Límites configurables por tenant
- Optimización de queries con LIMIT/OFFSET

Esta arquitectura garantiza un sistema robusto, mantenible y escalable que puede evolucionar con los requisitos del negocio. 