//go:build integration

// Package persistence — test de integración de PostgresReclassifyRepository contra Postgres real.
//
// Ejecutar con:
//
//	PIM_TEST_DATABASE_URL="postgres://postgres:postgres@localhost:5432/pim_reclassify_test?sslmode=disable" \
//	  go test -tags=integration ./src/product/global_catalog/infrastructure/persistence/... -run Reclassify -v
//
// SEGURIDAD: este test MUTA global_products (CREATE TABLE snapshot + UPDATE). Por eso EXIGE una
// DB de test aislada vía PIM_TEST_DATABASE_URL y se NIEGA a correr si la URL apunta a pim_db
// (el catálogo global real). Si la env var no está, hace t.Skip y no rompe `go test ./...`.
package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// setupReclassifyDB abre la conexión a la DB de test, valida que no sea pim_db,
// y deja un schema mínimo limpio (global_products + audit) en cada test.
func setupReclassifyDB(t *testing.T) *sql.DB {
	t.Helper()

	dsn := os.Getenv("PIM_TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("PIM_TEST_DATABASE_URL no seteada — se omite el test de integración del reclassify repo")
	}

	// Guardrail anti-pim_db: nunca correr applies contra el catálogo real.
	if strings.Contains(dsn, "/pim_db") {
		t.Fatalf("REHÚSO correr: PIM_TEST_DATABASE_URL apunta a pim_db (catálogo global real). Usá pim_reclassify_test")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("ping db: %v", err)
	}

	resetSchema(t, db)
	return db
}

