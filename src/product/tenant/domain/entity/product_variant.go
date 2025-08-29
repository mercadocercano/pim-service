package entity

import (
	"fmt"
	"time"

	"saas-mt-pim-service/src/product/tenant/domain/value_object"

	"github.com/google/uuid"
)

// ProductVariant representa una variante específica de un producto
type ProductVariant struct {
	id         uuid.UUID
	tenantID   string
	productID  uuid.UUID
	name       string
	sku        *value_object.ProductSKU
	status     *value_object.VariantStatus
	isDefault  bool
	sortOrder  int
	attributes *value_object.VariantAttributeCollection
	createdAt  time.Time
	updatedAt  time.Time
}

// NewProductVariant crea una nueva variante de producto
func NewProductVariant(
	tenantID string,
	productID uuid.UUID,
	name string,
	sku *value_object.ProductSKU,
	isDefault bool,
	sortOrder int,
	attributes *value_object.VariantAttributeCollection,
) (*ProductVariant, error) {
	if err := validateVariantCreation(tenantID, productID, name, attributes); err != nil {
		return nil, err
	}

	id := uuid.New()
	status := value_object.NewActiveVariantStatus()
	now := time.Now()

	if attributes == nil {
		attributes, _ = value_object.NewVariantAttributeCollection([]*value_object.VariantAttribute{})
	}

	return &ProductVariant{
		id:         id,
		tenantID:   tenantID,
		productID:  productID,
		name:       name,
		sku:        sku,
		status:     status,
		isDefault:  isDefault,
		sortOrder:  sortOrder,
		attributes: attributes,
		createdAt:  now,
		updatedAt:  now,
	}, nil
}

// NewProductVariantFromRepository reconstruye una variante desde el repositorio
func NewProductVariantFromRepository(
	id uuid.UUID,
	tenantID string,
	productID uuid.UUID,
	name string,
	sku *value_object.ProductSKU,
	status *value_object.VariantStatus,
	isDefault bool,
	sortOrder int,
	attributes *value_object.VariantAttributeCollection,
	createdAt, updatedAt time.Time,
) *ProductVariant {
	if attributes == nil {
		attributes, _ = value_object.NewVariantAttributeCollection([]*value_object.VariantAttribute{})
	}

	return &ProductVariant{
		id:         id,
		tenantID:   tenantID,
		productID:  productID,
		name:       name,
		sku:        sku,
		status:     status,
		isDefault:  isDefault,
		sortOrder:  sortOrder,
		attributes: attributes,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
	}
}

// validateVariantCreation valida los datos para crear una variante
func validateVariantCreation(tenantID string, productID uuid.UUID, name string, attributes *value_object.VariantAttributeCollection) error {
	if tenantID == "" {
		return fmt.Errorf("el tenant ID es requerido")
	}

	if productID == uuid.Nil {
		return fmt.Errorf("el ID del producto es requerido")
	}

	if err := validateVariantName(name); err != nil {
		return err
	}

	return nil
}

// validateVariantName valida el nombre de la variante
func validateVariantName(name string) error {
	if name == "" {
		return fmt.Errorf("el nombre de la variante es requerido")
	}

	if len(name) < 2 {
		return fmt.Errorf("el nombre de la variante debe tener al menos 2 caracteres")
	}

	if len(name) > 255 {
		return fmt.Errorf("el nombre de la variante no puede exceder 255 caracteres")
	}

	return nil
}

// Getters
func (pv *ProductVariant) ID() uuid.UUID {
	return pv.id
}

func (pv *ProductVariant) TenantID() string {
	return pv.tenantID
}

func (pv *ProductVariant) ProductID() uuid.UUID {
	return pv.productID
}

func (pv *ProductVariant) Name() string {
	return pv.name
}

func (pv *ProductVariant) SKU() *value_object.ProductSKU {
	return pv.sku
}

func (pv *ProductVariant) Status() *value_object.VariantStatus {
	return pv.status
}

func (pv *ProductVariant) IsDefault() bool {
	return pv.isDefault
}

func (pv *ProductVariant) SortOrder() int {
	return pv.sortOrder
}

func (pv *ProductVariant) Attributes() *value_object.VariantAttributeCollection {
	return pv.attributes
}

func (pv *ProductVariant) CreatedAt() time.Time {
	return pv.createdAt
}

func (pv *ProductVariant) UpdatedAt() time.Time {
	return pv.updatedAt
}

// Métodos de negocio

