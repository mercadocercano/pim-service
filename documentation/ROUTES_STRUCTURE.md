# Estructura de Rutas del PIM Service

## 🎯 Organización por Módulos

### **1. Productos del Tenant** (`/api/v1/products`)
**Módulo**: `product/tenant`  
**Propósito**: Gestión de productos específicos del tenant

```
POST   /api/v1/products                           # Crear producto
GET    /api/v1/products                           # Listar productos del tenant
GET    /api/v1/products/{id}                      # Obtener producto por ID
PUT    /api/v1/products/{id}                      # Actualizar producto
DELETE /api/v1/products/{id}                      # Eliminar producto

# Estados flexibles
PATCH  /api/v1/products/{id}/status               # Cambiar estado del producto
GET    /api/v1/products/{id}/status/transitions   # Obtener transiciones disponibles
```

**Headers requeridos**: `X-Tenant-ID`

---

### **2. Catálogo Global** (`/api/v1/global-catalog`)
**Módulo**: `product/global_catalog`  
**Propósito**: Productos template para quickstart

```
# Rutas públicas (sin autenticación)
GET    /api/v1/public/global-catalog/health                          # Health check
GET    /api/v1/public/global-catalog/search?ean={ean}                # Búsqueda por EAN
GET    /api/v1/public/global-catalog/suggestions                     # Sugerencias por tipo de negocio
GET    /api/v1/public/global-catalog/products/ean/{ean}              # Producto por EAN

# Rutas privadas (administración)
POST   /api/v1/global-catalog/products                               # Crear producto global
GET    /api/v1/global-catalog/products                               # Listar productos globales
GET    /api/v1/global-catalog/products/search?ean={ean}              # Búsqueda avanzada
GET    /api/v1/global-catalog/business-types/{id}/products           # Productos por tipo de negocio
```

**Headers requeridos**: Ninguno para rutas públicas, autenticación para privadas

---

### **3. Quickstart** (`/api/v1/quickstart`)
**Módulo**: `product/quickstart`  
**Propósito**: Funcionalidades de configuración inicial rápida

```
# Productos desde templates
POST   /api/v1/quickstart/products/from-template                     # Crear producto desde template
POST   /api/v1/quickstart/products/import-from-business-type         # Importación masiva

# Progreso del quickstart
GET    /api/v1/quickstart/progress                                   # Estado del quickstart
```

**Headers requeridos**: `X-Tenant-ID`

---

## 🔄 Flujo de Negocio

### **Configuración Inicial del Tenant**
```
1. Tenant se registra y selecciona business_type
2. GET /global-catalog/business-types/{id}/products → Ve productos sugeridos
3. POST /quickstart/products/from-template → Crea productos en estado "draft"
4. PATCH /products/{id}/status → Transiciones: draft → pending → active
5. GET /products → Ve sus productos activos en el marketplace
```

### **Estados de Productos**
```
draft → pending → active ↔ inactive
  ↓       ↓         ↓         ↓
deleted ← deleted ← discontinued ← discontinued
```

---

## 📊 Ejemplos de Uso

### **Crear producto desde template**
```bash
# 1. Ver templates disponibles
GET /api/v1/global-catalog/business-types/electronics/products

# 2. Crear producto desde template
POST /api/v1/quickstart/products/from-template
X-Tenant-ID: tenant-123
{
  "template_id": "template-iphone-15",
  "custom_name": "iPhone 15 Pro - Mi Tienda",
  "initial_status": "draft"
}
```

### **Gestionar estados del producto**
```bash
# 1. Ver transiciones disponibles
GET /api/v1/products/{id}/status/transitions
X-Tenant-ID: tenant-123

# 2. Cambiar estado
PATCH /api/v1/products/{id}/status
X-Tenant-ID: tenant-123
{
  "status": "active"
}
```

### **Listar productos del tenant**
```bash
GET /api/v1/products?status=active&category_id=smartphones
X-Tenant-ID: tenant-123
```

---

## 🏗️ Arquitectura

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   TENANT        │    │ GLOBAL CATALOG  │    │   QUICKSTART    │
│   /products     │    │ /global-catalog │    │  /quickstart    │
├─────────────────┤    ├─────────────────┤    ├─────────────────┤
│ • CRUD products │    │ • Templates     │    │ • From template │
│ • Status mgmt   │    │ • By bus. type  │    │ • Mass import   │
│ • Tenant scope  │    │ • Public API    │    │ • Progress      │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

---

## 🚀 Próximas Mejoras

1. **Autenticación**: Implementar JWT/OAuth para rutas privadas
2. **Rate Limiting**: Limitar requests por tenant
3. **Webhooks**: Notificaciones de cambios de estado
4. **Bulk Operations**: Operaciones masivas en productos
5. **Analytics**: Métricas de uso por tenant 