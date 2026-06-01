package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"saas-mt-pim-service/src/attribute/domain/entity"
	cr "github.com/mercadocercano/criteria"
	"time"
)

// MarketplaceAttributePostgresRepository implementa el repositorio de marketplace attributes usando PostgreSQL
type MarketplaceAttributePostgresRepository struct {
	db        *sql.DB
	converter *cr.SQLCriteriaConverter
}

// NewMarketplaceAttributePostgresRepository crea una nueva instancia del repositorio
func NewMarketplaceAttributePostgresRepository(db *sql.DB) *MarketplaceAttributePostgresRepository {
	return &MarketplaceAttributePostgresRepository{
		db:        db,
		converter: cr.NewSQLCriteriaConverter(),
	}
}

// Create crea un nuevo marketplace attribute
func (r *MarketplaceAttributePostgresRepository) Create(ctx context.Context, attribute *entity.MarketplaceAttribute) error {
	// Convertir validation_rules a JSON
	validationRulesJSON, err := json.Marshal(attribute.ValidationRules)
	if err != nil {
		return fmt.Errorf("error marshaling validation rules: %v", err)
	}

	query := `
		INSERT INTO marketplace_attributes (
			id, name, slug, type, is_filterable, is_searchable, 
			is_required_for_listing, validation_rules, sort_order, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)
	`

	_, err = r.db.ExecContext(ctx, query,
		attribute.ID,
		attribute.Name,
		attribute.Slug,
		attribute.Type,
		attribute.IsFilterable,
		attribute.IsSearchable,
		attribute.IsRequiredForListing,
		string(validationRulesJSON),
		attribute.SortOrder,
		attribute.CreatedAt,
		attribute.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error creando marketplace attribute: %v", err)
		return fmt.Errorf("error creating marketplace attribute: %v", err)
	}

	return nil
}

// Update actualiza un marketplace attribute existente
func (r *MarketplaceAttributePostgresRepository) Update(ctx context.Context, attribute *entity.MarketplaceAttribute) error {
	// Convertir validation_rules a JSON
	validationRulesJSON, err := json.Marshal(attribute.ValidationRules)
	if err != nil {
		return fmt.Errorf("error marshaling validation rules: %v", err)
	}

	query := `
		UPDATE marketplace_attributes SET
			name = $2,
			slug = $3,
			type = $4,
			is_filterable = $5,
			is_searchable = $6,
			is_required_for_listing = $7,
			validation_rules = $8,
			sort_order = $9,
			updated_at = $10
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		attribute.ID,
		attribute.Name,
		attribute.Slug,
		attribute.Type,
		attribute.IsFilterable,
		attribute.IsSearchable,
		attribute.IsRequiredForListing,
		string(validationRulesJSON),
		attribute.SortOrder,
		time.Now(),
	)

	if err != nil {
		log.Printf("Error actualizando marketplace attribute: %v", err)
		return fmt.Errorf("error updating marketplace attribute: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("marketplace attribute not found")
	}

	return nil
}

// FindByID busca un marketplace attribute por su ID
func (r *MarketplaceAttributePostgresRepository) FindByID(ctx context.Context, id string) (*entity.MarketplaceAttribute, error) {
	query := `
		SELECT id, name, slug, type, is_filterable, is_searchable, 
			   is_required_for_listing, validation_rules, sort_order, created_at, updated_at
		FROM marketplace_attributes 
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	return r.scanMarketplaceAttribute(row)
}

// FindByName busca un marketplace attribute por su nombre
func (r *MarketplaceAttributePostgresRepository) FindByName(ctx context.Context, name string) (*entity.MarketplaceAttribute, error) {
	query := `
		SELECT id, name, slug, type, is_filterable, is_searchable, 
			   is_required_for_listing, validation_rules, sort_order, created_at, updated_at
		FROM marketplace_attributes 
		WHERE name = $1
	`

	row := r.db.QueryRowContext(ctx, query, name)
	return r.scanMarketplaceAttribute(row)
}

// FindBySlug busca un marketplace attribute por su slug
func (r *MarketplaceAttributePostgresRepository) FindBySlug(ctx context.Context, slug string) (*entity.MarketplaceAttribute, error) {
	query := `
		SELECT id, name, slug, type, is_filterable, is_searchable, 
			   is_required_for_listing, validation_rules, sort_order, created_at, updated_at
		FROM marketplace_attributes 
		WHERE slug = $1
	`

	row := r.db.QueryRowContext(ctx, query, slug)
	return r.scanMarketplaceAttribute(row)
}

