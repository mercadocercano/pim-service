# 🗺️ MAPA ARQUITECTÓNICO DE ENDPOINTS - PIM SERVICE

## 📐 Arquitectura de Alto Nivel

```
┌─────────────────────────────────────────────────────────────────┐
│                         KONG API GATEWAY                         │
│                     http://localhost:8001/pim/                   │
└────────────────────────────┬────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      PIM SERVICE (Go + Gin)                      │
│                     http://localhost:8090                        │
│                                                                   │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │  Controllers │  │  Use Cases   │  │   Domain     │          │
│  │   (HTTP)     │→ │ (Business    │→ │  (Entities)  │          │
│  │              │  │   Logic)     │  │              │          │
│  └──────────────┘  └──────────────┘  └──────────────┘          │
│         │                   │                  │                 │
│         ▼                   ▼                  ▼                 │
│  ┌──────────────────────────────────────────────────┐          │
│  │           Infrastructure Layer                    │          │
│  │  - PostgreSQL Repositories                        │          │
│  │  - MongoDB for Global Catalog                    │          │
│  │  - Criteria Builders                              │          │
│  └──────────────────────────────────────────────────┘          │
└─────────────────────────────────────────────────────────────────┘
                             │
                ┌────────────┴────────────┐
                ▼                         ▼
      ┌──────────────────┐    ┌──────────────────┐
      │   PostgreSQL     │    │     MongoDB      │
      │   (Relational)   │    │  (Global Catalog)│
      └──────────────────┘    └──────────────────┘
```

---

## 🎯 Módulos y sus Endpoints

### 1. **HEALTH & MONITORING**

```
/health ─────────────────────► HealthHandler
    │
    ├─► Check PostgreSQL
    ├─► Check MongoDB
    └─► Response: {"status": "healthy"}

/metrics ────────────────────► Prometheus Handler
    │
    └─► Métricas de performance
```

---

### 2. **BUSINESS TYPES** (Tipos de Negocio)

```
/api/v1/business-types/
    │
    ├─► POST    ""           CreateBusinessTypeUseCase
    ├─► GET     ""           ListBusinessTypesUseCase (con Criteria)
    ├─► GET     "/:id"       GetBusinessTypeUseCase
    ├─► PUT     "/:id"       UpdateBusinessTypeUseCase
    └─► DELETE  "/:id"       ❌ Not Implemented
```

**Casos de Uso**:
- Crear tipos de negocio (ej: "Ferretería", "Almacén")
- Listar con filtros (is_active, search)
- Ordenar por `sort_order`

---

### 3. **CATEGORIES** (Categorías con Jerarquía)

```
/api/v1/categories/
    │
    ├─► POST    ""                CreateCategoryUseCase
    ├─► GET     ""                ListCategoriesByCriteriaUseCase
    ├─► GET     "/simple"         ListCategoriesUseCase
    ├─► GET     "/tree"           ListCategoriesTreeUseCase
    ├─► GET     "/:id"            GetCategoryByIDUseCase
    ├─► PUT     "/:id"            UpdateCategoryUseCase
    ├─► PATCH   "/:id/activate"   ActivateCategoryUseCase
    ├─► PATCH   "/:id/deactivate" DeactivateCategoryUseCase
    └─► PATCH   "/:id/move"       MoveCategoryUseCase (con validación de ciclos)
```

**Jerarquía de Categorías**:
```
Herramientas (nivel 0)
  ├── Eléctricas (nivel 1)
  │   ├── Taladros (nivel 2)
  │   └── Amoladoras (nivel 2)
  └── Manuales (nivel 1)
      ├── Llaves (nivel 2)
      └── Destornilladores (nivel 2)
```

---

### 4. **CATEGORY-ATTRIBUTES** (Relación Categoría ↔ Atributos)

```
/api/v1/category-attributes/
    │
    ├─► GET     ""          ListCategoryAttributesByCriteriaUseCase
    ├─► GET     "/simple"   GetCategoryAttributesUseCase
    ├─► GET     "/detailed" GetDetailedCategoryAttributesUseCase
    ├─► POST    ""          CreateCategoryAttributeUseCase
    ├─► PUT     "/:id"      UpdateCategoryAttributeUseCase
    └─► DELETE  "/:id"      DeleteCategoryAttributeUseCase
```

