package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"saas-mt-pim-service/src/batch/application/request"
	"saas-mt-pim-service/src/batch/application/usecase"
	batchPort "saas-mt-pim-service/src/batch/domain/port"
	brandEntity "saas-mt-pim-service/src/brand/domain/entity"
	categoryEntity "saas-mt-pim-service/src/category/domain/entity"
	productEntity "saas-mt-pim-service/src/product/tenant/domain/entity"
	productPort "saas-mt-pim-service/src/product/tenant/domain/port"

	googleUUID "github.com/google/uuid"
	cr "github.com/mercadocercano/criteria"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- Transaction Mocks ---

type mockTransaction struct {
	commitErr   error
	rollbackErr error
	committed   bool
}

func (m *mockTransaction) Commit() error {
	m.committed = true
	return m.commitErr
}

func (m *mockTransaction) Rollback() error {
	return m.rollbackErr
}

type mockTxBeginner struct {
	tx  batchPort.Transaction
	err error
}

func (m *mockTxBeginner) BeginTx(_ context.Context, _ *sql.TxOptions) (batchPort.Transaction, error) {
	return m.tx, m.err
}

// --- Category Repo Mock ---

type mockCategoryRepo struct {
	createErr error
	callCount int
}

func (m *mockCategoryRepo) Create(_ context.Context, _ *categoryEntity.Category) error {
	m.callCount++
	return m.createErr
}
func (m *mockCategoryRepo) FindByID(_ context.Context, _, _ string) (*categoryEntity.Category, error) {
	return nil, nil
}
func (m *mockCategoryRepo) FindAll(_ context.Context, _ string) ([]*categoryEntity.Category, error) {
	return nil, nil
}
func (m *mockCategoryRepo) Update(_ context.Context, _ *categoryEntity.Category) error { return nil }
func (m *mockCategoryRepo) Delete(_ context.Context, _, _ string) error                { return nil }

// --- Brand Repo Mock ---

type mockBrandRepo struct {
	createErr error
	callCount int
}

func (m *mockBrandRepo) Create(_ context.Context, _ *brandEntity.Brand) error {
	m.callCount++
	return m.createErr
}
func (m *mockBrandRepo) FindByID(_ context.Context, _, _ string) (*brandEntity.Brand, error) {
	return nil, nil
}
func (m *mockBrandRepo) FindByName(_ context.Context, _, _ string) (*brandEntity.Brand, error) {
	return nil, nil
}
func (m *mockBrandRepo) FindAll(_ context.Context, _ string) ([]*brandEntity.Brand, error) {
	return nil, nil
}
func (m *mockBrandRepo) Update(_ context.Context, _ *brandEntity.Brand) error { return nil }
func (m *mockBrandRepo) Delete(_ context.Context, _, _ string) error          { return nil }
func (m *mockBrandRepo) ExistsByName(_ context.Context, _, _ string, _ *string) (bool, error) {
	return false, nil
}

// --- Product Repo Mock ---

type mockProductRepo struct {
	saveErr   error
	callCount int
}

