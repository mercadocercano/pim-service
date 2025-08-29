package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/exception"
	"saas-mt-pim-service/src/template_ai/domain/port"
)

// AIGatewayClient implements the AIGenerationService interface
type AIGatewayClient struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
}

// NewAIGatewayClient creates a new instance of AI Gateway client
func NewAIGatewayClient(baseURL string, apiKey string) port.AIGenerationService {
	return &AIGatewayClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		apiKey: apiKey,
	}
}

// GenerateTemplate generates a template using AI
func (c *AIGatewayClient) GenerateTemplate(ctx context.Context, request port.GenerationRequest) (*entity.AITemplate, error) {
	endpoint := fmt.Sprintf("%s/api/v1/ai/generate-template", c.baseURL)
	
	// Prepare request body
	reqBody := AIGenerateTemplateRequest{
		BusinessTypeID:   request.BusinessTypeID.String(),
		TemplateName:     request.TemplateName,
		TargetSize:       request.TargetSize,
		Preferences:      request.Preferences,
		Constraints:      request.Constraints,
		OptimizationGoal: request.OptimizationGoal,
		Context: map[string]interface{}{
			"region":        request.Region,
			"season":        request.Season,
			"budget_range":  request.BudgetRange,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, exception.NewInternalError("failed to marshal request: " + err.Error())
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, exception.NewInternalError("failed to create request: " + err.Error())
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, exception.NewInternalError("failed to execute request: " + err.Error())
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		var errorResp AIErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return nil, exception.NewInternalError(fmt.Sprintf("AI service returned status %d", resp.StatusCode))
		}
		return nil, exception.NewInternalError("AI service error: " + errorResp.Error)
	}

	// Parse response
	var aiResp AIGenerateTemplateResponse
	if err := json.NewDecoder(resp.Body).Decode(&aiResp); err != nil {
		return nil, exception.NewInternalError("failed to decode response: " + err.Error())
	}

	// Convert to domain entity
	template := c.convertToTemplate(aiResp)
	return template, nil
}

// ScoreProducts scores products based on AI analysis
func (c *AIGatewayClient) ScoreProducts(ctx context.Context, products []*entity.GlobalProduct, params interface{}) ([]*entity.ScoredProduct, error) {
	endpoint := fmt.Sprintf("%s/api/v1/ai/score-products", c.baseURL)

	// Convert products to request format
	productRequests := make([]AIProductRequest, 0, len(products))
	for _, product := range products {
		productRequests = append(productRequests, AIProductRequest{
			ID:           product.ID.String(),
			Name:         product.Name,
			CategoryID:   product.CategoryID,
			CategoryName: product.CategoryName,
			BrandID:      product.BrandID,
			BrandName:    product.BrandName,
			Price:        product.Price,
			Attributes:   product.Attributes,
			Tags:         product.Tags,
		})
	}

	reqBody := AIScoreProductsRequest{
		Products:   productRequests,
		Parameters: params,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, exception.NewInternalError("failed to marshal request: " + err.Error())
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, exception.NewInternalError("failed to create request: " + err.Error())
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, exception.NewInternalError("failed to execute request: " + err.Error())
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		var errorResp AIErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return nil, exception.NewInternalError(fmt.Sprintf("AI service returned status %d", resp.StatusCode))
		}
		return nil, exception.NewInternalError("AI service error: " + errorResp.Error)
	}

	// Parse response
	var aiResp AIScoreProductsResponse
	if err := json.NewDecoder(resp.Body).Decode(&aiResp); err != nil {
		return nil, exception.NewInternalError("failed to decode response: " + err.Error())
	}

	// Convert to scored products
	scoredProducts := c.convertToScoredProducts(aiResp, products)
	return scoredProducts, nil
}

// OptimizeTemplate optimizes an existing template
func (c *AIGatewayClient) OptimizeTemplate(ctx context.Context, template *entity.AITemplate, feedback []port.TemplateFeedback) (*entity.AITemplate, error) {
	endpoint := fmt.Sprintf("%s/api/v1/ai/optimize-template", c.baseURL)

	reqBody := AIOptimizeTemplateRequest{
		TemplateID:   template.ID.String(),
		TemplateName: template.Name,
		CurrentProducts: c.convertProductsToRequest(template.Products),
		Feedback:        c.convertFeedbackToRequest(feedback),
		Metrics: map[string]interface{}{
			"usage_count":        template.UsageCount,
			"satisfaction_score": template.SatisfactionScore,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, exception.NewInternalError("failed to marshal request: " + err.Error())
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, exception.NewInternalError("failed to create request: " + err.Error())
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, exception.NewInternalError("failed to execute request: " + err.Error())
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		var errorResp AIErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return nil, exception.NewInternalError(fmt.Sprintf("AI service returned status %d", resp.StatusCode))
		}
		return nil, exception.NewInternalError("AI service error: " + errorResp.Error)
	}

	// Parse response
	var aiResp AIOptimizeTemplateResponse
	if err := json.NewDecoder(resp.Body).Decode(&aiResp); err != nil {
		return nil, exception.NewInternalError("failed to decode response: " + err.Error())
	}

	// Create optimized template
	optimizedTemplate := c.convertToOptimizedTemplate(template, aiResp)
	return optimizedTemplate, nil
}

