#!/bin/bash

# Script simplificado para verificar el estado de completitud de los tipos de negocio
echo "📊 Estado de archivos YAML del quickstart"
echo "=========================================="
echo ""

# Lista de tipos de negocio
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

echo "Tipo de Negocio      | Cat | Attr | Var | Prod | Brand | Estado"
echo "-------------------- | --- | ---- | --- | ---- | ----- | --------"

total_implemented=0
total_pending=0

for business_type in "${BUSINESS_TYPES[@]}"; do
    # Verificar archivos
    cat_file="categories/${business_type}.yaml"
    attr_file="attributes/${business_type}.yaml"
    var_file="variants/${business_type}.yaml"
    prod_file="products/${business_type}.yaml"
    brand_file="brands/${business_type}.yaml"
    
    # Contar elementos reales (líneas que empiezan con "  - id:")
    cat_count=$(grep -c "^  - id:" "$cat_file" 2>/dev/null || echo "0")
    attr_count=$(grep -c "^  - id:" "$attr_file" 2>/dev/null || echo "0")
    var_count=$(grep -c "^  - id:" "$var_file" 2>/dev/null || echo "0")
    prod_count=$(grep -c "^  - id:" "$prod_file" 2>/dev/null || echo "0")
    brand_count=$(grep -c "^  - id:" "$brand_file" 2>/dev/null || echo "0")
    
    # Determinar estado
    if [ "$cat_count" -gt 0 ] && [ "$attr_count" -gt 0 ] && [ "$var_count" -gt 0 ] && [ "$prod_count" -gt 0 ] && [ "$brand_count" -gt 0 ]; then
        status="✅ Completo"
        ((total_implemented++))
    elif [ "$cat_count" -eq 0 ] && [ "$attr_count" -eq 0 ] && [ "$var_count" -eq 0 ] && [ "$prod_count" -eq 0 ] && [ "$brand_count" -eq 0 ]; then
        status="📝 Pendiente"
        ((total_pending++))
    else
        status="🔄 Parcial"
    fi
    
    # Formatear nombre
    formatted_name=$(echo "$business_type" | sed 's/-/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2))}1')
    
    printf "%-20s | %3d | %4d | %3d | %4d | %5d | %s\n" \
        "$formatted_name" \
        "$cat_count" "$attr_count" "$var_count" "$prod_count" "$brand_count" \
        "$status"
done

echo ""
echo "📈 Resumen:"
echo "  ✅ Implementados: $total_implemented"
echo "  📝 Pendientes: $total_pending"
echo "  🔄 Parciales: $(( ${#BUSINESS_TYPES[@]} - total_implemented - total_pending ))"
echo ""

if [ $total_pending -gt 0 ]; then
    echo "🎯 Próximo paso sugerido:"
    for business_type in "${BUSINESS_TYPES[@]}"; do
        cat_count=$(grep -c "^  - id:" "categories/${business_type}.yaml" 2>/dev/null || echo "0")
        if [ "$cat_count" -eq 0 ] && [ "$business_type" != "retail" ]; then
            formatted_name=$(echo "$business_type" | sed 's/-/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) tolower(substr($i,2))}1')
            echo "  📝 Implementar: $formatted_name"
            echo "     Comenzar editando: categories/${business_type}.yaml"
            break
        fi
    done
fi

echo ""
echo "📚 Archivos creados: $(find . -name "*.yaml" | wc -l | tr -d ' ') archivos YAML"
echo "🚀 Listo para implementar los tipos de negocio restantes!" 