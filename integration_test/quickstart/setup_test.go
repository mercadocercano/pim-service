//go:build integration

// Package quickstart_test contiene tests de integración para los endpoints de business-types
// y business-type-templates del módulo quickstart.
// Usa TestContainers para levantar un PostgreSQL real y verificar el comportamiento completo
// del stack HTTP → Handler → UseCase → Repository → DB.
//
// Ejecutar con: go test -tags=integration ./integration_test/quickstart/... -v
package quickstart_test

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

	btUsecase "saas-mt-pim-service/src/businesstype/application/usecase"
	btController "saas-mt-pim-service/src/businesstype/infrastructure/controller"
	btRepo "saas-mt-pim-service/src/businesstype/infrastructure/persistence/repository"
)

// testServer agrupa el servidor HTTP y la DB para los tests de integración.
type testServer struct {
	Server *httptest.Server
	DB     *sql.DB
}

// newTestServer levanta un contenedor PostgreSQL, ejecuta migraciones
// y retorna un httptest.Server con los módulos business-types y templates configurados.
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
	setupQuickstartRoutes(apiV1, db)

	srv := httptest.NewServer(router)
	t.Cleanup(srv.Close)

	return &testServer{Server: srv, DB: db}
}

// setupQuickstartRoutes configura las rutas de business-types y templates en el router de test
func setupQuickstartRoutes(router *gin.RouterGroup, db *sql.DB) {
	// Repositorios
	businessTypeRepo := btRepo.NewBusinessTypePostgresRepository(db)
	templateRepo := btRepo.NewBusinessTypeTemplatePostgresRepository(db)
	analyticsRepo := btRepo.NewTemplateAnalyticsPostgresRepository(db)

	// Use cases de business types
	createBTUseCase := btUsecase.NewCreateBusinessTypeUseCase(businessTypeRepo)
	listBTUseCase := btUsecase.NewListBusinessTypesUseCase(businessTypeRepo)
	getBTUseCase := btUsecase.NewGetBusinessTypeUseCase(businessTypeRepo)
	updateBTUseCase := btUsecase.NewUpdateBusinessTypeUseCase(businessTypeRepo)

	// Handler de business types (delete, activate, deactivate se crean internamente)
	btHandler := btController.NewBusinessTypeHandler(
		createBTUseCase,
		listBTUseCase,
		getBTUseCase,
		updateBTUseCase,
		businessTypeRepo,
	)
	btHandler.RegisterRoutes(router)

	// Use cases de templates
	createTplUseCase := btUsecase.NewCreateBusinessTypeTemplateUseCase(templateRepo, businessTypeRepo)
	updateTplUseCase := btUsecase.NewUpdateBusinessTypeTemplateUseCase(templateRepo)
	listTplUseCase := btUsecase.NewListBusinessTypeTemplatesUseCase(templateRepo)
	getTplUseCase := btUsecase.NewGetBusinessTypeTemplateUseCase(templateRepo)
	deleteTplUseCase := btUsecase.NewDeleteBusinessTypeTemplateUseCase(templateRepo)
	analyticsUseCase := btUsecase.NewGetTemplateAnalyticsUseCase(templateRepo, analyticsRepo)
	duplicateUseCase := btUsecase.NewDuplicateTemplateUseCase(templateRepo)

	tplHandler := btController.NewBusinessTypeTemplateHandler(
		createTplUseCase,
		updateTplUseCase,
		listTplUseCase,
		getTplUseCase,
		deleteTplUseCase,
	).WithAnalyticsUseCase(analyticsUseCase).
		WithDuplicateUseCase(duplicateUseCase)

	tplHandler.RegisterRoutes(router)
}

// enableExtensions habilita extensiones de PostgreSQL necesarias
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
			t.Logf("warn: migration %s may have failed (possibly idempotent): %v", filepath.Base(f), err)
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
