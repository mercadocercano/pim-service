# Guía de Instalación del Sistema PIM

## Requisitos del Sistema

### Software Requerido

- **Docker**: 20.10 o superior
- **Docker Compose**: 2.0 o superior
- **Go**: 1.22 o superior (para desarrollo local)
- **PostgreSQL**: 15 o superior (si se ejecuta localmente)
- **Git**: Para clonar el repositorio

### Recursos Mínimos

- **RAM**: 4GB mínimo, 8GB recomendado
- **CPU**: 2 cores mínimo, 4 cores recomendado
- **Almacenamiento**: 10GB disponibles
- **Red**: Puertos 8001, 8090, 5432, 3000-3002, 9090 disponibles

## Instalación con Docker Compose (Recomendado)

### 1. Clonar el Repositorio

```bash
git clone <repository-url>
cd saas/
```

### 2. Configurar Variables de Entorno

Crear archivo `.env` en la raíz del proyecto:

```bash
# Base de datos
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=pim_db

# PIM Service
PIM_PORT=8090
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=pim_db
DB_SSLMODE=disable
GIN_MODE=debug

# Kong API Gateway
KONG_PORT=8001
KONG_ADMIN_PORT=8444

# Monitoreo
PROMETHEUS_ENABLED=true
GRAFANA_PORT=3002
PROMETHEUS_PORT=9090

# Frontend
FRONTEND_PORT=3001
BACKOFFICE_PORT=3000

# Logs
LOKI_PORT=3100
```

### 3. Levantar los Servicios

```bash
# Levantar todos los servicios
docker-compose up -d

# Verificar que todos los servicios estén ejecutándose
docker-compose ps
```

### 4. Verificar la Instalación

```bash
# Health check del servicio PIM
curl http://localhost:8090/api/v1/health

# Health check a través del API Gateway
curl http://localhost:8001/pim/api/v1/health

# Verificar documentación OpenAPI
curl http://localhost:8001/pim/api-docs
```

### 5. Servicios Disponibles

Una vez instalado, los siguientes servicios estarán disponibles:

| Servicio | URL | Descripción |
|----------|-----|-------------|
| PIM API | http://localhost:8090 | Servicio PIM directo |
| API Gateway | http://localhost:8001 | Kong API Gateway |
| Documentación | http://localhost:8001/pim/api-docs | Swagger UI |
| Frontend | http://localhost:3001 | Frontend principal |
| Backoffice | http://localhost:3000 | Panel de administración |
| Grafana | http://localhost:3002 | Dashboards de monitoreo |
| Prometheus | http://localhost:9090 | Métricas del sistema |

## Instalación para Desarrollo Local

### 1. Configurar Base de Datos

```bash
# Levantar solo PostgreSQL
docker-compose up -d postgres

# Esperar a que PostgreSQL esté listo
docker-compose logs postgres
```

### 2. Configurar Variables de Entorno Local

Crear archivo `.env.local`:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=pim_db
DB_SSLMODE=disable
GIN_MODE=debug
PROMETHEUS_ENABLED=false
```

### 3. Ejecutar Migraciones

```bash
cd pim/

# Instalar dependencias
go mod download

# Ejecutar migraciones
go run migrations/migrate.go up
```

### 4. Ejecutar el Servicio

```bash
# Desde el directorio pim/
go run main.go
```

### 5. Verificar Instalación Local

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Documentación
open http://localhost:8080/api-docs
```

## Configuración de Base de Datos

### Migraciones Automáticas

El sistema incluye migraciones automáticas que se ejecutan al iniciar:

```sql
-- 001_create_categories_table.sql
-- 002_create_brands_table.sql
-- 003_create_products_table.sql
-- 004_create_product_variants_table.sql
-- 005_create_variant_attributes_table.sql
```

### Seeds de Datos

Para cargar datos de prueba:

```bash
# Ejecutar seeds
docker-compose exec pim-service psql -h postgres -U postgres -d pim_db -f /app/seeds/categories.sql
docker-compose exec pim-service psql -h postgres -U postgres -d pim_db -f /app/seeds/brands.sql
docker-compose exec pim-service psql -h postgres -U postgres -d pim_db -f /app/seeds/products.sql
docker-compose exec pim-service psql -h postgres -U postgres -d pim_db -f /app/seeds/product_variants.sql
```

### Backup y Restore

```bash
# Crear backup
docker-compose exec postgres pg_dump -U postgres pim_db > backup.sql

# Restaurar backup
docker-compose exec -T postgres psql -U postgres pim_db < backup.sql
```

## Configuración del API Gateway (Kong)

### Configuración Automática

Kong se configura automáticamente con:

- **Servicio PIM**: Upstream al servicio PIM
- **Rutas**: Todas las rutas `/pim/*` se enrutan al servicio PIM
- **Plugins**: Rate limiting, CORS, logging

### Configuración Manual (Opcional)

```bash
# Agregar servicio
curl -X POST http://localhost:8444/services \
  --data name=pim-service \
  --data url=http://pim-service:8080

# Agregar ruta
curl -X POST http://localhost:8444/services/pim-service/routes \
  --data paths[]=/pim \
  --data strip_path=true
```

## Configuración de Monitoreo

### Prometheus

Métricas disponibles en: http://localhost:9090

