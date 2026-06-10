package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

// DeleteGlobalProductRequest contiene los datos para eliminar un producto global
type DeleteGlobalProductRequest struct {
	ID string `json:"id"`
}

// DeleteGlobalProduct implementa el caso de uso de eliminar un producto global
type DeleteGlobalProduct struct {
	globalProductRepository port.GlobalProductRepository
}

// NewDeleteGlobalProduct crea una nueva instancia del caso de uso
func NewDeleteGlobalProduct(globalProductRepository port.GlobalProductRepository) *DeleteGlobalProduct {
	return &DeleteGlobalProduct{
		globalProductRepository: globalProductRepository,
	}
}

// Execute ejecuta el caso de uso
func (uc *DeleteGlobalProduct) Execute(ctx context.Context, request DeleteGlobalProductRequest) error {
	if request.ID == "" {
		return exception.NewValidationError("ID del producto es obligatorio", nil)
	}

	existing, err := uc.globalProductRepository.FindByID(request.ID)
	if err != nil {
		return exception.NewInternalError("Error al buscar el producto", err)
	}
	if existing == nil {
		return exception.NewGlobalProductNotFoundByID(request.ID)
	}

	tenantCount, err := uc.globalProductRepository.CountTenantLinks(ctx, request.ID)
	if err != nil {
		return exception.NewInternalError("Error al verificar uso del producto", err)
	}
	if tenantCount > 0 {
		return exception.NewConflictError(
			fmt.Sprintf("Producto en uso por %d tenants, no se puede eliminar", tenantCount),
		)
	}

	if err := uc.globalProductRepository.Delete(request.ID); err != nil {
		return exception.NewInternalError("Error al eliminar el producto", err)
	}

	return nil
}
