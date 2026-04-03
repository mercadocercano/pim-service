package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"saas-mt-pim-service/src/quickstart/domain/port"
	"saas-mt-pim-service/src/shared/infrastructure/database"
)

const (
	legacyQuickstartParentID   = "f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00"
	legacyQuickstartTemplateID = "ferreteria-corralon"
)

// legacyTemplateCategories: categorías para templates con slug (no UUID)
var legacyTemplateCategories = map[string][]port.TemplateCategory{
	"ferreteria-corralon": {
		{Name: "Tornilleria", Slug: "tornilleria"},
		{Name: "Herramientas Manuales", Slug: "herramientas-manuales"},
		{Name: "Herramientas Electricas", Slug: "herramientas-electricas"},
		{Name: "Materiales de Construccion", Slug: "materiales-construccion"},
		{Name: "Pinturas", Slug: "pinturas"},
		{Name: "Plomeria y Sanitarios", Slug: "plomeria-sanitarios"},
	},
	"bazar": {
		{Name: "Cocina", Slug: "cocina"},
		{Name: "Bazar", Slug: "bazar"},
		{Name: "Decoracion", Slug: "decoracion"},
		{Name: "Organizacion", Slug: "organizacion"},
	},
	"jugueteria": {
		{Name: "Bebes y Ninos", Slug: "bebes-ninos"},
		{Name: "Juguetes", Slug: "juguetes"},
	},
	"ropa": {
		{Name: "Remeras", Slug: "remeras"},
		{Name: "Pantalones", Slug: "pantalones"},
		{Name: "Buzos", Slug: "buzos"},
		{Name: "Camperas", Slug: "camperas"},
		{Name: "Zapatillas", Slug: "zapatillas"},
	},
	"electricidad": {
		{Name: "Electricidad", Slug: "electricidad"},
		{Name: "Iluminacion", Slug: "iluminacion"},
	},
	"zapateria": {
		{Name: "Calzado", Slug: "calzado"},
		{Name: "Escarpines", Slug: "escarpines"},
		{Name: "Calzado Deportivo", Slug: "calzado-deportivo"},
	},
	"deportes": {
		{Name: "Ropa Deportiva", Slug: "ropa-deportiva"},
		{Name: "Indumentaria Deportiva", Slug: "indumentaria-deportiva"},
		{Name: "Deportes", Slug: "deportes"},
		{Name: "Accesorios Deportivos", Slug: "accesorios-deportivos"},
	},
}

// ApplyTemplatePostgresRepository implementa ApplyTemplateRepository para PostgreSQL
type ApplyTemplatePostgresRepository struct {
	db *sql.DB
}

// NewApplyTemplatePostgresRepository crea una nueva instancia del repositorio
func NewApplyTemplatePostgresRepository(db *sql.DB) port.ApplyTemplateRepository {
	return &ApplyTemplatePostgresRepository{db: db}
}

