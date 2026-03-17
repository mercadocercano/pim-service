package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"saas-mt-pim-service/src/overview/domain/entity"
	"saas-mt-pim-service/src/overview/domain/exception"
	cr "github.com/mercadocercano/criteria"
)

// AttributeStatsPostgresRepository implementa el repositorio usando PostgreSQL
type AttributeStatsPostgresRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewAttributeStatsPostgresRepository crea una nueva instancia del repositorio
func NewAttributeStatsPostgresRepository(db *sql.DB) *AttributeStatsPostgresRepository {
	return &AttributeStatsPostgresRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo AttributeStats
func (r *AttributeStatsPostgresRepository) Create(ctx context.Context, AttributeStats *entity.AttributeStats) error {
	query := `
		INSERT INTO AttributeStatss (
			id, tenant_id, name, active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := r.db.ExecContext(ctx, query,
		AttributeStats.ID,
		AttributeStats.TenantID,
		AttributeStats.Name,
		AttributeStats.Active,
		AttributeStats.CreatedAt,
		AttributeStats.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error creando AttributeStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrAttributeStatsCreateFailed, err)
	}

	return nil
}

// Update actualiza un AttributeStats existente
func (r *AttributeStatsPostgresRepository) Update(ctx context.Context, AttributeStats *entity.AttributeStats) error {
	query := `
		UPDATE AttributeStatss SET
			name = $3,
			active = $4,
			updated_at = $5
		WHERE id = $1 AND tenant_id = $2
	`

	result, err := r.db.ExecContext(ctx, query,
		AttributeStats.ID,
		AttributeStats.TenantID,
		AttributeStats.Name,
		AttributeStats.Active,
		AttributeStats.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error actualizando AttributeStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrAttributeStatsUpdateFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrAttributeStatsNotFound
	}

	return nil
}

// FindByID busca un AttributeStats por su ID
func (r *AttributeStatsPostgresRepository) FindByID(ctx context.Context, id string, tenantID string) (*entity.AttributeStats, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM AttributeStatss 
		WHERE id = $1 AND tenant_id = $2
	`

	row := r.db.QueryRowContext(ctx, query, id, tenantID)
	return r.scanAttributeStats(row)
}

// FindByTenant obtiene todos los AttributeStatss de un tenant
func (r *AttributeStatsPostgresRepository) FindByTenant(ctx context.Context, tenantID string) ([]*entity.AttributeStats, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM AttributeStatss 
		WHERE tenant_id = $1 AND active = true
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributeStatss(rows)
}

// Delete elimina un AttributeStats
func (r *AttributeStatsPostgresRepository) Delete(ctx context.Context, id string, tenantID string) error {
	query := `DELETE FROM AttributeStatss WHERE id = $1 AND tenant_id = $2`

	result, err := r.db.ExecContext(ctx, query, id, tenantID)
	if err != nil {
		log.Printf("Error eliminando AttributeStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrAttributeStatsDeleteFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrAttributeStatsNotFound
	}

	return nil
}

// SearchByCriteria busca AttributeStatss usando criterios
func (r *AttributeStatsPostgresRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.AttributeStats, error) {
	baseQuery := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM AttributeStatss
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

	return r.scanAttributeStatss(rows)
}

// CountByCriteria cuenta AttributeStatss usando criterios
func (r *AttributeStatsPostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM AttributeStatss"

	query, params, err := r.converter.ToCountSQL(baseCountQuery, crit)
	if err != nil {
		return 0, fmt.Errorf("invalid criteria: %w", err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanAttributeStats escanea una fila y devuelve un AttributeStats
func (r *AttributeStatsPostgresRepository) scanAttributeStats(row *sql.Row) (*entity.AttributeStats, error) {
	var AttributeStats entity.AttributeStats

	err := row.Scan(
		&AttributeStats.ID,
		&AttributeStats.TenantID,
		&AttributeStats.Name,
		&AttributeStats.Active,
		&AttributeStats.CreatedAt,
		&AttributeStats.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &AttributeStats, nil
}

// scanAttributeStatss escanea múltiples filas y devuelve una lista de AttributeStatss
func (r *AttributeStatsPostgresRepository) scanAttributeStatss(rows *sql.Rows) ([]*entity.AttributeStats, error) {
	var AttributeStatss []*entity.AttributeStats

	for rows.Next() {
		var AttributeStats entity.AttributeStats

		err := rows.Scan(
			&AttributeStats.ID,
			&AttributeStats.TenantID,
			&AttributeStats.Name,
			&AttributeStats.Active,
			&AttributeStats.CreatedAt,
			&AttributeStats.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		AttributeStatss = append(AttributeStatss, &AttributeStats)
	}

	return AttributeStatss, nil
}
