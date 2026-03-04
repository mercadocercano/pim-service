package database

import (
	"context"
	"database/sql"
)

// Executor define la interfaz mínima para ejecutar queries SQL.
// Tanto *sql.DB como *sql.Tx implementan esta interfaz, permitiendo
// usar el mismo código con o sin transacción.
type Executor interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}
