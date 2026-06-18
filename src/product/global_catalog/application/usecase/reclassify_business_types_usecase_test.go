package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"saas-mt-pim-service/src/product/global_catalog/application/usecase"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// --- Mock del repositorio ---

type mockReclassifyRepo struct {
	candidates       []value_object.ReclassifyCandidate
	countResult      int
	countErr         error
	fetchErr         error
	applyInTxErr     error
	applyAffected    int
	applyCollisions  int
	auditErr         error
	applyInTxCalls   int
	auditCalls       int
	simulateSnapshot bool // si true, simula fallo de snapshot en ApplyInTransaction
}

func (m *mockReclassifyRepo) CountCandidates(_ context.Context, _ value_object.ReclassifyScope) (int, error) {
	if m.countErr != nil {
		return 0, m.countErr
	}
	if m.countResult > 0 {
		return m.countResult, nil
	}
	return len(m.candidates), nil
}

func (m *mockReclassifyRepo) FetchCandidates(_ context.Context, _ value_object.ReclassifyScope) ([]value_object.ReclassifyCandidate, error) {
	if m.fetchErr != nil {
		return nil, m.fetchErr
	}
	return m.candidates, nil
}

func (m *mockReclassifyRepo) ApplyInTransaction(_ context.Context, snapshotName string, _ []string, updates []value_object.ReclassifyUpdate) (string, int, int, error) {
	m.applyInTxCalls++
	if m.simulateSnapshot {
		return "", 0, 0, fmt.Errorf("simulated snapshot failure")
	}
	if m.applyInTxErr != nil {
		return "", 0, 0, m.applyInTxErr
	}
	affected := len(updates)
	if m.applyAffected > 0 {
		affected = m.applyAffected
	}
	return "global_products_bkp_" + snapshotName, affected, m.applyCollisions, nil
}

func (m *mockReclassifyRepo) SaveAudit(_ context.Context, _ value_object.ReclassifyAuditRow) error {
	m.auditCalls++
	return m.auditErr
}

// buildScope crea un scope de prueba estándar.
func buildScope(t *testing.T) value_object.ReclassifyScope {
	t.Helper()
	scope, err := value_object.NewReclassifyScope("scraper", 1000)
	if err != nil {
		t.Fatalf("buildScope: %v", err)
	}
	return scope
}

// --- Tests T-016..T-025 ---

