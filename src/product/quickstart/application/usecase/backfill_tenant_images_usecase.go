package usecase

import (
	"context"
	"fmt"
	"log"

	globalPort "saas-mt-pim-service/src/product/global_catalog/domain/port"
	tenantPort "saas-mt-pim-service/src/product/tenant/domain/port"
)

// BackfillResult resume el resultado del backfill de imágenes para un tenant.
type BackfillResult struct {
	Updated int `json:"updated"`
	Skipped int `json:"skipped"`
	Errors  int `json:"errors"`
}

// BackfillTenantImagesUseCase copia image_url desde global_products a los
// productos del tenant que no tienen imagen, usando match por nombre y marca.
type BackfillTenantImagesUseCase struct {
	globalRepo globalPort.GlobalProductRepository
	tenantRepo tenantPort.ProductRepository
}

// NewBackfillTenantImagesUseCase crea el use case con sus dependencias.
func NewBackfillTenantImagesUseCase(
	globalRepo globalPort.GlobalProductRepository,
	tenantRepo tenantPort.ProductRepository,
) *BackfillTenantImagesUseCase {
	return &BackfillTenantImagesUseCase{
		globalRepo: globalRepo,
		tenantRepo: tenantRepo,
	}
}

// Execute realiza el backfill de imágenes para un tenant específico.
func (uc *BackfillTenantImagesUseCase) Execute(ctx context.Context, tenantID string) (*BackfillResult, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenantID es requerido")
	}

	products, err := uc.tenantRepo.FindWithoutImage(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("buscando productos sin imagen para tenant %s: %w", tenantID, err)
	}

	result := &BackfillResult{}

	for _, product := range products {
		brandName := ""
		if product.HasBrand() {
			brandName = product.BrandReference().Name()
		}

		globalProduct, err := uc.globalRepo.FindByNameAndBrand(ctx, product.Name(), brandName)
		if err != nil {
			log.Printf("[backfill] error buscando en global_products para product_id=%s: %v", product.IDString(), err)
			result.Errors++
			continue
		}

		if globalProduct == nil || globalProduct.ImageURL() == nil {
			result.Skipped++
			continue
		}

		if updateErr := uc.tenantRepo.UpdateImageURL(ctx, tenantID, product.IDString(), *globalProduct.ImageURL()); updateErr != nil {
			log.Printf("[backfill] error actualizando image_url product_id=%s: %v", product.IDString(), updateErr)
			result.Errors++
			continue
		}

		log.Printf("[backfill] imagen actualizada product_id=%s tenant=%s", product.IDString(), tenantID)
		result.Updated++
	}

	return result, nil
}

// ExecuteAll realiza el backfill para todos los tenants que tienen productos.
func (uc *BackfillTenantImagesUseCase) ExecuteAll(ctx context.Context) (map[string]*BackfillResult, error) {
	tenantIDs, err := uc.tenantRepo.FindDistinctTenantIDs(ctx)
	if err != nil {
		return nil, fmt.Errorf("obteniendo tenant IDs: %w", err)
	}

	results := make(map[string]*BackfillResult, len(tenantIDs))

	for _, tenantID := range tenantIDs {
		result, err := uc.Execute(ctx, tenantID)
		if err != nil {
			log.Printf("[backfill] error procesando tenant %s: %v", tenantID, err)
			results[tenantID] = &BackfillResult{Errors: 1}
			continue
		}
		results[tenantID] = result
	}

	return results, nil
}
