package service

import (
	"context"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/exception"
	"saas-mt-pim-service/src/template_ai/domain/port"
	"saas-mt-pim-service/src/template_ai/domain/service"
)

// AITemplateEngineImpl implements the TemplateEngine interface
type AITemplateEngineImpl struct {
	aiGenerator        port.AIGenerationService
	globalProductRepo  port.GlobalProductRepository
	templateRepo       port.AITemplateRepository
	businessTypeRepo   port.BusinessTypeRepository
}

// NewAITemplateEngine creates a new instance of AI Template Engine
func NewAITemplateEngine(
	aiGenerator port.AIGenerationService,
	globalProductRepo port.GlobalProductRepository,
	templateRepo port.AITemplateRepository,
	businessTypeRepo port.BusinessTypeRepository,
) service.TemplateEngine {
	return &AITemplateEngineImpl{
		aiGenerator:       aiGenerator,
		globalProductRepo: globalProductRepo,
		templateRepo:      templateRepo,
		businessTypeRepo:  businessTypeRepo,
	}
}

// GenerateFromCatalog generates a smart template based on global catalog and parameters
func (e *AITemplateEngineImpl) GenerateFromCatalog(ctx context.Context, params service.TemplateParams) (*service.TemplateResult, error) {
	// Validate parameters
	if err := e.validateParams(params); err != nil {
		return nil, err
	}

	// Get business type information
	businessType, err := e.businessTypeRepo.GetByID(ctx, params.BusinessTypeID)
	if err != nil {
		return nil, exception.NewNotFoundError("business type not found")
	}

	// Analyze context and determine optimal product count
	targetCount := e.calculateTargetProductCount(params)

	// Get global products based on filters
	globalProducts, err := e.getFilteredGlobalProducts(ctx, businessType, params)
	if err != nil {
		return nil, err
	}

	// Calculate product mix based on business rules
	productMix := e.calculateOptimalMix(params)

	// Select products using AI scoring
	selectedProducts, err := e.selectProductsWithAI(ctx, globalProducts, targetCount, productMix, params)
	if err != nil {
		return nil, err
	}

	// Organize products by category and brand
	categories := e.organizeByCategory(selectedProducts)
	brands := e.organizeByBrand(selectedProducts)

	// Calculate summary statistics
	summary := e.calculateSummary(selectedProducts, categories, brands)

	// Generate template ID
	templateID, err := uuid.NewV4()
	if err != nil {
		return nil, exception.NewInternalError("failed to generate template ID: " + err.Error())
	}

	// Create result
	result := &service.TemplateResult{
		TemplateID:      templateID,
		Name:            e.generateTemplateName(businessType.Name, params),
		Version:         1,
		GeneratedBy:     "ai",
		BusinessTypeID:  params.BusinessTypeID,
		Products:        selectedProducts,
		Categories:      categories,
		Brands:          brands,
		Summary:         summary,
		OptimizationScore: e.calculateOptimizationScore(selectedProducts, params),
		GenerationMetadata: map[string]interface{}{
			"generated_at":    time.Now(),
			"total_evaluated": len(globalProducts),
			"ai_model":        "gpt-4",
			"parameters":      params,
		},
	}

	return result, nil
}

// OptimizeForRegion adapts a template for specific regional preferences
func (e *AITemplateEngineImpl) OptimizeForRegion(ctx context.Context, template *service.TemplateResult, region string) (*service.TemplateResult, error) {
	// Get regional preferences
	regionalPrefs, err := e.getRegionalPreferences(ctx, region)
	if err != nil {
		return nil, err
	}

	// Adjust product selection based on regional preferences
	optimizedProducts := make([]service.ProductRecommendation, 0, len(template.Products))
	
	for _, product := range template.Products {
		// Check if product is suitable for region
		if e.isProductSuitableForRegion(product, regionalPrefs) {
			// Adjust quantity suggestions based on regional consumption patterns
			product.QuantitySuggestion = e.adjustQuantityForRegion(product.QuantitySuggestion, product.CategoryID, regionalPrefs)
			optimizedProducts = append(optimizedProducts, product)
		}
	}

	// Add region-specific products if needed
	regionSpecificProducts, err := e.getRegionSpecificProducts(ctx, template.BusinessTypeID, region)
	if err == nil && len(regionSpecificProducts) > 0 {
		optimizedProducts = append(optimizedProducts, regionSpecificProducts...)
	}

	// Recalculate distributions
	categories := e.organizeByCategory(optimizedProducts)
	brands := e.organizeByBrand(optimizedProducts)
	summary := e.calculateSummary(optimizedProducts, categories, brands)
	summary.RegionalAdaptations = regionalPrefs.Adaptations

	// Create optimized template
	optimizedTemplate := *template
	optimizedTemplate.Products = optimizedProducts
	optimizedTemplate.Categories = categories
	optimizedTemplate.Brands = brands
	optimizedTemplate.Summary = summary
	optimizedTemplate.Version++
	optimizedTemplate.GenerationMetadata["optimized_for_region"] = region
	optimizedTemplate.GenerationMetadata["optimization_date"] = time.Now()

	return &optimizedTemplate, nil
}

