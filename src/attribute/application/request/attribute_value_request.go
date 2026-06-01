package request

// CreateAttributeValueRequest petición para crear un valor de atributo
type CreateAttributeValueRequest struct {
	Value     string `json:"value" binding:"required"`
	SortOrder int    `json:"sort_order"`
}

// UpdateAttributeValueRequest petición para actualizar un valor de atributo
type UpdateAttributeValueRequest struct {
	Value     string `json:"value" binding:"required"`
	SortOrder int    `json:"sort_order"`
}
