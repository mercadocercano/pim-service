package usecase_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cr "github.com/mercadocercano/criteria"
	globalEntity "saas-mt-pim-service/src/product/global_catalog/domain/entity"
	globalVO "saas-mt-pim-service/src/product/global_catalog/domain/value_object"
	"saas-mt-pim-service/src/product/quickstart/application/usecase"
	tenantEntity "saas-mt-pim-service/src/product/tenant/domain/entity"
	tenantPort "saas-mt-pim-service/src/product/tenant/domain/port"
)

// --- Mocks ---

type mockGlobalProductRepo struct {
	product *globalEntity.GlobalProduct
	err     error
}

func (m *mockGlobalProductRepo) FindByID(id string) (*globalEntity.GlobalProduct, error) {
	return m.product, m.err
}

// Stub all other methods (not used by this use case)
func (m *mockGlobalProductRepo) Save(_ *globalEntity.GlobalProduct) (*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) Update(_ *globalEntity.GlobalProduct) (*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) Delete(_ string) error { return nil }
func (m *mockGlobalProductRepo) FindByEAN(_ string) (*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindActiveByEAN(_ string) (*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindByBusinessType(_ string, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindBySource(_ string, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindByQualityScoreRange(_, _, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) SearchByName(_ string, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) SearchByBrand(_ string, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) SearchByCategory(_ string, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) SearchByTags(_ []string, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindAll(_, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindActive(_, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindVerified(_, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindArgentineProducts(_, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindHighQualityProducts(_, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindNeedingUpdate(_ int, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindNeedingEnrichment(_ *string, _, _ int) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) CountNeedingEnrichment(_ *string) (int, error) { return 0, nil }
func (m *mockGlobalProductRepo) FindDistinctBrandsByBusinessType(_ string) ([]string, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindDistinctCategoriesByBusinessType(_ string) ([]string, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) CountTotal() (int, error)               { return 0, nil }
func (m *mockGlobalProductRepo) CountBySource(_ string) (int, error)    { return 0, nil }
func (m *mockGlobalProductRepo) CountByQualityScore(_ int) (int, error) { return 0, nil }
func (m *mockGlobalProductRepo) CountArgentineProducts() (int, error)   { return 0, nil }
func (m *mockGlobalProductRepo) SearchByCriteria(_ context.Context, _ cr.Criteria) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) CountByCriteria(_ context.Context, _ cr.Criteria) (int, error) {
	return 0, nil
}
func (m *mockGlobalProductRepo) FindByIDs(_ context.Context, _ []string) ([]*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindByNameAndBrand(_ context.Context, _, _ string) (*globalEntity.GlobalProduct, error) {
	return nil, nil
}
func (m *mockGlobalProductRepo) FindDistinctBusinessTypes() ([]string, error) { return nil, nil }
func (m *mockGlobalProductRepo) CountTenantLinks(_ context.Context, _ string) (int, error) {
	return 0, nil
}

type mockTenantProductRepo struct {
	saved *tenantEntity.Product
	err   error
}

func (m *mockTenantProductRepo) Save(_ context.Context, p *tenantEntity.Product) error {
	m.saved = p
	return m.err
}

// Stub other methods — ProductRepository interface
func (m *mockTenantProductRepo) FindByID(_ context.Context, _ uuid.UUID, _ string) (*tenantEntity.Product, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) FindByIDWithVariants(_ context.Context, _ uuid.UUID, _ string) (*tenantEntity.Product, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) FindBySKU(_ context.Context, _, _ string) (*tenantEntity.Product, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) Update(_ context.Context, _ *tenantEntity.Product) error { return nil }
func (m *mockTenantProductRepo) Delete(_ context.Context, _ uuid.UUID, _ string) error   { return nil }
func (m *mockTenantProductRepo) SaveVariant(_ context.Context, _ uuid.UUID, _ *tenantEntity.ProductVariant) error {
	return nil
}
func (m *mockTenantProductRepo) UpdateVariant(_ context.Context, _ *tenantEntity.ProductVariant) error {
	return nil
}
func (m *mockTenantProductRepo) DeleteVariant(_ context.Context, _ uuid.UUID) error { return nil }
func (m *mockTenantProductRepo) LoadVariantsForProduct(_ context.Context, _ uuid.UUID) ([]*tenantEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) ExistsByID(_ context.Context, _ uuid.UUID, _ string) (bool, error) {
	return false, nil
}
func (m *mockTenantProductRepo) ExistsBySKU(_ context.Context, _, _ string) (bool, error) {
	return false, nil
}
func (m *mockTenantProductRepo) ExistsByName(_ context.Context, _, _ string) (bool, error) {
	return false, nil
}
func (m *mockTenantProductRepo) ExistsByNameExcludingID(_ context.Context, _, _ string, _ uuid.UUID) (bool, error) {
	return false, nil
}
func (m *mockTenantProductRepo) ExistsBySKUExcludingID(_ context.Context, _, _ string, _ uuid.UUID) (bool, error) {
	return false, nil
}
func (m *mockTenantProductRepo) VariantExistsByName(_ context.Context, _ string, _ uuid.UUID, _ string) (bool, error) {
	return false, nil
}
func (m *mockTenantProductRepo) VariantExistsBySKU(_ context.Context, _, _ string) (bool, error) {
	return false, nil
}
func (m *mockTenantProductRepo) VariantExistsByNameExcludingID(_ context.Context, _ string, _ uuid.UUID, _ string, _ uuid.UUID) (bool, error) {
	return false, nil
}
func (m *mockTenantProductRepo) VariantExistsBySKUExcludingID(_ context.Context, _, _ string, _ uuid.UUID) (bool, error) {
	return false, nil
}
func (m *mockTenantProductRepo) FindBySKUs(_ context.Context, _ string, _ []string) ([]*tenantEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) FindVariantsEnrichedBySKUs(_ context.Context, _ string, _ []string) ([]tenantPort.VariantEnrichedRow, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) FindByProduct(_ context.Context, _ string, _ uuid.UUID) ([]*tenantEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) GetBySKU(_ context.Context, _ string, _ uuid.UUID) (*tenantEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) FindDistinctTenantIDs(_ context.Context) ([]string, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) FindWithoutImage(_ context.Context, _ string) ([]*tenantEntity.Product, error) {
	return nil, nil
}
func (m *mockTenantProductRepo) UpdateImageURL(_ context.Context, _, _, _ string) error { return nil }

// --- Helpers ---

func buildTestGlobalProduct(active bool) *globalEntity.GlobalProduct {
	brand := "Arcor"
	category := "golosinas"
	price := 1500.0
	desc := "Alfajor triple de chocolate"
	img := "https://example.com/alfajor.jpg"
	bt := "kiosco"
	source, _ := globalVO.NewProductSource("seed", nil, nil, 0.5)
	qs, _ := globalVO.NewQualityScore(75)

	gp, _ := globalEntity.NewGlobalProductFromRepository(
		uuid.New(), nil, "Alfajor Jorgito Triple", &desc, &brand, &category, &price,
		&img, nil, source, qs, false, active, &bt, nil, nil,
		time.Now(), time.Now(), nil,
	)
	return gp
}

// --- Tests ---

func TestImportFromGlobalCatalog_Success(t *testing.T) {
	// Arrange
	gp := buildTestGlobalProduct(true)
	globalRepo := &mockGlobalProductRepo{product: gp}
	tenantRepo := &mockTenantProductRepo{}
	uc := usecase.NewImportFromGlobalCatalogUseCase(tenantRepo, globalRepo)

	req := usecase.ImportFromGlobalCatalogRequest{
		TenantID:        "tenant-123",
		GlobalProductID: gp.IDString(),
	}

	// Act
	resp, err := uc.Execute(context.Background(), req)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, gp.IDString(), resp.GlobalProductID)
	assert.Equal(t, "Alfajor Jorgito Triple", resp.Name)
	assert.Equal(t, "Arcor", *resp.Brand)
	assert.Equal(t, "golosinas", *resp.Category)
	assert.Equal(t, 1500.0, *resp.Price)
	assert.NotEmpty(t, resp.ProductID)

	// Verificar que se guardó el producto
	require.NotNil(t, tenantRepo.saved)
	assert.Equal(t, "Alfajor Jorgito Triple", tenantRepo.saved.Name())
}

func TestImportFromGlobalCatalog_ProductNotFound(t *testing.T) {
	// Arrange
	globalRepo := &mockGlobalProductRepo{err: fmt.Errorf("not found")}
	tenantRepo := &mockTenantProductRepo{}
	uc := usecase.NewImportFromGlobalCatalogUseCase(tenantRepo, globalRepo)

	req := usecase.ImportFromGlobalCatalogRequest{
		TenantID:        "tenant-123",
		GlobalProductID: uuid.New().String(),
	}

	// Act
	resp, err := uc.Execute(context.Background(), req)

	// Assert
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "producto global no encontrado")
}

func TestImportFromGlobalCatalog_InactiveProduct(t *testing.T) {
	// Arrange
	gp := buildTestGlobalProduct(false)
	globalRepo := &mockGlobalProductRepo{product: gp}
	tenantRepo := &mockTenantProductRepo{}
	uc := usecase.NewImportFromGlobalCatalogUseCase(tenantRepo, globalRepo)

	req := usecase.ImportFromGlobalCatalogRequest{
		TenantID:        "tenant-123",
		GlobalProductID: gp.IDString(),
	}

	// Act
	resp, err := uc.Execute(context.Background(), req)

	// Assert
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no está activo")
}

func TestImportFromGlobalCatalog_MissingTenantID(t *testing.T) {
	// Arrange
	uc := usecase.NewImportFromGlobalCatalogUseCase(&mockTenantProductRepo{}, &mockGlobalProductRepo{})

	req := usecase.ImportFromGlobalCatalogRequest{
		GlobalProductID: uuid.New().String(),
	}

	// Act
	resp, err := uc.Execute(context.Background(), req)

	// Assert
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "tenant_id es requerido")
}

func TestImportFromGlobalCatalog_MissingGlobalProductID(t *testing.T) {
	// Arrange
	uc := usecase.NewImportFromGlobalCatalogUseCase(&mockTenantProductRepo{}, &mockGlobalProductRepo{})

	req := usecase.ImportFromGlobalCatalogRequest{
		TenantID: "tenant-123",
	}

	// Act
	resp, err := uc.Execute(context.Background(), req)

	// Assert
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "global_product_id es requerido")
}

func TestImportFromGlobalCatalog_SaveError(t *testing.T) {
	// Arrange
	gp := buildTestGlobalProduct(true)
	globalRepo := &mockGlobalProductRepo{product: gp}
	tenantRepo := &mockTenantProductRepo{err: fmt.Errorf("db connection error")}
	uc := usecase.NewImportFromGlobalCatalogUseCase(tenantRepo, globalRepo)

	req := usecase.ImportFromGlobalCatalogRequest{
		TenantID:        "tenant-123",
		GlobalProductID: gp.IDString(),
	}

	// Act
	resp, err := uc.Execute(context.Background(), req)

	// Assert
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "error guardando producto")
}