// CalculateProductMix determines the optimal product distribution for a business type
func (e *AITemplateEngineImpl) CalculateProductMix(ctx context.Context, businessType string, targetCount int) ([]service.ProductSelection, error) {
	// Get business type configuration
	btConfig, err := e.getBusinessTypeConfig(ctx, businessType)
	if err != nil {
		return nil, err
	}

	// Calculate distribution based on business rules
	essentialCount := int(math.Ceil(float64(targetCount) * btConfig.EssentialRatio))
	recommendedCount := int(math.Ceil(float64(targetCount) * btConfig.RecommendedRatio))
	optionalCount := targetCount - essentialCount - recommendedCount

	// Get products by priority
	essentialProducts, err := e.getProductsByPriority(ctx, businessType, 1, essentialCount)
	if err != nil {
		return nil, err
	}

	recommendedProducts, err := e.getProductsByPriority(ctx, businessType, 2, recommendedCount)
	if err != nil {
		return nil, err
	}

	optionalProducts, err := e.getProductsByPriority(ctx, businessType, 3, optionalCount)
	if err != nil {
		return nil, err
	}

	// Combine all selections
	selections := make([]service.ProductSelection, 0, targetCount)
	selections = append(selections, essentialProducts...)
	selections = append(selections, recommendedProducts...)
	selections = append(selections, optionalProducts...)

	return selections, nil
}

// AnalyzeBusinessContext analyzes the business context to provide template recommendations
func (e *AITemplateEngineImpl) AnalyzeBusinessContext(ctx context.Context, businessContext service.BusinessContext) (*service.ContextAnalysis, error) {
	// Analyze location-based factors
	locationAnalysis := e.analyzeLocation(businessContext.Location)

	// Analyze customer demographics
	customerAnalysis := e.analyzeTargetCustomers(businessContext.TargetCustomers)

	// Analyze competition
	competitionAnalysis := e.analyzeCompetition(businessContext.CompetitorAnalysis)

	// Analyze seasonal factors
	seasonalAnalysis := e.analyzeSeasonalFactors(businessContext.SeasonalFactors)

	// Combine analyses to generate recommendations
	analysis := &service.ContextAnalysis{
		RecommendedSize:      e.determineRecommendedSize(businessContext.BudgetConstraints, locationAnalysis),
		RecommendedMix:       e.calculateRecommendedMix(customerAnalysis, competitionAnalysis),
		KeyCategories:        e.identifyKeyCategories(businessContext.BusinessTypeID, customerAnalysis),
		LocalPreferences:     locationAnalysis.Preferences,
		SeasonalSuggestions:  seasonalAnalysis.Suggestions,
		CompetitiveAdvantage: competitionAnalysis.Advantages,
		RiskFactors:          e.identifyRisks(businessContext),
		Opportunities:        e.identifyOpportunities(businessContext, locationAnalysis),
		ConfidenceScore:      e.calculateConfidenceScore(businessContext),
		AnalysisMetadata: map[string]interface{}{
			"analyzed_at":     time.Now(),
			"location_score":  locationAnalysis.Score,
			"competition_level": competitionAnalysis.Level,
		},
	}

	return analysis, nil
}

