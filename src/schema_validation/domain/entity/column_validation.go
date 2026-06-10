package entity

import (
	"saas-mt-pim-service/src/schema_validation/domain/value_object"
)

// ColumnValidation representa la validación de una columna completa
type ColumnValidation struct {
	// Name es el nombre de la columna en el CSV
	Name string `json:"name"`

	// Index es la posición de la columna
	Index int `json:"index"`

	// Status es el estado general de la columna
	Status value_object.ValidationStatus `json:"status"`

	// TypeExpected es el tipo de dato esperado
	TypeExpected string `json:"type_expected"`

	// TypeDetected es el tipo de dato detectado
	TypeDetected string `json:"type_detected"`

	// Required indica si la columna es obligatoria
	Required bool `json:"required"`

	// MappedTo es el nombre del campo al que se mapea
	MappedTo string `json:"mapped_to,omitempty"`

	// Statistics contiene estadísticas de la columna
	Statistics ColumnStatistics `json:"statistics"`

	// Issues son los problemas encontrados
	Issues []string `json:"issues"`

	// SampleValues son valores de ejemplo
	SampleValues []string `json:"sample_values"`

	// InvalidExamples son ejemplos de valores inválidos
	InvalidExamples []InvalidExample `json:"invalid_examples,omitempty"`
}

// ColumnStatistics contiene estadísticas de una columna
type ColumnStatistics struct {
	ValidCount   int     `json:"valid_count"`
	InvalidCount int     `json:"invalid_count"`
	NullCount    int     `json:"null_count"`
	UniqueCount  int     `json:"unique_count"`
	TotalCount   int     `json:"total_count"`
	ValidRate    float64 `json:"valid_rate"`
}

// InvalidExample representa un ejemplo de valor inválido
type InvalidExample struct {
	Row    int    `json:"row"`
	Value  string `json:"value"`
	Reason string `json:"reason"`
}

// NewColumnValidation crea una nueva validación de columna
func NewColumnValidation(name string, index int) *ColumnValidation {
	return &ColumnValidation{
		Name:            name,
		Index:           index,
		Status:          value_object.ValidationStatusInfo,
		Issues:          make([]string, 0),
		SampleValues:    make([]string, 0),
		InvalidExamples: make([]InvalidExample, 0),
	}
}

// SetRequired marca la columna como requerida
func (c *ColumnValidation) SetRequired(required bool) {
	c.Required = required
}

// SetExpectedType establece el tipo esperado
func (c *ColumnValidation) SetExpectedType(expectedType string) {
	c.TypeExpected = expectedType
}

// SetDetectedType establece el tipo detectado
func (c *ColumnValidation) SetDetectedType(detectedType string) {
	c.TypeDetected = detectedType
}

// MapTo establece el mapeo de la columna
func (c *ColumnValidation) MapTo(fieldName string) {
	c.MappedTo = fieldName
}

// AddIssue agrega un problema a la columna
func (c *ColumnValidation) AddIssue(issue string) {
	c.Issues = append(c.Issues, issue)
	// Si hay issues, el estado no puede ser válido
	if c.Status == value_object.ValidationStatusValid {
		c.Status = value_object.ValidationStatusWarning
	}
}

// AddInvalidExample agrega un ejemplo de valor inválido
func (c *ColumnValidation) AddInvalidExample(row int, value, reason string) {
	c.InvalidExamples = append(c.InvalidExamples, InvalidExample{
		Row:    row,
		Value:  value,
		Reason: reason,
	})
	c.Statistics.InvalidCount++
	c.updateStatus()
}

// AddSampleValue agrega un valor de ejemplo
func (c *ColumnValidation) AddSampleValue(value string) {
	if len(c.SampleValues) < 5 { // Limitar a 5 ejemplos
		c.SampleValues = append(c.SampleValues, value)
	}
}

// UpdateStatistics actualiza las estadísticas de la columna
func (c *ColumnValidation) UpdateStatistics(validCount, invalidCount, nullCount, uniqueCount, totalCount int) {
	c.Statistics = ColumnStatistics{
		ValidCount:   validCount,
		InvalidCount: invalidCount,
		NullCount:    nullCount,
		UniqueCount:  uniqueCount,
		TotalCount:   totalCount,
		ValidRate:    float64(validCount) / float64(totalCount) * 100,
	}
	c.updateStatus()
}

// updateStatus actualiza el estado basado en las estadísticas
func (c *ColumnValidation) updateStatus() {
	if c.Statistics.InvalidCount == 0 && c.MappedTo != "" {
		c.Status = value_object.ValidationStatusValid
	} else if c.Statistics.InvalidCount > 0 && c.Required {
		c.Status = value_object.ValidationStatusError
	} else if c.Statistics.InvalidCount > 0 || c.MappedTo == "" {
		c.Status = value_object.ValidationStatusWarning
	} else {
		c.Status = value_object.ValidationStatusInfo
	}
}

// IsValid retorna true si la columna es válida
func (c *ColumnValidation) IsValid() bool {
	return c.Status.IsValid() && (c.MappedTo != "" || !c.Required)
}
