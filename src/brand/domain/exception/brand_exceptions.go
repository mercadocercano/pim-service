package exception

import "errors"

// Errores específicos del dominio Brand
var (
	ErrBrandNotFound       = errors.New("marca no encontrada")
	ErrBrandAlreadyExists  = errors.New("ya existe una marca con ese nombre")
	ErrBrandInvalidData    = errors.New("datos de marca inválidos")
	ErrBrandCannotDelete   = errors.New("no se puede eliminar la marca")
	ErrBrandInvalidStatus  = errors.New("estado de marca inválido")
	ErrBrandNameRequired   = errors.New("el nombre de la marca es obligatorio")
	ErrBrandTenantRequired = errors.New("el tenant ID es obligatorio")
)

// BrandDomainError representa un error específico del dominio Brand
type BrandDomainError struct {
	Code    string
	Message string
	Cause   error
}

// Error implementa la interfaz error
func (e *BrandDomainError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}
	return e.Message
}

// NewBrandDomainError crea un nuevo error de dominio
func NewBrandDomainError(code, message string, cause error) *BrandDomainError {
	return &BrandDomainError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

// Códigos de error específicos
const (
	BrandErrorCodeNotFound      = "BRAND_NOT_FOUND"
	BrandErrorCodeAlreadyExists = "BRAND_ALREADY_EXISTS"
	BrandErrorCodeInvalidData   = "BRAND_INVALID_DATA"
	BrandErrorCodeCannotDelete  = "BRAND_CANNOT_DELETE"
	BrandErrorCodeInvalidStatus = "BRAND_INVALID_STATUS"
)
