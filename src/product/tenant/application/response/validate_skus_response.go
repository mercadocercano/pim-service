package response

// ValidateSKUsResponse respuesta de validación de SKUs
type ValidateSKUsResponse struct {
	ExistingSKUs  []string `json:"existing_skus"`
	AvailableSKUs []string `json:"available_skus"`
}