// T-016: dry_run=true — el repo NO recibe ApplyInTransaction.
func TestReclassifyUseCase_DryRun_NoMutation(t *testing.T) {
	almacen := "almacen"
	repo := &mockReclassifyRepo{
		candidates: []value_object.ReclassifyCandidate{
			{ID: "p1", Name: "Leche Entera", Category: "/Lácteos/Leches/", BusinessType: &almacen},
		},
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	req := usecase.ReclassifyRequest{
		DryRun:  true,
		Confirm: false,
		Scope:   buildScope(t),
	}
	resp, err := uc.Execute(context.Background(), req)
	if err != nil {
		t.Fatalf("T-016: unexpected error: %v", err)
	}
	if repo.applyInTxCalls != 0 {
		t.Errorf("T-016: ApplyInTransaction called %d times in dry_run, want 0", repo.applyInTxCalls)
	}
	if resp.Mode != "dry_run" {
		t.Errorf("T-016: expected mode=dry_run, got %q", resp.Mode)
	}
}

// T-017: dry_run=true — summary devuelto tiene candidatos > 0 y mode=dry_run.
func TestReclassifyUseCase_DryRun_SummaryMatchesExpected(t *testing.T) {
	almacen := "almacen"
	repo := &mockReclassifyRepo{
		candidates: []value_object.ReclassifyCandidate{
			{ID: "p1", Name: "Leche Entera", Category: "/Lácteos/Leches/", BusinessType: &almacen},
		},
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	req := usecase.ReclassifyRequest{DryRun: true, Scope: buildScope(t)}
	resp, err := uc.Execute(context.Background(), req)
	if err != nil {
		t.Fatalf("T-017: %v", err)
	}
	if resp.Summary.Candidatos == 0 {
		t.Error("T-017: expected candidatos > 0 in dry_run summary")
	}
	if resp.Mode != "dry_run" {
		t.Errorf("T-017: expected mode=dry_run, got %q", resp.Mode)
	}
}

// T-018: dry_run=false AND confirm=false → no aplica (trata como dry_run).
func TestReclassifyUseCase_NoConfirm_DoesNotApply(t *testing.T) {
	almacen := "almacen"
	repo := &mockReclassifyRepo{
		candidates: []value_object.ReclassifyCandidate{
			{ID: "p1", Name: "Leche Entera", Category: "/Lácteos/Leches/", BusinessType: &almacen},
		},
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	req := usecase.ReclassifyRequest{DryRun: false, Confirm: false, Scope: buildScope(t)}
	resp, err := uc.Execute(context.Background(), req)
	if err != nil {
		t.Fatalf("T-018: %v", err)
	}
	if repo.applyInTxCalls != 0 {
		t.Errorf("T-018: ApplyInTransaction called without confirm, want 0 calls")
	}
	if resp.Mode != "dry_run" {
		t.Errorf("T-018: expected mode=dry_run when confirm=false, got %q", resp.Mode)
	}
}

// T-019: producto ya en rubro específico (no almacen, no nil) → skip ya_especifico.
func TestReclassifyUseCase_YaEspecifico_Skipped(t *testing.T) {
	fiambreria := "fiambreria"
	repo := &mockReclassifyRepo{
		candidates: []value_object.ReclassifyCandidate{
			{ID: "p1", Name: "Queso Cremoso", Category: "/Lácteos/Quesos/", BusinessType: &fiambreria},
		},
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	resp, err := uc.Execute(context.Background(), usecase.ReclassifyRequest{DryRun: true, Scope: buildScope(t)})
	if err != nil {
		t.Fatalf("T-019: %v", err)
	}
	if resp.Summary.Skips.YaEspecifico != 1 {
		t.Errorf("T-019: expected skips.ya_especifico=1, got %d", resp.Summary.Skips.YaEspecifico)
	}
	if resp.Summary.Candidatos != 0 {
		t.Errorf("T-019: expected 0 updates (ya específico skipped), got %d", resp.Summary.Candidatos)
	}
}

// T-020: producto ya en rubro correcto (almacen y el resolver también devuelve almacen) → skip ya_correcto.
func TestReclassifyUseCase_YaCorrecto_Skipped(t *testing.T) {
	almacen := "almacen"
	// Category con keyword "almacen" → resolver devuelve almacen → ya_correcto (ya tiene almacen).
	repo := &mockReclassifyRepo{
		candidates: []value_object.ReclassifyCandidate{
			{ID: "p1", Name: "Arroz", Category: "arroz integral", BusinessType: &almacen},
		},
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	resp, err := uc.Execute(context.Background(), usecase.ReclassifyRequest{DryRun: true, Scope: buildScope(t)})
	if err != nil {
		t.Fatalf("T-020: %v", err)
	}
	// arroz → almacen; el producto ya tiene almacen → ya_correcto
	if resp.Summary.Skips.YaCorrecto != 1 {
		t.Errorf("T-020: expected skips.ya_correcto=1, got %d (noResuelve=%d, yaEspecifico=%d)", resp.Summary.Skips.YaCorrecto, resp.Summary.Skips.NoResuelve, resp.Summary.Skips.YaEspecifico)
	}
	if resp.Summary.Candidatos != 0 {
		t.Errorf("T-020: expected 0 updates, got %d", resp.Summary.Candidatos)
	}
}

// T-021: categoría sin resolver → skip no_resuelve.
func TestReclassifyUseCase_NoResuelve_Skipped(t *testing.T) {
	repo := &mockReclassifyRepo{
		candidates: []value_object.ReclassifyCandidate{
			{ID: "p1", Name: "Xyzzy Inasignable 99", Category: "Xyzzy Inasignable 99", BusinessType: nil},
		},
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	resp, err := uc.Execute(context.Background(), usecase.ReclassifyRequest{DryRun: true, Scope: buildScope(t)})
	if err != nil {
		t.Fatalf("T-021: %v", err)
	}
	if resp.Summary.Skips.NoResuelve != 1 {
		t.Errorf("T-021: expected skips.no_resuelve=1, got %d", resp.Summary.Skips.NoResuelve)
	}
}

// T-022: colisión UNIQUE → skip del lote no se aborta.
// El mock simula que ApplyInTransaction aplica solo 1 de 2 (colisión en uno).
// El use case no debe panicar ni abortar.
func TestReclassifyUseCase_Collision_SkipNotAbort(t *testing.T) {
	almacen := "almacen"
	repo := &mockReclassifyRepo{
		candidates: []value_object.ReclassifyCandidate{
			{ID: "p1", Name: "Leche A", Category: "/Lácteos/Leches/", BusinessType: &almacen},
			{ID: "p2", Name: "Leche B", Category: "/Lácteos/Leches/", BusinessType: &almacen},
		},
		applyAffected:   1, // simula que solo 1 fue efectivo (1 colisionó)
		applyCollisions: 1, // y que esa 1 colisión se reporta de vuelta
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	req := usecase.ReclassifyRequest{DryRun: false, Confirm: true, Scope: buildScope(t)}
	resp, err := uc.Execute(context.Background(), req)
	if err != nil {
		t.Fatalf("T-022: use case aborted on collision: %v", err)
	}
	if resp == nil {
		t.Fatal("T-022: expected non-nil response")
	}
	if resp.Mode != "applied" {
		t.Errorf("T-022: expected mode=applied, got %q", resp.Mode)
	}
	if resp.Summary.ColisionesSkipeadas != 1 {
		t.Errorf("T-022: expected colisiones_skipeadas=1 propagada al summary, got %d", resp.Summary.ColisionesSkipeadas)
	}
}

// T-023: CountCandidates > max_rows → use case devuelve error 422 antes de fetch.
func TestReclassifyUseCase_ExceedsMaxRows_Returns422(t *testing.T) {
	repo := &mockReclassifyRepo{
		countResult: 1001, // más que el MaxRows del scope
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	scope, _ := value_object.NewReclassifyScope("scraper", 10) // MaxRows=10, count=1001
	req := usecase.ReclassifyRequest{DryRun: true, Scope: scope}
	_, err := uc.Execute(context.Background(), req)
	if err == nil {
		t.Fatal("T-023: expected error when candidates > max_rows")
	}
	var reclErr *usecase.ReclassifyError
	if !errors.As(err, &reclErr) {
		t.Fatalf("T-023: expected ReclassifyError, got %T: %v", err, err)
	}
	if reclErr.HTTPStatus != 422 {
		t.Errorf("T-023: expected HTTPStatus=422, got %d", reclErr.HTTPStatus)
	}
}

// T-024: snapshot falla → ApplyInTransaction retorna error, use case retorna error; 0 updates.
func TestReclassifyUseCase_SnapshotFails_NoUpdates(t *testing.T) {
	almacen := "almacen"
	repo := &mockReclassifyRepo{
		candidates: []value_object.ReclassifyCandidate{
			{ID: "p1", Name: "Leche Entera", Category: "/Lácteos/Leches/", BusinessType: &almacen},
		},
		simulateSnapshot: true, // el repo simula fallo de snapshot dentro de ApplyInTransaction
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	req := usecase.ReclassifyRequest{DryRun: false, Confirm: true, Scope: buildScope(t)}
	_, err := uc.Execute(context.Background(), req)
	if err == nil {
		t.Fatal("T-024: expected error when snapshot fails")
	}
	// El apply fue llamado (y falló en snapshot), pero el use case retorna el error
	if repo.applyInTxCalls != 1 {
		t.Errorf("T-024: expected ApplyInTransaction called once (and failed), got %d", repo.applyInTxCalls)
	}
}

// T-025: idempotencia — 2do apply con 0 candidatos → affected_count=0, mode=applied.
func TestReclassifyUseCase_Idempotency_ZeroUpdatesSecondRun(t *testing.T) {
	// Segunda corrida: 0 candidatos (todos ya clasificados)
	repo := &mockReclassifyRepo{
		candidates: []value_object.ReclassifyCandidate{},
	}
	uc := usecase.NewReclassifyBusinessTypesUseCase(repo, nil)

	req := usecase.ReclassifyRequest{DryRun: false, Confirm: true, Scope: buildScope(t)}
	resp, err := uc.Execute(context.Background(), req)
	if err != nil {
		t.Fatalf("T-025: %v", err)
	}
	if resp.Mode != "applied" {
		t.Errorf("T-025: expected mode=applied, got %q", resp.Mode)
	}
	if resp.Summary.Candidatos != 0 {
		t.Errorf("T-025: expected candidatos=0 on 2nd apply, got %d", resp.Summary.Candidatos)
	}
}