**Ejemplo de Uso**:
```json
{
  "category_id": "taladros-uuid",
  "attribute_id": "potencia-uuid",
  "allowed_values": ["500W", "750W", "1000W", "1500W"]
}
```

---

### 5. **ATTRIBUTES** (Atributos Globales del Marketplace)

```
/api/v1/marketplace/attributes/
    │
    ├─► GET     ""      ListAttributesUseCase
    ├─► POST    ""      CreateAttributeUseCase
    ├─► GET     "/:id"  GetAttributeByIDUseCase
    ├─► PUT     "/:id"  UpdateAttributeUseCase
    └─► DELETE  "/:id"  DeleteAttributeUseCase
```

**Tipos de Atributos**:
- `text`: Descripción libre
- `number`: Peso, dimensiones
- `boolean`: Tiene garantía?
- `select`: Color (Rojo, Azul, Verde)
- `multi_select`: Materiales (Acero, Plástico, Goma)

---

### 6. **BRANDS** (Marcas)

#### Tenant Brands (Por Tienda)

```
/api/v1/brands/
    │
    ├─► POST    ""      CreateBrandUseCase
    ├─► GET     ""      ListBrandsUseCase (con filtros)
    ├─► GET     "/:id"  GetBrandUseCase
    ├─► PUT     "/:id"  UpdateBrandUseCase
    └─► DELETE  "/:id"  DeleteBrandUseCase
```

#### Marketplace Brands (Globales)

```
/api/v1/marketplace-brands/
    │
    ├─► GET     ""             List
    ├─► POST    ""             Create
    ├─► GET     "/:id"         GetByID
    ├─► PUT     "/:id"         Update
    ├─► DELETE  "/:id"         Delete
    ├─► PUT     "/:id/verify"  VerifyBrand (validación oficial)
    └─► PUT     "/:id/unverify" UnverifyBrand
```

---

### 7. **PRODUCTS** (Productos del Tenant)

```
/api/v1/products/
    │
    ├─── CRUD ───────────────────────────────┐
    │   ├─► POST    ""          CreateProduct
    │   ├─► GET     ""          ListProducts (con Criteria)
    │   ├─► GET     "/:id"      GetProduct
    │   ├─► PUT     "/:id"      UpdateProduct
    │   └─► DELETE  "/:id"      DeleteProduct (soft)
    │
    ├─── STATUS ─────────────────────────────┐
    │   ├─► PATCH   "/:id/status"            UpdateStatus
    │   └─► GET     "/:id/status/transitions" GetTransitions
    │
    ├─── VALIDATION ─────────────────────────┐
    │   ├─► POST    "/validate-schema"       ValidateCSV
    │   ├─► POST    "/apply-mapping"         ApplyMapping
    │   ├─► GET     "/csv-template"          GetTemplate
    │   └─► GET     "/json-template"         GetJSONTemplate
    │
    └─── IMPORT (HITO 2.1) ──────────────────┐
        └─► POST    "/import"                BulkImport (CSV/JSON)
```

**Transiciones de Estado**:
```
┌────────┐      ┌─────────┐      ┌────────┐
│ draft  │─────►│ pending │─────►│ active │
└────────┘      └─────────┘      └────────┘
    │               │                 │
    │               │                 ▼
    │               │           ┌──────────┐
    │               └──────────►│ inactive │
    │                           └──────────┘
    │                                 │
    ▼                                 ▼
┌─────────────┐               ┌──────────────┐
│  deleted    │               │ discontinued │
└─────────────┘               └──────────────┘
```

---

### 8. **PRODUCT VARIANTS** (Variantes - HITO 2.1)

```
/api/v1/products/:id/variants/
    │
    ├─► POST    ""           CreateVariant
    ├─► GET     ""           ListVariants
    ├─► GET     "/:variant_id" GetVariant
    ├─► PUT     "/:variant_id" UpdateVariant
    └─► DELETE  "/:variant_id" DeleteVariant

/api/v1/variants/
    │
    └─► GET     ""           ListAllVariants (standalone)
```

