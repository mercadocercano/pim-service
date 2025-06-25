# 📚 Documentación del PIM Service

## 🎯 Descripción General

El **PIM Service (Product Information Management)** es un microservicio multitenant para la gestión completa de información de productos, desarrollado con **arquitectura hexagonal**, **Domain-Driven Design (DDD)** y principios **SOLID**. Forma parte del ecosistema SaaS Marketplace y permite a múltiples organizaciones gestionar sus catálogos de productos de forma independiente y segura.

## ✨ Características Principales

- **🏢 Multi-tenancy**: Aislamiento completo de datos por tenant
- **🏗️ Arquitectura Hexagonal**: Separación clara entre dominio, aplicación e infraestructura
- **🎯 Domain-Driven Design**: Modelado rico del dominio con entidades, value objects y servicios
- **🌐 API RESTful**: Endpoints bien documentados con OpenAPI/Swagger
- **🗃️ PostgreSQL + MongoDB**: Almacenamiento híbrido robusto y escalable
- **🐳 Containerización**: Despliegue con Docker y Docker Compose
- **🚪 API Gateway**: Kong para enrutamiento y gestión de APIs
- **📊 Monitoreo**: Prometheus, Grafana y Loki para observabilidad
- **🚀 Quickstart**: Configuración rápida de catálogos predefinidos por tipo de negocio
- **🔄 Estados Flexibles**: Gestión avanzada de estados de productos con transiciones validadas

## 📦 Módulos Implementados

### **🛍️ Product Management**
- **Products (Tenant)**: CRUD completo con estados flexibles (draft, pending, active, inactive, discontinued, deleted)
- **Product Variants**: Gestión de variantes con atributos dinámicos y SKU únicos
- **Product Status Service**: Validaciones de transición y lógica de negocio avanzada

### **🚀 Quickstart**
- **Template Import**: Crear productos individuales desde templates del catálogo global
- **Bulk Import**: Importación masiva basada en tipo de negocio con reporte detallado
- **Progress Tracking**: Seguimiento inteligente del progreso de configuración con recomendaciones

### **🗂️ Category Management**
- **Categories**: Gestión jerárquica con soporte multitenant
- **Marketplace Categories**: Categorías globales administradas centralmente
- **Tenant Category Mapping**: Mapeo personalizado de categorías marketplace a tenant

### **🏷️ Attribute Management**
- **Attributes**: Atributos base del marketplace
- **Marketplace Attributes**: Atributos globales con control centralizado
- **Tenant Custom Attributes**: Extensiones personalizadas por tenant
- **Category Attributes**: Relación atributos-categorías con valores permitidos y filtros avanzados

### **🌟 Brand Management**
- **Brands**: Gestión completa con validaciones de unicidad por tenant
- **Brand Status**: Estados y soft delete para preservar integridad referencial

### **📊 Business Types**
- **Business Types**: Clasificación de tipos de negocio para quickstart
- **Argentina Business Types**: Datos precargados específicos del mercado argentino

### **🌍 Global Catalog**
- **Global Products**: Catálogo de referencia con productos template
- **EAN Search**: Búsqueda por código de barras con APIs públicas y privadas
- **Business Type Suggestions**: Productos sugeridos por tipo de negocio

### **🔧 Infrastructure**
- **Health Check**: Monitoreo completo del estado del servicio y dependencias
- **OpenAPI Documentation**: Documentación Swagger interactiva
- **Metrics**: Integración opcional con Prometheus

## 🏗️ Arquitectura del Sistema

