# 🏗️ Arquitectura Final del PIM Service

## **📊 Resumen de la Reorganización**

Después de la reorganización arquitectónica completa, el módulo **marketplace** ha sido **eliminado** y su funcionalidad distribuida correctamente en módulos especializados.

### **🎯 Cambios Principales**

| **Antes** | **Después** | **Beneficio** |
|-----------|-------------|---------------|
| ❌ `/marketplace` (confuso) | ✅ **ELIMINADO** | Arquitectura más clara |
| 🔄 Lógica mezclada | ✅ Módulos especializados | Separación de responsabilidades |
| 📂 Configuración centralizada | ✅ Configuración distribuida | Escalabilidad |

---

## **🏗️ Estructura de Módulos Final**

### **📁 Módulos de Dominio Principal**

```
src/
├── 🏷️ /category/              # Gestión de categorías (tenant + marketplace)
├── 🔧 /attribute/             # Gestión de atributos
├── 🔗 /category_attribute/    # Relación categoría-atributo
├── 📦 /product/               # Gestión de productos
├── 🏪 /brand/                 # Gestión de marcas
├── 🌐 /global_catalog/        # Catálogo global
├── 🏢 /businesstype/          # Tipos de negocio
├── 🚀 /quickstart/            # Configuración rápida
├── 🤖 /template_ai/           # Templates inteligentes con AI
├── 🛠️ /shared/               # Componentes compartidos
└── 🌐 /api/                   # Configuración general de API
```

---

## **🔗 Módulo category_attribute - El Corazón de la Relación**

### **📋 Responsabilidades**
- ✅ Gestionar relación entre categorías y atributos
- ✅ Validar valores permitidos por categoría
- ✅ Validaciones cruzadas (categorías + atributos)
- ✅ Gestión avanzada de valores de atributos

### **📂 Estructura**
```
category_attribute/
├── domain/
│   ├── entity/
│   │   ├── category_attribute.go        # Entidad principal
│   │   ├── marketplace_validator.go     # ← MIGRADO desde marketplace
│   │   └── attribute_value.go           # ← MIGRADO desde marketplace
│   └── port/
├── application/
│   ├── usecase/
│   ├── request/
│   └── response/
└── infrastructure/
    ├── controller/
    ├── persistence/
    └── config/
```

### **🛠️ Archivos Migrados**
| **Archivo** | **Origen** | **Razón de Migración** |
|-------------|------------|------------------------|
| `marketplace_validator.go` | `/marketplace/domain/entity/` | Valida categorías + atributos |
| `attribute_value.go` | `/marketplace/domain/entity/` | Gestiona valores de atributos |

---

## **🌐 Módulo api - Infraestructura General**

### **📋 Responsabilidades**
- ✅ Middlewares generales (CORS, Auth, Tenant)
- ✅ Health checks
- ✅ Configuración de API
- ✅ Documentación

### **📂 Estructura**
```
api/
├── config/                    # Configuración general
├── docs/                      # Documentación
├── health/                    # Health checks
├── infrastructure/
│   └── middleware/
│       └── middleware.go      # ← MIGRADO desde marketplace
└── monitoring/                # Métricas y monitoreo
```

### **🛠️ Archivos Migrados**
| **Archivo** | **Origen** | **Razón de Migración** |
|-------------|------------|------------------------|
| `middleware.go` | `/marketplace/infrastructure/controller/` | Middleware general de API |

---

## **🤖 Módulo template_ai - Templates Inteligentes con AI**

### **📋 Responsabilidades**
- ✅ Generar templates optimizados con inteligencia artificial
- ✅ Aplicar templates al catálogo del tenant
- ✅ Analizar performance y métricas de uso
- ✅ Aprender de feedback para mejora continua
- ✅ Integración con catálogo global

### **📂 Estructura**
```
template_ai/
├── domain/
│   ├── entity/
│   │   ├── ai_template.go              # Entidad principal de template AI
│   │   ├── template_global_product.go  # Relación template-producto global
│   │   └── performance_metric.go       # Métricas de rendimiento
│   ├── value_object/
│   │   └── generation_status.go        # Estados de generación
│   ├── port/
│   │   ├── ai_template_repository.go   # Interface repositorio
│   │   └── ai_generation_service.go    # Interface servicio AI
│   ├── service/
│   │   ├── template_engine.go          # Motor de templates
│   │   └── ai_template_domain_service.go
│   └── exception/
│       └── errors.go                   # Errores específicos del dominio
├── application/
│   ├── usecase/
│   │   ├── generate_smart_template_usecase.go
│   │   ├── apply_dynamic_template_usecase.go
│   │   ├── analyze_template_performance_usecase.go
│   │   └── update_template_from_feedback_usecase.go
│   ├── request/                        # DTOs de entrada
│   ├── response/                       # DTOs de salida
│   └── mapper/                         # Mappers entre capas
└── infrastructure/
    ├── controller/
    │   └── ai_template_controller.go   # Endpoints HTTP
    ├── persistence/
    │   └── repository/                 # Implementaciones PostgreSQL
    ├── client/
    │   └── ai_gateway_client.go        # Cliente para AI Gateway
    ├── service/
    │   └── ai_template_engine_impl.go  # Implementación del engine
    └── config/
        └── wire.go                     # Dependency injection
```

