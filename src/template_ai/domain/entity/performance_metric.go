package entity

import (
	"time"
	"github.com/gofrs/uuid/v5"
)

// PerformanceMetric represents a performance metric for a template
type PerformanceMetric struct {
	ID             uuid.UUID
	TemplateID     uuid.UUID
	MetricType     string
	MetricValue    float64
	MetricMetadata map[string]interface{}
	PeriodStart    time.Time
	PeriodEnd      time.Time
	CreatedAt      time.Time
}