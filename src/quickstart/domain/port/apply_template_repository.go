package port

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"saas-mt-pim-service/src/shared/infrastructure/database"
)

// TemplateCategory representa una categoría del template (desde JSON)
type TemplateCategory struct {
	MarketplaceID string
	Name          string
	Slug          string
	ParentSlug    string `json:"parent_slug"`
	Level         int    `json:"level"`
}

// CreatedCategory representa una categoría creada en el tenant
type CreatedCategory struct {
	ID            string
	Name          string
	Slug          string
	MarketplaceID string
}

// GlobalProductCandidate representa un producto candidato desde global_products
type GlobalProductCandidate struct {
	Name                  string
	Brand                 string
	Description           sql.NullString
	MarketplaceCategoryID string
	EAN                   sql.NullString
	SkuGlobal             sql.NullString
	ImageURL              string
}

// TemplateBrand represents a brand from the template JSONB
type TemplateBrand struct {
	Name                   string   `json:"name"`
	LogoURL                string   `json:"logo_url,omitempty"`
	SuggestedForCategories []string `json:"suggested_for_categories"`
}

// TemplateProduct represents a curated product from the template JSONB
type TemplateProduct struct {
	Name           string  `json:"name"`
	CategorySlug   string  `json:"category_slug"`
	Brand          string  `json:"brand"`
	PriceReference float64 `json:"price_reference"`
	Unit           string  `json:"unit"`
	SkuPrefix      string  `json:"sku_prefix"`
	ImageURL       string  `json:"image_url,omitempty"`
}

// TemplateAttribute represents an attribute definition from the template JSONB
type TemplateAttribute struct {
	Name                 string   `json:"name"`
	Slug                 string   `json:"slug"`
	Values               []string `json:"values"`
	AppliesToCategories []string `json:"applies_to_categories"`
}

// FullTemplateData contains all curated data from a template JSONB
type FullTemplateData struct {
	Categories []TemplateCategory
	Brands     []TemplateBrand
	Products   []TemplateProduct
	Attributes []TemplateAttribute
}

// ApplyTemplateRepository define las operaciones de persistencia para aplicar templates
type ApplyTemplateRepository interface {
	// LoadTemplateCategories carga desde business_type_templates (usa db, fuera de tx)
	LoadTemplateCategories(ctx context.Context, templateID string) ([]TemplateCategory, []string, map[string]string, error)

	// CreateTenantCategoriesLegacy INSERT desde marketplace_categories
	CreateTenantCategoriesLegacy(ctx context.Context, exec database.Executor, tenantID uuid.UUID, parentID string) (int, []CreatedCategory, error)

	// CreateTenantCategoriesFromTemplate INSERT categorías desde template JSON
	CreateTenantCategoriesFromTemplate(ctx context.Context, exec database.Executor, tenantID uuid.UUID, categories []TemplateCategory) (int, []CreatedCategory, error)

	// GetMarketplaceCategoryIDsBySlug SELECT id FROM marketplace_categories
	GetMarketplaceCategoryIDsBySlug(ctx context.Context, exec database.Executor, slugs []string) ([]string, error)

	// CreateTenantBrandsFromGlobalProducts crea marcas desde global_products
	CreateTenantBrandsFromGlobalProducts(ctx context.Context, exec database.Executor, tenantID uuid.UUID, marketplaceIDs []string, categorySlugs []string, useMarketplaceID bool) (int, []string, error)

	// FindGlobalProduct busca un producto candidato
	FindGlobalProduct(ctx context.Context, exec database.Executor, marketplaceIDs []string, categorySlugs []string, useMarketplaceID bool, brandNames []string) (GlobalProductCandidate, error)

	// EnsureTenantBrand busca o inserta marca
	EnsureTenantBrand(ctx context.Context, exec database.Executor, tenantID uuid.UUID, brandName string) (string, string, error)

	// ResolveTenantCategory busca categoría por slug (o usa mapas en memoria)
	ResolveTenantCategory(ctx context.Context, exec database.Executor, tenantID uuid.UUID, tenantCategoriesByMarketplaceID map[string]CreatedCategory, tenantCategoriesBySlug map[string]CreatedCategory, categorySlugByMarketplaceID map[string]string, marketplaceCategoryID string) (sql.NullString, sql.NullString, error)

	// EnsureTenantProduct busca o inserta producto
	EnsureTenantProduct(ctx context.Context, exec database.Executor, tenantID uuid.UUID, candidate GlobalProductCandidate, categoryID sql.NullString, categoryName sql.NullString, brandID string, brandName string) (string, string, string, bool, error)

	// EnsureDefaultVariant inserta variante default
	EnsureDefaultVariant(ctx context.Context, exec database.Executor, tenantID uuid.UUID, productID string, productName string, productSKU string) (int, error)

	// GlobalProductsHasColumn verifica si global_products tiene una columna
	GlobalProductsHasColumn(ctx context.Context, exec database.Executor, columnName string) bool

	// GlobalProductsHasMarketplaceCategoryID atajo para marketplace_category_id
	GlobalProductsHasMarketplaceCategoryID(ctx context.Context, exec database.Executor) bool

	// LoadFullTemplateData loads categories + brands + products + attributes from template JSONB
	LoadFullTemplateData(ctx context.Context, templateID string) (*FullTemplateData, error)

	// CreateTenantBrandsFromTemplate creates all brands listed in the template
	CreateTenantBrandsFromTemplate(ctx context.Context, exec database.Executor, tenantID uuid.UUID, brands []TemplateBrand) (int, []string, error)

	// CreateTenantProductsFromTemplate creates all products listed in the template
	CreateTenantProductsFromTemplate(ctx context.Context, exec database.Executor, tenantID uuid.UUID, products []TemplateProduct, createdCategories []CreatedCategory, createdBrands []string) (int, int, []string, error)

	// CreateTenantAttributesFromTemplate creates attributes and links them to categories
	CreateTenantAttributesFromTemplate(ctx context.Context, exec database.Executor, tenantID uuid.UUID, attributes []TemplateAttribute, createdCategories []CreatedCategory) (int, int, error)
}