// AnalyzePerformance analyzes template performance using AI
func (c *AIGatewayClient) AnalyzePerformance(ctx context.Context, metrics []port.PerformanceMetric) (*port.PerformanceAnalysis, error) {
	endpoint := fmt.Sprintf("%s/api/v1/ai/analyze-performance", c.baseURL)

	reqBody := AIAnalyzePerformanceRequest{
		Metrics: c.convertMetricsToRequest(metrics),
		Period: map[string]string{
			"start": metrics[0].PeriodStart.Format(time.RFC3339),
			"end":   metrics[0].PeriodEnd.Format(time.RFC3339),
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, exception.NewInternalError("failed to marshal request: " + err.Error())
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, exception.NewInternalError("failed to create request: " + err.Error())
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, exception.NewInternalError("failed to execute request: " + err.Error())
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		var errorResp AIErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return nil, exception.NewInternalError(fmt.Sprintf("AI service returned status %d", resp.StatusCode))
		}
		return nil, exception.NewInternalError("AI service error: " + errorResp.Error)
	}

	// Parse response
	var aiResp AIAnalyzePerformanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&aiResp); err != nil {
		return nil, exception.NewInternalError("failed to decode response: " + err.Error())
	}

	// Convert to performance analysis
	analysis := c.convertToPerformanceAnalysis(aiResp)
	return analysis, nil
}

// Helper methods for conversion

