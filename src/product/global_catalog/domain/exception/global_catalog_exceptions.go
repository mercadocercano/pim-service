package exception

import "errors"

// Errores específicos del dominio Global Catalog
var (
	// Errores de validación
	ErrGlobalProductEANRequired    = errors.New("el código EAN-13 es obligatorio")
	ErrGlobalProductNameRequired   = errors.New("el nombre del producto es obligatorio")
	ErrGlobalProductSourceRequired = errors.New("la fuente del producto es obligatoria")
	ErrGlobalProductNotFound       = errors.New("producto global no encontrado")
	ErrGlobalProductAlreadyExists  = errors.New("ya existe un producto con ese EAN")
	ErrGlobalProductInvalidEAN     = errors.New("formato de EAN-13 inválido")
	ErrGlobalProductInvalidSource  = errors.New("fuente de producto inválida")
	ErrGlobalProductInvalidQuality = errors.New("quality score inválido")

	// Errores de operaciones
	ErrGlobalProductCannotBeDeleted = errors.New("no se puede eliminar el producto")
	ErrGlobalProductAlreadyDeleted  = errors.New("el producto ya está eliminado")
	ErrGlobalProductInactive        = errors.New("el producto está inactivo")

	// Errores de integración
	ErrScrapingSourceUnavailable = errors.New("fuente de scraping no disponible")
	ErrQualityScoreTooLow        = errors.New("quality score muy bajo para operación")
)

// ValidationError representa un error de validación
type ValidationError struct {
	Message string
	Cause   error
}

func (e ValidationError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}
	return e.Message
}

func NewValidationError(message string, cause error) *ValidationError {
	return &ValidationError{
		Message: message,
		Cause:   cause,
	}
}

// ConflictError representa un error de conflicto (recurso ya existe)
type ConflictError struct {
	Message string
}

func (e ConflictError) Error() string {
	return e.Message
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
	}
}

// InternalError representa un error interno del sistema
type InternalError struct {
	Message string
	Cause   error
}

func (e InternalError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}
	return e.Message
}

func NewInternalError(message string, cause error) *InternalError {
	return &InternalError{
		Message: message,
		Cause:   cause,
	}
}

// GlobalProductNotFoundError representa un error cuando no se encuentra un producto
type GlobalProductNotFoundError struct {
	EAN string
	ID  string
}

func (e GlobalProductNotFoundError) Error() string {
	if e.EAN != "" {
		return "producto global no encontrado con EAN: " + e.EAN
	}
	return "producto global no encontrado con ID: " + e.ID
}

func NewGlobalProductNotFoundByEAN(ean string) *GlobalProductNotFoundError {
	return &GlobalProductNotFoundError{EAN: ean}
}

func NewGlobalProductNotFoundByID(id string) *GlobalProductNotFoundError {
	return &GlobalProductNotFoundError{ID: id}
}

// GlobalProductAlreadyExistsError representa un error cuando ya existe un producto
type GlobalProductAlreadyExistsError struct {
	EAN string
}

func (e GlobalProductAlreadyExistsError) Error() string {
	return "ya existe un producto global con EAN: " + e.EAN
}

func NewGlobalProductAlreadyExists(ean string) *GlobalProductAlreadyExistsError {
	return &GlobalProductAlreadyExistsError{EAN: ean}
}
