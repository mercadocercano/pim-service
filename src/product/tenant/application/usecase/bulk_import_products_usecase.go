package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/product/tenant/domain/port"
	"saas-mt-pim-service/src/product/tenant/domain/value_object"
)

// ProductImportData representa los datos de un producto para importación bulk
type ProductImportData struct {
	SKU         string            `json:"sku" binding:"required"`
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description"`
	Price       float64           `json:"price" binding:"required"`
	Category    string            `json:"category" binding:"required"`  // Nombre o slug de categoría
	Brand       string            `json:"brand"`                         // Nombre de marca
	Unit        string            `json:"unit" binding:"required"`      // unidad, kg, bolsa, metro, etc.
	Attributes  map[string]string `json:"attributes,omitempty"`         // Atributos adicionales
	Active      bool              `json:"active"`                        // Estado activo/inactivo
}

// BulkImportProductsRequest es la petición para importación bulk
type BulkImportProductsRequest struct {
	TenantID string              `json:"tenant_id" binding:"required"`
	Products []ProductImportData `json:"products" binding:"required,min=1"`
}

// BulkImportProductsResponse es la respuesta de la importación bulk
type BulkImportProductsResponse struct {
	Success           bool     `json:"success"`
	TotalProducts     int      `json:"total_products"`
	ProductsCreated   int      `json:"products_created"`
	ProductsFailed    int      `json:"products_failed"`
	Errors            []string `json:"errors,omitempty"`
	CreatedProductIDs []string `json:"created_product_ids,omitempty"`
}

// BulkImportProductsUseCase caso de uso para importación bulk de productos desde JSON
type BulkImportProductsUseCase struct {
	productRepo  port.ProductCriteriaRepository
	categoryRepo interface {
		FindBySlug(ctx context.Context, tenantID uuid.UUID, slug string) (*entity.Category, error)
		FindByName(ctx context.Context, tenantID uuid.UUID, name string) (*entity.Category, error)
	}
}

// Category representa una categoría (estructura mínima para evitar importaciones circulares)
type Category struct {
	ID   uuid.UUID
	Name string
	Slug string
}

// NewBulkImportProductsUseCase crea una nueva instancia del caso de uso
func NewBulkImportProductsUseCase(
	productRepo port.ProductCriteriaRepository,
	categoryRepo interface {
		FindBySlug(ctx context.Context, tenantID uuid.UUID, slug string) (*entity.Category, error)
		FindByName(ctx context.Context, tenantID uuid.UUID, name string) (*entity.Category, error)
	},
) *BulkImportProductsUseCase {
	return &BulkImportProductsUseCase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

// Execute ejecuta la importación bulk de productos
func (uc *BulkImportProductsUseCase) Execute(ctx context.Context, req BulkImportProductsRequest) (*BulkImportProductsResponse, error) {
	// Validar tenant UUID
	tenantUUID, err := uuid.Parse(req.TenantID)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant_id format: %w", err)
	}

	response := &BulkImportProductsResponse{
		Success:           true,
		TotalProducts:     len(req.Products),
		ProductsCreated:   0,
		ProductsFailed:    0,
		Errors:            make([]string, 0),
		CreatedProductIDs: make([]string, 0),
	}

	// Procesar cada producto
	for idx, productData := range req.Products {
		productID, err := uc.createProduct(ctx, tenantUUID, productData)
		if err != nil {
			response.ProductsFailed++
			response.Errors = append(response.Errors, fmt.Sprintf(
				"Producto %d (SKU: %s): %v",
				idx+1,
				productData.SKU,
				err,
			))
			continue
		}

		response.ProductsCreated++
		response.CreatedProductIDs = append(response.CreatedProductIDs, productID)
	}

	// Si todos fallaron, marcar como no exitoso
	if response.ProductsCreated == 0 {
		response.Success = false
	}

	return response, nil
}

// createProduct crea un producto individual
func (uc *BulkImportProductsUseCase) createProduct(
	ctx context.Context,
	tenantID uuid.UUID,
	data ProductImportData,
) (string, error) {
	// 1. Buscar categoría por nombre o slug
	var category *entity.Category
	var err error

	// Intentar primero por slug (más preciso)
	category, err = uc.categoryRepo.FindBySlug(ctx, tenantID, data.Category)
	if err != nil {
		// Si no encuentra por slug, intentar por nombre
		category, err = uc.categoryRepo.FindByName(ctx, tenantID, data.Category)
		if err != nil {
			return "", fmt.Errorf("categoría no encontrada: %s", data.Category)
		}
	}

	// 2. Crear referencia de categoría
	categoryRef := value_object.NewCategoryReference(
		category.ID.String(),
		category.Name,
		"",  // path vacío por ahora
	)

	// 3. Crear referencia de marca (opcional)
	var brandRef *value_object.BrandReference
	if data.Brand != "" {
		// Por ahora crear referencia con ID vacío, solo el nombre
		brandRef = value_object.NewBrandReference("", data.Brand, "")
	}

	// 4. Determinar status
	status := value_object.StatusActive
	if !data.Active {
		status = value_object.StatusInactive
	}

	// 5. Crear producto
	now := time.Now()
	product := &entity.Product{
		ID:          uuid.New(),
		TenantID:    tenantID,
		SKU:         data.SKU,
		Name:        data.Name,
		Description: &data.Description,
		Price:       data.Price,
		Category:    *categoryRef,
		Brand:       brandRef,
		Metadata:    data.Attributes,
		Status:      status,
		IsActive:    data.Active,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// 6. Guardar producto
	if err := uc.productRepo.Save(ctx, product); err != nil {
		return "", fmt.Errorf("error al guardar producto: %w", err)
	}

	return product.ID.String(), nil
}

