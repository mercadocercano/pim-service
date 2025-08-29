package request

import (
	"errors"

	"github.com/gofrs/uuid/v5"
)

// UpdateTemplateFromFeedbackRequest represents the request to update a template based on feedback
type UpdateTemplateFromFeedbackRequest struct {
	TemplateID            string                 `json:"template_id" binding:"required" example:"550e8400-e29b-41d4-a716-446655440001"`
	FeedbackPeriodDays    int                    `json:"feedback_period_days,omitempty" example:"30"`
	MinFeedbackCount      int                    `json:"min_feedback_count,omitempty" example:"10"`
	UpdateStrategy        string                 `json:"update_strategy" binding:"required,oneof=incremental full regenerate" example:"incremental"`
	PreserveCustomizations bool                  `json:"preserve_customizations" example:"true"`
	OptimizationTargets   []OptimizationTarget   `json:"optimization_targets,omitempty"`
	ExcludeProducts       []string               `json:"exclude_products,omitempty"`
	ForceUpdate           bool                   `json:"force_update" example:"false"`
	TestMode              bool                   `json:"test_mode" example:"false"`
	UpdateReason          string                 `json:"update_reason,omitempty" example:"Monthly optimization based on user feedback"`
}

// OptimizationTarget represents a specific optimization goal
type OptimizationTarget struct {
	MetricType  string  `json:"metric_type" binding:"required" example:"satisfaction_score"`
	TargetValue float64 `json:"target_value" binding:"required,min=0,max=1" example:"0.85"`
	Weight      float64 `json:"weight,omitempty" example:"0.5"`
}

// Validate validates the request
func (r *UpdateTemplateFromFeedbackRequest) Validate() error {
	// Validate template ID
	if _, err := uuid.FromString(r.TemplateID); err != nil {
		return errors.New("invalid template ID format")
	}

	// Validate feedback period
	if r.FeedbackPeriodDays < 0 {
		return errors.New("feedback period days cannot be negative")
	}
	if r.FeedbackPeriodDays == 0 {
		r.FeedbackPeriodDays = 30 // Default to 30 days
	}
	if r.FeedbackPeriodDays > 365 {
		return errors.New("feedback period cannot exceed 365 days")
	}

	// Validate minimum feedback count
	if r.MinFeedbackCount < 0 {
		return errors.New("minimum feedback count cannot be negative")
	}
	if r.MinFeedbackCount == 0 {
		r.MinFeedbackCount = 10 // Default minimum
	}

	// Validate update strategy
	validStrategies := map[string]bool{
		"incremental": true, // Make small adjustments based on feedback
		"full":        true, // Complete reoptimization maintaining structure
		"regenerate":  true, // Full AI regeneration from scratch
	}
	if !validStrategies[r.UpdateStrategy] {
		return errors.New("invalid update strategy")
	}

	// Validate optimization targets
	if len(r.OptimizationTargets) > 0 {
		validMetrics := map[string]bool{
			"satisfaction_score": true,
			"modification_rate":  true,
			"adoption_rate":      true,
			"product_retention":  true,
			"category_balance":   true,
			"brand_diversity":    true,
		}

		totalWeight := 0.0
		for _, target := range r.OptimizationTargets {
			if !validMetrics[target.MetricType] {
				return errors.New("invalid metric type in optimization target: " + target.MetricType)
			}
			if target.TargetValue < 0 || target.TargetValue > 1 {
				return errors.New("target value must be between 0 and 1")
			}
			if target.Weight < 0 || target.Weight > 1 {
				return errors.New("weight must be between 0 and 1")
			}
			totalWeight += target.Weight
		}

		// Check if weights sum to approximately 1 (with tolerance for floating point)
		if len(r.OptimizationTargets) > 1 && totalWeight > 0 && (totalWeight < 0.99 || totalWeight > 1.01) {
			return errors.New("optimization target weights must sum to 1")
		}
	}

	// Validate exclude products
	for _, productID := range r.ExcludeProducts {
		if _, err := uuid.FromString(productID); err != nil {
			return errors.New("invalid product ID format in exclude products: " + productID)
		}
	}

	// Validate update reason length if provided
	if r.UpdateReason != "" && len(r.UpdateReason) > 500 {
		return errors.New("update reason cannot exceed 500 characters")
	}

	return nil
}

// UpdateStrategyType represents the type of update strategy
type UpdateStrategyType string

const (
	UpdateStrategyIncremental UpdateStrategyType = "incremental"
	UpdateStrategyFull        UpdateStrategyType = "full"
	UpdateStrategyRegenerate  UpdateStrategyType = "regenerate"
)

// GetUpdateStrategy returns the typed update strategy
func (r *UpdateTemplateFromFeedbackRequest) GetUpdateStrategy() UpdateStrategyType {
	return UpdateStrategyType(r.UpdateStrategy)
}