// LoadTemplateCategories carga categorías desde business_type_templates (usa db, fuera de tx)
// Para templates legacy con slug (bazar, jugueteria, etc.) retorna categorías estáticas.
// Para ferreteria-corralon legacy retorna nil (usa CreateTenantCategoriesLegacy).
func (r *ApplyTemplatePostgresRepository) LoadTemplateCategories(ctx context.Context, templateID string) ([]port.TemplateCategory, []string, map[string]string, error) {
	if templateID == legacyQuickstartTemplateID {
		return nil, nil, nil, nil
	}
	if cats, ok := legacyTemplateCategories[templateID]; ok {
		slugByID := make(map[string]string)
		for i, c := range cats {
			slugByID[fmt.Sprintf("legacy-%d", i)] = c.Slug
		}
		return cats, nil, slugByID, nil
	}
	if _, err := uuid.Parse(templateID); err != nil {
		return nil, nil, nil, nil
	}

	var categoriesRaw []byte
	err := r.db.QueryRowContext(ctx, `
		SELECT categories
		FROM business_type_templates
		WHERE id = $1 AND is_active = true
	`, templateID).Scan(&categoriesRaw)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, nil, nil
		}
		return nil, nil, nil, fmt.Errorf("failed to load template categories: %w", err)
	}

	type categoryPayload struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Slug       string `json:"slug"`
		ParentSlug string `json:"parent_slug"`
		Level      int    `json:"level"`
	}

	var payload []categoryPayload
	if err := json.Unmarshal(categoriesRaw, &payload); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to parse template categories: %w", err)
	}

	categories := make([]port.TemplateCategory, 0, len(payload))
	marketplaceIDs := make([]string, 0, len(payload))
	slugByMarketplaceID := make(map[string]string, len(payload))

	for _, item := range payload {
		name := strings.TrimSpace(item.Name)
		slug := strings.TrimSpace(item.Slug)
		marketplaceID := strings.TrimSpace(item.ID)

		if name == "" && slug == "" {
			continue
		}
		if slug == "" {
			slug = buildSlug(name)
		}

		categories = append(categories, port.TemplateCategory{
			MarketplaceID: marketplaceID,
			Name:          name,
			Slug:          slug,
			ParentSlug:    strings.TrimSpace(item.ParentSlug),
			Level:         item.Level,
		})

		if marketplaceID != "" {
			marketplaceIDs = append(marketplaceIDs, marketplaceID)
			if slug != "" {
				slugByMarketplaceID[marketplaceID] = slug
			}
		}
	}

	return categories, marketplaceIDs, slugByMarketplaceID, nil
}

// CreateTenantCategoriesLegacy INSERT desde marketplace_categories
func (r *ApplyTemplatePostgresRepository) CreateTenantCategoriesLegacy(ctx context.Context, exec database.Executor, tenantID uuid.UUID, parentID string) (int, []port.CreatedCategory, error) {
	if parentID == "" {
		return 0, nil, fmt.Errorf("template parent id is empty")
	}

	query := `
		INSERT INTO categories (id, tenant_id, name, slug, description, parent_id, status, created_at, updated_at)
		SELECT 
			gen_random_uuid(),
			$1,
			name,
			slug,
			description,
			NULL,
			'active',
			$2,
			$2
		FROM marketplace_categories
		WHERE parent_id = $3
		  AND is_active = true
		RETURNING id, name, slug
	`

	now := time.Now()
	rows, err := exec.QueryContext(ctx, query, tenantID.String(), now, parentID)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to insert categories: %w", err)
	}
	defer rows.Close()

	var categories []port.CreatedCategory
	count := 0

	for rows.Next() {
		var id, name, slug string
		if err := rows.Scan(&id, &name, &slug); err != nil {
			return 0, nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, port.CreatedCategory{
			ID:            id,
			Name:          name,
			Slug:          slug,
			MarketplaceID: "",
		})
		count++
	}

	if err := rows.Err(); err != nil {
		return 0, nil, fmt.Errorf("error iterating categories: %w", err)
	}

	return count, categories, nil
}

// CreateTenantCategoriesFromTemplate INSERT categorías desde template JSON
// Soporta jerarquía: primero crea categorías raíz (level=0 o sin parent_slug),
// luego crea hijas resolviendo parent_id desde el slug del padre.
func (r *ApplyTemplatePostgresRepository) CreateTenantCategoriesFromTemplate(ctx context.Context, exec database.Executor, tenantID uuid.UUID, categories []port.TemplateCategory) (int, []port.CreatedCategory, error) {
	if len(categories) == 0 {
		return 0, nil, nil
	}

	now := time.Now()
	created := 0
	var createdCategories []port.CreatedCategory
	// Map slug -> created ID for parent resolution
	slugToID := make(map[string]string)

	// Separate root categories (no parent) from children
	var roots, children []port.TemplateCategory
	for _, cat := range categories {
		if cat.ParentSlug == "" {
			roots = append(roots, cat)
		} else {
			children = append(children, cat)
		}
	}

	// Pass 1: Create root categories
	for _, category := range roots {
		id, name, slug, isNew, err := r.upsertCategory(ctx, exec, tenantID, category, "", now)
		if err != nil {
			return created, createdCategories, err
		}
		if isNew {
			created++
		}
		slugToID[slug] = id
		createdCategories = append(createdCategories, port.CreatedCategory{
			ID:            id,
			Name:          name,
			Slug:          slug,
			MarketplaceID: category.MarketplaceID,
		})
	}

	// Pass 2: Create child categories with resolved parent_id
	for _, category := range children {
		parentID := slugToID[category.ParentSlug]
		id, name, slug, isNew, err := r.upsertCategory(ctx, exec, tenantID, category, parentID, now)
		if err != nil {
			return created, createdCategories, err
		}
		if isNew {
			created++
		}
		slugToID[slug] = id
		createdCategories = append(createdCategories, port.CreatedCategory{
			ID:            id,
			Name:          name,
			Slug:          slug,
			MarketplaceID: category.MarketplaceID,
		})
	}

	return created, createdCategories, nil
}