// resetSchema deja el schema mínimo que toca el repo en estado conocido, dropeando
// cualquier snapshot global_products_bkp_* que haya quedado de corridas previas.
func resetSchema(t *testing.T, db *sql.DB) {
	t.Helper()
	ctx := context.Background()

	dropOrphanSnapshots(t, db)

	stmts := []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`,
		`DROP TABLE IF EXISTS global_product_reclassification_audit`,
		`DROP TABLE IF EXISTS global_products`,
		`CREATE TABLE global_products (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			ean VARCHAR(13) NOT NULL UNIQUE,
			name VARCHAR(500) NOT NULL,
			category VARCHAR(200),
			source VARCHAR(50) NOT NULL,
			business_type VARCHAR(100),
			is_active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		)`,
		`ALTER TABLE global_products
			ADD CONSTRAINT uq_global_products_name_business_type UNIQUE (name, business_type)`,
		// Migración real del audit (mismo CREATE que la migración 20260617000002).
		`CREATE TABLE IF NOT EXISTS global_product_reclassification_audit (
			id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			operator_id     TEXT NOT NULL,
			executed_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			mode            TEXT NOT NULL CHECK (mode IN ('dry_run', 'applied')),
			scope           JSONB NOT NULL,
			snapshot_ref    TEXT,
			summary         JSONB NOT NULL,
			affected_count  INT NOT NULL DEFAULT 0
		)`,
	}
	for _, s := range stmts {
		if _, err := db.ExecContext(ctx, s); err != nil {
			t.Fatalf("reset schema (%.40s...): %v", s, err)
		}
	}
}

// dropOrphanSnapshots elimina tablas global_products_bkp_* dejadas por corridas anteriores.
func dropOrphanSnapshots(t *testing.T, db *sql.DB) {
	t.Helper()
	ctx := context.Background()
	rows, err := db.QueryContext(ctx,
		`SELECT tablename FROM pg_tables WHERE schemaname='public' AND tablename LIKE 'global_products_bkp_%'`)
	if err != nil {
		t.Fatalf("list orphan snapshots: %v", err)
	}
	var tables []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			rows.Close()
			t.Fatalf("scan snapshot name: %v", err)
		}
		tables = append(tables, name)
	}
	rows.Close()
	for _, name := range tables {
		if _, err := db.ExecContext(ctx, fmt.Sprintf(`DROP TABLE IF EXISTS %q`, name)); err != nil {
			t.Fatalf("drop orphan snapshot %s: %v", name, err)
		}
	}
}

// seedProduct inserta un producto controlado y devuelve su id.
func seedProduct(t *testing.T, db *sql.DB, ean, name, category, source string, bt *string) string {
	t.Helper()
	var id string
	err := db.QueryRow(
		`INSERT INTO global_products (ean, name, category, source, business_type)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id::text`,
		ean, name, category, source, bt,
	).Scan(&id)
	if err != nil {
		t.Fatalf("seed product %s: %v", name, err)
	}
	return id
}

func strptr(s string) *string { return &s }

func mustScope(t *testing.T, prefix string, max int) value_object.ReclassifyScope {
	t.Helper()
	scope, err := value_object.NewReclassifyScope(prefix, max)
	if err != nil {
		t.Fatalf("new scope: %v", err)
	}
	return scope
}

// --- CASO 1: CountCandidates / FetchCandidates filtran por source_prefix y respetan el cap ---

func TestReclassifyRepo_CountAndFetchCandidates(t *testing.T) {
	db := setupReclassifyDB(t)
	repo := NewPostgresReclassifyRepository(db)
	ctx := context.Background()

	// Candidatos (business_type NULL o 'almacen') con source scraper%
	seedProduct(t, db, "7790000000001", "Prod A", "Lacteos", "scraper_disco", nil)
	seedProduct(t, db, "7790000000002", "Prod B", "Limpieza", "scraper_carrefour", strptr("almacen"))
	// Candidato pero source NO matchea el prefijo
	seedProduct(t, db, "7790000000003", "Prod C", "Lacteos", "manual", nil)
	// NO candidato: business_type específico (fiambreria)
	seedProduct(t, db, "7790000000004", "Prod D", "Fiambres", "scraper_disco", strptr("fiambreria"))

	scope := mustScope(t, "scraper", 1000)

	// CountCandidates: solo A y B (NULL/almacen + source scraper%) = 2
	count, err := repo.CountCandidates(ctx, scope)
	if err != nil {
		t.Fatalf("CountCandidates: %v", err)
	}
	if count != 2 {
		t.Errorf("CountCandidates filtro source_prefix: esperado 2, obtenido %d", count)
	}

	// FetchCandidates: mismas 2 filas
	cands, err := repo.FetchCandidates(ctx, scope)
	if err != nil {
		t.Fatalf("FetchCandidates: %v", err)
	}
	if len(cands) != 2 {
		t.Fatalf("FetchCandidates: esperado 2, obtenido %d", len(cands))
	}

	// El cap (MaxRows) se respeta como LIMIT.
	capped := mustScope(t, "scraper", 1)
	cappedCands, err := repo.FetchCandidates(ctx, capped)
	if err != nil {
		t.Fatalf("FetchCandidates capped: %v", err)
	}
	if len(cappedCands) != 1 {
		t.Errorf("FetchCandidates respeta cap LIMIT: esperado 1, obtenido %d", len(cappedCands))
	}

	// Sin prefijo: candidatos A, B, C (3). D sigue fuera por business_type específico.
	noPrefix := mustScope(t, "", 1000)
	allCount, err := repo.CountCandidates(ctx, noPrefix)
	if err != nil {
		t.Fatalf("CountCandidates sin prefijo: %v", err)
	}
	if allCount != 3 {
		t.Errorf("CountCandidates sin prefijo: esperado 3 (A,B,C), obtenido %d", allCount)
	}
}

// --- CASO 2: ApplyInTransaction happy path: snapshot con estado PREVIO + updates aplicados ---

func TestReclassifyRepo_ApplyInTransaction_HappyPath(t *testing.T) {
	db := setupReclassifyDB(t)
	repo := NewPostgresReclassifyRepository(db)
	ctx := context.Background()

	idA := seedProduct(t, db, "7790000000010", "Queso Cremoso", "Lacteos", "scraper_disco", nil)
	idB := seedProduct(t, db, "7790000000011", "Detergente", "Limpieza", "scraper_disco", strptr("almacen"))

	snapshotName := fmt.Sprintf("test_%d", time.Now().UnixNano())
	updates := []value_object.ReclassifyUpdate{
		{ID: idA, Name: "Queso Cremoso", FromType: nil, ToType: "fiambreria", Kind: "relleno"},
		{ID: idB, Name: "Detergente", FromType: strptr("almacen"), ToType: "limpieza", Kind: "correccion"},
	}

	snapRef, affected, collisions, err := repo.ApplyInTransaction(ctx, snapshotName, []string{idA, idB}, updates)
	if err != nil {
		t.Fatalf("ApplyInTransaction: %v", err)
	}
	t.Cleanup(func() { _, _ = db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %q`, snapRef)) })

	if affected != 2 {
		t.Errorf("affected: esperado 2, obtenido %d", affected)
	}
	if collisions != 0 {
		t.Errorf("collisions: esperado 0, obtenido %d", collisions)
	}
	if snapRef != "global_products_bkp_"+snapshotName {
		t.Errorf("snapRef: esperado global_products_bkp_%s, obtenido %s", snapshotName, snapRef)
	}

	// La snapshot debe contener el estado PREVIO (business_type NULL y 'almacen').
	var snapCount int
	if err := db.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM %q`, snapRef)).Scan(&snapCount); err != nil {
		t.Fatalf("count snapshot: %v", err)
	}
	if snapCount != 2 {
		t.Errorf("snapshot rows: esperado 2, obtenido %d", snapCount)
	}
	var prevBT sql.NullString
	if err := db.QueryRow(fmt.Sprintf(`SELECT business_type FROM %q WHERE id::text=$1`, snapRef), idB).Scan(&prevBT); err != nil {
		t.Fatalf("snapshot prev bt: %v", err)
	}
	if !prevBT.Valid || prevBT.String != "almacen" {
		t.Errorf("snapshot debe preservar estado PREVIO de B: esperado 'almacen', obtenido %v", prevBT)
	}

	// La tabla viva ya tiene los nuevos valores.
	var liveBT string
	if err := db.QueryRow(`SELECT business_type FROM global_products WHERE id::text=$1`, idA).Scan(&liveBT); err != nil {
		t.Fatalf("live bt A: %v", err)
	}
	if liveBT != "fiambreria" {
		t.Errorf("live business_type A: esperado fiambreria, obtenido %s", liveBT)
	}
}

// --- CASO 3: Colisión UNIQUE(name, business_type) → collisions++ y la tx NO se rompe ---

func TestReclassifyRepo_ApplyInTransaction_Collision(t *testing.T) {
	db := setupReclassifyDB(t)
	repo := NewPostgresReclassifyRepository(db)
	ctx := context.Background()

	// Dos productos MISMO name. Uno ya está en 'fiambreria'. El otro es candidato (almacen).
	// Si intentamos mover el candidato a 'fiambreria' → choca contra UNIQUE(name, business_type).
	idExisting := seedProduct(t, db, "7790000000020", "Jamon Cocido", "Fiambres", "scraper_disco", strptr("fiambreria"))
	idColliding := seedProduct(t, db, "7790000000021", "Jamon Cocido", "Fiambres", "scraper_disco", strptr("almacen"))
	// Un tercero que SÍ se aplica limpio, DESPUÉS del que colisiona, para probar que el lote sigue vivo.
	idOk := seedProduct(t, db, "7790000000022", "Salame", "Fiambres", "scraper_disco", nil)

	_ = idExisting

	snapshotName := fmt.Sprintf("test_col_%d", time.Now().UnixNano())
	updates := []value_object.ReclassifyUpdate{
		{ID: idColliding, Name: "Jamon Cocido", FromType: strptr("almacen"), ToType: "fiambreria", Kind: "correccion"},
		{ID: idOk, Name: "Salame", FromType: nil, ToType: "fiambreria", Kind: "relleno"},
	}

	snapRef, affected, collisions, err := repo.ApplyInTransaction(ctx, snapshotName, []string{idColliding, idOk}, updates)
	if err != nil {
		t.Fatalf("ApplyInTransaction con colisión NO debe fallar el lote, error: %v", err)
	}
	t.Cleanup(func() { _, _ = db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %q`, snapRef)) })

	if collisions != 1 {
		t.Errorf("collisions: esperado 1, obtenido %d", collisions)
	}
	if affected != 1 {
		t.Errorf("affected: esperado 1 (solo el OK), obtenido %d", affected)
	}

	// El que colisionó debe seguir en 'almacen' (no se aplicó).
	var btCol string
	if err := db.QueryRow(`SELECT business_type FROM global_products WHERE id::text=$1`, idColliding).Scan(&btCol); err != nil {
		t.Fatalf("bt colliding: %v", err)
	}
	if btCol != "almacen" {
		t.Errorf("producto en colisión debe quedar intacto en 'almacen', obtenido %s", btCol)
	}

	// El OK debe haberse aplicado a pesar de la colisión previa.
	var btOk string
	if err := db.QueryRow(`SELECT business_type FROM global_products WHERE id::text=$1`, idOk).Scan(&btOk); err != nil {
		t.Fatalf("bt ok: %v", err)
	}
	if btOk != "fiambreria" {
		t.Errorf("producto posterior a la colisión debe aplicarse, esperado fiambreria, obtenido %s", btOk)
	}
}

