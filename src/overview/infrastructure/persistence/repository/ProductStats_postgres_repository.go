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

// ProductStatsPostgresRepository implementa el repositorio usando PostgreSQL
type ProductStatsPostgresRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewProductStatsPostgresRepository crea una nueva instancia del repositorio
func NewProductStatsPostgresRepository(db *sql.DB) *ProductStatsPostgresRepository {
	return &ProductStatsPostgresRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo ProductStats
func (r *ProductStatsPostgresRepository) Create(ctx context.Context, ProductStats *entity.ProductStats) error {
	query := `
		INSERT INTO ProductStatss (
			id, tenant_id, name, active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`

	_, err := r.db.ExecContext(ctx, query,
		ProductStats.ID,
		ProductStats.TenantID,
		ProductStats.Name,
		ProductStats.Active,
		ProductStats.CreatedAt,
		ProductStats.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error creando ProductStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrProductStatsCreateFailed, err)
	}

	return nil
}

// Update actualiza un ProductStats existente
func (r *ProductStatsPostgresRepository) Update(ctx context.Context, ProductStats *entity.ProductStats) error {
	query := `
		UPDATE ProductStatss SET
			name = $3,
			active = $4,
			updated_at = $5
		WHERE id = $1 AND tenant_id = $2
	`

	result, err := r.db.ExecContext(ctx, query,
		ProductStats.ID,
		ProductStats.TenantID,
		ProductStats.Name,
		ProductStats.Active,
		ProductStats.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error actualizando ProductStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrProductStatsUpdateFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrProductStatsNotFound
	}

	return nil
}

// FindByID busca un ProductStats por su ID
func (r *ProductStatsPostgresRepository) FindByID(ctx context.Context, id string, tenantID string) (*entity.ProductStats, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM ProductStatss 
		WHERE id = $1 AND tenant_id = $2
	`

	row := r.db.QueryRowContext(ctx, query, id, tenantID)
	return r.scanProductStats(row)
}

// FindByTenant obtiene todos los ProductStatss de un tenant
func (r *ProductStatsPostgresRepository) FindByTenant(ctx context.Context, tenantID string) ([]*entity.ProductStats, error) {
	query := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM ProductStatss 
		WHERE tenant_id = $1 AND active = true
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanProductStatss(rows)
}

// Delete elimina un ProductStats
func (r *ProductStatsPostgresRepository) Delete(ctx context.Context, id string, tenantID string) error {
	query := `DELETE FROM ProductStatss WHERE id = $1 AND tenant_id = $2`

	result, err := r.db.ExecContext(ctx, query, id, tenantID)
	if err != nil {
		log.Printf("Error eliminando ProductStats: %v", err)
		return fmt.Errorf("%w: %v", exception.ErrProductStatsDeleteFailed, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrProductStatsNotFound
	}

	return nil
}

// SearchByCriteria busca ProductStatss usando criterios
func (r *ProductStatsPostgresRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.ProductStats, error) {
	baseQuery := `
		SELECT id, tenant_id, name, active, created_at, updated_at
		FROM ProductStatss
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

	return r.scanProductStatss(rows)
}

// CountByCriteria cuenta ProductStatss usando criterios
func (r *ProductStatsPostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM ProductStatss"

	query, params, err := r.converter.ToCountSQL(baseCountQuery, crit)
	if err != nil {
		return 0, fmt.Errorf("invalid criteria: %w", err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanProductStats escanea una fila y devuelve un ProductStats
func (r *ProductStatsPostgresRepository) scanProductStats(row *sql.Row) (*entity.ProductStats, error) {
	var ProductStats entity.ProductStats

	err := row.Scan(
		&ProductStats.ID,
		&ProductStats.TenantID,
		&ProductStats.Name,
		&ProductStats.Active,
		&ProductStats.CreatedAt,
		&ProductStats.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &ProductStats, nil
}

// scanProductStatss escanea múltiples filas y devuelve una lista de ProductStatss
func (r *ProductStatsPostgresRepository) scanProductStatss(rows *sql.Rows) ([]*entity.ProductStats, error) {
	var ProductStatss []*entity.ProductStats

	for rows.Next() {
		var ProductStats entity.ProductStats

		err := rows.Scan(
			&ProductStats.ID,
			&ProductStats.TenantID,
			&ProductStats.Name,
			&ProductStats.Active,
			&ProductStats.CreatedAt,
			&ProductStats.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		ProductStatss = append(ProductStatss, &ProductStats)
	}

	return ProductStatss, nil
}
