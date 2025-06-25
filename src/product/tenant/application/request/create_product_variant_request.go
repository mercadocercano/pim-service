package request

import (
	"github.com/google/uuid"
)

// CreateProductVariantRequest representa la solicitud para crear una nueva variante de producto
type CreateProductVariantRequest struct {
	ProductID  uuid.UUID                       `json:"product_id" binding:"required"`
	Name       string                          `json:"name" binding:"required,min=2,max=255"`
	SKU        *string                         `json:"sku,omitempty" binding:"omitempty,min=3,max=50"`
	IsDefault  *bool                           `json:"is_default,omitempty"`
	SortOrder  *int                            `json:"sort_order,omitempty" binding:"omitempty,min=0"`
	Attributes []CreateVariantAttributeRequest `json:"attributes,omitempty"`
}

// CreateVariantAttributeRequest representa un atributo de variante en la solicitud
type CreateVariantAttributeRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=100"`
	Value string `json:"value" binding:"required,min=1,max=255"`
}

// GetSKU devuelve el SKU si está presente
func (r *CreateProductVariantRequest) GetSKU() string {
	if r.SKU == nil {
		return ""
	}
	return *r.SKU
}

// GetIsDefault devuelve si es por defecto (false si no se especifica)
func (r *CreateProductVariantRequest) GetIsDefault() bool {
	if r.IsDefault == nil {
		return false
	}
	return *r.IsDefault
}

// GetSortOrder devuelve el orden de clasificación (0 si no se especifica)
func (r *CreateProductVariantRequest) GetSortOrder() int {
	if r.SortOrder == nil {
		return 0
	}
	return *r.SortOrder
}

// HasSKU verifica si se proporcionó un SKU
func (r *CreateProductVariantRequest) HasSKU() bool {
	return r.SKU != nil && *r.SKU != ""
}

// HasAttributes verifica si se proporcionaron atributos
func (r *CreateProductVariantRequest) HasAttributes() bool {
	return len(r.Attributes) > 0
}
