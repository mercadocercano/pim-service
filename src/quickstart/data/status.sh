#!/bin/bash

echo "🚀 Estado del Módulo Quickstart - Tipos de Negocio"
echo "=================================================="
echo ""

# Contar archivos totales
total_files=$(find . -name "*.yaml" | wc -l | tr -d ' ')
echo "📁 Total de archivos YAML: $total_files"
echo ""

# Cargar tipos de negocio desde el archivo principal
business_types=($(grep "^  - id:" business-types.yaml | cut -d '"' -f 2))

# Contador para tipos implementados
implemented_count=0
pending_count=0

echo "⏳ VERIFICANDO IMPLEMENTACIÓN DE TIPOS DE NEGOCIO..."
echo ""

echo "✅ IMPLEMENTADOS:"
for type in "${business_types[@]}"; do
    # Reemplazar guiones por guiones bajos para nombres de variables
    type_var=$(echo $type | tr '-' '_')
    
    # Contar elementos en cada archivo si existe
    categories_count=$(grep -c "^  - id:" "categories/$type.yaml" 2>/dev/null || echo "0")
    attributes_count=$(grep -c "^  - id:" "attributes/$type.yaml" 2>/dev/null || echo "0")
    variants_count=$(grep -c "^  - id:" "variants/$type.yaml" 2>/dev/null || echo "0")
    products_count=$(grep -c "^  - id:" "products/$type.yaml" 2>/dev/null || echo "0")
    brands_count=$(grep -c "^  - id:" "brands/$type.yaml" 2>/dev/null || echo "0")
    
    # Verificar si tiene contenido (al menos un ítem en cada archivo)
    if [[ -f "categories/$type.yaml" && -f "attributes/$type.yaml" && -f "variants/$type.yaml" && 
          -f "products/$type.yaml" && -f "brands/$type.yaml" ]]; then
        
        # Mostrar como implementado si tiene los archivos básicos
        case "$type" in
            "retail")
                emoji="🏪"
                ;;
            "food-beverage")
                emoji="🍔"
                ;;
            "fashion")
                emoji="👗"
                ;;
            "electronics")
                emoji="📱"
                ;;
            "automotive")
                emoji="🚗"
                ;;
            "sports-fitness")
                emoji="⚽"
                ;;
            "health-pharmacy")
                emoji="💊"
                ;;
            "books-media")
                emoji="📚"
                ;;
            "home-construction")
                emoji="🏠"
                ;;
            "beauty-cosmetics")
                emoji="💄"
                ;;
            "toys-games")
                emoji="🧸"
                ;;
            "pet-supplies")
                emoji="🐕"
                ;;
            "office-supplies")
                emoji="📎"
                ;;
            "jewelry-accessories")
                emoji="💍"
                ;;
            *)
                emoji="📦"
                ;;
        esac
        
        echo "   $emoji $type: $categories_count categorías, $attributes_count atributos, $variants_count variantes, $products_count productos, $brands_count marcas"
        implemented_count=$((implemented_count + 1))
    else
        pending_count=$((pending_count + 1))
    fi
done

echo ""
if [ $pending_count -gt 0 ]; then
    echo "📝 PENDIENTES DE IMPLEMENTAR ($pending_count tipos):"
    for type in "${business_types[@]}"; do
        if [[ ! -f "categories/$type.yaml" || ! -f "attributes/$type.yaml" || ! -f "variants/$type.yaml" || 
              ! -f "products/$type.yaml" || ! -f "brands/$type.yaml" ]]; then
            
            case "$type" in
                "retail")
                    emoji="🏪"
                    ;;
                "food-beverage")
                    emoji="🍔"
                    ;;
                "fashion")
                    emoji="👗"
                    ;;
                "electronics")
                    emoji="📱"
                    ;;
                "automotive")
                    emoji="🚗"
                    ;;
                "sports-fitness")
                    emoji="⚽"
                    ;;
                "health-pharmacy")
                    emoji="💊"
                    ;;
                "books-media")
                    emoji="📚"
                    ;;
                "home-construction")
                    emoji="🏠"
                    ;;
                "beauty-cosmetics")
                    emoji="💄"
                    ;;
                "toys-games")
                    emoji="🧸"
                    ;;
                "pet-supplies")
                    emoji="🐕"
                    ;;
                "office-supplies")
                    emoji="📎"
                    ;;
                "jewelry-accessories")
                    emoji="💍"
                    ;;
                *)
                    emoji="📦"
                    ;;
            esac
            
            echo "   $emoji $type"
        fi
    done
else
    echo "🎉 ¡Todos los tipos de negocio están implementados!"
fi

echo ""

echo "🎯 PRÓXIMOS PASOS:"
if [ $pending_count -gt 0 ]; then
    echo "1. Elegir un tipo de negocio pendiente para implementar"
    echo "2. Editar los 5 archivos YAML correspondientes:"
    echo "   - categories/[tipo].yaml"
    echo "   - attributes/[tipo].yaml"
    echo "   - variants/[tipo].yaml"
    echo "   - products/[tipo].yaml"
    echo "   - brands/[tipo].yaml"
    echo "3. Usar los archivos existentes como referencia (retail, sports-fitness, toys-games)"
else
    echo "1. Revisar la calidad de los datos implementados"
    echo "2. Verificar la consistencia entre archivos"
    echo "3. Probar la integración con la API"
fi
echo "4. Ejecutar ./status.sh para verificar progreso"
echo ""

echo "📚 COMANDOS ÚTILES:"
echo "   Ver estructura: ls -la categories/"
echo "   Editar archivo: nano categories/[tipo].yaml"
echo "   Ver ejemplo: cat categories/retail.yaml"
echo "   Verificar estado: ./status.sh"
echo "   Estado detallado: ./check_completion_status.sh"
echo ""

echo "🏗️ ESTRUCTURA DE ARCHIVOS:"
total_categories=$(ls categories/*.yaml 2>/dev/null | wc -l | tr -d ' ')
total_attributes=$(ls attributes/*.yaml 2>/dev/null | wc -l | tr -d ' ')
total_variants=$(ls variants/*.yaml 2>/dev/null | wc -l | tr -d ' ')
total_products=$(ls products/*.yaml 2>/dev/null | wc -l | tr -d ' ')
total_brands=$(ls brands/*.yaml 2>/dev/null | wc -l | tr -d ' ')

echo "   📂 categories/    - $total_categories archivos ($implemented_count completos, $pending_count pendientes)"
echo "   📂 attributes/    - $total_attributes archivos ($implemented_count completos, $pending_count pendientes)"
echo "   📂 variants/      - $total_variants archivos ($implemented_count completos, $pending_count pendientes)"
echo "   📂 products/      - $total_products archivos ($implemented_count completos, $pending_count pendientes)"
echo "   📂 brands/        - $total_brands archivos ($implemented_count completos, $pending_count pendientes)"
echo "   📄 business-types.yaml - Archivo principal (completo)"
echo ""

if [ $pending_count -gt 0 ]; then
    echo "✨ Progreso: $implemented_count/${#business_types[@]} tipos de negocio implementados ($(( implemented_count * 100 / ${#business_types[@]} ))%)"
else
    echo "🎉 ¡Felicidades! Todos los tipos de negocio están implementados."
fi 