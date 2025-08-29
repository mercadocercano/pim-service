#!/bin/bash

# Script para ejecutar la migración usando Docker

echo "🚀 Migración de Scraper Products a Global Products (Docker)"
echo "========================================================="

# Verificar conexiones primero
echo -e "\n🔍 Verificando conexiones..."

# MongoDB
MONGO_COUNT=$(docker exec -i dev-mongodb mongosh -u admin -p admin123 --authenticationDatabase admin pim_marketplace --quiet --eval "db.scraper_products.countDocuments({})" 2>/dev/null | tail -1)
if [ -z "$MONGO_COUNT" ]; then
    echo "❌ Error: No se pudo conectar a MongoDB"
    exit 1
fi
echo "✅ MongoDB conectado. Total productos: $MONGO_COUNT"

NOT_MIGRATED=$(docker exec -i dev-mongodb mongosh -u admin -p admin123 --authenticationDatabase admin pim_marketplace --quiet --eval "db.scraper_products.countDocuments({migrated_to_pim: {\$ne: true}})" 2>/dev/null | tail -1)
echo "📊 Productos pendientes de migrar: $NOT_MIGRATED"

# PostgreSQL
PG_COUNT=$(docker exec -i dev-postgres psql -U postgres -d pim_db -t -c "SELECT COUNT(*) FROM global_products;" 2>/dev/null | xargs)
echo "✅ PostgreSQL conectado. Productos actuales: $PG_COUNT"

# Preguntar si continuar
echo -e "\n⚠️  Esta operación migrará productos de MongoDB a PostgreSQL."
if [ "$1" == "--yes" ]; then
    REPLY="s"
else
    read -p "¿Deseas continuar? (s/N): " -n 1 -r
    echo
fi

if [[ $REPLY =~ ^[Ss]$ ]]; then
    # Ejecutar migración en contenedor Python
    echo -e "\n🐳 Creando contenedor para migración..."
    
    LIMIT=${MIGRATION_LIMIT:-10}
    
    docker run --rm \
        --name pim-migration \
        --network saas-mt_dev-network \
        -v "$(pwd):/scripts" \
        -w /scripts \
        -e MIGRATION_LIMIT=$LIMIT \
        python:3.11-slim bash -c "
            echo '📦 Instalando dependencias...'
            pip install --quiet pymongo psycopg2-binary
            
            LIMIT=\${MIGRATION_LIMIT:-10}
            echo -e \"\n🔄 Ejecutando migración (límite: \$LIMIT productos)...\"
            export MONGO_URI='mongodb://admin:admin123@dev-mongodb:27017/'
            export POSTGRES_HOST='dev-postgres'
            python3 migrate_scraper_to_global_products.py \$LIMIT
        "
    
    # Mostrar resultados
    echo -e "\n📊 Resultado de la migración:"
    
    NEW_PG_COUNT=$(docker exec -i dev-postgres psql -U postgres -d pim_db -t -c "SELECT COUNT(*) FROM global_products;" 2>/dev/null | xargs)
    echo "✅ Total en global_products: $NEW_PG_COUNT (antes: $PG_COUNT)"
    
    MIGRATED=$(docker exec -i dev-mongodb mongosh -u admin -p admin123 --authenticationDatabase admin pim_marketplace --quiet --eval "db.scraper_products.countDocuments({migrated_to_pim: true})" 2>/dev/null | tail -1)
    echo "✅ Productos marcados como migrados: $MIGRATED"
    
    # Mostrar algunos productos migrados
    echo -e "\n📋 Últimos productos migrados:"
    docker exec -i dev-postgres psql -U postgres -d pim_db -c "
        SELECT ean, name, brand, quality_score 
        FROM global_products 
        ORDER BY created_at DESC 
        LIMIT 5;
    "
    
    echo -e "\n✅ Migración de prueba completada!"
    echo "💡 Para migrar todos los productos, edita el script y quita el límite de 10"
    
else
    echo "❌ Migración cancelada"
fi