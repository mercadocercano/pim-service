package persistence

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"pim/src/marketplace/domain/entity"
	"pim/src/marketplace/domain/port"
	domainCriteria "pim/src/shared/domain/criteria"
	infraCriteria "pim/src/shared/infrastructure/criteria"

	"github.com/google/uuid"
)

// TenantCustomAttributePostgresRepository implementa el repositorio de atributos personalizados tenant para PostgreSQL
type TenantCustomAttributePostgresRepository struct {
	db        *sql.DB
	converter *infraCriteria.SQLCriteriaConverter
}

// NewTenantCustomAttributePostgresRepository crea una nueva instancia del repositorio
func NewTenantCustomAttributePostgresRepository(db *sql.DB) port.TenantCustomAttributeRepository {
	return &TenantCustomAttributePostgresRepository{
		db:        db,
		converter: infraCriteria.NewSQLCriteriaConverter(),
	}
}

// Save guarda un atributo personalizado tenant
func (r *TenantCustomAttributePostgresRepository) Save(ctx context.Context, attribute *entity.TenantCustomAttribute) error {
	// Generar ID si no existe
	if attribute.ID == "" {
		attribute.ID = uuid.New().String()
	}

	// Serializar ValidationRules a JSON
	validationRulesJSON, err := json.Marshal(attribute.ValidationRules)
	if err != nil {
		return fmt.Errorf("error serializing validation rules: %w", err)
	}

	query := `
		INSERT INTO tenant_custom_attributes (
			id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
			validation_rules, sort_order, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT (id) DO UPDATE SET
			tenant_id = EXCLUDED.tenant_id,
			marketplace_category_id = EXCLUDED.marketplace_category_id,
			name = EXCLUDED.name,
			slug = EXCLUDED.slug,
			type = EXCLUDED.type,
			is_filterable = EXCLUDED.is_filterable,
			is_searchable = EXCLUDED.is_searchable,
			validation_rules = EXCLUDED.validation_rules,
			sort_order = EXCLUDED.sort_order,
			updated_at = EXCLUDED.updated_at
	`

	_, err = r.db.ExecContext(ctx, query,
		attribute.ID,
		attribute.TenantID,
		attribute.MarketplaceCategoryID,
		attribute.Name,
		attribute.Slug,
		attribute.Type,
		attribute.IsFilterable,
		attribute.IsSearchable,
		validationRulesJSON,
		attribute.SortOrder,
		attribute.CreatedAt,
		attribute.UpdatedAt,
	)

	return err
}

// GetByID obtiene un atributo por su ID
func (r *TenantCustomAttributePostgresRepository) GetByID(ctx context.Context, id string) (*entity.TenantCustomAttribute, error) {
	query := `
		SELECT id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
		       validation_rules, sort_order, created_at, updated_at
		FROM tenant_custom_attributes
		WHERE id = $1 AND deleted_at IS NULL
	`

	row := r.db.QueryRowContext(ctx, query, id)
	return r.scanAttribute(row)
}

// GetByTenantAndSlug obtiene un atributo por tenant y slug
func (r *TenantCustomAttributePostgresRepository) GetByTenantAndSlug(ctx context.Context, tenantID, slug string) (*entity.TenantCustomAttribute, error) {
	query := `
		SELECT id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
		       validation_rules, sort_order, created_at, updated_at
		FROM tenant_custom_attributes
		WHERE tenant_id = $1 AND slug = $2 AND deleted_at IS NULL
	`

	row := r.db.QueryRowContext(ctx, query, tenantID, slug)
	return r.scanAttribute(row)
}

// GetByTenantID obtiene todos los atributos de un tenant
func (r *TenantCustomAttributePostgresRepository) GetByTenantID(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error) {
	query := `
		SELECT id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
		       validation_rules, sort_order, created_at, updated_at
		FROM tenant_custom_attributes
		WHERE tenant_id = $1 AND deleted_at IS NULL
		ORDER BY sort_order, name
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributes(rows)
}

// GetGlobalByTenant obtiene atributos globales de un tenant (sin categoría específica)
func (r *TenantCustomAttributePostgresRepository) GetGlobalByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error) {
	query := `
		SELECT id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
		       validation_rules, sort_order, created_at, updated_at
		FROM tenant_custom_attributes
		WHERE tenant_id = $1 AND marketplace_category_id IS NULL AND deleted_at IS NULL
		ORDER BY sort_order, name
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributes(rows)
}