```
┌─────────────────────────────────────────────────────────────┐
│                     API Gateway (Kong)                      │
│                    Puerto: 8001                            │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                  PIM Service                               │
│                 Puerto: 8080                               │
│                                                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │  Products    │  │  Categories  │  │  Attributes  │     │
│  │   Module     │  │    Module    │  │    Module    │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
│                                                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │    Brands    │  │ Business     │  │🚀 Quickstart │     │
│  │    Module    │  │  Types       │  │   Module     │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
│                                                            │
│  ┌─────────────────────────────────────────────────────┐   │
│  │            Arquitectura Hexagonal                  │   │
│  │                                                     │   │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │   │
│  │  │   Domain    │  │ Application │  │Infrastructure│ │   │
│  │  │             │  │             │  │             │ │   │
│  │  │ • Entities  │  │ • Use Cases │  │ • Controllers│ │   │
│  │  │ • Value Obj │  │ • Mappers   │  │ • Repository │ │   │
│  │  │ • Services  │  │ • DTOs      │  │ • Database   │ │   │
│  │  │ • Ports     │  │ • Criteria  │  │ • HTTP       │ │   │
│  │  └─────────────┘  └─────────────┘  └─────────────┘ │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                 Almacenamiento                             │
│                                                            │
│  ┌─────────────────────┐  ┌─────────────────────────────┐  │
│  │   PostgreSQL        │  │        MongoDB              │  │
│  │   Puerto: 5432      │  │      Puerto: 27017         │  │
│  │                     │  │                             │  │
│  │ • products          │  │ • global_catalog            │  │
│  │ • categories        │  │ • marketplace_products      │  │
│  │ • attributes        │  │ • scraped_data              │  │
│  │ • brands            │  │                             │  │
│  │ • business_types    │  │                             │  │
│  └─────────────────────┘  └─────────────────────────────┘  │
└────────────────────────────────────────────────────────────┘
```

## 🛣️ Estructura de Rutas Completa

### **🛍️ Productos (Tenant) - `/api/v1/products`**
```
POST   /products                    # Crear producto
GET    /products                    # Listar con filtros y paginación
GET    /products/:id                # Obtener producto específico
PUT    /products/:id                # Actualizar producto
DELETE /products/:id                # Eliminar producto (soft delete)
PATCH  /products/:id/status         # Cambiar estado con validaciones
GET    /products/:id/status/transitions # Obtener transiciones disponibles
```

### **🔧 Variantes de Producto - `/api/v1/products/:product_id/variants`**
```
POST   /variants                    # Crear variante
GET    /variants                    # Listar variantes del producto
GET    /variants/:variant_id        # Obtener variante específica
PUT    /variants/:variant_id        # Actualizar variante
DELETE /variants/:variant_id        # Eliminar variante
GET    /api/v1/variants            # Listar todas las variantes del tenant
```

### **🚀 Quickstart - `/api/v1/quickstart`**
```
POST   /products/from-template      # Crear producto desde template
POST   /products/import-from-business-type # Importación masiva
GET    /progress                    # Progreso del setup con recomendaciones
```

### **🗂️ Categorías - `/api/v1/categories`**
```
POST   /                           # Crear categoría
GET    /                           # Listar con filtros jerárquicos
GET    /:id                        # Obtener categoría específica
PUT    /:id                        # Actualizar categoría
DELETE /:id                        # Eliminar categoría
PATCH  /:id/status                 # Cambiar estado
PATCH  /:id/move                   # Mover en jerarquía
```

### **🏷️ Atributos - `/api/v1/attributes`**
```
GET    /                           # Listar atributos
POST   /                           # Crear atributo
GET    /:id                        # Obtener atributo
PUT    /:id                        # Actualizar atributo
DELETE /:id                        # Eliminar atributo
```

### **🔗 Atributos de Categoría - `/api/v1/category-attributes`**
```
GET    /                           # Listar con filtros avanzados y paginación
GET    /simple                     # Listado simple sin detalles
GET    /detailed                   # Listado detallado con relaciones
POST   /                           # Crear relación categoría-atributo
PUT    /:id                        # Actualizar relación
DELETE /:id                        # Eliminar relación
```

### **🌟 Marcas - `/api/v1/brands`**
```
POST   /                           # Crear marca
GET    /                           # Listar marcas del tenant
GET    /:id                        # Obtener marca específica
PUT    /:id                        # Actualizar marca
DELETE /:id                        # Eliminar marca (soft delete)
```

