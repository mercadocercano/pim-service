//go:build integration

// Package attributes_test contiene tests de integración para los endpoints de atributos marketplace.
// Usa TestContainers para levantar un PostgreSQL real y verificar el comportamiento completo
// del stack HTTP → Handler → UseCase → Repository → DB.
//
// Ejecutar con: go test -tags=integration ./integration_test/attributes/... -v
package attributes_test

import (
	"context"
	"database/sql"
	"fmt"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	attrUsecase "saas-mt-pim-service/src/attribute/application/usecase"
	attrCriteria "saas-mt-pim-service/src/attribute/infrastructure/criteria"
	attrController "saas-mt-pim-service/src/attribute/infrastructure/controller"
	attrRepo "saas-mt-pim-service/src/attribute/infrastructure/persistence/repository"
)

// testServer agrupa el servidor HTTP y la DB para los tests de integración.
type testServer struct {
	Server *httptest.Server
	DB     *sql.DB
}

// newTestServer levanta un contenedor PostgreSQL, ejecuta migraciones
// y retorna un httptest.Server con el módulo marketplace-attributes configurado.
func newTestServer(t *testing.T) *testServer {
	t.Helper()

	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("pim_test"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
		),
	)
	if err != nil {
		t.Fatalf("error starting postgres container: %v", err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Logf("warn: failed to terminate postgres container: %v", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("error getting connection string: %v", err)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}
	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("error pinging database: %v", err)
	}

	enableExtensions(t, db)
	runMigrations(t, db)
	ensureVariantAttributesTables(t, db)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(gin.Recovery())

	apiV1 := router.Group("/api/v1")
	setupAttributeRoutes(apiV1, db)

	srv := httptest.NewServer(router)
	t.Cleanup(srv.Close)

	return &testServer{Server: srv, DB: db}
}

// setupAttributeRoutes configura las rutas de atributos en el router de test
func setupAttributeRoutes(router *gin.RouterGroup, db *sql.DB) {
	marketplaceAttrRepo := attrRepo.NewMarketplaceAttributePostgresRepository(db)
	attrValueRepo := attrRepo.NewAttributeValuePostgresRepository(db)
	criteriaBuilder := attrCriteria.NewMarketplaceAttributeCriteriaBuilder()

	handler := attrController.NewMarketplaceAttributeHandler(
		attrUsecase.NewCreateMarketplaceAttributeUseCase(marketplaceAttrRepo),
		attrUsecase.NewListMarketplaceAttributesUseCase(marketplaceAttrRepo),
		attrUsecase.NewListMarketplaceAttributesByCriteriaUseCase(marketplaceAttrRepo),
		attrUsecase.NewGetMarketplaceAttributeByIDUseCase(marketplaceAttrRepo),
		attrUsecase.NewUpdateMarketplaceAttributeUseCase(marketplaceAttrRepo),
		attrUsecase.NewDeleteMarketplaceAttributeUseCase(marketplaceAttrRepo),
		criteriaBuilder,
	).WithValueUseCases(
		attrUsecase.NewListAttributeValuesUseCase(attrValueRepo),
		attrUsecase.NewCreateAttributeValueUseCase(attrValueRepo, marketplaceAttrRepo),
		attrUsecase.NewUpdateAttributeValueUseCase(attrValueRepo),
		attrUsecase.NewDeleteAttributeValueUseCase(attrValueRepo),
	)

	handler.RegisterRoutes(router)
}

