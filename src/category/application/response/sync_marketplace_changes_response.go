package response

import "time"

// SyncMarketplaceChangesResponse representa la respuesta de sincronización de cambios
type SyncMarketplaceChangesResponse struct {
	SyncID                string        `json:"sync_id"`
	TenantID              *string       `json:"tenant_id"`
	MarketplaceCategoryID *string       `json:"marketplace_category_id"`
	IsDryRun              bool          `json:"is_dry_run"`
	SyncStatus            string        `json:"sync_status"` // success, partial, failed
	TotalChanges          int           `json:"total_changes"`
	AppliedChanges        int           `json:"applied_changes"`
	FailedChanges         int           `json:"failed_changes"`
	SkippedChanges        int           `json:"skipped_changes"`
	SyncResults           []SyncResult  `json:"sync_results"`
	AffectedTenants       []string      `json:"affected_tenants"`
	ExecutionTime         time.Duration `json:"execution_time"`
	StartedAt             time.Time     `json:"started_at"`
	CompletedAt           time.Time     `json:"completed_at"`
	Errors                []SyncError   `json:"errors,omitempty"`
	Warnings              []SyncWarning `json:"warnings,omitempty"`
}

// SyncResult representa el resultado de una operación de sincronización específica
type SyncResult struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`      // mapping, attribute, category
	Operation   string                 `json:"operation"` // create, update, delete
	EntityID    string                 `json:"entity_id"`
	TenantID    string                 `json:"tenant_id"`
	Status      string                 `json:"status"` // success, failed, skipped
	Message     string                 `json:"message"`
	OldValue    map[string]interface{} `json:"old_value,omitempty"`
	NewValue    map[string]interface{} `json:"new_value,omitempty"`
	ProcessedAt time.Time              `json:"processed_at"`
}

// SyncError representa un error durante la sincronización
type SyncError struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	EntityID  string `json:"entity_id,omitempty"`
	TenantID  string `json:"tenant_id,omitempty"`
	Operation string `json:"operation,omitempty"`
	Details   string `json:"details,omitempty"`
}

// SyncWarning representa una advertencia durante la sincronización
type SyncWarning struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	EntityID   string `json:"entity_id,omitempty"`
	TenantID   string `json:"tenant_id,omitempty"`
	Suggestion string `json:"suggestion,omitempty"`
}

// SyncSummaryResponse representa un resumen de múltiples sincronizaciones
type SyncSummaryResponse struct {
	TotalSyncs          int                               `json:"total_syncs"`
	SuccessfulSyncs     int                               `json:"successful_syncs"`
	FailedSyncs         int                               `json:"failed_syncs"`
	PartialSyncs        int                               `json:"partial_syncs"`
	TotalChangesApplied int                               `json:"total_changes_applied"`
	SyncHistory         []*SyncMarketplaceChangesResponse `json:"sync_history"`
	LastSyncAt          *time.Time                        `json:"last_sync_at"`
	NextScheduledSync   *time.Time                        `json:"next_scheduled_sync,omitempty"`
	GeneratedAt         time.Time                         `json:"generated_at"`
}
