package persistence

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
	domainport "saas-mt-pim-service/src/product/global_catalog/domain/port"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// pgUniqueViolation es el código SQLSTATE de PostgreSQL para unicidad violada.
const pgUniqueViolation = "23505"

// PostgresReclassifyRepository implementa domainport.ReclassifyRepository usando PostgreSQL.
type PostgresReclassifyRepository struct {
	db *sql.DB
}

// NewPostgresReclassifyRepository crea una instancia del repositorio.
func NewPostgresReclassifyRepository(db *sql.DB) domainport.ReclassifyRepository {
	return &PostgresReclassifyRepository{db: db}
}

// CountCandidates devuelve el COUNT de productos candidatos según el scope.
// Candidatos: business_type IS NULL OR business_type = 'almacen'.
// Si SourcePrefix no está vacío, filtra adicionalmente source LIKE '<prefix>%'.
func (r *PostgresReclassifyRepository) CountCandidates(ctx context.Context, scope value_object.ReclassifyScope) (int, error) {
	q, args := buildCandidateQuery("COUNT(*)", scope, false)
	var count int
	err := r.db.QueryRowContext(ctx, q, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("count candidates: %w", err)
	}
	return count, nil
}

// FetchCandidates devuelve los productos candidatos con sus campos relevantes.
func (r *PostgresReclassifyRepository) FetchCandidates(ctx context.Context, scope value_object.ReclassifyScope) ([]value_object.ReclassifyCandidate, error) {
	q, args := buildCandidateQuery("id, name, COALESCE(category, ''), business_type", scope, true)
	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("fetch candidates: %w", err)
	}
	defer rows.Close()

	var candidates []value_object.ReclassifyCandidate
	for rows.Next() {
		var c value_object.ReclassifyCandidate
		var bt sql.NullString
		if err := rows.Scan(&c.ID, &c.Name, &c.Category, &bt); err != nil {
			return nil, fmt.Errorf("scan candidate: %w", err)
		}
		if bt.Valid && bt.String != "" {
			s := bt.String
			c.BusinessType = &s
		}
		candidates = append(candidates, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate candidates: %w", err)
	}
	return candidates, nil
}

// ApplyInTransaction crea el snapshot y aplica los updates en una transacción única.
// Si el snapshot falla, la tx aborta y nada se aplica (invariante ADR-005 §3, TEST-ID T-024).
// Colisión 23505 → skip individual, nunca rollback del lote (ADR-005 §8, TEST-ID T-022).
//
// TEST-IDs: T-026 (snapshot creado), T-027 (updates aplicados), T-028 (colisión → skip),
//           T-030 (snapshot falla → rollback).
func (r *PostgresReclassifyRepository) ApplyInTransaction(
	ctx context.Context,
	snapshotName string,
	candidateIDs []string,
	updates []value_object.ReclassifyUpdate,
) (snapshotRef string, affected int, collisions int, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", 0, 0, fmt.Errorf("begin tx: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	tableName := fmt.Sprintf("global_products_bkp_%s", snapshotName)

	// Paso 1: crear snapshot DENTRO de la tx (invariante: si falla → abort, nada se aplica)
	if snapErr := r.createSnapshot(ctx, tx, tableName, candidateIDs); snapErr != nil {
		return "", 0, 0, fmt.Errorf("create snapshot: %w", snapErr)
	}

	// Paso 2: aplicar updates DENTRO de la misma tx
	rowsUpdated, rowsCollided, updErr := r.applyUpdates(ctx, tx, updates)
	if updErr != nil {
		return "", 0, 0, fmt.Errorf("apply updates: %w", updErr)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return "", 0, 0, fmt.Errorf("commit tx: %w", commitErr)
	}
	committed = true

	return tableName, rowsUpdated, rowsCollided, nil
}

// createSnapshot crea global_products_bkp_<tableName> dentro de tx.
// TEST-ID: T-026, T-030.
func (r *PostgresReclassifyRepository) createSnapshot(ctx context.Context, tx *sql.Tx, tableName string, candidateIDs []string) error {
	if len(candidateIDs) == 0 {
		q := fmt.Sprintf(
			`CREATE TABLE %s AS SELECT * FROM global_products WHERE false`,
			pq.QuoteIdentifier(tableName),
		)
		_, err := tx.ExecContext(ctx, q)
		if err != nil {
			return fmt.Errorf("create empty snapshot %q: %w", tableName, err)
		}
		return nil
	}

	q := fmt.Sprintf(
		`CREATE TABLE %s AS SELECT * FROM global_products WHERE id::text = ANY($1)`,
		pq.QuoteIdentifier(tableName),
	)
	_, err := tx.ExecContext(ctx, q, pq.Array(candidateIDs))
	if err != nil {
		return fmt.Errorf("create snapshot %q: %w", tableName, err)
	}
	return nil
}

