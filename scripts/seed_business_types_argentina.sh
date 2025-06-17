#!/bin/bash

# Script: seed_business_types_argentina.sh
# Purpose: Ejecutar el seed de tipos de comercio argentinos
# Usage: ./scripts/seed_business_types_argentina.sh

set -e

echo "🏪 Iniciando seed de tipos de comercio argentinos..."

# Variables de configuración de la base de datos
DB_HOST=${DB_HOST:-"localhost"}
DB_PORT=${DB_PORT:-"5432"}
DB_NAME=${DB_NAME:-"pim_db"}
DB_USER=${DB_USER:-"postgres"}
DB_PASSWORD=${DB_PASSWORD:-"postgres"}

# Verificar si psql está disponible
if ! command -v psql &> /dev/null; then
    echo "❌ Error: psql no está instalado o no está en el PATH"
    echo "Instala PostgreSQL client o ejecuta el seed manualmente"
    exit 1
fi

echo "📊 Conectando a la base de datos: $DB_HOST:$DB_PORT/$DB_NAME"

# Ejecutar el seed
echo "🧹 Limpiando datos existentes e insertando tipos de comercio argentinos..."

PGPASSWORD=$DB_PASSWORD psql \
    -h $DB_HOST \
    -p $DB_PORT \
    -d $DB_NAME \
    -U $DB_USER \
    -f seeds/001_business_types_argentina_seed.sql

if [ $? -eq 0 ]; then
    echo "✅ Seed ejecutado exitosamente!"
    echo "📋 Se insertaron 28 tipos de comercio argentinos"
    
    # Verificar los datos insertados
    echo "🔍 Verificando datos insertados..."
    PGPASSWORD=$DB_PASSWORD psql \
        -h $DB_HOST \
        -p $DB_PORT \
        -d $DB_NAME \
        -U $DB_USER \
        -c "SELECT code, name FROM business_types ORDER BY sort_order LIMIT 10;"
        
else
    echo "❌ Error al ejecutar el seed"
    exit 1
fi

echo "🎉 ¡Seed completado! Los tipos de comercio argentinos están listos." 