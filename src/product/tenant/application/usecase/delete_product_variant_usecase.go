package usecase

import (
	"context"
	"fmt"

	"pim/src/product/tenant/domain/port"

	"github.com/google/uuid"
)

// DeleteProductVariantUseCase maneja la eliminación de variantes a través del agregado Product
type DeleteProductVariantUseCase struct {
	productRepo port.ProductCriteriaRepository
}

// NewDeleteProductVariantUseCase crea una nueva instancia del use case
func NewDeleteProductVariantUseCase(
	productRepo port.ProductCriteriaRepository,
) *DeleteProductVariantUseCase {
	return &DeleteProductVariantUseCase{
		productRepo: productRepo,
	}
}

// Execute ejecuta la eliminación de una variante a través del agregado Product
func (uc *DeleteProductVariantUseCase) Execute(
	ctx context.Context,
	productID uuid.UUID,
	variantID uuid.UUID,
	tenantID string,
) error {
	// Obtener el producto con sus variantes
	product, err := uc.productRepo.FindByIDWithVariants(ctx, productID, tenantID)
	if err != nil {
		return fmt.Errorf("error obteniendo producto: %w", err)
	}

	if product == nil {
		return fmt.Errorf("producto no encontrado")
	}

	// Verificar que la variante existe en el agregado
	variant := product.GetVariantByID(variantID)
	if variant == nil {
		return fmt.Errorf("variante no encontrada")
	}

	// Eliminar la variante a través del agregado (maneja todas las reglas de negocio)
	err = product.DeleteVariant(variantID)
	if err != nil {
		return fmt.Errorf("error eliminando variante: %w", err)
	}

	// Guardar los cambios a través del repositorio
	err = uc.productRepo.UpdateVariant(ctx, variant)
	if err != nil {
		return fmt.Errorf("error guardando eliminación de variante: %w", err)
	}

	// Actualizar el producto para reflejar los cambios
	err = uc.productRepo.Update(ctx, product)
	if err != nil {
		return fmt.Errorf("error actualizando producto: %w", err)
	}

	return nil
}
