package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"saas-mt-pim-service/src/category_attribute/domain/entity"
	"saas-mt-pim-service/src/category_attribute/domain/port"
	"saas-mt-pim-service/src/category_attribute/infrastructure/persistence/mapper"
	"saas-mt-pim-service/src/category_attribute/infrastructure/persistence/model"
	cr "github.com/mercadocercano/criteria"

	"github.com/lib/pq"
)

// CategoryAttributePostgresRepository implementa el repositorio de atributos de categoría usando PostgreSQL
type CategoryAttributePostgresRepository struct {
	db        *sql.DB
	mapper    *mapper.CategoryAttributeMapper
	converter *cr.SQLCriteriaConverter
	*cr.BaseListRepository[entity.CategoryAttribute]
}

// NewCategoryAttributePostgresRepository crea una nueva instancia del repositorio
func NewCategoryAttributePostgresRepository(db *sql.DB) *CategoryAttributePostgresRepository {
	repo := &CategoryAttributePostgresRepository{
		db:        db,
		mapper:    mapper.NewCategoryAttributeMapper(),
		converter: cr.NewSQLCriteriaConverter(),
	}

	// Inicializar el repositorio base con criteria
	repo.BaseListRepository = cr.NewBaseListRepository[entity.CategoryAttribute](repo)

	return repo
}

