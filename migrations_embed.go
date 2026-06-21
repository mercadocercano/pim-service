package main

import "embed"

// MigrationsFS embeds all migration files for pim-service.
// The "migrations" subdirectory name is required by the go-shared migrate helper
// (iofs.New expects the files under a named subdirectory of the provided FS).
//
// El entrypoint de pim-service es el main.go de la raíz (go build), que comparte
// paquete (main) con este archivo, por lo que main.go referencia MigrationsFS
// directamente sin import.
//
//go:embed migrations/*.up.sql migrations/*.down.sql
var MigrationsFS embed.FS
