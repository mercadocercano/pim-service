//go:build integration

package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/quickstart/infrastructure/persistence/repository"
)

func testDB(t *testing.T) *sql.DB {
	t.Helper()
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost port=5432 user=postgres password=postgres dbname=pim_test sslmode=disable"
	}
	db, err := sql.Open("postgres", dsn)
	require.NoError(t, err)
	require.NoError(t, db.Ping())
	return db
}

// insertBusinessTypeWithTemplate inserta un business_type, business_type_template activo/default,
// y retorna los IDs creados para cleanup posterior.
func insertBusinessTypeWithTemplate(t *testing.T, db *sql.DB, slug string, productsJSON string) (btID, bttID string) {
	t.Helper()

	err := db.QueryRow(`
		INSERT INTO business_types (name, code, description, icon, is_active, created_at, updated_at)
		VALUES ($1, $2, 'test', 'test', true, NOW(), NOW())
		RETURNING id`,
		fmt.Sprintf("Test-%s-%d", slug, time.Now().UnixNano()),
		fmt.Sprintf("%s-%d", slug, time.Now().UnixNano()),
	).Scan(&btID)
	require.NoError(t, err)

	err = db.QueryRow(`
		INSERT INTO business_type_templates (business_type_id, name, is_active, is_default, products, categories, region, created_at, updated_at)
		VALUES ($1, $2, true, true, $3::jsonb, '[]'::jsonb, 'AR', NOW(), NOW())
		RETURNING id`,
		btID, fmt.Sprintf("template-%s", slug), productsJSON,
	).Scan(&bttID)
	require.NoError(t, err)

	return btID, bttID
}

func insertGlobalProduct(t *testing.T, db *sql.DB, name, brand, category, businessType string, qualityScore float64) string {
	t.Helper()
	var id string
	err := db.QueryRow(`
		INSERT INTO global_products (name, brand, category, business_type, quality_score, is_active, is_verified, source_reliability, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, true, true, 1.0, NOW(), NOW())
		RETURNING id`,
		name, brand, category, businessType, qualityScore,
	).Scan(&id)
	require.NoError(t, err)
	return id
}

func cleanupIntegration(t *testing.T, db *sql.DB, btID, bttID string, globalProductIDs []string) {
	t.Helper()
	db.Exec(`DELETE FROM business_type_product_templates WHERE business_type_template_id = $1`, bttID)
	db.Exec(`DELETE FROM business_type_templates WHERE id = $1`, bttID)
	db.Exec(`DELETE FROM business_types WHERE id = $1`, btID)
	for _, gpID := range globalProductIDs {
		db.Exec(`DELETE FROM global_products WHERE id = $1`, gpID)
	}
}

// TestGetProductsByBusinessType_Integration_ComputedAvailable verifica que cuando
// business_type_product_templates tiene UUIDs válidos, el repositorio los resuelve
// desde global_products y retorna los productos computados.
func TestGetProductsByBusinessType_Integration_ComputedAvailable(t *testing.T) {
	db := testDB(t)
	defer db.Close()

	// Arrange: crear datos de test
	slug := fmt.Sprintf("test-almacen-%d", time.Now().UnixNano())
	btID, bttID := insertBusinessTypeWithTemplate(t, db, slug, `[]`)
	gpID := insertGlobalProduct(t, db, "Harina 000 Test", "Arcor", "harinas", slug, 0.9)
	defer cleanupIntegration(t, db, btID, bttID, []string{gpID})

	// Insertar en business_type_product_templates con el UUID del global_product
	_, err := db.Exec(`
		INSERT INTO business_type_product_templates (business_type_template_id, suggested_products, priority_brands, max_products_per_category, updated_at)
		VALUES ($1, $2::jsonb, '[]'::jsonb, 30, NOW())`,
		bttID, fmt.Sprintf(`["%s"]`, gpID),
	)
	require.NoError(t, err)

	repo := repository.NewGetProductsByBusinessTypePostgresRepository(db)

	// Act: usar el code real del business_type que se insertó
	var btCode string
	err = db.QueryRow(`SELECT code FROM business_types WHERE id = $1`, btID).Scan(&btCode)
	require.NoError(t, err)

	products, err := repo.GetProductsByBusinessType(context.Background(), btCode)

	// Assert
	require.NoError(t, err)
	require.Len(t, products, 1)
	assert.Equal(t, "Harina 000 Test", products[0].Name)
	assert.Equal(t, "Arcor", products[0].Brand)
}

// TestGetProductsByBusinessType_Integration_FallbackEditorial verifica que cuando
// no hay registro en business_type_product_templates, se usa el JSONB editorial.
func TestGetProductsByBusinessType_Integration_FallbackEditorial(t *testing.T) {
	db := testDB(t)
	defer db.Close()

	// Arrange: template con producto editorial, sin computed
	slug := fmt.Sprintf("test-editorial-%d", time.Now().UnixNano())
	editorialJSON := `[{"name":"ProductoEditorial","brand":"MarcaTest","category_slug":"test"}]`
	btID, bttID := insertBusinessTypeWithTemplate(t, db, slug, editorialJSON)
	defer cleanupIntegration(t, db, btID, bttID, nil)

	repo := repository.NewGetProductsByBusinessTypePostgresRepository(db)

	var btCode string
	err := db.QueryRow(`SELECT code FROM business_types WHERE id = $1`, btID).Scan(&btCode)
	require.NoError(t, err)

	// Act
	products, err := repo.GetProductsByBusinessType(context.Background(), btCode)

	// Assert
	require.NoError(t, err)
	require.Len(t, products, 1)
	assert.Equal(t, "ProductoEditorial", products[0].Name)
	assert.Equal(t, "MarcaTest", products[0].Brand)
}
