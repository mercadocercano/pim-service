package entity

import (
	"saas-mt-pim-service/src/schema_validation/domain/value_object"
)

// CellValidation representa la validación de una celda individual
type CellValidation struct {
	// Value es el valor original de la celda
	Value string `json:"value"`

	// Status es el estado de validación de la celda
	Status value_object.ValidationStatus `json:"status"`

	// Message es el mensaje de error o advertencia (opcional)
	Message string `json:"message,omitempty"`

	// CorrectedValue es el valor corregido sugerido (opcional)
	CorrectedValue string `json:"corrected_value,omitempty"`

	// ColumnIndex es el índice de la columna
	ColumnIndex int `json:"column_index"`

	// RowIndex es el índice de la fila
	RowIndex int `json:"row_index"`
}

// NewCellValidation crea una nueva validación de celda
func NewCellValidation(value string, columnIndex, rowIndex int) *CellValidation {
	return &CellValidation{
		Value:       value,
		Status:      value_object.ValidationStatusValid,
		ColumnIndex: columnIndex,
		RowIndex:    rowIndex,
	}
}

// SetError marca la celda como error con un mensaje
func (c *CellValidation) SetError(message string) {
	c.Status = value_object.ValidationStatusError
	c.Message = message
}

// SetWarning marca la celda como advertencia con un mensaje
func (c *CellValidation) SetWarning(message string) {
	c.Status = value_object.ValidationStatusWarning
	c.Message = message
}

// SetInfo marca la celda como información
func (c *CellValidation) SetInfo(message string) {
	c.Status = value_object.ValidationStatusInfo
	c.Message = message
}

// SuggestCorrection sugiere un valor corregido
func (c *CellValidation) SuggestCorrection(correctedValue string) {
	c.CorrectedValue = correctedValue
}

// IsValid retorna true si la celda es válida
func (c *CellValidation) IsValid() bool {
	return c.Status.IsValid()
}
