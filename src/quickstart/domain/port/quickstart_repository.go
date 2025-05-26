package port

import (
	"context"

	"pim/src/quickstart/domain/entity"
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

// YamlDataLoader define las operaciones para cargar datos desde archivos YAML
type YamlDataLoader interface {
	LoadBusinessTypes(ctx context.Context) ([]*entity.BusinessType, error)
	LoadCategoriesByBusinessType(ctx context.Context, businessType string) (interface{}, error)
	LoadAttributesByBusinessType(ctx context.Context, businessType string) (interface{}, error)
	LoadVariantsByBusinessType(ctx context.Context, businessType string) (interface{}, error)
	LoadProductsByBusinessType(ctx context.Context, businessType string) (interface{}, error)
	LoadBrandsByBusinessType(ctx context.Context, businessType string) (interface{}, error)
}