// --- CASO 4: Rollback — si el snapshot falla, NO queda mutación ni snapshot huérfana committeada ---

func TestReclassifyRepo_ApplyInTransaction_SnapshotFails_Rollback(t *testing.T) {
	db := setupReclassifyDB(t)
	repo := NewPostgresReclassifyRepository(db)
	ctx := context.Background()

	idA := seedProduct(t, db, "7790000000030", "Yogur", "Lacteos", "scraper_disco", nil)

	// Forzamos que createSnapshot falle: pre-creamos una tabla con el nombre exacto del snapshot.
	// CREATE TABLE global_products_bkp_<name> dará 42P07 (duplicate table) DENTRO de la tx → abort.
	snapshotName := fmt.Sprintf("dup_%d", time.Now().UnixNano())
	tableName := "global_products_bkp_" + snapshotName
	if _, err := db.Exec(fmt.Sprintf(`CREATE TABLE %q (x int)`, tableName)); err != nil {
		t.Fatalf("pre-create colliding snapshot table: %v", err)
	}
	t.Cleanup(func() { _, _ = db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %q`, tableName)) })

	updates := []value_object.ReclassifyUpdate{
		{ID: idA, Name: "Yogur", FromType: nil, ToType: "lacteos", Kind: "relleno"},
	}

	_, affected, _, err := repo.ApplyInTransaction(ctx, snapshotName, []string{idA}, updates)
	if err == nil {
		t.Fatalf("se esperaba error porque el snapshot falla, pero ApplyInTransaction devolvió nil")
	}
	if affected != 0 {
		t.Errorf("affected en rollback: esperado 0, obtenido %d", affected)
	}

	// El producto NO debe haber mutado.
	var bt sql.NullString
	if err := db.QueryRow(`SELECT business_type FROM global_products WHERE id::text=$1`, idA).Scan(&bt); err != nil {
		t.Fatalf("bt post-rollback: %v", err)
	}
	if bt.Valid {
		t.Errorf("producto NO debe mutar en rollback: business_type debería seguir NULL, obtenido %q", bt.String)
	}

	// Solo debe existir la tabla pre-creada por el test (el CREATE TABLE del repo fue parte de la tx abortada).
	var snapCount int
	if err := db.QueryRow(
		`SELECT COUNT(*) FROM pg_tables WHERE schemaname='public' AND tablename=$1`, tableName,
	).Scan(&snapCount); err != nil {
		t.Fatalf("count snapshot table: %v", err)
	}
	if snapCount != 1 {
		t.Errorf("snapshot huérfana: esperado 1 (la pre-creada), obtenido %d", snapCount)
	}
}

// --- CASO 5: SaveAudit inserta la fila con operator_id, mode, scope, snapshot_ref, summary, affected_count ---

func TestReclassifyRepo_SaveAudit(t *testing.T) {
	db := setupReclassifyDB(t)
	repo := NewPostgresReclassifyRepository(db)
	ctx := context.Background()

	scope := mustScope(t, "scraper", 5000)
	snapRef := "global_products_bkp_audit_test"
	audit := value_object.ReclassifyAuditRow{
		OperatorID:  "operator-123",
		ExecutedAt:  time.Now(),
		Mode:        "applied",
		Scope:       scope,
		SnapshotRef: &snapRef,
		Summary: value_object.ReclassifySummary{
			TotalEvaluados:      10,
			Candidatos:          4,
			ColisionesSkipeadas: 1,
			UpdatesPorRubro: value_object.UpdatesByRubro{
				"fiambreria": {Relleno: 2, Correccion: 1},
			},
		},
		AffectedCount: 3,
	}

	if err := repo.SaveAudit(ctx, audit); err != nil {
		t.Fatalf("SaveAudit: %v", err)
	}

	var (
		operatorID    string
		mode          string
		scopeJSON     string
		snapshotRef   sql.NullString
		summaryJSON   string
		affectedCount int
	)
	err := db.QueryRow(
		`SELECT operator_id, mode, scope::text, snapshot_ref, summary::text, affected_count
		 FROM global_product_reclassification_audit ORDER BY executed_at DESC LIMIT 1`,
	).Scan(&operatorID, &mode, &scopeJSON, &snapshotRef, &summaryJSON, &affectedCount)
	if err != nil {
		t.Fatalf("query audit row: %v", err)
	}

	if operatorID != "operator-123" {
		t.Errorf("operator_id: esperado operator-123, obtenido %s", operatorID)
	}
	if mode != "applied" {
		t.Errorf("mode: esperado applied, obtenido %s", mode)
	}
	if affectedCount != 3 {
		t.Errorf("affected_count: esperado 3, obtenido %d", affectedCount)
	}
	if !snapshotRef.Valid || snapshotRef.String != snapRef {
		t.Errorf("snapshot_ref: esperado %s, obtenido %v", snapRef, snapshotRef)
	}
	if !strings.Contains(scopeJSON, `"source_prefix": "scraper"`) && !strings.Contains(scopeJSON, `"source_prefix":"scraper"`) {
		t.Errorf("scope JSONB no contiene source_prefix esperado: %s", scopeJSON)
	}
	if !strings.Contains(summaryJSON, "fiambreria") {
		t.Errorf("summary JSONB no contiene el rubro esperado: %s", summaryJSON)
	}
}