// UpdateName actualiza el nombre de la variante
func (pv *ProductVariant) UpdateName(name string) error {
	if err := validateVariantName(name); err != nil {
		return err
	}

	pv.name = name
	pv.updatedAt = time.Now()
	return nil
}

// UpdateSKU actualiza el SKU de la variante
func (pv *ProductVariant) UpdateSKU(sku *value_object.ProductSKU) {
	pv.sku = sku
	pv.updatedAt = time.Now()
}

// UpdateSortOrder actualiza el orden de la variante
func (pv *ProductVariant) UpdateSortOrder(sortOrder int) {
	pv.sortOrder = sortOrder
	pv.updatedAt = time.Now()
}

// UpdateAttributes actualiza los atributos de la variante
func (pv *ProductVariant) UpdateAttributes(attributes *value_object.VariantAttributeCollection) {
	if attributes == nil {
		attributes, _ = value_object.NewVariantAttributeCollection([]*value_object.VariantAttribute{})
	}

	pv.attributes = attributes
	pv.updatedAt = time.Now()
}

// SetAsDefault marca esta variante como la por defecto
func (pv *ProductVariant) SetAsDefault() {
	pv.isDefault = true
	pv.updatedAt = time.Now()
}

// UnsetAsDefault desmarca esta variante como la por defecto
func (pv *ProductVariant) UnsetAsDefault() {
	pv.isDefault = false
	pv.updatedAt = time.Now()
}

// Activate activa la variante
func (pv *ProductVariant) Activate() error {
	if !pv.status.CanBeActivated() {
		return fmt.Errorf("la variante no puede ser activada desde el estado: %s", pv.status.Value())
	}

	status := value_object.NewActiveVariantStatus()
	pv.status = status
	pv.updatedAt = time.Now()
	return nil
}

// Deactivate desactiva la variante
func (pv *ProductVariant) Deactivate() error {
	if !pv.status.CanBeDeactivated() {
		return fmt.Errorf("la variante no puede ser desactivada desde el estado: %s", pv.status.Value())
	}

	status, err := value_object.NewVariantStatus(value_object.VariantStatusInactive)
	if err != nil {
		return err
	}

	pv.status = status
	pv.updatedAt = time.Now()
	return nil
}

// Discontinue descontinúa la variante
func (pv *ProductVariant) Discontinue() error {
	if !pv.status.CanBeDiscontinued() {
		return fmt.Errorf("la variante no puede ser descontinuada desde el estado: %s", pv.status.Value())
	}

	status, err := value_object.NewVariantStatus(value_object.VariantStatusDiscontinued)
	if err != nil {
		return err
	}

	pv.status = status
	pv.updatedAt = time.Now()
	return nil
}

// Delete elimina la variante (soft delete)
func (pv *ProductVariant) Delete() error {
	if !pv.status.CanBeDeleted() {
		return fmt.Errorf("la variante no puede ser eliminada desde el estado: %s", pv.status.Value())
	}

	status, err := value_object.NewVariantStatus(value_object.VariantStatusDeleted)
	if err != nil {
		return err
	}

	pv.status = status
	pv.updatedAt = time.Now()
	return nil
}

// Métodos de verificación

// HasSKU verifica si la variante tiene SKU
func (pv *ProductVariant) HasSKU() bool {
	return pv.sku != nil
}

// HasAttributes verifica si la variante tiene atributos
func (pv *ProductVariant) HasAttributes() bool {
	return pv.attributes != nil && !pv.attributes.IsEmpty()
}

// HasAttribute verifica si la variante tiene un atributo específico
func (pv *ProductVariant) HasAttribute(name string) bool {
	return pv.attributes != nil && pv.attributes.HasAttribute(name)
}

// GetAttributeValue obtiene el valor de un atributo específico
func (pv *ProductVariant) GetAttributeValue(name string) string {
	if pv.attributes == nil {
		return ""
	}

	attr := pv.attributes.GetByName(name)
	if attr == nil {
		return ""
	}

	return attr.Value()
}

// IsActive verifica si la variante está activa
func (pv *ProductVariant) IsActive() bool {
	return pv.status.IsActive()
}

// IsInactive verifica si la variante está inactiva
func (pv *ProductVariant) IsInactive() bool {
	return pv.status.IsInactive()
}

// IsDiscontinued verifica si la variante está descontinuada
func (pv *ProductVariant) IsDiscontinued() bool {
	return pv.status.IsDiscontinued()
}

// IsDeleted verifica si la variante está eliminada
func (pv *ProductVariant) IsDeleted() bool {
	return pv.status.IsDeleted()
}
