package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/hornosg/go-shared/domain/category"
	"saas-mt-pim-service/src/pim/domain/port"
	domainport "saas-mt-pim-service/src/product/global_catalog/domain/port"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// maxSinClasificarMuestra acota la lista de categorías sin resolver incluida en el summary
// (es el worklist de overrides; no hace falta el universo entero para curar).
const maxSinClasificarMuestra = 300

// NormalizeRequest es el input del use case.
type NormalizeRequest struct {
	DryRun     bool // true (default) simula sin mutar; el summary es idéntico al del apply.
	Confirm    bool // apply requiere DryRun=false AND Confirm=true.
	Scope      value_object.NormalizeCategoryScope
	OperatorID string // X-Operator-Id; obligatorio en apply (validado en el handler).
}

// NormalizeResponse es el output del use case.
type NormalizeResponse struct {
	Mode         string                                `json:"mode"`
	SnapshotRef  *string                               `json:"snapshot_ref"`
	Summary      value_object.NormalizeSummary         `json:"summary"`
	AntesDespues []value_object.NormalizeMappingSample `json:"antes_despues"`
}

// NormalizeError es el error tipado con código HTTP y semántico (mismo patrón que ReclassifyError).
type NormalizeError struct {
	HTTPStatus int
	Code       string
	Message    string
}

func (e *NormalizeError) Error() string {
	return fmt.Sprintf("[%d][%s] %s", e.HTTPStatus, e.Code, e.Message)
}

// NormalizeCategorySlugsUseCase normaliza global_products.category_slug a partir de category raw,
// usando el resolver determinístico de go-shared + la tabla de overrides (ADR-007 §4).
//
// Invariantes (ADR-007 §6):
//   - category_slug nunca queda NULL: o resuelve, o cae en "sin-clasificar".
//   - dry_run=true no muta; apply requiere DryRun=false AND Confirm=true.
//   - Idempotencia: 2do apply consecutivo → 0 filas (UPDATE con IS DISTINCT FROM).
//   - Snapshot falla → no se aplica nada (tx encapsulada en el repo).
type NormalizeCategorySlugsUseCase struct {
	repo   domainport.NormalizeCategoryRepository
	logger port.PIMEventLogger
}

// NewNormalizeCategorySlugsUseCase construye el use case.
func NewNormalizeCategorySlugsUseCase(
	repo domainport.NormalizeCategoryRepository,
	logger port.PIMEventLogger,
) *NormalizeCategorySlugsUseCase {
	return &NormalizeCategorySlugsUseCase{repo: repo, logger: logger}
}

// Execute corre la normalización según el request.
func (uc *NormalizeCategorySlugsUseCase) Execute(ctx context.Context, req NormalizeRequest) (*NormalizeResponse, error) {
	count, err := uc.repo.CountDistinctCategories(ctx, req.Scope)
	if err != nil {
		return nil, fmt.Errorf("count distinct categories: %w", err)
	}
	if count > req.Scope.MaxRows {
		return nil, &NormalizeError{
			HTTPStatus: 422,
			Code:       "CATEGORIES_EXCEED_MAX_ROWS",
			Message:    fmt.Sprintf("found %d distinct categories, exceeds max_rows=%d; narrow your scope", count, req.Scope.MaxRows),
		}
	}

	categories, err := uc.repo.FetchDistinctCategories(ctx, req.Scope)
	if err != nil {
		return nil, fmt.Errorf("fetch distinct categories: %w", err)
	}

	overrides, err := uc.repo.LoadOverrides(ctx)
	if err != nil {
		return nil, fmt.Errorf("load overrides: %w", err)
	}

	mappings, summary := resolveMappings(categories, overrides)
	antesDespues := buildNormalizeSample(mappings)

	isApply := !req.DryRun && req.Confirm
	mode := "dry_run"
	var snapshotRef *string
	affectedCount := 0

	if isApply {
		snapshotName := time.Now().Format("20060102_150405")
		ref, affected, applyErr := uc.repo.ApplyInTransaction(ctx, snapshotName, req.Scope, mappings)
		if applyErr != nil {
			uc.logEvent(port.PIMEvent{Event: "catalog.category_slug_normalization_failed", UserID: req.OperatorID, Reason: applyErr.Error()})
			return nil, applyErr
		}
		mode = "applied"
		snapshotRef = &ref
		affectedCount = affected
	}

	auditRow := value_object.NormalizeAuditRow{
		OperatorID:    req.OperatorID,
		ExecutedAt:    time.Now(),
		Mode:          mode,
		Scope:         req.Scope,
		SnapshotRef:   snapshotRef,
		Summary:       summary,
		AffectedCount: affectedCount,
	}
	if auditErr := uc.repo.SaveAudit(ctx, auditRow); auditErr != nil {
		uc.logEvent(port.PIMEvent{Event: "catalog.category_slug_normalization_audit_failed", UserID: req.OperatorID, Reason: auditErr.Error()})
	}

	uc.logEvent(port.PIMEvent{Event: "catalog.category_slug_normalization_completed", UserID: req.OperatorID, Count: affectedCount, Reason: mode})

	return &NormalizeResponse{
		Mode:         mode,
		SnapshotRef:  snapshotRef,
		Summary:      summary,
		AntesDespues: antesDespues,
	}, nil
}