func (m *mockProductRepo) Save(_ context.Context, _ *productEntity.Product) error {
	m.callCount++
	return m.saveErr
}
func (m *mockProductRepo) FindByID(_ context.Context, _ googleUUID.UUID, _ string) (*productEntity.Product, error) {
	return nil, nil
}
func (m *mockProductRepo) FindByIDWithVariants(_ context.Context, _ googleUUID.UUID, _ string) (*productEntity.Product, error) {
	return nil, nil
}
func (m *mockProductRepo) FindBySKU(_ context.Context, _, _ string) (*productEntity.Product, error) {
	return nil, nil
}
func (m *mockProductRepo) Update(_ context.Context, _ *productEntity.Product) error { return nil }
func (m *mockProductRepo) Delete(_ context.Context, _ googleUUID.UUID, _ string) error {
	return nil
}
func (m *mockProductRepo) SaveVariant(_ context.Context, _ googleUUID.UUID, _ *productEntity.ProductVariant) error {
	return nil
}
func (m *mockProductRepo) UpdateVariant(_ context.Context, _ *productEntity.ProductVariant) error {
	return nil
}
func (m *mockProductRepo) DeleteVariant(_ context.Context, _ googleUUID.UUID) error { return nil }
func (m *mockProductRepo) LoadVariantsForProduct(_ context.Context, _ googleUUID.UUID) ([]*productEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockProductRepo) ExistsByID(_ context.Context, _ googleUUID.UUID, _ string) (bool, error) {
	return false, nil
}
func (m *mockProductRepo) ExistsBySKU(_ context.Context, _, _ string) (bool, error) {
	return false, nil
}
func (m *mockProductRepo) ExistsByName(_ context.Context, _, _ string) (bool, error) {
	return false, nil
}
func (m *mockProductRepo) ExistsByNameExcludingID(_ context.Context, _, _ string, _ googleUUID.UUID) (bool, error) {
	return false, nil
}
func (m *mockProductRepo) ExistsBySKUExcludingID(_ context.Context, _, _ string, _ googleUUID.UUID) (bool, error) {
	return false, nil
}
func (m *mockProductRepo) VariantExistsByName(_ context.Context, _ string, _ googleUUID.UUID, _ string) (bool, error) {
	return false, nil
}
func (m *mockProductRepo) VariantExistsBySKU(_ context.Context, _, _ string) (bool, error) {
	return false, nil
}
func (m *mockProductRepo) VariantExistsByNameExcludingID(_ context.Context, _ string, _ googleUUID.UUID, _ string, _ googleUUID.UUID) (bool, error) {
	return false, nil
}
func (m *mockProductRepo) VariantExistsBySKUExcludingID(_ context.Context, _, _ string, _ googleUUID.UUID) (bool, error) {
	return false, nil
}
func (m *mockProductRepo) FindBySKUs(_ context.Context, _ string, _ []string) ([]*productEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockProductRepo) FindVariantsEnrichedBySKUs(_ context.Context, _ string, _ []string) ([]productPort.VariantEnrichedRow, error) {
	return nil, nil
}
func (m *mockProductRepo) FindByProduct(_ context.Context, _ string, _ googleUUID.UUID) ([]*productEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockProductRepo) GetBySKU(_ context.Context, _ string, _ googleUUID.UUID) (*productEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockProductRepo) SearchByCriteria(_ context.Context, _ cr.Criteria) ([]*productEntity.Product, error) {
	return nil, nil
}
func (m *mockProductRepo) CountByCriteria(_ context.Context, _ cr.Criteria) (int, error) {
	return 0, nil
}
func (m *mockProductRepo) FindVariantsByCriteria(_ context.Context, _ *cr.Criteria) ([]*productEntity.ProductVariant, error) {
	return nil, nil
}
func (m *mockProductRepo) CountVariantsByCriteria(_ context.Context, _ *cr.Criteria) (int, error) {
	return 0, nil
}

// --- Category Mapping Repo Mock ---

type mockCategoryMappingRepo struct {
	saveErr   error
	callCount int
}

func (m *mockCategoryMappingRepo) Save(_ context.Context, _ *categoryEntity.TenantCategoryMapping) error {
	m.callCount++
	return m.saveErr
}
func (m *mockCategoryMappingRepo) GetByID(_ context.Context, _ string) (*categoryEntity.TenantCategoryMapping, error) {
	return nil, nil
}
func (m *mockCategoryMappingRepo) GetByTenantAndMarketplaceCategory(_ context.Context, _, _ string) (*categoryEntity.TenantCategoryMapping, error) {
	return nil, nil
}
func (m *mockCategoryMappingRepo) GetByTenantID(_ context.Context, _ string) ([]*categoryEntity.TenantCategoryMapping, error) {
	return nil, nil
}
func (m *mockCategoryMappingRepo) GetByMarketplaceCategoryID(_ context.Context, _ string) ([]*categoryEntity.TenantCategoryMapping, error) {
	return nil, nil
}
func (m *mockCategoryMappingRepo) FindByCriteria(_ context.Context, _ cr.Criteria) ([]*categoryEntity.TenantCategoryMapping, error) {
	return nil, nil
}
func (m *mockCategoryMappingRepo) CountByCriteria(_ context.Context, _ cr.Criteria) (int, error) {
	return 0, nil
}
func (m *mockCategoryMappingRepo) Update(_ context.Context, _ *categoryEntity.TenantCategoryMapping) error {
	return nil
}
func (m *mockCategoryMappingRepo) Delete(_ context.Context, _ string) error { return nil }
func (m *mockCategoryMappingRepo) GetTenantTaxonomy(_ context.Context, _ string) ([]*categoryEntity.TenantCategoryMapping, error) {
	return nil, nil
}

// --- Helper ---

func validTenantID() string {
	return "550e8400-e29b-41d4-a716-446655440000"
}

