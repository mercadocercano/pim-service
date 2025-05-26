package usecase

import (
	"context"

	"pim/src/product/application/mapper"
	"pim/src/product/application/request"
	"pim/src/product/application/response"
	"pim/src/product/domain/entity"
	"pim/src/product/domain/port"
	"pim/src/product/domain/service"
	"pim/src/product/domain/value_object"
)

// CreateProductUseCase maneja la creación de productos
type CreateProductUseCase struct {
	productRepo   port.ProductRepository
	domainService *service.ProductDomainService
	mapper        *mapper.ProductMapper
}

// NewCreateProductUseCase crea una nueva instancia del caso de uso
func NewCreateProductUseCase(
	productRepo port.ProductRepository,
	domainService *service.ProductDomainService,
	mapper *mapper.ProductMapper,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepo:   productRepo,
		domainService: domainService,
		mapper:        mapper,
	}
}

// Execute ejecuta el caso de uso de creación de producto
func (uc *CreateProductUseCase) Execute(
	ctx context.Context,
	req *request.CreateProductRequest,
	tenantID string,
) (*response.ProductResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
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
		// En una implementación real, consultaríamos el servicio de categorías
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
		// En una implementación real, consultaríamos el servicio de marcas
		ref, err := value_object.NewBrandReference(*req.BrandID, "Marca") // Nombre temporal
		if err != nil {
			return nil, err
		}
		brandRef = ref
	}

	// Crear la entidad producto (el ID se genera automáticamente)
	product, err := entity.NewProduct(
		tenantID,
		req.Name,
		req.Description,
		productSKU,
		categoryRef,
		brandRef,
	)
	if err != nil {
		return nil, err
	}

	// Validar reglas de negocio
	if err := uc.domainService.ValidateForCreation(ctx, product); err != nil {
		return nil, err
	}

	// Guardar el producto
	if err := uc.productRepo.Save(ctx, product); err != nil {
		return nil, err
	}

	// Convertir a respuesta
	return uc.mapper.ToResponse(product), nil
}
