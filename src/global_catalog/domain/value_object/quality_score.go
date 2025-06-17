package value_object

import (
	"errors"
	"math"
)

// QualityScore representa una puntuación de calidad de 0 a 100
type QualityScore struct {
	value int
}

// QualityMetrics contiene las métricas usadas para calcular el quality score
type QualityMetrics struct {
	HasName           bool
	HasDescription    bool
	HasImage          bool
	HasPrice          bool
	HasBrand          bool
	HasCategory       bool
	HasEAN            bool
	HasStock          bool
	ImageCount        int
	DescriptionLength int
}

// NewQualityScore crea un nuevo QualityScore validando el rango
func NewQualityScore(value int) (*QualityScore, error) {
	if value < 0 || value > 100 {
		return nil, errors.New("el quality score debe estar entre 0 y 100")
	}

	return &QualityScore{value: value}, nil
}

// NewQualityScoreFromMetrics calcula automáticamente el score basado en métricas
func NewQualityScoreFromMetrics(metrics QualityMetrics) (*QualityScore, error) {
	score := calculateQualityScore(metrics)
	return NewQualityScore(score)
}

// Value retorna el valor del quality score
func (q *QualityScore) Value() int {
	return q.value
}

// Percentage retorna el score como porcentaje (0.0 - 1.0)
func (q *QualityScore) Percentage() float64 {
	return float64(q.value) / 100.0
}

// Level retorna el nivel de calidad como string
func (q *QualityScore) Level() string {
	switch {
	case q.value >= 90:
		return "Excelente"
	case q.value >= 70:
		return "Buena"
	case q.value >= 50:
		return "Regular"
	case q.value >= 30:
		return "Baja"
	default:
		return "Muy Baja"
	}
}

// IsHighQuality indica si el producto tiene alta calidad (>= 70)
func (q *QualityScore) IsHighQuality() bool {
	return q.value >= 70
}

// IsLowQuality indica si el producto tiene baja calidad (< 50)
func (q *QualityScore) IsLowQuality() bool {
	return q.value < 50
}

// calculateQualityScore calcula el score basado en las métricas del producto
func calculateQualityScore(metrics QualityMetrics) int {
	score := 0.0

	// Campos básicos obligatorios (60% del score)
	if metrics.HasName {
		score += 15.0
	}
	if metrics.HasEAN {
		score += 20.0 // EAN es crítico
	}
	if metrics.HasDescription {
		score += 15.0
		// Bonus por descripción detallada
		if metrics.DescriptionLength > 100 {
			score += 5.0
		}
	}
	if metrics.HasPrice {
		score += 10.0
	}

	// Campos de clasificación (25% del score)
	if metrics.HasCategory {
		score += 10.0
	}
	if metrics.HasBrand {
		score += 10.0
	}
	if metrics.HasStock {
		score += 5.0
	}

	// Multimedia (15% del score)
	if metrics.HasImage {
		score += 10.0
		// Bonus por múltiples imágenes
		if metrics.ImageCount > 1 {
			score += 5.0
		}
	}

	// Redondear y asegurar que esté en el rango válido
	finalScore := int(math.Round(score))
	if finalScore > 100 {
		finalScore = 100
	}
	if finalScore < 0 {
		finalScore = 0
	}

	return finalScore
}

// Equals compara dos QualityScore
func (q *QualityScore) Equals(other *QualityScore) bool {
	if other == nil {
		return false
	}
	return q.value == other.value
}

// String implementa la interfaz Stringer
func (q *QualityScore) String() string {
	return q.Level()
}