func newBatchUC(
	txBeginner batchPort.TxBeginner,
	catRepo *mockCategoryRepo,
	brandRepo *mockBrandRepo,
	prodRepo *mockProductRepo,
	mapRepo *mockCategoryMappingRepo,
) *usecase.BatchCreateUseCase {
	return usecase.NewBatchCreateUseCase(txBeginner, catRepo, brandRepo, prodRepo, mapRepo)
}

// --- Tests ---

func TestBatchCreate_WithCategories_CreatesSuccessfully(t *testing.T) {
	tx := &mockTransaction{}
	catRepo := &mockCategoryRepo{}
	uc := newBatchUC(&mockTxBeginner{tx: tx}, catRepo, &mockBrandRepo{}, &mockProductRepo{}, &mockCategoryMappingRepo{})

	req := &request.BatchCreateRequest{
		Categories: []request.CategoryBatchItem{
			{Name: "Electrónica", Description: "Productos electrónicos"},
			{Name: "Hogar", Description: "Productos del hogar"},
		},
	}

	result, err := uc.Execute(context.Background(), req, validTenantID())

	require.NoError(t, err)
	assert.Len(t, result.Created.Categories, 2)
	assert.Empty(t, result.Errors)
	assert.Equal(t, 2, catRepo.callCount)
	assert.True(t, tx.committed)
}

func TestBatchCreate_WithBrands_CreatesSuccessfully(t *testing.T) {
	tx := &mockTransaction{}
	brandRepo := &mockBrandRepo{}
	uc := newBatchUC(&mockTxBeginner{tx: tx}, &mockCategoryRepo{}, brandRepo, &mockProductRepo{}, &mockCategoryMappingRepo{})

	req := &request.BatchCreateRequest{
		Brands: []request.BrandBatchItem{
			{Name: "Samsung", Description: "Samsung Electronics"},
			{Name: "Apple", Description: "Apple Inc"},
		},
	}

	result, err := uc.Execute(context.Background(), req, validTenantID())

	require.NoError(t, err)
	assert.Len(t, result.Created.Brands, 2)
	assert.Empty(t, result.Errors)
	assert.Equal(t, 2, brandRepo.callCount)
	assert.True(t, tx.committed)
}

func TestBatchCreate_WithProducts_CreatesSuccessfully(t *testing.T) {
	tx := &mockTransaction{}
	prodRepo := &mockProductRepo{}
	uc := newBatchUC(&mockTxBeginner{tx: tx}, &mockCategoryRepo{}, &mockBrandRepo{}, prodRepo, &mockCategoryMappingRepo{})

	req := &request.BatchCreateRequest{
		Products: []request.ProductBatchItem{
			{Name: "Galaxy S24", Description: "Smartphone", SKU: "SAM-S24-001"},
		},
	}

	result, err := uc.Execute(context.Background(), req, validTenantID())

	require.NoError(t, err)
	assert.Len(t, result.Created.Products, 1)
	assert.Empty(t, result.Errors)
	assert.Equal(t, 1, prodRepo.callCount)
	assert.True(t, tx.committed)
}

func TestBatchCreate_WithMixedEntities_CreatesAll(t *testing.T) {
	tx := &mockTransaction{}
	uc := newBatchUC(&mockTxBeginner{tx: tx}, &mockCategoryRepo{}, &mockBrandRepo{}, &mockProductRepo{}, &mockCategoryMappingRepo{})

	req := &request.BatchCreateRequest{
		Categories: []request.CategoryBatchItem{{Name: "Electrónica", Description: "Test"}},
		Brands:     []request.BrandBatchItem{{Name: "Samsung", Description: "Test"}},
		Products:   []request.ProductBatchItem{{Name: "Galaxy S24", SKU: "SAM-001"}},
	}

	result, err := uc.Execute(context.Background(), req, validTenantID())

	require.NoError(t, err)
	assert.Len(t, result.Created.Categories, 1)
	assert.Len(t, result.Created.Brands, 1)
	assert.Len(t, result.Created.Products, 1)
	assert.Empty(t, result.Errors)
	assert.True(t, tx.committed)
}

func TestBatchCreate_BeginTxFails_ReturnsError(t *testing.T) {
	uc := newBatchUC(
		&mockTxBeginner{err: errors.New("connection refused")},
		&mockCategoryRepo{}, &mockBrandRepo{}, &mockProductRepo{}, &mockCategoryMappingRepo{},
	)

	req := &request.BatchCreateRequest{
		Categories: []request.CategoryBatchItem{{Name: "Test", Description: "Test"}},
	}

	result, err := uc.Execute(context.Background(), req, validTenantID())

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "error al iniciar transacción")
}

