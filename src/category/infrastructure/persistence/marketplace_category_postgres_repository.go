package persistence

import (
	"context"
	"database/sql"
	"fmt"

	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/port"

	"github.com/google/uuid"
)

// MarketplaceCategoryPostgresRepository implementa el repositorio de categorías marketplace para PostgreSQL
type MarketplaceCategoryPostgresRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewMarketplaceCategoryPostgresRepository crea una nueva instancia del repositorio
func NewMarketplaceCategoryPostgresRepository(db *sql.DB) port.MarketplaceCategoryRepository {
	return &MarketplaceCategoryPostgresRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Save guarda una categoría marketplace
func (r *MarketplaceCategoryPostgresRepository) Save(ctx context.Context, category *entity.MarketplaceCategory) error {
	// Generar ID si no existe
	if category.ID == "" {
		category.ID = uuid.New().String()
	}

	query := `
		INSERT INTO marketplace_categories (
			id, name, slug, description, parent_id, level, is_active, sort_order, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (id) DO UPDATE SET
			name = EXCLUDED.name,
			slug = EXCLUDED.slug,
			description = EXCLUDED.description,
			parent_id = EXCLUDED.parent_id,
			level = EXCLUDED.level,
			is_active = EXCLUDED.is_active,
			sort_order = EXCLUDED.sort_order,
			updated_at = EXCLUDED.updated_at
	`

	_, err := r.db.ExecContext(ctx, query,
		category.ID,
		category.Name,
		category.Slug,
		category.Description,
		category.ParentID,
		category.Level,
		category.IsActive,
		category.SortOrder,
		category.CreatedAt,
		category.UpdatedAt,
	)

	return err
}

// GetByID obtiene una categoría por su ID
func (r *MarketplaceCategoryPostgresRepository) GetByID(ctx context.Context, id string) (*entity.MarketplaceCategory, error) {
	query := `
		SELECT id, name, slug, description, parent_id, level, is_active, sort_order, created_at, updated_at
		FROM marketplace_categories
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	return r.scanCategory(row)
}

// GetBySlug obtiene una categoría por su slug
func (r *MarketplaceCategoryPostgresRepository) GetBySlug(ctx context.Context, slug string) (*entity.MarketplaceCategory, error) {
	query := `
		SELECT id, name, slug, description, parent_id, level, is_active, sort_order, created_at, updated_at
		FROM marketplace_categories
		WHERE slug = $1
	`

	row := r.db.QueryRowContext(ctx, query, slug)
	return r.scanCategory(row)
}

// GetByParentID obtiene las categorías hijas de un parent
func (r *MarketplaceCategoryPostgresRepository) GetByParentID(ctx context.Context, parentID *string) ([]*entity.MarketplaceCategory, error) {
	var query string
	var args []interface{}

	if parentID == nil {
		query = `
			SELECT id, name, slug, description, parent_id, level, is_active, sort_order, created_at, updated_at
			FROM marketplace_categories
			WHERE parent_id IS NULL
			ORDER BY sort_order, name
		`
	} else {
		query = `
			SELECT id, name, slug, description, parent_id, level, is_active, sort_order, created_at, updated_at
			FROM marketplace_categories
			WHERE parent_id = $1
			ORDER BY sort_order, name
		`
		args = append(args, *parentID)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanCategories(rows)
}

// GetRootCategories obtiene las categorías raíz (sin parent)
func (r *MarketplaceCategoryPostgresRepository) GetRootCategories(ctx context.Context) ([]*entity.MarketplaceCategory, error) {
	return r.GetByParentID(ctx, nil)
}

// GetTree obtiene el árbol completo de categorías
func (r *MarketplaceCategoryPostgresRepository) GetTree(ctx context.Context) ([]*entity.MarketplaceCategory, error) {
	query := `
		SELECT id, name, slug, description, parent_id, level, is_active, sort_order, created_at, updated_at
		FROM marketplace_categories
		ORDER BY level, sort_order, name
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanCategories(rows)
}

// FindByCriteria busca categorías según criterios
func (r *MarketplaceCategoryPostgresRepository) FindByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.MarketplaceCategory, error) {
	baseQuery := `
		SELECT id, name, slug, description, parent_id, level, is_active, sort_order, created_at, updated_at
		FROM marketplace_categories
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

	return r.scanCategories(rows)
}

// CountByCriteria cuenta categorías según criterios
func (r *MarketplaceCategoryPostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := `
		SELECT COUNT(*)
		FROM marketplace_categories
	`

	query, args, err := r.converter.ToCountSQL(baseCountQuery, crit)
	if err != nil {
		return 0, fmt.Errorf("invalid criteria: %w", err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	return count, err
}

// Update actualiza una categoría
func (r *MarketplaceCategoryPostgresRepository) Update(ctx context.Context, category *entity.MarketplaceCategory) error {
	query := `
		UPDATE marketplace_categories
		SET name = $2, slug = $3, description = $4, parent_id = $5, level = $6,
		    is_active = $7, sort_order = $8, updated_at = $9
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		category.ID,
		category.Name,
		category.Slug,
		category.Description,
		category.ParentID,
		category.Level,
		category.IsActive,
		category.SortOrder,
		category.UpdatedAt,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category with id %s not found", category.ID)
	}

	return nil
}

