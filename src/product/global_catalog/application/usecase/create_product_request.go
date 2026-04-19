package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

type CreateProductRequestUseCase struct {
	repo port.ProductRequestRepository
}

func NewCreateProductRequestUseCase(repo port.ProductRequestRepository) *CreateProductRequestUseCase {
	return &CreateProductRequestUseCase{repo: repo}
}

type CreateProductRequestRequest struct {
	TenantID     string  `json:"tenant_id"`
	Name         string  `json:"name" binding:"required"`
	Brand        *string `json:"brand,omitempty"`
	Category     *string `json:"category,omitempty"`
	Description  *string `json:"description,omitempty"`
	BusinessType *string `json:"business_type,omitempty"`
}

type CreateProductRequestResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (uc *CreateProductRequestUseCase) Execute(ctx context.Context, req CreateProductRequestRequest) (*CreateProductRequestResponse, error) {
	if req.TenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}

	productReq, err := entity.NewProductRequest(
		req.TenantID, req.Name, req.Brand, req.Category, req.Description, req.BusinessType,
	)
	if err != nil {
		return nil, fmt.Errorf("datos inválidos: %w", err)
	}

	if err := uc.repo.Save(ctx, productReq); err != nil {
		return nil, fmt.Errorf("error guardando solicitud: %w", err)
	}

	return &CreateProductRequestResponse{
		ID:      productReq.IDString(),
		Name:    productReq.Name(),
		Message: "Estamos buscando este producto. Te avisaremos cuando esté disponible en el catálogo.",
	}, nil
}
