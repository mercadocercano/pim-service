package request

// CreateBrandRequest representa la solicitud para crear una nueva marca
type CreateBrandRequest struct {
	Name        string  `json:"name" binding:"required,min=1,max=100" example:"Nike"`
	Description string  `json:"description" binding:"max=500" example:"Marca deportiva internacional"`
	LogoURL     *string `json:"logo_url,omitempty" binding:"omitempty,url" example:"https://example.com/logo.png"`
	Website     *string `json:"website,omitempty" binding:"omitempty,url" example:"https://nike.com"`
}

// Validate realiza validaciones adicionales del request
func (r *CreateBrandRequest) Validate() error {
	// Validaciones adicionales si son necesarias
	return nil
}
