package value_object

import "time"

// SkipKind clasifica por qué un candidato fue descartado (no actualizado).
type SkipKind string

const (
	// SkipNoResuelve: el resolver no encontró rubro para la categoría del producto.
	SkipNoResuelve SkipKind = "no_resuelve"

	// SkipYaEspecifico: el producto ya tiene un rubro específico (distinto de almacen y no nil).
	// Invariante: nunca se toca un producto en rubro específico. Solo se rellena nil o se mueve desde almacen.
	SkipYaEspecifico SkipKind = "ya_especifico"

	// SkipYaCorrecto: el resolver resolvió el mismo rubro que ya tiene el producto.
	SkipYaCorrecto SkipKind = "ya_correcto"

	// SkipColision: el UPDATE violaría UNIQUE(name, business_type) → skip individual,
	// nunca aborta el lote (ADR-005 §8, TEST-ID T-022).
	SkipColision SkipKind = "colision"
)

// ReclassifyCandidate representa un producto evaluable para re-clasificación.
// Candidatos son los que tienen business_type IS NULL o business_type = 'almacen'.
type ReclassifyCandidate struct {
	ID           string
	Name         string
	Category     string
	BusinessType *string // nil = no asignado; "almacen" = relleno genérico previo
}

// ReclassifyUpdate representa un producto que SÍ debe actualizarse.
type ReclassifyUpdate struct {
	ID          string
	Name        string
	FromType    *string  // valor anterior (nil o "almacen")
	ToType      string   // valor nuevo resuelto por el resolver
	Kind        string   // "relleno" (desde nil/vacío) o "correccion" (desde almacen → específico)
}

// ReclassifySkip representa un producto que NO se actualizó, con el motivo.
type ReclassifySkip struct {
	ID   string
	Name string
	Kind SkipKind
}

// UpdatesByRubro agrupa los updates por código de rubro destino.
// Claves: código del rubro (ej: "fiambreria", "limpieza").
// Valores: conteo de "relleno" y "correccion".
type UpdatesByRubro map[string]RubroUpdateCount

// RubroUpdateCount desglosa los updates de un rubro.
type RubroUpdateCount struct {
	Relleno   int `json:"relleno"`    // nil/vacío → rubro específico
	Correccion int `json:"correccion"` // almacen → rubro específico
}

// ReclassifySkipSummary agrega los skips por tipo.
type ReclassifySkipSummary struct {
	NoResuelve   int `json:"no_resuelve"`
	YaEspecifico int `json:"ya_especifico"`
	YaCorrecto   int `json:"ya_correcto"`
}

// ReclassifySummary es el resumen estructurado de una operación de re-clasificación.
// Formato idéntico en dry_run y apply (ADR-005 §2, invariante T-017).
type ReclassifySummary struct {
	TotalEvaluados     int                   `json:"total_evaluados"`
	Candidatos         int                   `json:"candidatos"`
	UpdatesPorRubro    UpdatesByRubro         `json:"updates_por_rubro"`
	ColisionesSkipeadas int                  `json:"colisiones_skipeadas"`
	Skips              ReclassifySkipSummary `json:"skips"`
}

// AntesDespuesRow es una fila del detalle before/after por producto (muestra acotada).
type AntesDespuesRow struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	From *string  `json:"from"`
	To   string   `json:"to"`
	Kind SkipKind `json:"kind"` // reutiliza el tipo, valor "update" para los efectivos
}

// ReclassifyAuditRow es el registro que se persiste en global_product_reclassification_audit.
type ReclassifyAuditRow struct {
	OperatorID    string
	ExecutedAt    time.Time
	Mode          string // "dry_run" | "applied"
	Scope         ReclassifyScope
	SnapshotRef   *string
	Summary       ReclassifySummary
	AffectedCount int
}
