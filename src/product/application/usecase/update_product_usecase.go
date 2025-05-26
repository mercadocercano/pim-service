package usecase

import (
	"context"

	"pim/src/product/application/mapper"
	"pim/src/product/application/request"
	"pim/src/product/application/response"
	"pim/src/product/domain/port"
	"pim/src/product/domain/service"
	"pim/src/product/domain/value_object"

	"github.com/google/uuid"
)

// UpdateProductUseCase maneja la actualización de productos
type UpdateProductUseCase struct {
	productRepo   port.ProductRepository
	domainService *service.ProductDomainService
	mapper        *mapper.ProductMapper
}

// NewUpdateProductUseCase crea una nueva instancia del caso de uso
func NewUpdateProductUseCase(
	productRepo port.ProductRepository,
	domainService *service.ProductDomainService,
	mapper *mapper.ProductMapper,
) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepo:   productRepo,
		domainService: domainService,
		mapper:        mapper,
	}
}

// Execute ejecuta el caso de uso de actualización de producto
func (uc *UpdateProductUseCase) Execute(
	ctx context.Context,
	productID uuid.UUID,
	req *request.UpdateProductRequest,
	tenantID string,
) (*response.ProductResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Buscar el producto existente
	product, err := uc.productRepo.FindByID(ctx, productID, tenantID)
	if err != nil {
		return nil, err
	}

	// Crear SKU si se proporciona
	var productSKU *value_object.ProductSKU
	if req.SKU != nil && *req.SKU != "" {
		sku, err := value_object.NewProductSKU(*req.SKU)
		if err != nil {
			return nil, err
		}
		productSKU = sku
	}

	// Crear referencia de categoría si se proporciona
	var categoryRef *value_object.CategoryReference
	if req.CategoryID != nil && *req.CategoryID != "" {
		// Aquí normalmente validaríamos que la categoría existe
		// Por ahora asumimos que el ID es válido y creamos una referencia básica
		ref, err := value_object.NewCategoryReference(*req.CategoryID, "Categoría") // Nombre temporal
		if err != nil {
			return nil, err
		}
		categoryRef = ref
	}

	// Crear referencia de marca si se proporciona
	var brandRef *value_object.BrandReference
	if req.BrandID != nil && *req.BrandID != "" {
		// Aquí normalmente validaríamos que la marca existe
		// Por ahora asumimos que el ID es válido y creamos una referencia básica
		ref, err := value_object.NewBrandReference(*req.BrandID, "Marca") // Nombre temporal
		if err != nil {
			return nil, err
		}
		brandRef = ref
	}

	// Actualizar el producto
	if err := product.Update(
		req.Name,
		req.Description,
		productSKU,
		categoryRef,
		brandRef,
	); err != nil {
		return nil, err
	}

	// Validar reglas de negocio
	if err := uc.domainService.ValidateForUpdate(ctx, product); err != nil {
		return nil, err
	}

	// Guardar los cambios
	if err := uc.productRepo.Update(ctx, product); err != nil {
		return nil, err
	}

	// Convertir a respuesta
	return uc.mapper.ToResponse(product), nil
}
