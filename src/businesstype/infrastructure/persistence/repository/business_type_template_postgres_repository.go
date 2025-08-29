package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/businesstype/domain/port"
	"saas-mt-pim-service/src/shared/domain/criteria"
	sharedCriteria "saas-mt-pim-service/src/shared/infrastructure/criteria"
)

// BusinessTypeTemplatePostgresRepository implementa el repositorio de BusinessTypeTemplate para PostgreSQL
type BusinessTypeTemplatePostgresRepository struct {
	db        *sql.DB
	converter *sharedCriteria.SQLCriteriaConverter
}

// NewBusinessTypeTemplatePostgresRepository crea una nueva instancia del repositorio
func NewBusinessTypeTemplatePostgresRepository(db *sql.DB) port.BusinessTypeTemplateRepository {
	return &BusinessTypeTemplatePostgresRepository{
		db:        db,
		converter: sharedCriteria.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo template
func (r *BusinessTypeTemplatePostgresRepository) Create(ctx context.Context, template *entity.BusinessTypeTemplate) error {
	categoriesJSON, err := json.Marshal(template.Categories)
	if err != nil {
		return fmt.Errorf("error marshaling categories: %w", err)
	}

	attributesJSON, err := json.Marshal(template.Attributes)
	if err != nil {
		return fmt.Errorf("error marshaling attributes: %w", err)
	}

	productsJSON, err := json.Marshal(template.Products)
	if err != nil {
		return fmt.Errorf("error marshaling products: %w", err)
	}

	brandsJSON, err := json.Marshal(template.Brands)
	if err != nil {
		return fmt.Errorf("error marshaling brands: %w", err)
	}

	metadataJSON, err := json.Marshal(template.Metadata)
	if err != nil {
		return fmt.Errorf("error marshaling metadata: %w", err)
	}

	query := `
		INSERT INTO business_type_templates (
			id, business_type_id, name, description, version, region,
			categories, attributes, products, brands, is_active, is_default,
			metadata, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	_, err = r.db.ExecContext(ctx, query,
		template.ID,
		template.BusinessTypeID,
		template.Name,
		template.Description,
		template.Version,
		template.Region,
		categoriesJSON,
		attributesJSON,
		productsJSON,
		brandsJSON,
		template.IsActive,
		template.IsDefault,
		metadataJSON,
		template.CreatedAt,
		template.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating template: %w", err)
	}

	return nil
}

// Update actualiza un template existente
func (r *BusinessTypeTemplatePostgresRepository) Update(ctx context.Context, template *entity.BusinessTypeTemplate) error {
	categoriesJSON, err := json.Marshal(template.Categories)
	if err != nil {
		return fmt.Errorf("error marshaling categories: %w", err)
	}

	attributesJSON, err := json.Marshal(template.Attributes)
	if err != nil {
		return fmt.Errorf("error marshaling attributes: %w", err)
	}

	productsJSON, err := json.Marshal(template.Products)
	if err != nil {
		return fmt.Errorf("error marshaling products: %w", err)
	}

	brandsJSON, err := json.Marshal(template.Brands)
	if err != nil {
		return fmt.Errorf("error marshaling brands: %w", err)
	}

	metadataJSON, err := json.Marshal(template.Metadata)
	if err != nil {
		return fmt.Errorf("error marshaling metadata: %w", err)
	}

	query := `
		UPDATE business_type_templates SET
			business_type_id = $2,
			name = $3,
			description = $4,
			version = $5,
			region = $6,
			categories = $7,
			attributes = $8,
			products = $9,
			brands = $10,
			is_active = $11,
			is_default = $12,
			metadata = $13,
			updated_at = $14
		WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query,
		template.ID,
		template.BusinessTypeID,
		template.Name,
		template.Description,
		template.Version,
		template.Region,
		categoriesJSON,
		attributesJSON,
		productsJSON,
		brandsJSON,
		template.IsActive,
		template.IsDefault,
		metadataJSON,
		template.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error updating template: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("template with id %s not found", template.ID)
	}

	return nil
}

// FindByID busca un template por ID
func (r *BusinessTypeTemplatePostgresRepository) FindByID(ctx context.Context, id string) (*entity.BusinessTypeTemplate, error) {
	query := `
		SELECT id, business_type_id, name, description, version, region,
			   categories, attributes, products, brands, is_active, is_default,
			   metadata, created_at, updated_at
		FROM business_type_templates
		WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)

	template, err := r.scanTemplate(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding template by id: %w", err)
	}

	return template, nil
}

// FindByBusinessTypeID busca templates por business type ID
func (r *BusinessTypeTemplatePostgresRepository) FindByBusinessTypeID(ctx context.Context, businessTypeID string) ([]*entity.BusinessTypeTemplate, error) {
	query := `
		SELECT id, business_type_id, name, description, version, region,
			   categories, attributes, products, brands, is_active, is_default,
			   metadata, created_at, updated_at
		FROM business_type_templates
		WHERE business_type_id = $1
		ORDER BY is_default DESC, created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, businessTypeID)
	if err != nil {
		return nil, fmt.Errorf("error finding templates by business type id: %w", err)
	}
	defer rows.Close()

	var templates []*entity.BusinessTypeTemplate
	for rows.Next() {
		template, err := r.scanTemplate(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning template: %w", err)
		}
		templates = append(templates, template)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating templates: %w", err)
	}

	return templates, nil
}

