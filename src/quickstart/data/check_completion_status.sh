#!/bin/bash

# Script para verificar el estado de completitud de los tipos de negocio del quickstart
# Uso: ./check_completion_status.sh

echo "📊 Verificando estado de completitud de tipos de negocio..."
echo ""

# Lista de todos los tipos de negocio
BUSINESS_TYPES=(
    "retail"
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

# Función para verificar si un archivo está implementado (no solo tiene TODOs)
is_implemented() {
    local file=$1
    if [ ! -f "$file" ]; then
        echo "❌"
        return 1
    fi
    
    # Verificar si tiene contenido real (no solo comentarios y TODOs)
    local has_real_content=$(grep -v "^#" "$file" | grep -v "TODO" | grep -v "^$" | grep -v "^business_type:" | grep -v "^categories:" | grep -v "^attributes:" | grep -v "^variants:" | grep -v "^products:" | grep -v "^brands:" | wc -l | tr -d ' ')
    
    if [ "$has_real_content" -gt 0 ]; then
        echo "✅"
        return 0
    else
        echo "📝"
        return 1
    fi
}

# Función para contar elementos implementados en un archivo
count_elements() {
    local file=$1
    local element_type=$2
    
    if [ ! -f "$file" ]; then
        echo "0"
        return
    fi
    
    case $element_type in
        "categories")
            grep -c "^  - id:" "$file" 2>/dev/null || echo "0"
            ;;
        "attributes")
            grep -c "^  - id:" "$file" 2>/dev/null || echo "0"
            ;;
        "variants")
            grep -c "^  - id:" "$file" 2>/dev/null || echo "0"
            ;;
        "products")
            grep -c "^  - id:" "$file" 2>/dev/null || echo "0"
            ;;
        "brands")
            grep -c "^  - id:" "$file" 2>/dev/null || echo "0"
            ;;
    esac
}

echo "┌─────────────────────┬────────────┬────────────┬──────────┬──────────┬────────────┬─────────┐"
echo "│ Tipo de Negocio     │ Categorías │ Atributos  │ Variantes│ Productos│ Marcas     │ Estado  │"
echo "├─────────────────────┼────────────┼────────────┼──────────┼──────────┼────────────┼─────────┤"

total_implemented=0
total_types=${#BUSINESS_TYPES[@]}

for business_type in "${BUSINESS_TYPES[@]}"; do
    # Verificar estado de cada archivo
    cat_status=$(is_implemented "categories/${business_type}.yaml")
    attr_status=$(is_implemented "attributes/${business_type}.yaml")
    var_status=$(is_implemented "variants/${business_type}.yaml")
    prod_status=$(is_implemented "products/${business_type}.yaml")
    brand_status=$(is_implemented "brands/${business_type}.yaml")
    
    # Contar elementos
    cat_count=$(count_elements "categories/${business_type}.yaml" "categories")
    attr_count=$(count_elements "attributes/${business_type}.yaml" "attributes")
    var_count=$(count_elements "variants/${business_type}.yaml" "variants")
    prod_count=$(count_elements "products/${business_type}.yaml" "products")
    brand_count=$(count_elements "brands/${business_type}.yaml" "brands")
    
    # Determinar estado general
    if [[ "$cat_status" == "✅" && "$attr_status" == "✅" && "$var_status" == "✅" && "$prod_status" == "✅" && "$brand_status" == "✅" ]]; then
        overall_status="✅ Completo"
        ((total_implemented++))
    elif [[ "$cat_status" == "📝" && "$attr_status" == "📝" && "$var_status" == "📝" && "$prod_status" == "📝" && "$brand_status" == "📝" ]]; then
        overall_status="📝 Pendiente"
    else
        overall_status="🔄 En progreso"
    fi
    
    # Formatear nombre del tipo de negocio
    formatted_name=$(echo "$business_type" | sed 's/-/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2))}1')
    
    printf "│ %-19s │ %2s (%2d)    │ %2s (%2d)    │ %2s (%2d)   │ %2s (%2d)   │ %2s (%2d)    │ %-7s │\n" \
        "$formatted_name" \
        "$cat_status" "$cat_count" \
        "$attr_status" "$attr_count" \
        "$var_status" "$var_count" \
        "$prod_status" "$prod_count" \
        "$brand_status" "$brand_count" \
        "$overall_status"
done

echo "└─────────────────────┴────────────┴────────────┴──────────┴──────────┴────────────┴─────────┘"
echo ""
echo "📈 Resumen:"
echo "  ✅ Tipos completamente implementados: $total_implemented/$total_types"
echo "  🔄 Progreso general: $(( total_implemented * 100 / total_types ))%"
echo ""
echo "📋 Leyenda:"
echo "  ✅ = Implementado con datos reales"
echo "  📝 = Archivo creado pero pendiente de implementar"
echo "  ❌ = Archivo no encontrado"
echo "  🔄 = Implementación parcial"
echo ""

# Mostrar próximos pasos
if [ $total_implemented -lt $total_types ]; then
    echo "🎯 Próximos pasos sugeridos:"
    echo ""
    
    # Encontrar el primer tipo no implementado
    for business_type in "${BUSINESS_TYPES[@]}"; do
        cat_status=$(is_implemented "categories/${business_type}.yaml")
        if [ "$cat_status" != "✅" ]; then
            formatted_name=$(echo "$business_type" | sed 's/-/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2))}1')
            echo "1. 📝 Implementar tipo: $formatted_name ($business_type)"
            echo "   - Archivo de referencia: categories/retail.yaml"
            echo "   - Comenzar con: categories/${business_type}.yaml"
            echo ""
            break
        fi
    done
    
    echo "2. 🔍 Para cada tipo de negocio:"
    echo "   - Investigar categorías típicas del sector"
    echo "   - Definir atributos relevantes para productos"
    echo "   - Configurar variantes comunes"
    echo "   - Agregar productos de ejemplo realistas"
    echo "   - Incluir marcas reconocidas del sector"
    echo ""
    echo "3. ✅ Verificar implementación:"
    echo "   - Ejecutar: ./check_completion_status.sh"
    echo "   - Probar endpoints del quickstart"
    echo "   - Validar estructura YAML"
else
    echo "🎉 ¡Felicidades! Todos los tipos de negocio están implementados."
    echo ""
    echo "🚀 Próximos pasos:"
    echo "1. Probar todos los endpoints del quickstart"
    echo "2. Validar la integración con el onboarding"
    echo "3. Ejecutar tests de integración"
    echo "4. Documentar cualquier tipo de negocio adicional"
fi

echo ""
echo "📚 Comandos útiles:"
echo "  - Ver archivo: cat categories/[tipo].yaml"
echo "  - Editar archivo: nano categories/[tipo].yaml"
echo "  - Verificar YAML: python -c \"import yaml; yaml.safe_load(open('categories/[tipo].yaml'))\""
echo "  - Contar líneas: wc -l categories/[tipo].yaml" 