### **🎯 Características Principales**

1. **Generación Inteligente**
   - Analiza tipo de negocio y preferencias
   - Selecciona productos del catálogo global
   - Optimiza por región y demografía
   - Balancea marcas y categorías

2. **Aplicación Flexible**
   - Customización de precios
   - Ajuste de cantidades
   - Exclusión de productos
   - Creación automática de entidades

3. **Análisis Continuo**
   - Métricas de adopción
   - Satisfacción del usuario
   - Modificaciones comunes
   - ROI y crecimiento

4. **Aprendizaje Adaptativo**
   - Feedback de usuarios
   - Optimización automática
   - Mejora continua
   - Adaptación regional

---

## **🛣️ Mapeo Completo de Rutas por Módulo**

### **🏷️ Categorías Tenant (`/api/v1/categories`)**
> **Propósito**: Categorías específicas de cada tenant  
> **Requiere**: `X-Tenant-ID` header  
> **Acceso**: Usuarios del tenant

```
Base: /api/v1/categories
├── POST   /                    # Crear categoría del tenant
├── GET    /                    # Listar con criterios avanzados
├── GET    /simple              # Listado simple
├── GET    /tree                # Vista de árbol jerárquico
├── GET    /:id                 # Obtener por ID
├── PUT    /:id                 # Actualizar categoría
├── PATCH  /:id/activate        # Activar categoría
├── PATCH  /:id/deactivate      # Desactivar categoría
└── PATCH  /:id/move            # Mover en jerarquía
```

### **🌐 Categorías Marketplace (`/api/v1/marketplace/categories`)**
> **Propósito**: Categorías globales del marketplace (plantillas)  
> **Requiere**: `X-User-Role: marketplace_admin` o `super_admin`  
> **Acceso**: Solo administradores del marketplace

```
Base: /api/v1/marketplace
├── GET    /categories          # Listar categorías globales
├── POST   /categories          # Crear categoría global
├── PUT    /categories/:id      # Actualizar categoría global
├── POST   /categories/validate-hierarchy  # Validar jerarquía
├── POST   /sync-changes        # Sincronizar cambios
└── GET    /taxonomy            # Obtener taxonomía tenant
```

### **🔧 Atributos Tenant (`/api/v1/attributes`)**
> **Propósito**: Atributos básicos del tenant  
> **Requiere**: `X-Tenant-ID` header  
> **Acceso**: Usuarios del tenant  
> **Estado**: ✅ IMPLEMENTADO

```
Base: /api/v1/attributes
├── POST   /                    # Crear atributo básico
├── GET    /                    # Listar atributos
├── GET    /:id                 # Obtener por ID
├── PUT    /:id                 # Actualizar atributo
└── DELETE /:id                 # Eliminar atributo
```

### **🎯 Atributos Personalizados Tenant (`/api/v1/tenant/custom-attributes`)**
> **Propósito**: Atributos personalizados/extendidos del tenant  
> **Requiere**: `X-Tenant-ID` header  
> **Acceso**: Usuarios del tenant  
> **Estado**: ✅ IMPLEMENTADO

```
Base: /api/v1/tenant/custom-attributes
├── POST   /                    # Crear atributo personalizado
├── GET    /                    # Listar atributos personalizados
├── PUT    /:attribute_id       # Actualizar atributo
└── DELETE /:attribute_id       # Eliminar atributo
```

### **🌐 Atributos Marketplace**
> **Propósito**: Atributos globales del marketplace (plantillas)  
> **Requiere**: `X-User-Role: marketplace_admin` o `super_admin`  
> **Acceso**: Solo administradores del marketplace  
> **Estado**: ✅ IMPLEMENTADO

```
Base: /api/v1/marketplace/attributes
├── POST   /                    # Crear atributo global
├── GET    /                    # Listar atributos globales
├── PUT    /:id                 # Actualizar atributo global
└── DELETE /:id                 # Eliminar atributo global
```

### **🔗 Category-Attribute Module**
```
Base: /api/v1/category-attributes
├── GET    /                    # Listar con filtros y paginación
├── GET    /simple              # Listado simple
├── POST   /                    # Crear relación
├── PUT    /:id                 # Actualizar valores permitidos
└── DELETE /:id                 # Eliminar relación
```

