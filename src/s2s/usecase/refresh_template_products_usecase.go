package usecase

import (
	"context"
	"fmt"

	pimport "saas-mt-pim-service/src/pim/domain/port"
	"saas-mt-pim-service/src/s2s/domain/port"
)

// RefreshTemplateProductsUseCase re-calcula business_type_product_templates
// desde global_products verificados. Idempotente — seguro de llamar repetidamente.
type RefreshTemplateProductsUseCase struct {
	repo   port.TemplateRepository
	logger pimport.PIMEventLogger
}

func NewRefreshTemplateProductsUseCase(repo port.TemplateRepository) *RefreshTemplateProductsUseCase {
	return &RefreshTemplateProductsUseCase{repo: repo}
}

// NewRefreshTemplateProductsUseCaseWithLogger crea el use case con logger canónico inyectado.
func NewRefreshTemplateProductsUseCaseWithLogger(
	repo port.TemplateRepository,
	logger pimport.PIMEventLogger,
) *RefreshTemplateProductsUseCase {
	return &RefreshTemplateProductsUseCase{repo: repo, logger: logger}
}

// logEvent emite un evento canónico si hay logger inyectado (nil-safe).
func (uc *RefreshTemplateProductsUseCase) logEvent(e pimport.PIMEvent) {
	if uc.logger != nil {
		uc.logger.Log(e)
	}
}

type RefreshResult struct {
	TemplatesUpdated int `json:"templates_updated"`
}

func (uc *RefreshTemplateProductsUseCase) Execute(ctx context.Context) (*RefreshResult, error) {
	rowsAffected, err := uc.repo.RefreshProductTemplates(ctx)
	if err != nil {
		return nil, fmt.Errorf("refresh template products: %w", err)
	}

	uc.logEvent(pimport.PIMEvent{
		Event: "pim.template_refresh_completed",
		Count: int(rowsAffected),
	})

	return &RefreshResult{TemplatesUpdated: int(rowsAffected)}, nil
}
