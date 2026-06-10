package repository

import (
	"context"
	"database/sql"
	"fmt"
	cr "github.com/mercadocercano/criteria"
	"log"
	"saas-mt-pim-service/src/overview/domain/entity"
	"saas-mt-pim-service/src/overview/domain/exception"
)

// MarketplaceSummaryPostgresRepository implementa el repositorio usando PostgreSQL
type MarketplaceSummaryPostgresRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewMarketplaceSummaryPostgresRepository crea una nueva instancia del repositorio
func NewMarketplaceSummaryPostgresRepository(db *sql.DB) *MarketplaceSummaryPostgresRepository {
	return &MarketplaceSummaryPostgresRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo MarketplaceSummary
func (r *MarketplaceSummaryPostgresRepository) Create(ctx context.Context, MarketplaceSummary *entity.MarketplaceSummary) error {
	query := `
		INSERT INTO MarketplaceSummarys (
			id, tenant_id, name, active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := r.db.ExecContext(ctx, query,
		MarketplaceSummary.ID,
		MarketplaceSummary.TenantID,
		MarketplaceSummary.Name,
		MarketplaceSummary.Active,
		MarketplaceSummary.CreatedAt,
		MarketplaceSummary.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error creando MarketplaceSummary: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrMarketplaceSummaryCreateFailed, err)
	}

	return nil
}

// Update actualiza un MarketplaceSummary existente
func (r *MarketplaceSummaryPostgresRepository) Update(ctx context.Context, MarketplaceSummary *entity.MarketplaceSummary) error {
	query := `
		UPDATE MarketplaceSummarys SET
			name = $3,
			active = $4,
			updated_at = $5
		WHERE id = $1 AND tenant_id = $2
	`

	result, err := r.db.ExecContext(ctx, query,
		MarketplaceSummary.ID,
		MarketplaceSummary.TenantID,
		MarketplaceSummary.Name,
		MarketplaceSummary.Active,
		MarketplaceSummary.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error actualizando MarketplaceSummary: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrMarketplaceSummaryUpdateFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrMarketplaceSummaryNotFound
	}

	return nil
}

// FindByID busca un MarketplaceSummary por su ID
func (r *MarketplaceSummaryPostgresRepository) FindByID(ctx context.Context, id string, tenantID string) (*entity.MarketplaceSummary, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM MarketplaceSummarys 
		WHERE id = $1 AND tenant_id = $2
	`

	row := r.db.QueryRowContext(ctx, query, id, tenantID)
	return r.scanMarketplaceSummary(row)
}

// FindByTenant obtiene todos los MarketplaceSummarys de un tenant
func (r *MarketplaceSummaryPostgresRepository) FindByTenant(ctx context.Context, tenantID string) ([]*entity.MarketplaceSummary, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM MarketplaceSummarys 
		WHERE tenant_id = $1 AND active = true
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMarketplaceSummarys(rows)
}

// Delete elimina un MarketplaceSummary
func (r *MarketplaceSummaryPostgresRepository) Delete(ctx context.Context, id string, tenantID string) error {
	query := `DELETE FROM MarketplaceSummarys WHERE id = $1 AND tenant_id = $2`

	result, err := r.db.ExecContext(ctx, query, id, tenantID)
	if err != nil {
		log.Printf("Error eliminando MarketplaceSummary: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrMarketplaceSummaryDeleteFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrMarketplaceSummaryNotFound
	}

	return nil
}

// SearchByCriteria busca MarketplaceSummarys usando criterios
func (r *MarketplaceSummaryPostgresRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.MarketplaceSummary, error) {
	baseQuery := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM MarketplaceSummarys
	`

	query, params, err := r.converter.ToSelectSQL(baseQuery, crit)
	if err != nil {
		return nil, fmt.Errorf("invalid criteria: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMarketplaceSummarys(rows)
}

// CountByCriteria cuenta MarketplaceSummarys usando criterios
func (r *MarketplaceSummaryPostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM MarketplaceSummarys"

	query, params, err := r.converter.ToCountSQL(baseCountQuery, crit)
	if err != nil {
		return 0, fmt.Errorf("invalid criteria: %w", err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanMarketplaceSummary escanea una fila y devuelve un MarketplaceSummary
func (r *MarketplaceSummaryPostgresRepository) scanMarketplaceSummary(row *sql.Row) (*entity.MarketplaceSummary, error) {
	var MarketplaceSummary entity.MarketplaceSummary

	err := row.Scan(
		&MarketplaceSummary.ID,
		&MarketplaceSummary.TenantID,
		&MarketplaceSummary.Name,
		&MarketplaceSummary.Active,
		&MarketplaceSummary.CreatedAt,
		&MarketplaceSummary.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &MarketplaceSummary, nil
}

// scanMarketplaceSummarys escanea múltiples filas y devuelve una lista de MarketplaceSummarys
func (r *MarketplaceSummaryPostgresRepository) scanMarketplaceSummarys(rows *sql.Rows) ([]*entity.MarketplaceSummary, error) {
	var MarketplaceSummarys []*entity.MarketplaceSummary

	for rows.Next() {
		var MarketplaceSummary entity.MarketplaceSummary

		err := rows.Scan(
			&MarketplaceSummary.ID,
			&MarketplaceSummary.TenantID,
			&MarketplaceSummary.Name,
			&MarketplaceSummary.Active,
			&MarketplaceSummary.CreatedAt,
			&MarketplaceSummary.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		MarketplaceSummarys = append(MarketplaceSummarys, &MarketplaceSummary)
	}

	return MarketplaceSummarys, nil
}
