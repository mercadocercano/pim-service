package request

// CreateCategoryAttributeRequest representa la solicitud para crear un nuevo atributo de categoría
type CreateCategoryAttributeRequest struct {
	CategoryID    string   `json:"category_id" binding:"required"`
	AttributeID   string   `json:"attribute_id" binding:"required"`
	AllowedValues []string `json:"allowed_values"`
}
