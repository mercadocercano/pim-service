package usecase

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/gofrs/uuid/v5"
	appPort "saas-mt-pim-service/src/template_ai/application/port"
	"saas-mt-pim-service/src/template_ai/application/request"
	"saas-mt-pim-service/src/template_ai/application/response"
	"saas-mt-pim-service/src/template_ai/domain/exception"
	"saas-mt-pim-service/src/template_ai/domain/port"
	"saas-mt-pim-service/src/template_ai/domain/value_object"
)

// AnalyzeTemplatePerformanceUseCase handles template performance analysis
type AnalyzeTemplatePerformanceUseCase struct {
	aiTemplateRepo port.AITemplateRepository
	domainService  port.AITemplateDomainServicePort
	mapper         appPort.TemplateMapperPort
}

// NewAnalyzeTemplatePerformanceUseCase creates a new instance of the use case
func NewAnalyzeTemplatePerformanceUseCase(
	aiTemplateRepo port.AITemplateRepository,
	domainService port.AITemplateDomainServicePort,
	mapper appPort.TemplateMapperPort,
) *AnalyzeTemplatePerformanceUseCase {
	return &AnalyzeTemplatePerformanceUseCase{
		aiTemplateRepo: aiTemplateRepo,
		domainService:  domainService,
		mapper:         mapper,
	}
}

// Execute analyzes template performance
func (uc *AnalyzeTemplatePerformanceUseCase) Execute(
	ctx context.Context,
	req *request.AnalyzeTemplatePerformanceRequest,
) (*response.TemplatePerformanceResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Parse template ID
	templateID, err := uuid.FromString(req.TemplateID)
	if err != nil {
		return nil, exception.NewValidationError("template_id", "invalid UUID format")
	}

	// Get template
	template, err := uc.aiTemplateRepo.FindByID(ctx, templateID)
	if err != nil {
		return nil, exception.ErrTemplateNotFound
	}

	// Get period
	startTime, endTime := req.GetPeriodOrDefault()

	// Get metrics
	metrics, err := uc.collectMetrics(ctx, templateID, req.GetMetricTypesOrDefault(), startTime, endTime)
	if err != nil {
		return nil, fmt.Errorf("failed to collect metrics: %w", err)
	}

	// Calculate performance score
	performanceScore := uc.calculatePerformanceScore(metrics)

	// Generate trends if requested
	var trends map[string][]response.TrendPointResponse
	if req.GroupBy != "" {
		trends = uc.generateTrends(ctx, templateID, req.GetMetricTypesOrDefault(), startTime, endTime, req.GroupBy)
	}

	// Generate comparisons if requested
	var comparisons []response.TemplateComparisonResponse
	if len(req.CompareWith) > 0 {
		comparisons = uc.generateComparisons(ctx, templateID, req.CompareWith, metrics)
	}

	// Generate recommendations
	recommendations := uc.generateRecommendations(metrics, template)

	// Prepare response
	return &response.TemplatePerformanceResponse{
		TemplateID:   templateID,
		TemplateName: template.Name,
		AnalysisPeriod: &response.PeriodResponse{
			Start: startTime,
			End:   endTime,
			Days:  int(endTime.Sub(startTime).Hours() / 24),
		},
		Metrics:          metrics,
		Trends:           trends,
		Comparisons:      comparisons,
		Recommendations:  recommendations,
		PerformanceScore: performanceScore,
		PerformanceRating: uc.mapper.GetPerformanceRating(performanceScore),
	}, nil
}

// collectMetrics collects all requested metrics for the template
func (uc *AnalyzeTemplatePerformanceUseCase) collectMetrics(
	ctx context.Context,
	templateID uuid.UUID,
	metricTypes []string,
	startTime, endTime time.Time,
) (map[string]response.MetricValueResponse, error) {
	result := make(map[string]response.MetricValueResponse)

	// Get aggregated metrics from domain service
	aggregatedMetrics, err := uc.domainService.CalculateTemplatePerformance(ctx, templateID)
	if err != nil {
		return nil, err
	}

	// Get stored metrics for the period
	for _, metricType := range metricTypes {
		storedMetrics, err := uc.aiTemplateRepo.FindPerformanceMetrics(ctx, templateID, metricType)
		if err != nil {
			continue
		}

		// Filter metrics by period and calculate average
		var totalValue float64
		var count int
		for _, metric := range storedMetrics {
			if metric.PeriodStart.After(startTime) && metric.PeriodEnd.Before(endTime) {
				totalValue += metric.MetricValue
				count++
			}
		}

		var value float64
		if count > 0 {
			value = totalValue / float64(count)
		} else if aggregatedValue, exists := aggregatedMetrics[metricType]; exists {
			value = aggregatedValue
		} else {
			// Default values for metrics without data
			value = uc.getDefaultMetricValue(metricType)
		}

		// Create metric response
		target := uc.getMetricTarget(metricType)
		metricResponse := uc.mapper.ToMetricValueResponse(value, target)

		// Add additional details based on metric type
		switch metricType {
		case string(value_object.MetricTypeUsageCount):
			metricResponse.Details = map[string]interface{}{
				"description": "Number of times this template has been applied",
				"unit":        "count",
			}
		case string(value_object.MetricTypeSatisfactionScore):
			metricResponse.Details = map[string]interface{}{
				"description": "User satisfaction based on feedback and modifications",
				"unit":        "percentage",
			}
		case string(value_object.MetricTypeModificationRate):
			metricResponse.Details = map[string]interface{}{
				"description": "Percentage of products modified after template application",
				"unit":        "percentage",
			}
		}

		result[metricType] = *metricResponse
	}

	return result, nil
}

