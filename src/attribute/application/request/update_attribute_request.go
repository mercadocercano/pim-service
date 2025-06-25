package request

// UpdateAttributeRequest representa la petición para actualizar un atributo básico
type UpdateAttributeRequest struct {
	Name string `json:"name" binding:"required" example:"color"`
}
