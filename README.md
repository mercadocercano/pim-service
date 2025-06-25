# SaaS Multitenant PIM Service

**Product Information Management (PIM) Service** para el ecosistema SaaS Marketplace multitenant. Este servicio gestiona productos, categorías, atributos, marcas y catálogo global con arquitectura hexagonal.

## 🏗️ Arquitectura

El servicio implementa **Arquitectura Hexagonal (Clean Architecture)** con separación clara de responsabilidades:

- **Domain**: Entidades de negocio, value objects, servicios de dominio
- **Application**: Casos de uso, mappers, DTOs
- **Infrastructure**: Controllers REST, repositorios, base de datos

## 📦 Módulos Implementados

### **🛍️ Product Management**
- **Products (Tenant)**: CRUD de productos del tenant con estados flexibles
- **Product Variants**: Gestión de variantes con atributos específicos
- **Product Status**: Estados avanzados (draft, pending, active, inactive, discontinued, deleted)

### **🚀 Quickstart**
- **Template Import**: Crear productos desde templates del catálogo global
- **Bulk Import**: Importación masiva por tipo de negocio
- **Progress Tracking**: Seguimiento inteligente del progreso de configuración

### **🗂️ Category Management**
- **Categories**: Gestión jerárquica de categorías con soporte multitenant
- **Marketplace Categories**: Categorías globales del marketplace
- **Tenant Category Mapping**: Mapeo de categorías marketplace a tenant

### **🏷️ Attribute Management**
- **Attributes**: Atributos base del marketplace
- **Marketplace Attributes**: Atributos globales administrados centralmente
- **Tenant Custom Attributes**: Extensiones personalizadas por tenant
- **Category Attributes**: Relación atributos-categorías con valores permitidos

### **🌟 Brand Management**
- **Brands**: Gestión de marcas con validaciones de unicidad por tenant

### **📊 Business Types**
- **Business Types**: Clasificación de tipos de negocio para quickstart
- **Argentina Business Types**: Datos precargados para el mercado argentino

### **🌍 Global Catalog**
- **Global Products**: Catálogo de referencia con productos template
- **EAN Search**: Búsqueda por código de barras
- **Business Type Suggestions**: Productos sugeridos por tipo de negocio

### **🔧 Infrastructure**
- **Health Check**: Monitoreo de estado del servicio y base de datos
- **OpenAPI Documentation**: Documentación Swagger completa
- **Metrics**: Integración con Prometheus (opcional)

## 🛣️ API Endpoints

### **Productos (Tenant)**
```
POST   /api/v1/products                    # Crear producto
GET    /api/v1/products                    # Listar productos con filtros
GET    /api/v1/products/:id                # Obtener producto
PUT    /api/v1/products/:id                # Actualizar producto
DELETE /api/v1/products/:id                # Eliminar producto
PATCH  /api/v1/products/:id/status         # Cambiar estado
GET    /api/v1/products/:id/status/transitions # Obtener transiciones disponibles
```

### **Variantes de Producto**
```
POST   /api/v1/products/:product_id/variants         # Crear variante
GET    /api/v1/products/:product_id/variants         # Listar variantes
GET    /api/v1/products/:product_id/variants/:id     # Obtener variante
PUT    /api/v1/products/:product_id/variants/:id     # Actualizar variante
DELETE /api/v1/products/:product_id/variants/:id     # Eliminar variante
GET    /api/v1/variants                              # Listar todas las variantes
```

### **Quickstart**
```
POST   /api/v1/quickstart/products/from-template           # Crear desde template
POST   /api/v1/quickstart/products/import-from-business-type # Importación masiva
GET    /api/v1/quickstart/progress                          # Progreso del setup
```

### **Categorías**
```
POST   /api/v1/categories                  # Crear categoría
GET    /api/v1/categories                  # Listar categorías
GET    /api/v1/categories/:id              # Obtener categoría
PUT    /api/v1/categories/:id              # Actualizar categoría
DELETE /api/v1/categories/:id              # Eliminar categoría
PATCH  /api/v1/categories/:id/status       # Cambiar estado
PATCH  /api/v1/categories/:id/move         # Mover en jerarquía
```

