package repository

import (
	"context"
	"errors"
	"sync"

	"pim/src/category/domain/entity"
)

// Errores mock
var (
	ErrMockFailedOp   = errors.New("operación fallida (simulada)")
	ErrMockNotFound   = errors.New("recurso no encontrado (simulado)")
	ErrMockDuplicated = errors.New("recurso duplicado (simulado)")
)

// MockCategoryRepository implementa un repositorio en memoria para pruebas
type MockCategoryRepository struct {
	mu            sync.RWMutex
	categories    map[string]*entity.Category
	shouldFail    bool
	failOnMethods map[string]bool
	callHistory   map[string]int
}

// NewMockCategoryRepository crea una nueva instancia del mock
func NewMockCategoryRepository() *MockCategoryRepository {
	return &MockCategoryRepository{
		categories:    make(map[string]*entity.Category),
		failOnMethods: make(map[string]bool),
		callHistory:   make(map[string]int),
	}
}

// SetShouldFail configura si todas las operaciones deberían fallar
func (r *MockCategoryRepository) SetShouldFail(shouldFail bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.shouldFail = shouldFail
}

// ShouldFailOn configura un método específico para que falle
func (r *MockCategoryRepository) ShouldFailOn(method string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.failOnMethods[method] = true
}

// ResetFailures limpia todas las configuraciones de fallo
func (r *MockCategoryRepository) ResetFailures() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.shouldFail = false
	r.failOnMethods = make(map[string]bool)
}

// ResetCallHistory reinicia los contadores de llamadas
func (r *MockCategoryRepository) ResetCallHistory() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.callHistory = make(map[string]int)
}

// GetCallCount retorna el número de veces que se ha llamado a un método
func (r *MockCategoryRepository) GetCallCount(method string) int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.callHistory[method]
}

// SetupCategories inicializa el repositorio con categorías predefinidas
func (r *MockCategoryRepository) SetupCategories(categories []*entity.Category) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.categories = make(map[string]*entity.Category)
	for _, category := range categories {
		// Clonar para evitar referencias compartidas
		r.categories[category.ID] = &entity.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ParentID:    category.ParentID,
			Status:      category.Status,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		}
	}
}

// GetCategories retorna todas las categorías almacenadas
func (r *MockCategoryRepository) GetCategories() []*entity.Category {
	r.mu.RLock()
	defer r.mu.RUnlock()

	categories := make([]*entity.Category, 0, len(r.categories))
	for _, category := range r.categories {
		categories = append(categories, category)
	}
	return categories
}

// shouldMethodFail comprueba si un método debería fallar
func (r *MockCategoryRepository) shouldMethodFail(method string) bool {
	return r.shouldFail || r.failOnMethods[method]
}

// incrementCallCount incrementa el contador de llamadas para un método
func (r *MockCategoryRepository) incrementCallCount(method string) {
	r.callHistory[method] = r.callHistory[method] + 1
}

// Create implementa la interfaz del repositorio
func (r *MockCategoryRepository) Create(_ context.Context, category *entity.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.incrementCallCount("Create")

	if r.shouldMethodFail("Create") {
		return ErrMockFailedOp
	}

	if _, exists := r.categories[category.ID]; exists {
		return ErrMockDuplicated
	}

	// Crear una copia para evitar referencia compartida
	r.categories[category.ID] = &entity.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
		Status:      category.Status,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	return nil
}

// FindByID implementa la interfaz del repositorio
func (r *MockCategoryRepository) FindByID(_ context.Context, id string) (*entity.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.incrementCallCount("FindByID")

	if r.shouldMethodFail("FindByID") {
		return nil, ErrMockFailedOp
	}

	category, exists := r.categories[id]
	if !exists {
		return nil, ErrMockNotFound
	}

	// Crear una copia para evitar referencia compartida
	return &entity.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
		Status:      category.Status,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

// FindAll implementa la interfaz del repositorio
func (r *MockCategoryRepository) FindAll(_ context.Context) ([]*entity.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.incrementCallCount("FindAll")

	if r.shouldMethodFail("FindAll") {
		return nil, ErrMockFailedOp
	}

	categories := make([]*entity.Category, 0, len(r.categories))
	for _, category := range r.categories {
		// Crear una copia para evitar referencia compartida
		categories = append(categories, &entity.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ParentID:    category.ParentID,
			Status:      category.Status,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		})
	}

	return categories, nil
}

// Update implementa la interfaz del repositorio
func (r *MockCategoryRepository) Update(_ context.Context, category *entity.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.incrementCallCount("Update")

	if r.shouldMethodFail("Update") {
		return ErrMockFailedOp
	}

	if _, exists := r.categories[category.ID]; !exists {
		return ErrMockNotFound
	}

	// Actualizar con una copia para evitar referencia compartida
	r.categories[category.ID] = &entity.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
		Status:      category.Status,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}

	return nil
}

// Delete implementa la interfaz del repositorio
func (r *MockCategoryRepository) Delete(_ context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.incrementCallCount("Delete")

	if r.shouldMethodFail("Delete") {
		return ErrMockFailedOp
	}

	if _, exists := r.categories[id]; !exists {
		return ErrMockNotFound
	}

	delete(r.categories, id)

	return nil
}