### **📦 Product Module**
```
Base: /api/v1/products
├── POST   /                    # Crear producto
├── GET    /                    # Listar productos
├── GET    /:id                 # Obtener por ID
├── PUT    /:id                 # Actualizar
├── DELETE /:id                 # Eliminar
└── Variants:
    ├── POST   /:product_id/variants           # Crear variante
    ├── GET    /:product_id/variants           # Listar variantes
    ├── GET    /:product_id/variants/:variant_id  # Obtener variante
    ├── PUT    /:product_id/variants/:variant_id  # Actualizar variante
    ├── DELETE /:product_id/variants/:variant_id  # Eliminar variante
    └── GET    /variants                       # Listar todas las variantes
```

### **🏪 Brand Module**
```
Base: /api/v1/brands
├── POST   /                    # Crear marca
├── GET    /                    # Listar marcas
├── GET    /:id                 # Obtener por ID
├── PUT    /:id                 # Actualizar
└── DELETE /:id                 # Eliminar
```

### **🌐 Global Catalog Module**
```
Público: /api/v1/public/global-catalog
├── GET    /health              # Health check
├── GET    /search              # Buscar por EAN
├── GET    /suggestions         # Sugerencias de productos
└── GET    /products/ean/:ean   # Producto por EAN

Privado: /api/v1/global-catalog
├── POST   /products            # Crear producto
├── GET    /products            # Listar productos
└── GET    /products/search     # Buscar por EAN
```

### **🏢 Business Type Module**
```
Base: /api/v1/business-types
├── POST   /                    # Crear tipo de negocio
├── GET    /                    # Listar tipos
├── GET    /:id                 # Obtener por ID
├── PUT    /:id                 # Actualizar
└── DELETE /:id                 # Eliminar
```

### **🚀 Quickstart Module**
```
Base: /api/v1/quickstart
├── GET    /business-types      # Tipos de negocio disponibles
├── GET    /categories/:businessType    # Categorías por tipo
├── GET    /attributes/:businessType    # Atributos por tipo
├── GET    /variants/:businessType      # Variantes por tipo
├── GET    /products/:businessType      # Productos por tipo
├── GET    /brands/:businessType        # Marcas por tipo
└── POST   /setup               # Configuración completa
```

### **🤖 AI Templates Module**
```
Base: /api/v1/templates
├── POST   /generate            # Generar template inteligente con AI
├── POST   /:id/apply           # Aplicar template al catálogo
├── GET    /:id/performance     # Obtener métricas de rendimiento
└── POST   /update-from-feedback # Actualizar template con feedback
```

---

## **🔒 Seguridad y Headers**

### **🛡️ Headers Requeridos**

#### **Para Operaciones Tenant**
```
X-Tenant-ID: {uuid}           # Obligatorio para categorías tenant
X-User-Role: {role}           # Para autorización
Content-Type: application/json # Para POST/PUT
```

#### **Para Operaciones Marketplace**
```
X-User-Role: marketplace_admin # O super_admin (obligatorio)
Content-Type: application/json # Para POST/PUT
```

### **👤 Roles Soportados**
- `super_admin`: Acceso completo
- `marketplace_admin`: Gestión de marketplace + categorías globales
- `tenant_admin`: Administración del tenant + categorías tenant
- `tenant_user`: Usuario del tenant + categorías tenant

---

## **📊 Diferencias Conceptuales Clave**

### **🏷️ Categorías: Tenant vs Marketplace**

| **Aspecto** | **Categorías Tenant** | **Categorías Marketplace** |
|-------------|----------------------|---------------------------|
| **Propósito** | Categorías específicas del tenant | Plantillas globales del marketplace |
| **Ruta** | `/api/v1/categories` | `/api/v1/marketplace/categories` |
| **Header** | `X-Tenant-ID` requerido | `X-User-Role: marketplace_admin` |
| **Acceso** | Usuarios del tenant | Solo administradores |
| **Scope** | Por tenant | Global (todos los tenants) |
| **Personalización** | Completa libertad | Definidas por marketplace |
| **Jerarquía** | Independiente por tenant | Jerarquía global |

### **🔧 Atributos: Tenant vs Personalizados vs Marketplace**

| **Aspecto** | **Atributos Tenant** | **Atributos Personalizados** | **Atributos Marketplace** |
|-------------|---------------------|----------------------------|--------------------------|
| **Propósito** | Atributos básicos del tenant | Extensiones personalizadas | Plantillas globales |
| **Ruta** | `/api/v1/attributes` | `/api/v1/tenant/custom-attributes` | `/api/v1/marketplace/attributes` |
| **Header** | `X-Tenant-ID` | `X-Tenant-ID` | `X-User-Role: marketplace_admin` |
| **Acceso** | Usuarios del tenant | Usuarios del tenant | Solo administradores |
| **Scope** | Por tenant | Por tenant | Global (todos los tenants) |
| **Estado** | ✅ IMPLEMENTADO | ✅ IMPLEMENTADO | ✅ IMPLEMENTADO |
| **Funcionalidad** | CRUD básico | Extensiones y personalizaciones | Plantillas y configuración global |

