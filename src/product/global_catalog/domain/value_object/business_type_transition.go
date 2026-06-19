package value_object

// businessTypeAlmacen es el rubro genérico de relleno previo: un producto en "almacen"
// se considera NO curado y puede corregirse a un rubro específico.
const businessTypeAlmacen = "almacen"

// Kinds de una transición de business_type. Los valores de los kinds que aplican
// ("relleno"/"correccion") coinciden con los strings que ya usa ReclassifyUpdate.Kind,
// y los de skip coinciden con los SkipKind de reclassify_result.go. La coincidencia es
// intencional: permite mapear el resultado de esta función pura tanto al camino batch
// (classifyCandidates) como al camino sync (UpdateGlobalProductByID) sin traducción.
const (
	TransitionRelleno    = "relleno"    // current nil/vacío → candidate
	TransitionCorreccion = "correccion" // current "almacen" → candidate específico

	TransitionSkipSinCandidate = "sin_candidate" // no hay candidate (vacío) que proponer
	TransitionSkipYaEspecifico = "ya_especifico" // current ya curado en rubro específico
	TransitionSkipYaCorrecto   = "ya_correcto"   // candidate == current; 0 escrituras
)

// ResolveSafeBusinessTypeTransition aplica la invariante de corrección segura (ADR-005 §8,
// extendida por ADR-006) sobre el par (estado actual del producto, candidate ya resuelto).
//
// Es la ÚNICA implementación de la política de aceptación/skip de business_type. La consumen:
//   - el reclassify batch (ReclassifyBusinessTypesUseCase.classifyCandidates), que resuelve el
//     candidate desde la categoría y luego decide con esta función;
//   - el update producto-a-producto (UpdateGlobalProductByID), que recibe el candidate ya
//     resuelto desde webdata en cada re-sync.
//
// Función pura: misma entrada → misma salida, sin efectos. Reglas (orden load-bearing):
//   - candidate vacío            → skip (sin_candidate); nada que proponer.
//   - current ya específico      → skip (ya_especifico); nunca degrada un rubro curado.
//   - candidate == current       → skip (ya_correcto); 0 escrituras.
//   - current nil/vacío          → apply relleno.
//   - current == "almacen"       → apply correccion.
func ResolveSafeBusinessTypeTransition(current *string, candidate string) (apply bool, newType string, kind string) {
	if candidate == "" {
		return false, "", TransitionSkipSinCandidate
	}

	// current específico (no nil, no vacío, no "almacen") → nunca se toca.
	if current != nil && *current != "" && *current != businessTypeAlmacen {
		return false, "", TransitionSkipYaEspecifico
	}

	// candidate coincide con el valor actual → no hay cambio.
	if current != nil && *current == candidate {
		return false, "", TransitionSkipYaCorrecto
	}

	// current == "almacen" → corrección desde genérico a específico.
	if current != nil && *current == businessTypeAlmacen {
		return true, candidate, TransitionCorreccion
	}

	// current nil/vacío → relleno.
	return true, candidate, TransitionRelleno
}
