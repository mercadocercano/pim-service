package request

// SetupTenantRequest representa la solicitud para configurar un tenant con quickstart
type SetupTenantRequest struct {
	BusinessType       string   `json:"businessType" binding:"required"`
	SelectedCategories []string `json:"selectedCategories" binding:"required,min=1"`
	SelectedAttributes []string `json:"selectedAttributes"`
	SelectedVariants   []string `json:"selectedVariants"`
	SelectedProducts   []string `json:"selectedProducts"`
}

// ToMap convierte el request a un mapa para ser procesado por el caso de uso
func (r *SetupTenantRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"businessType":       r.BusinessType,
		"selectedCategories": r.SelectedCategories,
		"selectedAttributes": r.SelectedAttributes,
		"selectedVariants":   r.SelectedVariants,
		"selectedProducts":   r.SelectedProducts,
	}
}
