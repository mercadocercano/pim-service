package request

// ImportProductsCSVRequest request para importar productos desde CSV
type ImportProductsCSVRequest struct {
	// CreateVariants indica si se deben crear variantes por defecto
	CreateVariants bool `form:"create_variants" json:"create_variants"`
	
	// SkipExisting indica si se deben omitir productos con SKU existente
	SkipExisting bool `form:"skip_existing" json:"skip_existing"`
	
	// UpdateExisting indica si se deben actualizar productos existentes
	UpdateExisting bool `form:"update_existing" json:"update_existing"`
}