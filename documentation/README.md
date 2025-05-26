# Sistema PIM (Product Information Management)

## Descripción General

El Sistema PIM es una solución multi-tenant para la gestión de información de productos, desarrollado con arquitectura hexagonal, Domain-Driven Design (DDD) y principios SOLID. Permite a múltiples organizaciones gestionar sus catálogos de productos de forma independiente y segura.

## Características Principales

- **Multi-tenancy**: Aislamiento completo de datos por tenant
- **Arquitectura Hexagonal**: Separación clara entre dominio, aplicación e infraestructura
- **Domain-Driven Design**: Modelado rico del dominio con entidades, value objects y servicios
- **API RESTful**: Endpoints bien documentados con OpenAPI/Swagger
- **Base de datos PostgreSQL**: Almacenamiento robusto y escalable
- **Containerización**: Despliegue con Docker y Docker Compose
- **API Gateway**: Kong para enrutamiento y gestión de APIs
- **Monitoreo**: Prometheus, Grafana y Loki para observabilidad
- **🚀 Quickstart**: Configuración rápida de catálogos predefinidos por tipo de negocio

## Módulos Implementados

### 1. Categorías
- Gestión jerárquica de categorías de productos
- Estructura de árbol con categorías padre e hijas
- Estados: ACTIVE, INACTIVE
- Operaciones: crear, actualizar, activar, desactivar, mover

### 2. Marcas (Brands)
- Gestión de marcas de productos
- Información: nombre, descripción, logo, sitio web
- Estados: active, inactive, deleted
- Soft delete para preservar integridad referencial

### 3. Productos
- Gestión de productos base
- Información: nombre, descripción, SKU, categoría, marca
- Estados: active, inactive, discontinued, deleted
- Referencias desacopladas a categorías y marcas

### 4. Variantes de Productos
- Gestión de variantes específicas de productos
- Atributos dinámicos (color, tamaño, capacidad, etc.)
- Variante por defecto automática
- SKU único por variante
- Estados independientes del producto base

### 5. 🚀 Quickstart (Nuevo)
- Configuración rápida de catálogos predefinidos
- 14 tipos de negocio soportados (retail completamente implementado)
- Datos YAML con categorías, atributos, variantes, productos y marcas
- Integración con sistema de onboarding multi-tenant
- Sistema de caché thread-safe para optimización

## Arquitectura del Sistema

```
┌─────────────────────────────────────────────────────────────┐
│                     API Gateway (Kong)                      │
│                    Puerto: 8001                            │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                  Servicio PIM                              │
│                 Puerto: 8090                               │
│                                                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │  Categories  │  │    Brands    │  │   Products   │     │
│  │    Module    │  │    Module    │  │    Module    │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
│                                                            │
│  ┌──────────────┐  ┌─────────────────────────────────┐     │
│  │   Variants   │  │        🚀 Quickstart           │     │
│  │    Module    │  │         Module                 │     │
│  └──────────────┘  └─────────────────────────────────┘     │
│                                                            │
│  ┌─────────────────────────────────────────────────────┐   │
│  │            Arquitectura Hexagonal                  │   │
│  │                                                     │   │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │   │
│  │  │   Domain    │  │ Application │  │Infrastructure│ │   │
│  │  │             │  │             │  │             │ │   │
│  │  │ • Entities  │  │ • Use Cases │  │ • Controllers│ │   │
│  │  │ • Value Obj │  │ • DTOs      │  │ • Repository │ │   │
│  │  │ • Services  │  │ • Mappers   │  │ • Database   │ │   │
│  │  │ • Ports     │  │ • Validators│  │ • HTTP       │ │   │
│  │  └─────────────┘  └─────────────┘  └─────────────┘ │   │
│  └─────────────────────────────────────────────────────┘   │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                PostgreSQL Database                         │
│                  Puerto: 5432                             │
│                                                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │  categories  │  │    brands    │  │   products   │     │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
│                                                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐     │
│  │product_variants│ │variant_attributes│ │quickstart_history│ │
│  └──────────────┘  └──────────────┘  └──────────────┘     │
└────────────────────────────────────────────────────────────┘
```

## Tecnologías Utilizadas

- **Backend**: Go 1.22
- **Framework Web**: Gin
- **Base de Datos**: PostgreSQL 15
- **ORM**: SQL nativo con database/sql
- **API Gateway**: Kong
- **Containerización**: Docker & Docker Compose
- **Documentación**: OpenAPI 3.1 / Swagger
- **Monitoreo**: Prometheus, Grafana, Loki
- **Logs**: Structured logging con Promtail
- **Configuración**: YAML para datos predefinidos

## Estructura del Proyecto

```
pim/
├── src/
│   ├── category/           # Módulo de categorías
│   ├── brand/             # Módulo de marcas
│   ├── product/           # Módulo de productos y variantes
│   ├── quickstart/        # 🚀 Módulo quickstart
│   │   ├── domain/        # Entidades y lógica de negocio
│   │   ├── application/   # Casos de uso
│   │   ├── infrastructure/# Controllers y loaders
│   │   └── data/          # Datos YAML predefinidos
│   ├── shared/            # Código compartido
│   └── api/               # Configuración de API
├── migrations/            # Migraciones de base de datos
├── api-docs/             # Documentación OpenAPI
├── documentation/        # Documentación del proyecto
├── docker-compose.yml    # Configuración de servicios
├── Dockerfile           # Imagen del servicio PIM
└── main.go             # Punto de entrada de la aplicación
```

## Guías de Uso