func TestBatchCreate_CategoryRepoFails_ReturnsPartialResult(t *testing.T) {
	tx := &mockTransaction{}
	uc := newBatchUC(
		&mockTxBeginner{tx: tx},
		&mockCategoryRepo{createErr: errors.New("duplicate key")},
		&mockBrandRepo{},
		&mockProductRepo{},
		&mockCategoryMappingRepo{},
	)

	req := &request.BatchCreateRequest{
		Categories: []request.CategoryBatchItem{{Name: "Electrónica", Description: "Test"}},
		Brands:     []request.BrandBatchItem{{Name: "Samsung", Description: "Test"}},
	}

	result, err := uc.Execute(context.Background(), req, validTenantID())

	require.NoError(t, err)
	assert.Empty(t, result.Created.Categories)
	assert.Len(t, result.Created.Brands, 1)
	assert.Len(t, result.Errors, 1)
	assert.Equal(t, "category", result.Errors[0].Type)
	assert.True(t, tx.committed)
}

func TestBatchCreate_AllFail_ReturnsError(t *testing.T) {
	tx := &mockTransaction{}
	uc := newBatchUC(
		&mockTxBeginner{tx: tx},
		&mockCategoryRepo{createErr: errors.New("fail")},
		&mockBrandRepo{createErr: errors.New("fail")},
		&mockProductRepo{saveErr: errors.New("fail")},
		&mockCategoryMappingRepo{},
	)

	req := &request.BatchCreateRequest{
		Categories: []request.CategoryBatchItem{{Name: "Cat", Description: "Test"}},
		Brands:     []request.BrandBatchItem{{Name: "Brand", Description: "Test"}},
		Products:   []request.ProductBatchItem{{Name: "Prod", SKU: "SKU-001"}},
	}

	_, err := uc.Execute(context.Background(), req, validTenantID())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ninguna entidad pudo ser creada")
}

func TestBatchCreate_CommitFails_ReturnsError(t *testing.T) {
	tx := &mockTransaction{commitErr: errors.New("commit failed")}
	uc := newBatchUC(&mockTxBeginner{tx: tx}, &mockCategoryRepo{}, &mockBrandRepo{}, &mockProductRepo{}, &mockCategoryMappingRepo{})

	req := &request.BatchCreateRequest{
		Categories: []request.CategoryBatchItem{{Name: "Electrónica", Description: "Test"}},
	}

	_, err := uc.Execute(context.Background(), req, validTenantID())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "error al confirmar transacción")
}

func TestBatchCreate_WithCategoryMapping_CreatesMappingToo(t *testing.T) {
	tx := &mockTransaction{}
	mapRepo := &mockCategoryMappingRepo{}
	uc := newBatchUC(&mockTxBeginner{tx: tx}, &mockCategoryRepo{}, &mockBrandRepo{}, &mockProductRepo{}, mapRepo)

	req := &request.BatchCreateRequest{
		Categories: []request.CategoryBatchItem{
			{
				Name:        "Electrónica",
				Description: "Test",
				Mapping: &request.CategoryMappingData{
					MarketplaceCategoryID: "mkt-cat-001",
					CustomName:            "Electro",
				},
			},
		},
	}

	result, err := uc.Execute(context.Background(), req, validTenantID())

	require.NoError(t, err)
	assert.Len(t, result.Created.Categories, 1)
	assert.Equal(t, 1, mapRepo.callCount)
	assert.True(t, tx.committed)
}

func TestBatchCreate_EmptyRequest_CommitsSuccessfully(t *testing.T) {
	tx := &mockTransaction{}
	uc := newBatchUC(&mockTxBeginner{tx: tx}, &mockCategoryRepo{}, &mockBrandRepo{}, &mockProductRepo{}, &mockCategoryMappingRepo{})

	result, err := uc.Execute(context.Background(), req(), validTenantID())

	require.NoError(t, err)
	assert.Empty(t, result.Created.Categories)
	assert.Empty(t, result.Created.Brands)
	assert.Empty(t, result.Created.Products)
	assert.True(t, tx.committed)
}

func req() *request.BatchCreateRequest {
	return &request.BatchCreateRequest{}
}