// ValidateTemplate validates a template against business rules and constraints
func (e *AITemplateEngineImpl) ValidateTemplate(ctx context.Context, template *service.TemplateResult) (*service.ValidationResult, error) {
	errors := []service.ValidationError{}
	warnings := []service.ValidationWarning{}
	recommendations := []string{}

	// Validate product count
	if template.Summary.TotalProducts < 10 {
		errors = append(errors, service.ValidationError{
			Field:   "products",
			Message: "Template must have at least 10 products",
			Code:    "MIN_PRODUCTS",
		})
	}

	// Validate category distribution
	if template.Summary.TotalCategories < 3 {
		warnings = append(warnings, service.ValidationWarning{
			Field:   "categories",
			Message: "Template should have at least 3 categories for variety",
			Code:    "LOW_CATEGORY_VARIETY",
		})
	}

	// Validate brand distribution
	genericPercent := e.calculateGenericPercent(template.Brands)
	if genericPercent > 30 {
		warnings = append(warnings, service.ValidationWarning{
			Field:   "brands",
			Message: fmt.Sprintf("Generic products represent %.1f%% of template, consider reducing", genericPercent),
			Code:    "HIGH_GENERIC_PERCENT",
		})
	}

	// Validate price distribution
	priceAnalysis := e.analyzePriceDistribution(template.Products)
	if priceAnalysis.Variance > 0.5 {
		recommendations = append(recommendations, "Consider balancing price ranges for better customer appeal")
	}

	// Calculate quality score
	qualityScore := e.calculateQualityScore(template, len(errors), len(warnings))

	result := &service.ValidationResult{
		IsValid:            len(errors) == 0,
		ValidationErrors:   errors,
		ValidationWarnings: warnings,
		QualityScore:       qualityScore,
		Recommendations:    recommendations,
	}

	return result, nil
}

// Helper methods

func (e *AITemplateEngineImpl) validateParams(params service.TemplateParams) error {
	if params.MinProducts > params.MaxProducts {
		return exception.NewValidationError("min_products", "cannot be greater than max_products")
	}
	
	if params.GenericPercent > 100 || params.GenericPercent < 0 {
		return exception.NewValidationError("generic_percent", "must be between 0 and 100")
	}

	return nil
}

func (e *AITemplateEngineImpl) calculateTargetProductCount(params service.TemplateParams) int {
	sizeMultipliers := map[string]float64{
		"small":  0.5,
		"medium": 1.0,
		"large":  1.5,
	}

	multiplier, exists := sizeMultipliers[params.TargetSize]
	if !exists {
		multiplier = 1.0
	}

	baseCount := (params.MinProducts + params.MaxProducts) / 2
	return int(float64(baseCount) * multiplier)
}

func (e *AITemplateEngineImpl) getFilteredGlobalProducts(ctx context.Context, businessType *entity.BusinessType, params service.TemplateParams) ([]*entity.GlobalProduct, error) {
	// Build filters based on parameters
	filters := map[string]interface{}{
		"business_type": businessType.Code,
		"is_active":     true,
	}

	if len(params.CategoryFocus) > 0 {
		filters["categories"] = params.CategoryFocus
	}

	if len(params.ExcludeBrands) > 0 {
		filters["exclude_brands"] = params.ExcludeBrands
	}

	// Get products from global catalog
	return e.globalProductRepo.GetFilteredProducts(ctx, filters)
}

func (e *AITemplateEngineImpl) calculateOptimalMix(params service.TemplateParams) map[string]float64 {
	// Default mix
	mix := map[string]float64{
		"essential":    0.6,
		"recommended":  0.3,
		"optional":     0.1,
	}

	// Adjust based on optimization goal
	switch params.OptimizationGoal {
	case "variety":
		mix["essential"] = 0.4
		mix["recommended"] = 0.4
		mix["optional"] = 0.2
	case "profit_margin":
		mix["essential"] = 0.5
		mix["recommended"] = 0.4
		mix["optional"] = 0.1
	case "fast_moving":
		mix["essential"] = 0.7
		mix["recommended"] = 0.2
		mix["optional"] = 0.1
	}

	return mix
}