// Delete elimina una categoría (hard delete)
func (r *MarketplaceCategoryPostgresRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM marketplace_categories
		WHERE id = $1
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
		return fmt.Errorf("category with id %s not found", id)
	}

	return nil
}

// ExistsBySlug verifica si ya existe una categoría con el slug
func (r *MarketplaceCategoryPostgresRepository) ExistsBySlug(ctx context.Context, slug string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM marketplace_categories
			WHERE slug = $1
		)
	`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, slug).Scan(&exists)
	return exists, err
}

// GetCategoryPath obtiene el path completo de una categoría
func (r *MarketplaceCategoryPostgresRepository) GetCategoryPath(ctx context.Context, categoryID string) ([]*entity.MarketplaceCategory, error) {
	query := `
		WITH RECURSIVE category_path AS (
			-- Caso base: la categoría solicitada
			SELECT id, name, slug, description, parent_id, level, is_active, sort_order, created_at, updated_at, 0 as depth
			FROM marketplace_categories
			WHERE id = $1
			
			UNION ALL
			
			-- Caso recursivo: obtener el parent
			SELECT mc.id, mc.name, mc.slug, mc.description, mc.parent_id, mc.level, mc.is_active, mc.sort_order, mc.created_at, mc.updated_at, cp.depth + 1
			FROM marketplace_categories mc
			INNER JOIN category_path cp ON mc.id = cp.parent_id
		)
		SELECT id, name, slug, description, parent_id, level, is_active, sort_order, created_at, updated_at
		FROM category_path
		ORDER BY depth DESC
	`

	rows, err := r.db.QueryContext(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanCategories(rows)
}

// scanCategory escanea una fila en una entidad MarketplaceCategory
func (r *MarketplaceCategoryPostgresRepository) scanCategory(row *sql.Row) (*entity.MarketplaceCategory, error) {
	var category entity.MarketplaceCategory
	var parentID sql.NullString

	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Slug,
		&category.Description,
		&parentID,
		&category.Level,
		&category.IsActive,
		&category.SortOrder,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if parentID.Valid {
		category.ParentID = &parentID.String
	}

	return &category, nil
}

// scanCategories escanea múltiples filas en entidades MarketplaceCategory
func (r *MarketplaceCategoryPostgresRepository) scanCategories(rows *sql.Rows) ([]*entity.MarketplaceCategory, error) {
	var categories []*entity.MarketplaceCategory

	for rows.Next() {
		var category entity.MarketplaceCategory
		var parentID sql.NullString

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Slug,
			&category.Description,
			&parentID,
			&category.Level,
			&category.IsActive,
			&category.SortOrder,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		if parentID.Valid {
			category.ParentID = &parentID.String
		}

		categories = append(categories, &category)
	}

	return categories, rows.Err()
}