// upsertCategory busca o crea una categoría individual, retornando (id, name, slug, isNew, error)
func (r *ApplyTemplatePostgresRepository) upsertCategory(ctx context.Context, exec database.Executor, tenantID uuid.UUID, category port.TemplateCategory, parentID string, now time.Time) (string, string, string, bool, error) {
	name := strings.TrimSpace(category.Name)
	slug := strings.TrimSpace(category.Slug)

	if name == "" && slug == "" {
		return "", "", "", false, nil
	}
	if slug == "" {
		slug = buildSlug(name)
	}
	if name == "" {
		name = slug
	}

	var id string
	err := exec.QueryRowContext(ctx, `
		SELECT id
		FROM categories
		WHERE tenant_id = $1 AND slug = $2 AND status = 'active'
		LIMIT 1
	`, tenantID.String(), slug).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return "", "", "", false, fmt.Errorf("failed to query category: %w", err)
	}

	if err == sql.ErrNoRows {
		var parentVal interface{}
		if parentID != "" {
			parentVal = parentID
		}
		err = exec.QueryRowContext(ctx, `
			INSERT INTO categories (id, tenant_id, name, slug, description, parent_id, status, created_at, updated_at)
			VALUES (gen_random_uuid(), $1, $2, $3, '', $4, 'active', $5, $5)
			RETURNING id
		`, tenantID.String(), name, slug, parentVal, now).Scan(&id)
		if err != nil {
			return "", "", "", false, fmt.Errorf("failed to insert category: %w", err)
		}
		return id, name, slug, true, nil
	}

	return id, name, slug, false, nil
}

// GetMarketplaceCategoryIDsBySlug SELECT id FROM marketplace_categories
func (r *ApplyTemplatePostgresRepository) GetMarketplaceCategoryIDsBySlug(ctx context.Context, exec database.Executor, slugs []string) ([]string, error) {
	if len(slugs) == 0 {
		return nil, nil
	}

	query := `
		SELECT id
		FROM marketplace_categories
		WHERE slug = ANY($1)
		  AND is_active = true
	`

	rows, err := exec.QueryContext(ctx, query, pq.Array(slugs))
	if err != nil {
		return nil, fmt.Errorf("failed to query marketplace categories: %w", err)
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan marketplace category id: %w", err)
		}
		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating marketplace categories: %w", err)
	}

	return ids, nil
}