**Ejemplo de Producto con Variantes**:
```json
{
  "product_name": "Taladro Bosch",
  "category": "Herramientas Eléctricas",
  "brand": "Bosch",
  "variants": [
    {
      "sku": "TALADRO-BOSCH-500W",
      "attributes": {"potencia": "500W", "voltaje": "220V"},
      "price": 35000
    },
    {
      "sku": "TALADRO-BOSCH-750W",
      "attributes": {"potencia": "750W", "voltaje": "220V"},
      "price": 45000
    },
    {
      "sku": "TALADRO-BOSCH-1000W",
      "attributes": {"potencia": "1000W", "voltaje": "220V"},
      "price": 65000
    }
  ]
}
```

---

### 9. **GLOBAL CATALOG** (Catálogo Global de Templates)

```
/api/v1/public/global-catalog/  (Sin autenticación)
    │
    ├─► GET     "/health"                HealthCheck
    ├─► GET     "/search?ean={ean}"      SearchByEAN
    ├─► GET     "/suggestions?business_type={type}" Suggestions
    └─► GET     "/products/ean/:ean"     GetByEAN

/api/v1/global-catalog/  (Privado - Admin)
    │
    ├─► POST    "/products"              CreateGlobalProduct
    ├─► GET     "/products"              ListGlobalProducts
    ├─► GET     "/products/search"       SearchByEAN
    ├─► GET     "/products/:id"          GetByID
    ├─► PUT     "/products/:id"          Update
    └─► DELETE  "/products/:id"          Delete
```

**Uso del Catálogo Global**:
1. Scraper service busca productos por EAN
2. Si existe en global catalog → obtener datos completos
3. Si no existe → crear entrada nueva para futuros tenants

---

### 🔟 **QUICKSTART** (Configuración Rápida - HITO 2)

```
/api/v1/quickstart/
    │
    ├─── TEMPLATES ──────────────────────┐
    │   ├─► GET  "/templates"            ListTemplates
    │   ├─► POST "/apply"                ApplyTemplate (body)
    │   └─► POST "/templates/:id/apply"  ApplyTemplateByID (path)
    │
    ├─── CONSULTAS POR BUSINESS TYPE ───┐
    │   ├─► GET  "/business-types"       GetBusinessTypes
    │   ├─► GET  "/categories/:type"     GetCategories
    │   ├─► GET  "/attributes/:type"     GetAttributes
    │   ├─► GET  "/variants/:type"       GetVariants
    │   ├─► GET  "/products/:type"       GetProducts
    │   └─► GET  "/brands/:type"         GetBrands
    │
    ├─── SETUP ──────────────────────────┐
    │   └─► POST "/setup"                SetupTenant
    │
    └─── LEGACY (Products) ──────────────┐
        ├─► POST "/products/from-template"       CreateFromTemplate
        ├─► POST "/products/import-from-business-type" ImportBulk
        └─► GET  "/progress"                     GetProgress
```

**Flujo Quickstart**:
```
1. Usuario selecciona "Ferretería"
   GET /quickstart/templates?business_type=ferreteria

2. Frontend muestra preview del template
   GET /quickstart/template/{ferreteria-id}

3. Usuario confirma
   POST /quickstart/apply
   {
     "template_id": "ferreteria-basica",
     "tenant_id": "mi-ferreteria-uuid"
   }

4. Sistema crea automáticamente:
   ✅ 15 categorías
   ✅ 25 marcas
   ✅ Estructura base lista en <10 segundos
```

---

### 1️⃣1️⃣ **WIZARD** (Wizard de Configuración)

```
/api/v1/wizard/
    │
    ├─► GET     "/status"                    GetWizardStatus
    ├─► POST    "/start"                     StartWizard
    ├─► PUT     "/step"                      UpdateStep
    ├─► GET     "/template/:businessTypeId"  GetTemplate
    ├─► GET     "/template/:businessTypeId/:section" GetSection
    └─► DELETE  "/reset"                     Reset (⚠️ TEMPORAL)
```

**Steps del Wizard**:
```
Step 1: Selección de Business Type
Step 2: Personalización de Categorías
Step 3: Selección de Marcas
Step 4: Configuración de Atributos
Step 5: Confirmación y Apply
```

---

### 1️⃣2️⃣ **AI TEMPLATES** (Templates con IA)