func (e *AITemplateEngineImpl) selectProductsWithAI(ctx context.Context, products []*entity.GlobalProduct, targetCount int, mix map[string]float64, params service.TemplateParams) ([]service.ProductRecommendation, error) {
	// Use AI service to score products
	scoredProducts, err := e.aiGenerator.ScoreProducts(ctx, products, params)
	if err != nil {
		return nil, err
	}

	// Sort by score
	sort.Slice(scoredProducts, func(i, j int) bool {
		return scoredProducts[i].Score > scoredProducts[j].Score
	})

	// Select products based on mix
	recommendations := make([]service.ProductRecommendation, 0, targetCount)
	
	essentialCount := int(float64(targetCount) * mix["essential"])
	recommendedCount := int(float64(targetCount) * mix["recommended"])

	// Add essential products
	for i := 0; i < essentialCount && i < len(scoredProducts); i++ {
		rec := e.createRecommendation(scoredProducts[i], 1)
		recommendations = append(recommendations, rec)
	}

	// Add recommended products
	for i := essentialCount; i < essentialCount+recommendedCount && i < len(scoredProducts); i++ {
		rec := e.createRecommendation(scoredProducts[i], 2)
		recommendations = append(recommendations, rec)
	}

	// Add optional products
	for i := essentialCount + recommendedCount; i < targetCount && i < len(scoredProducts); i++ {
		rec := e.createRecommendation(scoredProducts[i], 3)
		recommendations = append(recommendations, rec)
	}

	return recommendations, nil
}

func (e *AITemplateEngineImpl) createRecommendation(scoredProduct *entity.ScoredProduct, priority int) service.ProductRecommendation {
	return service.ProductRecommendation{
		GlobalProductID:    scoredProduct.Product.ID,
		Name:               scoredProduct.Product.Name,
		CategoryID:         scoredProduct.Product.CategoryID,
		CategoryName:       scoredProduct.Product.CategoryName,
		BrandID:            scoredProduct.Product.BrandID,
		BrandName:          scoredProduct.Product.BrandName,
		Priority:           priority,
		QuantitySuggestion: e.calculateQuantitySuggestion(scoredProduct.Product, priority),
		Price:              scoredProduct.Product.Price,
		RelevanceScore:     scoredProduct.Score,
		AIReasoning:        scoredProduct.Reasoning,
		Tags:               scoredProduct.Product.Tags,
		Attributes:         scoredProduct.Product.Attributes,
	}
}

func (e *AITemplateEngineImpl) calculateQuantitySuggestion(product *entity.GlobalProduct, priority int) int {
	// Base quantity on priority and product type
	baseQuantity := map[int]int{
		1: 12, // Essential
		2: 6,  // Recommended
		3: 3,  // Optional
	}

	quantity := baseQuantity[priority]

	// Adjust based on product attributes
	if product.IsBulk {
		quantity = quantity / 2
	}

	if product.IsPerishable {
		quantity = quantity / 3
	}

	return quantity
}

func (e *AITemplateEngineImpl) organizeByCategory(products []service.ProductRecommendation) []service.CategoryDistribution {
	categoryMap := make(map[string]*service.CategoryDistribution)

	for _, product := range products {
		if dist, exists := categoryMap[product.CategoryID]; exists {
			dist.ProductCount++
		} else {
			categoryMap[product.CategoryID] = &service.CategoryDistribution{
				CategoryID:   product.CategoryID,
				CategoryName: product.CategoryName,
				ProductCount: 1,
				Priority:     e.getCategoryPriority(product.CategoryID),
			}
		}
	}

	// Convert to slice and calculate percentages
	distributions := make([]service.CategoryDistribution, 0, len(categoryMap))
	total := float64(len(products))

	for _, dist := range categoryMap {
		dist.Percentage = (float64(dist.ProductCount) / total) * 100
		distributions = append(distributions, *dist)
	}

	// Sort by product count
	sort.Slice(distributions, func(i, j int) bool {
		return distributions[i].ProductCount > distributions[j].ProductCount
	})

	return distributions
}