// applyUpdates ejecuta los UPDATEs de business_type dentro de tx.
// Colisión 23505 → skip individual, nunca aborta el lote (ADR-005 §8, TEST-ID T-022, T-028).
//
// IMPORTANTE: en PostgreSQL, un error dentro de una transacción la deja en estado
// abortado ("current transaction is aborted") y todo comando posterior — incluido el
// COMMIT — falla con 25P02. Por eso un simple `continue` tras un 23505 NO basta: hay que
// envolver cada UPDATE en un SAVEPOINT y hacer ROLLBACK TO SAVEPOINT ante la colisión,
// para que el resto del lote pueda continuar y la tx siga siendo committeable.
func (r *PostgresReclassifyRepository) applyUpdates(ctx context.Context, tx *sql.Tx, updates []value_object.ReclassifyUpdate) (affected int, collisions int, err error) {
	if len(updates) == 0 {
		return 0, 0, nil
	}

	for _, u := range updates {
		if _, spErr := tx.ExecContext(ctx, "SAVEPOINT sp_reclassify"); spErr != nil {
			return affected, collisions, fmt.Errorf("savepoint: %w", spErr)
		}

		_, updErr := tx.ExecContext(ctx,
			`UPDATE global_products SET business_type = $1, updated_at = $2 WHERE id::text = $3`,
			u.ToType, time.Now(), u.ID,
		)
		if updErr != nil {
			if isUniqueViolation(updErr) {
				// Revertir solo este UPDATE; la tx sigue viva para el resto del lote.
				if _, rbErr := tx.ExecContext(ctx, "ROLLBACK TO SAVEPOINT sp_reclassify"); rbErr != nil {
					return affected, collisions, fmt.Errorf("rollback to savepoint: %w", rbErr)
				}
				collisions++
				continue
			}
			return affected, collisions, fmt.Errorf("update product %q: %w", u.ID, updErr)
		}

		if _, relErr := tx.ExecContext(ctx, "RELEASE SAVEPOINT sp_reclassify"); relErr != nil {
			return affected, collisions, fmt.Errorf("release savepoint: %w", relErr)
		}
		affected++
	}
	return affected, collisions, nil
}

// SaveAudit persiste el registro en global_product_reclassification_audit.
// TEST-ID: T-029.
func (r *PostgresReclassifyRepository) SaveAudit(ctx context.Context, audit value_object.ReclassifyAuditRow) error {
	summaryJSON, err := json.Marshal(audit.Summary)
	if err != nil {
		return fmt.Errorf("marshal summary: %w", err)
	}

	scopeJSON, err := json.Marshal(map[string]interface{}{
		"source_prefix": audit.Scope.SourcePrefix,
		"max_rows":      audit.Scope.MaxRows,
	})
	if err != nil {
		return fmt.Errorf("marshal scope: %w", err)
	}

	_, err = r.db.ExecContext(ctx,
		`INSERT INTO global_product_reclassification_audit
		 (operator_id, executed_at, mode, scope, snapshot_ref, summary, affected_count)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		audit.OperatorID,
		audit.ExecutedAt,
		audit.Mode,
		scopeJSON,
		audit.SnapshotRef,
		summaryJSON,
		audit.AffectedCount,
	)
	if err != nil {
		return fmt.Errorf("save audit: %w", err)
	}
	return nil
}

// buildCandidateQuery construye el SELECT/COUNT con los filtros del scope.
// selectClause puede ser "COUNT(*)" o los campos a seleccionar.
// withLimit=true agrega LIMIT con el MaxRows del scope.
func buildCandidateQuery(selectClause string, scope value_object.ReclassifyScope, withLimit bool) (string, []interface{}) {
	var conditions []string
	var args []interface{}

	// Filtro: candidatos son los que tienen business_type NULL o 'almacen'
	conditions = append(conditions, "(business_type IS NULL OR business_type = 'almacen')")

	// Filtro opcional: source prefix
	if scope.SourcePrefix != "" {
		args = append(args, scope.SourcePrefix+"%")
		conditions = append(conditions, fmt.Sprintf("source LIKE $%d", len(args)))
	}

	where := strings.Join(conditions, " AND ")
	q := fmt.Sprintf("SELECT %s FROM global_products WHERE %s", selectClause, where)

	if withLimit {
		args = append(args, scope.MaxRows)
		q += fmt.Sprintf(" LIMIT $%d", len(args))
	}

	return q, args
}

// isUniqueViolation reporta si el error es una violación de constraint UNIQUE (SQLSTATE 23505).
func isUniqueViolation(err error) bool {
	if err == nil {
		return false
	}
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == pgUniqueViolation
	}
	// Fallback por string (para casos donde lib/pq no wrappea el error)
	return strings.Contains(err.Error(), pgUniqueViolation)
}