// CreateTenantBrandsFromGlobalProducts crea marcas desde global_products
func (r *ApplyTemplatePostgresRepository) CreateTenantBrandsFromGlobalProducts(ctx context.Context, exec database.Executor, tenantID uuid.UUID, marketplaceIDs []string, categorySlugs []string, useMarketplaceID bool) (int, []string, error) {
	if useMarketplaceID && len(marketplaceIDs) == 0 {
		return 0, nil, nil
	}
	if !useMarketplaceID && len(categorySlugs) == 0 {
		return 0, nil, nil
	}

	query := `
		SELECT DISTINCT brand
		FROM global_products
		WHERE is_active = true
		  AND brand IS NOT NULL
		  AND TRIM(brand) != ''
	`

	var rows *sql.Rows
	var err error
	if useMarketplaceID {
		query += " AND marketplace_category_id = ANY($1) LIMIT 3"
		rows, err = exec.QueryContext(ctx, query, pq.Array(marketplaceIDs))
	} else {
		query += " AND category = ANY($1) LIMIT 3"
		rows, err = exec.QueryContext(ctx, query, pq.Array(categorySlugs))
	}
	if err != nil {
		return 0, nil, fmt.Errorf("failed to query global brands: %w", err)
	}
	defer rows.Close()

	now := time.Now()
	created := 0
	var createdBrands []string

	for rows.Next() {
		var brandName string
		if err := rows.Scan(&brandName); err != nil {
			return 0, nil, fmt.Errorf("failed to scan brand: %w", err)
		}

		if brandName == "" {
			continue
		}

		result, err := exec.ExecContext(ctx, `
			INSERT INTO brands (id, tenant_id, name, description, status, created_at, updated_at)
			VALUES ($1, $2, $3, '', 'active', $4, $4)
			ON CONFLICT (tenant_id, name) DO NOTHING
		`, uuid.New().String(), tenantID.String(), brandName, now)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to insert brand: %w", err)
		}

		if rowsAffected, err := result.RowsAffected(); err == nil && rowsAffected > 0 {
			created++
			createdBrands = append(createdBrands, brandName)
		}
	}

	if err := rows.Err(); err != nil {
		return 0, nil, fmt.Errorf("error iterating brands: %w", err)
	}

	return created, createdBrands, nil
}

// FindGlobalProduct busca un producto candidato
func (r *ApplyTemplatePostgresRepository) FindGlobalProduct(ctx context.Context, exec database.Executor, marketplaceIDs []string, categorySlugs []string, useMarketplaceID bool, brandNames []string) (port.GlobalProductCandidate, error) {
	var candidate port.GlobalProductCandidate
	hasSkuGlobal := r.GlobalProductsHasColumn(ctx, exec, "sku_global")
	hasEAN := r.GlobalProductsHasColumn(ctx, exec, "ean")

	query := `
		SELECT name, brand, description
	`
	if hasEAN {
		query += ", ean"
	}
	if hasSkuGlobal {
		query += ", sku_global"
	}
	query += `
		FROM global_products
		WHERE is_active = true
	`

	args := make([]interface{}, 0, 2)
	if useMarketplaceID {
		query += " AND marketplace_category_id = ANY($1)"
		args = append(args, pq.Array(marketplaceIDs))
	} else {
		query += " AND category = ANY($1)"
		args = append(args, pq.Array(categorySlugs))
	}

	if len(brandNames) > 0 {
		query += " AND brand = ANY($2)"
		args = append(args, pq.Array(brandNames))
	}

	if r.GlobalProductsHasColumn(ctx, exec, "popularity_rank") {
		query += " ORDER BY COALESCE(quality_score, 0) DESC, popularity_rank DESC LIMIT 1"
	} else {
		query += " ORDER BY COALESCE(quality_score, 0) DESC LIMIT 1"
	}

	row := exec.QueryRowContext(ctx, query, args...)
	scanTargets := []interface{}{&candidate.Name, &candidate.Brand, &candidate.Description}
	if hasEAN {
		scanTargets = append(scanTargets, &candidate.EAN)
	}
	if hasSkuGlobal {
		scanTargets = append(scanTargets, &candidate.SkuGlobal)
	}
	err := row.Scan(scanTargets...)
	if err != nil {
		return candidate, err
	}

	if useMarketplaceID && len(marketplaceIDs) > 0 {
		candidate.MarketplaceCategoryID = marketplaceIDs[0]
	}

	return candidate, nil
}

