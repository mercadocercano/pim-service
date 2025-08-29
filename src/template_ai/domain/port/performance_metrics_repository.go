package port

import (
	"context"
	"time"
	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
)

// PerformanceMetricsRepository defines the interface for performance metrics persistence
type PerformanceMetricsRepository interface {
	SaveMetric(ctx context.Context, metric *entity.PerformanceMetric) error
	GetMetricsByTemplate(ctx context.Context, templateID uuid.UUID, startDate, endDate time.Time) ([]*entity.PerformanceMetric, error)
	GetAggregatedMetrics(ctx context.Context, templateID uuid.UUID, metricType string) (map[string]float64, error)
	GetLatestMetrics(ctx context.Context, templateID uuid.UUID) (map[string]*entity.PerformanceMetric, error)
}