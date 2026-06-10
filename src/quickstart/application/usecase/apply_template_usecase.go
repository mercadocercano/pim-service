package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"saas-mt-pim-service/src/quickstart/domain/port"
	sharedport "github.com/mercadocercano/go-shared/domain/port"
)

// ApplyTemplateRequest es la petición para aplicar un template
type ApplyTemplateRequest struct {
	TemplateID string `json:"template_id" binding:"required"`
	TenantID   string `json:"tenant_id" binding:"required"`
}

// ApplyTemplateResponse es la respuesta al aplicar un template
type ApplyTemplateResponse struct {
	Success            bool     `json:"success"`
	TemplateID         string   `json:"template_id"`
	TenantID           string   `json:"tenant_id"`
	CategoriesCreated  int      `json:"categories_created"`
	BrandsCreated      int      `json:"brands_created"`
	ProductsCreated    int      `json:"products_created"`
	VariantsCreated    int      `json:"variants_created"`
	AttributesCreated  int      `json:"attributes_created"`
	CategoryAttributes int      `json:"category_attributes_created"`
	Message            string   `json:"message"`
	CreatedCategories  []string `json:"created_categories,omitempty"`
	CreatedBrands      []string `json:"created_brands,omitempty"`
	CreatedProducts    []string `json:"created_products,omitempty"`
}

// ApplyTemplateUseCase aplica un template de quickstart a un tenant
type ApplyTemplateUseCase struct {
	db   *sql.DB
	repo port.ApplyTemplateRepository
}

const (
	legacyQuickstartTemplateID = "ferreteria-corralon"
	legacyQuickstartParentID   = "f1e8f2a3-4b6c-4d5e-8f9a-1b2c3d4e5f00"
)

// NewApplyTemplateUseCase crea una nueva instancia del caso de uso
func NewApplyTemplateUseCase(db *sql.DB, repo port.ApplyTemplateRepository) *ApplyTemplateUseCase {
	return &ApplyTemplateUseCase{
		db:   db,
		repo: repo,
	}
}

