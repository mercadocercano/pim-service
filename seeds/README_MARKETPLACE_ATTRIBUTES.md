# Sistema de Atributos Marketplace Argentina

## 🎯 Propósito

Sistema completo de atributos y categorías para marketplace argentino, con datos localizados y filtros contextuales específicos para cada tipo de producto.

## 📊 Estructura del Sistema

### 🗃️ **Datos Cargados**
- **50 categorías** organizadas en 3 niveles jerárquicos
- **15+ atributos** específicos para el mercado argentino
- **200+ valores** de atributos localizados
- **200+ relaciones** categoría-atributo configuradas

### 🏗️ **Arquitectura de Tablas**

```
marketplace_categories (50 registros)
├── Nivel 0: 10 categorías principales
├── Nivel 1: 20+ subcategorías
└── Nivel 2: 20+ categorías específicas

marketplace_attributes (15+ registros)
├── Atributos generales (marca, condición, color)
├── Atributos de moda (talle, material textil, género)
├── Atributos de tecnología (almacenamiento, RAM, marca tech)
├── Atributos de hogar (material, ambiente, tamaño)
└── Atributos específicos (peso, volumen, marca alimentos)

marketplace_attribute_values (200+ registros)
├── Valores por atributo con slugs
└── Sort order para ordenamiento

marketplace_category_attributes (200+ registros)
├── Relaciones categoría ↔ atributo
├── is_required flag
└── sort_order para UI
```

## 🚀 Ejecución de Seeders

### **Método Automático (Recomendado)**
```bash
cd services/saas-mt-pim-service
./scripts/run-marketplace-seeders.sh
```

### **Método Manual**
```bash
# 1. Categorías (50 categorías en 3 niveles)
psql -f seeds/013_marketplace_categories_seeder.sql

# 2. Atributos (15+ atributos con 200+ valores)
psql -f seeds/014_marketplace_attributes_argentina.sql  

# 3. Relaciones (200+ mapeos categoría-atributo)
psql -f seeds/015_marketplace_category_attributes_relations.sql
```

## 📋 Categorías por Industria

### 👕 **Moda y Accesorios**
```
Moda y Accesorios/
├── Ropa Mujer/
│   ├── Remeras y Tops
│   ├── Vestidos  
│   ├── Pantalones
│   ├── Abrigos
│   └── Faldas
├── Ropa Hombre/
│   ├── Remeras y Polos
│   ├── Camisas
│   ├── Pantalones Hombre
│   └── Abrigos Hombre
├── Calzado/
│   ├── Zapatillas
│   ├── Zapatos Formales
│   ├── Sandalias
│   └── Botas
├── Accesorios/
└── Lencería/
```

**Atributos aplicables:**
- Talle: XS, S, M, L, XL, XXL, XXXL, 1-18, Único
- Talle Calzado: 32-46 (sistema argentino)
- Material: Algodón, Poliéster, Lycra, Lana, Cuero, Denim, etc.
- Género: Mujer, Hombre, Unisex, Niña, Niño, Bebé
- Color: 14 colores principales

### 💻 **Tecnología**
```
Tecnología/
├── Celulares y Tablets
├── Computadoras
├── Gaming
└── Audio y Video
```

**Atributos aplicables:**
- Almacenamiento: 16GB-4TB
- RAM: 2GB-64GB  
- Marca Tech: Apple, Samsung, Huawei, Xiaomi, etc.
- Condición: Nuevo, Usado, Reacondicionado

### 🏠 **Hogar y Jardín**
```
Hogar y Jardín/
├── Muebles/
│   ├── Sillones y Sofás
│   ├── Mesas
│   └── Dormitorio
├── Decoración
├── Electrodomésticos
├── Jardín
└── Cocina
```

**Atributos aplicables:**
- Material: Madera, Metal, Plástico, Vidrio, MDF, etc.
- Ambiente: Living, Comedor, Dormitorio, Cocina, etc.
- Tamaño: Pequeño, Mediano, Grande, Extra Grande

### 🍔 **Comida y Bebidas**
**Atributos aplicables:**
- Peso: 50g-5kg
- Volumen: 250ml-3L
- Marca Alimentos: Arcor, Mastellone, Molinos, Nestlé, etc.

## 🔗 Endpoints Disponibles

### **Backend (PIM Service)**
```
GET /marketplace/categories
- Lista todas las categorías con jerarquía

GET /marketplace/categories/{id}/attributes  
- Atributos específicos de una categoría

GET /marketplace/attributes
- Lista todos los atributos disponibles

GET /marketplace/attributes/{id}/values
- Valores de un atributo específico
```

## 🎨 Filtros Contextuales por Categoría

El sistema aplica **filtros inteligentes** según la categoría:

### **Ejemplo: Remeras Mujer**
```json
{
  "category": "Remeras y Tops",
  "filters": [
    {"name": "Talle", "values": ["XS", "S", "M", "L", "XL"]},
    {"name": "Color", "values": ["Negro", "Blanco", "Azul"]},
    {"name": "Material", "values": ["Algodón", "Poliéster"]},
    {"name": "Género", "values": ["Mujer"]}
  ]
}
```

### **Ejemplo: iPhone**
```json
{
  "category": "Celulares y Tablets", 
  "filters": [
    {"name": "Almacenamiento", "values": ["64GB", "128GB", "256GB"]},
    {"name": "RAM", "values": ["4GB", "6GB", "8GB"]},
    {"name": "Marca", "values": ["Apple", "Samsung"]},
    {"name": "Condición", "values": ["Nuevo", "Usado"]}
  ]
}
```

## 📱 Integración Frontend

### **Backoffice (Next.js)**
```typescript
// Páginas CRUD disponibles
/dashboard/pim/marketplace-attributes     // Gestión de atributos
/dashboard/pim/marketplace-categories     // Gestión de categorías

// API calls
const attributes = await fetchMarketplaceAttributes();
const categoryFilters = await fetchCategoryAttributes(categoryId);
```

### **Marketplace Frontend**
```typescript
// Filtros dinámicos por categoría
const ProductFilters = ({ categoryId }) => {
  const filters = useCategoryFilters(categoryId);
  
  return (
    <FilterPanel>
      {filters.map(filter => (
        <FilterGroup key={filter.id} filter={filter} />
      ))}
    </FilterPanel>
  );
};
```

## 🔍 Verificación de Datos

### **Comandos de Verificación**
```bash
# Contar categorías cargadas
psql -c "SELECT count(*) FROM marketplace_categories WHERE is_active = true;"

# Ver distribución por niveles
psql -c "SELECT level, count(*) FROM marketplace_categories GROUP BY level ORDER BY level;"

# Contar atributos y valores
psql -c "SELECT count(*) FROM marketplace_attributes;"
psql -c "SELECT count(*) FROM marketplace_attribute_values;"

# Ver relaciones configuradas
psql -c "SELECT count(*) FROM marketplace_category_attributes;"

# Ejemplo: Atributos de "Remeras y Tops"
psql -c "
SELECT ma.name, ma.type, COUNT(mav.id) as value_count
FROM marketplace_categories mc
JOIN marketplace_category_attributes mca ON mc.id = mca.category_id  
JOIN marketplace_attributes ma ON mca.attribute_id = ma.id
LEFT JOIN marketplace_attribute_values mav ON ma.id = mav.attribute_id
WHERE mc.slug = 'remeras-tops'
GROUP BY ma.id, ma.name, ma.type
ORDER BY mca.sort_order;
"
```

## 🛠️ Mantenimiento

### **Agregar Nueva Categoría**
```sql
INSERT INTO marketplace_categories (id, name, slug, description, parent_id, level, sort_order) 
VALUES ('new-cat-id', 'Nueva Categoría', 'nueva-categoria', 'Descripción', 'parent-id', 2, 10);

-- Asignar atributos
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order)
VALUES ('new-cat-id', 'attr-color', false, 1);
```

### **Agregar Nuevo Atributo**
```sql
-- Crear atributo
INSERT INTO marketplace_attributes (id, name, slug, type, is_filterable, is_searchable) 
VALUES ('attr-nuevo', 'Nuevo Atributo', 'nuevo-atributo', 'select', true, false);

-- Agregar valores
INSERT INTO marketplace_attribute_values (attribute_id, value, slug, sort_order) VALUES
('attr-nuevo', 'Valor 1', 'valor-1', 1),
('attr-nuevo', 'Valor 2', 'valor-2', 2);

-- Asignar a categorías
INSERT INTO marketplace_category_attributes (category_id, attribute_id, is_required, sort_order)
SELECT id, 'attr-nuevo', false, 10 FROM marketplace_categories WHERE slug LIKE 'categoria-target%';
```

## 🎉 Estado del Sistema

✅ **COMPLETADO:**
- Estructura de tablas
- 50 categorías marketplace
- 15+ atributos argentinizados  
- 200+ valores localizados
- 200+ relaciones categoría-atributo
- Endpoints backend funcionales
- Script de carga automático

🚀 **PRÓXIMOS PASOS:**
- Páginas CRUD en backoffice
- Componentes de filtros frontend
- Integración con productos
- Analytics de uso de filtros

---

## 💡 Tips de Uso

1. **Siempre usar el script automático** para cargar datos
2. **Verificar relaciones** antes de crear productos
3. **Testear filtros** en cada categoría específica
4. **Mantener consistencia** en naming de slugs
5. **Documentar cambios** en atributos custom

Este sistema está **listo para producción** con datos argentinos reales y filtros contextuales inteligentes. 🇦🇷 