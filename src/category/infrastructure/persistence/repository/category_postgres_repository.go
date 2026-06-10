package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	cr "github.com/mercadocercano/criteria"
	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/exception"
	"saas-mt-pim-service/src/category/infrastructure/persistence/mapper"
	"saas-mt-pim-service/src/category/infrastructure/persistence/model"
)

// CategoryPostgresRepository implementa el repositorio de categorías usando PostgreSQL
type CategoryPostgresRepository struct {
	db        *sql.DB
	mapper    *mapper.CategoryMapper
	converter *cr.SQLCriteriaConverter
	*cr.BaseListRepository[entity.Category]
}

// NewCategoryPostgresRepository crea una nueva instancia del repositorio
func NewCategoryPostgresRepository(db *sql.DB) *CategoryPostgresRepository {
	repo := &CategoryPostgresRepository{
		db:        db,
		mapper:    mapper.NewCategoryMapper(),
		converter: cr.NewSQLCriteriaConverter(),
	}

	// Inicializar el repositorio base con criteria
	repo.BaseListRepository = cr.NewBaseListRepository[entity.Category](repo)

	return repo
}

// SearchByCriteria implementa la búsqueda usando criteria
func (r *CategoryPostgresRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.Category, error) {
	baseQuery := `
		SELECT id, tenant_id, name, slug, description, parent_id, status, created_at, updated_at
		FROM categories
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

	var categories []*model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(
			&category.ID,
			&category.TenantID,
			&category.Name,
			&category.Slug,
			&category.Description,
			&category.ParentID,
			&category.Status,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return r.mapper.ToEntityList(categories), nil
}

// CountByCriteria implementa el conteo usando criteria
func (r *CategoryPostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM categories"

	query, params, err := r.converter.ToCountSQL(baseCountQuery, crit)
	if err != nil {
		return 0, fmt.Errorf("invalid criteria: %w", err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Create implementa la interfaz CategoryRepository creando una nueva categoría
func (r *CategoryPostgresRepository) Create(ctx context.Context, category *entity.Category) error {
	return r.Save(ctx, category)
}

// Update implementa la interfaz CategoryRepository actualizando una categoría existente
func (r *CategoryPostgresRepository) Update(ctx context.Context, category *entity.Category) error {
	return r.Save(ctx, category)
}

// Save guarda una categoría en la base de datos
func (r *CategoryPostgresRepository) Save(ctx context.Context, category *entity.Category) error {
	query := `
		INSERT INTO categories (id, tenant_id, name, slug, description, parent_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE
		SET name = $3, slug = $4, description = $5, parent_id = $6, status = $7, updated_at = $9
	`

	categoryModel := r.mapper.ToModel(category)

	var parentIDValue interface{}
	if categoryModel.ParentID == nil {
		parentIDValue = nil
	} else {
		parentIDValue = *categoryModel.ParentID
	}

	_, err := r.db.ExecContext(ctx, query,
		categoryModel.ID,
		categoryModel.TenantID,
		categoryModel.Name,
		categoryModel.Slug,
		categoryModel.Description,
		parentIDValue,
		categoryModel.Status,
		categoryModel.CreatedAt,
		categoryModel.UpdatedAt,
	)

	return err
}

// FindByID busca una categoría por su ID
func (r *CategoryPostgresRepository) FindByID(ctx context.Context, id string, tenantID string) (*entity.Category, error) {
	query := `
		SELECT id, tenant_id, name, slug, description, parent_id, status, created_at, updated_at
		FROM categories
		WHERE id = $1 AND tenant_id = $2
	`

	var categoryModel model.Category
	err := r.db.QueryRowContext(ctx, query, id, tenantID).Scan(
		&categoryModel.ID,
		&categoryModel.TenantID,
		&categoryModel.Name,
		&categoryModel.Slug,
		&categoryModel.Description,
		&categoryModel.ParentID,
		&categoryModel.Status,
		&categoryModel.CreatedAt,
		&categoryModel.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exception.ErrCategoryNotFound
		}
		return nil, err
	}

	return r.mapper.ToEntity(&categoryModel), nil
}

// FindAll retorna todas las categorías de un tenant
func (r *CategoryPostgresRepository) FindAll(ctx context.Context, tenantID string) ([]*entity.Category, error) {
	query := `
		SELECT id, tenant_id, name, slug, description, parent_id, status, created_at, updated_at
		FROM categories
		WHERE tenant_id = $1
		ORDER BY name
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(
			&category.ID,
			&category.TenantID,
			&category.Name,
			&category.Slug,
			&category.Description,
			&category.ParentID,
			&category.Status,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return r.mapper.ToEntityList(categories), nil
}

// Delete elimina una categoría por su ID
func (r *CategoryPostgresRepository) Delete(ctx context.Context, id string, tenantID string) error {
	query := `DELETE FROM categories WHERE id = $1 AND tenant_id = $2`
	result, err := r.db.ExecContext(ctx, query, id, tenantID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return exception.ErrCategoryNotFound
	}

	return nil
}

// FindBySlug busca una categoría por su slug (HITO 2)
func (r *CategoryPostgresRepository) FindBySlug(ctx context.Context, tenantID uuid.UUID, slug string) (*entity.Category, error) {
	query := `
		SELECT id, tenant_id, name, slug, description, parent_id, status, created_at, updated_at
		FROM categories
		WHERE tenant_id = $1 AND slug = $2
	`

	var categoryModel model.Category
	err := r.db.QueryRowContext(ctx, query, tenantID.String(), slug).Scan(
		&categoryModel.ID,
		&categoryModel.TenantID,
		&categoryModel.Name,
		&categoryModel.Slug,
		&categoryModel.Description,
		&categoryModel.ParentID,
		&categoryModel.Status,
		&categoryModel.CreatedAt,
		&categoryModel.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exception.ErrCategoryNotFound
		}
		return nil, err
	}

	return r.mapper.ToEntity(&categoryModel), nil
}

// FindByName busca una categoría por su nombre exacto (HITO 2)
func (r *CategoryPostgresRepository) FindByName(ctx context.Context, tenantID uuid.UUID, name string) (*entity.Category, error) {
	query := `
		SELECT id, tenant_id, name, slug, description, parent_id, status, created_at, updated_at
		FROM categories
		WHERE tenant_id = $1 AND LOWER(name) = LOWER($2)
		LIMIT 1
	`

	var categoryModel model.Category
	err := r.db.QueryRowContext(ctx, query, tenantID.String(), name).Scan(
		&categoryModel.ID,
		&categoryModel.TenantID,
		&categoryModel.Name,
		&categoryModel.Slug,
		&categoryModel.Description,
		&categoryModel.ParentID,
		&categoryModel.Status,
		&categoryModel.CreatedAt,
		&categoryModel.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exception.ErrCategoryNotFound
		}
		return nil, err
	}

	return r.mapper.ToEntity(&categoryModel), nil
}