// FindAll obtiene todos los marketplace attributes
func (r *MarketplaceAttributePostgresRepository) FindAll(ctx context.Context) ([]*entity.MarketplaceAttribute, error) {
	query := `
		SELECT id, name, slug, type, is_filterable, is_searchable, 
			   is_required_for_listing, validation_rules, sort_order, created_at, updated_at
		FROM marketplace_attributes 
		ORDER BY sort_order ASC, name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanMarketplaceAttributes(rows)
}

// Delete elimina un marketplace attribute
func (r *MarketplaceAttributePostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM marketplace_attributes WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error eliminando marketplace attribute: %v", err)
		return fmt.Errorf("error deleting marketplace attribute: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("marketplace attribute not found")
	}

	return nil
}

// SearchByCriteria busca marketplace attributes usando criterios
func (r *MarketplaceAttributePostgresRepository) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.MarketplaceAttribute, error) {
	baseQuery := `
		SELECT id, name, slug, type, is_filterable, is_searchable, 
			   is_required_for_listing, validation_rules, sort_order, created_at, updated_at
		FROM marketplace_attributes
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

	return r.scanMarketplaceAttributes(rows)
}

// CountByCriteria cuenta marketplace attributes usando criterios
func (r *MarketplaceAttributePostgresRepository) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	baseCountQuery := "SELECT COUNT(*) FROM marketplace_attributes"

	query, params, err := r.converter.ToCountSQL(baseCountQuery, crit)
	if err != nil {
		return 0, fmt.Errorf("invalid criteria: %w", err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	return count, err
}

// scanMarketplaceAttribute escanea una fila y devuelve un marketplace attribute
func (r *MarketplaceAttributePostgresRepository) scanMarketplaceAttribute(row *sql.Row) (*entity.MarketplaceAttribute, error) {
	var id, name, slug, attrType string
	var isFilterable, isSearchable, isRequiredForListing bool
	var validationRulesJSON string
	var sortOrder int
	var createdAt, updatedAt time.Time

	err := row.Scan(
		&id, &name, &slug, &attrType, &isFilterable, &isSearchable,
		&isRequiredForListing, &validationRulesJSON, &sortOrder, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Parsear validation_rules JSON
	var validationRules map[string]interface{}
	if err := json.Unmarshal([]byte(validationRulesJSON), &validationRules); err != nil {
		log.Printf("Error parsing validation rules JSON: %v", err)
		validationRules = make(map[string]interface{})
	}

	return &entity.MarketplaceAttribute{
		ID:                   id,
		Name:                 name,
		Slug:                 slug,
		Type:                 attrType,
		IsFilterable:         isFilterable,
		IsSearchable:         isSearchable,
		IsRequiredForListing: isRequiredForListing,
		ValidationRules:      validationRules,
		SortOrder:            sortOrder,
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
	}, nil
}

// IsInUse verifica si el atributo está asignado a alguna variante de producto
func (r *MarketplaceAttributePostgresRepository) IsInUse(ctx context.Context, id string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM variant_marketplace_attributes WHERE marketplace_attribute_id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking attribute usage: %w", err)
	}
	return count > 0, nil
}

// scanMarketplaceAttributes escanea múltiples filas y devuelve un slice de marketplace attributes
func (r *MarketplaceAttributePostgresRepository) scanMarketplaceAttributes(rows *sql.Rows) ([]*entity.MarketplaceAttribute, error) {
	var attributes []*entity.MarketplaceAttribute

	for rows.Next() {
		var id, name, slug, attrType string
		var isFilterable, isSearchable, isRequiredForListing bool
		var validationRulesJSON string
		var sortOrder int
		var createdAt, updatedAt time.Time

		err := rows.Scan(
			&id, &name, &slug, &attrType, &isFilterable, &isSearchable,
			&isRequiredForListing, &validationRulesJSON, &sortOrder, &createdAt, &updatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Parsear validation_rules JSON
		var validationRules map[string]interface{}
		if err := json.Unmarshal([]byte(validationRulesJSON), &validationRules); err != nil {
			log.Printf("Error parsing validation rules JSON: %v", err)
			validationRules = make(map[string]interface{})
		}

		attribute := &entity.MarketplaceAttribute{
			ID:                   id,
			Name:                 name,
			Slug:                 slug,
			Type:                 attrType,
			IsFilterable:         isFilterable,
			IsSearchable:         isSearchable,
			IsRequiredForListing: isRequiredForListing,
			ValidationRules:      validationRules,
			SortOrder:            sortOrder,
			CreatedAt:            createdAt,
			UpdatedAt:            updatedAt,
		}

		attributes = append(attributes, attribute)
	}

	return attributes, nil
}
