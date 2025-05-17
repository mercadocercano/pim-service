#!/bin/bash

# Las variables de entorno ya deberían estar disponibles en el contenedor Docker
# Directorio de migraciones
MIGRATIONS_DIR="/migrations"

echo "Migrando base de datos $POSTGRES_DB en $POSTGRES_HOST:$POSTGRES_PORT como usuario $POSTGRES_USER"

# Ejecutar migraciones
for migration in $(ls $MIGRATIONS_DIR/*.sql | grep -v "seed_" | sort); do
    echo "Ejecutando migración: $migration"
    PGPASSWORD=$POSTGRES_PASSWORD psql -h $POSTGRES_HOST -p $POSTGRES_PORT -U $POSTGRES_USER -d $POSTGRES_DB -f $migration
done

# Ejecutar seeds después de todas las migraciones
for seed in $(ls $MIGRATIONS_DIR/seed_*.sql | sort); do
    echo "Ejecutando seed: $seed"
    PGPASSWORD=$POSTGRES_PASSWORD psql -h $POSTGRES_HOST -p $POSTGRES_PORT -U $POSTGRES_USER -d $POSTGRES_DB -f $seed
done

echo "Migraciones y seeds completados" 