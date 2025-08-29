package port

import (
	"context"
	"time"
	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
)

// FeedbackRepository defines the interface for feedback persistence
type FeedbackRepository interface {
	Save(ctx context.Context, feedback *entity.ProductFeedback) error
	GetByTemplate(ctx context.Context, templateID uuid.UUID) ([]*entity.ProductFeedback, error)
	GetByTenant(ctx context.Context, tenantID uuid.UUID, startDate, endDate time.Time) ([]*entity.ProductFeedback, error)
	GetCommonModifications(ctx context.Context, tenantID uuid.UUID, limit int) ([]CommonModification, error)
}