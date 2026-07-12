package persistence

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
	domainport "saas-mt-pim-service/src/product/global_catalog/domain/port"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// PostgresBulkVerifyRepository implementa domainport.BulkVerifyRepository usando PostgreSQL.
type PostgresBulkVerifyRepository struct {
	db *sql.DB
}

// NewPostgresBulkVerifyRepository crea una instancia del repositorio.
func NewPostgresBulkVerifyRepository(db *sql.DB) domainport.BulkVerifyRepository {
	return &PostgresBulkVerifyRepository{db: db}
}

// ApplyInTransaction crea el snapshot y aplica los UPDATEs de is_verified en una tx única.
// Si el snapshot falla, la tx aborta y nada se aplica.
func (r *PostgresBulkVerifyRepository) ApplyInTransaction(
	ctx context.Context,
	snapshotName string,
	ids []string,
	mode value_object.BulkVerifyMode,
) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("begin tx: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	tableName := fmt.Sprintf("global_products_bkp_bulkverify_%s", snapshotName)

	// Paso 1: snapshot dentro de la tx
	if snapErr := r.createSnapshot(ctx, tx, tableName, ids); snapErr != nil {
		return 0, fmt.Errorf("create snapshot: %w", snapErr)
	}

	// Paso 2: aplicar updates dentro de la misma tx
	isVerified := mode == value_object.BulkVerifyModeVerify
	result, updErr := tx.ExecContext(ctx,
		`UPDATE global_products SET is_verified = $1, updated_at = $2 WHERE id::text = ANY($3)`,
		isVerified, time.Now(), pq.Array(ids),
	)
	if updErr != nil {
		return 0, fmt.Errorf("apply updates: %w", updErr)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("rows affected: %w", err)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return 0, fmt.Errorf("commit tx: %w", commitErr)
	}
	committed = true

	return int(rowsAffected), nil
}

// createSnapshot crea la tabla backup dentro de tx.
func (r *PostgresBulkVerifyRepository) createSnapshot(ctx context.Context, tx *sql.Tx, tableName string, ids []string) error {
	q := fmt.Sprintf(
		`CREATE TABLE %s AS SELECT * FROM global_products WHERE id::text = ANY($1)`,
		pq.QuoteIdentifier(tableName),
	)
	_, err := tx.ExecContext(ctx, q, pq.Array(ids))
	if err != nil {
		return fmt.Errorf("create snapshot %q: %w", tableName, err)
	}
	return nil
}

// SaveAudit persiste el registro de auditoría en global_product_bulk_verify_audit.
func (r *PostgresBulkVerifyRepository) SaveAudit(ctx context.Context, audit value_object.BulkVerifyAuditRow) error {
	summaryJSON, err := json.Marshal(audit.Summary)
	if err != nil {
		return fmt.Errorf("marshal summary: %w", err)
	}

	requestIDsJSON, err := json.Marshal(audit.RequestIDs)
	if err != nil {
		return fmt.Errorf("marshal request ids: %w", err)
	}

	_, err = r.db.ExecContext(ctx,
		`INSERT INTO global_product_bulk_verify_audit
		 (operator_id, executed_at, mode, request_ids, snapshot_ref, summary, affected_count)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		audit.OperatorID,
		audit.ExecutedAt,
		string(audit.Mode),
		requestIDsJSON,
		audit.SnapshotRef,
		summaryJSON,
		audit.AffectedCount,
	)
	if err != nil {
		return fmt.Errorf("save audit: %w", err)
	}
	return nil
}
