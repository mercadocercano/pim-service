package response

import "time"

// CategoryHierarchyValidationResponse representa la respuesta de validación de jerarquía
type CategoryHierarchyValidationResponse struct {
	IsValid          bool                    `json:"is_valid"`
	CategoryID       string                  `json:"category_id"`
	CurrentParentID  *string                 `json:"current_parent_id"`
	NewParentID      *string                 `json:"new_parent_id"`
	CurrentLevel     int                     `json:"current_level"`
	NewLevel         int                     `json:"new_level"`
	MaxDepth         int                     `json:"max_depth"`
	ValidationErrors []ValidationError       `json:"validation_errors,omitempty"`
	AffectedChildren []AffectedChildCategory `json:"affected_children,omitempty"`
	CategoryPath     []CategoryPathItem      `json:"category_path"`
	ValidatedAt      time.Time               `json:"validated_at"`
}

// ValidationError representa un error específico de validación
type ValidationError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Field      string `json:"field,omitempty"`
	CategoryID string `json:"category_id,omitempty"`
}

// AffectedChildCategory representa una categoría hija que sería afectada por el cambio
type AffectedChildCategory struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CurrentLevel int    `json:"current_level"`
	NewLevel     int    `json:"new_level"`
	WouldExceed  bool   `json:"would_exceed_max_depth"`
}

// CategoryPathItem representa un elemento en el path de la categoría
type CategoryPathItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// HierarchyValidationSummary representa un resumen de la validación
type HierarchyValidationSummary struct {
	TotalCategoriesValidated int                                    `json:"total_categories_validated"`
	ValidCategories          int                                    `json:"valid_categories"`
	InvalidCategories        int                                    `json:"invalid_categories"`
	ValidationResults        []*CategoryHierarchyValidationResponse `json:"validation_results"`
	GeneratedAt              time.Time                              `json:"generated_at"`
}
