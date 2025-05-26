package entity

import (
	"errors"
	"time"

	"pim/src/product/domain/value_object"

	"github.com/google/uuid"
)

// Product representa un producto en el sistema PIM y actúa como agregado raíz
type Product struct {
	id          uuid.UUID
	tenantID    string
	name        string
	description *string
	sku         *value_object.ProductSKU
	categoryRef *value_object.CategoryReference
	brandRef    *value_object.BrandReference
	status      value_object.ProductStatus
	variants    []*ProductVariant // Variantes como parte del agregado
	createdAt   time.Time
	updatedAt   time.Time
}

// NewProduct crea una nueva instancia de Product
func NewProduct(
	tenantID, name string,
	description *string,
	sku *value_object.ProductSKU,
	categoryRef *value_object.CategoryReference,
	brandRef *value_object.BrandReference,
) (*Product, error) {
	if tenantID == "" {
		return nil, errors.New("el tenant ID es obligatorio")
	}

	if name == "" {
		return nil, errors.New("el nombre del producto es obligatorio")
	}

	id := uuid.New()
	status := value_object.ProductStatusActive()

	product := &Product{
		id:          id,
		tenantID:    tenantID,
		name:        name,
		description: description,
		sku:         sku,
		categoryRef: categoryRef,
		brandRef:    brandRef,
		status:      status,
		variants:    make([]*ProductVariant, 0),
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}

	// Crear variante por defecto automáticamente
	defaultVariant, err := product.createDefaultVariant()
	if err != nil {
		return nil, err
	}
	product.variants = append(product.variants, defaultVariant)

	return product, nil
}