// enableExtensions habilita las extensiones de PostgreSQL necesarias y crea tablas auxiliares
func enableExtensions(t *testing.T, db *sql.DB) {
	t.Helper()
	stmts := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`,
		`CREATE EXTENSION IF NOT EXISTS "unaccent"`,
		`CREATE TABLE IF NOT EXISTS schema_migrations (
			filename   TEXT PRIMARY KEY,
			applied_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		)`,
	}
	for _, stmt := range stmts {
		if _, err := db.Exec(stmt); err != nil {
			t.Fatalf("error enabling extension (%s): %v", stmt, err)
		}
	}
}

// ensureVariantAttributesTables crea las tablas auxiliares que pueden no haberse creado
// correctamente por limitaciones del driver al ejecutar migraciones multi-statement.
func ensureVariantAttributesTables(t *testing.T, db *sql.DB) {
	t.Helper()

	// tenant_custom_attributes es necesaria como FK de variant_marketplace_attributes
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tenant_custom_attributes (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			tenant_id UUID NOT NULL,
			marketplace_category_id UUID,
			name VARCHAR(255) NOT NULL,
			slug VARCHAR(255) NOT NULL,
			type VARCHAR(50) NOT NULL,
			is_filterable BOOLEAN DEFAULT FALSE,
			is_searchable BOOLEAN DEFAULT FALSE,
			validation_rules JSONB DEFAULT '{}',
			sort_order INTEGER DEFAULT 0,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			CONSTRAINT tca_type_check CHECK (type IN ('text', 'number', 'boolean', 'select', 'multi_select')),
			CONSTRAINT tca_name_not_empty CHECK (LENGTH(TRIM(name)) > 0),
			CONSTRAINT tca_slug_not_empty CHECK (LENGTH(TRIM(slug)) > 0)
		)
	`)
	if err != nil {
		t.Fatalf("error creating tenant_custom_attributes: %v", err)
	}

	// variant_marketplace_attributes es usada por IsInUse
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS variant_marketplace_attributes (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			variant_id UUID NOT NULL REFERENCES product_variants(id) ON DELETE CASCADE,
			marketplace_attribute_id UUID REFERENCES marketplace_attributes(id) ON DELETE CASCADE,
			tenant_custom_attribute_id UUID REFERENCES tenant_custom_attributes(id) ON DELETE CASCADE,
			value TEXT NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			CONSTRAINT vma_single_attribute CHECK (
				(marketplace_attribute_id IS NOT NULL AND tenant_custom_attribute_id IS NULL) OR
				(marketplace_attribute_id IS NULL AND tenant_custom_attribute_id IS NOT NULL)
			),
			UNIQUE(variant_id, marketplace_attribute_id),
			UNIQUE(variant_id, tenant_custom_attribute_id)
		)
	`)
	if err != nil {
		t.Fatalf("error creating variant_marketplace_attributes: %v", err)
	}
}

// runMigrations ejecuta los archivos SQL de migraciones en orden ascendente
func runMigrations(t *testing.T, db *sql.DB) {
	t.Helper()

	migrationsDir := findMigrationsDir(t)

	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		t.Fatalf("error reading migrations dir %s: %v", migrationsDir, err)
	}

	var sqlFiles []string
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".sql" {
			continue
		}
		name := entry.Name()
		if isDownMigration(name) || isSeedFile(name) {
			continue
		}
		sqlFiles = append(sqlFiles, filepath.Join(migrationsDir, name))
	}
	sort.Strings(sqlFiles)

	for _, f := range sqlFiles {
		content, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("error reading migration %s: %v", f, err)
		}
		if _, err := db.Exec(string(content)); err != nil {
			t.Fatalf("error executing migration %s: %v", f, err)
		}
	}
}

func isDownMigration(name string) bool {
	return len(name) > 9 && name[len(name)-9:] == ".down.sql"
}

func isSeedFile(name string) bool {
	return len(name) > 5 && name[:5] == "seed_"
}

// findMigrationsDir localiza el directorio migrations relativo a este archivo de test
func findMigrationsDir(t *testing.T) string {
	t.Helper()

	candidates := []string{
		"../../migrations",
		"../migrations",
		"migrations",
	}
	for _, c := range candidates {
		if info, err := os.Stat(c); err == nil && info.IsDir() {
			return c
		}
	}
	t.Fatalf("could not find migrations directory")
	return ""
}

// baseURL construye la URL base del servidor de test
func baseURL(srv *testServer) string {
	return fmt.Sprintf("%s/api/v1", srv.Server.URL)
}
