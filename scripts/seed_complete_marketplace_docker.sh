#!/bin/bash

# Script: seed_complete_marketplace_docker.sh
# Purpose: Ejecutar seed completo de marketplace argentino usando Docker Compose
# Incluye: business_types, business_types adicionales y marketplace_categories
# Usage: ./scripts/seed_complete_marketplace_docker.sh

set -e

echo "🛍️ Iniciando seed completo del marketplace argentino (Docker)..."

# Verificar si docker-compose está disponible
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Error: docker-compose no está instalado"
    exit 1
fi

# Verificar si estamos en el directorio correcto
if [ ! -f "docker-compose.yml" ]; then
    echo "❌ Error: Ejecuta este script desde el directorio del servicio PIM"
    exit 1
fi

# Buscar el contenedor de PostgreSQL
PG_CONTAINER=$(docker-compose ps -q postgres 2>/dev/null || docker-compose ps -q db 2>/dev/null || echo "")

if [ -z "$PG_CONTAINER" ]; then
    echo "❌ Error: No se encontró contenedor de PostgreSQL"
    echo "Asegúrate de que docker-compose esté corriendo"
    exit 1
fi

echo "📊 Ejecutando seeds en contenedor PostgreSQL..."

# Ejecutar seed de business_types argentinos
echo "🏪 1/3 - Cargando tipos de comercio argentinos..."
docker-compose exec -T postgres psql -U postgres -d pim_db < seeds/001_business_types_argentina_seed.sql

if [ $? -ne 0 ]; then
    echo "❌ Error al cargar tipos de comercio"
    exit 1
fi

echo "✅ Tipos de comercio cargados exitosamente"

# Ejecutar seed de business_types adicionales
echo "🚀 2/3 - Agregando tipos de comercio adicionales..."
docker-compose exec -T postgres psql -U postgres -d pim_db < seeds/002_business_types_additional_seed.sql

if [ $? -ne 0 ]; then
    echo "❌ Error al cargar tipos de comercio adicionales"
    exit 1
fi

echo "✅ Tipos de comercio adicionales cargados exitosamente"

# Ejecutar seed de marketplace_categories
echo "📂 3/4 - Cargando categorías globales del marketplace..."
docker-compose exec -T postgres psql -U postgres -d pim_db < seeds/003_marketplace_categories_argentina_seed.sql

if [ $? -ne 0 ]; then
    echo "❌ Error al cargar categorías del marketplace"
    exit 1
fi

echo "✅ Categorías del marketplace cargadas exitosamente"

# Ejecutar seed de business_type_templates (quickstart)
echo "🚀 4/6 - Creando templates de quickstart (vinculando tipos con categorías)..."
docker-compose exec -T postgres psql -U postgres -d pim_db < seeds/004_business_type_quickstart_templates.sql

if [ $? -ne 0 ]; then
    echo "❌ Error al cargar templates de quickstart"
    exit 1
fi

echo "✅ Templates de quickstart creados exitosamente"

# Ejecutar migración de catálogo global
echo "🌍 5/6 - Creando estructura de catálogo global de productos..."
docker-compose exec -T postgres psql -U postgres -d pim_db < migrations/005_global_catalog_structure.sql

if [ $? -ne 0 ]; then
    echo "❌ Error al crear estructura de catálogo global"
    exit 1
fi

echo "✅ Estructura de catálogo global creada exitosamente"

# Ejecutar seed de productos globales
echo "📦 6/6 - Cargando productos populares argentinos..."
docker-compose exec -T postgres psql -U postgres -d pim_db < seeds/006_global_products_seed.sql

if [ $? -ne 0 ]; then
    echo "❌ Error al cargar productos globales"
    exit 1
fi

echo "✅ Productos globales cargados exitosamente"

# Verificar los datos insertados
echo "🔍 Verificando datos insertados..."

echo ""
echo "📋 TIPOS DE COMERCIO CARGADOS:"
docker-compose exec -T postgres psql -U postgres -d pim_db -c "SELECT code, name FROM business_types ORDER BY sort_order;"

echo ""
echo "📂 CATEGORÍAS PRINCIPALES DEL MARKETPLACE:"
docker-compose exec -T postgres psql -U postgres -d pim_db -c "SELECT name, slug FROM marketplace_categories WHERE parent_id IS NULL ORDER BY sort_order LIMIT 10;"

echo ""
echo "📊 RESUMEN:"
echo "$(docker-compose exec -T postgres psql -U postgres -d pim_db -t -c "SELECT COUNT(*) FROM business_types;") tipos de comercio cargados"
echo "$(docker-compose exec -T postgres psql -U postgres -d pim_db -t -c "SELECT COUNT(*) FROM marketplace_categories WHERE parent_id IS NULL;") categorías principales cargadas"
echo "$(docker-compose exec -T postgres psql -U postgres -d pim_db -t -c "SELECT COUNT(*) FROM marketplace_categories WHERE parent_id IS NOT NULL;") subcategorías cargadas"
echo "$(docker-compose exec -T postgres psql -U postgres -d pim_db -t -c "SELECT COUNT(*) FROM business_type_templates;") templates de quickstart creados"
echo "$(docker-compose exec -T postgres psql -U postgres -d pim_db -t -c "SELECT COUNT(*) FROM global_products;") productos en catálogo global"
echo "$(docker-compose exec -T postgres psql -U postgres -d pim_db -t -c "SELECT COUNT(*) FROM business_type_product_templates;") templates de productos configurados"

echo ""
echo "🎉 ¡Seed completo del marketplace argentino terminado!"
echo "🏪 35 tipos de comercio físicos y digitales"
echo "📂 14 categorías principales con múltiples subcategorías"
echo "🚀 35 templates de quickstart con mapeos inteligentes"
echo "📦 30+ productos populares argentinos en catálogo global"
echo "🔗 Templates de productos vinculados por tipo de negocio"
echo "📈 Basado en estudios de mercado argentino 2024-2025"
echo ""
echo "✨ FUNCIONALIDAD QUICKSTART AVANZADA LISTA:"
echo "   • Seller elige tipo de negocio en onboarding"
echo "   • Sistema sugiere categorías relevantes automáticamente"
echo "   • Muestra productos populares para ese tipo de negocio"
echo "   • Seller puede vincular productos del catálogo global"
echo "   • Solo necesita agregar precio y stock → producto online en segundos"
echo ""
echo "🌟 PRÓXIMOS PASOS:"
echo "   1. Implementar APIs de catálogo global en PIM Service"
echo "   2. Desarrollar scraper para expandir catálogo automáticamente"  
echo "   3. Integrar frontend con flujo de productos sugeridos"
echo "   4. Configurar búsqueda por EAN para scaneo de códigos" 