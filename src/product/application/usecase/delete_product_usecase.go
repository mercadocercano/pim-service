package usecase

import (
	"context"

	"pim/src/product/domain/port"
	"pim/src/product/domain/service"

	"github.com/google/uuid"
)

// DeleteProductUseCase maneja la eliminación de productos
type DeleteProductUseCase struct {
	productRepo   port.ProductRepository
	domainService *service.ProductDomainService
}

// NewDeleteProductUseCase crea una nueva instancia del caso de uso
func NewDeleteProductUseCase(
	productRepo port.ProductRepository,
	domainService *service.ProductDomainService,
) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productRepo:   productRepo,
		domainService: domainService,
	}
}

// Execute ejecuta el caso de uso de eliminación de producto
func (uc *DeleteProductUseCase) Execute(
	ctx context.Context,
	productID uuid.UUID,
	tenantID string,
) error {
	// Buscar el producto existente
	product, err := uc.productRepo.FindByID(ctx, productID, tenantID)
	if err != nil {
		return err
	}

	// Validar reglas de negocio para eliminación
	if err := uc.domainService.ValidateForDeletion(ctx, product); err != nil {
		return err
	}

	// Eliminar el producto (soft delete)
	product.Delete()

	// Guardar los cambios
	if err := uc.productRepo.Update(ctx, product); err != nil {
		return err
	}

	return nil
}
