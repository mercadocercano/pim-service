# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 🏗️ Servicio PIM (Product Information Management)

Este es el servicio de gestión de información de productos del ecosistema SaaS Multitenant. Implementa arquitectura hexagonal y gestiona productos, categorías, atributos, marcas y catálogo global.

## 📦 Comandos de Desarrollo

### Desarrollo Local
```bash
# Ejecutar servicio localmente (puerto 8090)
go run main.go

# Ejecutar con hot reload
air

# Build del servicio
go build -o pim-service main.go

# Build optimizado para producción
go build -ldflags="-s -w" -o pim-service main.go
```

### Docker
```bash
# Build de desarrollo con hot reload
docker-compose build

# Ejecutar en Docker
docker-compose up -d

# Ver logs
docker-compose logs -f pim-service

# Ejecutar solo infraestructura (BD)
docker-compose up -d postgres mongodb
```

### Tests
```bash
# Tests unitarios
go test ./...

# Tests con coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Tests de integración
./run_tests.sh

# Tests específicos de integración
cd test-integration
./test-all-endpoints.sh          # Todos los endpoints
./test-marketplace-endpoints.sh   # Solo marketplace
./test-mongodb-endpoints.sh       # Solo MongoDB
```

### Migraciones y Seeds
```bash
# Ejecutar migraciones
./scripts/migrate.sh

# Cargar datos de Argentina
./scripts/seed_business_types_argentina.sh

# Seeds completos de marketplace
./scripts/seed_complete_marketplace_docker.sh

# Ejecutar seeds en Docker
./scripts/seed_business_types_docker.sh
```

## 🏛️ Arquitectura Hexagonal

### Estructura de Módulos
```
src/
├── api/                    # Infraestructura general de API
│   ├── config/            # Configuración general
│   ├── health/            # Health checks
│   └── infrastructure/    # Middlewares globales
│
├── attribute/             # Gestión de atributos
├── brand/                 # Gestión de marcas  
├── businesstype/          # Tipos de negocio
├── category/              # Categorías (tenant + marketplace)
├── category_attribute/    # Relación categoría-atributo
├── product/               # Gestión de productos
│   ├── tenant/           # Productos del tenant
│   ├── quickstart/       # Funcionalidad quickstart
│   └── global_catalog/   # Catálogo global
├── quickstart/           # Sistema de templates YAML
└── shared/               # Componentes compartidos
```

### Capas por Módulo
Cada módulo sigue la estructura:
```
módulo/
├── domain/               # Lógica de negocio pura
│   ├── entity/          # Entidades
│   ├── value_object/    # Value objects
│   ├── port/            # Interfaces (repositorios)
│   ├── service/         # Servicios de dominio
│   └── exception/       # Excepciones del dominio
│
├── application/          # Casos de uso
│   ├── usecase/         # Implementación de casos de uso
│   ├── request/         # DTOs de entrada
│   ├── response/        # DTOs de salida
│   └── mapper/          # Mappers DTO ↔ Entity
│
└── infrastructure/       # Implementaciones técnicas
    ├── controller/      # Controladores HTTP
    ├── persistence/     # Repositorios
    ├── config/          # Configuración del módulo
    └── criteria/        # Builders de criterios
```

## 🔧 Uso del MCP Go Generator

### Analizar Workflow del PIM
```bash
# Desde el directorio mcp-go-generator-node
analyzeUsecaseWorkflow --service_name="pim" --entity_name="product"
analyzeUsecaseWorkflow --service_name="pim" --entity_name="category"
analyzeUsecaseWorkflow --service_name="pim" --entity_name="attribute"
```

### Generar Componentes
```bash
# Generar DTOs
generateComponentByStep --step_type="dto" --entity_name="product_variant"

# Generar mapper
generateComponentByStep --step_type="mapper" --entity_name="product_variant"

# Generar caso de uso
generateComponentByStep --step_type="usecase" --entity_name="marketplace_brand"
```

### Generar Roadmap
```bash
generateWorkflowRoadmap --service_name="pim"
```

## 🛣️ API Endpoints Principales

