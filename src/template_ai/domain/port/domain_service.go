package port

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
)

// AITemplateDomainServicePort define la interfaz del domain service de templates AI.
type AITemplateDomainServicePort interface {
	ValidateTemplateForCreation(ctx context.Context, template *entity.AITemplate) error
	ValidateProductForTemplate(ctx context.Context, templateProduct *entity.TemplateGlobalProduct) error
	CalculateTemplatePerformance(ctx context.Context, templateID uuid.UUID) (map[string]float64, error)
	ShouldRegenerateTemplate(metrics map[string]float64) bool
	ValidateFeedback(feedback *entity.AIProductFeedback) error
}