// EnsureTenantBrand busca o inserta marca
func (r *ApplyTemplatePostgresRepository) EnsureTenantBrand(ctx context.Context, exec database.Executor, tenantID uuid.UUID, brandName string) (string, string, error) {
	if brandName == "" {
		return "", "", nil
	}

	var existingID string
	err := exec.QueryRowContext(ctx, `
		SELECT id FROM brands
		WHERE tenant_id = $1 AND name = $2 AND status = 'active'
		LIMIT 1
	`, tenantID.String(), brandName).Scan(&existingID)
	if err == nil {
		return existingID, brandName, nil
	}
	if err != sql.ErrNoRows {
		return "", "", fmt.Errorf("failed to query brand: %w", err)
	}

	newID := uuid.New().String()
	now := time.Now()
	_, err = exec.ExecContext(ctx, `
		INSERT INTO brands (id, tenant_id, name, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, '', 'active', $4, $4)
		ON CONFLICT (tenant_id, name) DO NOTHING
	`, newID, tenantID.String(), brandName, now)
	if err != nil {
		return "", "", fmt.Errorf("failed to insert brand: %w", err)
	}

	return newID, brandName, nil
}

// ResolveTenantCategory busca categoría por slug (o usa mapas en memoria)
func (r *ApplyTemplatePostgresRepository) ResolveTenantCategory(ctx context.Context, exec database.Executor, tenantID uuid.UUID, tenantCategoriesByMarketplaceID map[string]port.CreatedCategory, tenantCategoriesBySlug map[string]port.CreatedCategory, categorySlugByMarketplaceID map[string]string, marketplaceCategoryID string) (sql.NullString, sql.NullString, error) {
	if category, ok := tenantCategoriesByMarketplaceID[marketplaceCategoryID]; ok {
		return sql.NullString{String: category.ID, Valid: true}, sql.NullString{String: category.Name, Valid: true}, nil
	}

	categorySlug := categorySlugByMarketplaceID[marketplaceCategoryID]
	if categorySlug != "" {
		if category, ok := tenantCategoriesBySlug[categorySlug]; ok {
			return sql.NullString{String: category.ID, Valid: true}, sql.NullString{String: category.Name, Valid: true}, nil
		}
	}

	var categoryID sql.NullString
	var categoryName sql.NullString
	err := exec.QueryRowContext(ctx, `
		SELECT id, name FROM categories
		WHERE tenant_id = $1 AND slug = $2 AND status = 'active'
		LIMIT 1
	`, tenantID.String(), categorySlug).Scan(&categoryID, &categoryName)
	if err != nil && err != sql.ErrNoRows {
		return sql.NullString{}, sql.NullString{}, fmt.Errorf("failed to query tenant category: %w", err)
	}

	return categoryID, categoryName, nil
}

// EnsureTenantProduct busca o inserta producto
func (r *ApplyTemplatePostgresRepository) EnsureTenantProduct(ctx context.Context, exec database.Executor, tenantID uuid.UUID, candidate port.GlobalProductCandidate, categoryID sql.NullString, categoryName sql.NullString, brandID string, brandName string) (string, string, string, bool, error) {
	var existingID string
	err := exec.QueryRowContext(ctx, `
		SELECT id FROM products
		WHERE tenant_id = $1 AND name = $2 AND status != 'deleted'
		LIMIT 1
	`, tenantID.String(), candidate.Name).Scan(&existingID)
	if err == nil {
		return existingID, candidate.Name, "", false, nil
	}
	if err != sql.ErrNoRows {
		return "", "", "", false, fmt.Errorf("failed to query product: %w", err)
	}

	productID := uuid.New().String()
	productSKU := pickProductSKU(candidate)
	now := time.Now()
	_, err = exec.ExecContext(ctx, `
		INSERT INTO products (
			id, tenant_id, name, description, sku, category_id, category_name,
			brand_id, brand_name, status, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, 'active', $10, $10)
	`, productID, tenantID.String(), candidate.Name, candidate.Description.String, productSKU, categoryID, categoryName, nullableUUID(brandID), sql.NullString{String: brandName, Valid: brandName != ""}, now)
	if err != nil {
		return "", "", "", false, fmt.Errorf("failed to insert product: %w", err)
	}

	return productID, candidate.Name, productSKU, true, nil
}

