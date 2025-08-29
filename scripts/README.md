# 🛠️ Scripts de Utilidades - PIM Service

Este directorio contiene scripts de utilidad para el servicio PIM.

## 📦 Scripts de Migración y Seeds

### `migrate.sh`
Script principal de migración de base de datos PostgreSQL.

### `run_migration.sh`
Ejecuta migraciones específicas del servicio.

### `run_migration_docker.sh`
Ejecuta migraciones en el entorno Docker.

### Seeds de Datos
- `seed_business_types_argentina.sh` - Carga datos de tipos de negocio para Argentina
- `seed_business_types_docker.sh` - Ejecuta seeds de business types en Docker
- `seed_complete_marketplace_docker.sh` - Carga completa de datos del marketplace en Docker
- `run-marketplace-seeders.sh` - Ejecuta todos los seeders del marketplace

## 🔧 Scripts de Utilidad

### `wait-for-db.sh`
Script que espera a que PostgreSQL esté disponible antes de continuar. Usado en Docker Compose.

### `update-openapi.sh`
Actualiza la documentación OpenAPI del servicio.

### `setup-ai-templates-e2e.sh`
Configura el entorno para pruebas E2E de templates AI.

## 📦 Tests de Integración

**⚠️ IMPORTANTE**: Los tests de integración están en el directorio `test-integration/`

Para ejecutar tests de integración:

```bash
# Desde el directorio de tests
cd test-integration/
./run_integration_tests.sh
```

Ver documentación completa en: [`test-integration/README.md`](../test-integration/README.md)

## 📋 Uso en Docker

Estos scripts son utilizados automáticamente por Docker Compose:

```yaml
# docker-compose.yml
services:
  pim-service:
    depends_on:
      - postgres
    command: ["./scripts/wait-for-db.sh", "./main"]
```

## 🔧 Desarrollo

### Agregar Nuevos Scripts

1. Crear el script en este directorio
2. Hacer el archivo ejecutable: `chmod +x script-name.sh`
3. Documentar su uso en este README
4. Si es un test de integración, moverlo a `test-integration/`

### Convenciones

- Usar extensión `.sh` para scripts bash
- Incluir shebang: `#!/bin/bash`
- Agregar comentarios descriptivos
- Manejar errores apropiadamente
- Usar códigos de salida estándar (0 = éxito, 1+ = error)

---

**Para tests de integración, ver**: [`test-integration/README.md`](../test-integration/README.md) 