### **📊 Tipos de Negocio - `/api/v1/business-types`**
```
POST   /                           # Crear tipo de negocio (admin)
GET    /                           # Listar tipos disponibles
GET    /:id                        # Obtener tipo específico
PUT    /:id                        # Actualizar tipo (admin)
DELETE /:id                        # Eliminar tipo (admin)
```

### **🌍 Catálogo Global**
```
# Rutas Públicas - /api/v1/public/global-catalog
GET    /health                     # Health check público
GET    /search                     # Búsqueda por EAN
GET    /suggestions                # Sugerencias por business type
GET    /products/ean/:ean          # Producto por EAN

# Rutas Privadas - /api/v1/global-catalog
POST   /products                   # Crear producto global (scrapers)
GET    /products                   # Listar productos con filtros
GET    /products/search            # Búsqueda avanzada
```

### **🔧 Sistema - `/api/v1`**
```
GET    /health                     # Health check completo
GET    /api-docs                   # Documentación Swagger UI
GET    /openapi.yaml               # Especificación OpenAPI
GET    /metrics                    # Métricas Prometheus (opcional)
```

## 🔄 Estados y Flujos de Negocio

### **Estados de Productos**
```
draft        → pending, active, deleted
pending      → draft, active, deleted  
active       → inactive, discontinued, deleted
inactive     → active, discontinued, deleted
discontinued → deleted
deleted      → (sin transiciones)
```

### **Flujo de Quickstart**
```
1. Tenant registra → selecciona business_type
2. GET /global-catalog/business-types/{id}/products → Ve productos sugeridos
3. POST /quickstart/products/import-from-business-type → Importa masivamente
4. POST /quickstart/products/from-template → Crea productos específicos
5. PATCH /products/{id}/status → draft → pending → active
6. GET /quickstart/progress → Monitorea progreso
7. GET /products → Ve productos activos para marketplace
```

## 🛠️ Tecnologías Utilizadas

- **Backend**: Go 1.21+
- **Framework Web**: Gin (HTTP Router)
- **Base de Datos**: PostgreSQL 15+ (principal), MongoDB 6.0+ (global catalog)
- **ORM**: SQL nativo con database/sql + drivers
- **API Gateway**: Kong
- **Containerización**: Docker & Docker Compose
- **Documentación**: OpenAPI 3.0 / Swagger
- **Monitoreo**: Prometheus, Grafana, Loki
- **Logs**: Structured logging JSON
- **Testing**: Go testing + coverage reports

## 📁 Estructura del Proyecto

```
saas-mt-pim-service/
├── src/
│   ├── api/                    # API infrastructure (health, docs)
│   ├── attribute/              # Attribute management
│   │   ├── domain/             # Entities, value objects, ports
│   │   ├── application/        # Use cases, DTOs, mappers
│   │   └── infrastructure/     # Controllers, repositories, persistence
│   ├── brand/                  # Brand management
│   ├── businesstype/           # Business types
│   ├── category/               # Category management
│   ├── category_attribute/     # Category-Attribute relations
│   ├── product/
│   │   ├── tenant/             # Tenant products (main functionality)
│   │   ├── quickstart/         # Quickstart functionality
│   │   └── global_catalog/     # Global catalog
│   ├── quickstart/             # Legacy quickstart (YAML-based)
│   └── shared/                 # Shared infrastructure (middleware, criteria)
├── documentation/              # Complete project documentation
│   ├── ARCHITECTURE_FINAL.md   # Detailed architecture
│   ├── QUICKSTART_ENDPOINTS.md # Quickstart API documentation
│   ├── ROUTES_STRUCTURE.md     # Complete routes reference
│   └── [otros archivos]
├── migrations/                 # Database migrations
├── seeds/                      # Seed data (business types, etc.)
├── scripts/                    # Utility scripts
├── test/                       # Unit tests
├── test-integration/           # Integration tests
├── api-docs/                   # OpenAPI specifications
├── templates/                  # HTML templates (Swagger UI)
├── docker-compose.yml          # Services configuration
├── Dockerfile                  # Service image
└── main.go                     # Application entry point
```

## 📚 Índice de Documentación