// SearchByCriteria implementa la búsqueda usando criteria
func (r *CategoryAttributePostgresRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.CategoryAttribute, error) {
	baseQuery := `
		SELECT id, tenant_id, category_id, attribute_id, allowed_values, status, created_at, updated_at
		FROM category_attributes
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

	var categoryAttributes []*model.CategoryAttribute
	for rows.Next() {
		var categoryAttribute model.CategoryAttribute
		err := rows.Scan(
			&categoryAttribute.ID,
			&categoryAttribute.TenantID,
			&categoryAttribute.CategoryID,
			&categoryAttribute.AttributeID,
			&categoryAttribute.AllowedValues,
			&categoryAttribute.Status,
			&categoryAttribute.CreatedAt,
			&categoryAttribute.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categoryAttributes = append(categoryAttributes, &categoryAttribute)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return r.mapper.ToEntityList(categoryAttributes), nil
}

// CountByCriteria implementa el conteo usando criteria
func (r *CategoryAttributePostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM category_attributes"

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

// Create implementa la interfaz CategoryAttributeRepository creando un nuevo atributo de categoría
func (r *CategoryAttributePostgresRepository) Create(ctx context.Context, categoryAttribute *entity.CategoryAttribute) error {
	query := `
		INSERT INTO category_attributes (id, tenant_id, category_id, attribute_id, allowed_values, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	categoryAttributeModel := r.mapper.ToModel(categoryAttribute)

	_, err := r.db.ExecContext(ctx, query,
		categoryAttributeModel.ID,
		categoryAttributeModel.TenantID,
		categoryAttributeModel.CategoryID,
		categoryAttributeModel.AttributeID,
		pq.Array(categoryAttributeModel.AllowedValues),
		categoryAttributeModel.Status,
		categoryAttributeModel.CreatedAt,
		categoryAttributeModel.UpdatedAt,
	)

	return err
}

// Update implementa la interfaz CategoryAttributeRepository actualizando un atributo de categoría existente
func (r *CategoryAttributePostgresRepository) Update(ctx context.Context, categoryAttribute *entity.CategoryAttribute) error {
	query := `
		UPDATE category_attributes 
		SET allowed_values = $3, status = $4, updated_at = $5
		WHERE id = $1 AND tenant_id = $2
	`

	categoryAttributeModel := r.mapper.ToModel(categoryAttribute)

	result, err := r.db.ExecContext(ctx, query,
		categoryAttributeModel.ID,
		categoryAttributeModel.TenantID,
		pq.Array(categoryAttributeModel.AllowedValues),
		categoryAttributeModel.Status,
		categoryAttributeModel.UpdatedAt,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("atributo de categoría no encontrado")
	}

	return nil
}

// FindByID busca un atributo de categoría por su ID
func (r *CategoryAttributePostgresRepository) FindByID(ctx context.Context, id string, tenantID string) (*entity.CategoryAttribute, error) {
	query := `
		SELECT id, tenant_id, category_id, attribute_id, allowed_values, status, created_at, updated_at
		FROM category_attributes
		WHERE id = $1 AND tenant_id = $2
	`

	var categoryAttributeModel model.CategoryAttribute
	err := r.db.QueryRowContext(ctx, query, id, tenantID).Scan(
		&categoryAttributeModel.ID,
		&categoryAttributeModel.TenantID,
		&categoryAttributeModel.CategoryID,
		&categoryAttributeModel.AttributeID,
		&categoryAttributeModel.AllowedValues,
		&categoryAttributeModel.Status,
		&categoryAttributeModel.CreatedAt,
		&categoryAttributeModel.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return r.mapper.ToEntity(&categoryAttributeModel), nil
}

// FindByCategoryAndTenant busca atributos de categoría por categoryID y tenantID
func (r *CategoryAttributePostgresRepository) FindByCategoryAndTenant(ctx context.Context, categoryID string, tenantID string) ([]*entity.CategoryAttribute, error) {
	query := `
		SELECT id, tenant_id, category_id, attribute_id, allowed_values, status, created_at, updated_at
		FROM category_attributes
		WHERE category_id = $1 AND tenant_id = $2
		ORDER BY created_at
	`

	rows, err := r.db.QueryContext(ctx, query, categoryID, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categoryAttributes []*model.CategoryAttribute
	for rows.Next() {
		var categoryAttribute model.CategoryAttribute
		err := rows.Scan(
			&categoryAttribute.ID,
			&categoryAttribute.TenantID,
			&categoryAttribute.CategoryID,
			&categoryAttribute.AttributeID,
			&categoryAttribute.AllowedValues,
			&categoryAttribute.Status,
			&categoryAttribute.CreatedAt,
			&categoryAttribute.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categoryAttributes = append(categoryAttributes, &categoryAttribute)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return r.mapper.ToEntityList(categoryAttributes), nil
}

// FindByTenant recupera todos los atributos de categoría de un tenant
func (r *CategoryAttributePostgresRepository) FindByTenant(ctx context.Context, tenantID string) ([]*entity.CategoryAttribute, error) {
	query := `
		SELECT id, tenant_id, category_id, attribute_id, allowed_values, status, created_at, updated_at
		FROM category_attributes
		WHERE tenant_id = $1
		ORDER BY created_at
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categoryAttributes []*model.CategoryAttribute
	for rows.Next() {
		var categoryAttribute model.CategoryAttribute
		err := rows.Scan(
			&categoryAttribute.ID,
			&categoryAttribute.TenantID,
			&categoryAttribute.CategoryID,
			&categoryAttribute.AttributeID,
			&categoryAttribute.AllowedValues,
			&categoryAttribute.Status,
			&categoryAttribute.CreatedAt,
			&categoryAttribute.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categoryAttributes = append(categoryAttributes, &categoryAttribute)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return r.mapper.ToEntityList(categoryAttributes), nil
}

// FindByAttributeAndCategory busca una relación específica por attributeID, categoryID y tenantID
func (r *CategoryAttributePostgresRepository) FindByAttributeAndCategory(ctx context.Context, attributeID, categoryID, tenantID string) (*entity.CategoryAttribute, error) {
	query := `
		SELECT id, tenant_id, category_id, attribute_id, allowed_values, status, created_at, updated_at
		FROM category_attributes
		WHERE attribute_id = $1 AND category_id = $2 AND tenant_id = $3
	`

	var categoryAttributeModel model.CategoryAttribute
	err := r.db.QueryRowContext(ctx, query, attributeID, categoryID, tenantID).Scan(
		&categoryAttributeModel.ID,
		&categoryAttributeModel.TenantID,
		&categoryAttributeModel.CategoryID,
		&categoryAttributeModel.AttributeID,
		&categoryAttributeModel.AllowedValues,
		&categoryAttributeModel.Status,
		&categoryAttributeModel.CreatedAt,
		&categoryAttributeModel.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return r.mapper.ToEntity(&categoryAttributeModel), nil
}

// Delete elimina un atributo de categoría por su ID
func (r *CategoryAttributePostgresRepository) Delete(ctx context.Context, id string, tenantID string) error {
	query := `DELETE FROM category_attributes WHERE id = $1 AND tenant_id = $2`
	result, err := r.db.ExecContext(ctx, query, id, tenantID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("atributo de categoría no encontrado")
	}

	return nil
}

// FindDetailedByCategoryAndTenant busca atributos completos con JOIN para una categoría
func (r *CategoryAttributePostgresRepository) FindDetailedByCategoryAndTenant(ctx context.Context, categoryID string, tenantID string) ([]*port.DetailedCategoryAttribute, error) {
	query := `
		SELECT 
			ca.id,
			ca.category_id,
			ca.attribute_id,
			a.name as attribute_name,
			a.type as attribute_type,
			a.description,
			a.required,
			a.options as attribute_options,
			ca.allowed_values,
			ca.status,
			ca.created_at,
			ca.updated_at
		FROM category_attributes ca
		INNER JOIN attributes a ON ca.attribute_id = a.id AND ca.tenant_id = a.tenant_id
		WHERE ca.category_id = $1 AND ca.tenant_id = $2 AND ca.status = 'active'
		ORDER BY ca.created_at
	`

	rows, err := r.db.QueryContext(ctx, query, categoryID, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var detailedAttributes []*port.DetailedCategoryAttribute
	for rows.Next() {
		var detailed port.DetailedCategoryAttribute
		var attributeOptions pq.StringArray
		var allowedValues pq.StringArray
		var status string

		err := rows.Scan(
			&detailed.ID,
			&detailed.CategoryID,
			&detailed.AttributeID,
			&detailed.AttributeName,
			&detailed.AttributeType,
			&detailed.Description,
			&detailed.Required,
			&attributeOptions,
			&allowedValues,
			&status,
			&detailed.CreatedAt,
			&detailed.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		detailed.AttributeOptions = []string(attributeOptions)
		detailed.AllowedValues = []string(allowedValues)
		detailed.Active = status == "active"

		detailedAttributes = append(detailedAttributes, &detailed)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return detailedAttributes, nil
}
