package response

import "time"

// TenantTaxonomyResponse representa la respuesta con la taxonomía completa de un tenant
type TenantTaxonomyResponse struct {
	TenantID              string                      `json:"tenant_id"`
	Format                string                      `json:"format"`
	TotalCategories       int                         `json:"total_categories"`
	TotalMappings         int                         `json:"total_mappings"`
	TotalCustomAttributes int                         `json:"total_custom_attributes"`
	Categories            []TenantCategoryNode        `json:"categories"`
	CustomAttributes      []TenantCustomAttributeInfo `json:"custom_attributes,omitempty"`
	Metadata              TaxonomyMetadata            `json:"metadata"`
	GeneratedAt           time.Time                   `json:"generated_at"`
}

// TenantCategoryNode representa un nodo de categoría en la taxonomía del tenant
type TenantCategoryNode struct {
	// Datos de la categoría tenant
	CategoryID string  `json:"category_id"`
	Name       string  `json:"name"`
	Level      int     `json:"level"`
	ParentID   *string `json:"parent_id"`

	// Datos del mapeo marketplace
	MappingID             *string `json:"mapping_id,omitempty"`
	MarketplaceCategoryID *string `json:"marketplace_category_id,omitempty"`
	CustomName            *string `json:"custom_name,omitempty"`

	// Datos de la categoría marketplace (si se incluyen)
	MarketplaceData *MarketplaceCategoryInfo `json:"marketplace_data,omitempty"`

	// Atributos personalizados de esta categoría
	CustomAttributes []TenantCustomAttributeInfo `json:"custom_attributes,omitempty"`

	// Hijos (solo para formato tree)
	Children []TenantCategoryNode `json:"children,omitempty"`

	// Metadatos
	IsActive   bool      `json:"is_active"`
	HasMapping bool      `json:"has_mapping"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MarketplaceCategoryInfo representa información de la categoría marketplace
type MarketplaceCategoryInfo struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description *string   `json:"description"`
	Level       int       `json:"level"`
	ParentID    *string   `json:"parent_id"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TenantCustomAttributeInfo representa información de un atributo personalizado
type TenantCustomAttributeInfo struct {
	ID              string                 `json:"id"`
	CategoryID      string                 `json:"category_id"`
	Name            string                 `json:"name"`
	Slug            string                 `json:"slug"`
	Type            string                 `json:"type"`
	IsRequired      bool                   `json:"is_required"`
	DefaultValue    *string                `json:"default_value"`
	ValidationRules map[string]interface{} `json:"validation_rules,omitempty"`
	Options         []string               `json:"options,omitempty"` // para select/multi_select
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

// TaxonomyMetadata representa metadatos de la taxonomía
type TaxonomyMetadata struct {
	MaxDepth                int                    `json:"max_depth"`
	RootCategoriesCount     int                    `json:"root_categories_count"`
	MappedCategoriesCount   int                    `json:"mapped_categories_count"`
	UnmappedCategoriesCount int                    `json:"unmapped_categories_count"`
	CustomAttributesCount   int                    `json:"custom_attributes_count"`
	LastSyncAt              *time.Time             `json:"last_sync_at,omitempty"`
	IncludeOptions          TaxonomyIncludeOptions `json:"include_options"`
}

// TaxonomyIncludeOptions representa las opciones de inclusión utilizadas
type TaxonomyIncludeOptions struct {
	CustomAttributes   bool `json:"custom_attributes"`
	MarketplaceData    bool `json:"marketplace_data"`
	InactiveCategories bool `json:"inactive_categories"`
}

// TenantTaxonomySummary representa un resumen de la taxonomía del tenant
type TenantTaxonomySummary struct {
	TenantID              string    `json:"tenant_id"`
	TotalCategories       int       `json:"total_categories"`
	MappedCategories      int       `json:"mapped_categories"`
	UnmappedCategories    int       `json:"unmapped_categories"`
	CustomAttributesCount int       `json:"custom_attributes_count"`
	MaxCategoryDepth      int       `json:"max_category_depth"`
	LastUpdated           time.Time `json:"last_updated"`
	CompletionPercentage  float64   `json:"completion_percentage"` // % de categorías mapeadas
}