// FindByBusinessTypeAndRegion busca templates por business type y región
func (r *BusinessTypeTemplatePostgresRepository) FindByBusinessTypeAndRegion(ctx context.Context, businessTypeID, region string) ([]*entity.BusinessTypeTemplate, error) {
	query := `
		SELECT id, business_type_id, name, description, version, region,
			   categories, attributes, products, brands, is_active, is_default,
			   metadata, created_at, updated_at
		FROM business_type_templates
		WHERE business_type_id = $1 AND region = $2
		ORDER BY is_default DESC, created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, businessTypeID, region)
	if err != nil {
		return nil, fmt.Errorf("error finding templates by business type and region: %w", err)
	}
	defer rows.Close()

	var templates []*entity.BusinessTypeTemplate
	for rows.Next() {
		template, err := r.scanTemplate(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning template: %w", err)
		}
		templates = append(templates, template)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating templates: %w", err)
	}

	return templates, nil
}

// FindDefault busca el template por defecto para un business type y región
func (r *BusinessTypeTemplatePostgresRepository) FindDefault(ctx context.Context, businessTypeID, region string) (*entity.BusinessTypeTemplate, error) {
	query := `
		SELECT id, business_type_id, name, description, version, region,
			   categories, attributes, products, brands, is_active, is_default,
			   metadata, created_at, updated_at
		FROM business_type_templates
		WHERE business_type_id = $1 AND region = $2 AND is_default = true AND is_active = true`

	row := r.db.QueryRowContext(ctx, query, businessTypeID, region)

	template, err := r.scanTemplate(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding default template: %w", err)
	}

	return template, nil
}

// Delete elimina un template
func (r *BusinessTypeTemplatePostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM business_type_templates WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting template: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("template with id %s not found", id)
	}

	return nil
}

// SearchByCriteria busca templates usando criteria
func (r *BusinessTypeTemplatePostgresRepository) SearchByCriteria(ctx context.Context, crit criteria.Criteria) ([]*entity.BusinessTypeTemplate, error) {
	baseQuery := `
		SELECT id, business_type_id, name, description, version, region,
			   categories, attributes, products, brands, is_active, is_default,
			   metadata, created_at, updated_at
		FROM business_type_templates`

	query, params := r.converter.ToSelectSQL(baseQuery, crit)

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("error searching templates by criteria: %w", err)
	}
	defer rows.Close()

	templates := make([]*entity.BusinessTypeTemplate, 0)
	for rows.Next() {
		template, err := r.scanTemplate(rows)
		if err != nil {
			return nil, fmt.Errorf("error scanning template: %w", err)
		}
		templates = append(templates, template)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating templates: %w", err)
	}

	return templates, nil
}

// CountByCriteria cuenta templates usando criteria
func (r *BusinessTypeTemplatePostgresRepository) CountByCriteria(ctx context.Context, crit criteria.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM business_type_templates"

	query, params := r.converter.ToCountSQL(baseCountQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error counting templates by criteria: %w", err)
	}

	return count, nil
}

// scanTemplate escanea una fila y crea una entidad BusinessTypeTemplate
func (r *BusinessTypeTemplatePostgresRepository) scanTemplate(scanner interface {
	Scan(dest ...interface{}) error
}) (*entity.BusinessTypeTemplate, error) {
	var template entity.BusinessTypeTemplate
	var categoriesJSON, attributesJSON, productsJSON, brandsJSON, metadataJSON []byte

	err := scanner.Scan(
		&template.ID,
		&template.BusinessTypeID,
		&template.Name,
		&template.Description,
		&template.Version,
		&template.Region,
		&categoriesJSON,
		&attributesJSON,
		&productsJSON,
		&brandsJSON,
		&template.IsActive,
		&template.IsDefault,
		&metadataJSON,
		&template.CreatedAt,
		&template.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Deserializar JSON fields
	if err = json.Unmarshal(categoriesJSON, &template.Categories); err != nil {
		return nil, fmt.Errorf("error unmarshaling categories: %w", err)
	}

	if err = json.Unmarshal(attributesJSON, &template.Attributes); err != nil {
		return nil, fmt.Errorf("error unmarshaling attributes: %w", err)
	}

	if err = json.Unmarshal(productsJSON, &template.Products); err != nil {
		return nil, fmt.Errorf("error unmarshaling products: %w", err)
	}

	if err = json.Unmarshal(brandsJSON, &template.Brands); err != nil {
		return nil, fmt.Errorf("error unmarshaling brands: %w", err)
	}

	if err = json.Unmarshal(metadataJSON, &template.Metadata); err != nil {
		return nil, fmt.Errorf("error unmarshaling metadata: %w", err)
	}

	return &template, nil
}