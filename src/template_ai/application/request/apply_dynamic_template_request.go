package request

import (
	"errors"

	"github.com/gofrs/uuid/v5"
)

// ApplyDynamicTemplateRequest represents the request to apply a dynamic template
type ApplyDynamicTemplateRequest struct {
	TemplateID          string                 `json:"template_id" binding:"required" example:"550e8400-e29b-41d4-a716-446655440001"`
	TenantID            string                 `json:"tenant_id" binding:"required" example:"550e8400-e29b-41d4-a716-446655440002"`
	ApplyMode           string                 `json:"apply_mode" binding:"required,oneof=full partial merge" example:"full"`
	ProductSelections   []ProductSelection     `json:"product_selections,omitempty"`
	OverrideQuantities  map[string]int         `json:"override_quantities,omitempty"`
	ExcludeProducts     []string               `json:"exclude_products,omitempty"`
	CategoryAdjustments map[string]float64     `json:"category_adjustments,omitempty"`
	CreateVariants      bool                   `json:"create_variants" example:"true"`
	ApplyPricing        bool                   `json:"apply_pricing" example:"true"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
}

// ProductSelection represents manual product selection override
type ProductSelection struct {
	GlobalProductID string `json:"global_product_id" binding:"required" example:"550e8400-e29b-41d4-a716-446655440003"`
	Quantity        int    `json:"quantity" binding:"required,min=1" example:"10"`
	Priority        int    `json:"priority" binding:"required,min=1,max=3" example:"1"`
}

// Validate validates the request
func (r *ApplyDynamicTemplateRequest) Validate() error {
	// Validate template ID
	if _, err := uuid.FromString(r.TemplateID); err != nil {
		return errors.New("invalid template ID format")
	}

	// Validate tenant ID
	if _, err := uuid.FromString(r.TenantID); err != nil {
		return errors.New("invalid tenant ID format")
	}

	// Validate apply mode
	validModes := map[string]bool{
		"full":    true,
		"partial": true,
		"merge":   true,
	}
	if !validModes[r.ApplyMode] {
		return errors.New("invalid apply mode, must be one of: full, partial, merge")
	}

	// Validate product selections
	for i, selection := range r.ProductSelections {
		if _, err := uuid.FromString(selection.GlobalProductID); err != nil {
			return errors.New("invalid global product ID format in product selection")
		}
		if selection.Quantity < 1 {
			return errors.New("product quantity must be at least 1")
		}
		if selection.Priority < 1 || selection.Priority > 3 {
			return errors.New("product priority must be between 1 and 3")
		}
		// Check for duplicates
		for j := i + 1; j < len(r.ProductSelections); j++ {
			if selection.GlobalProductID == r.ProductSelections[j].GlobalProductID {
				return errors.New("duplicate product in selections: " + selection.GlobalProductID)
			}
		}
	}

	// Validate override quantities
	for productID, quantity := range r.OverrideQuantities {
		if _, err := uuid.FromString(productID); err != nil {
			return errors.New("invalid product ID format in override quantities")
		}
		if quantity < 0 {
			return errors.New("override quantity cannot be negative")
		}
	}

	// Validate exclude products
	for _, productID := range r.ExcludeProducts {
		if _, err := uuid.FromString(productID); err != nil {
			return errors.New("invalid product ID format in exclude products")
		}
	}

	// Validate category adjustments
	for category, adjustment := range r.CategoryAdjustments {
		if category == "" {
			return errors.New("category name cannot be empty in adjustments")
		}
		if adjustment < 0 || adjustment > 2 {
			return errors.New("category adjustment must be between 0 and 2 (0% to 200%)")
		}
	}

	return nil
}

// ApplyModeType represents the type of template application
type ApplyModeType string

const (
	ApplyModeFull    ApplyModeType = "full"    // Replace all existing products
	ApplyModePartial ApplyModeType = "partial" // Add only new products
	ApplyModeMerge   ApplyModeType = "merge"   // Merge with existing, update quantities
)

// GetApplyMode returns the typed apply mode
func (r *ApplyDynamicTemplateRequest) GetApplyMode() ApplyModeType {
	return ApplyModeType(r.ApplyMode)
}