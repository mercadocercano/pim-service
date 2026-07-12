package usecase

import (
	"context"
	"fmt"
	"time"

	"saas-mt-pim-service/src/pim/domain/port"
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	domainport "saas-mt-pim-service/src/product/global_catalog/domain/port"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

const (
	// MaxBulkVerifyIDs es la cantidad máxima de productos que se pueden verificar/desverificar
	// en una sola operación. Coincide con el límite de selección de mc_admin.
	MaxBulkVerifyIDs = 100
)

// BulkVerifyGlobalProductsRequest identifica los productos y la acción a aplicar.
type BulkVerifyGlobalProductsRequest struct {
	IDs        []string `json:"ids"`
	Verify     bool     `json:"verify"`
	OperatorID string   `json:"operator_id"`
}

// BulkVerifyGlobalProductsResponse resume el resultado de la operación masiva.
type BulkVerifyGlobalProductsResponse struct {
	Mode           value_object.BulkVerifyMode `json:"mode"`
	SnapshotRef    *string                     `json:"snapshot_ref,omitempty"`
	Summary        value_object.BulkVerifySummary `json:"summary"`
}

// BulkVerifyGlobalProducts implementa la verificación/desverificación en lote de productos globales.
type BulkVerifyGlobalProducts struct {
	globalProductRepo domainport.GlobalProductRepository
	bulkVerifyRepo    domainport.BulkVerifyRepository
	logger            port.PIMEventLogger
}

// NewBulkVerifyGlobalProducts crea una nueva instancia del caso de uso.
func NewBulkVerifyGlobalProducts(
	globalProductRepo domainport.GlobalProductRepository,
	bulkVerifyRepo domainport.BulkVerifyRepository,
	logger port.PIMEventLogger,
) *BulkVerifyGlobalProducts {
	return &BulkVerifyGlobalProducts{
		globalProductRepo: globalProductRepo,
		bulkVerifyRepo:    bulkVerifyRepo,
		logger:            logger,
	}
}

// Execute verifica o desverifica los productos indicados de forma masiva y auditada.
//
// Flujo:
//  1. Validar request (IDs no vacíos, cantidad <= MaxBulkVerifyIDs, operador presente).
//  2. Cargar productos existentes por IDs.
//  3. Clasificar en: faltantes, ya en estado destino, y a actualizar.
//  4. Aplicar updates atómicamente (snapshot + UPDATEs en tx).
//  5. Guardar audit inmutable.
//  6. Emitir log canónico ADR-001.
func (uc *BulkVerifyGlobalProducts) Execute(ctx context.Context, req BulkVerifyGlobalProductsRequest) (*BulkVerifyGlobalProductsResponse, error) {
	if err := uc.validateRequest(req); err != nil {
		return nil, err
	}

	mode := value_object.BulkVerifyModeUnverify
	if req.Verify {
		mode = value_object.BulkVerifyModeVerify
	}

	products, err := uc.globalProductRepo.FindByIDs(ctx, req.IDs)
	if err != nil {
		return nil, exception.NewInternalError("Error al cargar productos del lote", err)
	}

	productByID := make(map[string]*entity.GlobalProduct, len(products))
	for _, p := range products {
		productByID[p.IDString()] = p
	}

	var missingIDs, toUpdateIDs []string
	skippedCount := 0

	for _, id := range req.IDs {
		p, ok := productByID[id]
		if !ok {
			missingIDs = append(missingIDs, id)
			continue
		}
		if p.IsVerified() == req.Verify {
			skippedCount++
			continue
		}
		toUpdateIDs = append(toUpdateIDs, id)
	}

	summary := value_object.BulkVerifySummary{
		TotalRequested: len(req.IDs),
		Skipped:        skippedCount,
		Failed:         len(missingIDs),
	}

	var snapshotRef *string
	affected := 0

	if len(toUpdateIDs) > 0 {
		// Incluimos nanosegundos para evitar colisiones si se ejecutan varias
		// operaciones dentro del mismo segundo (caso común en tests y en UI rápida).
		now := time.Now()
		snapshotName := fmt.Sprintf("%s_%09d", now.Format("20060102_150405"), now.Nanosecond())
		affected, err = uc.bulkVerifyRepo.ApplyInTransaction(ctx, snapshotName, toUpdateIDs, mode)
		if err != nil {
			uc.logEvent(port.PIMEvent{
				Event:  "catalog.bulk_verify_failed",
				UserID: req.OperatorID,
				Count:  len(toUpdateIDs),
				Reason: err.Error(),
			})
			return nil, exception.NewInternalError("Error al aplicar verificación en lote", err)
		}
		ref := fmt.Sprintf("global_products_bkp_bulkverify_%s", snapshotName)
		snapshotRef = &ref
	}

	summary.Processed = affected

	audit := value_object.BulkVerifyAuditRow{
		OperatorID:    req.OperatorID,
		ExecutedAt:    time.Now(),
		Mode:          mode,
		RequestIDs:    req.IDs,
		SnapshotRef:   snapshotRef,
		Summary:       summary,
		AffectedCount: affected,
	}
	if auditErr := uc.bulkVerifyRepo.SaveAudit(ctx, audit); auditErr != nil {
		uc.logEvent(port.PIMEvent{
			Event:  "catalog.bulk_verify_audit_failed",
			UserID: req.OperatorID,
			Reason: auditErr.Error(),
		})
	}

	uc.logEvent(port.PIMEvent{
		Event:  "catalog.bulk_verify_completed",
		UserID: req.OperatorID,
		Count:  affected,
		Reason: string(mode),
	})

	return &BulkVerifyGlobalProductsResponse{
		Mode:        mode,
		SnapshotRef: snapshotRef,
		Summary:     summary,
	}, nil
}

func (uc *BulkVerifyGlobalProducts) validateRequest(req BulkVerifyGlobalProductsRequest) error {
	if len(req.IDs) == 0 {
		return exception.NewValidationError("IDs es obligatorio y no puede estar vacío", nil)
	}
	if len(req.IDs) > MaxBulkVerifyIDs {
		return exception.NewValidationError(
			fmt.Sprintf("se permite un máximo de %d IDs por operación", MaxBulkVerifyIDs),
			nil,
		)
	}
	if req.OperatorID == "" {
		return exception.NewValidationError("OperatorID es obligatorio", nil)
	}
	return nil
}

func (uc *BulkVerifyGlobalProducts) logEvent(e port.PIMEvent) {
	if uc.logger != nil {
		uc.logger.Log(e)
	}
}
