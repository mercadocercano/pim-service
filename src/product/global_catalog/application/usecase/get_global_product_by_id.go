package usecase

import (
	"context"
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

// GetGlobalProductByIDRequest representa la solicitud para obtener un producto por ID
type GetGlobalProductByIDRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

// GetGlobalProductByIDResponse representa la respuesta con el producto encontrado
type GetGlobalProductByIDResponse struct {
	Product *entity.GlobalProduct `json:"product"`
}

// GetGlobalProductByID use case para obtener un producto del catálogo global por su ID
type GetGlobalProductByID struct {
	globalProductRepository port.GlobalProductRepository
}

// NewGetGlobalProductByID crea una nueva instancia del use case
func NewGetGlobalProductByID(globalProductRepository port.GlobalProductRepository) *GetGlobalProductByID {
	return &GetGlobalProductByID{
		globalProductRepository: globalProductRepository,
	}
}

// Execute ejecuta el caso de uso para obtener un producto por ID
func (uc *GetGlobalProductByID) Execute(ctx context.Context, request GetGlobalProductByIDRequest) (*GetGlobalProductByIDResponse, error) {
	// Validar que el ID no esté vacío
	if request.ID == "" {
		return nil, exception.NewValidationError("ID del producto es obligatorio", nil)
	}

	// Buscar el producto en el repositorio
	product, err := uc.globalProductRepository.FindByID(request.ID)
	if err != nil {
		return nil, exception.NewInternalError("Error al buscar producto por ID", err)
	}

	// Si no se encuentra el producto, devolver error específico
	if product == nil {
		return nil, exception.NewGlobalProductNotFoundByID(request.ID)
	}

	// Devolver el producto encontrado
	return &GetGlobalProductByIDResponse{
		Product: product,
	}, nil
}
