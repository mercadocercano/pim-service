#!/bin/bash

# Script para crear todos los archivos YAML necesarios para los tipos de negocio del quickstart
# Uso: ./create_business_type_files.sh

echo "🚀 Creando archivos YAML para todos los tipos de negocio del quickstart..."

# Lista de todos los tipos de negocio (excluyendo retail que ya existe)
BUSINESS_TYPES=(
    "food-beverage"
    "fashion"
    "electronics"
    "automotive"
    "sports-fitness"
    "health-pharmacy"
    "books-media"
    "home-construction"
    "beauty-cosmetics"
    "toys-games"
    "pet-supplies"
    "office-supplies"
    "jewelry-accessories"
)

# Crear directorios si no existen
echo "📁 Creando directorios..."
mkdir -p categories attributes variants products brands

# Función para crear archivo de categorías
create_categories_file() {
    local business_type=$1
    local file="categories/${business_type}.yaml"
    
    cat > "$file" << EOF
business_type: "${business_type}"
categories:
  # TODO: Implementar categorías para ${business_type}
  # Estructura de ejemplo:
  # - id: "main-category"
  #   name: "Categoría Principal"
  #   description: "Descripción de la categoría"
  #   parent_id: null
  #   subcategories:
  #     - id: "subcategory-1"
  #       name: "Subcategoría 1"
  #       description: "Descripción de subcategoría"
  #       subcategories:
  #         - id: "sub-subcategory-1"
  #           name: "Sub-subcategoría 1"
  #           description: "Descripción de sub-subcategoría"
EOF
    echo "✅ Creado: $file"
}

# Función para crear archivo de atributos
create_attributes_file() {
    local business_type=$1
    local file="attributes/${business_type}.yaml"
    
    cat > "$file" << EOF
business_type: "${business_type}"
attributes:
  # TODO: Implementar atributos para ${business_type}
  # Estructura de ejemplo:
  # - id: "color"
  #   name: "Color"
  #   type: "select"
  #   required: false
  #   description: "Color del producto"
  #   values:
  #     - "Rojo"
  #     - "Azul"
  #     - "Verde"
  #
  # - id: "weight"
  #   name: "Peso"
  #   type: "number"
  #   required: false
  #   description: "Peso del producto en gramos"
  #   unit: "g"
  #   min_value: 0
  #   max_value: 10000
  #
  # - id: "description"
  #   name: "Descripción Detallada"
  #   type: "text"
  #   required: false
  #   description: "Descripción detallada del producto"
  #   placeholder: "Ingrese descripción..."
  #
  # - id: "premium"
  #   name: "Producto Premium"
  #   type: "boolean"
  #   required: false
  #   description: "Indica si es un producto premium"
EOF
    echo "✅ Creado: $file"
}

# Función para crear archivo de variantes
create_variants_file() {
    local business_type=$1
    local file="variants/${business_type}.yaml"
    
    cat > "$file" << EOF
business_type: "${business_type}"
variants:
  # TODO: Implementar variantes para ${business_type}
  # Estructura de ejemplo:
  # - id: "color-size"
  #   name: "Color y Tamaño"
  #   description: "Variante por color y tamaño del producto"
  #   attributes:
  #     - "color"
  #     - "size"
  #   combinations:
  #     - color: "Rojo"
  #       size: "S"
  #       sku_suffix: "RED-S"
  #     - color: "Rojo"
  #       size: "M"
  #       sku_suffix: "RED-M"
  #     - color: "Azul"
  #       size: "S"
  #       sku_suffix: "BLU-S"
  #     - color: "Azul"
  #       size: "M"
  #       sku_suffix: "BLU-M"
EOF
    echo "✅ Creado: $file"
}

# Función para crear archivo de productos
create_products_file() {
    local business_type=$1
    local file="products/${business_type}.yaml"
    
    cat > "$file" << EOF
business_type: "${business_type}"
products:
  # TODO: Implementar productos para ${business_type}
  # Estructura de ejemplo:
  # - id: "product-example-1"
  #   name: "Producto de Ejemplo 1"
  #   description: "Descripción del producto de ejemplo"
  #   category: "main-category"
  #   base_price: 9999
  #   currency: "ARS"
  #   sku: "PROD-001"
  #   attributes:
  #     color: "Rojo"
  #     weight: 500
  #     premium: true
  #   variants:
  #     - "color-size"
  #   stock: 100
  #
  # - id: "product-example-2"
  #   name: "Producto de Ejemplo 2"
  #   description: "Otro producto de ejemplo"
  #   category: "subcategory-1"
  #   base_price: 15999
  #   currency: "ARS"
  #   sku: "PROD-002"
  #   attributes:
  #     color: "Azul"
  #     weight: 750
  #     premium: false
  #   variants:
  #     - "color-size"
  #   stock: 50
EOF
    echo "✅ Creado: $file"
}

# Función para crear archivo de marcas
create_brands_file() {
    local business_type=$1
    local file="brands/${business_type}.yaml"
    
    cat > "$file" << EOF
business_type: "${business_type}"
brands:
  # TODO: Implementar marcas para ${business_type}
  # Estructura de ejemplo:
  # - id: "brand-example-1"
  #   name: "Marca Ejemplo 1"
  #   description: "Descripción de la marca ejemplo"
  #   country_origin: "Argentina"
  #   category: "Categoría de la Marca"
  #   website: "https://www.ejemplo.com"
  #   logo_url: "/brands/ejemplo-logo.png"
  #
  # - id: "brand-example-2"
  #   name: "Marca Ejemplo 2"
  #   description: "Otra marca de ejemplo"
  #   country_origin: "Estados Unidos"
  #   category: "Otra Categoría"
  #   website: "https://www.ejemplo2.com"
  #   logo_url: "/brands/ejemplo2-logo.png"
EOF
    echo "✅ Creado: $file"
}

# Crear archivos para cada tipo de negocio
echo ""
echo "📝 Creando archivos para cada tipo de negocio..."

for business_type in "${BUSINESS_TYPES[@]}"; do
    echo ""
    echo "🏢 Procesando: $business_type"
    
    create_categories_file "$business_type"
    create_attributes_file "$business_type"
    create_variants_file "$business_type"
    create_products_file "$business_type"
    create_brands_file "$business_type"
done

echo ""
echo "📊 Resumen de archivos creados:"
echo "==============================================="

total_files=0
for business_type in "${BUSINESS_TYPES[@]}"; do
    count=0
    for dir in categories attributes variants products brands; do
        if [ -f "${dir}/${business_type}.yaml" ]; then
            ((count++))
            ((total_files++))
        fi
    done
    echo "📁 $business_type: $count archivos"
done

echo "==============================================="
echo "✨ Total de archivos creados: $total_files"
echo ""
echo "🎯 Próximos pasos:"
echo "1. Completar los archivos YAML con datos reales para cada tipo de negocio"
echo "2. Seguir la estructura del archivo 'retail.yaml' como referencia"
echo "3. Actualizar business-types.yaml si es necesario"
echo "4. Probar cada tipo de negocio conforme se complete"
echo ""
echo "📚 Archivos de referencia:"
echo "- categories/retail.yaml"
echo "- attributes/retail.yaml" 
echo "- variants/retail.yaml"
echo "- products/retail.yaml"
echo "- brands/retail.yaml"
echo ""
echo "🚀 ¡Listo para comenzar a implementar los tipos de negocio!" 