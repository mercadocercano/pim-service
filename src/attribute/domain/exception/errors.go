package exception

import "errors"

// Errores de validación
var (
	ErrAttributeInvalidName  = errors.New("nombre de attribute inválido")
	ErrAttributeNameRequired = errors.New("nombre de attribute es requerido")
)

// Errores de negocio
var (
	ErrAttributeNotFound      = errors.New("attribute no encontrado")
	ErrAttributeAlreadyExists = errors.New("attribute ya existe")
)

// Errores de persistencia
var (
	ErrAttributeCreateFailed = errors.New("error al crear attribute")
	ErrAttributeUpdateFailed = errors.New("error al actualizar attribute")
	ErrAttributeDeleteFailed = errors.New("error al eliminar attribute")
)
