package request

// UpdateProductVariantRequest representa la solicitud para actualizar una variante de producto
type UpdateProductVariantRequest struct {
	Name       *string                         `json:"name,omitempty" binding:"omitempty,min=2,max=255"`
	SKU        *string                         `json:"sku,omitempty" binding:"omitempty,min=3,max=50"`
	Price      *float64                        `json:"price,omitempty" binding:"omitempty,min=0"`
	Stock      *int                            `json:"stock,omitempty" binding:"omitempty,min=0"`
	SortOrder  *int                            `json:"sort_order,omitempty" binding:"omitempty,min=0"`
	Attributes []UpdateVariantAttributeRequest `json:"attributes,omitempty"`
}

// UpdateVariantAttributeRequest representa un atributo de variante en la solicitud de actualización
type UpdateVariantAttributeRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=100"`
	Value string `json:"value" binding:"required,min=1,max=255"`
}

// GetName devuelve el nombre si está presente
func (r *UpdateProductVariantRequest) GetName() string {
	if r.Name == nil {
		return ""
	}
	return *r.Name
}

// GetSKU devuelve el SKU si está presente
func (r *UpdateProductVariantRequest) GetSKU() string {
	if r.SKU == nil {
		return ""
	}
	return *r.SKU
}

// GetSortOrder devuelve el orden de clasificación si está presente
func (r *UpdateProductVariantRequest) GetSortOrder() int {
	if r.SortOrder == nil {
		return 0
	}
	return *r.SortOrder
}

// HasName verifica si se proporcionó un nombre
func (r *UpdateProductVariantRequest) HasName() bool {
	return r.Name != nil && *r.Name != ""
}

// HasSKU verifica si se proporcionó un SKU
func (r *UpdateProductVariantRequest) HasSKU() bool {
	return r.SKU != nil
}

// HasSortOrder verifica si se proporcionó un orden de clasificación
func (r *UpdateProductVariantRequest) HasSortOrder() bool {
	return r.SortOrder != nil
}

// HasAttributes verifica si se proporcionaron atributos
func (r *UpdateProductVariantRequest) HasAttributes() bool {
	return len(r.Attributes) > 0
}

// GetPrice devuelve el precio si está presente
func (r *UpdateProductVariantRequest) GetPrice() float64 {
	if r.Price == nil {
		return 0
	}
	return *r.Price
}

// GetStock devuelve el stock si está presente
func (r *UpdateProductVariantRequest) GetStock() int {
	if r.Stock == nil {
		return 0
	}
	return *r.Stock
}

// HasPrice verifica si se proporcionó un precio
func (r *UpdateProductVariantRequest) HasPrice() bool {
	return r.Price != nil
}

// HasStock verifica si se proporcionó un stock
func (r *UpdateProductVariantRequest) HasStock() bool {
	return r.Stock != nil
}