// Execute ejecuta el caso de uso para aplicar un template
func (uc *ApplyTemplateUseCase) Execute(ctx context.Context, req ApplyTemplateRequest) (*ApplyTemplateResponse, error) {
	// Validar tenant UUID
	tenantUUID, err := uuid.Parse(req.TenantID)
	if err != nil {
		return nil, fmt.Errorf("invalid tenant_id format: %w", err)
	}

	// Intentar cargar datos curados del template (brands, products, attributes)
	fullData, _ := uc.repo.LoadFullTemplateData(ctx, req.TemplateID)
	useCuratedFlow := fullData != nil && len(fullData.Products) > 0

	// Cargar categorías (siempre necesario)
	templateCategories, marketplaceCategoryIDs, categorySlugByMarketplaceID, err := uc.repo.LoadTemplateCategories(ctx, req.TemplateID)
	if err != nil {
		return nil, err
	}

	// Si hay datos curados, usar sus categorías (pueden ser más completas)
	if useCuratedFlow && len(fullData.Categories) > 0 {
		templateCategories = fullData.Categories
	}

	useLegacy := len(templateCategories) == 0 && req.TemplateID == legacyQuickstartTemplateID
	if !useLegacy && len(templateCategories) == 0 {
		return nil, fmt.Errorf("template not found: %s", req.TemplateID)
	}

	tx, err := uc.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	response := &ApplyTemplateResponse{
		Success:    true,
		TemplateID: req.TemplateID,
		TenantID:   req.TenantID,
	}

	// PASO 1: Crear categorías del tenant
	var categoriesCreated int
	var createdCategories []port.CreatedCategory
	if useLegacy {
		categoriesCreated, createdCategories, err = uc.repo.CreateTenantCategoriesLegacy(ctx, tx, tenantUUID, legacyQuickstartParentID)
	} else {
		categoriesCreated, createdCategories, err = uc.repo.CreateTenantCategoriesFromTemplate(ctx, tx, tenantUUID, templateCategories)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create categories: %w", err)
	}
	response.CategoriesCreated = categoriesCreated
	response.CreatedCategories = createdCategoryList(createdCategories).Names()

	if useCuratedFlow {
		// FLUJO CURADO: usar datos del template JSONB

		// PASO 2: Crear marcas del template + marcas extraídas de productos
		allBrands := extractBrandsFromProducts(fullData.Brands, fullData.Products)
		brandsCreated, createdBrands, err := uc.repo.CreateTenantBrandsFromTemplate(ctx, tx, tenantUUID, allBrands)
		if err != nil {
			return nil, fmt.Errorf("failed to create brands from template: %w", err)
		}
		response.BrandsCreated = brandsCreated
		response.CreatedBrands = createdBrands

		// PASO 3: Crear productos del template
		productsCreated, variantsCreated, createdProducts, err := uc.repo.CreateTenantProductsFromTemplate(
			ctx, tx, tenantUUID, fullData.Products, createdCategories, createdBrands,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create products from template: %w", err)
		}
		response.ProductsCreated = productsCreated
		response.VariantsCreated = variantsCreated
		response.CreatedProducts = createdProducts

		// PASO 4: Crear atributos y vincularlos a categorías
		attrsCreated, linksCreated, err := uc.repo.CreateTenantAttributesFromTemplate(
			ctx, tx, tenantUUID, fullData.Attributes, createdCategories,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create attributes from template: %w", err)
		}
		response.AttributesCreated = attrsCreated
		response.CategoryAttributes = linksCreated

	} else {
		// FLUJO LEGACY: buscar en global_products (retrocompatibilidad)

		if useLegacy {
			marketplaceCategoryIDs, err = uc.repo.GetMarketplaceCategoryIDsBySlug(ctx, tx, createdCategoryList(createdCategories).Slugs())
			if err != nil {
				return nil, fmt.Errorf("failed to load marketplace categories: %w", err)
			}
		}

		useMarketplaceCategoryID := uc.repo.GlobalProductsHasMarketplaceCategoryID(ctx, tx)

		categorySlugs := createdCategoryList(createdCategories).Slugs()
		if !useLegacy {
			categorySlugs = templateCategoriesSlugs(templateCategories)
		}

		brandsCreated, createdBrands, err := uc.repo.CreateTenantBrandsFromGlobalProducts(
			ctx, tx, tenantUUID, marketplaceCategoryIDs, categorySlugs, useMarketplaceCategoryID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create brands: %w", err)
		}
		response.BrandsCreated = brandsCreated
		response.CreatedBrands = createdBrands

		productsCreated, variantsCreated, createdProducts, err := uc.createTenantProductFromGlobalProducts(
			ctx, tx, tenantUUID, marketplaceCategoryIDs, categorySlugs,
			useMarketplaceCategoryID, createdCategories, categorySlugByMarketplaceID, createdBrands,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create product: %w", err)
		}
		response.ProductsCreated = productsCreated
		response.VariantsCreated = variantsCreated
		response.CreatedProducts = createdProducts

		response.AttributesCreated = 0
		response.CategoryAttributes = 0
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	response.Message = fmt.Sprintf(
		"Template aplicado exitosamente: %d categorías, %d marcas, %d productos, %d variantes, %d atributos",
		response.CategoriesCreated,
		response.BrandsCreated,
		response.ProductsCreated,
		response.VariantsCreated,
		response.AttributesCreated,
	)

	return response, nil
}

type createdCategoryList []port.CreatedCategory

func (list createdCategoryList) Names() []string {
	names := make([]string, 0, len(list))
	for _, item := range list {
		names = append(names, item.Name)
	}
	return names
}

func (list createdCategoryList) Slugs() []string {
	slugs := make([]string, 0, len(list))
	for _, item := range list {
		slugs = append(slugs, item.Slug)
	}
	return slugs
}

func (list createdCategoryList) BySlug() map[string]port.CreatedCategory {
	result := make(map[string]port.CreatedCategory, len(list))
	for _, item := range list {
		result[item.Slug] = item
	}
	return result
}

func (list createdCategoryList) ByMarketplaceID() map[string]port.CreatedCategory {
	result := make(map[string]port.CreatedCategory, len(list))
	for _, item := range list {
		if item.MarketplaceID != "" {
			result[item.MarketplaceID] = item
		}
	}
	return result
}

// extractBrandsFromProducts combina las marcas explícitas del template con las
// marcas que aparecen en los productos pero no en el array de brands.
func extractBrandsFromProducts(templateBrands []port.TemplateBrand, products []port.TemplateProduct) []port.TemplateBrand {
	seen := make(map[string]bool, len(templateBrands))
	for _, b := range templateBrands {
		seen[strings.TrimSpace(b.Name)] = true
	}

	result := make([]port.TemplateBrand, len(templateBrands))
	copy(result, templateBrands)

	for _, p := range products {
		brand := strings.TrimSpace(p.Brand)
		if brand == "" || seen[brand] {
			continue
		}
		seen[brand] = true
		result = append(result, port.TemplateBrand{Name: brand})
	}
	return result
}

func templateCategoriesSlugs(categories []port.TemplateCategory) []string {
	slugs := make([]string, 0, len(categories))
	seen := make(map[string]struct{}, len(categories))
	for _, category := range categories {
		slug := strings.TrimSpace(category.Slug)
		if slug == "" {
			continue
		}
		if _, ok := seen[slug]; ok {
			continue
		}
		seen[slug] = struct{}{}
		slugs = append(slugs, slug)
	}
	return slugs
}

func (uc *ApplyTemplateUseCase) createTenantProductFromGlobalProducts(
	ctx context.Context,
	exec sharedport.Executor,
	tenantID uuid.UUID,
	marketplaceCategoryIDs []string,
	categorySlugs []string,
	useMarketplaceCategoryID bool,
	createdCategories []port.CreatedCategory,
	categorySlugByMarketplaceID map[string]string,
	brandNames []string,
) (int, int, []string, error) {
	if useMarketplaceCategoryID && len(marketplaceCategoryIDs) == 0 {
		return 0, 0, nil, nil
	}
	if !useMarketplaceCategoryID && len(categorySlugs) == 0 {
		return 0, 0, nil, nil
	}

	var candidate port.GlobalProductCandidate
	var err error

	if len(brandNames) > 0 {
		candidate, err = uc.repo.FindGlobalProduct(ctx, exec, marketplaceCategoryIDs, categorySlugs, useMarketplaceCategoryID, brandNames)
	} else {
		candidate, err = uc.repo.FindGlobalProduct(ctx, exec, marketplaceCategoryIDs, categorySlugs, useMarketplaceCategoryID, nil)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, nil, nil
		}
		return 0, 0, nil, err
	}

	brandID, brandName, err := uc.repo.EnsureTenantBrand(ctx, exec, tenantID, candidate.Brand)
	if err != nil {
		return 0, 0, nil, err
	}

	categoryID, categoryName, err := uc.repo.ResolveTenantCategory(
		ctx,
		exec,
		tenantID,
		createdCategoryList(createdCategories).ByMarketplaceID(),
		createdCategoryList(createdCategories).BySlug(),
		categorySlugByMarketplaceID,
		candidate.MarketplaceCategoryID,
	)
	if err != nil {
		return 0, 0, nil, err
	}

	productID, productName, productSKU, created, err := uc.repo.EnsureTenantProduct(ctx, exec, tenantID, candidate, categoryID, categoryName, brandID, brandName)
	if err != nil {
		return 0, 0, nil, err
	}

	variantsCreated := 0
	if created {
		variantsCreated, err = uc.repo.EnsureDefaultVariant(ctx, exec, tenantID, productID, productName, productSKU)
		if err != nil {
			return 0, 0, nil, err
		}
	}

	productsCreated := 0
	if created {
		productsCreated = 1
	}

	createdProducts := []string{}
	if created {
		createdProducts = append(createdProducts, productName)
	}

	return productsCreated, variantsCreated, createdProducts, nil
}
