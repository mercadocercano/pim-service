package port

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
)

// AITemplateRepository defines the interface for AI template persistence
type AITemplateRepository interface {
	// Template operations
	Save(ctx context.Context, template *entity.AITemplate) error
	Update(ctx context.Context, template *entity.AITemplate) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.AITemplate, error)
	FindByBusinessTypeAndTenant(ctx context.Context, businessTypeID uuid.UUID, tenantID *uuid.UUID) ([]*entity.AITemplate, error)
	Delete(ctx context.Context, id uuid.UUID) error

	// Template-Product associations
	SaveTemplateProduct(ctx context.Context, templateProduct *entity.TemplateGlobalProduct) error
	UpdateTemplateProduct(ctx context.Context, templateProduct *entity.TemplateGlobalProduct) error
	FindTemplateProducts(ctx context.Context, templateID uuid.UUID) ([]*entity.TemplateGlobalProduct, error)
	DeleteTemplateProduct(ctx context.Context, templateID, globalProductID uuid.UUID) error

	// Generation history
	SaveGenerationHistory(ctx context.Context, history *entity.AIGenerationHistory) error
	FindGenerationHistory(ctx context.Context, templateID *uuid.UUID, tenantID *uuid.UUID, limit int) ([]*entity.AIGenerationHistory, error)

	// Performance metrics
	SavePerformanceMetric(ctx context.Context, metric *entity.AIPerformanceMetric) error
	FindPerformanceMetrics(ctx context.Context, templateID uuid.UUID, metricType string) ([]*entity.AIPerformanceMetric, error)
	GetAggregatedMetrics(ctx context.Context, templateID uuid.UUID) (map[string]float64, error)

	// Product feedback
	SaveProductFeedback(ctx context.Context, feedback *entity.AIProductFeedback) error
	FindProductFeedback(ctx context.Context, templateID *uuid.UUID, tenantID uuid.UUID) ([]*entity.AIProductFeedback, error)
	GetFeedbackSummary(ctx context.Context, templateID uuid.UUID) (map[string]int, error)
}

// GlobalProductRepository defines the interface for accessing global products
type GlobalProductRepository interface {
	FindByIDs(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]interface{}, error)
	SearchByCategory(ctx context.Context, categoryID string, limit int) ([]interface{}, error)
	SearchByBrand(ctx context.Context, brandID string, limit int) ([]interface{}, error)
	GetFilteredProducts(ctx context.Context, filters map[string]interface{}) ([]*entity.GlobalProduct, error)
}