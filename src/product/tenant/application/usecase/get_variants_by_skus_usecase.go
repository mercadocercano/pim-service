package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/port"
)

type GetVariantsBySKUsUseCase struct {
	productRepo port.ProductRepository
}

func NewGetVariantsBySKUsUseCase(productRepo port.ProductRepository) *GetVariantsBySKUsUseCase {
	return &GetVariantsBySKUsUseCase{productRepo: productRepo}
}

type GetVariantsBySKUsRequest struct {
	SKUs []string `json:"skus" binding:"required"`
}

type GetVariantsBySKUsResponse struct {
	Variants []response.VariantEnrichedResponse `json:"variants"`
}

func (uc *GetVariantsBySKUsUseCase) Execute(ctx context.Context, req *GetVariantsBySKUsRequest, tenantID string) (*GetVariantsBySKUsResponse, error) {
	if len(req.SKUs) == 0 {
		return &GetVariantsBySKUsResponse{Variants: []response.VariantEnrichedResponse{}}, nil
	}

	if len(req.SKUs) > 1000 {
		return nil, fmt.Errorf("maximum 1000 SKUs per request")
	}

	rows, err := uc.productRepo.FindVariantsEnrichedBySKUs(ctx, tenantID, req.SKUs)
	if err != nil {
		return nil, fmt.Errorf("error finding variants by SKUs: %w", err)
	}

	variants := make([]response.VariantEnrichedResponse, 0, len(rows))
	for _, row := range rows {
		variants = append(variants, response.VariantEnrichedResponse{
			VariantID:    row.VariantID,
			ProductID:    row.ProductID,
			SKU:          row.SKU,
			VariantName:  row.VariantName,
			ProductName:  row.ProductName,
			CategoryID:   row.CategoryID,
			CategoryName: row.CategoryName,
			Price:        row.Price,
		})
	}

	return &GetVariantsBySKUsResponse{Variants: variants}, nil
}
