package request

// CreateAttributeRequest representa la petición para crear un atributo básico
type CreateAttributeRequest struct {
	Name string `json:"name" binding:"required" example:"color"`
}
