package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/lib/pq"
)

// AITemplate represents an AI-generated or AI-enhanced template
type AITemplate struct {
	ID                    uuid.UUID              `json:"id"`
	BusinessTypeID        uuid.UUID              `json:"business_type_id"`
	TenantID              *uuid.UUID             `json:"tenant_id,omitempty"`
	Name                  string                 `json:"name"`
	Description           string                 `json:"description"`
	GeneratedBy           string                 `json:"generated_by"` // 'ai', 'manual', 'hybrid'
	AIGenerationParams    map[string]interface{} `json:"ai_generation_params,omitempty"`
	PerformanceMetrics    map[string]interface{} `json:"performance_metrics,omitempty"`
	LastAIUpdate          *time.Time             `json:"last_ai_update,omitempty"`
	ProductSelectionRules map[string]interface{} `json:"product_selection_rules,omitempty"`
	CategoryDistribution  map[string]interface{} `json:"category_distribution,omitempty"`
	BrandPreferences      map[string]interface{} `json:"brand_preferences,omitempty"`
	RegionalAdaptations   map[string]interface{} `json:"regional_adaptations,omitempty"`
	Version               int                    `json:"version"`
	Products              []TemplateGlobalProduct `json:"products,omitempty"`
	GenerationParams      map[string]interface{} `json:"generation_params,omitempty"`
	UsageCount            int                    `json:"usage_count"`
	SatisfactionScore     float64                `json:"satisfaction_score"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

// NewAITemplate creates a new AI template
func NewAITemplate(
	businessTypeID uuid.UUID,
	tenantID *uuid.UUID,
	name string,
	description string,
	generatedBy string,
) (*AITemplate, error) {
	newID, _ := uuid.NewV4()
	return &AITemplate{
		ID:                 newID,
		BusinessTypeID:     businessTypeID,
		TenantID:           tenantID,
		Name:               name,
		Description:        description,
		GeneratedBy:        generatedBy,
		AIGenerationParams: make(map[string]interface{}),
		PerformanceMetrics: make(map[string]interface{}),
		Version:            1,
		UsageCount:         0,
		SatisfactionScore:  0.0,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}, nil
}

// TemplateGlobalProduct represents the link between templates and global products
type TemplateGlobalProduct struct {
	ID                      uuid.UUID  `json:"id"`
	TemplateID              uuid.UUID  `json:"template_id"`
	GlobalProductID         uuid.UUID  `json:"global_product_id"`
	Priority                int        `json:"priority"` // 1=essential, 2=recommended, 3=optional
	QuantitySuggestion      int        `json:"quantity_suggestion"`
	AIReasoning             string     `json:"ai_reasoning"`
	RelevanceScore          float64    `json:"relevance_score"`
	CategoryMatchScore      *float64   `json:"category_match_score,omitempty"`
	BrandMatchScore         *float64   `json:"brand_match_score,omitempty"`
	RegionalPreferenceScore *float64   `json:"regional_preference_score,omitempty"`
	CreatedAt               time.Time  `json:"created_at"`
	UpdatedAt               time.Time  `json:"updated_at"`
}

// AIGenerationHistory tracks AI generation attempts
type AIGenerationHistory struct {
	ID               uuid.UUID              `json:"id"`
	TemplateID       *uuid.UUID             `json:"template_id,omitempty"`
	TenantID         *uuid.UUID             `json:"tenant_id,omitempty"`
	BusinessTypeID   uuid.UUID              `json:"business_type_id"`
	GenerationParams map[string]interface{} `json:"generation_params"`
	AIModel          string                 `json:"ai_model"`
	PromptTemplate   string                 `json:"prompt_template"`
	GeneratedContent map[string]interface{} `json:"generated_content"`
	GenerationStatus string                 `json:"generation_status"` // pending, processing, completed, failed
	ErrorMessage     *string                `json:"error_message,omitempty"`
	GenerationTimeMs *int                   `json:"generation_time_ms,omitempty"`
	AppliedAt        *time.Time             `json:"applied_at,omitempty"`
	CreatedAt        time.Time              `json:"created_at"`
}

// AIPerformanceMetric tracks template performance
type AIPerformanceMetric struct {
	ID             uuid.UUID              `json:"id"`
	TemplateID     uuid.UUID              `json:"template_id"`
	MetricType     string                 `json:"metric_type"` // usage_count, satisfaction_score, modification_rate
	MetricValue    float64                `json:"metric_value"`
	MetricMetadata map[string]interface{} `json:"metric_metadata,omitempty"`
	PeriodStart    time.Time              `json:"period_start"`
	PeriodEnd      time.Time              `json:"period_end"`
	CreatedAt      time.Time              `json:"created_at"`
}

// AIProductFeedback captures user modifications to AI suggestions
type AIProductFeedback struct {
	ID                   uuid.UUID  `json:"id"`
	TenantID             uuid.UUID  `json:"tenant_id"`
	TemplateID           *uuid.UUID `json:"template_id,omitempty"`
	GlobalProductID      *uuid.UUID `json:"global_product_id,omitempty"`
	Action               string     `json:"action"` // 'kept', 'removed', 'quantity_changed', 'replaced'
	OriginalQuantity     *int       `json:"original_quantity,omitempty"`
	NewQuantity          *int       `json:"new_quantity,omitempty"`
	ReplacementProductID *uuid.UUID `json:"replacement_product_id,omitempty"`
	FeedbackReason       *string    `json:"feedback_reason,omitempty"`
	CreatedAt            time.Time  `json:"created_at"`
}

// GenerationParams represents parameters for AI template generation
type GenerationParams struct {
	BusinessTypeID      uuid.UUID              `json:"business_type_id"`
	TenantID            *uuid.UUID             `json:"tenant_id,omitempty"`
	ProductCount        int                    `json:"product_count"`
	BudgetRange         *BudgetRange           `json:"budget_range,omitempty"`
	RegionalPreferences map[string]interface{} `json:"regional_preferences,omitempty"`
	CategoryPriorities  pq.StringArray         `json:"category_priorities,omitempty"`
	BrandPreferences    map[string]float64     `json:"brand_preferences,omitempty"`
	CustomConstraints   map[string]interface{} `json:"custom_constraints,omitempty"`
}

// BudgetRange represents a budget constraint
type BudgetRange struct {
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
	Currency string  `json:"currency"`
}