// generateTrends generates trend data for metrics
func (uc *AnalyzeTemplatePerformanceUseCase) generateTrends(
	ctx context.Context,
	templateID uuid.UUID,
	metricTypes []string,
	startTime, endTime time.Time,
	groupBy string,
) map[string][]response.TrendPointResponse {
	trends := make(map[string][]response.TrendPointResponse)

	// For each metric type, generate trend points
	for _, metricType := range metricTypes {
		trendPoints := uc.generateTrendPoints(ctx, templateID, metricType, startTime, endTime, groupBy)
		if len(trendPoints) > 0 {
			trends[metricType] = trendPoints
		}
	}

	return trends
}

// generateTrendPoints generates trend points for a specific metric
func (uc *AnalyzeTemplatePerformanceUseCase) generateTrendPoints(
	ctx context.Context,
	templateID uuid.UUID,
	metricType string,
	startTime, endTime time.Time,
	groupBy string,
) []response.TrendPointResponse {
	// In a real implementation, this would query historical data
	// For now, we'll generate sample trend data
	var points []response.TrendPointResponse

	current := startTime
	for current.Before(endTime) {
		// Generate a sample value (in real implementation, would query actual data)
		value := uc.getDefaultMetricValue(metricType) + (0.1 * float64(current.Unix()%10))
		
		points = append(points, response.TrendPointResponse{
			Timestamp: current,
			Value:     value,
			Label:     uc.formatTimeLabel(current, groupBy),
		})

		// Advance time based on groupBy
		switch groupBy {
		case "day":
			current = current.AddDate(0, 0, 1)
		case "week":
			current = current.AddDate(0, 0, 7)
		case "month":
			current = current.AddDate(0, 1, 0)
		case "quarter":
			current = current.AddDate(0, 3, 0)
		case "year":
			current = current.AddDate(1, 0, 0)
		}
	}

	return points
}

// generateComparisons generates comparisons with other templates
func (uc *AnalyzeTemplatePerformanceUseCase) generateComparisons(
	ctx context.Context,
	templateID uuid.UUID,
	compareWithIDs []string,
	currentMetrics map[string]response.MetricValueResponse,
) []response.TemplateComparisonResponse {
	var comparisons []response.TemplateComparisonResponse

	for _, compareIDStr := range compareWithIDs {
		compareID, err := uuid.FromString(compareIDStr)
		if err != nil {
			continue
		}

		// Get comparison template
		compareTemplate, err := uc.aiTemplateRepo.FindByID(ctx, compareID)
		if err != nil {
			continue
		}

		// Get comparison metrics
		compareMetrics, err := uc.domainService.CalculateTemplatePerformance(ctx, compareID)
		if err != nil {
			continue
		}

		// Calculate differences
		differences := make(map[string]float64)
		winners := make(map[string]string)

		for metricType, currentValue := range currentMetrics {
			if compareValue, exists := compareMetrics[metricType]; exists {
				diff := currentValue.Value - compareValue
				differences[metricType] = diff

				// Determine winner based on metric type
				if uc.isHigherBetter(metricType) {
					if diff > 0 {
						winners[metricType] = templateID.String()
					} else {
						winners[metricType] = compareID.String()
					}
				} else {
					if diff < 0 {
						winners[metricType] = templateID.String()
					} else {
						winners[metricType] = compareID.String()
					}
				}
			}
		}

		comparisons = append(comparisons, response.TemplateComparisonResponse{
			TemplateID:   compareID,
			TemplateName: compareTemplate.Name,
			Metrics:      compareMetrics,
			Difference:   differences,
			Winner:       winners,
		})
	}

	return comparisons
}

