package usecase

import (
	"context"
	"fmt"

	"pim/src/product/tenant/domain/port"
	"pim/src/product/tenant/domain/service"

	"github.com/google/uuid"
)

// UpdateProductStatusUseCase caso de uso para actualizar el estado de un producto
type UpdateProductStatusUseCase struct {
	productRepository    port.ProductRepository
	productStatusService *service.ProductStatusService
}

// NewUpdateProductStatusUseCase crea una nueva instancia del caso de uso
func NewUpdateProductStatusUseCase(
	productRepository port.ProductRepository,
	productStatusService *service.ProductStatusService,
) *UpdateProductStatusUseCase {
	return &UpdateProductStatusUseCase{
		productRepository:    productRepository,
		productStatusService: productStatusService,
	}
}

// UpdateProductStatusRequest request para actualizar estado
type UpdateProductStatusRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	TenantID  string    `json:"tenant_id"`
	NewStatus string    `json:"new_status"`
}

// UpdateProductStatusResponse response de actualización de estado
type UpdateProductStatusResponse struct {
	ProductID            uuid.UUID `json:"product_id"`
	PreviousStatus       string    `json:"previous_status"`
	NewStatus            string    `json:"new_status"`
	AvailableTransitions []string  `json:"available_transitions"`
	Message              string    `json:"message"`
}

// Execute ejecuta el caso de uso de actualización de estado
func (uc *UpdateProductStatusUseCase) Execute(
	ctx context.Context,
	request UpdateProductStatusRequest,
) (*UpdateProductStatusResponse, error) {
	// Obtener el producto
	product, err := uc.productRepository.FindByID(ctx, request.ProductID, request.TenantID)
	if err != nil {
		return nil, fmt.Errorf("producto no encontrado: %w", err)
	}

	// Guardar estado anterior
	previousStatus := product.Status().Value()

	// Intentar la transición de estado
	if err := uc.productStatusService.TransitionTo(product, request.NewStatus); err != nil {
		return nil, fmt.Errorf("error en transición de estado: %w", err)
	}

	// Guardar el producto actualizado
	if err := uc.productRepository.Update(ctx, product); err != nil {
		return nil, fmt.Errorf("error al guardar producto: %w", err)
	}

	// Obtener transiciones disponibles después del cambio
	availableTransitions := uc.productStatusService.GetAvailableTransitions(product)

	return &UpdateProductStatusResponse{
		ProductID:            product.ID(),
		PreviousStatus:       previousStatus,
		NewStatus:            product.Status().Value(),
		AvailableTransitions: availableTransitions,
		Message:              fmt.Sprintf("Estado actualizado exitosamente de %s a %s", previousStatus, product.Status().Value()),
	}, nil
}

// GetAvailableTransitions retorna las transiciones disponibles para un producto
func (uc *UpdateProductStatusUseCase) GetAvailableTransitions(
	ctx context.Context,
	productID uuid.UUID,
	tenantID string,
) ([]string, error) {
	// Obtener el producto
	product, err := uc.productRepository.FindByID(ctx, productID, tenantID)
	if err != nil {
		return nil, fmt.Errorf("producto no encontrado: %w", err)
	}

	return uc.productStatusService.GetAvailableTransitions(product), nil
}
