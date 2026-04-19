package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

type ListProductRequestsUseCase struct {
	repo port.ProductRequestRepository
}

func NewListProductRequestsUseCase(repo port.ProductRequestRepository) *ListProductRequestsUseCase {
	return &ListProductRequestsUseCase{repo: repo}
}

type ProductRequestItem struct {
	ID           string  `json:"id"`
	TenantID     string  `json:"tenant_id"`
	Name         string  `json:"name"`
	Brand        *string `json:"brand,omitempty"`
	Category     *string `json:"category,omitempty"`
	Description  *string `json:"description,omitempty"`
	BusinessType *string `json:"business_type,omitempty"`
	Status       string  `json:"status"`
	CreatedAt    string  `json:"created_at"`
}

type ListProductRequestsResponse struct {
	Items        []ProductRequestItem `json:"items"`
	PendingCount int                  `json:"pending_count"`
}

func (uc *ListProductRequestsUseCase) Execute(ctx context.Context, limit, offset int) (*ListProductRequestsResponse, error) {
	if limit <= 0 {
		limit = 50
	}

	requests, err := uc.repo.FindPending(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error listando solicitudes: %w", err)
	}

	count, err := uc.repo.CountPending(ctx)
	if err != nil {
		return nil, fmt.Errorf("error contando solicitudes: %w", err)
	}

	items := make([]ProductRequestItem, len(requests))
	for i, r := range requests {
		items[i] = ProductRequestItem{
			ID:           r.IDString(),
			TenantID:     r.TenantID(),
			Name:         r.Name(),
			Brand:        r.Brand(),
			Category:     r.Category(),
			Description:  r.Description(),
			BusinessType: r.BusinessType(),
			Status:       string(r.Status()),
			CreatedAt:    r.CreatedAt().Format("2006-01-02T15:04:05Z"),
		}
	}

	return &ListProductRequestsResponse{
		Items:        items,
		PendingCount: count,
	}, nil
}
