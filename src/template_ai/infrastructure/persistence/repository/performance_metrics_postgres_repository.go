package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/exception"
	"saas-mt-pim-service/src/template_ai/domain/port"
)

// PerformanceMetricsPostgresRepository implements port.PerformanceMetricsRepository
type PerformanceMetricsPostgresRepository struct {
	db *sql.DB
}

// NewPerformanceMetricsPostgresRepository creates a new instance
func NewPerformanceMetricsPostgresRepository(db *sql.DB) port.PerformanceMetricsRepository {
	return &PerformanceMetricsPostgresRepository{
		db: db,
	}
}

// SaveMetric saves a performance metric
func (r *PerformanceMetricsPostgresRepository) SaveMetric(ctx context.Context, metric *entity.PerformanceMetric) error {
	query := `
		INSERT INTO ai_template_performance_metrics 
		(id, template_id, metric_type, metric_value, metric_metadata, period_start, period_end, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (template_id, metric_type, period_start, period_end) 
		DO UPDATE SET 
			metric_value = EXCLUDED.metric_value,
			metric_metadata = EXCLUDED.metric_metadata
	`

	metadataJSON, err := json.Marshal(metric.MetricMetadata)
	if err != nil {
		return exception.NewInternalError("failed to marshal metadata: " + err.Error())
	}

	_, err = r.db.ExecContext(ctx, query,
		metric.ID,
		metric.TemplateID,
		metric.MetricType,
		metric.MetricValue,
		metadataJSON,
		metric.PeriodStart,
		metric.PeriodEnd,
		metric.CreatedAt,
	)

	if err != nil {
		return exception.NewInternalError("failed to save metric: " + err.Error())
	}

	return nil
}

// GetMetricsByTemplate retrieves metrics for a template
func (r *PerformanceMetricsPostgresRepository) GetMetricsByTemplate(ctx context.Context, templateID uuid.UUID, startDate, endDate time.Time) ([]*entity.PerformanceMetric, error) {
	query := `
		SELECT id, template_id, metric_type, metric_value, metric_metadata, 
			   period_start, period_end, created_at
		FROM ai_template_performance_metrics
		WHERE template_id = $1 
		AND period_start >= $2 
		AND period_end <= $3
		ORDER BY period_start DESC
	`

	rows, err := r.db.QueryContext(ctx, query, templateID, startDate, endDate)
	if err != nil {
		return nil, exception.NewInternalError("failed to query metrics: " + err.Error())
	}
	defer rows.Close()

	var metrics []*entity.PerformanceMetric
	for rows.Next() {
		metric := &entity.PerformanceMetric{}
		var metadataJSON []byte

		err := rows.Scan(
			&metric.ID,
			&metric.TemplateID,
			&metric.MetricType,
			&metric.MetricValue,
			&metadataJSON,
			&metric.PeriodStart,
			&metric.PeriodEnd,
			&metric.CreatedAt,
		)
		if err != nil {
			return nil, exception.NewInternalError("failed to scan metric: " + err.Error())
		}

		if err := json.Unmarshal(metadataJSON, &metric.MetricMetadata); err != nil {
			return nil, exception.NewInternalError("failed to unmarshal metadata: " + err.Error())
		}

		metrics = append(metrics, metric)
	}

	return metrics, nil
}

// GetAggregatedMetrics retrieves aggregated metrics
func (r *PerformanceMetricsPostgresRepository) GetAggregatedMetrics(ctx context.Context, templateID uuid.UUID, metricType string) (map[string]float64, error) {
	query := `
		SELECT 
			AVG(metric_value) as avg_value,
			MIN(metric_value) as min_value,
			MAX(metric_value) as max_value,
			COUNT(*) as count
		FROM ai_template_performance_metrics
		WHERE template_id = $1 AND metric_type = $2
	`

	var avgValue, minValue, maxValue sql.NullFloat64
	var count int

	err := r.db.QueryRowContext(ctx, query, templateID, metricType).Scan(
		&avgValue, &minValue, &maxValue, &count,
	)
	if err != nil {
		return nil, exception.NewInternalError("failed to get aggregated metrics: " + err.Error())
	}

	result := make(map[string]float64)
	if avgValue.Valid {
		result["average"] = avgValue.Float64
	}
	if minValue.Valid {
		result["min"] = minValue.Float64
	}
	if maxValue.Valid {
		result["max"] = maxValue.Float64
	}
	result["count"] = float64(count)

	return result, nil
}

// GetLatestMetrics retrieves the most recent metrics for a template
func (r *PerformanceMetricsPostgresRepository) GetLatestMetrics(ctx context.Context, templateID uuid.UUID) (map[string]*entity.PerformanceMetric, error) {
	query := `
		SELECT DISTINCT ON (metric_type) 
			id, template_id, metric_type, metric_value, metric_metadata, 
			period_start, period_end, created_at
		FROM ai_template_performance_metrics
		WHERE template_id = $1
		ORDER BY metric_type, period_end DESC
	`

	rows, err := r.db.QueryContext(ctx, query, templateID)
	if err != nil {
		return nil, exception.NewInternalError("failed to query latest metrics: " + err.Error())
	}
	defer rows.Close()

	metricsMap := make(map[string]*entity.PerformanceMetric)
	for rows.Next() {
		metric := &entity.PerformanceMetric{}
		var metadataJSON []byte

		err := rows.Scan(
			&metric.ID,
			&metric.TemplateID,
			&metric.MetricType,
			&metric.MetricValue,
			&metadataJSON,
			&metric.PeriodStart,
			&metric.PeriodEnd,
			&metric.CreatedAt,
		)
		if err != nil {
			return nil, exception.NewInternalError("failed to scan metric: " + err.Error())
		}

		if err := json.Unmarshal(metadataJSON, &metric.MetricMetadata); err != nil {
			return nil, exception.NewInternalError("failed to unmarshal metadata: " + err.Error())
		}

		metricsMap[metric.MetricType] = metric
	}

	return metricsMap, nil
}