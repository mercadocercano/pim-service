#!/bin/bash
# ⚠️ OBSOLETO / RETIRADO (ADR-001).
#
# Las migraciones de pim-service ahora corren IN-APP vía sharedmigrate.RunMigrations
# (golang-migrate) en main.go, con tabla de control schema_migrations(version, dirty)
# y baseline en prod. El antiguo Job pim-migrate y la stage 'migrate' del Dockerfile
# fueron eliminados.
#
# Este script ya NO aplica migraciones (re-corría todos los .sql sin tracking, y tras
# el rename a .up/.down podía ejecutar DROPs por accidente). Se deja solo como marcador.

echo "migrate.sh está RETIRADO. Las migraciones corren in-app (ADR-001, golang-migrate)."
echo "Para una migración manual de emergencia, usar el binario del servicio o psql directo."
exit 1
