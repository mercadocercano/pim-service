# 🚀 Datos del Módulo Quickstart - PIM

Este directorio contiene todos los datos YAML necesarios para el módulo quickstart del PIM, que permite la configuración rápida de catálogos predefinidos basados en tipos de negocio.

## 📁 Estructura de Archivos

```
data/
├── business-types.yaml           # ✅ Archivo principal con todos los tipos de negocio
├── categories/                   # 📂 Categorías por tipo de negocio
│   ├── retail.yaml              # ✅ Implementado completamente
│   ├── food-beverage.yaml       # 📝 Pendiente
│   ├── fashion.yaml             # 📝 Pendiente
│   ├── electronics.yaml         # 📝 Pendiente
│   ├── automotive.yaml          # 📝 Pendiente
│   ├── sports-fitness.yaml      # 📝 Pendiente
│   ├── health-pharmacy.yaml     # 📝 Pendiente
│   ├── books-media.yaml         # 📝 Pendiente
│   ├── home-construction.yaml   # 📝 Pendiente
│   ├── beauty-cosmetics.yaml    # 📝 Pendiente
│   ├── toys-games.yaml          # 📝 Pendiente
│   ├── pet-supplies.yaml        # 📝 Pendiente
│   ├── office-supplies.yaml     # 📝 Pendiente
│   └── jewelry-accessories.yaml # 📝 Pendiente
├── attributes/                   # 📂 Atributos por tipo de negocio
│   ├── retail.yaml              # ✅ 20 atributos implementados
│   └── [otros 13 archivos]      # 📝 Pendientes
├── variants/                     # 📂 Configuraciones de variantes
│   ├── retail.yaml              # ✅ 10 variantes implementadas
│   └── [otros 13 archivos]      # 📝 Pendientes
├── products/                     # 📂 Productos de ejemplo
│   ├── retail.yaml              # ✅ 21 productos implementados
│   └── [otros 13 archivos]      # 📝 Pendientes
├── brands/                       # 📂 Marcas reconocidas
│   ├── retail.yaml              # ✅ 37 marcas implementadas
│   └── [otros 13 archivos]      # 📝 Pendientes
├── create_business_type_files.sh # 🛠️ Script para crear archivos
├── check_completion_status.sh    # 🛠️ Script de verificación avanzado
├── simple_status_check.sh        # 🛠️ Script de verificación simple
├── status.sh                     # 🛠️ Script de estado final
└── README.md                     # 📚 Este archivo
```

## 📊 Estado Actual

- **✅ Completamente implementado:** 1 tipo (Retail)
- **📝 Pendientes de implementar:** 13 tipos
- **📁 Total de archivos:** 71 archivos YAML
- **🎯 Progreso:** 7% (1/14 tipos)

## 🏪 Tipo Retail (Implementado)

El tipo **retail** está completamente implementado con:

- **5 categorías principales** con estructura jerárquica de 3 niveles
- **20 atributos específicos** para productos de retail
- **10 configuraciones de variantes** (color-talla, material-peso, etc.)
- **21 productos reales** con precios en ARS
- **37 marcas reconocidas** mundialmente y argentinas

### Categorías Retail
1. **Hogar y Jardín** (Muebles, Decoración, Jardín)
2. **Salud y Belleza** (Cuidado de la Piel, Maquillaje, Higiene)
3. **Electrónicos y Electrodomésticos**
4. **Ropa y Accesorios** (Masculina, Femenina, Calzado)
5. **Alimentos y Bebidas**

## 🛠️ Scripts Disponibles

### 1. `create_business_type_files.sh`
Crea todos los archivos YAML necesarios para los 14 tipos de negocio.

```bash
./create_business_type_files.sh
```

**Funcionalidad:**
- Crea 65 archivos YAML (5 por cada uno de los 13 tipos pendientes)
- Incluye estructura de ejemplo y comentarios TODO
- Crea directorios si no existen

### 2. `status.sh` (Recomendado)
Script simple y claro que muestra el estado actual.

```bash
./status.sh
```

**Muestra:**
- Archivos implementados vs pendientes
- Próximos pasos sugeridos
- Comandos útiles
- Estructura de archivos

### 3. `check_completion_status.sh`
Script avanzado con tabla detallada de estado.

```bash
./check_completion_status.sh
```

### 4. `simple_status_check.sh`
Script intermedio con conteo de elementos.

```bash
./simple_status_check.sh
```

## 🎯 Cómo Implementar un Nuevo Tipo de Negocio

