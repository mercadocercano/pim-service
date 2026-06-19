package value_object

import "fmt"

const (
	// NormalizeCategoryMaxRowsCap es el límite de CATEGORÍAS DISTINTAS que una operación de
	// normalización puede evaluar. La unidad es la categoría cruda distinta (no la fila): el
	// resolver es función pura del string, así que todas las filas con la misma category raw
	// reciben el mismo slug. Hoy hay ~2.283 categorías distintas en global_products (ADR-007).
	NormalizeCategoryMaxRowsCap = 10000

	// ErrCodeNormalizeScopeExceedsCap es el código semántico cuando max_rows excede el cap.
	ErrCodeNormalizeScopeExceedsCap = "SCOPE_EXCEEDS_CAP"
)

// NormalizeCategoryScope define el criterio de selección para la normalización de category_slug.
// Value object inmutable, validado en el constructor (mismo patrón que ReclassifyScope, ADR-005).
type NormalizeCategoryScope struct {
	// SourcePrefix filtra global_products por source LIKE '<prefix>%'. Vacío = sin filtro.
	SourcePrefix string

	// MaxRows es el cap de CATEGORÍAS DISTINTAS evaluadas. Entre 1 y NormalizeCategoryMaxRowsCap.
	MaxRows int
}

// NormalizeScopeError es el error de validación del scope.
type NormalizeScopeError struct {
	Code    string
	Message string
}

func (e *NormalizeScopeError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewNormalizeCategoryScope construye un scope válido. Error si MaxRows <= 0 o supera el cap.
func NewNormalizeCategoryScope(sourcePrefix string, maxRows int) (NormalizeCategoryScope, error) {
	if maxRows <= 0 {
		return NormalizeCategoryScope{}, &NormalizeScopeError{
			Code:    "SCOPE_INVALID_MAX_ROWS",
			Message: "max_rows must be greater than 0",
		}
	}
	if maxRows > NormalizeCategoryMaxRowsCap {
		return NormalizeCategoryScope{}, &NormalizeScopeError{
			Code:    ErrCodeNormalizeScopeExceedsCap,
			Message: fmt.Sprintf("max_rows %d exceeds the allowed cap of %d", maxRows, NormalizeCategoryMaxRowsCap),
		}
	}
	return NormalizeCategoryScope{SourcePrefix: sourcePrefix, MaxRows: maxRows}, nil
}
