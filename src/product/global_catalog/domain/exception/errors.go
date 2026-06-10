package exception

import "errors"

// Errores de validación
var (
	ErrGlobalproductInvalidName  = errors.New("nombre de global_product inválido")
	ErrGlobalproductNameRequired = errors.New("nombre de global_product es requerido")
)

// Errores de negocio
var (
	ErrGlobalproductNotFound      = errors.New("global_product no encontrado")
	ErrGlobalproductAlreadyExists = errors.New("global_product ya existe")
)

// Errores de persistencia
var (
	ErrGlobalproductCreateFailed = errors.New("error al crear global_product")
	ErrGlobalproductUpdateFailed = errors.New("error al actualizar global_product")
	ErrGlobalproductDeleteFailed = errors.New("error al eliminar global_product")
)
