package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/batch/application/request"
	"saas-mt-pim-service/src/batch/application/response"
	batchPort "saas-mt-pim-service/src/batch/domain/port"
	"saas-mt-pim-service/src/brand/domain/entity"
	brandPort "saas-mt-pim-service/src/brand/domain/port"
	categoryEntity "saas-mt-pim-service/src/category/domain/entity"
	categoryPort "saas-mt-pim-service/src/category/domain/port"
	productEntity "saas-mt-pim-service/src/product/tenant/domain/entity"
	productPort "saas-mt-pim-service/src/product/tenant/domain/port"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"

	"github.com/google/uuid"
)

// BatchCreateUseCase maneja la creación masiva de entidades
type BatchCreateUseCase struct {
	txBeginner          batchPort.TxBeginner
	categoryRepo        categoryPort.CategoryRepository
	brandRepo           brandPort.BrandRepository
	productRepo         productPort.ProductCriteriaRepository
	categoryMappingRepo categoryPort.TenantCategoryMappingRepository
}

// NewBatchCreateUseCase crea una nueva instancia del caso de uso
func NewBatchCreateUseCase(
	txBeginner batchPort.TxBeginner,
	categoryRepo categoryPort.CategoryRepository,
	brandRepo brandPort.BrandRepository,
	productRepo productPort.ProductCriteriaRepository,
	categoryMappingRepo categoryPort.TenantCategoryMappingRepository,
) *BatchCreateUseCase {
	return &BatchCreateUseCase{
		txBeginner:          txBeginner,
		categoryRepo:        categoryRepo,
		brandRepo:           brandRepo,
		productRepo:         productRepo,
		categoryMappingRepo: categoryMappingRepo,
	}
}