### Paso 1: Elegir un Tipo
Selecciona uno de los 13 tipos pendientes:
- 🍔 Food & Beverage
- 👗 Fashion  
- 📱 Electronics
- 🚗 Automotive
- ⚽ Sports & Fitness
- 💊 Health & Pharmacy
- 📚 Books & Media
- 🏠 Home & Construction
- 💄 Beauty & Cosmetics
- 🧸 Toys & Games
- 🐕 Pet Supplies
- 📎 Office Supplies
- 💍 Jewelry & Accessories

### Paso 2: Editar los 5 Archivos YAML

#### 2.1 Categorías (`categories/[tipo].yaml`)
```yaml
business_type: "fashion"
categories:
  - id: "womens-clothing"
    name: "Ropa Femenina"
    description: "Ropa y accesorios para mujeres"
    parent_id: null
    subcategories:
      - id: "dresses"
        name: "Vestidos"
        description: "Vestidos casuales y formales"
        # ... más subcategorías
```

#### 2.2 Atributos (`attributes/[tipo].yaml`)
```yaml
business_type: "fashion"
attributes:
  - id: "size"
    name: "Talla"
    type: "select"
    required: true
    values: ["XS", "S", "M", "L", "XL", "XXL"]
  - id: "color"
    name: "Color"
    type: "select"
    required: false
    values: ["Negro", "Blanco", "Azul", "Rojo"]
```

#### 2.3 Variantes (`variants/[tipo].yaml`)
```yaml
business_type: "fashion"
variants:
  - id: "size-color"
    name: "Talla y Color"
    attributes: ["size", "color"]
    combinations:
      - size: "S"
        color: "Negro"
        sku_suffix: "S-BLK"
```

#### 2.4 Productos (`products/[tipo].yaml`)
```yaml
business_type: "fashion"
products:
  - id: "summer-dress-001"
    name: "Vestido de Verano Floral"
    category: "dresses"
    base_price: 12999
    currency: "ARS"
    sku: "DRESS-001"
```

#### 2.5 Marcas (`brands/[tipo].yaml`)
```yaml
business_type: "fashion"
brands:
  - id: "zara"
    name: "Zara"
    description: "Moda contemporánea"
    country_origin: "España"
```

### Paso 3: Verificar Implementación
```bash
./status.sh
```

## 📚 Archivos de Referencia

Usa estos archivos como referencia para implementar nuevos tipos:

- **Estructura completa:** `categories/retail.yaml`
- **Atributos diversos:** `attributes/retail.yaml`
- **Variantes complejas:** `variants/retail.yaml`
- **Productos reales:** `products/retail.yaml`
- **Marcas reconocidas:** `brands/retail.yaml`

## 🔍 Comandos Útiles

```bash
# Ver estructura de directorios
ls -la categories/

# Ver archivo de ejemplo
cat categories/retail.yaml

# Editar un archivo
nano categories/fashion.yaml

# Verificar sintaxis YAML
python -c "import yaml; yaml.safe_load(open('categories/fashion.yaml'))"

# Contar elementos implementados
grep -c "^  - id:" categories/retail.yaml

# Ver estado general
./status.sh
```

## 🎨 Convenciones de Datos

### Precios
- Usar pesos argentinos (ARS)
- Precios realistas del mercado actual
- Formato: números enteros (ej: 15999 para $159.99)

### IDs
- Usar kebab-case (ej: "womens-clothing")
- Descriptivos y únicos
- Sin espacios ni caracteres especiales

### Nombres
- En español
- Descriptivos y claros
- Usar mayúsculas apropiadas

### Marcas
- Incluir marcas reconocidas mundialmente
- Agregar marcas argentinas relevantes
- Proporcionar descripción y país de origen

## 🚀 Integración con el Sistema

Una vez implementados, estos datos se integran automáticamente con:

- **API Endpoints del Quickstart**
- **Sistema de Onboarding Multi-tenant**
- **Configuración Automática de Catálogos**
- **Carga de Datos Predefinidos**

## 📈 Próximos Pasos

1. **Implementar Fashion** (sugerido como siguiente)
2. **Implementar Electronics** 
3. **Implementar Food & Beverage**
4. **Continuar con los tipos restantes**
5. **Probar integración completa**
6. **Documentar casos de uso específicos**

## 🤝 Contribución

Para contribuir con nuevos tipos de negocio:

1. Ejecutar `./create_business_type_files.sh` si no está hecho
2. Elegir un tipo pendiente
3. Implementar los 5 archivos YAML
4. Verificar con `./status.sh`
5. Probar con los endpoints del quickstart
6. Documentar cualquier consideración especial

---

**¡Listo para implementar los tipos de negocio restantes! 🚀** 