func (e *AITemplateEngineImpl) organizeByBrand(products []service.ProductRecommendation) []service.BrandDistribution {
	brandMap := make(map[string]*service.BrandDistribution)

	for _, product := range products {
		if dist, exists := brandMap[product.BrandID]; exists {
			dist.ProductCount++
		} else {
			brandMap[product.BrandID] = &service.BrandDistribution{
				BrandID:      product.BrandID,
				BrandName:    product.BrandName,
				ProductCount: 1,
				BrandType:    e.getBrandType(product.BrandName),
			}
		}
	}

	// Convert to slice and calculate percentages
	distributions := make([]service.BrandDistribution, 0, len(brandMap))
	total := float64(len(products))

	for _, dist := range brandMap {
		dist.Percentage = (float64(dist.ProductCount) / total) * 100
		distributions = append(distributions, *dist)
	}

	// Sort by product count
	sort.Slice(distributions, func(i, j int) bool {
		return distributions[i].ProductCount > distributions[j].ProductCount
	})

	return distributions
}

func (e *AITemplateEngineImpl) calculateSummary(
	products []service.ProductRecommendation,
	categories []service.CategoryDistribution,
	brands []service.BrandDistribution,
) service.TemplateSummary {
	// Calculate total investment
	totalInvestment := 0.0
	priorityBreakdown := make(map[string]int)
	
	for _, product := range products {
		totalInvestment += product.Price * float64(product.QuantitySuggestion)
		
		priorityKey := fmt.Sprintf("priority_%d", product.Priority)
		priorityBreakdown[priorityKey]++
	}

	// Calculate brand distribution by type
	brandDistribution := make(map[string]int)
	for _, brand := range brands {
		brandDistribution[brand.BrandType] += brand.ProductCount
	}

	// Calculate category distribution
	categoryDistribution := make(map[string]int)
	for _, category := range categories {
		categoryDistribution[category.CategoryName] = category.ProductCount
	}

	return service.TemplateSummary{
		TotalProducts:        len(products),
		TotalCategories:      len(categories),
		TotalBrands:          len(brands),
		EstimatedInvestment:  totalInvestment,
		BrandDistribution:    brandDistribution,
		CategoryDistribution: categoryDistribution,
		PriorityBreakdown:    priorityBreakdown,
		RegionalAdaptations:  []string{},
		OptimizationMetrics: map[string]interface{}{
			"average_products_per_category": float64(len(products)) / float64(len(categories)),
			"average_products_per_brand":    float64(len(products)) / float64(len(brands)),
			"diversity_score":               e.calculateDiversityScore(categories, brands),
		},
	}
}

func (e *AITemplateEngineImpl) calculateOptimizationScore(products []service.ProductRecommendation, params service.TemplateParams) float64 {
	score := 0.0
	
	// Score based on meeting target count
	targetCount := e.calculateTargetProductCount(params)
	countScore := 1.0 - math.Abs(float64(len(products)-targetCount))/float64(targetCount)
	score += countScore * 0.3

	// Score based on relevance scores
	avgRelevance := 0.0
	for _, product := range products {
		avgRelevance += product.RelevanceScore
	}
	avgRelevance /= float64(len(products))
	score += avgRelevance * 0.4

	// Score based on diversity
	categoryCount := make(map[string]int)
	for _, product := range products {
		categoryCount[product.CategoryID]++
	}
	diversityScore := float64(len(categoryCount)) / float64(len(products))
	score += diversityScore * 0.3

	return math.Min(score, 1.0)
}

func (e *AITemplateEngineImpl) generateTemplateName(businessTypeName string, params service.TemplateParams) string {
	sizeMap := map[string]string{
		"small":  "Básico",
		"medium": "Estándar",
		"large":  "Premium",
	}

	size := sizeMap[params.TargetSize]
	if size == "" {
		size = "Estándar"
	}

	return fmt.Sprintf("%s %s", businessTypeName, size)
}

// Additional helper methods for regional optimization, business context analysis, etc.
// These would be implemented based on specific business rules and AI integration

func (e *AITemplateEngineImpl) getRegionalPreferences(ctx context.Context, region string) (*RegionalPreferences, error) {
	// Mock implementation - would integrate with regional data service
	return &RegionalPreferences{
		Region: region,
		Preferences: map[string]float64{
			"local_products": 0.3,
			"imported":       0.7,
		},
		Adaptations: []string{
			"Productos locales preferidos",
			"Marcas regionales incluidas",
		},
	}, nil
}

