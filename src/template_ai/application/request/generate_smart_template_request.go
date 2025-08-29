package request

import (
	"errors"

	"github.com/gofrs/uuid/v5"
)

// GenerateSmartTemplateRequest represents the request to generate a smart template
type GenerateSmartTemplateRequest struct {
	BusinessTypeID      string                 `json:"business_type_id" binding:"required" example:"550e8400-e29b-41d4-a716-446655440001"`
	TenantID            *string                `json:"tenant_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440002"`
	Name                string                 `json:"name" binding:"required" example:"Plantilla Kiosco Premium"`
	Description         string                 `json:"description" binding:"required" example:"Plantilla optimizada para kioscos con productos premium"`
	ProductCount        int                    `json:"product_count" binding:"required,min=10,max=500" example:"100"`
	BudgetRange         *BudgetRangeRequest    `json:"budget_range,omitempty"`
	RegionalPreferences map[string]interface{} `json:"regional_preferences,omitempty"`
	CategoryPriorities  []string               `json:"category_priorities,omitempty" example:"[\"bebidas\", \"snacks\", \"golosinas\"]"`
	BrandPreferences    map[string]float64     `json:"brand_preferences,omitempty"`
	CustomConstraints   map[string]interface{} `json:"custom_constraints,omitempty"`
	OptimizationGoals   []string               `json:"optimization_goals,omitempty" example:"[\"maximize_variety\", \"optimize_margins\"]"`
}

// BudgetRangeRequest represents budget constraints
type BudgetRangeRequest struct {
	Min      float64 `json:"min" binding:"required,min=0" example:"5000"`
	Max      float64 `json:"max" binding:"required,gtfield=Min" example:"50000"`
	Currency string  `json:"currency" binding:"required" example:"ARS"`
}

// Validate validates the request
func (r *GenerateSmartTemplateRequest) Validate() error {
	// Validate business type ID
	if _, err := uuid.FromString(r.BusinessTypeID); err != nil {
		return errors.New("invalid business type ID format")
	}

	// Validate tenant ID if provided
	if r.TenantID != nil && *r.TenantID != "" {
		if _, err := uuid.FromString(*r.TenantID); err != nil {
			return errors.New("invalid tenant ID format")
		}
	}

	// Validate name
	if len(r.Name) < 3 {
		return errors.New("template name must be at least 3 characters")
	}
	if len(r.Name) > 100 {
		return errors.New("template name cannot exceed 100 characters")
	}

	// Validate description
	if len(r.Description) < 10 {
		return errors.New("template description must be at least 10 characters")
	}
	if len(r.Description) > 500 {
		return errors.New("template description cannot exceed 500 characters")
	}

	// Validate product count
	if r.ProductCount < 10 || r.ProductCount > 500 {
		return errors.New("product count must be between 10 and 500")
	}

	// Validate budget range if provided
	if r.BudgetRange != nil {
		if r.BudgetRange.Min < 0 {
			return errors.New("minimum budget cannot be negative")
		}
		if r.BudgetRange.Max <= r.BudgetRange.Min {
			return errors.New("maximum budget must be greater than minimum budget")
		}
		if r.BudgetRange.Currency == "" {
			return errors.New("currency is required when budget range is specified")
		}
	}

	// Validate brand preferences
	if r.BrandPreferences != nil {
		totalWeight := 0.0
		for brand, weight := range r.BrandPreferences {
			if weight < 0 || weight > 1 {
				return errors.New("brand preference weights must be between 0 and 1")
			}
			if brand == "" {
				return errors.New("brand name cannot be empty")
			}
			totalWeight += weight
		}
		// Allow some tolerance for floating point arithmetic
		if totalWeight > 1.01 {
			return errors.New("total brand preference weights cannot exceed 1")
		}
	}

	// Validate category priorities
	if len(r.CategoryPriorities) > 20 {
		return errors.New("cannot specify more than 20 category priorities")
	}

	// Validate optimization goals
	validGoals := map[string]bool{
		"maximize_variety":     true,
		"optimize_margins":     true,
		"balance_categories":   true,
		"focus_essentials":     true,
		"premium_selection":    true,
		"budget_optimization":  true,
		"regional_preferences": true,
	}

	for _, goal := range r.OptimizationGoals {
		if !validGoals[goal] {
			return errors.New("invalid optimization goal: " + goal)
		}
	}

	return nil
}