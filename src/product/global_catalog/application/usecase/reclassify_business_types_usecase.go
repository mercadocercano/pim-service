package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/hornosg/go-shared/domain/businesstype"
	"saas-mt-pim-service/src/pim/domain/port"
	domainport "saas-mt-pim-service/src/product/global_catalog/domain/port"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

const (
	// maxAntesDespues es la cantidad máxima de filas incluidas en AntesDespues de la response.
	maxAntesDespues = 100
)

// ReclassifyBusinessTypesUseCase orquesta la re-clasificación masiva de business_type
// en global_products usando el resolver compartido de go-shared (ADR-005 §5).
//
// Invariantes (ADR-005 §8):
//   - dry_run=true NO muta y devuelve el MISMO summary que el apply.
//   - apply requiere DryRun=false AND Confirm=true.
//   - Idempotencia: 2do apply consecutivo → affected_count=0.
//   - Ningún rubro específico pierde productos (ya_especifico → skip).
//   - Colisión UNIQUE → skip individual, nunca aborta el lote.
//   - Snapshot falla → no se aplica nada (misma tx encapsulada en el repo).
type ReclassifyBusinessTypesUseCase struct {
	repo   domainport.ReclassifyRepository
	logger port.PIMEventLogger
}

// NewReclassifyBusinessTypesUseCase construye el use case con sus dependencias.
func NewReclassifyBusinessTypesUseCase(
	repo domainport.ReclassifyRepository,
	logger port.PIMEventLogger,
) *ReclassifyBusinessTypesUseCase {
	return &ReclassifyBusinessTypesUseCase{
		repo:   repo,
		logger: logger,
	}
}

// Execute corre la re-clasificación según el request.
//
// Flujo:
//  1. CountCandidates → 422 si excede MaxRows (TEST-ID T-023).
//  2. FetchCandidates.
//  3. Resolver cada producto con businesstype.ResolveBusinessTypeFromProductCategory.
//  4. Clasificar en updates/skips por kind.
//  5. Si dry_run=true o Confirm=false → devolver summary sin mutar (TEST-ID T-016, T-017, T-018).
//  6. Apply: ApplyInTransaction (snapshot + updates en tx única; TEST-ID T-024).
//  7. SaveAudit (siempre, incluso en dry_run).
//  8. Log canónico ADR-001.
func (uc *ReclassifyBusinessTypesUseCase) Execute(ctx context.Context, req ReclassifyRequest) (*ReclassifyResponse, error) {
	// Paso 1: contar candidatos y aplicar cap
	count, err := uc.repo.CountCandidates(ctx, req.Scope)
	if err != nil {
		return nil, fmt.Errorf("count candidates: %w", err)
	}

	if count > req.Scope.MaxRows {
		return nil, &ReclassifyError{
			HTTPStatus: 422,
			Code:       "CANDIDATES_EXCEED_MAX_ROWS",
			Message:    fmt.Sprintf("found %d candidates, exceeds max_rows=%d; narrow your scope", count, req.Scope.MaxRows),
		}
	}

	// Paso 2: obtener candidatos
	candidates, err := uc.repo.FetchCandidates(ctx, req.Scope)
	if err != nil {
		return nil, fmt.Errorf("fetch candidates: %w", err)
	}

	// Paso 3-4: resolver y clasificar
	updates, _, summary := uc.classifyCandidates(candidates)

	// Construir muestra acotada de antes/después
	antesDespues := buildAntesDespues(updates)

	// Determinar si realmente aplica: DryRun=false AND Confirm=true
	isApply := !req.DryRun && req.Confirm

	mode := "dry_run"
	var snapshotRef *string
	affectedCount := 0

	if isApply {
		// Paso 6: ejecutar en transacción (snapshot + updates atómicos)
		snapshotName := time.Now().Format("20060102_150405")
		candidateIDs := make([]string, len(updates))
		for i, u := range updates {
			candidateIDs[i] = u.ID
		}

		ref, affected, collisions, applyErr := uc.repo.ApplyInTransaction(ctx, snapshotName, candidateIDs, updates)
		if applyErr != nil {
			uc.logEvent(port.PIMEvent{
				Event:  "catalog.reclassification_failed",
				UserID: req.OperatorID,
				Count:  len(candidates),
				Reason: applyErr.Error(),
			})
			return nil, applyErr
		}
		mode = "applied"
		snapshotRef = &ref
		affectedCount = affected
		// Las colisiones UNIQUE solo se conocen al aplicar (son a nivel DB); en dry_run
		// quedan en 0 porque no se intenta el UPDATE. Esto hace que el conteo de
		// colisiones sea el único campo del summary que puede diferir entre dry_run y apply.
		summary.ColisionesSkipeadas = collisions
	}

	// Paso 7: guardar audit (siempre, incluso en dry_run)
	auditRow := value_object.ReclassifyAuditRow{
		OperatorID:    req.OperatorID,
		ExecutedAt:    time.Now(),
		Mode:          mode,
		Scope:         req.Scope,
		SnapshotRef:   snapshotRef,
		Summary:       summary,
		AffectedCount: affectedCount,
	}
	if auditErr := uc.repo.SaveAudit(ctx, auditRow); auditErr != nil {
		// El audit no puede fallar silenciosamente, pero tampoco aborta la response
		// ya committeada. Logueamos el error.
		uc.logEvent(port.PIMEvent{
			Event:  "catalog.reclassification_audit_failed",
			UserID: req.OperatorID,
			Reason: auditErr.Error(),
		})
	}

	// Paso 8: log canónico ADR-001
	uc.logEvent(port.PIMEvent{
		Event:  "catalog.reclassification_completed",
		UserID: req.OperatorID,
		Count:  affectedCount,
		Reason: mode,
	})

	return &ReclassifyResponse{
		Mode:         mode,
		SnapshotRef:  snapshotRef,
		Summary:      summary,
		AntesDespues: antesDespues,
	}, nil
}