func (e *AITemplateEngineImpl) isProductSuitableForRegion(product service.ProductRecommendation, prefs *RegionalPreferences) bool {
	// Mock implementation - would check product attributes against regional preferences
	return true
}

func (e *AITemplateEngineImpl) adjustQuantityForRegion(baseQuantity int, categoryID string, prefs *RegionalPreferences) int {
	// Mock implementation - would adjust based on regional consumption patterns
	return baseQuantity
}

func (e *AITemplateEngineImpl) getRegionSpecificProducts(ctx context.Context, businessTypeID uuid.UUID, region string) ([]service.ProductRecommendation, error) {
	// Mock implementation - would fetch region-specific products
	return []service.ProductRecommendation{}, nil
}

func (e *AITemplateEngineImpl) getBusinessTypeConfig(ctx context.Context, businessType string) (*BusinessTypeConfig, error) {
	// Mock implementation - would fetch from configuration service
	return &BusinessTypeConfig{
		EssentialRatio:   0.6,
		RecommendedRatio: 0.3,
		OptionalRatio:    0.1,
	}, nil
}

func (e *AITemplateEngineImpl) getProductsByPriority(ctx context.Context, businessType string, priority int, count int) ([]service.ProductSelection, error) {
	// Mock implementation - would fetch from repository
	selections := make([]service.ProductSelection, 0, count)
	
	for i := 0; i < count && i < 10; i++ {
		productID, err := uuid.NewV4()
		if err != nil {
			continue // Skip this iteration if UUID generation fails
		}
		selections = append(selections, service.ProductSelection{
			ProductID:      productID,
			CategoryID:     fmt.Sprintf("category_%d", i%5),
			Priority:       priority,
			SelectionScore: 0.8,
			Reason:         "Selected based on popularity and relevance",
		})
	}
	
	return selections, nil
}

func (e *AITemplateEngineImpl) getCategoryPriority(categoryID string) int {
	// Mock implementation - would be based on business rules
	priorityMap := map[string]int{
		"bebidas":   1,
		"alimentos": 1,
		"limpieza":  2,
		"otros":     3,
	}
	
	if priority, exists := priorityMap[categoryID]; exists {
		return priority
	}
	return 2
}

func (e *AITemplateEngineImpl) getBrandType(brandName string) string {
	// Mock implementation - would classify brands
	if brandName == "Genérico" || brandName == "S/Marca" {
		return "generic"
	}
	// Would have more sophisticated classification
	return "standard"
}

func (e *AITemplateEngineImpl) calculateDiversityScore(categories []service.CategoryDistribution, brands []service.BrandDistribution) float64 {
	// Shannon diversity index
	categoryShannonIndex := 0.0
	for _, cat := range categories {
		p := cat.Percentage / 100.0
		if p > 0 {
			categoryShannonIndex -= p * math.Log(p)
		}
	}
	
	brandShannonIndex := 0.0
	for _, brand := range brands {
		p := brand.Percentage / 100.0
		if p > 0 {
			brandShannonIndex -= p * math.Log(p)
		}
	}
	
	// Normalize and combine
	maxCategoryIndex := math.Log(float64(len(categories)))
	maxBrandIndex := math.Log(float64(len(brands)))
	
	categoryDiversity := categoryShannonIndex / maxCategoryIndex
	brandDiversity := brandShannonIndex / maxBrandIndex
	
	return (categoryDiversity + brandDiversity) / 2.0
}

func (e *AITemplateEngineImpl) calculateGenericPercent(brands []service.BrandDistribution) float64 {
	genericCount := 0.0
	totalCount := 0.0
	
	for _, brand := range brands {
		totalCount += float64(brand.ProductCount)
		if brand.BrandType == "generic" {
			genericCount += float64(brand.ProductCount)
		}
	}
	
	if totalCount == 0 {
		return 0
	}
	
	return (genericCount / totalCount) * 100
}

