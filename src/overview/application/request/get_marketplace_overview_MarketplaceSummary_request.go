package request

// GetMarketplaceOverviewRequest representa la petición para obtener overview del marketplace
type GetMarketplaceOverviewRequest struct {
	// Sections especifica qué secciones incluir en el overview
	Sections []string `json:"sections" form:"sections"`

	// IncludeStats indica si incluir estadísticas detalladas
	IncludeStats bool `json:"include_stats" form:"include_stats"`

	// TimeRangeDays especifica el rango de días para métricas temporales
	TimeRangeDays int `json:"time_range_days" form:"time_range_days"`

	// Limit especifica el límite de elementos para listas (top items, recent items, etc.)
	Limit int `json:"limit" form:"limit"`

	// Parallel indica si ejecutar las consultas en paralelo
	Parallel bool `json:"parallel" form:"parallel"`

	// TenantID específico para filtros (opcional - para marketplace admin es "global")
	TenantID string `json:"tenant_id" form:"tenant_id"`
}
