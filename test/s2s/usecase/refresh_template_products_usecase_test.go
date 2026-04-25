package usecase_test

import (
	"context"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/s2s/usecase"
)

// TestRefreshTemplateProductsUseCase_Execute_FiltersProductsByBusinessType verifica que el
// query incluye el JOIN `gp.business_type = tc.business_type_code`, la condición que evita
// que productos de un rubro (ej. kiosco) aparezcan en templates de otro (ej. almacen)
// aunque compartan el mismo slug de categoría.
func TestRefreshTemplateProductsUseCase_Execute_FiltersProductsByBusinessType(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`gp\.business_type = tc\.business_type_code`).
		WillReturnResult(sqlmock.NewResult(0, 2))

	uc := usecase.NewRefreshTemplateProductsUseCase(db)

	// Act
	result, err := uc.Execute(context.Background())

	// Assert
	assert.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.TemplatesUpdated)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestRefreshTemplateProductsUseCase_CrossRubro_IncludesSoldInProducts verifica que el
// query incluye la condición also_sold_in para permitir que productos de un rubro (ej.
// vinoteca) aparezcan en el quickstart de otro rubro (ej. almacen) cuando el campo
// also_sold_in del producto contiene el business_type_code del template destino.
//
// ESTADO ESPERADO: ROJO (TDD red). El UseCase todavía no implementa lógica cross-rubro.
// Este test debe fallar hasta que se agregue la condición `also_sold_in` al query SQL.
func TestRefreshTemplateProductsUseCase_CrossRubro_IncludesSoldInProducts(t *testing.T) {
	// Arrange — mock que espera un query con lógica also_sold_in (cross-rubro)
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// El query correcto (post-implementación) deberá unir productos vía also_sold_in.
	// Regexp: buscamos que el query mencione also_sold_in para el JOIN cross-rubro.
	mock.ExpectExec(`also_sold_in`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	uc := usecase.NewRefreshTemplateProductsUseCase(db)

	// Act
	result, err := uc.Execute(context.Background())

	// Assert — cuando el UseCase implemente cross-rubro, esto debe pasar:
	// - no error
	// - al menos 1 template actualizado (el de almacen con producto de vinoteca)
	// - todas las expectativas del mock satisfechas
	assert.NoError(t, err)
	require.NotNil(t, result)
	assert.GreaterOrEqual(t, result.TemplatesUpdated, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestRefreshTemplateProductsUseCase_Execute_WrapsDBError verifica que un error de base de
// datos se propaga envuelto con contexto ("refresh template products: ...").
func TestRefreshTemplateProductsUseCase_Execute_WrapsDBError(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`.*`).WillReturnError(fmt.Errorf("connection refused"))

	uc := usecase.NewRefreshTemplateProductsUseCase(db)

	// Act
	result, err := uc.Execute(context.Background())

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "refresh template products")
}
