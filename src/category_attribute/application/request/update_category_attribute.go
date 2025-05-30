package request

// UpdateCategoryAttributeRequest representa la solicitud para actualizar un atributo de categoría
type UpdateCategoryAttributeRequest struct {
	AllowedValues []string `json:"allowed_values" binding:"required"`
}
