package port

import (
	"context"
	"database/sql"
)

// Transaction abstrae las operaciones de commit/rollback.
// *sql.Tx implementa esta interfaz nativamente.
type Transaction interface {
	Commit() error
	Rollback() error
}

// TxBeginner abstrae la capacidad de iniciar transacciones.
type TxBeginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Transaction, error)
}

// SQLDBTxBeginner adapta *sql.DB para implementar TxBeginner.
type SQLDBTxBeginner struct {
	DB *sql.DB
}

// BeginTx inicia una transacción. *sql.Tx implementa Transaction.
func (s *SQLDBTxBeginner) BeginTx(ctx context.Context, opts *sql.TxOptions) (Transaction, error) {
	return s.DB.BeginTx(ctx, opts)
}