### **Atributos**
```
GET    /api/v1/attributes                  # Listar atributos
POST   /api/v1/attributes                  # Crear atributo
GET    /api/v1/attributes/:id              # Obtener atributo
PUT    /api/v1/attributes/:id              # Actualizar atributo
DELETE /api/v1/attributes/:id              # Eliminar atributo
```

### **Atributos de Categoría**
```
GET    /api/v1/category-attributes         # Listar con filtros y paginación
GET    /api/v1/category-attributes/simple  # Listado simple
GET    /api/v1/category-attributes/detailed # Listado detallado
POST   /api/v1/category-attributes         # Crear relación
PUT    /api/v1/category-attributes/:id     # Actualizar relación
DELETE /api/v1/category-attributes/:id     # Eliminar relación
```

### **Marcas**
```
POST   /api/v1/brands                      # Crear marca
GET    /api/v1/brands                      # Listar marcas
GET    /api/v1/brands/:id                  # Obtener marca
PUT    /api/v1/brands/:id                  # Actualizar marca
DELETE /api/v1/brands/:id                  # Eliminar marca
```

### **Tipos de Negocio**
```
POST   /api/v1/business-types              # Crear tipo de negocio
GET    /api/v1/business-types              # Listar tipos
GET    /api/v1/business-types/:id          # Obtener tipo
PUT    /api/v1/business-types/:id          # Actualizar tipo
DELETE /api/v1/business-types/:id          # Eliminar tipo
```

### **Catálogo Global**
```
# Rutas Públicas
GET    /api/v1/public/global-catalog/health           # Health check
GET    /api/v1/public/global-catalog/search           # Búsqueda por EAN
GET    /api/v1/public/global-catalog/suggestions      # Sugerencias por business type
GET    /api/v1/public/global-catalog/products/ean/:ean # Producto por EAN

# Rutas Privadas (Admin/Scrapers)
POST   /api/v1/global-catalog/products                # Crear producto global
GET    /api/v1/global-catalog/products                # Listar productos
GET    /api/v1/global-catalog/products/search         # Búsqueda avanzada
```

### **Sistema**
```
GET    /api/v1/health                      # Health check del servicio
GET    /api-docs                           # Documentación Swagger
GET    /openapi.yaml                       # Especificación OpenAPI
GET    /metrics                            # Métricas Prometheus (opcional)
```

## 🔄 Estados de Productos

### **Matriz de Transiciones**
```
draft        → pending, active, deleted
pending      → draft, active, deleted  
active       → inactive, discontinued, deleted
inactive     → active, discontinued, deleted
discontinued → deleted
deleted      → (sin transiciones)
```

### **Criterios por Estado**
- **draft**: Mínimo nombre requerido
- **pending**: Nombre + descripción + categoría
- **active**: Todos los campos + al menos 1 variante activa
- **inactive**: Producto temporalmente deshabilitado
- **discontinued**: Producto descontinuado permanentemente
- **deleted**: Soft delete

## 🚀 Tecnologías

- **Lenguaje**: Go 1.21+
- **Framework**: Gin (HTTP Router)
- **Base de Datos**: PostgreSQL 15+
- **MongoDB**: Para catálogo global (marketplace)
- **Documentación**: OpenAPI 3.0 / Swagger
- **Métricas**: Prometheus (opcional)
- **Contenedores**: Docker

## 🛠️ Desarrollo

### **Prerrequisitos**
```bash
# Software requerido
Go 1.21+
PostgreSQL 15+
MongoDB 6.0+ (para global catalog)
Docker & Docker Compose (opcional)
```

### **Configuración Local**
```bash
# 1. Clonar repositorio
git clone <repository-url>
cd saas-mt-pim-service

# 2. Instalar dependencias
go mod download

# 3. Configurar variables de entorno
cp .env.example .env
# Editar .env con tus configuraciones

# 4. Ejecutar migraciones
./scripts/migrate.sh

# 5. Cargar datos semilla (opcional)
./scripts/seed_business_types_argentina.sh

# 6. Ejecutar servicio
go run main.go
```

### **Variables de Entorno**
```bash
# Base de datos PostgreSQL
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=pim_db

# MongoDB (Global Catalog)
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=marketplace_global

# Servicio
PORT=8080
GIN_MODE=debug

# Métricas (opcional)
PROMETHEUS_ENABLED=true
```

### **Docker Compose**
```bash
# Levantar servicios completos
docker-compose up -d

# Solo base de datos
docker-compose up -d postgres mongodb

# Construir y ejecutar servicio
docker-compose up --build pim-service
```

