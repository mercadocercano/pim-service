package persistence

import (
	"context"
	"database/sql"
	"fmt"

	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/port"

	"github.com/google/uuid"
)

// TenantCategoryMappingPostgresRepository implementa el repositorio de mapeos de categorías tenant para PostgreSQL
type TenantCategoryMappingPostgresRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewTenantCategoryMappingPostgresRepository crea una nueva instancia del repositorio
func NewTenantCategoryMappingPostgresRepository(db *sql.DB) port.TenantCategoryMappingRepository {
	return &TenantCategoryMappingPostgresRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Save guarda un mapeo de categoría tenant
func (r *TenantCategoryMappingPostgresRepository) Save(ctx context.Context, mapping *entity.TenantCategoryMapping) error {
	// Generar ID si no existe
	if mapping.ID == "" {
		mapping.ID = uuid.New().String()
	}

	query := `
		INSERT INTO tenant_category_mappings (
			id, tenant_id, category_id, marketplace_category_id, custom_name, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET
			tenant_id = EXCLUDED.tenant_id,
			category_id = EXCLUDED.category_id,
			marketplace_category_id = EXCLUDED.marketplace_category_id,
			custom_name = EXCLUDED.custom_name,
			updated_at = EXCLUDED.updated_at
	`

	_, err := r.db.ExecContext(ctx, query,
		mapping.ID,
		mapping.TenantID,
		mapping.CategoryID,
		mapping.MarketplaceCategoryID,
		mapping.CustomName,
		mapping.CreatedAt,
		mapping.UpdatedAt,
	)

	return err
}

// GetByID obtiene un mapeo por su ID
func (r *TenantCategoryMappingPostgresRepository) GetByID(ctx context.Context, id string) (*entity.TenantCategoryMapping, error) {
	query := `
		SELECT id, tenant_id, category_id, marketplace_category_id, custom_name, created_at, updated_at
		FROM tenant_category_mappings
		WHERE id = $1 AND deleted_at IS NULL
	`

	row := r.db.QueryRowContext(ctx, query, id)
	return r.scanMapping(row)
}

// GetByTenantAndMarketplaceCategory obtiene un mapeo específico
func (r *TenantCategoryMappingPostgresRepository) GetByTenantAndMarketplaceCategory(ctx context.Context, tenantID, marketplaceCategoryID string) (*entity.TenantCategoryMapping, error) {
	query := `
		SELECT id, tenant_id, category_id, marketplace_category_id, custom_name, created_at, updated_at
		FROM tenant_category_mappings
		WHERE tenant_id = $1 AND marketplace_category_id = $2 AND deleted_at IS NULL
	`

	row := r.db.QueryRowContext(ctx, query, tenantID, marketplaceCategoryID)
	return r.scanMapping(row)
}

// GetByTenantID obtiene todos los mapeos de un tenant
func (r *TenantCategoryMappingPostgresRepository) GetByTenantID(ctx context.Context, tenantID string) ([]*entity.TenantCategoryMapping, error) {
	query := `
		SELECT id, tenant_id, category_id, marketplace_category_id, custom_name, created_at, updated_at
		FROM tenant_category_mappings
		WHERE tenant_id = $1 AND deleted_at IS NULL
		ORDER BY created_at
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMappings(rows)
}

// GetByMarketplaceCategoryID obtiene todos los mapeos de una categoría marketplace
func (r *TenantCategoryMappingPostgresRepository) GetByMarketplaceCategoryID(ctx context.Context, marketplaceCategoryID string) ([]*entity.TenantCategoryMapping, error) {
	query := `
		SELECT id, tenant_id, category_id, marketplace_category_id, custom_name, created_at, updated_at
		FROM tenant_category_mappings
		WHERE marketplace_category_id = $1 AND deleted_at IS NULL
		ORDER BY tenant_id, created_at
	`

	rows, err := r.db.QueryContext(ctx, query, marketplaceCategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMappings(rows)
}

// FindByCriteria busca mapeos según criterios
func (r *TenantCategoryMappingPostgresRepository) FindByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.TenantCategoryMapping, error) {
	baseQuery := `
		SELECT id, tenant_id, category_id, marketplace_category_id, custom_name, created_at, updated_at
		FROM tenant_category_mappings
		WHERE deleted_at IS NULL
	`

	query, args, err := r.converter.ToSelectSQL(baseQuery, crit)
	if err != nil {
		return nil, fmt.Errorf("invalid criteria: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMappings(rows)
}

// CountByCriteria cuenta mapeos según criterios
func (r *TenantCategoryMappingPostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := `
		SELECT COUNT(*)
		FROM tenant_category_mappings
		WHERE deleted_at IS NULL
	`

	query, args, err := r.converter.ToCountSQL(baseCountQuery, crit)
	if err != nil {
		return 0, fmt.Errorf("invalid criteria: %w", err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	return count, err
}

// Update actualiza un mapeo
func (r *TenantCategoryMappingPostgresRepository) Update(ctx context.Context, mapping *entity.TenantCategoryMapping) error {
	query := `
		UPDATE tenant_category_mappings
		SET tenant_id = $2, category_id = $3, marketplace_category_id = $4, custom_name = $5, updated_at = $6
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query,
		mapping.ID,
		mapping.TenantID,
		mapping.CategoryID,
		mapping.MarketplaceCategoryID,
		mapping.CustomName,
		mapping.UpdatedAt,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("mapping with id %s not found", mapping.ID)
	}

	return nil
}