// EnsureDefaultVariant inserta variante default
func (r *ApplyTemplatePostgresRepository) EnsureDefaultVariant(ctx context.Context, exec database.Executor, tenantID uuid.UUID, productID string, productName string, productSKU string) (int, error) {
	var existingID string
	err := exec.QueryRowContext(ctx, `
		SELECT id FROM product_variants
		WHERE product_id = $1 AND is_default = true AND status != 'deleted'
		LIMIT 1
	`, productID).Scan(&existingID)
	if err == nil {
		return 0, nil
	}
	if err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to query variant: %w", err)
	}

	variantSKU := productSKU
	if variantSKU == "" {
		variantSKU = "VAR-" + uuid.New().String()[:8]
	} else {
		variantSKU = fmt.Sprintf("%s-DEF", productSKU)
	}

	now := time.Now()
	_, err = exec.ExecContext(ctx, `
		INSERT INTO product_variants (
			id, tenant_id, product_id, name, sku, status, is_default, sort_order, price, stock, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, 'active', true, 0, 0, 0, $6, $6)
	`, uuid.New().String(), tenantID.String(), productID, productName, variantSKU, now)
	if err != nil {
		return 0, fmt.Errorf("failed to insert variant: %w", err)
	}

	return 1, nil
}

// GlobalProductsHasColumn verifica si global_products tiene una columna (esquema public)
func (r *ApplyTemplatePostgresRepository) GlobalProductsHasColumn(ctx context.Context, exec database.Executor, columnName string) bool {
	var exists bool
	err := exec.QueryRowContext(ctx, `
		SELECT EXISTS(
			SELECT 1
			FROM information_schema.columns
			WHERE table_schema = 'public'
			  AND table_name = 'global_products'
			  AND column_name = $1
		)
	`, columnName).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// GlobalProductsHasMarketplaceCategoryID atajo para marketplace_category_id
func (r *ApplyTemplatePostgresRepository) GlobalProductsHasMarketplaceCategoryID(ctx context.Context, exec database.Executor) bool {
	return r.GlobalProductsHasColumn(ctx, exec, "marketplace_category_id")
}

func buildSlug(value string) string {
	normalized := strings.ToLower(strings.TrimSpace(value))
	if normalized == "" {
		return "category-" + uuid.New().String()[:8]
	}
	nonAlnum := regexp.MustCompile(`[^a-z0-9]+`)
	normalized = nonAlnum.ReplaceAllString(normalized, "-")
	normalized = strings.Trim(normalized, "-")
	if normalized == "" {
		return "category-" + uuid.New().String()[:8]
	}
	return normalized
}

func pickProductSKU(candidate port.GlobalProductCandidate) string {
	if candidate.EAN.Valid && candidate.EAN.String != "" {
		return candidate.EAN.String
	}
	if candidate.SkuGlobal.Valid && candidate.SkuGlobal.String != "" {
		return candidate.SkuGlobal.String
	}
	return "PRD-" + uuid.New().String()[:8]
}

func nullableUUID(value string) sql.NullString {
	if value == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: value, Valid: true}
}

// LoadFullTemplateData loads all JSONB fields from business_type_templates
func (r *ApplyTemplatePostgresRepository) LoadFullTemplateData(ctx context.Context, templateID string) (*port.FullTemplateData, error) {
	if _, err := uuid.Parse(templateID); err != nil {
		return nil, nil
	}

	var categoriesRaw, brandsRaw, productsRaw, attributesRaw []byte
	err := r.db.QueryRowContext(ctx, `
		SELECT 
			COALESCE(categories, '[]'::jsonb),
			COALESCE(brands, '[]'::jsonb),
			COALESCE(products, '[]'::jsonb),
			COALESCE(attributes, '[]'::jsonb)
		FROM business_type_templates
		WHERE id = $1 AND is_active = true
	`, templateID).Scan(&categoriesRaw, &brandsRaw, &productsRaw, &attributesRaw)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to load full template data: %w", err)
	}

	data := &port.FullTemplateData{}

	type catPayload struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Slug       string `json:"slug"`
		ParentSlug string `json:"parent_slug"`
		Level      int    `json:"level"`
	}
	var cats []catPayload
	if err := json.Unmarshal(categoriesRaw, &cats); err == nil {
		for _, c := range cats {
			data.Categories = append(data.Categories, port.TemplateCategory{
				MarketplaceID: c.ID,
				Name:          c.Name,
				Slug:          c.Slug,
				ParentSlug:    c.ParentSlug,
				Level:         c.Level,
			})
		}
	}

	json.Unmarshal(brandsRaw, &data.Brands)
	json.Unmarshal(productsRaw, &data.Products)
	json.Unmarshal(attributesRaw, &data.Attributes)

	return data, nil
}

