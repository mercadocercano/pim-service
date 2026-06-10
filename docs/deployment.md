# Deployment — PIM Service

## Variables de entorno

| Variable | Default | Descripción |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | Host de PostgreSQL |
| `DB_PORT` | `5432` | Puerto de PostgreSQL |
| `DB_USER` | `postgres` | Usuario de BD |
| `DB_PASSWORD` | `postgres` | Contraseña de BD |
| `DB_NAME` | `pim_db` | Nombre de la base de datos |
| `PORT` | `8090` | Puerto HTTP del servicio |
| `GIN_MODE` | `debug` | `debug` \| `release` |
| `PROMETHEUS_ENABLED` | `false` | `true` para exponer `/metrics` |
| `JWT_SECRET` | — | Secreto para validar JWTs del tenant |

---

## Desarrollo local

```bash
# Crear la BD si no existe
psql -U postgres -c "CREATE DATABASE pim_db;"

# Correr migraciones
./scripts/migrate.sh

# Iniciar el servicio
go run main.go
# → http://localhost:8090/api/v1/health
# → http://localhost:8090/api-docs  (Swagger UI)
```

---

## Laboratorio (Docker + lab-network)

El servicio corre en el laboratorio compartido contra `lab-postgres:5432`.

```bash
# Desde la raíz del lab
cd ~/Projects && make infra   # levanta postgres, redis, kong
make mc                        # levanta iam-service + pim-service

# O solo pim-service
cd ~/Projects/active/mercado-cercano
docker compose up -d pim-service
```

El contenedor se llama `mc-pim-service` y expone el puerto 8090.
Kong rutea `/pim/api/v1` → `mc-pim-service:8090/api/v1`.

---

## Build de producción

```bash
# Binary estático
go build -ldflags="-s -w" -o pim-service main.go

# Docker
docker build -t pim-service:latest .
docker run -p 8090:8090 \
  -e DB_HOST=<host> \
  -e DB_USER=<user> \
  -e DB_PASSWORD=<pass> \
  -e DB_NAME=pim_db \
  -e JWT_SECRET=<secret> \
  -e GIN_MODE=release \
  pim-service:latest
```

---

## CI/CD

Push a `main` → `.github/workflows/deploy-pim.yml` (build → test → deploy).

---

## Migraciones en producción

```bash
./scripts/migrate.sh
```

Las migraciones son **evolutivas, nunca destructivas**. No hay rollback automático — para revertir, crear una nueva migración que deshaga el cambio.

---

## Monitoreo

| Endpoint | Descripción |
|----------|-------------|
| `GET /api/v1/health` | Estado del servicio + conexión BD |
| `GET /metrics` | Métricas Prometheus (requiere `PROMETHEUS_ENABLED=true`) |
| `GET /api-docs` | Swagger UI (no usar en producción) |

Si Prometheus está habilitado, las métricas del servicio se registran via `PrometheusRecorder` (go-shared).
Grafana en el lab: http://localhost:3002 (admin / admin123).
