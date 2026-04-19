package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/attribute/domain/entity"
	brandEntity "saas-mt-pim-service/src/brand/domain/entity"
	categoryEntity "saas-mt-pim-service/src/category/domain/entity"
	globalEntity "saas-mt-pim-service/src/product/global_catalog/domain/entity"

	"saas-mt-pim-service/src/overview/application/request"
	"saas-mt-pim-service/src/overview/application/usecase"
	cr "github.com/mercadocercano/criteria"
)

// MockMarketplaceCategoryRepo mock para categorías marketplace
type MockMarketplaceCategoryRepo struct {
	mock.Mock
}

func (m *MockMarketplaceCategoryRepo) Save(ctx context.Context, category *categoryEntity.MarketplaceCategory) error {
	return m.Called(ctx, category).Error(0)
}
func (m *MockMarketplaceCategoryRepo) GetByID(ctx context.Context, id string) (*categoryEntity.MarketplaceCategory, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*categoryEntity.MarketplaceCategory), args.Error(1)
}
func (m *MockMarketplaceCategoryRepo) GetBySlug(ctx context.Context, slug string) (*categoryEntity.MarketplaceCategory, error) {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*categoryEntity.MarketplaceCategory), args.Error(1)
}
func (m *MockMarketplaceCategoryRepo) GetByParentID(ctx context.Context, parentID *string) ([]*categoryEntity.MarketplaceCategory, error) {
	args := m.Called(ctx, parentID)
	return args.Get(0).([]*categoryEntity.MarketplaceCategory), args.Error(1)
}
func (m *MockMarketplaceCategoryRepo) GetRootCategories(ctx context.Context) ([]*categoryEntity.MarketplaceCategory, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*categoryEntity.MarketplaceCategory), args.Error(1)
}
func (m *MockMarketplaceCategoryRepo) GetTree(ctx context.Context) ([]*categoryEntity.MarketplaceCategory, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*categoryEntity.MarketplaceCategory), args.Error(1)
}
func (m *MockMarketplaceCategoryRepo) FindByCriteria(ctx context.Context, crit cr.Criteria) ([]*categoryEntity.MarketplaceCategory, error) {
	args := m.Called(ctx, crit)
	return args.Get(0).([]*categoryEntity.MarketplaceCategory), args.Error(1)
}
func (m *MockMarketplaceCategoryRepo) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	args := m.Called(ctx, crit)
	return args.Int(0), args.Error(1)
}
func (m *MockMarketplaceCategoryRepo) Update(ctx context.Context, category *categoryEntity.MarketplaceCategory) error {
	return m.Called(ctx, category).Error(0)
}
func (m *MockMarketplaceCategoryRepo) Delete(ctx context.Context, id string) error {
	return m.Called(ctx, id).Error(0)
}
func (m *MockMarketplaceCategoryRepo) ExistsBySlug(ctx context.Context, slug string) (bool, error) {
	args := m.Called(ctx, slug)
	return args.Bool(0), args.Error(1)
}
func (m *MockMarketplaceCategoryRepo) GetCategoryPath(ctx context.Context, categoryID string) ([]*categoryEntity.MarketplaceCategory, error) {
	args := m.Called(ctx, categoryID)
	return args.Get(0).([]*categoryEntity.MarketplaceCategory), args.Error(1)
}

// MockMarketplaceBrandRepo mock para marcas marketplace
type MockMarketplaceBrandRepo struct {
	mock.Mock
}

func (m *MockMarketplaceBrandRepo) Create(ctx context.Context, b *brandEntity.Marketplacebrand) error {
	return m.Called(ctx, b).Error(0)
}
func (m *MockMarketplaceBrandRepo) Update(ctx context.Context, b *brandEntity.Marketplacebrand) error {
	return m.Called(ctx, b).Error(0)
}
func (m *MockMarketplaceBrandRepo) FindByID(ctx context.Context, id string) (*brandEntity.Marketplacebrand, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*brandEntity.Marketplacebrand), args.Error(1)
}
func (m *MockMarketplaceBrandRepo) FindAll(ctx context.Context) ([]*brandEntity.Marketplacebrand, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*brandEntity.Marketplacebrand), args.Error(1)
}
func (m *MockMarketplaceBrandRepo) Delete(ctx context.Context, id string) error {
	return m.Called(ctx, id).Error(0)
}
func (m *MockMarketplaceBrandRepo) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*brandEntity.Marketplacebrand, error) {
	args := m.Called(ctx, crit)
	return args.Get(0).([]*brandEntity.Marketplacebrand), args.Error(1)
}
func (m *MockMarketplaceBrandRepo) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	args := m.Called(ctx, crit)
	return args.Int(0), args.Error(1)
}

// MockMarketplaceAttributeRepo mock para atributos marketplace
type MockMarketplaceAttributeRepo struct {
	mock.Mock
}

