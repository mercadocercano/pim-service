package database

import (
	"database/sql"

	sharedport "github.com/hornosg/go-shared/domain/port"
)

// Compile-time checks: both *sql.DB and *sql.Tx satisfy the shared Executor port.
var (
	_ sharedport.Executor = (*sql.DB)(nil)
	_ sharedport.Executor = (*sql.Tx)(nil)
)
