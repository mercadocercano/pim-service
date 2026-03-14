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

// BrandStatsPostgresRepository implementa el repositorio usando PostgreSQL
type BrandStatsPostgresRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewBrandStatsPostgresRepository crea una nueva instancia del repositorio
func NewBrandStatsPostgresRepository(db *sql.DB) *BrandStatsPostgresRepository {
	return &BrandStatsPostgresRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo BrandStats
func (r *BrandStatsPostgresRepository) Create(ctx context.Context, BrandStats *entity.BrandStats) error {
	query := `
		INSERT INTO BrandStatss (
			id, tenant_id, name, active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := r.db.ExecContext(ctx, query,
		BrandStats.ID,
		BrandStats.TenantID,
		BrandStats.Name,
		BrandStats.Active,
		BrandStats.CreatedAt,
		BrandStats.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error creando BrandStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrBrandStatsCreateFailed, err)
	}

	return nil
}

// Update actualiza un BrandStats existente
func (r *BrandStatsPostgresRepository) Update(ctx context.Context, BrandStats *entity.BrandStats) error {
	query := `
		UPDATE BrandStatss SET
			name = $3,
			active = $4,
			updated_at = $5
		WHERE id = $1 AND tenant_id = $2
	`

	result, err := r.db.ExecContext(ctx, query,
		BrandStats.ID,
		BrandStats.TenantID,
		BrandStats.Name,
		BrandStats.Active,
		BrandStats.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error actualizando BrandStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrBrandStatsUpdateFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrBrandStatsNotFound
	}

	return nil
}

// FindByID busca un BrandStats por su ID
func (r *BrandStatsPostgresRepository) FindByID(ctx context.Context, id string, tenantID string) (*entity.BrandStats, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM BrandStatss 
		WHERE id = $1 AND tenant_id = $2
	`

	row := r.db.QueryRowContext(ctx, query, id, tenantID)
	return r.scanBrandStats(row)
}

// FindByTenant obtiene todos los BrandStatss de un tenant
func (r *BrandStatsPostgresRepository) FindByTenant(ctx context.Context, tenantID string) ([]*entity.BrandStats, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM BrandStatss 
		WHERE tenant_id = $1 AND active = true
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanBrandStatss(rows)
}

// Delete elimina un BrandStats
func (r *BrandStatsPostgresRepository) Delete(ctx context.Context, id string, tenantID string) error {
	query := `DELETE FROM BrandStatss WHERE id = $1 AND tenant_id = $2`

	result, err := r.db.ExecContext(ctx, query, id, tenantID)
	if err != nil {
		log.Printf("Error eliminando BrandStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrBrandStatsDeleteFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrBrandStatsNotFound
	}

	return nil
}

// SearchByCriteria busca BrandStatss usando criterios
func (r *BrandStatsPostgresRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.BrandStats, error) {
	baseQuery := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM BrandStatss
	`

	query, params := r.converter.ToSelectSQL(baseQuery, crit)

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanBrandStatss(rows)
}

// CountByCriteria cuenta BrandStatss usando criterios
func (r *BrandStatsPostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM BrandStatss"

	query, params := r.converter.ToCountSQL(baseCountQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanBrandStats escanea una fila y devuelve un BrandStats
func (r *BrandStatsPostgresRepository) scanBrandStats(row *sql.Row) (*entity.BrandStats, error) {
	var BrandStats entity.BrandStats

	err := row.Scan(
		&BrandStats.ID,
		&BrandStats.TenantID,
		&BrandStats.Name,
		&BrandStats.Active,
		&BrandStats.CreatedAt,
		&BrandStats.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &BrandStats, nil
}

// scanBrandStatss escanea múltiples filas y devuelve una lista de BrandStatss
func (r *BrandStatsPostgresRepository) scanBrandStatss(rows *sql.Rows) ([]*entity.BrandStats, error) {
	var BrandStatss []*entity.BrandStats

	for rows.Next() {
		var BrandStats entity.BrandStats

		err := rows.Scan(
			&BrandStats.ID,
			&BrandStats.TenantID,
			&BrandStats.Name,
			&BrandStats.Active,
			&BrandStats.CreatedAt,
			&BrandStats.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		BrandStatss = append(BrandStatss, &BrandStats)
	}

	return BrandStatss, nil
}
