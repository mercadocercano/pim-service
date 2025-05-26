package exception

import "errors"

// Errores específicos del dominio Product
var (
	// Errores de validación
	ErrProductNameRequired     = errors.New("el nombre del producto es obligatorio")
	ErrProductTenantRequired   = errors.New("el tenant ID es obligatorio")
	ErrProductNotFound         = errors.New("producto no encontrado")
	ErrProductAlreadyExists    = errors.New("ya existe un producto con ese nombre")
	ErrProductSKUAlreadyExists = errors.New("ya existe un producto con ese SKU")
	ErrProductInvalidSKU       = errors.New("formato de SKU inválido")
	ErrProductInvalidStatus    = errors.New("estado de producto inválido")

	// Errores de operaciones
	ErrProductCannotBeDeleted    = errors.New("no se puede eliminar el producto")
	ErrProductAlreadyDeleted     = errors.New("el producto ya está eliminado")
	ErrProductCannotChangeStatus = errors.New("no se puede cambiar el estado del producto")

	// Errores de referencias
	ErrInvalidCategoryReference = errors.New("referencia de categoría inválida")
	ErrInvalidBrandReference    = errors.New("referencia de marca inválida")
)

// ProductNotFoundError representa un error cuando no se encuentra un producto
type ProductNotFoundError struct {
	ID       string
	TenantID string
}

func (e ProductNotFoundError) Error() string {
	return "producto no encontrado"
}

// ProductAlreadyExistsError representa un error cuando ya existe un producto
type ProductAlreadyExistsError struct {
	Name     string
	TenantID string
}

func (e ProductAlreadyExistsError) Error() string {
	return "ya existe un producto con ese nombre"
}

// ProductSKUAlreadyExistsError representa un error cuando ya existe un SKU
type ProductSKUAlreadyExistsError struct {
	SKU      string
	TenantID string
}

func (e ProductSKUAlreadyExistsError) Error() string {
	return "ya existe un producto con ese SKU"
}