func (m *MockMarketplaceAttributeRepo) Create(ctx context.Context, a *entity.MarketplaceAttribute) error {
	return m.Called(ctx, a).Error(0)
}
func (m *MockMarketplaceAttributeRepo) Update(ctx context.Context, a *entity.MarketplaceAttribute) error {
	return m.Called(ctx, a).Error(0)
}
func (m *MockMarketplaceAttributeRepo) FindByID(ctx context.Context, id string) (*entity.MarketplaceAttribute, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.MarketplaceAttribute), args.Error(1)
}
func (m *MockMarketplaceAttributeRepo) FindAll(ctx context.Context) ([]*entity.MarketplaceAttribute, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.MarketplaceAttribute), args.Error(1)
}
func (m *MockMarketplaceAttributeRepo) FindByName(ctx context.Context, name string) (*entity.MarketplaceAttribute, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.MarketplaceAttribute), args.Error(1)
}
func (m *MockMarketplaceAttributeRepo) FindBySlug(ctx context.Context, slug string) (*entity.MarketplaceAttribute, error) {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.MarketplaceAttribute), args.Error(1)
}
func (m *MockMarketplaceAttributeRepo) Delete(ctx context.Context, id string) error {
	return m.Called(ctx, id).Error(0)
}
func (m *MockMarketplaceAttributeRepo) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*entity.MarketplaceAttribute, error) {
	args := m.Called(ctx, crit)
	return args.Get(0).([]*entity.MarketplaceAttribute), args.Error(1)
}
func (m *MockMarketplaceAttributeRepo) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	args := m.Called(ctx, crit)
	return args.Int(0), args.Error(1)
}

// MockGlobalProductRepo mock para productos globales
type MockGlobalProductRepo struct {
	mock.Mock
}