// resolveMappings aplica override > resolver > sin-clasificar a cada categoría cruda y acumula
// el summary. Es función pura (testeable sin DB).
func resolveMappings(
	categories []value_object.RawCategoryCount,
	overrides map[string]string,
) ([]value_object.CategoryMapping, value_object.NormalizeSummary) {
	summary := value_object.NormalizeSummary{
		CategoriasPorFuente: map[string]int{},
		ProductosPorFuente:  map[string]int{},
		TopSlugs:            map[string]int{},
	}
	mappings := make([]value_object.CategoryMapping, 0, len(categories))

	for _, c := range categories {
		raw := ""
		if c.RawCategory != nil {
			raw = *c.RawCategory
		}

		var slug, source string
		if ov, ok := overrides[raw]; ok && raw != "" {
			slug, source = ov, value_object.SourceOverride
		} else if resolved, ok := category.ResolveCategorySlug(raw); ok {
			slug, source = resolved, value_object.SourceResolver
		} else {
			slug, source = category.Unclassified, value_object.SourceUnclassified
		}

		changed := c.CurrentSlug == nil || *c.CurrentSlug != slug

		mappings = append(mappings, value_object.CategoryMapping{
			RawCategory:  c.RawCategory,
			ProductCount: c.ProductCount,
			CurrentSlug:  c.CurrentSlug,
			NewSlug:      slug,
			Source:       source,
			Changed:      changed,
		})

		summary.TotalCategorias++
		summary.TotalProductos += c.ProductCount
		summary.CategoriasPorFuente[source]++
		summary.ProductosPorFuente[source] += c.ProductCount
		summary.TopSlugs[slug] += c.ProductCount
		if changed {
			summary.ProductosAfectados += c.ProductCount
		}
		if source == value_object.SourceUnclassified && raw != "" && len(summary.SinClasificarMuestra) < maxSinClasificarMuestra {
			summary.SinClasificarMuestra = append(summary.SinClasificarMuestra, raw)
		}
	}

	return mappings, summary
}

// buildNormalizeSample arma una muestra acotada (máx maxAntesDespues) de los mappings que cambian.
func buildNormalizeSample(mappings []value_object.CategoryMapping) []value_object.NormalizeMappingSample {
	sample := make([]value_object.NormalizeMappingSample, 0, maxAntesDespues)
	for _, m := range mappings {
		if !m.Changed {
			continue
		}
		if len(sample) >= maxAntesDespues {
			break
		}
		sample = append(sample, value_object.NormalizeMappingSample{
			RawCategory:  m.RawCategory,
			ProductCount: m.ProductCount,
			From:         m.CurrentSlug,
			To:           m.NewSlug,
			Source:       m.Source,
		})
	}
	return sample
}

func (uc *NormalizeCategorySlugsUseCase) logEvent(e port.PIMEvent) {
	if uc.logger != nil {
		uc.logger.Log(e)
	}
}
