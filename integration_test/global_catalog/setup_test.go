//go:build integration

// Package global_catalog_test contiene tests de integración para los endpoints del catálogo global.
// Usa TestContainers para levantar un PostgreSQL real y verificar el comportamiento completo
// del stack HTTP → Handler → Repository → DB.
//
// Ejecutar con: go test -tags=integration ./integration_test/global_catalog/... -v
package global_catalog_test

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

	pimport "saas-mt-pim-service/src/pim/domain/port"
	gcUsecase "saas-mt-pim-service/src/product/global_catalog/application/usecase"
	gcController "saas-mt-pim-service/src/product/global_catalog/infrastructure/controller"
	gcCriteria "saas-mt-pim-service/src/product/global_catalog/infrastructure/criteria"
	gcPersistence "saas-mt-pim-service/src/product/global_catalog/infrastructure/persistence"
)

// testServer agrupa el servidor HTTP y la DB para los tests de integración.
type testServer struct {
	Server *httptest.Server
	DB     *sql.DB
}

// newTestServer levanta un contenedor PostgreSQL, ejecuta migraciones
// y retorna un httptest.Server con el módulo global-catalog configurado.
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

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(gin.Recovery())

	apiV1 := router.Group("/api/v1")
	repo := gcPersistence.NewPostgresGlobalProductRepository(db)
	bulkVerifyRepo := gcPersistence.NewPostgresBulkVerifyRepository(db)
	noopLog := testLogger{}

	deps := gcController.GlobalCatalogControllerDeps{
		CreateGlobalProduct:           gcUsecase.NewCreateGlobalProduct(repo),
		SearchByEAN:                   gcUsecase.NewSearchByEAN(repo),
		ListGlobalProducts:            gcUsecase.NewListGlobalProducts(repo),
		ListGlobalProductsByCriteria:  gcUsecase.NewListGlobalProductsByCriteriaUseCase(repo),
		GetGlobalProductByID:          gcUsecase.NewGetGlobalProductByID(repo),
		UpdateGlobalProductByID:       gcUsecase.NewUpdateGlobalProductByID(repo),
		DeleteGlobalProduct:           gcUsecase.NewDeleteGlobalProduct(repo),
		VerifyGlobalProduct:           gcUsecase.NewVerifyGlobalProduct(repo),
		UnverifyGlobalProduct:         gcUsecase.NewUnverifyGlobalProduct(repo),
		BulkImportGlobalProducts:      gcUsecase.NewBulkImportGlobalProducts(repo),
		GetBusinessTypeFacets:         gcUsecase.NewGetBusinessTypeFacets(repo),
		ListProductsNeedingEnrichment: gcUsecase.NewListProductsNeedingEnrichment(repo),
		GetGlobalProductsByIDs:        gcUsecase.NewGetGlobalProductsByIDs(repo),
		GetDistinctBusinessTypes:      gcUsecase.NewGetDistinctBusinessTypes(repo),
		BulkVerifyGlobalProducts:      gcUsecase.NewBulkVerifyGlobalProducts(repo, bulkVerifyRepo, noopLog),
		CriteriaBuilder:               gcCriteria.NewGlobalProductCriteriaBuilder(),
	}

	handler := gcController.NewGlobalCatalogControllerWithDeps(deps)
	handler.RegisterRoutes(apiV1)

	srv := httptest.NewServer(router)
	t.Cleanup(srv.Close)

	return &testServer{Server: srv, DB: db}
}

// testLogger implementa pimport.PIMEventLogger sin efectos colaterales.
type testLogger struct{}

func (testLogger) Log(e pimport.PIMEvent) {}


// enableExtensions habilita las extensiones de PostgreSQL y crea tablas de infraestructura.
func enableExtensions(t *testing.T, db *sql.DB) {
	t.Helper()
	stmts := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`,
		`CREATE EXTENSION IF NOT EXISTS "unaccent"`,
		// Tabla requerida por algunas migraciones que usan el migration runner
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

// runMigrations ejecuta los archivos SQL de migraciones en orden.
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

// findMigrationsDir localiza el directorio migrations relativo a este archivo de test.
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

// baseURL construye la URL base del servidor de test.
func baseURL(srv *testServer) string {
	return fmt.Sprintf("%s/api/v1", srv.Server.URL)
}
