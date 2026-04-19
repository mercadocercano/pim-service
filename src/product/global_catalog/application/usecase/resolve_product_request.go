package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"

	"github.com/google/uuid"
)

type ResolveProductRequestUseCase struct {
	repo port.ProductRequestRepository
}

func NewResolveProductRequestUseCase(repo port.ProductRequestRepository) *ResolveProductRequestUseCase {
	return &ResolveProductRequestUseCase{repo: repo}
}

type ResolveProductRequestRequest struct {
	RequestID       string  `json:"-"`
	Action          string  `json:"action" binding:"required"`
	AdminNotes      *string `json:"admin_notes,omitempty"`
	GlobalProductID *string `json:"global_product_id,omitempty"`
}

type ResolveProductRequestResponse struct {
	ID              string  `json:"id"`
	Status          string  `json:"status"`
	AdminNotes      *string `json:"admin_notes,omitempty"`
	GlobalProductID *string `json:"global_product_id,omitempty"`
}

func (uc *ResolveProductRequestUseCase) Execute(ctx context.Context, req ResolveProductRequestRequest) (*ResolveProductRequestResponse, error) {
	productReq, err := uc.repo.FindByID(ctx, req.RequestID)
	if err != nil {
		return nil, fmt.Errorf("solicitud no encontrada: %w", err)
	}

	if productReq.Status() != entity.RequestStatusPending && productReq.Status() != entity.RequestStatusApproved {
		return nil, fmt.Errorf("la solicitud ya fue resuelta con status: %s", productReq.Status())
	}

	notes := ""
	if req.AdminNotes != nil {
		notes = *req.AdminNotes
	}

	switch req.Action {
	case "approve":
		productReq.Approve(notes)
	case "reject":
		productReq.Reject(notes)
	case "fulfill":
		if req.GlobalProductID == nil {
			return nil, fmt.Errorf("global_product_id es requerido para fulfill")
		}
		gpID, err := uuid.Parse(*req.GlobalProductID)
		if err != nil {
			return nil, fmt.Errorf("global_product_id inválido: %w", err)
		}
		productReq.Fulfill(gpID)
	default:
		return nil, fmt.Errorf("acción inválida: %s (usar approve, reject o fulfill)", req.Action)
	}

	if err := uc.repo.Update(ctx, productReq); err != nil {
		return nil, fmt.Errorf("error actualizando solicitud: %w", err)
	}

	resp := &ResolveProductRequestResponse{
		ID:     productReq.IDString(),
		Status: string(productReq.Status()),
	}
	if productReq.AdminNotes() != nil {
		resp.AdminNotes = productReq.AdminNotes()
	}
	if productReq.GlobalProductID() != nil {
		gpIDStr := productReq.GlobalProductID().String()
		resp.GlobalProductID = &gpIDStr
	}

	return resp, nil
}
