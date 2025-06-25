package service

import (
	"context"
	"errors"

	"pim/src/product/tenant/domain/entity"
	"pim/src/product/tenant/domain/port"
)

// ProductDomainService maneja las reglas de negocio del dominio Product
type ProductDomainService struct {
	productRepo port.ProductRepository
}

// NewProductDomainService crea una nueva instancia del servicio de dominio
func NewProductDomainService(productRepo port.ProductRepository) *ProductDomainService {
	return &ProductDomainService{
		productRepo: productRepo,
	}
}

// ValidateForCreation valida que un producto puede ser creado
func (s *ProductDomainService) ValidateForCreation(ctx context.Context, product *entity.Product) error {
	// Validar que no existe otro producto con el mismo nombre en el tenant
	exists, err := s.productRepo.ExistsByName(ctx, product.Name(), product.TenantID())
	if err != nil {
		return err
	}
	if exists {
		return errors.New("ya existe un producto con ese nombre")
	}

	// Validar que no existe otro producto con el mismo SKU en el tenant (si tiene SKU)
	if product.HasSKU() {
		exists, err := s.productRepo.ExistsBySKU(ctx, product.SKU().Value(), product.TenantID())
		if err != nil {
			return err
		}
		if exists {
			return errors.New("ya existe un producto con ese SKU")
		}
	}

	return nil
}

// ValidateForUpdate valida que un producto puede ser actualizado
func (s *ProductDomainService) ValidateForUpdate(ctx context.Context, product *entity.Product) error {
	// Validar que no existe otro producto con el mismo nombre en el tenant (excluyendo el actual)
	exists, err := s.productRepo.ExistsByNameExcludingID(ctx, product.Name(), product.TenantID(), product.ID())
	if err != nil {
		return err
	}
	if exists {
		return errors.New("ya existe un producto con ese nombre")
	}

	// Validar que no existe otro producto con el mismo SKU en el tenant (si tiene SKU, excluyendo el actual)
	if product.HasSKU() {
		exists, err := s.productRepo.ExistsBySKUExcludingID(ctx, product.SKU().Value(), product.TenantID(), product.ID())
		if err != nil {
			return err
		}
		if exists {
			return errors.New("ya existe un producto con ese SKU")
		}
	}

	return nil
}

// ValidateForDeletion valida que un producto puede ser eliminado
func (s *ProductDomainService) ValidateForDeletion(ctx context.Context, product *entity.Product) error {
	if !product.CanBeDeleted() {
		return errors.New("no se puede eliminar el producto")
	}

	// Aquí se pueden agregar más validaciones, como:
	// - Verificar que no tenga órdenes pendientes
	// - Verificar que no tenga stock en inventario
	// - Verificar que no esté siendo usado en promociones activas
	// Estas validaciones se harían mediante eventos o consultas a otros bounded contexts

	return nil
}

// CanChangeStatus verifica si se puede cambiar el estado del producto
func (s *ProductDomainService) CanChangeStatus(product *entity.Product, newStatus string) error {
	if product.IsDeleted() {
		return errors.New("no se puede cambiar el estado de un producto eliminado")
	}

	// Validaciones específicas por estado
	switch newStatus {
	case "active":
		// Un producto puede ser activado si no está eliminado
		return nil
	case "inactive":
		// Un producto puede ser desactivado si no está eliminado
		return nil
	case "discontinued":
		// Un producto puede ser descontinuado si no está eliminado
		return nil
	case "deleted":
		// Usar ValidateForDeletion para esta validación
		return s.ValidateForDeletion(context.Background(), product)
	default:
		return errors.New("estado de producto inválido")
	}
}