### **📋 Documentación Principal**
- [**README Principal**](../README.md) - Visión general y quick start
- [**Arquitectura Final**](./ARCHITECTURE_FINAL.md) - Arquitectura detallada del sistema
- [**Estructura de Rutas**](./ROUTES_STRUCTURE.md) - Referencia completa de endpoints
- [**Guía de Instalación**](./installation.md) - Setup completo paso a paso

### **🚀 Quickstart**
- [**Endpoints Quickstart**](./QUICKSTART_ENDPOINTS.md) - API del módulo quickstart
- [**Guía de Integración Quickstart**](./QUICKSTART_INTEGRATION_GUIDE.md) - Integración con onboarding
- [**Módulo Quickstart Legacy**](./quickstart-module.md) - Documentación del módulo YAML

### **🌍 Global Catalog**
- [**Global Catalog Design**](./GLOBAL_CATALOG_DESIGN.md) - Diseño del catálogo global
- [**Global Catalog Roadmap**](./GLOBAL_CATALOG_ROADMAP.md) - Roadmap de implementación
- [**Ejemplos cURL Global Catalog**](./curl-examples-global-catalog.md) - Ejemplos de uso

### **🧪 Testing & QA**
- [**Reporte de Coverage**](./test_coverage_report.md) - Cobertura de tests actual
- [**Resumen de Tests**](./TESTING_COMPLETION_SUMMARY.md) - Estado de testing completo
- [**Tests del Marketplace**](./MARKETPLACE_TESTS_SUMMARY.md) - Tests específicos del marketplace

### **🔧 Desarrollo**
- [**Arquitectura Detallada**](./architecture.md) - Arquitectura técnica profunda
- [**Resumen de Reorganización**](./REORGANIZATION_SUMMARY.md) - Historial de cambios arquitectónicos
- [**Estado de Integración**](./MARKETPLACE_INTEGRATION_STATUS.md) - Estado de integración con otros servicios

## 🔐 Multi-Tenancy

El sistema implementa **multi-tenancy a nivel de aplicación**:

### **Implementación**
- **Header requerido**: `X-Tenant-ID` en todas las peticiones que requieren contexto de tenant
- **Aislamiento de datos**: Cada tenant solo accede a sus propios datos
- **Base de datos compartida**: PostgreSQL compartida con filtrado automático por `tenant_id`
- **Validaciones**: Todas las operaciones validan pertenencia al tenant
- **Global Catalog**: Datos compartidos accesibles por todos los tenants

### **Rutas por Contexto**
- **Tenant Context**: `/api/v1/products`, `/api/v1/categories`, etc. (requieren `X-Tenant-ID`)
- **Global Context**: `/api/v1/public/global-catalog/*` (públicas, sin tenant)
- **Admin Context**: `/api/v1/business-types`, `/api/v1/global-catalog/*` (administración)

## 🚀 Getting Started

### **Quick Start**
```bash
# 1. Clonar y configurar
git clone <repository-url>
cd saas-mt-pim-service
cp .env.example .env

# 2. Levantar servicios
docker-compose up -d

# 3. Ejecutar migraciones
./scripts/migrate.sh

# 4. Cargar datos semilla
./scripts/seed_business_types_argentina.sh

# 5. Acceder a documentación
open http://localhost:8080/api-docs
```

### **Desarrollo Local**
```bash
# Instalar dependencias
go mod download

# Ejecutar tests
go test ./...

# Ejecutar servicio
go run main.go

# Ver cobertura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🤝 Contribución

1. **Fork** el proyecto
2. **Crear** rama feature (`git checkout -b feature/amazing-feature`)
3. **Commit** cambios (`git commit -m 'Add amazing feature'`)
4. **Push** a la rama (`git push origin feature/amazing-feature`)
5. **Abrir** Pull Request

### **Estándares**
- **Arquitectura Hexagonal**: Mantener separación de capas
- **Tests**: Cobertura mínima 80%
- **Documentación**: Actualizar docs en cada cambio
- **API**: Seguir convenciones REST y OpenAPI

---

**📖 Documentación mantenida al día - Última actualización: Enero 2024** 