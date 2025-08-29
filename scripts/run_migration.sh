#!/bin/bash

# Script para ejecutar la migración de scraper_products a global_products

echo "🚀 Migración de Scraper Products a Global Products"
echo "=================================================="

# Instalar dependencias si no existen
if ! python3 -c "import pymongo" 2>/dev/null; then
    echo "📦 Instalando dependencias Python..."
    pip3 install -r requirements_migration.txt
fi

# Variables de entorno
export MONGO_URI="mongodb://admin:admin123@localhost:27017/"
export MONGO_DB="pim_marketplace"
export POSTGRES_HOST="localhost"
export POSTGRES_PORT="5432"
export POSTGRES_DB="pim_db"
export POSTGRES_USER="postgres"
export POSTGRES_PASSWORD="postgres"

# Verificar conexión a MongoDB
echo -e "\n🔍 Verificando conexión a MongoDB..."
python3 -c "
from pymongo import MongoClient
client = MongoClient('$MONGO_URI')
db = client['$MONGO_DB']
count = db['scraper_products'].count_documents({})
print(f'✅ MongoDB conectado. Total de productos en scraper_products: {count}')
not_migrated = db['scraper_products'].count_documents({'migrated_to_pim': {'$ne': True}})
print(f'📊 Productos pendientes de migrar: {not_migrated}')
"

# Verificar conexión a PostgreSQL
echo -e "\n🔍 Verificando conexión a PostgreSQL..."
docker exec -i dev-postgres psql -U postgres -d pim_db -c "SELECT COUNT(*) as total FROM global_products;" | grep -E "^\s*[0-9]+" | xargs -I {} echo "✅ PostgreSQL conectado. Productos en global_products: {}"

# Preguntar si ejecutar migración
echo -e "\n⚠️  Esta operación migrará productos de MongoDB a PostgreSQL."
read -p "¿Deseas continuar? (s/N): " -n 1 -r
echo

if [[ $REPLY =~ ^[Ss]$ ]]; then
    # Preguntar límite
    echo -e "\n📊 ¿Cuántos productos deseas migrar?"
    echo "   - Presiona ENTER para migrar todos"
    echo "   - O ingresa un número para limitar la migración"
    read -p "Límite: " LIMIT
    
    echo -e "\n🔄 Iniciando migración..."
    
    if [ -z "$LIMIT" ]; then
        python3 migrate_scraper_to_global_products.py
    else
        python3 migrate_scraper_to_global_products.py $LIMIT
    fi
    
    # Mostrar resultado final
    echo -e "\n📊 Estado final:"
    docker exec -i dev-postgres psql -U postgres -d pim_db -c "SELECT COUNT(*) as total FROM global_products;" | grep -E "^\s*[0-9]+" | xargs -I {} echo "✅ Total en global_products: {}"
    
    python3 -c "
from pymongo import MongoClient
client = MongoClient('$MONGO_URI')
db = client['$MONGO_DB']
migrated = db['scraper_products'].count_documents({'migrated_to_pim': True})
not_migrated = db['scraper_products'].count_documents({'migrated_to_pim': {'$ne': True}})
print(f'✅ Productos migrados en MongoDB: {migrated}')
print(f'⏳ Productos pendientes: {not_migrated}')
"
else
    echo "❌ Migración cancelada"
fi