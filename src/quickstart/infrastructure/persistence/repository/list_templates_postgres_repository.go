package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"saas-mt-pim-service/src/quickstart/domain/port"
)

// ListTemplatesPostgresRepository implementa ListTemplatesRepository para PostgreSQL
type ListTemplatesPostgresRepository struct {
	db *sql.DB
}

// NewListTemplatesPostgresRepository crea una nueva instancia del repositorio
func NewListTemplatesPostgresRepository(db *sql.DB) port.ListTemplatesRepository {
	return &ListTemplatesPostgresRepository{db: db}
}

type templateCategoryPayload struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	ParentSlug string `json:"parent_slug"`
	Level      int    `json:"level"`
}

// LoadTemplatesFromBusinessTypeTemplates carga templates desde business_type_templates
func (r *ListTemplatesPostgresRepository) LoadTemplatesFromBusinessTypeTemplates(ctx context.Context) ([]port.ListTemplate, error) {
	query := `
		SELECT btt.id, btt.name, btt.description, btt.categories, btt.brands, btt.is_default, btt.region,
		       COALESCE(bt.code, '') as slug, COALESCE(bt.icon, '') as icon,
		       COALESCE(jsonb_array_length(btt.categories), 0) as total_categories,
		       COALESCE(jsonb_array_length(btt.products), 0) as total_products
		FROM business_type_templates btt
		LEFT JOIN business_types bt ON bt.id = btt.business_type_id
		WHERE btt.is_active = true
		  AND btt.region IN ('AR', 'GLOBAL')
		  AND COALESCE(jsonb_array_length(btt.categories), 0) >= 1
		ORDER BY btt.is_default DESC,
		         CASE WHEN btt.region = 'AR' THEN 0 ELSE 1 END,
		         btt.name
		LIMIT 50
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying business_type_templates: %w", err)
	}
	defer rows.Close()

	templates := make([]port.ListTemplate, 0)
	for rows.Next() {
		var id, name, description, slug, icon sql.NullString
		var categoriesRaw, brandsRaw []byte
		var isDefault sql.NullBool
		var region sql.NullString
		var totalCategories, totalProducts int
		if err := rows.Scan(&id, &name, &description, &categoriesRaw, &brandsRaw, &isDefault, &region, &slug, &icon, &totalCategories, &totalProducts); err != nil {
			return nil, fmt.Errorf("error scanning business_type_template: %w", err)
		}
		if !id.Valid || id.String == "" {
			continue
		}

		categories, err := parseTemplateCategoryNames(categoriesRaw)
		if err != nil {
			return nil, err
		}

		brands := parseTemplateBrands(brandsRaw)

		// Slug para CSV: business_type code si existe, sino id
		templateSlug := id.String
		if slug.Valid && slug.String != "" {
			templateSlug = slug.String
		}

		templates = append(templates, port.ListTemplate{
			ID:              id.String,
			Name:            name.String,
			Slug:            templateSlug,
			Description:     description.String,
			Icon:            icon.String,
			Categories:      categories,
			Brands:          brands,
			TotalCategories: totalCategories,
			TotalProducts:   totalProducts,
			IsActive:        true,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating business_type_templates: %w", err)
	}

	return templates, nil
}

// LoadTemplatesComputed carga templates usando el surtido COMPUTADO desde
// global_products (business_type_product_templates.priority_brands +
// suggested_products). Mantiene la taxonomía (categories) y la metadata
// (name/description/icon/slug) del editorial. El refresh solo computa los
// templates is_default → para los que no tienen surtido computado, cae a las
// marcas y el conteo de productos editoriales por-template.
func (r *ListTemplatesPostgresRepository) LoadTemplatesComputed(ctx context.Context) ([]port.ListTemplate, error) {
	query := `
		SELECT btt.id, btt.name, btt.description, btt.categories, btt.brands, btt.is_default, btt.region,
		       COALESCE(bt.code, '') as slug, COALESCE(bt.icon, '') as icon,
		       COALESCE(jsonb_array_length(btt.categories), 0) as total_categories,
		       COALESCE(jsonb_array_length(btt.products), 0) as editorial_products,
		       btpt.priority_brands,
		       COALESCE(jsonb_array_length(btpt.suggested_products), 0) as computed_products
		FROM business_type_templates btt
		LEFT JOIN business_types bt ON bt.id = btt.business_type_id
		LEFT JOIN business_type_product_templates btpt ON btpt.business_type_template_id = btt.id
		WHERE btt.is_active = true
		  AND btt.region IN ('AR', 'GLOBAL')
		  AND COALESCE(jsonb_array_length(btt.categories), 0) >= 1
		ORDER BY btt.is_default DESC,
		         CASE WHEN btt.region = 'AR' THEN 0 ELSE 1 END,
		         btt.name
		LIMIT 50
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying computed templates: %w", err)
	}
	defer rows.Close()

	templates := make([]port.ListTemplate, 0)
	for rows.Next() {
		var id, name, description, slug, icon sql.NullString
		var categoriesRaw, editorialBrandsRaw, priorityBrandsRaw []byte
		var isDefault sql.NullBool
		var region sql.NullString
		var totalCategories, editorialProducts, computedProducts int
		if err := rows.Scan(&id, &name, &description, &categoriesRaw, &editorialBrandsRaw, &isDefault, &region, &slug, &icon, &totalCategories, &editorialProducts, &priorityBrandsRaw, &computedProducts); err != nil {
			return nil, fmt.Errorf("error scanning computed template: %w", err)
		}
		if !id.Valid || id.String == "" {
			continue
		}

		categories, err := parseTemplateCategoryNames(categoriesRaw)
		if err != nil {
			return nil, err
		}

		// Fallback por-template: si no hay surtido computado para este template,
		// usar marcas y conteo editoriales.
		var brands []port.ListTemplateBrand
		totalProducts := editorialProducts
		if computedProducts > 0 {
			brands = parsePriorityBrands(priorityBrandsRaw)
			totalProducts = computedProducts
		} else {
			brands = parseTemplateBrands(editorialBrandsRaw)
		}

		templateSlug := id.String
		if slug.Valid && slug.String != "" {
			templateSlug = slug.String
		}

		templates = append(templates, port.ListTemplate{
			ID:              id.String,
			Name:            name.String,
			Slug:            templateSlug,
			Description:     description.String,
			Icon:            icon.String,
			Categories:      categories,
			Brands:          brands,
			TotalCategories: totalCategories,
			TotalProducts:   totalProducts,
			IsActive:        true,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating computed templates: %w", err)
	}

	return templates, nil
}

// parsePriorityBrands convierte priority_brands (array JSON de strings, ordenado
// por frecuencia descendente) en ListTemplateBrand. Sin logo: el computado solo
// tiene el nombre de la marca (follow-up: enriquecer con logos desde brands).
func parsePriorityBrands(raw []byte) []port.ListTemplateBrand {
	if len(raw) == 0 {
		return nil
	}
	var names []string
	if err := json.Unmarshal(raw, &names); err != nil {
		return nil
	}
	brands := make([]port.ListTemplateBrand, 0, len(names))
	for _, n := range names {
		if n == "" {
			continue
		}
		brands = append(brands, port.ListTemplateBrand{Name: n})
	}
	return brands
}

type templateBrandPayload struct {
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
}

func parseTemplateBrands(brandsRaw []byte) []port.ListTemplateBrand {
	if len(brandsRaw) == 0 {
		return nil
	}
	var payload []templateBrandPayload
	if err := json.Unmarshal(brandsRaw, &payload); err != nil {
		return nil
	}
	brands := make([]port.ListTemplateBrand, 0, len(payload))
	for _, b := range payload {
		if b.Name == "" {
			continue
		}
		brands = append(brands, port.ListTemplateBrand{
			Name:    b.Name,
			LogoURL: b.LogoURL,
		})
	}
	return brands
}

func parseTemplateCategoryNames(categoriesRaw []byte) ([]string, error) {
	if len(categoriesRaw) == 0 {
		return nil, nil
	}

	var payload []templateCategoryPayload
	if err := json.Unmarshal(categoriesRaw, &payload); err != nil {
		return nil, fmt.Errorf("error parsing template categories: %w", err)
	}

	categories := make([]string, 0, len(payload))
	for _, category := range payload {
		// Solo mostrar categorías raíz (level 0 o sin parent_slug) en el listado
		if category.ParentSlug != "" {
			continue
		}
		if category.Name != "" {
			categories = append(categories, category.Name)
			continue
		}
		if category.Slug != "" {
			categories = append(categories, category.Slug)
		}
	}

	return categories, nil
}