// CreateTenantBrandsFromTemplate creates all brands listed in the template
func (r *ApplyTemplatePostgresRepository) CreateTenantBrandsFromTemplate(ctx context.Context, exec database.Executor, tenantID uuid.UUID, brands []port.TemplateBrand) (int, []string, error) {
	if len(brands) == 0 {
		return 0, nil, nil
	}

	created := 0
	var createdNames []string
	now := time.Now()

	for _, b := range brands {
		name := strings.TrimSpace(b.Name)
		if name == "" {
			continue
		}

		result, err := exec.ExecContext(ctx, `
			INSERT INTO brands (id, tenant_id, name, description, status, created_at, updated_at)
			VALUES ($1, $2, $3, '', 'active', $4, $4)
			ON CONFLICT (tenant_id, name) DO NOTHING
		`, uuid.New().String(), tenantID.String(), name, now)
		if err != nil {
			return created, createdNames, fmt.Errorf("failed to insert brand %s: %w", name, err)
		}

		if rowsAffected, err := result.RowsAffected(); err == nil && rowsAffected > 0 {
			created++
			createdNames = append(createdNames, name)
		}
	}

	return created, createdNames, nil
}

// CreateTenantProductsFromTemplate creates all curated products from the template
func (r *ApplyTemplatePostgresRepository) CreateTenantProductsFromTemplate(ctx context.Context, exec database.Executor, tenantID uuid.UUID, products []port.TemplateProduct, createdCategories []port.CreatedCategory, createdBrands []string) (int, int, []string, error) {
	if len(products) == 0 {
		return 0, 0, nil, nil
	}

	catBySlug := make(map[string]port.CreatedCategory)
	for _, c := range createdCategories {
		catBySlug[c.Slug] = c
	}

	now := time.Now()
	productsCreated := 0
	variantsCreated := 0
	var productNames []string

	for i, p := range products {
		name := strings.TrimSpace(p.Name)
		if name == "" {
			continue
		}

		var categoryID, categoryName sql.NullString
		if cat, ok := catBySlug[p.CategorySlug]; ok {
			categoryID = sql.NullString{String: cat.ID, Valid: true}
			categoryName = sql.NullString{String: cat.Name, Valid: true}
		}

		var brandID, brandName sql.NullString
		if p.Brand != "" {
			brandName = sql.NullString{String: p.Brand, Valid: true}
			var foundBrandID string
			err := exec.QueryRowContext(ctx, `
				SELECT id FROM brands WHERE tenant_id = $1 AND name = $2 AND status = 'active' LIMIT 1
			`, tenantID.String(), p.Brand).Scan(&foundBrandID)
			if err == nil {
				brandID = sql.NullString{String: foundBrandID, Valid: true}
			} else if err != sql.ErrNoRows {
				return productsCreated, variantsCreated, productNames, fmt.Errorf("failed to lookup brand %s: %w", p.Brand, err)
			}
		}

		sku := fmt.Sprintf("%s-%03d", p.SkuPrefix, i+1)

		var existingID string
		err := exec.QueryRowContext(ctx, `
			SELECT id FROM products WHERE tenant_id = $1 AND name = $2 AND status != 'deleted' LIMIT 1
		`, tenantID.String(), name).Scan(&existingID)

		if err == sql.ErrNoRows {
			productID := uuid.New().String()
			_, err = exec.ExecContext(ctx, `
				INSERT INTO products (id, tenant_id, name, description, sku, category_id, category_name, brand_id, brand_name, status, created_at, updated_at)
				VALUES ($1, $2, $3, '', $4, $5, $6, $7, $8, 'active', $9, $9)
			`, productID, tenantID.String(), name, sku, categoryID, categoryName, brandID, brandName, now)
			if err != nil {
				return productsCreated, variantsCreated, productNames, fmt.Errorf("failed to insert product %s: %w", name, err)
			}
			productsCreated++
			productNames = append(productNames, name)

			variantSKU := sku + "-DEF"
			_, err = exec.ExecContext(ctx, `
				INSERT INTO product_variants (id, tenant_id, product_id, name, sku, status, is_default, sort_order, price, stock, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, 'active', true, 0, $6, 0, $7, $7)
			`, uuid.New().String(), tenantID.String(), productID, name, variantSKU, p.PriceReference, now)
			if err != nil {
				return productsCreated, variantsCreated, productNames, fmt.Errorf("failed to insert default variant for product %s: %w", name, err)
			}
			variantsCreated++
		} else if err != nil {
			return productsCreated, variantsCreated, productNames, fmt.Errorf("failed to check existing product %s: %w", name, err)
		}
	}

	return productsCreated, variantsCreated, productNames, nil
}

