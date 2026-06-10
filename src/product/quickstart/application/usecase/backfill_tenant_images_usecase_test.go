package usecase_test

import (
	"context"
	"testing"
	"time"

	cr "github.com/mercadocercano/criteria"
	globalEntity "saas-mt-pim-service/src/product/global_catalog/domain/entity"
	globalVO "saas-mt-pim-service/src/product/global_catalog/domain/value_object"
	"saas-mt-pim-service/src/product/quickstart/application/usecase"
	tenantEntity "saas-mt-pim-service/src/product/tenant/domain/entity"
	tenantPort "saas-mt-pim-service/src/product/tenant/domain/port"
	tenantVO "saas-mt-pim-service/src/product/tenant/domain/value_object"

	"github.com/google/uuid"
)

// ---- Mocks inline ----

type mockGlobalRepo struct {
	findByNameAndBrandFn func(ctx context.Context, name, brand string) (*globalEntity.GlobalProduct, error)
}

func (m *mockGlobalRepo) FindByNameAndBrand(ctx context.Context, name, brand string) (*globalEntity.GlobalProduct, error) {
	return m.findByNameAndBrandFn(ctx, name, brand)
}

// Métodos no usados en estos tests — implementaciones vacías para satisfacer la interface.
func (m *mockGlobalRepo) Save(gp *globalEntity.GlobalProduct) (*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) Update(gp *globalEntity.GlobalProduct) (*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindByID(id string) (*globalEntity.GlobalProduct, error)   { return nil, nil }
func (m *mockGlobalRepo) Delete(id string) error                                    { return nil }
func (m *mockGlobalRepo) FindByEAN(ean string) (*globalEntity.GlobalProduct, error) { return nil, nil }
func (m *mockGlobalRepo) FindActiveByEAN(ean string) (*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindByBusinessType(bt string, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindBySource(source string, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindByQualityScoreRange(min, max, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) SearchByName(name string, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) SearchByBrand(brand string, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) SearchByCategory(cat string, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) SearchByTags(tags []string, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindAll(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindActive(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindVerified(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindArgentineProducts(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindHighQualityProducts(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindNeedingUpdate(maxAgeHours, limit int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindNeedingEnrichment(bt *string, limit, offset int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) CountNeedingEnrichment(bt *string) (int, error) { return 0, nil }
func (m *mockGlobalRepo) FindDistinctBrandsByBusinessType(bt string) ([]string, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindDistinctCategoriesByBusinessType(bt string) ([]string, error) {
	return nil, nil
}
func (m *mockGlobalRepo) FindDistinctBusinessTypes() ([]string, error)  { return nil, nil }
func (m *mockGlobalRepo) CountTotal() (int, error)                      { return 0, nil }
func (m *mockGlobalRepo) CountBySource(source string) (int, error)      { return 0, nil }
func (m *mockGlobalRepo) CountByQualityScore(minScore int) (int, error) { return 0, nil }
func (m *mockGlobalRepo) CountArgentineProducts() (int, error)          { return 0, nil }
func (m *mockGlobalRepo) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalRepo) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	return 0, nil
}
func (m *mockGlobalRepo) FindByIDs(ctx context.Context, ids []string) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}

type mockTenantRepo struct {
	findWithoutImageFn    func(ctx context.Context, tenantID string) ([]*tenantEntity.Product, error)
	updateImageURLFn      func(ctx context.Context, tenantID, productID, imageURL string) error
	findDistinctTenantsFn func(ctx context.Context) ([]string, error)
	updatedProductIDs     []string
}

func (m *mockTenantRepo) FindWithoutImage(ctx context.Context, tenantID string) ([]*tenantEntity.Product, error) {
	return m.findWithoutImageFn(ctx, tenantID)
}

func (m *mockTenantRepo) UpdateImageURL(ctx context.Context, tenantID, productID, imageURL string) error {
	m.updatedProductIDs = append(m.updatedProductIDs, productID)
	if m.updateImageURLFn != nil {
		return m.updateImageURLFn(ctx, tenantID, productID, imageURL)
	}
	return nil
}

func (m *mockTenantRepo) FindDistinctTenantIDs(ctx context.Context) ([]string, error) {
	if m.findDistinctTenantsFn != nil {
		return m.findDistinctTenantsFn(ctx)
	}
	return nil, nil
}

// Métodos no usados — implementaciones vacías para satisfacer ProductRepository.
func (m *mockTenantRepo) Save(ctx context.Context, p *tenantEntity.Product) error { return nil }
func (m *mockTenantRepo) FindByID(ctx context.Context, id uuid.UUID, tenantID string) (*tenantEntity.Product, error) {
	return nil, nil
}
func (m *mockTenantRepo) FindByIDWithVariants(ctx context.Context, id uuid.UUID, tenantID string) (*tenantEntity.Product, error) {
	return nil, nil
}
func (m *mockTenantRepo) FindBySKU(ctx context.Context, sku, tenantID string) (*tenantEntity.Product, error) {
	return nil, nil
}
func (m *mockTenantRepo) Update(ctx context.Context, p *tenantEntity.Product) error { return nil }
func (m *mockTenantRepo) Delete(ctx context.Context, id uuid.UUID, tenantID string) error {
	return nil
}
func (m *mockTenantRepo) SaveVariant(ctx context.Context, productID uuid.UUID, v *tenantEntity.ProductVariant) error {
	return nil
}
func (m *mockTenantRepo) UpdateVariant(ctx context.Context, v *tenantEntity.ProductVariant) error {
	return nil
}
func (m *mockTenantRepo) DeleteVariant(ctx context.Context, variantID uuid.UUID) error { return nil }
func (m *mockTenantRepo) LoadVariantsForProduct(ctx context.Context, productID uuid.UUID) ([]*tenantEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockTenantRepo) ExistsByID(ctx context.Context, id uuid.UUID, tenantID string) (bool, error) {
	return false, nil
}
func (m *mockTenantRepo) ExistsBySKU(ctx context.Context, sku, tenantID string) (bool, error) {
	return false, nil
}
func (m *mockTenantRepo) ExistsByName(ctx context.Context, name, tenantID string) (bool, error) {
	return false, nil
}
func (m *mockTenantRepo) ExistsByNameExcludingID(ctx context.Context, name, tenantID string, excludeID uuid.UUID) (bool, error) {
	return false, nil
}
func (m *mockTenantRepo) ExistsBySKUExcludingID(ctx context.Context, sku, tenantID string, excludeID uuid.UUID) (bool, error) {
	return false, nil
}
func (m *mockTenantRepo) VariantExistsByName(ctx context.Context, name string, productID uuid.UUID, tenantID string) (bool, error) {
	return false, nil
}
func (m *mockTenantRepo) VariantExistsBySKU(ctx context.Context, sku, tenantID string) (bool, error) {
	return false, nil
}
func (m *mockTenantRepo) VariantExistsByNameExcludingID(ctx context.Context, name string, productID uuid.UUID, tenantID string, excludeID uuid.UUID) (bool, error) {
	return false, nil
}
func (m *mockTenantRepo) VariantExistsBySKUExcludingID(ctx context.Context, sku, tenantID string, excludeID uuid.UUID) (bool, error) {
	return false, nil
}
func (m *mockTenantRepo) FindBySKUs(ctx context.Context, tenantID string, skus []string) ([]*tenantEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockTenantRepo) FindVariantsEnrichedBySKUs(ctx context.Context, tenantID string, skus []string) ([]tenantPort.VariantEnrichedRow, error) {
	return nil, nil
}
func (m *mockTenantRepo) FindByProduct(ctx context.Context, tenantID string, productID uuid.UUID) ([]*tenantEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockTenantRepo) GetBySKU(ctx context.Context, sku string, tenantID uuid.UUID) (*tenantEntity.ProductVariant, error) {
	return nil, nil
}

// ---- Helpers ----

func buildProductWithoutImage(t *testing.T, tenantID, name string) *tenantEntity.Product {
	t.Helper()
	status, err := tenantVO.NewProductStatus("active")
	if err != nil {
		t.Fatalf("NewProductStatus: %v", err)
	}
	p, err := tenantEntity.NewProductFromRepository(
		uuid.New(), tenantID, name,
		nil, nil, nil,
		nil, nil,
		status,
		time.Now(), time.Now(),
	)
	if err != nil {
		t.Fatalf("NewProductFromRepository: %v", err)
	}
	return p
}

func buildGlobalProductWithImage(t *testing.T, name, imageURL string) *globalEntity.GlobalProduct {
	t.Helper()
	qs, _ := globalVO.NewQualityScore(80)
	src, _ := globalVO.NewProductSource("test", nil, nil, 0.9)
	img := imageURL
	gp, err := globalEntity.NewGlobalProductFromRepository(
		uuid.New(), nil, name, nil, nil, nil, nil,
		&img, nil,
		src, qs,
		false, true,
		nil, nil, nil,
		time.Now(), time.Now(), nil,
	)
	if err != nil {
		t.Fatalf("NewGlobalProductFromRepository: %v", err)
	}
	return gp
}

// ---- Tests ----

func TestBackfill_MatchFound_UpdatesImage(t *testing.T) {
	// Arrange
	tenantID := "tenant-001"
	productName := "Leche Entera"
	imageURL := "https://cdn.example.com/leche.jpg"

	product := buildProductWithoutImage(t, tenantID, productName)
	globalProduct := buildGlobalProductWithImage(t, productName, imageURL)

	tenantRepo := &mockTenantRepo{
		findWithoutImageFn: func(ctx context.Context, tid string) ([]*tenantEntity.Product, error) {
			return []*tenantEntity.Product{product}, nil
		},
	}
	globalRepo := &mockGlobalRepo{
		findByNameAndBrandFn: func(ctx context.Context, name, brand string) (*globalEntity.GlobalProduct, error) {
			return globalProduct, nil
		},
	}

	uc := usecase.NewBackfillTenantImagesUseCase(globalRepo, tenantRepo)

	// Act
	result, err := uc.Execute(context.Background(), tenantID)

	// Assert
	if err != nil {
		t.Fatalf("esperaba nil error, obtuvo: %v", err)
	}
	if result.Updated != 1 {
		t.Errorf("esperaba Updated=1, obtuvo %d", result.Updated)
	}
	if result.Skipped != 0 {
		t.Errorf("esperaba Skipped=0, obtuvo %d", result.Skipped)
	}
	if result.Errors != 0 {
		t.Errorf("esperaba Errors=0, obtuvo %d", result.Errors)
	}
	if len(tenantRepo.updatedProductIDs) != 1 || tenantRepo.updatedProductIDs[0] != product.IDString() {
		t.Errorf("se esperaba actualizar product_id=%s", product.IDString())
	}
}

func TestBackfill_NoMatch_Skips(t *testing.T) {
	// Arrange
	tenantID := "tenant-002"
	product := buildProductWithoutImage(t, tenantID, "Producto Sin Match")

	tenantRepo := &mockTenantRepo{
		findWithoutImageFn: func(ctx context.Context, tid string) ([]*tenantEntity.Product, error) {
			return []*tenantEntity.Product{product}, nil
		},
	}
	globalRepo := &mockGlobalRepo{
		findByNameAndBrandFn: func(ctx context.Context, name, brand string) (*globalEntity.GlobalProduct, error) {
			return nil, nil // sin match
		},
	}

	uc := usecase.NewBackfillTenantImagesUseCase(globalRepo, tenantRepo)

	// Act
	result, err := uc.Execute(context.Background(), tenantID)

	// Assert
	if err != nil {
		t.Fatalf("esperaba nil error, obtuvo: %v", err)
	}
	if result.Skipped != 1 {
		t.Errorf("esperaba Skipped=1, obtuvo %d", result.Skipped)
	}
	if result.Updated != 0 {
		t.Errorf("esperaba Updated=0, obtuvo %d", result.Updated)
	}
	if len(tenantRepo.updatedProductIDs) != 0 {
		t.Error("no se debería haber llamado a UpdateImageURL")
	}
}

func TestBackfill_AlreadyHasImage_NotIncluded(t *testing.T) {
	// Arrange: FindWithoutImage retorna lista vacía (el producto con imagen no es incluido)
	tenantID := "tenant-003"

	tenantRepo := &mockTenantRepo{
		findWithoutImageFn: func(ctx context.Context, tid string) ([]*tenantEntity.Product, error) {
			return []*tenantEntity.Product{}, nil // ningún producto sin imagen
		},
	}
	globalRepo := &mockGlobalRepo{
		findByNameAndBrandFn: func(ctx context.Context, name, brand string) (*globalEntity.GlobalProduct, error) {
			t.Error("FindByNameAndBrand no debería haberse llamado")
			return nil, nil
		},
	}

	uc := usecase.NewBackfillTenantImagesUseCase(globalRepo, tenantRepo)

	// Act
	result, err := uc.Execute(context.Background(), tenantID)

	// Assert
	if err != nil {
		t.Fatalf("esperaba nil error, obtuvo: %v", err)
	}
	if result.Updated != 0 || result.Skipped != 0 || result.Errors != 0 {
		t.Errorf("resultado esperado {0,0,0}, obtuvo %+v", result)
	}
}
