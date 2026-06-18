package usecase

import "saas-mt-pim-service/src/product/global_catalog/domain/value_object"

// ReclassifyRequest es el input del use case ReclassifyBusinessTypesUseCase.
type ReclassifyRequest struct {
	// DryRun true (default) simula la operación sin mutar global_products.
	// El summary devuelto es idéntico al que produciría el apply real (invariante ADR-005 §8).
	DryRun bool

	// Confirm must be true para que un apply (DryRun=false) proceda.
	// apply requiere DryRun=false AND Confirm=true; de lo contrario, no muta.
	Confirm bool

	// Scope define los criterios de selección y el cap máximo de filas.
	Scope value_object.ReclassifyScope

	// OperatorID es el ID del operador humano propagado via header X-Operator-Id.
	// En L3: se almacena en el audit aunque esté vacío.
	// En L4: rechazado con 400 si está vacío en un apply real.
	OperatorID string
}

// ReclassifyResponse es el output del use case.
type ReclassifyResponse struct {
	// Mode indica si fue simulación ("dry_run") o aplicación real ("applied").
	Mode string `json:"mode"`

	// SnapshotRef es el nombre de la tabla de backup creada durante el apply.
	// nil en dry_run (no se crea snapshot).
	SnapshotRef *string `json:"snapshot_ref"`

	// Summary es el resumen estructurado de la operación.
	Summary value_object.ReclassifySummary `json:"summary"`

	// AntesDespues es una muestra acotada del detalle antes/después por fila.
	AntesDespues []value_object.AntesDespuesRow `json:"antes_despues"`
}