func (e *AITemplateEngineImpl) analyzePriceDistribution(products []service.ProductRecommendation) *PriceAnalysis {
	prices := make([]float64, len(products))
	totalPrice := 0.0
	
	for i, product := range products {
		prices[i] = product.Price
		totalPrice += product.Price
	}
	
	mean := totalPrice / float64(len(products))
	
	// Calculate variance
	variance := 0.0
	for _, price := range prices {
		variance += math.Pow(price-mean, 2)
	}
	variance /= float64(len(products))
	
	return &PriceAnalysis{
		Mean:     mean,
		Variance: variance / mean, // Coefficient of variation
		Min:      minFloat64(prices),
		Max:      maxFloat64(prices),
	}
}

func (e *AITemplateEngineImpl) calculateQualityScore(template *service.TemplateResult, errorCount, warningCount int) float64 {
	// Base score
	score := 1.0
	
	// Deduct for errors and warnings
	score -= float64(errorCount) * 0.2
	score -= float64(warningCount) * 0.05
	
	// Add points for good characteristics
	if template.Summary.TotalCategories >= 5 {
		score += 0.1
	}
	
	if template.OptimizationScore > 0.8 {
		score += 0.1
	}
	
	// Ensure score is between 0 and 1
	return math.Max(0, math.Min(1, score))
}

// Helper structs

type RegionalPreferences struct {
	Region      string
	Preferences map[string]float64
	Adaptations []string
}

type BusinessTypeConfig struct {
	EssentialRatio   float64
	RecommendedRatio float64
	OptionalRatio    float64
}

type LocationAnalysis struct {
	Score       float64
	Preferences []string
}

type CustomerAnalysis struct {
	Segments map[string]float64
}

type CompetitionAnalysis struct {
	Level      string
	Advantages []string
}

type SeasonalAnalysis struct {
	Suggestions []string
}

type PriceAnalysis struct {
	Mean     float64
	Variance float64
	Min      float64
	Max      float64
}

// Utility functions

func minFloat64(slice []float64) float64 {
	if len(slice) == 0 {
		return 0
	}
	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func maxFloat64(slice []float64) float64 {
	if len(slice) == 0 {
		return 0
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Mock implementations for context analysis methods

func (e *AITemplateEngineImpl) analyzeLocation(location service.Location) LocationAnalysis {
	return LocationAnalysis{
		Score:       0.85,
		Preferences: []string{"Productos locales", "Marcas regionales"},
	}
}

func (e *AITemplateEngineImpl) analyzeTargetCustomers(customers []string) CustomerAnalysis {
	return CustomerAnalysis{
		Segments: map[string]float64{
			"families":      0.4,
			"young_adults":  0.3,
			"professionals": 0.3,
		},
	}
}

func (e *AITemplateEngineImpl) analyzeCompetition(competition map[string]interface{}) CompetitionAnalysis {
	return CompetitionAnalysis{
		Level:      "medium",
		Advantages: []string{"Mejor selección de productos", "Precios competitivos"},
	}
}

func (e *AITemplateEngineImpl) analyzeSeasonalFactors(factors []string) SeasonalAnalysis {
	return SeasonalAnalysis{
		Suggestions: []string{"Productos de verano", "Bebidas refrescantes"},
	}
}

func (e *AITemplateEngineImpl) determineRecommendedSize(budget service.BudgetConstraints, location LocationAnalysis) string {
	if budget.MaxBudget < 500000 {
		return "small"
	} else if budget.MaxBudget < 1500000 {
		return "medium"
	}
	return "large"
}

func (e *AITemplateEngineImpl) calculateRecommendedMix(customer CustomerAnalysis, competition CompetitionAnalysis) map[string]float64 {
	return map[string]float64{
		"essential":   0.5,
		"recommended": 0.35,
		"optional":    0.15,
	}
}

func (e *AITemplateEngineImpl) identifyKeyCategories(businessTypeID uuid.UUID, customer CustomerAnalysis) []string {
	return []string{"bebidas", "snacks", "limpieza", "alimentos_basicos"}
}

func (e *AITemplateEngineImpl) identifyRisks(context service.BusinessContext) []string {
	return []string{"Competencia establecida", "Estacionalidad de productos"}
}

func (e *AITemplateEngineImpl) identifyOpportunities(context service.BusinessContext, location LocationAnalysis) []string {
	return []string{"Demanda de productos locales", "Nicho no atendido"}
}

func (e *AITemplateEngineImpl) calculateConfidenceScore(context service.BusinessContext) float64 {
	return 0.85
}