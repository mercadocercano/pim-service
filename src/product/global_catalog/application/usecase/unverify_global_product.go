package usecase

import (
	"context"

	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

// UnverifyGlobalProductRequest identifica el producto a desverificar
type UnverifyGlobalProductRequest struct {
	ID string `json:"id"`
}

// UnverifyGlobalProduct implementa el caso de uso de desverificar un producto global
type UnverifyGlobalProduct struct {
	globalProductRepository port.GlobalProductRepository
}

// NewUnverifyGlobalProduct crea una nueva instancia del caso de uso
func NewUnverifyGlobalProduct(globalProductRepository port.GlobalProductRepository) *UnverifyGlobalProduct {
	return &UnverifyGlobalProduct{globalProductRepository: globalProductRepository}
}

// Execute desverifica el producto con el ID dado y lo persiste
func (uc *UnverifyGlobalProduct) Execute(ctx context.Context, request UnverifyGlobalProductRequest) error {
	if request.ID == "" {
		return exception.NewValidationError("ID del producto es obligatorio", nil)
	}

	product, err := uc.globalProductRepository.FindByID(request.ID)
	if err != nil {
		return exception.NewInternalError("Error al buscar el producto", err)
	}
	if product == nil {
		return exception.NewGlobalProductNotFoundByID(request.ID)
	}

	product.Unverify()

	if _, err := uc.globalProductRepository.Update(product); err != nil {
		return exception.NewInternalError("Error al desverificar el producto", err)
	}

	return nil
}
