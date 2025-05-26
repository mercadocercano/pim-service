package request

// UpdateProductVariantRequest representa la solicitud para actualizar una variante de producto
type UpdateProductVariantRequest struct {
	Name       *string                         `json:"name,omitempty" binding:"omitempty,min=2,max=255"`
	SKU        *string                         `json:"sku,omitempty" binding:"omitempty,min=3,max=50"`
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