// CreateTenantAttributesFromTemplate creates attributes and links them to categories
func (r *ApplyTemplatePostgresRepository) CreateTenantAttributesFromTemplate(ctx context.Context, exec database.Executor, tenantID uuid.UUID, attributes []port.TemplateAttribute, createdCategories []port.CreatedCategory) (int, int, error) {
	if len(attributes) == 0 {
		return 0, 0, nil
	}

	catBySlug := make(map[string]port.CreatedCategory)
	for _, c := range createdCategories {
		catBySlug[c.Slug] = c
	}

	now := time.Now()
	attributesCreated := 0
	linksCreated := 0

	for _, attr := range attributes {
		name := strings.TrimSpace(attr.Name)
		if name == "" {
			continue
		}

		var attrID string
		err := exec.QueryRowContext(ctx, `
			SELECT id FROM attributes WHERE tenant_id = $1 AND name = $2 AND status = 'active' LIMIT 1
		`, tenantID.String(), name).Scan(&attrID)

		if err == sql.ErrNoRows {
			attrID = uuid.New().String()
			_, err = exec.ExecContext(ctx, `
				INSERT INTO attributes (id, tenant_id, name, description, type, required, options, status, created_at, updated_at)
				VALUES ($1, $2, $3, '', 'select', false, $4, 'active', $5, $5)
			`, attrID, tenantID.String(), name, pq.Array(attr.Values), now)
			if err != nil {
				return attributesCreated, linksCreated, fmt.Errorf("failed to insert attribute %s: %w", name, err)
			}
			attributesCreated++
		} else if err != nil {
			return attributesCreated, linksCreated, fmt.Errorf("failed to check existing attribute %s: %w", name, err)
		}

		for _, catSlug := range attr.AppliesToCategories {
			cat, ok := catBySlug[catSlug]
			if !ok {
				continue
			}
			_, err = exec.ExecContext(ctx, `
				INSERT INTO category_attributes (id, tenant_id, category_id, attribute_id, status, created_at, updated_at)
				VALUES ($1, $2, $3, $4, 'active', $5, $5)
				ON CONFLICT (tenant_id, category_id, attribute_id) DO NOTHING
			`, uuid.New().String(), tenantID.String(), cat.ID, attrID, now)
			if err != nil {
				return attributesCreated, linksCreated, fmt.Errorf("failed to link attribute %s to category %s: %w", name, catSlug, err)
			}
			linksCreated++
		}
	}

	return attributesCreated, linksCreated, nil
}
