package logging

import (
	"io"

	"saas-mt-pim-service/src/pim/domain/port"

	sharedlog "github.com/hornosg/go-shared/infrastructure/logging"
)

// PIMLogger implementa port.PIMEventLogger emitiendo una línea JSON canónica (ADR-001)
// por evento, delegando el envelope (ts/level/service/event + campos flat omitempty) en
// go-shared CanonicalLogger (>= v0.8.0). El mapeo struct→fields y las reglas de nivel
// por evento viven acá; el formato canónico es compartido por la flota.
type PIMLogger struct {
	canonical *sharedlog.CanonicalLogger
}

// NewPIMLogger crea el adapter escribiendo a stdout. El service se fija acá, nunca por-call.
func NewPIMLogger(service string) *PIMLogger {
	return &PIMLogger{canonical: sharedlog.NewCanonicalLogger(service)}
}

// NewPIMLoggerWithWriter permite inyectar un io.Writer (tests).
func NewPIMLoggerWithWriter(service string, w io.Writer) *PIMLogger {
	return &PIMLogger{canonical: sharedlog.NewCanonicalLoggerWithWriter(service, w)}
}

// levelFor aplica las reglas de nivel del ADR-001 por tipo de evento.
//
//   - info  → outcome OK del dominio (producto creado, import completado, etc.)
//   - warn  → anomalía recuperable (backfill error, template refresh sin resultados)
//   - error → estado inesperado / violación de invariante (import fallido, persistencia fallida)
func levelFor(event string) string {
	switch event {
	// catalog outcomes OK
	case "catalog.product_created",
		"catalog.product_updated",
		"catalog.product_deleted",
		"catalog.global_product_created",
		"catalog.global_product_updated",
		"catalog.global_product_deleted":
		return "info"
	// quickstart outcomes OK
	case "pim.quickstart_completed",
		"pim.import_from_global_catalog_completed",
		"pim.import_from_business_type_completed",
		"pim.template_refresh_completed",
		"pim.backfill_completed":
		return "info"
	// outcomes fallidos — error de negocio o persistencia
	case "pim.import_failed",
		"pim.quickstart_failed":
		return "error"
	// anomalías recuperables
	case "pim.backfill_product_error",
		"pim.backfill_tenant_error",
		"pim.template_status_error":
		return "warn"
	default:
		return "info"
	}
}

// Log implementa port.PIMEventLogger.
func (l *PIMLogger) Log(e port.PIMEvent) {
	fields := map[string]any{
		"tenant_id":  e.TenantID,
		"user_id":    e.UserID,
		"product_id": e.ProductID,
		"sku":        e.SKU,
		"job_id":     e.JobID,
		"reason":     e.Reason,
	}
	if e.Count > 0 {
		fields["count"] = e.Count
	}
	l.canonical.Emit(levelFor(e.Event), e.Event, fields)
}
