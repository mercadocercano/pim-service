package port

import (
	"context"

	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// BulkVerifyRepository define las operaciones de persistencia para la verificación/desverificación
// masiva de global_products.
//
// Contratos de implementación:
//   - ApplyInTransaction encapsula la creación del snapshot y los UPDATEs en una sola transacción
//     atómica. Si el snapshot falla, la tx aborta y nada se aplica.
//   - SaveAudit persiste fuera de la tx del apply.
type BulkVerifyRepository interface {
	// ApplyInTransaction crea el snapshot y aplica los UPDATEs de is_verified en una tx única.
	// snapshotName es el sufijo de timestamp para global_products_bkp_bulkverify_<snapshotName>.
	// Devuelve la cantidad de filas efectivamente actualizadas.
	ApplyInTransaction(ctx context.Context, snapshotName string, ids []string, mode value_object.BulkVerifyMode) (int, error)

	// SaveAudit persiste el registro de auditoría en global_product_bulk_verify_audit.
	SaveAudit(ctx context.Context, audit value_object.BulkVerifyAuditRow) error
}
