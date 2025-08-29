package request

import (
	"errors"
	"time"

	"github.com/gofrs/uuid/v5"
)

// AnalyzeTemplatePerformanceRequest represents the request to analyze template performance
type AnalyzeTemplatePerformanceRequest struct {
	TemplateID   string                `json:"template_id" binding:"required" example:"550e8400-e29b-41d4-a716-446655440001"`
	PeriodStart  *time.Time            `json:"period_start,omitempty" example:"2024-01-01T00:00:00Z"`
	PeriodEnd    *time.Time            `json:"period_end,omitempty" example:"2024-12-31T23:59:59Z"`
	MetricTypes  []string              `json:"metric_types,omitempty" example:"[\"usage_count\", \"satisfaction_score\", \"modification_rate\"]"`
	GroupBy      string                `json:"group_by,omitempty" example:"month"`
	IncludeDetails bool                `json:"include_details" example:"true"`
	CompareWith  []string              `json:"compare_with,omitempty" example:"[\"550e8400-e29b-41d4-a716-446655440002\"]"`
}

// Validate validates the request
func (r *AnalyzeTemplatePerformanceRequest) Validate() error {
	// Validate template ID
	if _, err := uuid.FromString(r.TemplateID); err != nil {
		return errors.New("invalid template ID format")
	}

	// Validate period
	if r.PeriodStart != nil && r.PeriodEnd != nil {
		if r.PeriodStart.After(*r.PeriodEnd) {
			return errors.New("period start must be before period end")
		}
		// Check if period is not too long (e.g., max 2 years)
		if r.PeriodEnd.Sub(*r.PeriodStart) > 2*365*24*time.Hour {
			return errors.New("analysis period cannot exceed 2 years")
		}
	}

	// Validate metric types if specified
	if len(r.MetricTypes) > 0 {
		validMetrics := map[string]bool{
			"usage_count":         true,
			"satisfaction_score":  true,
			"modification_rate":   true,
			"adoption_rate":       true,
			"product_retention":   true,
			"average_order_value": true,
			"category_balance":    true,
			"brand_diversity":     true,
		}

		for _, metric := range r.MetricTypes {
			if !validMetrics[metric] {
				return errors.New("invalid metric type: " + metric)
			}
		}
	}

	// Validate group by
	if r.GroupBy != "" {
		validGroupBy := map[string]bool{
			"day":     true,
			"week":    true,
			"month":   true,
			"quarter": true,
			"year":    true,
		}
		if !validGroupBy[r.GroupBy] {
			return errors.New("invalid group by value, must be one of: day, week, month, quarter, year")
		}
	}

	// Validate compare with template IDs
	if len(r.CompareWith) > 0 {
		if len(r.CompareWith) > 5 {
			return errors.New("cannot compare with more than 5 templates")
		}
		
		seenIDs := make(map[string]bool)
		for _, templateID := range r.CompareWith {
			if _, err := uuid.FromString(templateID); err != nil {
				return errors.New("invalid template ID format in compare_with: " + templateID)
			}
			if templateID == r.TemplateID {
				return errors.New("cannot compare template with itself")
			}
			if seenIDs[templateID] {
				return errors.New("duplicate template ID in compare_with: " + templateID)
			}
			seenIDs[templateID] = true
		}
	}

	return nil
}

// GetPeriodOrDefault returns the period with defaults if not specified
func (r *AnalyzeTemplatePerformanceRequest) GetPeriodOrDefault() (start, end time.Time) {
	now := time.Now()
	
	if r.PeriodEnd != nil {
		end = *r.PeriodEnd
	} else {
		end = now
	}
	
	if r.PeriodStart != nil {
		start = *r.PeriodStart
	} else {
		// Default to last 30 days
		start = end.AddDate(0, 0, -30)
	}
	
	return start, end
}

// GetMetricTypesOrDefault returns metric types or defaults if not specified
func (r *AnalyzeTemplatePerformanceRequest) GetMetricTypesOrDefault() []string {
	if len(r.MetricTypes) > 0 {
		return r.MetricTypes
	}
	
	// Default metrics
	return []string{
		"usage_count",
		"satisfaction_score",
		"modification_rate",
		"product_retention",
	}
}