func (c *AIGatewayClient) convertToTemplate(resp AIGenerateTemplateResponse) *entity.AITemplate {
	templateID, err := uuid.NewV4()
	if err != nil {
		templateID = uuid.Nil // Use nil UUID if generation fails
	}
	businessTypeID, _ := uuid.FromString(resp.BusinessTypeID)

	products := make([]entity.TemplateGlobalProduct, 0, len(resp.Products))
	for _, p := range resp.Products {
		productID, _ := uuid.FromString(p.GlobalProductID)
		products = append(products, entity.TemplateGlobalProduct{
			GlobalProductID:    productID,
			Priority:           p.Priority,
			QuantitySuggestion: p.QuantitySuggestion,
			AIReasoning:        p.AIReasoning,
			RelevanceScore:     p.RelevanceScore,
		})
	}

	return &entity.AITemplate{
		ID:               templateID,
		BusinessTypeID:   businessTypeID,
		Name:             resp.TemplateName,
		GeneratedBy:      "ai",
		Products:         products,
		GenerationParams: resp.GenerationParams,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

func (c *AIGatewayClient) convertToScoredProducts(resp AIScoreProductsResponse, originalProducts []*entity.GlobalProduct) []*entity.ScoredProduct {
	// Create a map for quick lookup
	productMap := make(map[string]*entity.GlobalProduct)
	for _, p := range originalProducts {
		productMap[p.ID.String()] = p
	}

	scoredProducts := make([]*entity.ScoredProduct, 0, len(resp.ScoredProducts))
	for _, sp := range resp.ScoredProducts {
		if product, exists := productMap[sp.ProductID]; exists {
			scoredProducts = append(scoredProducts, &entity.ScoredProduct{
				Product:   product,
				Score:     sp.Score,
				Reasoning: sp.Reasoning,
				Tags:      sp.SuggestedTags,
			})
		}
	}

	return scoredProducts
}

func (c *AIGatewayClient) convertProductsToRequest(products []entity.TemplateGlobalProduct) []AIProductRequest {
	requests := make([]AIProductRequest, 0, len(products))
	for _, p := range products {
		requests = append(requests, AIProductRequest{
			ID:       p.GlobalProductID.String(),
			Priority: p.Priority,
			Quantity: p.QuantitySuggestion,
		})
	}
	return requests
}

func (c *AIGatewayClient) convertFeedbackToRequest(feedback []port.TemplateFeedback) []AIFeedbackRequest {
	requests := make([]AIFeedbackRequest, 0, len(feedback))
	for _, f := range feedback {
		requests = append(requests, AIFeedbackRequest{
			ProductID:    f.ProductID.String(),
			Action:       f.Action,
			Reason:       f.Reason,
			NewQuantity:  f.NewQuantity,
			Timestamp:    f.Timestamp,
		})
	}
	return requests
}

func (c *AIGatewayClient) convertMetricsToRequest(metrics []port.PerformanceMetric) []AIMetricRequest {
	requests := make([]AIMetricRequest, 0, len(metrics))
	for _, m := range metrics {
		requests = append(requests, AIMetricRequest{
			MetricType:  m.MetricType,
			MetricValue: m.MetricValue,
			Metadata:    m.Metadata,
		})
	}
	return requests
}

func (c *AIGatewayClient) convertToOptimizedTemplate(original *entity.AITemplate, resp AIOptimizeTemplateResponse) *entity.AITemplate {
	optimized := *original
	optimized.Version++
	optimized.UpdatedAt = time.Now()
	optimized.LastAIUpdate = &time.Time{}
	*optimized.LastAIUpdate = time.Now()

	// Update products based on AI recommendations
	newProducts := make([]entity.TemplateGlobalProduct, 0, len(resp.OptimizedProducts))
	for _, p := range resp.OptimizedProducts {
		productID, _ := uuid.FromString(p.GlobalProductID)
		newProducts = append(newProducts, entity.TemplateGlobalProduct{
			GlobalProductID:    productID,
			Priority:           p.Priority,
			QuantitySuggestion: p.QuantitySuggestion,
			AIReasoning:        p.AIReasoning,
			RelevanceScore:     p.RelevanceScore,
		})
	}
	optimized.Products = newProducts

	// Update generation params with optimization data
	if optimized.GenerationParams == nil {
		optimized.GenerationParams = make(map[string]interface{})
	}
	optimized.GenerationParams["optimization_reason"] = resp.OptimizationReason
	optimized.GenerationParams["changes_summary"] = resp.ChangesSummary

	return &optimized
}

func (c *AIGatewayClient) convertToPerformanceAnalysis(resp AIAnalyzePerformanceResponse) *port.PerformanceAnalysis {
	return &port.PerformanceAnalysis{
		Score:           resp.OverallScore,
		Insights:        resp.Insights,
		Recommendations: resp.Recommendations,
		Trends:          resp.Trends,
		Metadata:        resp.Metadata,
	}
}

// Request/Response DTOs for AI Gateway

type AIGenerateTemplateRequest struct {
	BusinessTypeID   string                 `json:"business_type_id"`
	TemplateName     string                 `json:"template_name"`
	TargetSize       string                 `json:"target_size"`
	Preferences      map[string]interface{} `json:"preferences"`
	Constraints      map[string]interface{} `json:"constraints"`
	OptimizationGoal string                 `json:"optimization_goal"`
	Context          map[string]interface{} `json:"context"`
}

type AIGenerateTemplateResponse struct {
	TemplateName     string                 `json:"template_name"`
	BusinessTypeID   string                 `json:"business_type_id"`
	Products         []AIProductResponse    `json:"products"`
	GenerationParams map[string]interface{} `json:"generation_params"`
	Metadata         map[string]interface{} `json:"metadata"`
}

type AIProductResponse struct {
	GlobalProductID    string  `json:"global_product_id"`
	Priority           int     `json:"priority"`
	QuantitySuggestion int     `json:"quantity_suggestion"`
	AIReasoning        string  `json:"ai_reasoning"`
	RelevanceScore     float64 `json:"relevance_score"`
}

type AIProductRequest struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	CategoryID   string                 `json:"category_id"`
	CategoryName string                 `json:"category_name"`
	BrandID      string                 `json:"brand_id"`
	BrandName    string                 `json:"brand_name"`
	Price        float64                `json:"price"`
	Attributes   map[string]interface{} `json:"attributes"`
	Tags         []string               `json:"tags"`
	Priority     int                    `json:"priority,omitempty"`
	Quantity     int                    `json:"quantity,omitempty"`
}

type AIScoreProductsRequest struct {
	Products   []AIProductRequest `json:"products"`
	Parameters interface{}        `json:"parameters"`
}

type AIScoreProductsResponse struct {
	ScoredProducts []AIScoredProduct `json:"scored_products"`
}

type AIScoredProduct struct {
	ProductID     string   `json:"product_id"`
	Score         float64  `json:"score"`
	Reasoning     string   `json:"reasoning"`
	SuggestedTags []string `json:"suggested_tags"`
}

type AIOptimizeTemplateRequest struct {
	TemplateID      string                 `json:"template_id"`
	TemplateName    string                 `json:"template_name"`
	CurrentProducts []AIProductRequest     `json:"current_products"`
	Feedback        []AIFeedbackRequest    `json:"feedback"`
	Metrics         map[string]interface{} `json:"metrics"`
}

type AIOptimizeTemplateResponse struct {
	OptimizedProducts  []AIProductResponse `json:"optimized_products"`
	OptimizationReason string              `json:"optimization_reason"`
	ChangesSummary     map[string]int      `json:"changes_summary"`
}

type AIFeedbackRequest struct {
	ProductID   string    `json:"product_id"`
	Action      string    `json:"action"`
	Reason      string    `json:"reason"`
	NewQuantity int       `json:"new_quantity,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}

type AIAnalyzePerformanceRequest struct {
	Metrics []AIMetricRequest     `json:"metrics"`
	Period  map[string]string     `json:"period"`
}

type AIAnalyzePerformanceResponse struct {
	OverallScore    float64                `json:"overall_score"`
	Insights        []string               `json:"insights"`
	Recommendations []string               `json:"recommendations"`
	Trends          map[string]interface{} `json:"trends"`
	Metadata        map[string]interface{} `json:"metadata"`
}

type AIMetricRequest struct {
	MetricType  string                 `json:"metric_type"`
	MetricValue float64                `json:"metric_value"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type AIErrorResponse struct {
	Error   string `json:"error"`
	Code    string `json:"code"`
	Details string `json:"details,omitempty"`
}