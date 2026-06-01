package usecase

import (
	"context"
	"fmt"
	"time"

	"saas-mt-pim-service/src/businesstype/domain/port"
)

// TemplateAnalytics contiene las métricas de uso de un template
type TemplateAnalytics struct {
	TemplateID     string     `json:"template_id"`
	TenantsUsed    int        `json:"tenants_used"`
	LastActivated  *time.Time `json:"last_activated"`
	CompletionRate float64    `json:"completion_rate"`
}

// TemplateAnalyticsRepository define el puerto para obtener analytics
type TemplateAnalyticsRepository interface {
	GetTemplateAnalytics(ctx context.Context, templateID string) (*TemplateAnalytics, error)
}

// GetTemplateAnalyticsUseCase obtiene analíticas de uso de un template
type GetTemplateAnalyticsUseCase struct {
	templateRepo  port.BusinessTypeTemplateRepository
	analyticsRepo TemplateAnalyticsRepository
}

// NewGetTemplateAnalyticsUseCase crea una nueva instancia del caso de uso
func NewGetTemplateAnalyticsUseCase(
	templateRepo port.BusinessTypeTemplateRepository,
	analyticsRepo TemplateAnalyticsRepository,
) *GetTemplateAnalyticsUseCase {
	return &GetTemplateAnalyticsUseCase{
		templateRepo:  templateRepo,
		analyticsRepo: analyticsRepo,
	}
}

// Execute obtiene las analíticas de un template por ID
func (uc *GetTemplateAnalyticsUseCase) Execute(ctx context.Context, templateID string) (*TemplateAnalytics, error) {
	if templateID == "" {
		return nil, fmt.Errorf("template_id es requerido")
	}

	template, err := uc.templateRepo.FindByID(ctx, templateID)
	if err != nil {
		return nil, fmt.Errorf("error verificando template: %w", err)
	}
	if template == nil {
		return nil, fmt.Errorf("template no encontrado")
	}

	analytics, err := uc.analyticsRepo.GetTemplateAnalytics(ctx, templateID)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo analytics: %w", err)
	}

	return analytics, nil
}
