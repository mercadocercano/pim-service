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

// PostgresNormalizeCategoryRepository implementa NormalizeCategoryRepository (ADR-007 §4).
// Opera sobre categorías crudas distintas, no por fila: el resolver es función pura del
// string, así que un UPDATE por categoría cubre todas sus filas.
type PostgresNormalizeCategoryRepository struct {
	db *sql.DB
}

// NewPostgresNormalizeCategoryRepository crea una instancia del repositorio.
func NewPostgresNormalizeCategoryRepository(db *sql.DB) domainport.NormalizeCategoryRepository {
	return &PostgresNormalizeCategoryRepository{db: db}
}

// CountDistinctCategories cuenta las categorías crudas distintas (incluye NULL como una).
func (r *PostgresNormalizeCategoryRepository) CountDistinctCategories(ctx context.Context, scope value_object.NormalizeCategoryScope) (int, error) {
	where, args := normalizeWhere(scope)
	q := fmt.Sprintf(
		`SELECT COUNT(*) FROM (SELECT category FROM global_products WHERE %s GROUP BY category) t`,
		where,
	)
	var count int
	if err := r.db.QueryRowContext(ctx, q, args...).Scan(&count); err != nil {
		return 0, fmt.Errorf("count distinct categories: %w", err)
	}
	return count, nil
}

// FetchDistinctCategories devuelve cada categoría cruda distinta con su conteo de productos y
// el category_slug actual (sólo si TODAS sus filas comparten el mismo; si no, nil para forzar
// recálculo). Ordena por conteo desc para que la muestra antes/después sea representativa.
func (r *PostgresNormalizeCategoryRepository) FetchDistinctCategories(ctx context.Context, scope value_object.NormalizeCategoryScope) ([]value_object.RawCategoryCount, error) {
	where, args := normalizeWhere(scope)
	args = append(args, scope.MaxRows)
	q := fmt.Sprintf(`
		SELECT category, COUNT(*) AS n,
		       CASE WHEN COUNT(DISTINCT category_slug) = 1 THEN MAX(category_slug) END AS current_slug
		FROM global_products
		WHERE %s
		GROUP BY category
		ORDER BY n DESC
		LIMIT $%d`, where, len(args))

	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("fetch distinct categories: %w", err)
	}
	defer rows.Close()

	var out []value_object.RawCategoryCount
	for rows.Next() {
		var cat, curSlug sql.NullString
		var n int
		if err := rows.Scan(&cat, &n, &curSlug); err != nil {
			return nil, fmt.Errorf("scan distinct category: %w", err)
		}
		rc := value_object.RawCategoryCount{ProductCount: n}
		if cat.Valid {
			s := cat.String
			rc.RawCategory = &s
		}
		if curSlug.Valid {
			s := curSlug.String
			rc.CurrentSlug = &s
		}
		out = append(out, rc)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate distinct categories: %w", err)
	}
	return out, nil
}

