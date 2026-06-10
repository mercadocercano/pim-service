package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	s2sport "saas-mt-pim-service/src/s2s/domain/port"
	"saas-mt-pim-service/src/s2s/usecase"
)

// mockTemplateRepo implementa port.TemplateRepository para tests de usecase.
type mockTemplateRepo struct {
	getStatusFn func(ctx context.Context, slug string) (*s2sport.TemplateStatusRow, error)
	refreshFn   func(ctx context.Context) (int64, error)
}

func (m *mockTemplateRepo) GetTemplateStatus(ctx context.Context, slug string) (*s2sport.TemplateStatusRow, error) {
	return m.getStatusFn(ctx, slug)
}

func (m *mockTemplateRepo) RefreshProductTemplates(ctx context.Context) (int64, error) {
	return m.refreshFn(ctx)
}

// TestGetTemplateStatusUseCase_ComputedExists_ReturnsComputedSource verifica que cuando
// computed_count > 0 el resultado tiene source="computed" y LastRefresh no nulo.
func TestGetTemplateStatusUseCase_ComputedExists_ReturnsComputedSource(t *testing.T) {
	now := time.Now().UTC().Truncate(time.Second)
	repo := &mockTemplateRepo{
		getStatusFn: func(_ context.Context, _ string) (*s2sport.TemplateStatusRow, error) {
			return &s2sport.TemplateStatusRow{ComputedCount: 28, EditorialCount: 20, LastRefresh: &now}, nil
		},
	}
	uc := usecase.NewGetTemplateStatusUseCase(repo)

	result, err := uc.Execute(context.Background(), "almacen")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "computed", result.Source)
	assert.Equal(t, 28, result.ComputedCount)
	assert.Equal(t, 20, result.EditorialCount)
	assert.NotNil(t, result.LastRefresh)
	assert.Equal(t, "almacen", result.BusinessTypeSlug)
}

// TestGetTemplateStatusUseCase_NoComputed_ReturnsEditorialSource verifica que cuando
// computed_count = 0, source="editorial" y LastRefresh es nil.
func TestGetTemplateStatusUseCase_NoComputed_ReturnsEditorialSource(t *testing.T) {
	repo := &mockTemplateRepo{
		getStatusFn: func(_ context.Context, _ string) (*s2sport.TemplateStatusRow, error) {
			return &s2sport.TemplateStatusRow{ComputedCount: 0, EditorialCount: 15, LastRefresh: nil}, nil
		},
	}
	uc := usecase.NewGetTemplateStatusUseCase(repo)

	result, err := uc.Execute(context.Background(), "relojeria")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "editorial", result.Source)
	assert.Equal(t, 0, result.ComputedCount)
	assert.Equal(t, 15, result.EditorialCount)
	assert.Nil(t, result.LastRefresh)
}

// TestGetTemplateStatusUseCase_UnknownSlug_ReturnsNotFoundError verifica que un slug
// inexistente (repo retorna nil, nil) resulta en ErrBusinessTypeNotFound.
func TestGetTemplateStatusUseCase_UnknownSlug_ReturnsNotFoundError(t *testing.T) {
	repo := &mockTemplateRepo{
		getStatusFn: func(_ context.Context, _ string) (*s2sport.TemplateStatusRow, error) {
			return nil, nil // nil, nil = slug no encontrado (igual que PostgresTemplateRepository)
		},
	}
	uc := usecase.NewGetTemplateStatusUseCase(repo)

	result, err := uc.Execute(context.Background(), "rubros-inexistente")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.True(t, errors.Is(err, usecase.ErrBusinessTypeNotFound))
}

// TestGetTemplateStatusUseCase_RepoError_WrapsError verifica que un error del repositorio
// se propaga envuelto con contexto.
func TestGetTemplateStatusUseCase_RepoError_WrapsError(t *testing.T) {
	repoErr := errors.New("connection refused")
	repo := &mockTemplateRepo{
		getStatusFn: func(_ context.Context, _ string) (*s2sport.TemplateStatusRow, error) {
			return nil, repoErr
		},
	}
	uc := usecase.NewGetTemplateStatusUseCase(repo)

	result, err := uc.Execute(context.Background(), "almacen")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "get template status")
}