// classifyCandidates aplica el resolver a cada candidato y los clasifica en updates o skips.
//
// Reglas de clasificación (ADR-005 §8):
//   - Si resolver no resuelve → SkipNoResuelve.
//   - Si producto ya tiene rubro específico (no nil, no "almacen") → SkipYaEspecifico.
//   - Si resolver devuelve el mismo rubro que ya tiene → SkipYaCorrecto.
//   - Caso restante (nil/almacen → rubro específico) → update.
func (uc *ReclassifyBusinessTypesUseCase) classifyCandidates(
	candidates []value_object.ReclassifyCandidate,
) ([]value_object.ReclassifyUpdate, []value_object.ReclassifySkip, value_object.ReclassifySummary) {
	var updates []value_object.ReclassifyUpdate
	var skips []value_object.ReclassifySkip

	summary := value_object.ReclassifySummary{
		TotalEvaluados:  len(candidates),
		UpdatesPorRubro: make(value_object.UpdatesByRubro),
	}

	for _, c := range candidates {
		assignment, resolved := businesstype.ResolveBusinessTypeFromProductCategory(c.Category)

		// Caso: no resuelve
		if !resolved {
			skips = append(skips, value_object.ReclassifySkip{ID: c.ID, Name: c.Name, Kind: value_object.SkipNoResuelve})
			summary.Skips.NoResuelve++
			continue
		}

		newType := assignment.BusinessTypeCode

		// Política de corrección segura §8: única implementación, compartida con el
		// camino del re-sync (UpdateGlobalProductByID). No se duplica el criterio acá.
		apply, toType, kind := value_object.ResolveSafeBusinessTypeTransition(c.BusinessType, newType)
		if !apply {
			switch kind {
			case value_object.TransitionSkipYaEspecifico:
				// Invariante: nunca se toca un producto ya en rubro específico.
				skips = append(skips, value_object.ReclassifySkip{ID: c.ID, Name: c.Name, Kind: value_object.SkipYaEspecifico})
				summary.Skips.YaEspecifico++
			case value_object.TransitionSkipYaCorrecto:
				skips = append(skips, value_object.ReclassifySkip{ID: c.ID, Name: c.Name, Kind: value_object.SkipYaCorrecto})
				summary.Skips.YaCorrecto++
			default:
				// sin_candidate: candidate vacío. No ocurre acá (resolved==true ⇒ candidate no vacío),
				// pero se contabiliza como no_resuelve por completitud.
				skips = append(skips, value_object.ReclassifySkip{ID: c.ID, Name: c.Name, Kind: value_object.SkipNoResuelve})
				summary.Skips.NoResuelve++
			}
			continue
		}

		updates = append(updates, value_object.ReclassifyUpdate{
			ID:       c.ID,
			Name:     c.Name,
			FromType: c.BusinessType,
			ToType:   toType,
			Kind:     kind,
		})

		// Acumular en summary por rubro.
		rubroCount := summary.UpdatesPorRubro[toType]
		if kind == value_object.TransitionRelleno {
			rubroCount.Relleno++
		} else {
			rubroCount.Correccion++
		}
		summary.UpdatesPorRubro[toType] = rubroCount
	}

	summary.Candidatos = len(updates)
	return updates, skips, summary
}

// buildAntesDespues construye una muestra acotada (máximo maxAntesDespues filas)
// del detalle antes/después para la response.
func buildAntesDespues(updates []value_object.ReclassifyUpdate) []value_object.AntesDespuesRow {
	limit := len(updates)
	if limit > maxAntesDespues {
		limit = maxAntesDespues
	}
	rows := make([]value_object.AntesDespuesRow, 0, limit)
	for i := 0; i < limit; i++ {
		u := updates[i]
		rows = append(rows, value_object.AntesDespuesRow{
			ID:   u.ID,
			Name: u.Name,
			From: u.FromType,
			To:   u.ToType,
		})
	}
	return rows
}

// logEvent emite un evento canónico ADR-001; no paniquea si el logger es nil.
func (uc *ReclassifyBusinessTypesUseCase) logEvent(e port.PIMEvent) {
	if uc.logger != nil {
		uc.logger.Log(e)
	}
}

// ReclassifyError es el error tipado del use case con código HTTP y semántico.
type ReclassifyError struct {
	HTTPStatus int
	Code       string
	Message    string
}

func (e *ReclassifyError) Error() string {
	return fmt.Sprintf("[%d][%s] %s", e.HTTPStatus, e.Code, e.Message)
}