// NewProductFromRepository crea una instancia de Product desde datos de repositorio
func NewProductFromRepository(
	id uuid.UUID,
	tenantID, name string,
	description *string,
	sku *value_object.ProductSKU,
	categoryRef *value_object.CategoryReference,
	brandRef *value_object.BrandReference,
	status value_object.ProductStatus,
	createdAt, updatedAt time.Time,
) (*Product, error) {
	if tenantID == "" {
		return nil, errors.New("el tenant ID es obligatorio")
	}

	if name == "" {
		return nil, errors.New("el nombre del producto es obligatorio")
	}

	return &Product{
		id:          id,
		tenantID:    tenantID,
		name:        name,
		description: description,
		sku:         sku,
		categoryRef: categoryRef,
		brandRef:    brandRef,
		status:      status,
		variants:    make([]*ProductVariant, 0), // Se cargarán por separado
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil
}

// createDefaultVariant crea una variante por defecto para el producto
func (p *Product) createDefaultVariant() (*ProductVariant, error) {
	variantName := p.name + " - Default"

	// Crear atributos vacíos
	attributes, _ := value_object.NewVariantAttributeCollection([]*value_object.VariantAttribute{})

	return NewProductVariant(
		p.tenantID,
		p.id, // Ahora ambos son UUID
		variantName,
		p.sku, // Usar el mismo SKU del producto
		true,  // Es la variante por defecto
		1,     // Primer orden
		attributes,
	)
}

// Getters
func (p *Product) ID() uuid.UUID                                      { return p.id }
func (p *Product) IDString() string                                   { return p.id.String() }
func (p *Product) TenantID() string                                   { return p.tenantID }
func (p *Product) Name() string                                       { return p.name }
func (p *Product) Description() *string                               { return p.description }
func (p *Product) SKU() *value_object.ProductSKU                      { return p.sku }
func (p *Product) CategoryReference() *value_object.CategoryReference { return p.categoryRef }
func (p *Product) BrandReference() *value_object.BrandReference       { return p.brandRef }
func (p *Product) Status() value_object.ProductStatus                 { return p.status }
func (p *Product) Variants() []*ProductVariant                        { return p.variants }
func (p *Product) CreatedAt() time.Time                               { return p.createdAt }
func (p *Product) UpdatedAt() time.Time                               { return p.updatedAt }

// LoadVariants carga las variantes del producto (usado por el repositorio)
func (p *Product) LoadVariants(variants []*ProductVariant) {
	p.variants = variants
}

// GetVariants retorna todas las variantes activas
func (p *Product) GetVariants() []*ProductVariant {
	activeVariants := make([]*ProductVariant, 0)
	for _, variant := range p.variants {
		if !variant.IsDeleted() {
			activeVariants = append(activeVariants, variant)
		}
	}
	return activeVariants
}

// GetDefaultVariant retorna la variante por defecto
func (p *Product) GetDefaultVariant() *ProductVariant {
	for _, variant := range p.variants {
		if variant.IsDefault() && !variant.IsDeleted() {
			return variant
		}
	}
	return nil
}

// GetVariantByID retorna una variante específica por ID
func (p *Product) GetVariantByID(variantID uuid.UUID) *ProductVariant {
	for _, variant := range p.variants {
		if variant.ID() == variantID && !variant.IsDeleted() {
			return variant
		}
	}
	return nil
}

// AddVariant agrega una nueva variante al producto
func (p *Product) AddVariant(
	name string,
	sku *value_object.ProductSKU,
	isDefault bool,
	sortOrder int,
	attributes *value_object.VariantAttributeCollection,
) (*ProductVariant, error) {
	// Validar que no existe otra variante con el mismo nombre
	if p.hasVariantWithName(name) {
		return nil, errors.New("ya existe una variante con ese nombre")
	}

	// Si se marca como default, desmarcar la actual
	if isDefault {
		p.unsetCurrentDefaultVariant()
	}

	// Si no hay variantes activas y no se especifica isDefault, hacerla por defecto
	if len(p.GetVariants()) == 0 && !isDefault {
		isDefault = true
	}

	// Crear la nueva variante
	variant, err := NewProductVariant(
		p.tenantID,
		p.id, // Ya es UUID, no necesita parsing
		name,
		sku,
		isDefault,
		sortOrder,
		attributes,
	)
	if err != nil {
		return nil, err
	}

	// Agregar al agregado
	p.variants = append(p.variants, variant)
	p.updatedAt = time.Now()

	return variant, nil
}

// UpdateVariant actualiza una variante existente
func (p *Product) UpdateVariant(
	variantID uuid.UUID,
	name *string,
	sku *value_object.ProductSKU,
	sortOrder *int,
	attributes *value_object.VariantAttributeCollection,
) error {
	variant := p.GetVariantByID(variantID)
	if variant == nil {
		return errors.New("variante no encontrada")
	}

	// Actualizar campos si se proporcionan
	if name != nil {
		if p.hasVariantWithNameExcluding(*name, variantID) {
			return errors.New("ya existe una variante con ese nombre")
		}
		err := variant.UpdateName(*name)
		if err != nil {
			return err
		}
	}

	if sku != nil {
		variant.UpdateSKU(sku)
	}

	if sortOrder != nil {
		variant.UpdateSortOrder(*sortOrder)
	}

	if attributes != nil {
		variant.UpdateAttributes(attributes)
	}

	p.updatedAt = time.Now()
	return nil
}

// SetVariantAsDefault marca una variante como por defecto
func (p *Product) SetVariantAsDefault(variantID uuid.UUID) error {
	variant := p.GetVariantByID(variantID)
	if variant == nil {
		return errors.New("variante no encontrada")
	}

	// Desmarcar la variante por defecto actual
	p.unsetCurrentDefaultVariant()

	// Marcar la nueva como por defecto
	variant.SetAsDefault()
	p.updatedAt = time.Now()

	return nil
}

// DeleteVariant elimina una variante (soft delete)
func (p *Product) DeleteVariant(variantID uuid.UUID) error {
	variant := p.GetVariantByID(variantID)
	if variant == nil {
		return errors.New("variante no encontrada")
	}

	// No permitir eliminar la última variante activa
	activeVariants := p.GetVariants()
	if len(activeVariants) <= 1 {
		return errors.New("no se puede eliminar la última variante activa")
	}

	// Si es la variante por defecto, asignar otra como por defecto
	wasDefault := variant.IsDefault()

	// Eliminar la variante
	err := variant.Delete()
	if err != nil {
		return err
	}

	// Si era la por defecto, asignar otra
	if wasDefault {
		p.assignNewDefaultVariant()
	}

	p.updatedAt = time.Now()
	return nil
}

// hasVariantWithName verifica si existe una variante con el nombre dado
func (p *Product) hasVariantWithName(name string) bool {
	for _, variant := range p.variants {
		if variant.Name() == name && !variant.IsDeleted() {
			return true
		}
	}
	return false
}

// hasVariantWithNameExcluding verifica si existe una variante con el nombre dado, excluyendo una específica
func (p *Product) hasVariantWithNameExcluding(name string, excludeID uuid.UUID) bool {
	for _, variant := range p.variants {
		if variant.Name() == name && variant.ID() != excludeID && !variant.IsDeleted() {
			return true
		}
	}
	return false
}

// unsetCurrentDefaultVariant desmarca la variante por defecto actual
func (p *Product) unsetCurrentDefaultVariant() {
	for _, variant := range p.variants {
		if variant.IsDefault() {
			variant.UnsetAsDefault()
			break
		}
	}
}

// assignNewDefaultVariant asigna una nueva variante por defecto
func (p *Product) assignNewDefaultVariant() {
	activeVariants := p.GetVariants()
	if len(activeVariants) > 0 {
		// Asignar la primera variante activa como por defecto
		activeVariants[0].SetAsDefault()
	}
}

// Update actualiza los datos del producto
func (p *Product) Update(
	name string,
	description *string,
	sku *value_object.ProductSKU,
	categoryRef *value_object.CategoryReference,
	brandRef *value_object.BrandReference,
) error {
	if name == "" {
		return errors.New("el nombre del producto es obligatorio")
	}

	p.name = name
	p.description = description
	p.sku = sku
	p.categoryRef = categoryRef
	p.brandRef = brandRef
	p.updatedAt = time.Now()

	return nil
}

// Activate activa el producto
func (p *Product) Activate() {
	if !p.status.IsDeleted() {
		p.status = value_object.ProductStatusActive()
		p.updatedAt = time.Now()
	}
}

// Deactivate desactiva el producto
func (p *Product) Deactivate() {
	if !p.status.IsDeleted() {
		p.status = value_object.ProductStatusInactive()
		p.updatedAt = time.Now()
	}
}

// Discontinue marca el producto como descontinuado
func (p *Product) Discontinue() {
	if !p.status.IsDeleted() {
		p.status = value_object.ProductStatusDiscontinued()
		p.updatedAt = time.Now()
	}
}

// Delete elimina el producto (soft delete)
func (p *Product) Delete() {
	p.status = value_object.ProductStatusDeleted()
	p.updatedAt = time.Now()

	// También eliminar todas las variantes
	for _, variant := range p.variants {
		variant.Delete()
	}
}

// IsActive verifica si el producto está activo
func (p *Product) IsActive() bool {
	return p.status.IsActive()
}

// IsDeleted verifica si el producto está eliminado
func (p *Product) IsDeleted() bool {
	return p.status.IsDeleted()
}

// IsDiscontinued verifica si el producto está descontinuado
func (p *Product) IsDiscontinued() bool {
	return p.status.IsDiscontinued()
}

// CanBeDeleted verifica si el producto puede ser eliminado
func (p *Product) CanBeDeleted() bool {
	return !p.status.IsDeleted()
}

// HasSKU verifica si el producto tiene SKU
func (p *Product) HasSKU() bool {
	return p.sku != nil && !p.sku.IsEmpty()
}

// HasCategory verifica si el producto tiene categoría
func (p *Product) HasCategory() bool {
	return p.categoryRef != nil && !p.categoryRef.IsEmpty()
}

// HasBrand verifica si el producto tiene marca
func (p *Product) HasBrand() bool {
	return p.brandRef != nil && !p.brandRef.IsEmpty()
}

// HasVariants verifica si el producto tiene variantes
func (p *Product) HasVariants() bool {
	return len(p.GetVariants()) > 0
}

// GetCategoryID retorna el ID de la categoría si existe
func (p *Product) GetCategoryID() *string {
	if !p.HasCategory() {
		return nil
	}
	id := p.categoryRef.ID()
	return &id
}

// GetBrandID retorna el ID de la marca si existe
func (p *Product) GetBrandID() *string {
	if !p.HasBrand() {
		return nil
	}
	id := p.brandRef.ID()
	return &id
}

// GetCategoryName retorna el nombre de la categoría si existe
func (p *Product) GetCategoryName() *string {
	if !p.HasCategory() {
		return nil
	}
	name := p.categoryRef.Name()
	return &name
}

// GetBrandName retorna el nombre de la marca si existe
func (p *Product) GetBrandName() *string {
	if !p.HasBrand() {
		return nil
	}
	name := p.brandRef.Name()
	return &name
}