// Execute ejecuta la creación en batch con transacción
func (uc *BatchCreateUseCase) Execute(ctx context.Context, req *request.BatchCreateRequest, tenantID string) (*response.BatchCreateResponse, error) {
	result := &response.BatchCreateResponse{
		Created: response.BatchCreatedItems{
			Categories: []string{},
			Brands:     []string{},
			Products:   []string{},
		},
		Errors: []response.BatchError{},
	}

	// Iniciar transacción
	tx, err := uc.txBeginner.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error al iniciar transacción: %w", err)
	}
	defer tx.Rollback()

	// Mapas para guardar IDs creados (para resolver referencias)
	categoryIDMap := make(map[string]string) // nombre -> ID
	brandIDMap := make(map[string]string)    // nombre -> ID

	// 1. Crear categorías (respetando jerarquía)
	for _, catItem := range req.Categories {
		// Usar el constructor NewCategory
		category, err := categoryEntity.NewCategory(
			tenantID,
			catItem.Name,
			catItem.Description,
			catItem.ParentID,
		)
		if err != nil {
			result.Errors = append(result.Errors, response.BatchError{
				Type:  "category",
				Name:  catItem.Name,
				Error: fmt.Sprintf("error al crear categoría: %v", err),
			})
			continue
		}

		// TODO: Usar repositorio con transacción
		if err := uc.categoryRepo.Create(ctx, category); err != nil {
			result.Errors = append(result.Errors, response.BatchError{
				Type:  "category",
				Name:  catItem.Name,
				Error: err.Error(),
			})
			continue
		}

		categoryIDMap[catItem.Name] = category.ID
		result.Created.Categories = append(result.Created.Categories, category.ID)

		// Crear mapeo con marketplace si se especifica
		if catItem.Mapping != nil {
			// Usar el constructor NewTenantCategoryMapping
			var customName *string
			if catItem.Mapping.CustomName != "" {
				customName = &catItem.Mapping.CustomName
			}
			
			mapping, err := categoryEntity.NewTenantCategoryMapping(
				tenantID,
				category.ID,
				catItem.Mapping.MarketplaceCategoryID,
				customName,
			)
			if err != nil {
				result.Errors = append(result.Errors, response.BatchError{
					Type:  "category_mapping",
					Name:  catItem.Name,
					Error: fmt.Sprintf("error al crear mapeo: %v", err),
				})
				continue
			}
			
			// Asignar ID al mapping
			mapping.ID = uuid.New().String()

			if err := uc.categoryMappingRepo.Save(ctx, mapping); err != nil {
				// Log error pero no fallar la categoría
				result.Errors = append(result.Errors, response.BatchError{
					Type:  "category_mapping",
					Name:  catItem.Name,
					Error: fmt.Sprintf("error al crear mapeo: %v", err),
				})
			}
		}
	}

	// 2. Crear marcas
	for _, brandItem := range req.Brands {
		// Preparar logoURL y website como punteros si no están vacíos
		var logoURL, website *string
		if brandItem.LogoURL != "" {
			logoURL = &brandItem.LogoURL
		}
		if brandItem.Website != "" {
			website = &brandItem.Website
		}

		// Usar el constructor NewBrand
		brand, err := entity.NewBrand(
			tenantID,
			brandItem.Name,
			brandItem.Description,
			logoURL,
			website,
		)
		if err != nil {
			result.Errors = append(result.Errors, response.BatchError{
				Type:  "brand",
				Name:  brandItem.Name,
				Error: fmt.Sprintf("error al crear marca: %v", err),
			})
			continue
		}

		if err := uc.brandRepo.Create(ctx, brand); err != nil {
			result.Errors = append(result.Errors, response.BatchError{
				Type:  "brand",
				Name:  brandItem.Name,
				Error: err.Error(),
			})
			continue
		}

		brandIDMap[brandItem.Name] = brand.ID
		result.Created.Brands = append(result.Created.Brands, brand.ID)
	}

	// 3. Crear productos
	for _, prodItem := range req.Products {
		// Preparar description como puntero si no está vacío
		var description *string
		if prodItem.Description != "" {
			description = &prodItem.Description
		}

		// Preparar SKU si se proporciona
		var sku *value_object.ProductSKU
		if prodItem.SKU != "" {
			productSKU, err := value_object.NewProductSKU(prodItem.SKU)
			if err != nil {
				result.Errors = append(result.Errors, response.BatchError{
					Type:  "product",
					Name:  prodItem.Name,
					Error: fmt.Sprintf("SKU inválido: %v", err),
				})
				continue
			}
			sku = productSKU
		}

		// Preparar referencia de categoría si se proporciona
		var categoryRef *value_object.CategoryReference
		if prodItem.CategoryID != "" {
			// Obtener el nombre de la categoría del mapa o usar vacío
			categoryName := ""
			for catName, catID := range categoryIDMap {
				if catID == prodItem.CategoryID {
					categoryName = catName
					break
				}
			}
			ref, err := value_object.NewCategoryReference(prodItem.CategoryID, categoryName)
			if err != nil {
				result.Errors = append(result.Errors, response.BatchError{
					Type:  "product",
					Name:  prodItem.Name,
					Error: fmt.Sprintf("referencia de categoría inválida: %v", err),
				})
				continue
			}
			categoryRef = ref
		}

		// Preparar referencia de marca si se proporciona
		var brandRef *value_object.BrandReference
		if prodItem.BrandID != "" {
			// Obtener el nombre de la marca del mapa o usar vacío
			brandName := ""
			for bName, bID := range brandIDMap {
				if bID == prodItem.BrandID {
					brandName = bName
					break
				}
			}
			ref, err := value_object.NewBrandReference(prodItem.BrandID, brandName)
			if err != nil {
				result.Errors = append(result.Errors, response.BatchError{
					Type:  "product",
					Name:  prodItem.Name,
					Error: fmt.Sprintf("referencia de marca inválida: %v", err),
				})
				continue
			}
			brandRef = ref
		}

		// Crear el producto usando el constructor
		product, err := productEntity.NewProduct(
			tenantID,
			prodItem.Name,
			description,
			sku,
			categoryRef,
			brandRef,
		)
		if err != nil {
			result.Errors = append(result.Errors, response.BatchError{
				Type:  "product",
				Name:  prodItem.Name,
				Error: fmt.Sprintf("error al crear producto: %v", err),
			})
			continue
		}

		if err := uc.productRepo.Save(ctx, product); err != nil {
			result.Errors = append(result.Errors, response.BatchError{
				Type:  "product",
				Name:  prodItem.Name,
				Error: err.Error(),
			})
			continue
		}

		result.Created.Products = append(result.Created.Products, product.IDString())

		// TODO: Crear variante por defecto si se especifica precio y stock
		// NOTA: El constructor NewProduct ya crea una variante por defecto automáticamente
	}

	// Si hay errores críticos, hacer rollback
	if len(result.Errors) > 0 && len(result.Created.Categories) == 0 && 
	   len(result.Created.Brands) == 0 && len(result.Created.Products) == 0 {
		return result, fmt.Errorf("ninguna entidad pudo ser creada")
	}

	// Commit de la transacción
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error al confirmar transacción: %w", err)
	}

	return result, nil
}