# Seeds para Marketplace Argentino

Este directorio contiene los seeds para poblar la base de datos con tipos de comercio y categorías basados en investigación de mercado argentino 2024-2025.

## 📋 Contenido

### `001_business_types_argentina_seed.sql`

Seed que incluye **28 tipos de comercio físicos argentinos** organizados por categorías:

### `002_business_types_additional_seed.sql`

Seed adicional con **7 tipos de comercio digitales y servicios** identificados en la investigación:
- Agencia de Viajes
- Delivery de Comidas  
- Servicios Digitales
- Centro Educativo
- Entretenimiento
- Lavadero de Autos
- Servicios Profesionales

### `003_marketplace_categories_argentina_seed.sql`

Seed de **categorías globales del marketplace** basado en estudios CACE, Kantar TNS, PCMI y Americas Market Intelligence:
- **14 categorías principales** ordenadas por facturación e importancia
- **Más de 70 subcategorías** organizadas jerárquicamente
- Estructura optimizada para navegación de compradores

### `004_business_type_quickstart_templates.sql`

Seed de **templates de quickstart** que vincula inteligentemente tipos de negocio con categorías relevantes:
- **35 templates personalizados** (uno por cada tipo de negocio)
- **Mapeos inteligentes** basados en lógica de negocio argentina
- **Categorías pre-sugeridas** para facilitar onboarding de sellers
- **Sistema de quickstart** listo para usar en el proceso de registro

#### 🍽️ Comercios Alimentarios
- Almacén de Barrio
- Supermercado  
- Carnicería
- Panadería
- Verdulería
- Fiambrería
- Heladería

#### 💊 Farmacias y Salud
- Farmacia
- Perfumería

#### 👕 Indumentaria y Calzado
- Tienda de Ropa
- Zapatería
- Artículos Deportivos

#### 🏠 Hogar y Construcción
- Ferretería
- Mueblería
- Bazar

#### 📱 Tecnología y Comunicaciones
- Casa de Electrodomésticos
- Casa de Celulares
- Computación

#### 🚗 Automotriz
- Casa de Repuestos
- Lubricentro

#### 🤓 Servicios Especializados
- Óptica
- Relojería
- Librería
- Juguetería

#### 🐕 Mascotas
- Veterinaria

#### 🌸 Servicios Diversos
- Kiosco
- Florería
- Polirubro

## 🚀 Ejecución

### Opción 1: Seed Completo (Recomendado)

```bash
# Ejecuta TODOS los seeds: business_types + categorías
./scripts/seed_complete_marketplace_docker.sh
```

### Opción 2: Solo Business Types (Docker)

```bash
# Solo tipos de comercio argentinos
./scripts/seed_business_types_docker.sh
```

### Opción 3: Solo Business Types (Conexión directa)

```bash
# Solo tipos de comercio argentinos
./scripts/seed_business_types_argentina.sh
```

### Opción 4: Manual por partes

```bash
# 1. Tipos de comercio base
psql -h localhost -p 5432 -d pim_db -U postgres -f seeds/001_business_types_argentina_seed.sql

# 2. Tipos de comercio adicionales
psql -h localhost -p 5432 -d pim_db -U postgres -f seeds/002_business_types_additional_seed.sql

# 3. Categorías del marketplace
psql -h localhost -p 5432 -d pim_db -U postgres -f seeds/003_marketplace_categories_argentina_seed.sql

# 4. Templates de quickstart (¡IMPORTANTE para onboarding!)
psql -h localhost -p 5432 -d pim_db -U postgres -f seeds/004_business_type_quickstart_templates.sql
```

## ⚠️ Importante

- **Los seeds BORRAN todos los datos existentes** en las tablas correspondientes
- `001_business_types_argentina_seed.sql` limpia: `business_types`, `business_type_templates`, `tenant_business_type_setup`
- `003_marketplace_categories_argentina_seed.sql` limpia: `marketplace_categories`
- Asegúrate de hacer backup si tienes datos importantes
- Todos los datos están diseñados específicamente para el mercado argentino

## 🎯 Características

### Business Types
- **35 tipos de comercio** (28 físicos + 7 digitales/servicios)
- **Códigos únicos** para identificación
- **Iconos apropiados** para cada comercio
- **Colores distintivos** para UI
- **Descripciones realistas** del mercado argentino

### Marketplace Categories  
- **14 categorías principales** basadas en estudios de facturación
- **Más de 70 subcategorías** organizadas jerárquicamente
- **Slugs optimizados** para URLs amigables
- **Estructura de 2 niveles** (principal + subcategoría)
- **Ordenamiento por relevancia** de mercado

### Business Type Templates (Quickstart)
- **35 templates personalizados** vinculando tipos con categorías
- **Mapeos inteligentes** por lógica de negocio
- **Categorías pre-sugeridas** para cada tipo de comercio
- **Sistema de onboarding** facilitado para sellers
- **Configuración automática** de categorías iniciales

## 📊 Verificación

Después de ejecutar los seeds, puedes verificar los datos con:

```sql
-- Verificar tipos de comercio
SELECT code, name, description FROM business_types ORDER BY sort_order;

-- Verificar categorías principales
SELECT name, slug, description FROM marketplace_categories 
WHERE parent_id IS NULL ORDER BY sort_order;

-- Verificar subcategorías de una categoría
SELECT c1.name as categoria, c2.name as subcategoria
FROM marketplace_categories c1 
JOIN marketplace_categories c2 ON c1.id = c2.parent_id 
WHERE c1.slug = 'alimentos-bebidas' 
ORDER BY c2.sort_order;

-- Resumen general
SELECT 
  (SELECT COUNT(*) FROM business_types) as tipos_comercio,
  (SELECT COUNT(*) FROM marketplace_categories WHERE parent_id IS NULL) as categorias_principales,
  (SELECT COUNT(*) FROM marketplace_categories WHERE parent_id IS NOT NULL) as subcategorias,
  (SELECT COUNT(*) FROM business_type_templates) as templates_quickstart;

-- Verificar templates de quickstart por tipo de negocio
SELECT 
  bt.name as tipo_negocio,
  btt.name as template_name,
  jsonb_array_length(btt.categories) as num_categorias_sugeridas
FROM business_types bt
JOIN business_type_templates btt ON bt.id = btt.business_type_id
ORDER BY bt.sort_order;

-- Ver categorías sugeridas para un tipo específico (ejemplo: almacén)
SELECT 
  bt.name as tipo_negocio,
  mc.name as categoria_sugerida,
  mc.slug as categoria_slug
FROM business_types bt
JOIN business_type_templates btt ON bt.id = btt.business_type_id,
jsonb_array_elements(btt.categories) as cat_obj
JOIN marketplace_categories mc ON mc.id = (cat_obj->>'id')::uuid
WHERE bt.code = 'almacen'
ORDER BY mc.sort_order;
```

## 📈 Fuentes de Investigación

Los seeds están basados en estudios de mercado argentino:
- **CACE (Cámara Argentina de Comercio Electrónico)** - Estudios 2024-2025
- **Kantar TNS** - Mid Term 2025 de Comercio Electrónico  
- **PCMI** - Análisis primer semestre 2024
- **Americas Market Intelligence** - Categorías más compradas 2022
- **ClearSale** - Principales categorías e-commerce 2019
- **Mailchimp** - Tendencias globales y nichos en crecimiento 