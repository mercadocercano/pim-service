package usecase

import (
	"context"

	"pim/src/product/application/mapper"
	"pim/src/product/application/response"
	"pim/src/product/domain/port"

	"github.com/google/uuid"
)

// GetProductByIDUseCase maneja la obtención de un producto por ID
type GetProductByIDUseCase struct {
	productRepo port.ProductRepository
	mapper      *mapper.ProductMapper
}

// NewGetProductByIDUseCase crea una nueva instancia del caso de uso
func NewGetProductByIDUseCase(
	productRepo port.ProductRepository,
	mapper *mapper.ProductMapper,
) *GetProductByIDUseCase {
	return &GetProductByIDUseCase{
		productRepo: productRepo,
		mapper:      mapper,
	}
}

// Execute ejecuta el caso de uso de obtención de producto por ID
func (uc *GetProductByIDUseCase) Execute(
	ctx context.Context,
	productID uuid.UUID,
	tenantID string,
) (*response.ProductResponse, error) {
	// Buscar el producto
	product, err := uc.productRepo.FindByID(ctx, productID, tenantID)
	if err != nil {
		return nil, err
	}

	// Convertir a respuesta
	return uc.mapper.ToResponse(product), nil
}
