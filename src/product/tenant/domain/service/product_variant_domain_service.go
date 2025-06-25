package service

import (
	"context"
	"fmt"

	"pim/src/product/tenant/domain/entity"
	"pim/src/product/tenant/domain/port"

	"github.com/google/uuid"
)

// ProductVariantDomainService maneja las reglas de negocio para variantes de productos
type ProductVariantDomainService struct {
	productRepo port.ProductRepository
}

// NewProductVariantDomainService crea una nueva instancia del servicio de dominio
func NewProductVariantDomainService(
	productRepo port.ProductRepository,
) *ProductVariantDomainService {
	return &ProductVariantDomainService{
		productRepo: productRepo,
	}
}

// ValidateVariantCreation valida que se puede crear una nueva variante
func (s *ProductVariantDomainService) ValidateVariantCreation(
	ctx context.Context,
	variant *entity.ProductVariant,
) error {
	// Verificar que el producto existe
	product, err := s.productRepo.FindByID(ctx, variant.ProductID(), variant.TenantID())
	if err != nil {
		return fmt.Errorf("error verificando producto: %w", err)
	}
	if product == nil {
		return fmt.Errorf("el producto no existe")
	}

	// Verificar que el tenant coincide
	if product.TenantID() != variant.TenantID() {
		return fmt.Errorf("el tenant de la variante no coincide con el del producto")
	}

	// Verificar que el nombre de la variante es único para el producto
	exists, err := s.productRepo.VariantExistsByName(
		ctx,
		variant.Name(),
		variant.ProductID(),
		variant.TenantID(),
	)
	if err != nil {
		return fmt.Errorf("error verificando nombre de variante: %w", err)
	}
	if exists {
		return fmt.Errorf("ya existe una variante con el nombre '%s' para este producto", variant.Name())
	}

	// Verificar que el SKU es único (si se proporciona)
	if variant.HasSKU() {
		exists, err := s.productRepo.VariantExistsBySKU(
			ctx,
			variant.SKU().Value(),
			variant.TenantID(),
		)
		if err != nil {
			return fmt.Errorf("error verificando SKU de variante: %w", err)
		}
		if exists {
			return fmt.Errorf("ya existe una variante con el SKU '%s'", variant.SKU().Value())
		}
	}

	return nil
}

// ValidateVariantUpdate valida que se puede actualizar una variante
func (s *ProductVariantDomainService) ValidateVariantUpdate(
	ctx context.Context,
	variant *entity.ProductVariant,
) error {
	// Verificar que el nombre de la variante es único para el producto (excluyendo la variante actual)
	exists, err := s.productRepo.VariantExistsByNameExcludingID(
		ctx,
		variant.Name(),
		variant.ProductID(),
		variant.TenantID(),
		variant.ID(),
	)
	if err != nil {
		return fmt.Errorf("error verificando nombre de variante: %w", err)
	}
	if exists {
		return fmt.Errorf("ya existe otra variante con el nombre '%s' para este producto", variant.Name())
	}

	// Verificar que el SKU es único (si se proporciona, excluyendo la variante actual)
	if variant.HasSKU() {
		exists, err := s.productRepo.VariantExistsBySKUExcludingID(
			ctx,
			variant.SKU().Value(),
			variant.TenantID(),
			variant.ID(),
		)
		if err != nil {
			return fmt.Errorf("error verificando SKU de variante: %w", err)
		}
		if exists {
			return fmt.Errorf("ya existe otra variante con el SKU '%s'", variant.SKU().Value())
		}
	}

	return nil
}

// ValidateVariantDeletion valida que se puede eliminar una variante
func (s *ProductVariantDomainService) ValidateVariantDeletion(
	ctx context.Context,
	productID uuid.UUID,
	variantID uuid.UUID,
	tenantID string,
) error {
	// Obtener el producto con sus variantes
	product, err := s.productRepo.FindByIDWithVariants(ctx, productID, tenantID)
	if err != nil {
		return fmt.Errorf("error obteniendo producto: %w", err)
	}
	if product == nil {
		return fmt.Errorf("el producto no existe")
	}

	// Verificar que la variante existe
	variant := product.GetVariantByID(variantID)
	if variant == nil {
		return fmt.Errorf("la variante no existe")
	}

	// Verificar si es la única variante activa del producto
	activeVariants := product.GetVariants()
	if len(activeVariants) <= 1 && variant.IsActive() {
		return fmt.Errorf("no se puede eliminar la única variante activa del producto")
	}

	return nil
}