// GetByTenantAndCategory obtiene atributos de un tenant para una categoría específica
func (r *TenantCustomAttributePostgresRepository) GetByTenantAndCategory(ctx context.Context, tenantID, marketplaceCategoryID string) ([]*entity.TenantCustomAttribute, error) {
	query := `
		SELECT id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
		       validation_rules, sort_order, created_at, updated_at
		FROM tenant_custom_attributes
		WHERE tenant_id = $1 AND marketplace_category_id = $2 AND deleted_at IS NULL
		ORDER BY sort_order, name
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID, marketplaceCategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributes(rows)
}

// GetFilterableByTenant obtiene atributos filtrables de un tenant
func (r *TenantCustomAttributePostgresRepository) GetFilterableByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error) {
	query := `
		SELECT id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
		       validation_rules, sort_order, created_at, updated_at
		FROM tenant_custom_attributes
		WHERE tenant_id = $1 AND is_filterable = true AND deleted_at IS NULL
		ORDER BY sort_order, name
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributes(rows)
}

// GetSearchableByTenant obtiene atributos buscables de un tenant
func (r *TenantCustomAttributePostgresRepository) GetSearchableByTenant(ctx context.Context, tenantID string) ([]*entity.TenantCustomAttribute, error) {
	query := `
		SELECT id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
		       validation_rules, sort_order, created_at, updated_at
		FROM tenant_custom_attributes
		WHERE tenant_id = $1 AND is_searchable = true AND deleted_at IS NULL
		ORDER BY sort_order, name
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributes(rows)
}

// GetTenantAttributesForCategory obtiene todos los atributos aplicables a una categoría
func (r *TenantCustomAttributePostgresRepository) GetTenantAttributesForCategory(ctx context.Context, tenantID, marketplaceCategoryID string) ([]*entity.TenantCustomAttribute, error) {
	query := `
		SELECT id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
		       validation_rules, sort_order, created_at, updated_at
		FROM tenant_custom_attributes
		WHERE tenant_id = $1 AND (marketplace_category_id = $2 OR marketplace_category_id IS NULL) AND deleted_at IS NULL
		ORDER BY marketplace_category_id NULLS FIRST, sort_order, name
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID, marketplaceCategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributes(rows)
}

// FindByCriteria busca atributos según criterios
func (r *TenantCustomAttributePostgresRepository) FindByCriteria(ctx context.Context, crit domainCriteria.Criteria) ([]*entity.TenantCustomAttribute, error) {
	baseQuery := `
		SELECT id, tenant_id, marketplace_category_id, name, slug, type, is_filterable, is_searchable, 
		       validation_rules, sort_order, created_at, updated_at
		FROM tenant_custom_attributes
		WHERE deleted_at IS NULL
	`

	query, args := r.converter.ToSelectSQL(baseQuery, crit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.scanAttributes(rows)
}

// CountByCriteria cuenta atributos según criterios
func (r *TenantCustomAttributePostgresRepository) CountByCriteria(ctx context.Context, crit domainCriteria.Criteria) (int, error) {
	baseCountQuery := `
		SELECT COUNT(*)
		FROM tenant_custom_attributes
		WHERE deleted_at IS NULL
	`

	query, args := r.converter.ToCountSQL(baseCountQuery, crit)

	var count int
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&count)
	return count, err
}

