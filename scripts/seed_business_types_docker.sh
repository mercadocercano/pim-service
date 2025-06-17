#!/bin/bash

# Script: seed_business_types_docker.sh
# Purpose: Ejecutar seed de tipos de comercio argentinos usando Docker Compose
# Usage: ./scripts/seed_business_types_docker.sh

set -e

echo "🏪 Iniciando seed de tipos de comercio argentinos (Docker)..."

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

echo "📊 Ejecutando seed en contenedor PostgreSQL..."

# Ejecutar el seed
echo "🧹 Limpiando datos existentes e insertando tipos de comercio argentinos..."

docker-compose exec -T postgres psql -U postgres -d pim_db < seeds/001_business_types_argentina_seed.sql

if [ $? -eq 0 ]; then
    echo "✅ Seed ejecutado exitosamente!"
    echo "📋 Se insertaron 28 tipos de comercio argentinos"
    
    # Verificar los datos insertados
    echo "🔍 Verificando datos insertados..."
    docker-compose exec -T postgres psql -U postgres -d pim_db -c "SELECT code, name FROM business_types ORDER BY sort_order LIMIT 10;"
        
else
    echo "❌ Error al ejecutar el seed"
    exit 1
fi

echo "🎉 ¡Seed completado! Los tipos de comercio argentinos están listos." 