// Delete elimina un mapeo (soft delete)
func (r *TenantCategoryMappingPostgresRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE tenant_category_mappings
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("mapping with id %s not found", id)
	}

	return nil
}

// GetTenantTaxonomy obtiene la taxonomía completa de un tenant (categorías + mapeos)
func (r *TenantCategoryMappingPostgresRepository) GetTenantTaxonomy(ctx context.Context, tenantID string) ([]*entity.TenantCategoryMapping, error) {
	query := `
		SELECT tcm.id, tcm.tenant_id, tcm.category_id, tcm.marketplace_category_id, tcm.custom_name, tcm.created_at, tcm.updated_at
		FROM tenant_category_mappings tcm
		INNER JOIN marketplace_categories mc ON tcm.marketplace_category_id = mc.id
		WHERE tcm.tenant_id = $1 AND tcm.deleted_at IS NULL AND mc.deleted_at IS NULL
		ORDER BY mc.level, mc.sort_order, mc.name
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMappings(rows)
}

// scanMapping escanea una fila en una entidad TenantCategoryMapping
func (r *TenantCategoryMappingPostgresRepository) scanMapping(row *sql.Row) (*entity.TenantCategoryMapping, error) {
	var mapping entity.TenantCategoryMapping
	var customName sql.NullString

	err := row.Scan(
		&mapping.ID,
		&mapping.TenantID,
		&mapping.CategoryID,
		&mapping.MarketplaceCategoryID,
		&customName,
		&mapping.CreatedAt,
		&mapping.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if customName.Valid {
		mapping.CustomName = &customName.String
	}

	return &mapping, nil
}

// scanMappings escanea múltiples filas en entidades TenantCategoryMapping
func (r *TenantCategoryMappingPostgresRepository) scanMappings(rows *sql.Rows) ([]*entity.TenantCategoryMapping, error) {
	var mappings []*entity.TenantCategoryMapping

	for rows.Next() {
		var mapping entity.TenantCategoryMapping
		var customName sql.NullString

		err := rows.Scan(
			&mapping.ID,
			&mapping.TenantID,
			&mapping.CategoryID,
			&mapping.MarketplaceCategoryID,
			&customName,
			&mapping.CreatedAt,
			&mapping.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		if customName.Valid {
			mapping.CustomName = &customName.String
		}

		mappings = append(mappings, &mapping)
	}

	return mappings, rows.Err()
}
