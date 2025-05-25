package repository

import (
	"context"
	"database/sql"
	"errors"

	"pim/src/category/domain/entity"
	"pim/src/category/domain/exception"
	"pim/src/category/infrastructure/persistence/mapper"
	"pim/src/category/infrastructure/persistence/model"
	"pim/src/shared/domain/criteria"
	sharedCriteria "pim/src/shared/infrastructure/criteria"
)

// CategoryPostgresRepository implementa el repositorio de categorías usando PostgreSQL
type CategoryPostgresRepository struct {
	db        *sql.DB
	mapper    *mapper.CategoryMapper
	converter *sharedCriteria.SQLCriteriaConverter
	*criteria.BaseListRepository[entity.Category]
}

// NewCategoryPostgresRepository crea una nueva instancia del repositorio
func NewCategoryPostgresRepository(db *sql.DB) *CategoryPostgresRepository {
	repo := &CategoryPostgresRepository{
		db:        db,
		mapper:    mapper.NewCategoryMapper(),
		converter: sharedCriteria.NewSQLCriteriaConverter(),
	}

	// Inicializar el repositorio base con criteria
	repo.BaseListRepository = criteria.NewBaseListRepository[entity.Category](repo)

	return repo
}

// SearchByCriteria implementa la búsqueda usando criteria
func (r *CategoryPostgresRepository) SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.Category, error) {
	baseQuery := `
		SELECT id, tenant_id, name, description, parent_id, status, created_at, updated_at
		FROM categories
	`

	query, params := r.converter.ToSelectSQL(baseQuery, crit)

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
func (r *CategoryPostgresRepository) CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM categories"

	query, params := r.converter.ToCountSQL(baseCountQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, params...).Scan(&count)
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
		INSERT INTO categories (id, tenant_id, name, description, parent_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO UPDATE
		SET name = $3, description = $4, parent_id = $5, status = $6, updated_at = $8
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
		SELECT id, tenant_id, name, description, parent_id, status, created_at, updated_at
		FROM categories
		WHERE id = $1 AND tenant_id = $2
	`

	var categoryModel model.Category
	err := r.db.QueryRowContext(ctx, query, id, tenantID).Scan(
		&categoryModel.ID,
		&categoryModel.TenantID,
		&categoryModel.Name,
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
		SELECT id, tenant_id, name, description, parent_id, status, created_at, updated_at
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
