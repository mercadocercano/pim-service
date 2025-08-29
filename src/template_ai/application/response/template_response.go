package response

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

// AITemplateResponse represents the response for an AI template
type AITemplateResponse struct {
	ID                    uuid.UUID              `json:"id"`
	BusinessTypeID        uuid.UUID              `json:"business_type_id"`
	BusinessTypeName      string                 `json:"business_type_name"`
	TenantID              *uuid.UUID             `json:"tenant_id,omitempty"`
	Name                  string                 `json:"name"`
	Description           string                 `json:"description"`
	GeneratedBy           string                 `json:"generated_by"`
	ProductCount          int                    `json:"product_count"`
	EstimatedValue        *float64               `json:"estimated_value,omitempty"`
	AIGenerationParams    map[string]interface{} `json:"ai_generation_params,omitempty"`
	PerformanceMetrics    map[string]interface{} `json:"performance_metrics,omitempty"`
	LastAIUpdate          *time.Time             `json:"last_ai_update,omitempty"`
	ProductSelectionRules map[string]interface{} `json:"product_selection_rules,omitempty"`
	CategoryDistribution  map[string]interface{} `json:"category_distribution,omitempty"`
	BrandPreferences      map[string]interface{} `json:"brand_preferences,omitempty"`
	RegionalAdaptations   map[string]interface{} `json:"regional_adaptations,omitempty"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

// TemplateProductResponse represents a product in a template
type TemplateProductResponse struct {
	ID                      uuid.UUID                `json:"id"`
	GlobalProductID         uuid.UUID                `json:"global_product_id"`
	ProductDetails          *ProductDetailsResponse  `json:"product_details,omitempty"`
	Priority                int                      `json:"priority"`
	PriorityLabel           string                   `json:"priority_label"`
	QuantitySuggestion      int                      `json:"quantity_suggestion"`
	AIReasoning             string                   `json:"ai_reasoning"`
	RelevanceScore          float64                  `json:"relevance_score"`
	CategoryMatchScore      *float64                 `json:"category_match_score,omitempty"`
	BrandMatchScore         *float64                 `json:"brand_match_score,omitempty"`
	RegionalPreferenceScore *float64                 `json:"regional_preference_score,omitempty"`
	CreatedAt               time.Time                `json:"created_at"`
	UpdatedAt               time.Time                `json:"updated_at"`
}

// ProductDetailsResponse represents basic product information
type ProductDetailsResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	SKU         string    `json:"sku,omitempty"`
	Category    string    `json:"category,omitempty"`
	Brand       string    `json:"brand,omitempty"`
	Price       *float64  `json:"price,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
}

// GenerateSmartTemplateResponse represents the response for template generation
type GenerateSmartTemplateResponse struct {
	Template          *AITemplateResponse        `json:"template"`
	Products          []TemplateProductResponse  `json:"products"`
	GenerationSummary *GenerationSummaryResponse `json:"generation_summary"`
	Success           bool                       `json:"success"`
	Message           string                     `json:"message,omitempty"`
}

// GenerationSummaryResponse contains summary of the generation process
type GenerationSummaryResponse struct {
	TotalProducts        int                    `json:"total_products"`
	EssentialProducts    int                    `json:"essential_products"`
	RecommendedProducts  int                    `json:"recommended_products"`
	OptionalProducts     int                    `json:"optional_products"`
	TotalEstimatedValue  float64                `json:"total_estimated_value"`
	CategoryBreakdown    map[string]int         `json:"category_breakdown"`
	BrandBreakdown       map[string]int         `json:"brand_breakdown"`
	GenerationTimeMs     int                    `json:"generation_time_ms"`
	AIModel              string                 `json:"ai_model"`
	OptimizationScore    float64                `json:"optimization_score"`
}

// ApplyDynamicTemplateResponse represents the response for template application
type ApplyDynamicTemplateResponse struct {
	Success              bool                         `json:"success"`
	Message              string                       `json:"message"`
	AppliedProducts      int                          `json:"applied_products"`
	SkippedProducts      int                          `json:"skipped_products"`
	FailedProducts       int                          `json:"failed_products"`
	CreatedVariants      int                          `json:"created_variants,omitempty"`
	ApplicationDetails   *ApplicationDetailsResponse  `json:"application_details"`
	Errors               []ApplicationErrorResponse   `json:"errors,omitempty"`
}

// ApplicationDetailsResponse contains details of the application process
type ApplicationDetailsResponse struct {
	TemplateID       uuid.UUID              `json:"template_id"`
	TemplateName     string                 `json:"template_name"`
	TenantID         uuid.UUID              `json:"tenant_id"`
	ApplyMode        string                 `json:"apply_mode"`
	StartedAt        time.Time              `json:"started_at"`
	CompletedAt      time.Time              `json:"completed_at"`
	DurationMs       int                    `json:"duration_ms"`
	ProductsCreated  []string               `json:"products_created,omitempty"`
	ProductsUpdated  []string               `json:"products_updated,omitempty"`
	ProductsSkipped  []string               `json:"products_skipped,omitempty"`
}

