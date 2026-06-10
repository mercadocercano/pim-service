package value_object

// ValidationStatus representa el estado de validación de una columna o celda
type ValidationStatus string

const (
	// ValidationStatusValid indica que el valor es correcto y válido
	ValidationStatusValid ValidationStatus = "valid"

	// ValidationStatusWarning indica que el valor es aceptable pero con advertencias
	ValidationStatusWarning ValidationStatus = "warning"

	// ValidationStatusError indica que el valor es inválido y requiere corrección
	ValidationStatusError ValidationStatus = "error"

	// ValidationStatusInfo indica información adicional, como columna no mapeada
	ValidationStatusInfo ValidationStatus = "info"
)

// IsValid retorna true si el estado es válido o warning
func (s ValidationStatus) IsValid() bool {
	return s == ValidationStatusValid || s == ValidationStatusWarning
}

// String retorna la representación en string del estado
func (s ValidationStatus) String() string {
	return string(s)
}

// GetColor retorna el color asociado al estado para visualización
func (s ValidationStatus) GetColor() string {
	switch s {
	case ValidationStatusValid:
		return "#d4edda" // Verde
	case ValidationStatusWarning:
		return "#fff3cd" // Amarillo
	case ValidationStatusError:
		return "#f8d7da" // Rojo
	case ValidationStatusInfo:
		return "#d1ecf1" // Azul
	default:
		return "#ffffff" // Blanco
	}
}

// GetEmoji retorna el emoji asociado al estado
func (s ValidationStatus) GetEmoji() string {
	switch s {
	case ValidationStatusValid:
		return "🟢"
	case ValidationStatusWarning:
		return "🟡"
	case ValidationStatusError:
		return "🔴"
	case ValidationStatusInfo:
		return "🔵"
	default:
		return "⚪"
	}
}