### Documentación Principal
- [Guía de Instalación](./installation.md)
- [Guía de Desarrollo](./development.md)
- [Arquitectura Detallada](./architecture.md)
- [API Reference](./api-reference.md)
- [Ejemplos de Uso](./examples.md)
- [Troubleshooting](./troubleshooting.md)

### 🚀 Documentación Quickstart
- [**Módulo Quickstart**](./quickstart-module.md) - Documentación completa del módulo
- [**Integración con Onboarding**](../../documentation/Onboarding.md#-integración-con-el-módulo-quickstart-del-pim) - Cómo se integra con el proceso de onboarding
- [**Integración Técnica**](../../documentation/PIM-Onboarding-Integration.md) - Detalles técnicos de la integración

## Endpoints Principales

### Categorías
- `GET /api/v1/categories` - Listar categorías
- `POST /api/v1/categories` - Crear categoría
- `GET /api/v1/categories/{id}` - Obtener categoría
- `PUT /api/v1/categories/{id}` - Actualizar categoría
- `DELETE /api/v1/categories/{id}` - Eliminar categoría

### Marcas
- `GET /api/v1/brands` - Listar marcas
- `POST /api/v1/brands` - Crear marca
- `GET /api/v1/brands/{id}` - Obtener marca
- `PUT /api/v1/brands/{id}` - Actualizar marca
- `DELETE /api/v1/brands/{id}` - Eliminar marca

### Productos
- `GET /api/v1/products` - Listar productos
- `POST /api/v1/products` - Crear producto
- `GET /api/v1/products/{id}` - Obtener producto
- `PUT /api/v1/products/{id}` - Actualizar producto
- `DELETE /api/v1/products/{id}` - Eliminar producto

### Variantes de Productos
- `GET /api/v1/product-variants` - Listar variantes
- `POST /api/v1/product-variants` - Crear variante
- `GET /api/v1/product-variants/{id}` - Obtener variante
- `PUT /api/v1/product-variants/{id}` - Actualizar variante
- `DELETE /api/v1/product-variants/{id}` - Eliminar variante

### 🚀 Quickstart
- `GET /api/quickstart/business-types` - Listar tipos de negocio disponibles
- `GET /api/quickstart/categories/{businessType}` - Obtener categorías predefinidas
- `GET /api/quickstart/attributes/{businessType}` - Obtener atributos predefinidos
- `GET /api/quickstart/variants/{businessType}` - Obtener configuraciones de variantes
- `GET /api/quickstart/products/{businessType}` - Obtener productos de ejemplo
- `POST /api/quickstart/setup` - Aplicar configuración quickstart completa

## Multi-Tenancy

El sistema implementa multi-tenancy a nivel de aplicación:

- **Header requerido**: `X-Tenant-ID` en todas las peticiones
- **Aislamiento de datos**: Cada tenant solo ve sus propios datos
- **Base de datos compartida**: Una sola base de datos con filtrado por tenant_id
- **Validaciones**: Todas las operaciones validan pertenencia al tenant
- **🚀 Quickstart**: Configuraciones independientes por tenant con historial

## Seguridad

- **Autenticación**: Bearer Token (JWT) - configurado en API Gateway
- **Autorización**: Control de acceso por tenant
- **Validación**: Validación estricta de entrada en todos los endpoints
- **SQL Injection**: Uso de prepared statements
- **CORS**: Configurado en API Gateway
- **🚀 Quickstart**: Validaciones de límites de recursos por tenant

## Monitoreo y Observabilidad

- **Métricas**: Prometheus en `/metrics`
- **Logs**: Structured logging con niveles
- **Health Check**: Endpoint `/health` para verificar estado
- **Tracing**: Logs de requests con correlation IDs
- **Dashboards**: Grafana para visualización
- **🚀 Quickstart**: Métricas específicas de adopción y performance

## 🚀 Tipos de Negocio Quickstart

### Completamente Implementado
- ✅ **Retail** - Comercio minorista con 5 categorías, 20 atributos, 10 variantes, 20 productos y 40+ marcas

### En Desarrollo
- 🔄 **Food & Beverage** - Alimentos y bebidas
- 🔄 **Fashion** - Moda y vestimenta
- 🔄 **Electronics** - Electrónicos y tecnología
- 🔄 **Automotive** - Automotriz y repuestos
- 🔄 **Sports & Fitness** - Deportes y fitness
- 🔄 **Health & Pharmacy** - Salud y farmacia
- 🔄 **Books & Media** - Libros y medios
- 🔄 **Home & Construction** - Hogar y construcción
- 🔄 **Beauty & Cosmetics** - Belleza y cosméticos
- 🔄 **Toys & Games** - Juguetes y juegos
- 🔄 **Pet Supplies** - Mascotas y suministros
- 🔄 **Office Supplies** - Oficina y papelería
- 🔄 **Jewelry & Accessories** - Joyería y accesorios

## Integración con Onboarding

El módulo quickstart se integra perfectamente con el sistema de onboarding multi-tenant:

- **Detección automática**: Basada en el tipo de industria seleccionado
- **Configuración opcional**: Los usuarios pueden elegir aplicar o no el quickstart
- **Aplicación selectiva**: Categorías, atributos, variantes, productos y marcas por separado
- **Time-to-value**: Reducción de días a minutos para tener un catálogo funcional

## Contribución

Para contribuir al proyecto:

1. Seguir la arquitectura hexagonal establecida
2. Implementar tests unitarios para nuevas funcionalidades
3. Documentar cambios en OpenAPI
4. Seguir convenciones de naming y estructura
5. Validar multi-tenancy en nuevos endpoints
6. **🚀 Para Quickstart**: Agregar nuevos tipos de negocio siguiendo la estructura YAML establecida

## Licencia

[Especificar licencia del proyecto] 