func (m *MockGlobalProductRepo) Save(p *globalEntity.GlobalProduct) (*globalEntity.GlobalProduct, error) {
	args := m.Called(p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) Update(p *globalEntity.GlobalProduct) (*globalEntity.GlobalProduct, error) {
	args := m.Called(p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindByID(id string) (*globalEntity.GlobalProduct, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) Delete(id string) error                   { return m.Called(id).Error(0) }
func (m *MockGlobalProductRepo) FindByEAN(ean string) (*globalEntity.GlobalProduct, error) {
	args := m.Called(ean)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindActiveByEAN(ean string) (*globalEntity.GlobalProduct, error) {
	args := m.Called(ean)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindByBusinessType(bt string, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(bt, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindBySource(s string, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(s, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindByQualityScoreRange(min, max, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(min, max, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) SearchByName(name string, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(name, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) SearchByBrand(brand string, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(brand, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) SearchByCategory(cat string, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(cat, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) SearchByTags(tags []string, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(tags, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindAll(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindActive(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindVerified(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindArgentineProducts(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindHighQualityProducts(offset, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(offset, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) FindNeedingUpdate(maxAge, limit int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(maxAge, limit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) CountTotal() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
func (m *MockGlobalProductRepo) CountBySource(source string) (int, error) {
	args := m.Called(source)
	return args.Int(0), args.Error(1)
}
func (m *MockGlobalProductRepo) CountByQualityScore(minScore int) (int, error) {
	args := m.Called(minScore)
	return args.Int(0), args.Error(1)
}
func (m *MockGlobalProductRepo) CountArgentineProducts() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
func (m *MockGlobalProductRepo) SearchByCriteria(ctx context.Context, crit cr.Criteria) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(ctx, crit)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) CountByCriteria(ctx context.Context, crit cr.Criteria) (int, error) {
	args := m.Called(ctx, crit)
	return args.Int(0), args.Error(1)
}

func (m *MockGlobalProductRepo) FindDistinctBrandsByBusinessType(businessType string) ([]string, error) {
	args := m.Called(businessType)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockGlobalProductRepo) FindDistinctCategoriesByBusinessType(businessType string) ([]string, error) {
	args := m.Called(businessType)
	return args.Get(0).([]string), args.Error(1)
}
func (m *MockGlobalProductRepo) FindNeedingEnrichment(businessType *string, limit, offset int) ([]*globalEntity.GlobalProduct, error) {
	args := m.Called(businessType, limit, offset)
	return args.Get(0).([]*globalEntity.GlobalProduct), args.Error(1)
}
func (m *MockGlobalProductRepo) CountNeedingEnrichment(businessType *string) (int, error) {
	args := m.Called(businessType)
	return args.Int(0), args.Error(1)
}

func TestGetMarketplaceOverviewUseCase_Execute_DashboardSection(t *testing.T) {
	// Arrange
	mockCategoryRepo := new(MockMarketplaceCategoryRepo)
	mockBrandRepo := new(MockMarketplaceBrandRepo)
	mockAttributeRepo := new(MockMarketplaceAttributeRepo)
	mockProductRepo := new(MockGlobalProductRepo)

	uc := usecase.NewGetMarketplaceOverviewUseCase(
		mockCategoryRepo, mockBrandRepo, mockAttributeRepo, mockProductRepo,
	)

	mockCategoryRepo.On("CountByCriteria", mock.Anything, mock.Anything).Return(50, nil)
	mockBrandRepo.On("CountByCriteria", mock.Anything, mock.Anything).Return(25, nil)
	mockAttributeRepo.On("CountByCriteria", mock.Anything, mock.Anything).Return(30, nil)
	mockProductRepo.On("CountTotal").Return(200, nil)

	req := &request.GetMarketplaceOverviewRequest{
		Sections:      []string{"dashboard"},
		Parallel:      false,
		TimeRangeDays: 7,
		Limit:         10,
	}

	// Act
	result, err := uc.Execute(context.Background(), req)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Contains(t, result.Data, "dashboard")

	dashboard := result.Data["dashboard"].(map[string]interface{})
	assert.Equal(t, 50, dashboard["total_categories"])
	assert.Equal(t, 25, dashboard["total_brands"])
	assert.Equal(t, 30, dashboard["total_attributes"])
	assert.Equal(t, 200, dashboard["total_global_products"])
}

func TestGetMarketplaceOverviewUseCase_Execute_NoSections_ShouldFail(t *testing.T) {
	// Arrange
	uc := usecase.NewGetMarketplaceOverviewUseCase(
		new(MockMarketplaceCategoryRepo),
		new(MockMarketplaceBrandRepo),
		new(MockMarketplaceAttributeRepo),
		new(MockGlobalProductRepo),
	)

	req := &request.GetMarketplaceOverviewRequest{
		Sections: []string{},
	}

	// Act
	result, err := uc.Execute(context.Background(), req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "al menos una sección")
}

func TestGetMarketplaceOverviewUseCase_Execute_InvalidSection_ShouldFail(t *testing.T) {
	// Arrange
	uc := usecase.NewGetMarketplaceOverviewUseCase(
		new(MockMarketplaceCategoryRepo),
		new(MockMarketplaceBrandRepo),
		new(MockMarketplaceAttributeRepo),
		new(MockGlobalProductRepo),
	)

	req := &request.GetMarketplaceOverviewRequest{
		Sections:      []string{"invalid-section"},
		TimeRangeDays: 7,
		Limit:         10,
	}

	// Act
	result, err := uc.Execute(context.Background(), req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "inválida")
}

func TestGetMarketplaceOverviewUseCase_Execute_DashboardWithRepositoryErrors(t *testing.T) {
	// Arrange
	mockCategoryRepo := new(MockMarketplaceCategoryRepo)
	mockBrandRepo := new(MockMarketplaceBrandRepo)
	mockAttributeRepo := new(MockMarketplaceAttributeRepo)
	mockProductRepo := new(MockGlobalProductRepo)

	uc := usecase.NewGetMarketplaceOverviewUseCase(
		mockCategoryRepo, mockBrandRepo, mockAttributeRepo, mockProductRepo,
	)

	// Simular errores - el use case usa fallback a 0
	mockCategoryRepo.On("CountByCriteria", mock.Anything, mock.Anything).Return(0, errors.New("db error"))
	mockBrandRepo.On("CountByCriteria", mock.Anything, mock.Anything).Return(0, errors.New("db error"))
	mockAttributeRepo.On("CountByCriteria", mock.Anything, mock.Anything).Return(0, errors.New("db error"))
	mockProductRepo.On("CountTotal").Return(0, errors.New("db error"))

	req := &request.GetMarketplaceOverviewRequest{
		Sections:      []string{"dashboard"},
		Parallel:      false,
		TimeRangeDays: 7,
		Limit:         10,
	}

	// Act
	result, err := uc.Execute(context.Background(), req)

	// Assert - use case should NOT fail, it uses fallback values
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)

	dashboard := result.Data["dashboard"].(map[string]interface{})
	assert.Equal(t, 0, dashboard["total_categories"])
	assert.Equal(t, 0, dashboard["total_brands"])
}

func TestGetMarketplaceOverviewUseCase_Execute_DefaultsApplied(t *testing.T) {
	// Arrange
	mockCategoryRepo := new(MockMarketplaceCategoryRepo)
	mockBrandRepo := new(MockMarketplaceBrandRepo)
	mockAttributeRepo := new(MockMarketplaceAttributeRepo)
	mockProductRepo := new(MockGlobalProductRepo)

	uc := usecase.NewGetMarketplaceOverviewUseCase(
		mockCategoryRepo, mockBrandRepo, mockAttributeRepo, mockProductRepo,
	)

	mockCategoryRepo.On("CountByCriteria", mock.Anything, mock.Anything).Return(10, nil)
	mockBrandRepo.On("CountByCriteria", mock.Anything, mock.Anything).Return(5, nil)
	mockAttributeRepo.On("CountByCriteria", mock.Anything, mock.Anything).Return(8, nil)
	mockProductRepo.On("CountTotal").Return(100, nil)

	req := &request.GetMarketplaceOverviewRequest{
		Sections:      []string{"dashboard"},
		TimeRangeDays: 0,  // Should be defaulted to 7
		Limit:         0,  // Should be defaulted to 10
	}

	// Act
	result, err := uc.Execute(context.Background(), req)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
}