// generateRecommendations generates performance recommendations
func (uc *AnalyzeTemplatePerformanceUseCase) generateRecommendations(
	metrics map[string]response.MetricValueResponse,
	template interface{},
) []response.RecommendationResponse {
	var recommendations []response.RecommendationResponse

	// Check modification rate
	if modRate, exists := metrics[string(value_object.MetricTypeModificationRate)]; exists {
		if modRate.Value > 0.3 {
			recommendations = append(recommendations, *uc.mapper.ToRecommendationResponse(
				"improvement",
				"high",
				"High Modification Rate Detected",
				fmt.Sprintf("%.0f%% of products are being modified after template application", modRate.Value*100),
				"Consider updating template based on user feedback",
				[]string{string(value_object.MetricTypeModificationRate)},
			))
		}
	}

	// Check satisfaction score
	if satScore, exists := metrics[string(value_object.MetricTypeSatisfactionScore)]; exists {
		if satScore.Value < 0.7 {
			recommendations = append(recommendations, *uc.mapper.ToRecommendationResponse(
				"warning",
				"high",
				"Low Satisfaction Score",
				fmt.Sprintf("User satisfaction is at %.0f%%, below the target of 70%%", satScore.Value*100),
				"Analyze user feedback and regenerate template with improvements",
				[]string{string(value_object.MetricTypeSatisfactionScore)},
			))
		}
	}

	// Check usage count
	if usage, exists := metrics[string(value_object.MetricTypeUsageCount)]; exists {
		if usage.Value < 5 {
			recommendations = append(recommendations, *uc.mapper.ToRecommendationResponse(
				"opportunity",
				"medium",
				"Low Template Usage",
				"This template has been used fewer than 5 times",
				"Promote template to increase adoption or review if it meets business needs",
				[]string{string(value_object.MetricTypeUsageCount)},
			))
		}
	}

	// Sort recommendations by priority
	sort.Slice(recommendations, func(i, j int) bool {
		priorityOrder := map[string]int{"high": 0, "medium": 1, "low": 2}
		return priorityOrder[recommendations[i].Priority] < priorityOrder[recommendations[j].Priority]
	})

	return recommendations
}

// Helper methods

func (uc *AnalyzeTemplatePerformanceUseCase) calculatePerformanceScore(
	metrics map[string]response.MetricValueResponse,
) float64 {
	var totalScore float64
	var count float64

	// Weight different metrics
	weights := map[string]float64{
		string(value_object.MetricTypeSatisfactionScore): 0.4,
		string(value_object.MetricTypeModificationRate):  0.3,
		string(value_object.MetricTypeProductRetention):  0.2,
		string(value_object.MetricTypeUsageCount):        0.1,
	}

	for metricType, weight := range weights {
		if metric, exists := metrics[metricType]; exists {
			// Normalize based on metric type
			normalizedValue := metric.Value
			if metricType == string(value_object.MetricTypeModificationRate) {
				// Lower modification rate is better
				normalizedValue = 1.0 - metric.Value
			}
			totalScore += normalizedValue * weight
			count += weight
		}
	}

	if count > 0 {
		return totalScore / count
	}
	return 0.5 // Default middle score
}

func (uc *AnalyzeTemplatePerformanceUseCase) getDefaultMetricValue(metricType string) float64 {
	defaults := map[string]float64{
		string(value_object.MetricTypeUsageCount):        0,
		string(value_object.MetricTypeSatisfactionScore): 0.5,
		string(value_object.MetricTypeModificationRate):  0.5,
		string(value_object.MetricTypeProductRetention):  0.5,
		string(value_object.MetricTypeAdoptionRate):      0.5,
	}
	if value, exists := defaults[metricType]; exists {
		return value
	}
	return 0
}

func (uc *AnalyzeTemplatePerformanceUseCase) getMetricTarget(metricType string) *float64 {
	targets := map[string]float64{
		string(value_object.MetricTypeSatisfactionScore): 0.8,
		string(value_object.MetricTypeModificationRate):  0.2,
		string(value_object.MetricTypeProductRetention):  0.8,
		string(value_object.MetricTypeAdoptionRate):      0.7,
	}
	if target, exists := targets[metricType]; exists {
		return &target
	}
	return nil
}

func (uc *AnalyzeTemplatePerformanceUseCase) isHigherBetter(metricType string) bool {
	// For most metrics, higher is better
	// Exception: modification rate (lower is better)
	return metricType != string(value_object.MetricTypeModificationRate)
}

func (uc *AnalyzeTemplatePerformanceUseCase) formatTimeLabel(t time.Time, groupBy string) string {
	switch groupBy {
	case "day":
		return t.Format("2006-01-02")
	case "week":
		year, week := t.ISOWeek()
		return fmt.Sprintf("%d-W%02d", year, week)
	case "month":
		return t.Format("2006-01")
	case "quarter":
		quarter := (t.Month()-1)/3 + 1
		return fmt.Sprintf("%d-Q%d", t.Year(), quarter)
	case "year":
		return t.Format("2006")
	default:
		return t.Format("2006-01-02")
	}
}