### Productos
```
POST   /api/v1/products                         # Crear producto
GET    /api/v1/products                         # Listar con criteria
GET    /api/v1/products/:id                     # Obtener por ID
PUT    /api/v1/products/:id                     # Actualizar
DELETE /api/v1/products/:id                     # Eliminar (soft)
PATCH  /api/v1/products/:id/status              # Cambiar estado
GET    /api/v1/products/:id/status/transitions  # Transiciones disponibles
```

### Quickstart
```
POST   /api/v1/quickstart/setup                           # Setup inicial del tenant
GET    /api/v1/quickstart/business-types                  # Tipos de negocio disponibles
GET    /api/v1/quickstart/business-types/:id/products    # Productos por tipo
GET    /api/v1/quickstart/business-types/:id/categories  # Categorías por tipo
POST   /api/v1/quickstart/products/from-template         # Crear desde template
POST   /api/v1/quickstart/products/import-from-business-type # Importación masiva
```

### Catálogo Global
```
GET    /api/v1/public/global-catalog/search              # Búsqueda por EAN
GET    /api/v1/public/global-catalog/suggestions         # Sugerencias por tipo
GET    /api/v1/global-catalog/products                   # Listar productos globales
POST   /api/v1/global-catalog/products                   # Crear producto global (admin)
```

## 📊 Patrón Criteria

Todos los endpoints de listado soportan:
```
?page=1&page_size=10&sort_by=created_at&sort_dir=desc
?search=termo&status=active&category_id=uuid
```

Respuesta estándar:
```json
{
  "items": [...],
  "total_count": 100,
  "page": 1,
  "page_size": 10,
  "total_pages": 10
}
```

## 🔄 Estados de Productos

### Transiciones Permitidas
```
draft        → pending, active, deleted
pending      → draft, active, deleted  
active       → inactive, discontinued, deleted
inactive     → active, discontinued, deleted
discontinued → deleted
deleted      → (sin transiciones)
```

### Validaciones por Estado
- **draft**: Solo nombre requerido
- **pending**: Nombre + descripción + categoría
- **active**: Todos los campos + mínimo 1 variante activa
- **inactive**: Producto temporalmente deshabilitado
- **discontinued**: Descontinuado permanentemente
- **deleted**: Soft delete

## 🗄️ Bases de Datos

### PostgreSQL (Principal)
```bash
# Variables de entorno
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=pim_db
```

### MongoDB (Catálogo Global)
```bash
# Variables de entorno
MONGO_HOST=mongodb
MONGO_PORT=27017
MONGO_USER=admin
MONGO_PASSWORD=admin123
MONGO_DATABASE=pim_marketplace
```

## 🏢 Multi-Tenancy

- Todas las tablas incluyen `tenant_id` UUID
- Headers obligatorios: `X-Tenant-ID` y `Authorization`
- Datos globales sin tenant_id: business_types, global_products
- Validación automática en middlewares

## 📚 Documentación Adicional

### Arquitectura
- `documentation/ARCHITECTURE_FINAL.md` - Arquitectura completa
- `documentation/ROUTES_STRUCTURE.md` - Estructura de rutas
- `documentation/GLOBAL_CATALOG_DESIGN.md` - Diseño del catálogo

### Integración
- `documentation/QUICKSTART_INTEGRATION_GUIDE.md` - Guía quickstart
- `documentation/MARKETPLACE_INTEGRATION_STATUS.md` - Estado marketplace
- `documentation/curl-examples-global-catalog.md` - Ejemplos cURL

### Testing
- `documentation/TESTING_COMPLETION_SUMMARY.md` - Resumen de tests
- `documentation/test_coverage_report.md` - Reporte de cobertura

## 🚀 Flujo de Desarrollo Recomendado

1. **Usar MCP primero**: Verificar si el mcp-go-generator puede generar el componente
2. **Arquitectura hexagonal**: Mantener separación estricta de capas
3. **Tests primero**: Escribir tests antes de implementar
4. **Criteria pattern**: Usar para todos los listados
5. **Documentar APIs**: Actualizar OpenAPI con cada cambio

## ⚠️ Consideraciones Importantes

- El servicio corre en puerto **8090** (no 8080)
- MongoDB es requerido para el catálogo global
- Las migraciones son evolutivas, nunca destructivas
- Los estados de productos tienen reglas estrictas
- El quickstart usa templates YAML y datos del catálogo global