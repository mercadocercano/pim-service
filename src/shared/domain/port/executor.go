package port

import (
	"context"
	"database/sql"
)

// Executor abstrae una transacción o conexión de base de datos.
// Tanto *sql.DB como *sql.Tx satisfacen esta interfaz.
type Executor interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}
