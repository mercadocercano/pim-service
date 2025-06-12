# 🛠️ Scripts de Utilidades - PIM Service

Este directorio contiene scripts de utilidades para el servicio PIM.

## 📁 Archivos

- `migrate.sh` - Script de migración de base de datos
- `wait-for-db.sh` - Script para esperar que la base de datos esté lista
- `README.md` - Este archivo

## 📦 Tests de Integración

**⚠️ IMPORTANTE**: Los tests de integración se han movido al directorio `test-integration/`

Para ejecutar tests de integración, usa:

```bash
# Desde el directorio raíz del proyecto
./run_tests.sh

# O desde el directorio de tests
cd test-integration/
./run_integration_tests.sh
```

Ver documentación completa en: [`test-integration/README.md`](../test-integration/README.md)

## 🚀 Scripts Disponibles

### migrate.sh
Script para ejecutar migraciones de base de datos.

```bash
./scripts/migrate.sh
```

### wait-for-db.sh
Script para esperar que la base de datos esté lista antes de iniciar el servicio.

```bash
./scripts/wait-for-db.sh
```

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