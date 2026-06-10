package request

// ValidateSKUsRequest solicitud para validar SKUs existentes
type ValidateSKUsRequest struct {
	SKUs []string `json:"skus" binding:"required,min=1"`
}
