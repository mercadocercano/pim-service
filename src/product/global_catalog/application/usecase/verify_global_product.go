package usecase

import (
	"context"

	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

// VerifyGlobalProductRequest identifica el producto a verificar
type VerifyGlobalProductRequest struct {
	ID string `json:"id"`
}

// VerifyGlobalProduct implementa el caso de uso de verificar un producto global
type VerifyGlobalProduct struct {
	globalProductRepository port.GlobalProductRepository
}

// NewVerifyGlobalProduct crea una nueva instancia del caso de uso
func NewVerifyGlobalProduct(globalProductRepository port.GlobalProductRepository) *VerifyGlobalProduct {
	return &VerifyGlobalProduct{globalProductRepository: globalProductRepository}
}

// Execute verifica el producto con el ID dado y lo persiste
func (uc *VerifyGlobalProduct) Execute(ctx context.Context, request VerifyGlobalProductRequest) error {
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

	product.Verify()

	if _, err := uc.globalProductRepository.Update(product); err != nil {
		return exception.NewInternalError("Error al verificar el producto", err)
	}

	return nil
}
