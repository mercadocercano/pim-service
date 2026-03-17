package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	brandPort "saas-mt-pim-service/src/brand/domain/port"
	categoryPort "saas-mt-pim-service/src/category/domain/port"
	"saas-mt-pim-service/src/product/tenant/domain/port"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"
)

type BulkUpdateProductsRequest struct {
	TenantID   string   `json:"-"`
	ProductIDs []string `json:"product_ids" binding:"required,min=1"`
	BrandID    *string  `json:"brand_id,omitempty"`
	CategoryID *string  `json:"category_id,omitempty"`
}

type BulkUpdateProductsResponse struct {
	Success         bool     `json:"success"`
	UpdatedCount    int      `json:"updated_count"`
	SkippedCount    int      `json:"skipped_count"`
	Errors          []string `json:"errors,omitempty"`
}

type BulkUpdateProductsUseCase struct {
	productRepo  port.ProductCriteriaRepository
	brandRepo    brandPort.BrandRepository
	categoryRepo categoryPort.CategoryRepository
}

func NewBulkUpdateProductsUseCase(
	productRepo port.ProductCriteriaRepository,
	brandRepo brandPort.BrandRepository,
	categoryRepo categoryPort.CategoryRepository,
) *BulkUpdateProductsUseCase {
	return &BulkUpdateProductsUseCase{
		productRepo:  productRepo,
		brandRepo:    brandRepo,
		categoryRepo: categoryRepo,
	}
}

func (uc *BulkUpdateProductsUseCase) Execute(ctx context.Context, req BulkUpdateProductsRequest) (*BulkUpdateProductsResponse, error) {
	if req.BrandID == nil && req.CategoryID == nil {
		return nil, fmt.Errorf("at least one of brand_id or category_id must be provided")
	}

	// Validate brand exists in tenant
	var brandRef *value_object.BrandReference
	if req.BrandID != nil && *req.BrandID != "" {
		brand, err := uc.brandRepo.FindByID(ctx, *req.BrandID, req.TenantID)
		if err != nil {
			return nil, fmt.Errorf("brand_id '%s' not found in tenant", *req.BrandID)
		}
		brandRef, _ = value_object.NewBrandReference(brand.ID, brand.Name)
	}

	// Validate category exists in tenant
	var categoryRef *value_object.CategoryReference
	if req.CategoryID != nil && *req.CategoryID != "" {
		cat, err := uc.categoryRepo.FindByID(ctx, *req.CategoryID, req.TenantID)
		if err != nil {
			return nil, fmt.Errorf("category_id '%s' not found in tenant", *req.CategoryID)
		}
		categoryRef, _ = value_object.NewCategoryReference(cat.ID, cat.Name)
	}

	response := &BulkUpdateProductsResponse{
		Success: true,
		Errors:  make([]string, 0),
	}

	for _, productIDStr := range req.ProductIDs {
		productID, err := uuid.Parse(productIDStr)
		if err != nil {
			response.SkippedCount++
			response.Errors = append(response.Errors, fmt.Sprintf("invalid product ID: %s", productIDStr))
			continue
		}

		product, err := uc.productRepo.FindByID(ctx, productID, req.TenantID)
		if err != nil {
			response.SkippedCount++
			continue
		}

		newBrandRef := product.BrandReference()
		if brandRef != nil {
			newBrandRef = brandRef
		}
		newCategoryRef := product.CategoryReference()
		if categoryRef != nil {
			newCategoryRef = categoryRef
		}

		if err := product.Update(product.Name(), product.Description(), product.SKU(), newCategoryRef, newBrandRef); err != nil {
			response.SkippedCount++
			response.Errors = append(response.Errors, fmt.Sprintf("failed to update product %s: %v", productIDStr, err))
			continue
		}

		if err := uc.productRepo.Update(ctx, product); err != nil {
			response.Errors = append(response.Errors, fmt.Sprintf("failed to update product %s: %v", productIDStr, err))
			response.SkippedCount++
			continue
		}

		response.UpdatedCount++
	}

	if response.UpdatedCount == 0 {
		response.Success = false
	}

	return response, nil
}
