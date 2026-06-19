package value_object

import "time"

// Fuentes de resolución de un slug (de dónde salió el category_slug computado).
const (
	SourceOverride     = "override"       // tabla category_slug_overrides (gana sobre el resolver)
	SourceResolver     = "resolver"       // resolver determinístico de go-shared
	SourceUnclassified = "sin-clasificar" // no resolvió → catch-all
)

// RawCategoryCount es una categoría cruda distinta con la cantidad de productos que la usan.
// RawCategory es puntero porque global_products.category puede ser NULL.
type RawCategoryCount struct {
	RawCategory  *string
	ProductCount int
	// CurrentSlug es el category_slug actual si TODAS las filas de esta categoría lo comparten;
	// nil si es NULL o está mezclado (en cuyo caso se fuerza recálculo).
	CurrentSlug *string
}

// CategoryMapping es el resultado de resolver una categoría cruda a su slug normalizado.
type CategoryMapping struct {
	RawCategory  *string // nil = filas con category NULL
	ProductCount int
	CurrentSlug  *string // category_slug actual (nil si nunca se normalizó)
	NewSlug      string  // slug computado (override > resolver > sin-clasificar)
	Source       string  // SourceOverride | SourceResolver | SourceUnclassified
	Changed      bool    // NewSlug != CurrentSlug
}

// NormalizeSummary es el resumen estructurado de la operación (dry_run y apply son idénticos,
// salvo el modo). El detalle por-categoría de los sin-clasificar es el worklist de overrides.
type NormalizeSummary struct {
	TotalCategorias      int            `json:"total_categorias"`       // categorías distintas evaluadas
	TotalProductos       int            `json:"total_productos"`        // suma de productos cubiertos
	CategoriasPorFuente  map[string]int `json:"categorias_por_fuente"`  // {override, resolver, sin-clasificar}
	ProductosPorFuente   map[string]int `json:"productos_por_fuente"`   // idem, ponderado por productos
	ProductosAfectados   int            `json:"productos_afectados"`    // productos cuyo slug cambia
	TopSlugs             map[string]int `json:"top_slugs"`              // slug destino -> productos
	SinClasificarMuestra []string       `json:"sin_clasificar_muestra"` // categorías crudas sin resolver (worklist)
}

// NormalizeMappingSample es una fila de la muestra antes/después de la response.
type NormalizeMappingSample struct {
	RawCategory  *string `json:"raw_category"`
	ProductCount int     `json:"product_count"`
	From         *string `json:"from"`
	To           string  `json:"to"`
	Source       string  `json:"source"`
}

// NormalizeAuditRow es el registro persistido en global_product_reclassification_audit
// (tabla reusada, ADR-007 §4). El campo scope/summary JSONB distinguen la operación.
type NormalizeAuditRow struct {
	OperatorID    string
	ExecutedAt    time.Time
	Mode          string // dry_run | applied
	Scope         NormalizeCategoryScope
	SnapshotRef   *string
	Summary       NormalizeSummary
	AffectedCount int
}
