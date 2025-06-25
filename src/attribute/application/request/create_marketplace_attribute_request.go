package request

// CreateMarketplaceAttributeRequest representa la petición para crear un atributo marketplace
type CreateMarketplaceAttributeRequest struct {
	Name          string   `json:"name" binding:"required" example:"color"`
	Type          string   `json:"type" binding:"required" example:"enum"`
	Description   *string  `json:"description" example:"Color del producto"`
	IsRequired    bool     `json:"is_required" example:"false"`
	IsSearchable  bool     `json:"is_searchable" example:"true"`
	IsFilterable  bool     `json:"is_filterable" example:"true"`
	AllowedValues []string `json:"allowed_values" example:"[\"rojo\", \"azul\", \"verde\"]"`
}
