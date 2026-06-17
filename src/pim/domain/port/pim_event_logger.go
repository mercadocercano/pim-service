package port

// PIMEvent es el payload canónico para eventos de dominio del servicio PIM (ADR-001).
// Campos flat, named. Los nombres comunes (tenant_id, user_id) son idénticos al resto
// de la flota para que el LogQL cross-service funcione. Todos opcionales salvo Event.
//
// Bounded contexts cubiertos: catalog (product/global-catalog), quickstart, s2s/backfill.
// Naming: <domain>.<action>_<result> — verbos en pasado; _failed sólo cuando hay fallo.
type PIMEvent struct {
	Event     string // ej: "catalog.product_created", "pim.quickstart_completed"
	TenantID  string
	UserID    string
	ProductID string
	SKU       string
	JobID     string
	Count     int
	Reason    string
}

// PIMEventLogger es el puerto para emitir eventos canónicos del dominio PIM.
// El código de aplicación depende de esta interfaz; el adapter (JSON a stdout, etc.)
// la implementa. Nunca al revés.
type PIMEventLogger interface {
	Log(e PIMEvent)
}