// Update actualiza un atributo
func (r *TenantCustomAttributePostgresRepository) Update(ctx context.Context, attribute *entity.TenantCustomAttribute) error {
	// Serializar ValidationRules a JSON
	validationRulesJSON, err := json.Marshal(attribute.ValidationRules)
	if err != nil {
		return fmt.Errorf("error serializing validation rules: %w", err)
	}

	query := `
		UPDATE tenant_custom_attributes
		SET tenant_id = $2, marketplace_category_id = $3, name = $4, slug = $5, type = $6,
		    is_filterable = $7, is_searchable = $8, validation_rules = $9, sort_order = $10, updated_at = $11
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query,
		attribute.ID,
		attribute.TenantID,
		attribute.MarketplaceCategoryID,
		attribute.Name,
		attribute.Slug,
		attribute.Type,
		attribute.IsFilterable,
		attribute.IsSearchable,
		validationRulesJSON,
		attribute.SortOrder,
		attribute.UpdatedAt,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("attribute with id %s not found", attribute.ID)
	}

	return nil
}

// Delete elimina un atributo (soft delete)
func (r *TenantCustomAttributePostgresRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE tenant_custom_attributes
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
		return fmt.Errorf("attribute with id %s not found", id)
	}

	return nil
}

// ExistsByTenantAndSlug verifica si ya existe un atributo con el slug en el tenant
func (r *TenantCustomAttributePostgresRepository) ExistsByTenantAndSlug(ctx context.Context, tenantID, slug string, marketplaceCategoryID *string) (bool, error) {
	var query string
	var args []interface{}

	if marketplaceCategoryID == nil {
		query = `
			SELECT EXISTS(
				SELECT 1 FROM tenant_custom_attributes
				WHERE tenant_id = $1 AND slug = $2 AND marketplace_category_id IS NULL AND deleted_at IS NULL
			)
		`
		args = append(args, tenantID, slug)
	} else {
		query = `
			SELECT EXISTS(
				SELECT 1 FROM tenant_custom_attributes
				WHERE tenant_id = $1 AND slug = $2 AND marketplace_category_id = $3 AND deleted_at IS NULL
			)
		`
		args = append(args, tenantID, slug, *marketplaceCategoryID)
	}

	var exists bool
	err := r.db.QueryRowContext(ctx, query, args...).Scan(&exists)
	return exists, err
}

// scanAttribute escanea una fila en una entidad TenantCustomAttribute
func (r *TenantCustomAttributePostgresRepository) scanAttribute(row *sql.Row) (*entity.TenantCustomAttribute, error) {
	var attribute entity.TenantCustomAttribute
	var marketplaceCategoryID sql.NullString
	var validationRulesJSON string

	err := row.Scan(
		&attribute.ID,
		&attribute.TenantID,
		&marketplaceCategoryID,
		&attribute.Name,
		&attribute.Slug,
		&attribute.Type,
		&attribute.IsFilterable,
		&attribute.IsSearchable,
		&validationRulesJSON,
		&attribute.SortOrder,
		&attribute.CreatedAt,
		&attribute.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if marketplaceCategoryID.Valid {
		attribute.MarketplaceCategoryID = &marketplaceCategoryID.String
	}

	// Deserializar ValidationRules desde JSON
	if validationRulesJSON != "" {
		err = json.Unmarshal([]byte(validationRulesJSON), &attribute.ValidationRules)
		if err != nil {
			return nil, fmt.Errorf("error deserializing validation rules: %w", err)
		}
	} else {
		attribute.ValidationRules = make(map[string]interface{})
	}

	return &attribute, nil
}

// scanAttributes escanea múltiples filas en entidades TenantCustomAttribute
func (r *TenantCustomAttributePostgresRepository) scanAttributes(rows *sql.Rows) ([]*entity.TenantCustomAttribute, error) {
	var attributes []*entity.TenantCustomAttribute

	for rows.Next() {
		var attribute entity.TenantCustomAttribute
		var marketplaceCategoryID sql.NullString
		var validationRulesJSON string

		err := rows.Scan(
			&attribute.ID,
			&attribute.TenantID,
			&marketplaceCategoryID,
			&attribute.Name,
			&attribute.Slug,
			&attribute.Type,
			&attribute.IsFilterable,
			&attribute.IsSearchable,
			&validationRulesJSON,
			&attribute.SortOrder,
			&attribute.CreatedAt,
			&attribute.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		if marketplaceCategoryID.Valid {
			attribute.MarketplaceCategoryID = &marketplaceCategoryID.String
		}

		// Deserializar ValidationRules desde JSON
		if validationRulesJSON != "" {
			err = json.Unmarshal([]byte(validationRulesJSON), &attribute.ValidationRules)
			if err != nil {
				return nil, fmt.Errorf("error deserializing validation rules: %w", err)
			}
		} else {
			attribute.ValidationRules = make(map[string]interface{})
		}

		attributes = append(attributes, &attribute)
	}

	return attributes, rows.Err()
}
