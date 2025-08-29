package service

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

// TemplateEngine defines the interface for AI-powered template generation and optimization
type TemplateEngine interface {
	// GenerateFromCatalog generates a smart template based on global catalog and parameters
	GenerateFromCatalog(ctx context.Context, params TemplateParams) (*TemplateResult, error)
	
	// OptimizeForRegion adapts a template for specific regional preferences
	OptimizeForRegion(ctx context.Context, template *TemplateResult, region string) (*TemplateResult, error)
	
	// CalculateProductMix determines the optimal product distribution for a business type
	CalculateProductMix(ctx context.Context, businessType string, targetCount int) ([]ProductSelection, error)
	
	// AnalyzeBusinessContext analyzes the business context to provide template recommendations
	AnalyzeBusinessContext(ctx context.Context, context BusinessContext) (*ContextAnalysis, error)
	
	// ValidateTemplate validates a template against business rules and constraints
	ValidateTemplate(ctx context.Context, template *TemplateResult) (*ValidationResult, error)
}

// TemplateParams contains parameters for template generation
type TemplateParams struct {
	BusinessTypeID   uuid.UUID              `json:"business_type_id"`
	TargetSize       string                 `json:"target_size"` // small, medium, large
	PriceRange       string                 `json:"price_range"` // economy, standard, premium
	IncludeGenerics  bool                   `json:"include_generics"`
	GenericPercent   int                    `json:"generic_percentage"`
	CategoryFocus    []string               `json:"categories_focus"`
	ExcludeBrands    []string               `json:"exclude_brands"`
	ExcludeProducts  []uuid.UUID            `json:"exclude_products"`
	RegionalPrefs    string                 `json:"regional_preferences"`
	MinProducts      int                    `json:"min_products"`
	MaxProducts      int                    `json:"max_products"`
	CustomRules      map[string]interface{} `json:"custom_rules"`
	OptimizationGoal string                 `json:"optimization_goal"` // variety, profit_margin, fast_moving
}

// TemplateResult represents the generated template
type TemplateResult struct {
	TemplateID         uuid.UUID                  `json:"template_id"`
	Name               string                     `json:"name"`
	Version            int                        `json:"version"`
	GeneratedBy        string                     `json:"generated_by"`
	BusinessTypeID     uuid.UUID                  `json:"business_type_id"`
	Products           []ProductRecommendation    `json:"products"`
	Categories         []CategoryDistribution     `json:"categories"`
	Brands             []BrandDistribution        `json:"brands"`
	Summary            TemplateSummary            `json:"summary"`
	OptimizationScore  float64                    `json:"optimization_score"`
	GenerationMetadata map[string]interface{}     `json:"generation_metadata"`
}

// ProductRecommendation represents a recommended product with metadata
type ProductRecommendation struct {
	GlobalProductID     uuid.UUID              `json:"global_product_id"`
	Name                string                 `json:"name"`
	CategoryID          string                 `json:"category_id"`
	CategoryName        string                 `json:"category_name"`
	BrandID             string                 `json:"brand_id"`
	BrandName           string                 `json:"brand_name"`
	Priority            int                    `json:"priority"` // 1=essential, 2=recommended, 3=optional
	QuantitySuggestion  int                    `json:"quantity_suggestion"`
	Price               float64                `json:"price"`
	RelevanceScore      float64                `json:"relevance_score"`
	AIReasoning         string                 `json:"ai_reasoning"`
	Tags                []string               `json:"tags"`
	Attributes          map[string]interface{} `json:"attributes"`
}

// ProductSelection represents a product selection for mix calculation
type ProductSelection struct {
	ProductID      uuid.UUID `json:"product_id"`
	CategoryID     string    `json:"category_id"`
	Priority       int       `json:"priority"`
	SelectionScore float64   `json:"selection_score"`
	Reason         string    `json:"reason"`
}

// CategoryDistribution shows product distribution by category
type CategoryDistribution struct {
	CategoryID   string  `json:"category_id"`
	CategoryName string  `json:"category_name"`
	ProductCount int     `json:"product_count"`
	Percentage   float64 `json:"percentage"`
	Priority     int     `json:"priority"`
}

// BrandDistribution shows product distribution by brand
type BrandDistribution struct {
	BrandID      string  `json:"brand_id"`
	BrandName    string  `json:"brand_name"`
	ProductCount int     `json:"product_count"`
	Percentage   float64 `json:"percentage"`
	BrandType    string  `json:"brand_type"` // premium, standard, generic
}

// TemplateSummary provides overview statistics
type TemplateSummary struct {
	TotalProducts        int                    `json:"total_products"`
	TotalCategories      int                    `json:"total_categories"`
	TotalBrands          int                    `json:"total_brands"`
	EstimatedInvestment  float64                `json:"estimated_investment"`
	BrandDistribution    map[string]int         `json:"brand_distribution"`
	CategoryDistribution map[string]int         `json:"category_distribution"`
	PriorityBreakdown    map[string]int         `json:"priority_breakdown"`
	RegionalAdaptations  []string               `json:"regional_adaptations"`
	OptimizationMetrics  map[string]interface{} `json:"optimization_metrics"`
}

// BusinessContext provides context for template generation
type BusinessContext struct {
	BusinessTypeID     uuid.UUID              `json:"business_type_id"`
	Location           Location               `json:"location"`
	TargetCustomers    []string               `json:"target_customers"`
	CompetitorAnalysis map[string]interface{} `json:"competitor_analysis"`
	SeasonalFactors    []string               `json:"seasonal_factors"`
	BudgetConstraints  BudgetConstraints      `json:"budget_constraints"`
}

// Location represents business location details
type Location struct {
	Country      string  `json:"country"`
	Region       string  `json:"region"`
	City         string  `json:"city"`
	Neighborhood string  `json:"neighborhood"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

// BudgetConstraints defines budget limitations
type BudgetConstraints struct {
	MinBudget       float64 `json:"min_budget"`
	MaxBudget       float64 `json:"max_budget"`
	PreferredBudget float64 `json:"preferred_budget"`
	Currency        string  `json:"currency"`
}

// ContextAnalysis provides analysis results
type ContextAnalysis struct {
	RecommendedSize      string                 `json:"recommended_size"`
	RecommendedMix       map[string]float64     `json:"recommended_mix"`
	KeyCategories        []string               `json:"key_categories"`
	LocalPreferences     []string               `json:"local_preferences"`
	SeasonalSuggestions  []string               `json:"seasonal_suggestions"`
	CompetitiveAdvantage []string               `json:"competitive_advantage"`
	RiskFactors          []string               `json:"risk_factors"`
	Opportunities        []string               `json:"opportunities"`
	ConfidenceScore      float64                `json:"confidence_score"`
	AnalysisMetadata     map[string]interface{} `json:"analysis_metadata"`
}

// ValidationResult contains template validation results
type ValidationResult struct {
	IsValid           bool                `json:"is_valid"`
	ValidationErrors  []ValidationError   `json:"validation_errors"`
	ValidationWarnings []ValidationWarning `json:"validation_warnings"`
	QualityScore      float64             `json:"quality_score"`
	Recommendations   []string            `json:"recommendations"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// ValidationWarning represents a validation warning
type ValidationWarning struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}