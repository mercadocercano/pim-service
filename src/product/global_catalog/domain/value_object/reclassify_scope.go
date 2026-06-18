package value_object

import "fmt"

const (
	// ReclassifyMaxRowsCap es el límite absoluto de filas que puede evaluar una operación
	// de re-clasificación. Protege contra operaciones desbocadas sobre la tabla global_products.
	// Si CountCandidates supera este valor, el use case devuelve 422 (ADR-005 §6).
	ReclassifyMaxRowsCap = 50000

	// ErrCodeScopeExceedsCap es el código semántico devuelto cuando max_rows excede el cap.
	ErrCodeScopeExceedsCap = "SCOPE_EXCEEDS_CAP"
)

// ReclassifyScope define el criterio de selección de candidatos para re-clasificación.
// Es un value object inmutable: se valida en el constructor y no cambia después.
type ReclassifyScope struct {
	// SourcePrefix filtra global_products por source LIKE '<prefix>%'.
	// Vacío string acepta cualquier source (sin filtro de prefijo).
	SourcePrefix string

	// MaxRows es el cap máximo de candidatos permitidos en una operación.
	// Debe estar entre 1 y ReclassifyMaxRowsCap (50000).
	MaxRows int
}

// ReclassScopeError es el error de validación del scope.
type ReclassScopeError struct {
	Code    string
	Message string
}

func (e *ReclassScopeError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewReclassifyScope construye un ReclassifyScope válido.
// Devuelve error si MaxRows es 0 o supera ReclassifyMaxRowsCap.
//
// TEST-ID: T-013 (MaxRows=0 → error), T-014 (MaxRows=50001 → SCOPE_EXCEEDS_CAP), T-015 (MaxRows=50000 → ok).
func NewReclassifyScope(sourcePrefix string, maxRows int) (ReclassifyScope, error) {
	if maxRows <= 0 {
		return ReclassifyScope{}, &ReclassScopeError{
			Code:    "SCOPE_INVALID_MAX_ROWS",
			Message: "max_rows must be greater than 0",
		}
	}
	if maxRows > ReclassifyMaxRowsCap {
		return ReclassifyScope{}, &ReclassScopeError{
			Code:    ErrCodeScopeExceedsCap,
			Message: fmt.Sprintf("max_rows %d exceeds the allowed cap of %d", maxRows, ReclassifyMaxRowsCap),
		}
	}
	return ReclassifyScope{
		SourcePrefix: sourcePrefix,
		MaxRows:      maxRows,
	}, nil
}