## 🧪 Testing

### **Ejecutar Tests**
```bash
# Tests unitarios
go test ./...

# Tests con coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Tests de integración
./test-integration/run_integration_tests.sh
```

### **Coverage Actual**
- **Cobertura total**: ~85%
- **Módulos críticos**: >90%
- **Tests de integración**: Completos para todos los endpoints

## 📁 Estructura del Proyecto

```
saas-mt-pim-service/
├── src/
│   ├── api/                    # API infrastructure
│   ├── attribute/              # Attribute management
│   ├── brand/                  # Brand management
│   ├── businesstype/           # Business types
│   ├── category/               # Category management
│   ├── category_attribute/     # Category-Attribute relations
│   ├── product/
│   │   ├── tenant/             # Tenant products
│   │   ├── quickstart/         # Quickstart functionality
│   │   └── global_catalog/     # Global catalog
│   ├── quickstart/             # Legacy quickstart (YAML-based)
│   └── shared/                 # Shared infrastructure
├── documentation/              # Complete documentation
├── migrations/                 # Database migrations
├── seeds/                      # Seed data
├── scripts/                    # Utility scripts
├── test/                       # Unit tests
├── test-integration/           # Integration tests
├── api-docs/                   # OpenAPI specifications
└── templates/                  # HTML templates
```

## 📚 Documentación

### **Documentación Técnica**
- [Arquitectura Final](./documentation/ARCHITECTURE_FINAL.md)
- [Estructura de Rutas](./documentation/ROUTES_STRUCTURE.md)
- [Endpoints Quickstart](./documentation/QUICKSTART_ENDPOINTS.md)
- [Global Catalog Design](./documentation/GLOBAL_CATALOG_DESIGN.md)
- [Guía de Instalación](./documentation/installation.md)

### **Documentación de API**
- **Swagger UI**: http://localhost:8080/api-docs
- **OpenAPI Spec**: http://localhost:8080/openapi.yaml
- [Ejemplos cURL](./documentation/curl-examples-global-catalog.md)

### **Testing & QA**
- [Reporte de Coverage](./documentation/test_coverage_report.md)
- [Resumen de Tests](./documentation/TESTING_COMPLETION_SUMMARY.md)
- [Estado de Integración](./documentation/MARKETPLACE_INTEGRATION_STATUS.md)

## 🔄 Flujo de Negocio

### **Configuración Inicial (Quickstart)**
```
1. Tenant registra → selecciona business_type
2. GET /global-catalog/business-types/{id}/products → Ve productos sugeridos
3. POST /quickstart/products/import-from-business-type → Importa masivamente
4. POST /quickstart/products/from-template → Crea productos específicos
5. PATCH /products/{id}/status → draft → pending → active
6. GET /quickstart/progress → Monitorea progreso
7. GET /products → Ve productos activos para marketplace
```

### **Gestión Diaria**
```
1. POST /products → Crear nuevos productos
2. POST /products/{id}/variants → Agregar variantes
3. PUT /products/{id} → Actualizar información
4. PATCH /products/{id}/status → Gestionar estados
5. GET /products → Consultar inventario
```

## 🚀 Deployment

### **Producción**
```bash
# Build optimizado
go build -ldflags="-s -w" -o pim-service main.go

# Docker
docker build -t saas-mt-pim-service:latest .
docker run -p 8080:8080 saas-mt-pim-service:latest
```

### **Monitoreo**
- **Health Check**: `/api/v1/health`
- **Métricas**: `/metrics` (si Prometheus está habilitado)
- **Logs**: Estructurados en JSON para producción

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama feature (`git checkout -b feature/amazing-feature`)
3. Commit cambios (`git commit -m 'Add amazing feature'`)
4. Push a la rama (`git push origin feature/amazing-feature`)
5. Abre un Pull Request

### **Estándares de Código**
- **Arquitectura Hexagonal**: Mantener separación de capas
- **Tests**: Cobertura mínima 80%
- **Documentación**: Actualizar README y docs/ en cada cambio
- **API**: Seguir convenciones REST y OpenAPI

## 📄 Licencia

Este proyecto es parte del ecosistema SaaS Marketplace.

---

**Desarrollado con ❤️ para el ecosistema SaaS Multitenant**
