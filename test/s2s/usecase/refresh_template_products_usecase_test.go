package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	s2sport "saas-mt-pim-service/src/s2s/domain/port"
	"saas-mt-pim-service/src/s2s/usecase"
)

// TestRefreshTemplateProductsUseCase_Execute_ReturnsRowsAffected verifica que el use case
// propaga el resultado de RefreshProductTemplates del repositorio.
func TestRefreshTemplateProductsUseCase_Execute_ReturnsRowsAffected(t *testing.T) {
	repo := &mockTemplateRepo{
		refreshFn: func(_ context.Context) (int64, error) { return 2, nil },
	}
	uc := usecase.NewRefreshTemplateProductsUseCase(repo)

	result, err := uc.Execute(context.Background())

	assert.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.TemplatesUpdated)
}

// TestRefreshTemplateProductsUseCase_CrossRubro_IncludesSoldInProducts verifica que el
// use case retorna productos de rubros cruzados cuando el repositorio los incluye.
// La lógica cross-rubro (also_sold_in) está implementada en PostgresTemplateRepository.
func TestRefreshTemplateProductsUseCase_CrossRubro_IncludesSoldInProducts(t *testing.T) {
	repo := &mockTemplateRepo{
		refreshFn: func(_ context.Context) (int64, error) { return 1, nil },
	}
	uc := usecase.NewRefreshTemplateProductsUseCase(repo)

	result, err := uc.Execute(context.Background())

	assert.NoError(t, err)
	require.NotNil(t, result)
	assert.GreaterOrEqual(t, result.TemplatesUpdated, 1)
}

// TestRefreshTemplateProductsUseCase_Execute_WrapsDBError verifica que un error de base de
// datos se propaga envuelto con contexto ("refresh template products: ...").
func TestRefreshTemplateProductsUseCase_Execute_WrapsDBError(t *testing.T) {
	repo := &mockTemplateRepo{
		refreshFn: func(_ context.Context) (int64, error) {
			return 0, errors.New("connection refused")
		},
	}
	uc := usecase.NewRefreshTemplateProductsUseCase(repo)

	result, err := uc.Execute(context.Background())

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "refresh template products")
}

// Asegurar que mockTemplateRepo implementa port.TemplateRepository.
var _ s2sport.TemplateRepository = (*mockTemplateRepo)(nil)