Configuración en `prometheus.yml`:

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'pim-service'
    static_configs:
      - targets: ['pim-service:8080']
    metrics_path: '/metrics'
```

### Grafana

Dashboards disponibles en: http://localhost:3002

Credenciales por defecto:
- Usuario: `admin`
- Contraseña: `admin`

### Loki (Logs)

Logs centralizados en: http://localhost:3100

## Configuración Multi-Tenant

### Crear Tenant

```bash
# Generar UUID para tenant
TENANT_ID=$(uuidgen)
echo "Tenant ID: $TENANT_ID"

# Usar en peticiones
curl -X POST http://localhost:8001/pim/api/v1/products \
  -H "X-Tenant-ID: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -d '{"name": "Producto de Prueba"}'
```

### Validar Aislamiento

```bash
# Crear producto para tenant 1
TENANT_1=$(uuidgen)
curl -X POST http://localhost:8001/pim/api/v1/products \
  -H "X-Tenant-ID: $TENANT_1" \
  -H "Content-Type: application/json" \
  -d '{"name": "Producto Tenant 1"}'

# Crear producto para tenant 2
TENANT_2=$(uuidgen)
curl -X POST http://localhost:8001/pim/api/v1/products \
  -H "X-Tenant-ID: $TENANT_2" \
  -H "Content-Type: application/json" \
  -d '{"name": "Producto Tenant 2"}'

# Verificar que cada tenant solo ve sus productos
curl -X GET http://localhost:8001/pim/api/v1/products \
  -H "X-Tenant-ID: $TENANT_1"

curl -X GET http://localhost:8001/pim/api/v1/products \
  -H "X-Tenant-ID: $TENANT_2"
```

## Troubleshooting

### Problemas Comunes

#### 1. Puerto en Uso

```bash
# Verificar puertos en uso
netstat -tulpn | grep :8090

# Cambiar puerto en docker-compose.yml
ports:
  - "8091:8080"  # Cambiar puerto externo
```

#### 2. Base de Datos No Conecta

```bash
# Verificar logs de PostgreSQL
docker-compose logs postgres

# Verificar conectividad
docker-compose exec pim-service pg_isready -h postgres -p 5432
```

#### 3. Migraciones Fallan

```bash
# Verificar estado de migraciones
docker-compose exec postgres psql -U postgres -d pim_db -c "\dt"

# Ejecutar migraciones manualmente
docker-compose exec pim-service /app/main migrate
```

#### 4. Kong No Responde

```bash
# Verificar logs de Kong
docker-compose logs api-gateway

# Verificar configuración
curl http://localhost:8444/services
```

### Logs y Debugging

```bash
# Ver logs de todos los servicios
docker-compose logs -f

# Ver logs de un servicio específico
docker-compose logs -f pim-service

# Ver logs en tiempo real
docker-compose logs -f --tail=100 pim-service
```

### Reiniciar Servicios

```bash
# Reiniciar servicio específico
docker-compose restart pim-service

# Reiniciar todos los servicios
docker-compose restart

# Reconstruir y reiniciar
docker-compose down
docker-compose build
docker-compose up -d
```

## Configuración de Producción

### Variables de Entorno de Producción

```bash
# Seguridad
GIN_MODE=release
DB_SSLMODE=require
POSTGRES_PASSWORD=<contraseña-segura>

# Performance
DB_MAX_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10

# Monitoreo
PROMETHEUS_ENABLED=true
LOG_LEVEL=info
```

### Consideraciones de Seguridad

1. **Cambiar contraseñas por defecto**
2. **Configurar SSL/TLS**
3. **Configurar firewall**
4. **Usar secrets management**
5. **Configurar backup automático**

### Escalabilidad

```yaml
# docker-compose.prod.yml
services:
  pim-service:
    deploy:
      replicas: 3
      resources:
        limits:
          memory: 512M
        reservations:
          memory: 256M
```

### Health Checks

```yaml
services:
  pim-service:
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/api/v1/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

## Actualización del Sistema

### Actualizar Código

```bash
# Detener servicios
docker-compose down

# Actualizar código
git pull origin main

# Reconstruir imágenes
docker-compose build

# Levantar servicios
docker-compose up -d
```

### Ejecutar Nuevas Migraciones

```bash
# Las migraciones se ejecutan automáticamente al iniciar
# Verificar en logs
docker-compose logs pim-service | grep migration
```

### Rollback

```bash
# Volver a versión anterior
git checkout <commit-anterior>
docker-compose build
docker-compose up -d

# Rollback de migraciones (manual)
docker-compose exec postgres psql -U postgres -d pim_db
# Ejecutar rollback SQL manualmente
```

## Monitoreo Post-Instalación

### Métricas Clave

- **Response Time**: < 200ms promedio
- **Error Rate**: < 1%
- **CPU Usage**: < 70%
- **Memory Usage**: < 80%
- **Database Connections**: < 80% del máximo

### Alertas Recomendadas

1. **Servicio caído**
2. **Alto tiempo de respuesta**
3. **Errores 5xx frecuentes**
4. **Uso alto de recursos**
5. **Fallos de conexión a BD**

Con esta guía, el sistema PIM debería estar completamente funcional y listo para uso en desarrollo o producción. 