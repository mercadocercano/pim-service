package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/product/tenant/application/mapper"
	"saas-mt-pim-service/src/product/tenant/application/request"
	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/product/tenant/domain/port"
	"saas-mt-pim-service/src/product/tenant/domain/service"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"
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
	productMapper *mapper.ProductMapper,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepo:   productRepo,
		domainService: domainService,
		mapper:        productMapper,
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

	// Configurar variantes: si se proporcionan en variants[], usarlas; si no, aplicar price/stock top-level
	if len(req.Variants) > 0 {
		if err := uc.applyVariantsFromRequest(product, req.Variants); err != nil {
			return nil, err
		}
	} else if req.Price != nil || req.Stock != nil {
		defaultVariant := product.GetDefaultVariant()
		if defaultVariant != nil {
			if req.Price != nil {
				defaultVariant.UpdatePrice(*req.Price)
			}
			if req.Stock != nil {
				defaultVariant.UpdateStock(*req.Stock)
			}
		}
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

// applyVariantsFromRequest configura las variantes del producto a partir del request.
// La primera variante del array sobreescribe la variante default auto-generada.
// Las variantes adicionales se agregan al producto.
func (uc *CreateProductUseCase) applyVariantsFromRequest(
	product *entity.Product,
	variants []request.CreateVariantInProductRequest,
) error {
	// Configurar la variante default con los datos de la primera variante del request
	defaultVariant := product.GetDefaultVariant()
	if defaultVariant != nil {
		first := variants[0]
		if first.Name != nil && *first.Name != "" {
			if err := defaultVariant.UpdateName(*first.Name); err != nil {
				return err
			}
		}
		if first.SKU != nil && *first.SKU != "" {
			sku, err := value_object.NewProductSKU(*first.SKU)
			if err != nil {
				return err
			}
			defaultVariant.UpdateSKU(sku)
		}
		if first.Price != nil {
			defaultVariant.UpdatePrice(*first.Price)
		}
		if first.Barcode != nil && *first.Barcode != "" {
			attrs, err := buildBarcodeAttributeCollection(*first.Barcode)
			if err != nil {
				return err
			}
			defaultVariant.UpdateAttributes(attrs)
		}
	}

	// Agregar variantes adicionales (índice 1 en adelante)
	for i, variantReq := range variants[1:] {
		variantName := fmt.Sprintf("%s - Variante %d", product.Name(), i+2)
		if variantReq.Name != nil && *variantReq.Name != "" {
			variantName = *variantReq.Name
		}

		var variantSKU *value_object.ProductSKU
		if variantReq.SKU != nil && *variantReq.SKU != "" {
			sku, err := value_object.NewProductSKU(*variantReq.SKU)
			if err != nil {
				return err
			}
			variantSKU = sku
		}

		var attrCollection *value_object.VariantAttributeCollection
		if variantReq.Barcode != nil && *variantReq.Barcode != "" {
			var err error
			attrCollection, err = buildBarcodeAttributeCollection(*variantReq.Barcode)
			if err != nil {
				return err
			}
		} else {
			attrCollection, _ = value_object.NewVariantAttributeCollection([]*value_object.VariantAttribute{})
		}

		newVariant, err := product.AddVariant(variantName, variantSKU, false, i+2, attrCollection)
		if err != nil {
			return err
		}
		if variantReq.Price != nil {
			newVariant.UpdatePrice(*variantReq.Price)
		}
	}

	return nil
}

// buildBarcodeAttributeCollection crea una colección de atributos con el barcode dado
func buildBarcodeAttributeCollection(barcode string) (*value_object.VariantAttributeCollection, error) {
	barcodeAttr, err := value_object.NewVariantAttribute("barcode", barcode)
	if err != nil {
		return nil, err
	}
	return value_object.NewVariantAttributeCollection([]*value_object.VariantAttribute{barcodeAttr})
}
