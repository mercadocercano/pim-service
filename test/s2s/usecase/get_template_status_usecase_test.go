package usecase_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/s2s/usecase"
)

// TestGetTemplateStatusUseCase_ComputedExists_ReturnsComputedSource verifica que cuando
// computed_count > 0 el resultado tiene source="computed" y LastRefresh no nulo.
func TestGetTemplateStatusUseCase_ComputedExists_ReturnsComputedSource(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	now := time.Now().UTC().Truncate(time.Second)
	rows := sqlmock.NewRows([]string{"computed_count", "editorial_count", "last_refresh"}).
		AddRow(28, 20, sql.NullTime{Time: now, Valid: true})
	mock.ExpectQuery(`business_type_product_templates`).
		WithArgs("almacen").
		WillReturnRows(rows)

	uc := usecase.NewGetTemplateStatusUseCase(db)

	// Act
	result, err := uc.Execute(context.Background(), "almacen")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "computed", result.Source)
	assert.Equal(t, 28, result.ComputedCount)
	assert.Equal(t, 20, result.EditorialCount)
	assert.NotNil(t, result.LastRefresh)
	assert.Equal(t, "almacen", result.BusinessTypeSlug)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestGetTemplateStatusUseCase_NoComputed_ReturnsEditorialSource verifica que cuando
// computed_count = 0, source="editorial" y LastRefresh es nil.
func TestGetTemplateStatusUseCase_NoComputed_ReturnsEditorialSource(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"computed_count", "editorial_count", "last_refresh"}).
		AddRow(0, 15, sql.NullTime{Valid: false})
	mock.ExpectQuery(`business_type_product_templates`).
		WithArgs("relojeria").
		WillReturnRows(rows)

	uc := usecase.NewGetTemplateStatusUseCase(db)

	// Act
	result, err := uc.Execute(context.Background(), "relojeria")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "editorial", result.Source)
	assert.Equal(t, 0, result.ComputedCount)
	assert.Equal(t, 15, result.EditorialCount)
	assert.Nil(t, result.LastRefresh)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestGetTemplateStatusUseCase_UnknownSlug_ReturnsNotFoundError verifica que un slug
// inexistente retorna ErrBusinessTypeNotFound.
func TestGetTemplateStatusUseCase_UnknownSlug_ReturnsNotFoundError(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`business_type_product_templates`).
		WithArgs("rubros-inexistente").
		WillReturnError(sql.ErrNoRows)

	uc := usecase.NewGetTemplateStatusUseCase(db)

	// Act
	result, err := uc.Execute(context.Background(), "rubros-inexistente")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, usecase.ErrBusinessTypeNotFound)
	assert.NoError(t, mock.ExpectationsWereMet())
}
