package response

import "time"

// GetMarketplaceOverviewResponse representa la respuesta para obtener overview del marketplace
type GetMarketplaceOverviewResponse struct {
	// Success indica si la operación fue exitosa
	Success bool `json:"success"`

	// Message mensaje descriptivo (opcional)
	Message string `json:"message,omitempty"`

	// Data contiene los datos del overview organizados por sección
	Data map[string]interface{} `json:"data"`

	// GeneratedAt timestamp de cuando se generó la respuesta
	GeneratedAt time.Time `json:"generated_at"`

	// Metadata contiene información sobre la consulta ejecutada
	Metadata OverviewMetadata `json:"metadata"`
}

// OverviewMetadata contiene metadata sobre la ejecución del overview
type OverviewMetadata struct {
	// ParallelExecution indica si se ejecutó en paralelo
	ParallelExecution bool `json:"parallel_execution"`

	// Sections secciones que fueron procesadas
	Sections []string `json:"sections"`

	// TotalTime tiempo total de ejecución
	TotalTime time.Duration `json:"total_time"`

	// CacheUsed indica si se usó información de cache
	CacheUsed bool `json:"cache_used,omitempty"`
}