// LoadOverrides devuelve el mapa raw_category -> category_slug.
func (r *PostgresNormalizeCategoryRepository) LoadOverrides(ctx context.Context) (map[string]string, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT raw_category, category_slug FROM category_slug_overrides`)
	if err != nil {
		return nil, fmt.Errorf("load overrides: %w", err)
	}
	defer rows.Close()

	m := make(map[string]string)
	for rows.Next() {
		var raw, slug string
		if err := rows.Scan(&raw, &slug); err != nil {
			return nil, fmt.Errorf("scan override: %w", err)
		}
		m[raw] = slug
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate overrides: %w", err)
	}
	return m, nil
}

// ApplyInTransaction crea el snapshot de la tabla completa y aplica un UPDATE por mapping
// (sólo los Changed) en una transacción única. Si el snapshot falla, la tx aborta (ADR-005 §3).
// No hay colisiones posibles (category_slug no es UNIQUE), así que no se usan savepoints.
func (r *PostgresNormalizeCategoryRepository) ApplyInTransaction(
	ctx context.Context,
	snapshotName string,
	scope value_object.NormalizeCategoryScope,
	mappings []value_object.CategoryMapping,
) (snapshotRef string, affected int, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", 0, fmt.Errorf("begin tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	tableName := fmt.Sprintf("global_products_bkp_%s", snapshotName)

	// Paso 1: snapshot de la tabla completa DENTRO de la tx (si falla → abort, nada se aplica).
	snapQ := fmt.Sprintf(`CREATE TABLE %s AS SELECT * FROM global_products`, pq.QuoteIdentifier(tableName))
	if _, snapErr := tx.ExecContext(ctx, snapQ); snapErr != nil {
		return "", 0, fmt.Errorf("create snapshot %q: %w", tableName, snapErr)
	}

	// Paso 2: un UPDATE por categoría (replica el WHERE del fetch para conteos consistentes).
	_, baseArgs := normalizeWhere(scope) // baseArgs: source prefix si existe; WHERE vía shiftNormalizeWhere
	for _, m := range mappings {
		if !m.Changed {
			continue
		}
		// Argumentos: [newSlug] + baseArgs(source prefix) + [rawCategory?]
		args := make([]interface{}, 0, len(baseArgs)+2)
		args = append(args, m.NewSlug)
		args = append(args, baseArgs...)

		var catCond string
		if m.RawCategory == nil {
			catCond = "category IS NULL"
		} else {
			args = append(args, *m.RawCategory)
			catCond = fmt.Sprintf("category = $%d", len(args))
		}

		// Placeholders: $1 = newSlug (fijo); el WHERE base (is_active + source prefix opcional)
		// se desplaza +1 para no pisar $1; category y updated_at toman los siguientes índices.
		// `IS DISTINCT FROM $1` garantiza idempotencia: una 2da corrida no toca filas ya correctas.
		whereClause := shiftNormalizeWhere(scope, 1)
		q := fmt.Sprintf(
			`UPDATE global_products SET category_slug = $1, updated_at = $%d
			 WHERE %s AND %s AND category_slug IS DISTINCT FROM $1`,
			len(args)+1, whereClause, catCond,
		)
		args = append(args, time.Now())

		res, updErr := tx.ExecContext(ctx, q, args...)
		if updErr != nil {
			return "", 0, fmt.Errorf("update category %v: %w", safeStr(m.RawCategory), updErr)
		}
		n, _ := res.RowsAffected()
		affected += int(n)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return "", 0, fmt.Errorf("commit tx: %w", commitErr)
	}
	committed = true
	return tableName, affected, nil
}

// SaveAudit persiste el registro en global_product_reclassification_audit (tabla reusada, ADR-007 §4).
func (r *PostgresNormalizeCategoryRepository) SaveAudit(ctx context.Context, audit value_object.NormalizeAuditRow) error {
	summaryJSON, err := json.Marshal(audit.Summary)
	if err != nil {
		return fmt.Errorf("marshal summary: %w", err)
	}
	scopeJSON, err := json.Marshal(map[string]interface{}{
		"operation":     "normalize_category_slugs",
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
		audit.OperatorID, audit.ExecutedAt, audit.Mode, scopeJSON,
		audit.SnapshotRef, summaryJSON, audit.AffectedCount,
	)
	if err != nil {
		return fmt.Errorf("save audit: %w", err)
	}
	return nil
}

// normalizeWhere construye el WHERE base (is_active + source prefix opcional) con placeholders
// empezando en $1.
func normalizeWhere(scope value_object.NormalizeCategoryScope) (string, []interface{}) {
	conditions := []string{"is_active = true"}
	var args []interface{}
	if scope.SourcePrefix != "" {
		args = append(args, scope.SourcePrefix+"%")
		conditions = append(conditions, fmt.Sprintf("source LIKE $%d", len(args)))
	}
	return strings.Join(conditions, " AND "), args
}

// shiftNormalizeWhere reconstruye el WHERE base con los placeholders desplazados por `offset`
// (usado en el UPDATE donde $1 está ocupado por category_slug).
func shiftNormalizeWhere(scope value_object.NormalizeCategoryScope, offset int) string {
	conditions := []string{"is_active = true"}
	if scope.SourcePrefix != "" {
		conditions = append(conditions, fmt.Sprintf("source LIKE $%d", offset+1))
	}
	return strings.Join(conditions, " AND ")
}

func safeStr(s *string) string {
	if s == nil {
		return "<NULL>"
	}
	return *s
}
