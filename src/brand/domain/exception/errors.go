package exception

import "errors"

// Errores de validación
var (
	ErrMarketplacebrandInvalidName  = errors.New("nombre de marketplace_brand inválido")
	ErrMarketplacebrandNameRequired = errors.New("nombre de marketplace_brand es requerido")
)

// Errores de negocio
var (
	ErrMarketplacebrandNotFound      = errors.New("marketplace_brand no encontrado")
	ErrMarketplacebrandAlreadyExists = errors.New("marketplace_brand ya existe")
)

// Errores de persistencia
var (
	ErrMarketplacebrandCreateFailed = errors.New("error al crear marketplace_brand")
	ErrMarketplacebrandUpdateFailed = errors.New("error al actualizar marketplace_brand")
	ErrMarketplacebrandDeleteFailed = errors.New("error al eliminar marketplace_brand")
)