// ApplicationErrorResponse represents an error during application
type ApplicationErrorResponse struct {
	ProductID   string    `json:"product_id,omitempty"`
	ProductName string    `json:"product_name,omitempty"`
	Error       string    `json:"error"`
	ErrorType   string    `json:"error_type"`
	Timestamp   time.Time `json:"timestamp"`
}

// TemplatePerformanceResponse represents the performance analysis response
type TemplatePerformanceResponse struct {
	TemplateID           uuid.UUID                         `json:"template_id"`
	TemplateName         string                            `json:"template_name"`
	AnalysisPeriod       *PeriodResponse                   `json:"analysis_period"`
	Metrics              map[string]MetricValueResponse    `json:"metrics"`
	Trends               map[string][]TrendPointResponse   `json:"trends,omitempty"`
	Comparisons          []TemplateComparisonResponse      `json:"comparisons,omitempty"`
	Recommendations      []RecommendationResponse          `json:"recommendations"`
	PerformanceScore     float64                           `json:"performance_score"`
	PerformanceRating    string                            `json:"performance_rating"`
}

// PeriodResponse represents a time period
type PeriodResponse struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Days  int       `json:"days"`
}

// MetricValueResponse represents a metric value
type MetricValueResponse struct {
	Value       float64                `json:"value"`
	Change      *float64               `json:"change,omitempty"`
	ChangeType  string                 `json:"change_type,omitempty"` // "increase", "decrease", "stable"
	Target      *float64               `json:"target,omitempty"`
	Status      string                 `json:"status"` // "good", "warning", "critical"
	Details     map[string]interface{} `json:"details,omitempty"`
}

// TrendPointResponse represents a point in a trend
type TrendPointResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
	Label     string    `json:"label,omitempty"`
}

// TemplateComparisonResponse represents comparison with another template
type TemplateComparisonResponse struct {
	TemplateID   uuid.UUID              `json:"template_id"`
	TemplateName string                 `json:"template_name"`
	Metrics      map[string]float64     `json:"metrics"`
	Difference   map[string]float64     `json:"difference"`
	Winner       map[string]string      `json:"winner"`
}

// RecommendationResponse represents a performance recommendation
type RecommendationResponse struct {
	Type        string   `json:"type"` // "improvement", "warning", "opportunity"
	Priority    string   `json:"priority"` // "high", "medium", "low"
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Action      string   `json:"action"`
	Impact      string   `json:"impact,omitempty"`
	Metrics     []string `json:"affected_metrics,omitempty"`
}

// UpdateTemplateFromFeedbackResponse represents the response for template update
type UpdateTemplateFromFeedbackResponse struct {
	Success          bool                       `json:"success"`
	Message          string                     `json:"message"`
	UpdatedTemplate  *AITemplateResponse        `json:"updated_template,omitempty"`
	UpdateSummary    *UpdateSummaryResponse     `json:"update_summary"`
	TestResults      *TestResultsResponse       `json:"test_results,omitempty"`
}

// UpdateSummaryResponse contains summary of the update process
type UpdateSummaryResponse struct {
	UpdateStrategy       string                 `json:"update_strategy"`
	FeedbackAnalyzed     int                    `json:"feedback_analyzed"`
	ProductsAdded        int                    `json:"products_added"`
	ProductsRemoved      int                    `json:"products_removed"`
	ProductsModified     int                    `json:"products_modified"`
	QuantitiesAdjusted   int                    `json:"quantities_adjusted"`
	PreviousMetrics      map[string]float64     `json:"previous_metrics"`
	ProjectedMetrics     map[string]float64     `json:"projected_metrics"`
	ImprovementAreas     []string               `json:"improvement_areas"`
	UpdateReason         string                 `json:"update_reason,omitempty"`
	UpdatedAt            time.Time              `json:"updated_at"`
}

// TestResultsResponse contains results from test mode
type TestResultsResponse struct {
	SimulatedChanges     []SimulatedChangeResponse `json:"simulated_changes"`
	ProjectedImprovement map[string]float64        `json:"projected_improvement"`
	RiskAssessment       string                    `json:"risk_assessment"`
	Recommendation       string                    `json:"recommendation"`
}

// SimulatedChangeResponse represents a simulated change
type SimulatedChangeResponse struct {
	ChangeType  string                 `json:"change_type"`
	ProductID   string                 `json:"product_id,omitempty"`
	ProductName string                 `json:"product_name,omitempty"`
	OldValue    interface{}            `json:"old_value,omitempty"`
	NewValue    interface{}            `json:"new_value,omitempty"`
	Reason      string                 `json:"reason"`
	Impact      map[string]interface{} `json:"impact,omitempty"`
}