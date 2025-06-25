#!/bin/bash

# Script para actualizar imports de la estructura antigua a la nueva
# Opción 2: Actualización automática de imports

echo "🔄 Actualizando imports de la estructura de productos..."

# Función para actualizar imports en un archivo
update_imports_in_file() {
    local file="$1"
    local temp_file="${file}.tmp"
    
    echo "  📝 Procesando: $file"
    
    # Crear copia temporal
    cp "$file" "$temp_file"
    
    # Aplicar transformaciones específicas una por una
    # Product tenant module
    sed -i '' 's|"pim/src/product/application/usecase"|"pim/src/product/tenant/application/usecase"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/application/request"|"pim/src/product/tenant/application/request"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/application/response"|"pim/src/product/tenant/application/response"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/application/mapper"|"pim/src/product/tenant/application/mapper"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/domain/entity"|"pim/src/product/tenant/domain/entity"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/domain/value_object"|"pim/src/product/tenant/domain/value_object"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/domain/port"|"pim/src/product/tenant/domain/port"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/domain/service"|"pim/src/product/tenant/domain/service"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/infrastructure/controller"|"pim/src/product/tenant/infrastructure/controller"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/infrastructure/persistence"|"pim/src/product/tenant/infrastructure/persistence"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/infrastructure/config"|"pim/src/product/tenant/infrastructure/config"|g' "$temp_file"
    sed -i '' 's|"pim/src/product/infrastructure/criteria"|"pim/src/product/tenant/infrastructure/criteria"|g' "$temp_file"
    
    # Global catalog module
    sed -i '' 's|"pim/src/global_catalog/application/usecase"|"pim/src/product/global_catalog/application/usecase"|g' "$temp_file"
    sed -i '' 's|"pim/src/global_catalog/infrastructure/controller"|"pim/src/product/global_catalog/infrastructure/controller"|g' "$temp_file"
    sed -i '' 's|"pim/src/global_catalog/infrastructure/persistence"|"pim/src/product/global_catalog/infrastructure/persistence"|g' "$temp_file"
    sed -i '' 's|"pim/src/global_catalog/infrastructure/config"|"pim/src/product/global_catalog/infrastructure/config"|g' "$temp_file"
    sed -i '' 's|"pim/src/global_catalog/domain/entity"|"pim/src/product/global_catalog/domain/entity"|g' "$temp_file"
    sed -i '' 's|"pim/src/global_catalog/domain/exception"|"pim/src/product/global_catalog/domain/exception"|g' "$temp_file"
    sed -i '' 's|"pim/src/global_catalog/domain/port"|"pim/src/product/global_catalog/domain/port"|g' "$temp_file"
    sed -i '' 's|"pim/src/global_catalog/domain/value_object"|"pim/src/product/global_catalog/domain/value_object"|g' "$temp_file"
    
    # Solo sobrescribir si hubo cambios
    if ! cmp -s "$file" "$temp_file"; then
        mv "$temp_file" "$file"
        echo "    ✅ Actualizado"
        return 0
    else
        rm "$temp_file"
        echo "    ⏭️  Sin cambios"
        return 1
    fi
}

# Buscar y actualizar todos los archivos .go
echo "🔍 Buscando archivos Go para actualizar..."

updated_count=0
total_count=0

find src -name "*.go" -type f | while read -r file; do
    total_count=$((total_count + 1))
    if update_imports_in_file "$file"; then
        updated_count=$((updated_count + 1))
    fi
done

echo ""
echo "✅ Actualización de imports completada!"
echo "📊 Archivos procesados: $total_count"
echo "📊 Archivos actualizados: $updated_count"
echo ""
echo "🧪 Probando compilación..."

# Intentar compilar para verificar
if go build -o tmp/main main.go 2>/dev/null; then
    echo "✅ Compilación exitosa!"
    rm -f tmp/main
else
    echo "❌ Errores de compilación detectados. Ejecutando go build para ver detalles..."
    go build -o tmp/main main.go
fi

echo ""
echo "📋 Transformaciones aplicadas:"
echo "  pim/src/product/application/* -> pim/src/product/tenant/application/*"
echo "  pim/src/product/domain/* -> pim/src/product/tenant/domain/*"
echo "  pim/src/product/infrastructure/* -> pim/src/product/tenant/infrastructure/*"
echo "  pim/src/global_catalog/* -> pim/src/product/global_catalog/*" 