### **🔄 Flujo de Trabajo**

#### **📋 Para Categorías**
1. **Administradores Marketplace**:
   - Crean categorías globales en `/marketplace/categories`
   - Definen jerarquías y plantillas
   - Los tenants pueden usar estas como base

2. **Tenants**:
   - Crean sus propias categorías en `/categories`
   - Pueden basarse en plantillas marketplace
   - Personalizan según sus necesidades

#### **🔧 Para Atributos**
1. **Administradores Marketplace** ✅:
   - Crean atributos globales en `/marketplace/attributes`
   - Definen tipos y validaciones estándar
   - Los tenants pueden usar como base

2. **Tenants**:
   - Crean atributos básicos en `/attributes` ✅
   - Extienden con atributos personalizados en `/tenant/custom-attributes` ✅
   - Personalizan según sus necesidades específicas

---

## **📊 Filtros y Paginación**

### **🔍 Parámetros Estándar**
```
?page=1                       # Número de página
&page_size=10                 # Elementos por página (max: 100)
&sort_by=created_at           # Campo de ordenamiento
&sort_dir=desc                # Dirección (asc/desc)
```

### **🎯 Filtros Específicos por Módulo**

#### **Categories (Tenant)**
```
?status=ACTIVE                # Estado
&parent_id=uuid               # ID del padre
&name=texto                   # Búsqueda por nombre
&active=true                  # Solo activas
&root_only=true               # Solo raíz
```

#### **Categories (Marketplace)**
```
?include_inactive=true        # Incluir inactivas
&include_marketplace_data=true # Incluir datos marketplace
&format=tree                  # Formato (tree/flat)
```

#### **Attributes (Tenant)**
```
?type=string                  # Tipo de atributo
&required=true                # Solo obligatorios
&active=true                  # Solo activos
&searchable=true              # Solo buscables
```

#### **Tenant Custom Attributes**
```
?include_inactive=true        # Incluir inactivos
&marketplace_category_id=uuid # Por categoría marketplace
&attribute_type=string        # Tipo de atributo
&is_filterable=true           # Solo filtrables
&is_searchable=true           # Solo buscables
```

#### **Marketplace Attributes**
```
?type=string                  # Tipo de atributo
&category_id=uuid             # Por categoría
&required=true                # Solo obligatorios
&global=true                  # Solo globales
```

#### **Category-Attributes**
```
?tenant_id=uuid               # ID del tenant
&category_id=uuid             # ID de categoría
&attribute_id=uuid            # ID de atributo
&allowed_value=valor          # Contiene valor específico
```

#### **Products**
```
?status=active                # Estado
&category_id=uuid             # Categoría
&brand_id=uuid                # Marca
&sku=codigo                   # SKU
&name=texto                   # Nombre
```

---

## **✅ Beneficios de la Arquitectura Final**

### **🎯 Separación de Responsabilidades**
- ✅ Cada módulo tiene un propósito específico
- ✅ No hay solapamiento de funcionalidades
- ✅ Fácil localización de código
- ✅ **Distinción clara entre tenant y marketplace**

### **📈 Escalabilidad**
- ✅ Fácil agregar nuevos módulos
- ✅ Configuración distribuida
- ✅ Independencia entre módulos

### **🧪 Testabilidad**
- ✅ Módulos independientes
- ✅ Casos de uso aislados
- ✅ Mocking simplificado

### **🔧 Mantenibilidad**
- ✅ Código organizado
- ✅ Responsabilidades claras
- ✅ Arquitectura hexagonal consistente

---

## **🚀 Estado de Compilación**

```bash
✅ go build ./...              # Compilación exitosa
✅ Estructura validada         # Sin dependencias circulares
✅ Tests preparados            # Listos para actualización
```

---

## **📝 Próximos Pasos**

1. ✅ **Estructura reorganizada** - COMPLETADO
2. ✅ **OpenAPI actualizado** - COMPLETADO
3. 🧪 **Corregir tests** - EN PROGRESO
4. 📊 **Validar endpoints** - PENDIENTE
5. 🚀 **Deploy y testing** - PENDIENTE

La arquitectura del PIM Service ahora es **robusta, escalable y mantenible** con **distinción clara entre operaciones tenant y marketplace**. ¡Lista para continuar con el desarrollo! 🎯 