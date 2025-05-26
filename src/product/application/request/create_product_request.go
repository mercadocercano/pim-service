package request

import "errors"

// CreateProductRequest representa la petición para crear un producto
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"iPhone 15 Pro"`
	Description *string `json:"description,omitempty" example:"Smartphone Apple con chip A17 Pro"`
	SKU         *string `json:"sku,omitempty" example:"IPHONE-15-PRO-256GB"`
	CategoryID  *string `json:"category_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440001"`
	BrandID     *string `json:"brand_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440006"`
}

// Validate valida los datos de la petición
func (r *CreateProductRequest) Validate() error {
	if r.Name == "" {
		return errors.New("el nombre del producto es obligatorio")
	}

	if len(r.Name) < 2 {
		return errors.New("el nombre del producto debe tener al menos 2 caracteres")
	}

	if len(r.Name) > 255 {
		return errors.New("el nombre del producto no puede exceder 255 caracteres")
	}

	if r.Description != nil && len(*r.Description) > 1000 {
		return errors.New("la descripción no puede exceder 1000 caracteres")
	}

	if r.SKU != nil && *r.SKU != "" {
		if len(*r.SKU) < 3 {
			return errors.New("el SKU debe tener al menos 3 caracteres")
		}
		if len(*r.SKU) > 50 {
			return errors.New("el SKU no puede exceder 50 caracteres")
		}
	}

	return nil
}
