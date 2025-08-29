package port

import (
	"context"

	"saas-mt-pim-service/src/quickstart/domain/entity"
)

// QuickstartTemplateRepository define las operaciones de persistencia para plantillas de quickstart
type QuickstartTemplateRepository interface {
	Create(ctx context.Context, template *entity.QuickstartTemplate) error
	GetByBusinessTypeAndTemplateType(ctx context.Context, businessType string, templateType entity.TemplateType) ([]*entity.QuickstartTemplate, error)
	GetByID(ctx context.Context, id string) (*entity.QuickstartTemplate, error)
	Update(ctx context.Context, template *entity.QuickstartTemplate) error
	Delete(ctx context.Context, id string) error
}

// TenantQuickstartHistoryRepository define las operaciones de persistencia para el historial de quickstart
type TenantQuickstartHistoryRepository interface {
	Create(ctx context.Context, history *entity.TenantQuickstartHistory) error
	GetByTenantID(ctx context.Context, tenantID string) ([]*entity.TenantQuickstartHistory, error)
	GetByID(ctx context.Context, id string) (*entity.TenantQuickstartHistory, error)
	Update(ctx context.Context, history *entity.TenantQuickstartHistory) error
	GetLatestByTenantID(ctx context.Context, tenantID string) (*entity.TenantQuickstartHistory, error)
}

