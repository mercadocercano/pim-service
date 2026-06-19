package port

import (
	"context"

	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// NormalizeCategoryRepository define la persistencia del use case NormalizeCategorySlugsUseCase
// (ADR-007 §4). La unidad de trabajo es la CATEGORÍA cruda distinta, no la fila.
//
// Contratos:
//   - ApplyInTransaction crea el snapshot de toda la tabla y aplica los UPDATEs por categoría
//     en una sola transacción. Si el snapshot falla, la tx aborta y nada se aplica (ADR-005 §3).
//   - SaveAudit persiste fuera de la tx del apply (se llama siempre, incluso en dry_run).
type NormalizeCategoryRepository interface {
	// CountDistinctCategories devuelve la cantidad de categorías crudas distintas según el scope
	// (incluye NULL como una categoría más).
	CountDistinctCategories(ctx context.Context, scope value_object.NormalizeCategoryScope) (int, error)

	// FetchDistinctCategories devuelve las categorías crudas distintas con su conteo de productos
	// y el category_slug actual (si todas las filas de esa categoría comparten el mismo; si no,
	// devuelve nil para forzar recálculo).
	FetchDistinctCategories(ctx context.Context, scope value_object.NormalizeCategoryScope) ([]value_object.RawCategoryCount, error)

	// LoadOverrides devuelve el mapa raw_category -> category_slug de category_slug_overrides.
	LoadOverrides(ctx context.Context) (map[string]string, error)

	// ApplyInTransaction crea el snapshot (tabla completa) y aplica un UPDATE por mapping en
	// una transacción única. El scope replica el WHERE del fetch para que el conteo de filas
	// afectadas coincida con el dry_run. Devuelve el snapshot y la cantidad de FILAS actualizadas.
	ApplyInTransaction(ctx context.Context, snapshotName string, scope value_object.NormalizeCategoryScope, mappings []value_object.CategoryMapping) (snapshotRef string, affected int, err error)

	// SaveAudit persiste el registro en global_product_reclassification_audit (tabla reusada).
	SaveAudit(ctx context.Context, audit value_object.NormalizeAuditRow) error
}