```
/api/v1/templates/
    │
    ├─► POST    "/generate"                  GenerateSmartTemplate (IA)
    ├─► POST    "/:id/apply"                 ApplyDynamicTemplate
    ├─► GET     "/:id/performance"           AnalyzePerformance
    └─► POST    "/update-from-feedback"      UpdateFromFeedback
```

**Generación con IA**:
```json
POST /templates/generate
{
  "business_type_id": "ferreteria-uuid",
  "target_size": "medium",
  "preferences": {
    "price_range": "standard",
    "include_generics": true,
    "generic_percentage": 25,
    "categories_focus": ["herramientas", "electricidad"],
    "regional_preferences": "buenos_aires"
  }
}

Response:
{
  "template_id": "generated-uuid",
  "products": [...], // 150 productos optimizados
  "summary": {
    "total_products": 150,
    "categories": 12,
    "brands": 20,
    "estimated_investment": 2500000
  }
}
```

---

## 🔄 FLUJOS DE NEGOCIO COMPLETOS

### Flujo 1: Onboarding de Nuevo Tenant (Quickstart)

```
┌─────────────┐
│  Usuario    │
│  se registra│
└──────┬──────┘
       │
       ▼
┌─────────────────────────────────┐
│ Selecciona tipo de negocio      │
│ GET /business-types              │
│ → ["Ferretería", "Almacén", ...] │
└──────┬──────────────────────────┘
       │
       ▼
┌──────────────────────────────────┐
│ Ve template sugerido             │
│ GET /quickstart/templates        │
│ ?business_type=ferreteria        │
└──────┬───────────────────────────┘
       │
       ▼
┌──────────────────────────────────┐
│ Aplica template                  │
│ POST /quickstart/apply           │
│ Body: {template_id, tenant_id}   │
└──────┬───────────────────────────┘
       │
       ▼
┌──────────────────────────────────┐
│ Sistema crea automáticamente:    │
│ ✅ Categorías                     │
│ ✅ Marcas                         │
│ ✅ Atributos base                 │
└──────┬───────────────────────────┘
       │
       ▼
┌──────────────────────────────────┐
│ Tenant listo para cargar         │
│ productos en <10 minutos          │
└──────────────────────────────────┘
```

---

### Flujo 2: Importación Masiva de Productos con Variantes

```
┌─────────────┐
│   CSV File  │
│ (con variantes)
└──────┬──────┘
       │
       ▼
┌─────────────────────────────────┐
│ 1. Validar Schema                │
│ POST /products/validate-schema   │
│ → {is_valid, can_import, errors} │
└──────┬──────────────────────────┘
       │
       ▼
┌─────────────────────────────────┐
│ 2. Corregir mapeos (si necesario)│
│ POST /products/apply-mapping     │
│ Body: {validation_id, mappings}  │
└──────┬──────────────────────────┘
       │
       ▼
┌─────────────────────────────────┐
│ 3. Importar productos            │
│ POST /products/import            │
│ File: multipart/form-data        │
└──────┬──────────────────────────┘
       │
       ▼
┌─────────────────────────────────┐
│ Sistema procesa:                 │
│ ✅ Crea productos                 │
│ ✅ Crea variantes                 │
│ ✅ Auto-crea categorías/marcas    │
│ ✅ Reporta errores por fila       │
└──────────────────────────────────┘
```

**Ejemplo CSV**:
```csv
product_name,variant_sku,category,brand,potencia,voltaje,price
Taladro Bosch,TALADRO-500W,Herramientas,Bosch,500W,220V,35000
Taladro Bosch,TALADRO-750W,Herramientas,Bosch,750W,220V,45000
Taladro Bosch,TALADRO-1000W,Herramientas,Bosch,1000W,220V,65000
Amoladora Makita,AMOLADORA-900W,Herramientas,Makita,900W,220V,55000
```

**Sistema detecta automáticamente**:
- "Taladro Bosch" = 1 producto con 3 variantes
- "Amoladora Makita" = 1 producto con 1 variante
- Atributos dinámicos: `potencia`, `voltaje`
- Auto-crea categoría "Herramientas" si no existe
- Auto-crea marcas "Bosch" y "Makita" si no existen

---

