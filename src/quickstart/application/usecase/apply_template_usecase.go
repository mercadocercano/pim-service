package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ApplyTemplateRequest es la petición para aplicar un template
type ApplyTemplateRequest struct {
	TemplateID string `json:"template_id" binding:"required"`
	TenantID   string `json:"tenant_id" binding:"required"`
}

// ApplyTemplateResponse es la respuesta al aplicar un template
type ApplyTemplateResponse struct {
	Success            bool     `json:"success"`
	TemplateID         string   `json:"template_id"`
	TenantID           string   `json:"tenant_id"`
	CategoriesCreated  int      `json:"categories_created"`
	AttributesCreated  int      `json:"attributes_created"`
	CategoryAttributes int      `json:"category_attributes_created"`
	Message            string   `json:"message"`
	CreatedCategories  []string `json:"created_categories,omitempty"`
}

// ApplyTemplateUseCase aplica un template de quickstart a un tenant
type ApplyTemplateUseCase struct {
	db *sql.DB
}

// NewApplyTemplateUseCase crea una nueva instancia del caso de uso
func NewApplyTemplateUseCase(db *sql.DB) *ApplyTemplateUseCase {
	return &ApplyTemplateUseCase{
		db: db,
	}
}

// Execute ejecuta el caso de uso para aplicar un template
func (uc *ApplyTemplateUseCase) Execute(ctx context.Context, req ApplyTemplateRequest) (*ApplyTemplateResponse, error) {
	// Validar tenant UUID
	tenantUUID, err := uuid.Parse(req.TenantID)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant_id format: %w", err)
	}

	// Verificar que el template existe
	if req.TemplateID != "ferreteria-corralon" {
		return nil, fmt.Errorf("template not found: %s", req.TemplateID)
	}

	// Iniciar transacción
	tx, err := uc.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	response := &ApplyTemplateResponse{
		Success:    true,
		TemplateID: req.TemplateID,
		TenantID:   req.TenantID,
	}

	// PASO 1: Crear categorías del tenant desde marketplace_categories
	categoriesCreated, categoryIDs, err := uc.createTenantCategories(ctx, tx, tenantUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to create categories: %w", err)
	}
	response.CategoriesCreated = categoriesCreated
	response.CreatedCategories = categoryIDs

	// PASO 2: Crear atributos del tenant desde marketplace_attributes
	attributesCreated, err := uc.createTenantAttributes(ctx, tx, tenantUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to create attributes: %w", err)
	}
	response.AttributesCreated = attributesCreated

	// PASO 3: Crear relaciones categoría-atributo
	relationsCreated, err := uc.createCategoryAttributeRelations(ctx, tx, tenantUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to create category-attribute relations: %w", err)
	}
	response.CategoryAttributes = relationsCreated

	// Commit de la transacción
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	response.Message = fmt.Sprintf(
		"Template aplicado exitosamente: %d categorías, %d atributos, %d relaciones",
		categoriesCreated, attributesCreated, relationsCreated,
	)

	return response, nil
}

// createTenantCategories crea las categorías del tenant desde marketplace_categories
func (uc *ApplyTemplateUseCase) createTenantCategories(ctx context.Context, tx *sql.Tx, tenantID uuid.UUID) (int, []string, error) {
	// Obtener categorías del template de ferretería desde marketplace_categories
	query := `
		INSERT INTO categories (id, tenant_id, name, slug, description, parent_id, level, sort_order, status, created_at, updated_at)
		SELECT 
			gen_random_uuid()::text,
			$1,
			name,
			slug,
			description,
			NULL,  -- Por ahora sin jerarquía, todas al mismo nivel
			0,
			sort_order,
			'active',
			$2,
			$2
		FROM marketplace_categories
		WHERE parent_id = 'f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00'
		  AND is_active = true
		ON CONFLICT (tenant_id, slug) DO NOTHING
		RETURNING id, name
	`

	now := time.Now()
	rows, err := tx.QueryContext(ctx, query, tenantID.String(), now)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to insert categories: %w", err)
	}
	defer rows.Close()

	var categoryIDs []string
	var categoryNames []string
	count := 0

	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			return 0, nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categoryIDs = append(categoryIDs, id)
		categoryNames = append(categoryNames, name)
		count++
	}

	if err := rows.Err(); err != nil {
		return 0, nil, fmt.Errorf("error iterating categories: %w", err)
	}

	return count, categoryNames, nil
}

// createTenantAttributes crea los atributos del tenant desde marketplace_attributes
func (uc *ApplyTemplateUseCase) createTenantAttributes(ctx context.Context, tx *sql.Tx, tenantID uuid.UUID) (int, error) {
	// Crear atributos del tenant desde marketplace_attributes relacionados con ferretería
	query := `
		INSERT INTO attributes (id, tenant_id, name, description, type, required, options, status, created_at, updated_at)
		SELECT 
			gen_random_uuid()::text,
			$1,
			ma.name,
			ma.description,
			ma.attribute_type::text,  -- Cast del tipo enum a texto
			ma.is_required,
			COALESCE(
				ARRAY(
					SELECT mav.value 
					FROM marketplace_attribute_values mav 
					WHERE mav.attribute_id = ma.id 
					ORDER BY mav.sort_order
				),
				'{}'::text[]
			),
			'active',
			$2,
			$2
		FROM marketplace_attributes ma
		WHERE ma.id LIKE 'fa1e8f2a-%'
		  AND NOT EXISTS (
			SELECT 1 FROM attributes a 
			WHERE a.tenant_id = $1 
			  AND a.name = ma.name
		  )
	`

	now := time.Now()
	result, err := tx.ExecContext(ctx, query, tenantID.String(), now)
	if err != nil {
		return 0, fmt.Errorf("failed to insert attributes: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}

	return int(rowsAffected), nil
}

// createCategoryAttributeRelations crea las relaciones categoría-atributo para el tenant
func (uc *ApplyTemplateUseCase) createCategoryAttributeRelations(ctx context.Context, tx *sql.Tx, tenantID uuid.UUID) (int, error) {
	// Crear relaciones categoría-atributo basadas en marketplace_category_attributes
	// Necesitamos mapear los IDs de marketplace a los IDs del tenant
	query := `
		INSERT INTO category_attributes (id, tenant_id, category_id, attribute_id, allowed_values, status, created_at, updated_at)
		SELECT 
			gen_random_uuid()::text,
			$1,
			tc.id,
			ta.id,
			'{}'::text[],
			'active',
			$2,
			$2
		FROM marketplace_category_attributes mca
		INNER JOIN marketplace_categories mc ON mca.category_id = mc.id
		INNER JOIN marketplace_attributes ma ON mca.attribute_id = ma.id
		INNER JOIN categories tc ON tc.tenant_id = $1 AND tc.slug = mc.slug
		INNER JOIN attributes ta ON ta.tenant_id = $1 AND ta.name = ma.name
		WHERE mca.id LIKE 'fca-%'
		ON CONFLICT (tenant_id, category_id, attribute_id) DO NOTHING
	`

	now := time.Now()
	result, err := tx.ExecContext(ctx, query, tenantID.String(), now)
	if err != nil {
		return 0, fmt.Errorf("failed to insert category-attribute relations: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}

	return int(rowsAffected), nil
}

