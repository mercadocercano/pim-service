#!/bin/bash

# Run Marketplace Seeders Script
# PROPÓSITO: Ejecutar todos los seeders marketplace en orden correcto
# USO: ./scripts/run-marketplace-seeders.sh

set -e  # Exit on any error

echo "🚀 INICIANDO SEEDERS MARKETPLACE ARGENTINA"
echo "=========================================="

# Database connection variables
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_NAME=${DB_NAME:-pim_saas_mt}
DB_USER=${DB_USER:-postgres}

echo "📊 Base de datos: $DB_NAME@$DB_HOST:$DB_PORT"
echo "👤 Usuario: $DB_USER"
echo ""

# Function to execute SQL file
execute_sql() {
    local file_path=$1
    local description=$2
    
    echo "📄 Ejecutando: $description"
    echo "   Archivo: $file_path"
    
    if [ ! -f "$file_path" ]; then
        echo "❌ ERROR: Archivo no encontrado: $file_path"
        exit 1
    fi
    
    # Execute SQL file
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$file_path"
    
    if [ $? -eq 0 ]; then
        echo "✅ Completado exitosamente"
    else
        echo "❌ ERROR: Falló la ejecución del seeder"
        exit 1
    fi
    echo ""
}

# Check if database is accessible
echo "🔍 Verificando conexión a base de datos..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "SELECT 1;" > /dev/null

if [ $? -ne 0 ]; then
    echo "❌ ERROR: No se puede conectar a la base de datos"
    echo "   Verifica que el servicio esté corriendo y las credenciales sean correctas"
    exit 1
fi
echo "✅ Conexión exitosa"
echo ""

# Execute seeders in correct order
echo "🌱 EJECUTANDO SEEDERS EN ORDEN..."
echo ""

# 1. Categories seeder
execute_sql "seeds/013_marketplace_categories_seeder.sql" "Categorías Marketplace (50 categorías en 3 niveles)"

# 2. Attributes seeder  
execute_sql "seeds/014_marketplace_attributes_argentina.sql" "Atributos Argentina (15+ atributos con 200+ valores)"

# 3. Category-Attributes relations
execute_sql "seeds/015_marketplace_category_attributes_relations.sql" "Relaciones Categoría-Atributo (200+ mapeos)"

echo "🎉 SEEDERS COMPLETADOS EXITOSAMENTE!"
echo "===================================="
echo ""
echo "📊 RESUMEN DE DATOS CARGADOS:"
echo "• 50 categorías marketplace organizadas en 3 niveles"
echo "• 15+ atributos específicos para Argentina"
echo "• 200+ valores de atributos localizados"
echo "• 200+ relaciones categoría-atributo configuradas"
echo ""
echo "🔍 PRÓXIMOS PASOS:"
echo "• Verificar datos: psql -c 'SELECT count(*) FROM marketplace_categories;'"
echo "• Ver atributos: psql -c 'SELECT count(*) FROM marketplace_attributes;'"
echo "• Testear relaciones: psql -c 'SELECT count(*) FROM marketplace_category_attributes;'"
echo ""
echo "🌐 ENDPOINTS DISPONIBLES:"
echo "• GET /marketplace/categories - Listar categorías"
echo "• GET /marketplace/attributes - Listar atributos" 
echo "• GET /marketplace/categories/{id}/attributes - Atributos de categoría"
echo ""
echo "✨ Sistema marketplace listo para usar con datos argentinos!" 