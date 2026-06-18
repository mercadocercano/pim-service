package port

import (
	"context"

	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// ReclassifyRepository define las operaciones de persistencia necesarias para
// el use case ReclassifyBusinessTypesUseCase.
//
// Contratos de implementación:
//   - ApplyInTransaction encapsula la creación del snapshot y los UPDATEs en una sola
//     transacción atómica. Si el snapshot falla, la tx aborta y nada se aplica (ADR-005 §3).
//   - ApplyInTransaction maneja colisión 23505 como skip individual; no aborta el lote (ADR-005 §8).
//   - SaveAudit persiste fuera de la tx del apply (se llama siempre, incluso en dry_run).
type ReclassifyRepository interface {
	// CountCandidates devuelve la cantidad de productos candidatos según el scope.
	// Candidatos: business_type IS NULL OR business_type = 'almacen', source LIKE prefix%.
	CountCandidates(ctx context.Context, scope value_object.ReclassifyScope) (int, error)

	// FetchCandidates devuelve los productos candidatos para evaluación.
	FetchCandidates(ctx context.Context, scope value_object.ReclassifyScope) ([]value_object.ReclassifyCandidate, error)

	// ApplyInTransaction crea el snapshot y aplica los updates en una transacción única.
	// snapshotName es el sufijo de timestamp para global_products_bkp_<snapshotName>.
	// Invariante: si el snapshot falla → rollback completo, 0 updates (ADR-005 §3, TEST-ID T-024).
	// Colisión 23505 → skip individual del producto (vía SAVEPOINT), nunca rollback del lote (ADR-005 §8, T-022).
	// Devuelve el nombre completo del snapshot creado, el conteo de rows actualizados y el de colisiones skipeadas.
	ApplyInTransaction(ctx context.Context, snapshotName string, candidateIDs []string, updates []value_object.ReclassifyUpdate) (snapshotRef string, affected int, collisions int, err error)

	// SaveAudit persiste el registro de auditoría en global_product_reclassification_audit.
	// Se llama fuera de la tx del apply (el audit se guarda siempre, incluso en dry_run).
	SaveAudit(ctx context.Context, audit value_object.ReclassifyAuditRow) error
}
