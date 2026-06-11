package repository

import (
	"context"
	"errors"
	"sync"

	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/brand/domain/entity"
)

// Errores mock
var (
	ErrMockFailedOp   = errors.New("operación fallida (simulada)")
	ErrMockNotFound   = errors.New("recurso no encontrado (simulado)")
	ErrMockDuplicated = errors.New("recurso duplicado (simulado)")
)

// MockBrandRepository implementa BrandCriteriaRepository en memoria para pruebas
type MockBrandRepository struct {
	mu            sync.RWMutex
	brands        map[string]*entity.Brand
	failOnMethods map[string]bool
	callHistory   map[string]int
}

// NewMockBrandRepository crea una nueva instancia del mock
func NewMockBrandRepository() *MockBrandRepository {
	return &MockBrandRepository{
		brands:        make(map[string]*entity.Brand),
		failOnMethods: make(map[string]bool),
		callHistory:   make(map[string]int),
	}
}

// ShouldFailOn configura un método específico para que falle
func (r *MockBrandRepository) ShouldFailOn(method string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.failOnMethods[method] = true
}

// GetCallCount retorna el número de veces que se ha llamado a un método
func (r *MockBrandRepository) GetCallCount(method string) int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.callHistory[method]
}

// SetupBrands inicializa el repositorio con marcas predefinidas
func (r *MockBrandRepository) SetupBrands(brands []*entity.Brand) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.brands = make(map[string]*entity.Brand)
	for _, brand := range brands {
		r.brands[brand.ID] = cloneBrand(brand)
	}
}

// GetBrands retorna todas las marcas almacenadas
func (r *MockBrandRepository) GetBrands() []*entity.Brand {
	r.mu.RLock()
	defer r.mu.RUnlock()
	brands := make([]*entity.Brand, 0, len(r.brands))
	for _, brand := range r.brands {
		brands = append(brands, brand)
	}
	return brands
}

func (r *MockBrandRepository) shouldMethodFail(method string) bool {
	return r.failOnMethods[method]
}

func (r *MockBrandRepository) incrementCallCount(method string) {
	r.callHistory[method]++
}

func cloneBrand(b *entity.Brand) *entity.Brand {
	clone := *b
	if b.LogoURL != nil {
		v := *b.LogoURL
		clone.LogoURL = &v
	}
	if b.Website != nil {
		v := *b.Website
		clone.Website = &v
	}
	return &clone
}

// Create implementa BrandRepository
func (r *MockBrandRepository) Create(_ context.Context, brand *entity.Brand) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.incrementCallCount("Create")
	if r.shouldMethodFail("Create") {
		return ErrMockFailedOp
	}
	r.brands[brand.ID] = cloneBrand(brand)
	return nil
}

// FindByID implementa BrandRepository
func (r *MockBrandRepository) FindByID(_ context.Context, id string, tenantID string) (*entity.Brand, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.incrementCallCount("FindByID")
	if r.shouldMethodFail("FindByID") {
		return nil, ErrMockFailedOp
	}
	brand, exists := r.brands[id]
	if !exists || brand.TenantID != tenantID {
		return nil, nil
	}
	return cloneBrand(brand), nil
}

// FindByName implementa BrandRepository
func (r *MockBrandRepository) FindByName(_ context.Context, name string, tenantID string) (*entity.Brand, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.incrementCallCount("FindByName")
	if r.shouldMethodFail("FindByName") {
		return nil, ErrMockFailedOp
	}
	for _, brand := range r.brands {
		if brand.Name == name && brand.TenantID == tenantID {
			return cloneBrand(brand), nil
		}
	}
	return nil, nil
}

// FindAll implementa BrandRepository
func (r *MockBrandRepository) FindAll(_ context.Context, tenantID string) ([]*entity.Brand, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.incrementCallCount("FindAll")
	if r.shouldMethodFail("FindAll") {
		return nil, ErrMockFailedOp
	}
	result := make([]*entity.Brand, 0)
	for _, brand := range r.brands {
		if brand.TenantID == tenantID {
			result = append(result, cloneBrand(brand))
		}
	}
	return result, nil
}

// Update implementa BrandRepository
func (r *MockBrandRepository) Update(_ context.Context, brand *entity.Brand) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.incrementCallCount("Update")
	if r.shouldMethodFail("Update") {
		return ErrMockFailedOp
	}
	if _, exists := r.brands[brand.ID]; !exists {
		return ErrMockNotFound
	}
	r.brands[brand.ID] = cloneBrand(brand)
	return nil
}

// Delete implementa BrandRepository
func (r *MockBrandRepository) Delete(_ context.Context, id string, tenantID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.incrementCallCount("Delete")
	if r.shouldMethodFail("Delete") {
		return ErrMockFailedOp
	}
	brand, exists := r.brands[id]
	if !exists || brand.TenantID != tenantID {
		return ErrMockNotFound
	}
	delete(r.brands, id)
	return nil
}

// ExistsByName implementa BrandRepository
func (r *MockBrandRepository) ExistsByName(_ context.Context, name string, tenantID string, excludeID *string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.incrementCallCount("ExistsByName")
	if r.shouldMethodFail("ExistsByName") {
		return false, ErrMockFailedOp
	}
	for _, brand := range r.brands {
		if brand.Name == name && brand.TenantID == tenantID {
			if excludeID != nil && brand.ID == *excludeID {
				continue
			}
			return true, nil
		}
	}
	return false, nil
}

// SearchByCriteria implementa CriteriaRepository
func (r *MockBrandRepository) SearchByCriteria(_ context.Context, _ cr.Criteria) ([]*entity.Brand, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.incrementCallCount("SearchByCriteria")
	if r.shouldMethodFail("SearchByCriteria") {
		return nil, ErrMockFailedOp
	}
	result := make([]*entity.Brand, 0, len(r.brands))
	for _, brand := range r.brands {
		result = append(result, cloneBrand(brand))
	}
	return result, nil
}

// CountByCriteria implementa CriteriaRepository
func (r *MockBrandRepository) CountByCriteria(_ context.Context, _ cr.Criteria) (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.incrementCallCount("CountByCriteria")
	if r.shouldMethodFail("CountByCriteria") {
		return 0, ErrMockFailedOp
	}
	return len(r.brands), nil
}

// ListByCriteria implementa ListRepository
func (r *MockBrandRepository) ListByCriteria(_ context.Context, crit cr.Criteria) (*cr.ListResponse[entity.Brand], error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.incrementCallCount("ListByCriteria")
	if r.shouldMethodFail("ListByCriteria") {
		return nil, ErrMockFailedOp
	}
	result := make([]*entity.Brand, 0, len(r.brands))
	for _, brand := range r.brands {
		result = append(result, cloneBrand(brand))
	}
	return cr.NewListResponseFromCriteria(result, len(result), crit), nil
}
