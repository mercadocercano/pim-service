package repository_test

import (
	"context"
	"encoding/json"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/quickstart/domain/port"
	"saas-mt-pim-service/src/quickstart/infrastructure/persistence/repository"
)

// TestGetProductsByBusinessType_ComputedAvailable_ReturnsComputedProducts
// verifica que cuando business_type_product_template_details tiene product_ids,
// se resuelven desde global_products (fuente computed).
func TestGetProductsByBusinessType_ComputedAvailable_ReturnsComputedProducts(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	productID1 := "11111111-1111-1111-1111-111111111111"
	productID2 := "22222222-2222-2222-2222-222222222222"

	// Primer query: fetchComputedTemplateIDs retorna 2 UUIDs
	computedRows := sqlmock.NewRows([]string{"product_id"}).
		AddRow(productID1).
		AddRow(productID2)
	mock.ExpectQuery(`business_type_product_template_details`).
		WithArgs("almacen").
		WillReturnRows(computedRows)

	// Segundo query: resolveGlobalProducts retorna 2 productos
	globalRows := sqlmock.NewRows([]string{"name", "brand", "category", "quality_score"}).
		AddRow("Harina 000", "Arcor", "harinas", 0.9).
		AddRow("Azúcar 1kg", "Ledesma", "azucar", 0.85)
	mock.ExpectQuery(`global_products`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(globalRows)

	repo := repository.NewGetProductsByBusinessTypePostgresRepository(db)

	// Act
	products, err := repo.GetProductsByBusinessType(context.Background(), "almacen")

	// Assert
	require.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Harina 000", products[0].Name)
	assert.Equal(t, "Arcor", products[0].Brand)
	assert.Equal(t, "harinas", products[0].CategorySlug)
	assert.Equal(t, "Azúcar 1kg", products[1].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestGetProductsByBusinessType_NoComputed_FallsBackToEditorial
// verifica que cuando no hay detalle computado (ErrNoRows),
// se consulta el JSONB editorial y se devuelven esos productos.
func TestGetProductsByBusinessType_NoComputed_FallsBackToEditorial(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Primer query: fetchComputedTemplateIDs retorna 0 filas (sin detalle computado)
	computedRows := sqlmock.NewRows([]string{"product_id"})
	mock.ExpectQuery(`business_type_product_template_details`).
		WithArgs("relojeria").
		WillReturnRows(computedRows)

	// Segundo query: fetchEditorialProducts retorna JSONB editorial
	editorialProducts := []port.TemplateProduct{
		{Name: "Reloj Casio", Brand: "Casio", CategorySlug: "relojes"},
		{Name: "Pila AA", Brand: "Duracell", CategorySlug: "pilas"},
		{Name: "Correa Cuero", Brand: "Genérico", CategorySlug: "accesorios"},
	}
	editorialJSON, _ := json.Marshal(editorialProducts)
	editorialRows := sqlmock.NewRows([]string{"products"}).AddRow(editorialJSON)
	mock.ExpectQuery(`business_type_templates`).
		WithArgs("relojeria").
		WillReturnRows(editorialRows)

	repo := repository.NewGetProductsByBusinessTypePostgresRepository(db)

	// Act
	products, err := repo.GetProductsByBusinessType(context.Background(), "relojeria")

	// Assert
	require.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Reloj Casio", products[0].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestGetProductsByBusinessType_EmptyComputed_FallsBackToEditorial
// verifica que cuando el detalle computado no retorna filas, se usa el fallback editorial.
func TestGetProductsByBusinessType_EmptyComputed_FallsBackToEditorial(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Primer query: fetchComputedTemplateIDs retorna 0 filas
	computedRows := sqlmock.NewRows([]string{"product_id"})
	mock.ExpectQuery(`business_type_product_template_details`).
		WithArgs("bazar").
		WillReturnRows(computedRows)

	// Segundo query: fetchEditorialProducts retorna JSONB editorial
	editorialProducts := []port.TemplateProduct{
		{Name: "Jarro Esmaltado", Brand: "Tramontina", CategorySlug: "cocina"},
	}
	editorialJSON, _ := json.Marshal(editorialProducts)
	editorialRows := sqlmock.NewRows([]string{"products"}).AddRow(editorialJSON)
	mock.ExpectQuery(`business_type_templates`).
		WithArgs("bazar").
		WillReturnRows(editorialRows)

	repo := repository.NewGetProductsByBusinessTypePostgresRepository(db)

	// Act
	products, err := repo.GetProductsByBusinessType(context.Background(), "bazar")

	// Assert
	require.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, "Jarro Esmaltado", products[0].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}
