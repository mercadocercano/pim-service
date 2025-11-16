package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"pim/src/overview/domain/entity"
	"pim/src/overview/domain/exception"
	"pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"
)

// CategoryStatsPostgresRepository implementa el repositorio usando PostgreSQL
type CategoryStatsPostgresRepository struct {
	db        *sql.DB
	converter *sharedCriteria.SQLCriteriaConverter
}

// NewCategoryStatsPostgresRepository crea una nueva instancia del repositorio
func NewCategoryStatsPostgresRepository(db *sql.DB) *CategoryStatsPostgresRepository {
	return &CategoryStatsPostgresRepository{
		db:        db,
		converter: sharedCriteria.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo CategoryStats
func (r *CategoryStatsPostgresRepository) Create(ctx context.Context, CategoryStats *entity.CategoryStats) error {
	query := `
		INSERT INTO CategoryStatss (
			id, tenant_id, name, active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := r.db.ExecContext(ctx, query,
		CategoryStats.ID,
		CategoryStats.TenantID,
		CategoryStats.Name,
		CategoryStats.Active,
		CategoryStats.CreatedAt,
		CategoryStats.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error creando CategoryStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrCategoryStatsCreateFailed, err)
	}

	return nil
}

// Update actualiza un CategoryStats existente
func (r *CategoryStatsPostgresRepository) Update(ctx context.Context, CategoryStats *entity.CategoryStats) error {
	query := `
		UPDATE CategoryStatss SET
			name = $3,
			active = $4,
			updated_at = $5
		WHERE id = $1 AND tenant_id = $2
	`

	result, err := r.db.ExecContext(ctx, query,
		CategoryStats.ID,
		CategoryStats.TenantID,
		CategoryStats.Name,
		CategoryStats.Active,
		CategoryStats.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error actualizando CategoryStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrCategoryStatsUpdateFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrCategoryStatsNotFound
	}

	return nil
}

// FindByID busca un CategoryStats por su ID
func (r *CategoryStatsPostgresRepository) FindByID(ctx context.Context, id string, tenantID string) (*entity.CategoryStats, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM CategoryStatss 
		WHERE id = $1 AND tenant_id = $2
	`

	row := r.db.QueryRowContext(ctx, query, id, tenantID)
	return r.scanCategoryStats(row)
}

// FindByTenant obtiene todos los CategoryStatss de un tenant
func (r *CategoryStatsPostgresRepository) FindByTenant(ctx context.Context, tenantID string) ([]*entity.CategoryStats, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM CategoryStatss 
		WHERE tenant_id = $1 AND active = true
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanCategoryStatss(rows)
}

// Delete elimina un CategoryStats
func (r *CategoryStatsPostgresRepository) Delete(ctx context.Context, id string, tenantID string) error {
	query := `DELETE FROM CategoryStatss WHERE id = $1 AND tenant_id = $2`

	result, err := r.db.ExecContext(ctx, query, id, tenantID)
	if err != nil {
		log.Printf("Error eliminando CategoryStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrCategoryStatsDeleteFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrCategoryStatsNotFound
	}

	return nil
}

// SearchByCriteria busca CategoryStatss usando criterios
func (r *CategoryStatsPostgresRepository) SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.CategoryStats, error) {
	baseQuery := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM CategoryStatss
	`

	query, params := r.converter.ToSelectSQL(baseQuery, crit)

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanCategoryStatss(rows)
}

// CountByCriteria cuenta CategoryStatss usando criterios
func (r *CategoryStatsPostgresRepository) CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM CategoryStatss"

	query, params := r.converter.ToCountSQL(baseCountQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanCategoryStats escanea una fila y devuelve un CategoryStats
func (r *CategoryStatsPostgresRepository) scanCategoryStats(row *sql.Row) (*entity.CategoryStats, error) {
	var CategoryStats entity.CategoryStats

	err := row.Scan(
		&CategoryStats.ID,
		&CategoryStats.TenantID,
		&CategoryStats.Name,
		&CategoryStats.Active,
		&CategoryStats.CreatedAt,
		&CategoryStats.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &CategoryStats, nil
}

// scanCategoryStatss escanea múltiples filas y devuelve una lista de CategoryStatss
func (r *CategoryStatsPostgresRepository) scanCategoryStatss(rows *sql.Rows) ([]*entity.CategoryStats, error) {
	var CategoryStatss []*entity.CategoryStats

	for rows.Next() {
		var CategoryStats entity.CategoryStats

		err := rows.Scan(
			&CategoryStats.ID,
			&CategoryStats.TenantID,
			&CategoryStats.Name,
			&CategoryStats.Active,
			&CategoryStats.CreatedAt,
			&CategoryStats.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		CategoryStatss = append(CategoryStatss, &CategoryStats)
	}

	return CategoryStatss, nil
}
