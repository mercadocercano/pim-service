package usecase

import (
	"context"
	"fmt"
	"log"

	globalPort "saas-mt-pim-service/src/product/global_catalog/domain/port"
	tenantEntity "saas-mt-pim-service/src/product/tenant/domain/entity"
	tenantPort "saas-mt-pim-service/src/product/tenant/domain/port"
)

// ImportFromGlobalCatalogUseCase importa un producto individual del catálogo global al catálogo del tenant.
type ImportFromGlobalCatalogUseCase struct {
	tenantProductRepo tenantPort.ProductRepository
	globalCatalogRepo globalPort.GlobalProductRepository
}

func NewImportFromGlobalCatalogUseCase(
	tenantProductRepo tenantPort.ProductRepository,
	globalCatalogRepo globalPort.GlobalProductRepository,
) *ImportFromGlobalCatalogUseCase {
	return &ImportFromGlobalCatalogUseCase{
		tenantProductRepo: tenantProductRepo,
		globalCatalogRepo: globalCatalogRepo,
	}
}

type ImportFromGlobalCatalogRequest struct {
	TenantID        string `json:"tenant_id"`
	GlobalProductID string `json:"global_product_id" binding:"required"`
}

type ImportFromGlobalCatalogResponse struct {
	ProductID       string  `json:"product_id"`
	GlobalProductID string  `json:"global_product_id"`
	Name            string  `json:"name"`
	Brand           *string `json:"brand,omitempty"`
	Category        *string `json:"category,omitempty"`
	ImageURL        *string `json:"image_url,omitempty"`
	Price           *float64 `json:"price,omitempty"`
}

func (uc *ImportFromGlobalCatalogUseCase) Execute(
	ctx context.Context,
	request ImportFromGlobalCatalogRequest,
) (*ImportFromGlobalCatalogResponse, error) {
	if request.TenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}
	if request.GlobalProductID == "" {
		return nil, fmt.Errorf("global_product_id es requerido")
	}

	gp, err := uc.globalCatalogRepo.FindByID(request.GlobalProductID)
	if err != nil {
		return nil, fmt.Errorf("producto global no encontrado: %w", err)
	}
	if !gp.IsActive() {
		return nil, fmt.Errorf("el producto global no está activo")
	}

	categoryRef := buildCategoryRef(gp)
	brandRef := buildBrandRef(gp)

	product, err := tenantEntity.NewProductWithImage(
		request.TenantID,
		gp.Name(),
		gp.Description(),
		gp.ImageURL(),
		nil,
		categoryRef,
		brandRef,
	)
	if err != nil {
		return nil, fmt.Errorf("error creando producto tenant: %w", err)
	}

	if price := gp.Price(); price != nil && *price > 0 {
		if defaultVariant := product.GetDefaultVariant(); defaultVariant != nil {
			defaultVariant.UpdatePrice(*price)
		}
	}

	if err := uc.tenantProductRepo.Save(ctx, product); err != nil {
		return nil, fmt.Errorf("error guardando producto: %w", err)
	}

	log.Printf("[catalog-import] imported global product %s → tenant product %s for tenant %s",
		gp.IDString(), product.IDString(), request.TenantID)

	return &ImportFromGlobalCatalogResponse{
		ProductID:       product.IDString(),
		GlobalProductID: gp.IDString(),
		Name:            gp.Name(),
		Brand:           gp.Brand(),
		Category:        gp.Category(),
		ImageURL:        gp.ImageURL(),
		Price:           gp.Price(),
	}, nil
}