### Flujo 3: Búsqueda de Producto Global por EAN

```
┌─────────────┐
│   Scraper   │
│  Service    │
└──────┬──────┘
       │
       ▼
┌─────────────────────────────────────┐
│ Buscar producto por EAN              │
│ GET /public/global-catalog/search    │
│ ?ean=7790315108901                   │
└──────┬──────────────────────────────┘
       │
       ├─► ✅ Existe → Retorna datos completos
       │   {
       │     "name": "Coca Cola 2.25L",
       │     "brand": "Coca Cola",
       │     "category": "Bebidas",
       │     "attributes": {...}
       │   }
       │
       └─► ❌ No existe → 404
           → Scraper crea entrada nueva
           → POST /global-catalog/products
```

---

## 🎨 DISEÑO DE RESPUESTAS

### Respuesta Paginada Estándar

```json
{
  "items": [...],
  "total_count": 150,
  "page": 1,
  "page_size": 10,
  "total_pages": 15
}
```

### Respuesta de Error Estándar

```json
{
  "error": "Producto no encontrado",
  "code": "PRODUCT_NOT_FOUND",
  "details": {
    "product_id": "abc-123-xyz"
  }
}
```

### Respuesta de Importación

```json
{
  "success": true,
  "summary": {
    "products_created": 25,
    "variants_created": 78,
    "categories_created": 3,
    "brands_created": 5,
    "products_skipped": 2,
    "errors": [
      {
        "row": 15,
        "error": "SKU duplicado: TALADRO-500W"
      }
    ]
  }
}
```

---

## 🔐 HEADERS REQUERIDOS

### Multi-Tenancy

```
X-Tenant-ID: <uuid>
```

**Obligatorio en**:
- Todos los endpoints de tenant-specific resources (products, brands, categories)
- No requerido en endpoints públicos (`/public/global-catalog`)

### Autorización

```
Authorization: Bearer <jwt_token>
```

**Roles**:
- `super_admin`: Acceso completo
- `marketplace_admin`: Gestión de marketplace (global catalog, business types)
- `tenant_admin`: Administración del tenant
- `tenant_user`: Usuario estándar del tenant

### User Role (Opcional)

```
X-User-Role: super_admin|marketplace_admin|tenant_admin|tenant_user
```

---

## 📊 TABLA DE COMPATIBILIDAD

| Endpoint | Requiere X-Tenant-ID | Requiere Auth | Rol Mínimo |
|----------|---------------------|---------------|------------|
| `/health` | ❌ | ❌ | - |
| `/public/global-catalog/*` | ❌ | ❌ | - |
| `/products` | ✅ | ✅ | tenant_user |
| `/brands` | ✅ | ✅ | tenant_user |
| `/categories` | ✅ | ✅ | tenant_user |
| `/business-types` (GET) | ❌ | ✅ | tenant_user |
| `/business-types` (POST/PUT) | ❌ | ✅ | marketplace_admin |
| `/global-catalog/products` | ❌ | ✅ | marketplace_admin |
| `/quickstart/*` | ✅ | ✅ | tenant_admin |
| `/wizard/*` | ✅ | ✅ | tenant_admin |

---

## 🚀 PERFORMANCE TIPS

### 1. Usar Criterios para Filtros Complejos

**❌ Malo**:
```
GET /products?name=taladro&brand=bosch&category=herramientas&min_price=1000
```

**✅ Bueno**:
```
GET /products?page=1&page_size=20&sort_by=created_at&sort_dir=desc
+ Body (POST a /products/search):
{
  "filters": {
    "name": {"like": "taladro"},
    "brand_id": {"eq": "bosch-uuid"},
    "category_id": {"eq": "herramientas-uuid"},
    "price": {"gte": 1000}
  }
}
```

### 2. Usar Endpoints `/simple` para UIs Simples

```
GET /categories/simple  → Sin paginación, rápido para dropdowns
GET /categories         → Con paginación y filtros, para tablas grandes
```

### 3. Cachear Business Types y Global Catalog

- Business Types cambian raramente → TTL de 1 hora
- Global Catalog es read-heavy → CDN o Redis

---

**Documento generado**: 2026-02-05  
**Versión**: 1.0  
**Autor**: Cursor AI Assistant
