package exception

import "fmt"

// ProductVariantNotFound se lanza cuando no se encuentra una variante
type ProductVariantNotFound struct {
	ID       string
	TenantID string
}

func (e ProductVariantNotFound) Error() string {
	return fmt.Sprintf("variante de producto no encontrada: ID=%s, TenantID=%s", e.ID, e.TenantID)
}

// ProductVariantAlreadyExists se lanza cuando ya existe una variante con el mismo nombre o SKU
type ProductVariantAlreadyExists struct {
	Field    string // "name" o "sku"
	Value    string
	TenantID string
}

func (e ProductVariantAlreadyExists) Error() string {
	return fmt.Sprintf("ya existe una variante con %s '%s' en el tenant %s", e.Field, e.Value, e.TenantID)
}

// ProductVariantInvalidStatus se lanza cuando se intenta una transición de estado inválida
type ProductVariantInvalidStatus struct {
	CurrentStatus string
	TargetStatus  string
	Operation     string
}

func (e ProductVariantInvalidStatus) Error() string {
	return fmt.Sprintf("no se puede %s: transición inválida de estado '%s' a '%s'", e.Operation, e.CurrentStatus, e.TargetStatus)
}

// ProductVariantDefaultRequired se lanza cuando se intenta eliminar la única variante por defecto
type ProductVariantDefaultRequired struct {
	ProductID string
}

func (e ProductVariantDefaultRequired) Error() string {
	return fmt.Sprintf("el producto %s debe tener al menos una variante por defecto", e.ProductID)
}

// ProductVariantLastActiveVariant se lanza cuando se intenta eliminar la última variante activa
type ProductVariantLastActiveVariant struct {
	ProductID string
}

func (e ProductVariantLastActiveVariant) Error() string {
	return fmt.Sprintf("no se puede eliminar la única variante activa del producto %s", e.ProductID)
}

// ProductVariantInvalidAttribute se lanza cuando un atributo de variante es inválido
type ProductVariantInvalidAttribute struct {
	AttributeName  string
	AttributeValue string
	Reason         string
}

func (e ProductVariantInvalidAttribute) Error() string {
	return fmt.Sprintf("atributo de variante inválido '%s'='%s': %s", e.AttributeName, e.AttributeValue, e.Reason)
}

// ProductVariantDuplicateAttribute se lanza cuando se intenta agregar un atributo duplicado
type ProductVariantDuplicateAttribute struct {
	AttributeName string
}

func (e ProductVariantDuplicateAttribute) Error() string {
	return fmt.Sprintf("atributo de variante duplicado: '%s'", e.AttributeName)
}

// ProductVariantInvalidSKU se lanza cuando el SKU de la variante es inválido
type ProductVariantInvalidSKU struct {
	SKU    string
	Reason string
}

func (e ProductVariantInvalidSKU) Error() string {
	return fmt.Sprintf("SKU de variante inválido '%s': %s", e.SKU, e.Reason)
}

// ProductVariantProductNotFound se lanza cuando el producto padre no existe
type ProductVariantProductNotFound struct {
	ProductID string
	TenantID  string
}

func (e ProductVariantProductNotFound) Error() string {
	return fmt.Sprintf("producto padre no encontrado: ID=%s, TenantID=%s", e.ProductID, e.TenantID)
}

// ProductVariantTenantMismatch se lanza cuando el tenant de la variante no coincide con el del producto
type ProductVariantTenantMismatch struct {
	VariantTenantID string
	ProductTenantID string
}

func (e ProductVariantTenantMismatch) Error() string {
	return fmt.Sprintf("el tenant de la variante (%s) no coincide con el del producto (%s)", e.VariantTenantID, e.ProductTenantID)
}
