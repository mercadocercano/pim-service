package usecase

import (
	"context"
	"fmt"
	"log"

	"saas-mt-pim-service/src/s2s/domain/port"
)

// RefreshTemplateProductsUseCase re-calcula business_type_product_templates
// desde global_products verificados. Idempotente — seguro de llamar repetidamente.
type RefreshTemplateProductsUseCase struct {
	repo port.TemplateRepository
}

func NewRefreshTemplateProductsUseCase(repo port.TemplateRepository) *RefreshTemplateProductsUseCase {
	return &RefreshTemplateProductsUseCase{repo: repo}
}

type RefreshResult struct {
	TemplatesUpdated int `json:"templates_updated"`
}

func (uc *RefreshTemplateProductsUseCase) Execute(ctx context.Context) (*RefreshResult, error) {
	log.Println("[refresh-template-products] Starting refresh from global_products...")

	rowsAffected, err := uc.repo.RefreshProductTemplates(ctx)
	if err != nil {
		return nil, fmt.Errorf("refresh template products: %w", err)
	}

	log.Printf("[refresh-template-products] Done: %d templates updated", rowsAffected)
	return &RefreshResult{TemplatesUpdated: int(rowsAffected)}, nil
}
