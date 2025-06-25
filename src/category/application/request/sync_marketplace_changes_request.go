package request

import (
	"errors"
	"strings"
)

// SyncMarketplaceChangesRequest representa la petición para sincronizar cambios del marketplace
type SyncMarketplaceChangesRequest struct {
	TenantID              *string           `json:"tenant_id,omitempty"`               // null = todos los tenants
	MarketplaceCategoryID *string           `json:"marketplace_category_id,omitempty"` // null = todas las categorías
	ChangeTypes           []string          `json:"change_types,omitempty"`            // tipos de cambios a sincronizar
	ForceSync             bool              `json:"force_sync"`                        // forzar sincronización aunque no haya cambios
	DryRun                bool              `json:"dry_run"`                           // solo simular, no aplicar cambios
	SyncOptions           SyncOptionsConfig `json:"sync_options"`                      // opciones de sincronización
}

// SyncOptionsConfig representa las opciones de configuración para la sincronización
type SyncOptionsConfig struct {
	UpdateMappings        bool `json:"update_mappings"`         // actualizar mapeos existentes
	CreateMissingMappings bool `json:"create_missing_mappings"` // crear mapeos faltantes
	RemoveOrphanMappings  bool `json:"remove_orphan_mappings"`  // remover mapeos huérfanos
	UpdateAttributes      bool `json:"update_attributes"`       // sincronizar atributos
	NotifyTenants         bool `json:"notify_tenants"`          // notificar a tenants de cambios
}

// Validate valida los datos de la petición
func (r *SyncMarketplaceChangesRequest) Validate() error {
	// Validar tenant_id si se proporciona
	if r.TenantID != nil && strings.TrimSpace(*r.TenantID) == "" {
		return errors.New("tenant_id cannot be empty string")
	}

	// Validar marketplace_category_id si se proporciona
	if r.MarketplaceCategoryID != nil && strings.TrimSpace(*r.MarketplaceCategoryID) == "" {
		return errors.New("marketplace_category_id cannot be empty string")
	}

	// Validar change_types
	validChangeTypes := map[string]bool{
		"category_created":  true,
		"category_updated":  true,
		"category_deleted":  true,
		"category_moved":    true,
		"attribute_created": true,
		"attribute_updated": true,
		"attribute_deleted": true,
		"hierarchy_changed": true,
	}

	for _, changeType := range r.ChangeTypes {
		if !validChangeTypes[changeType] {
			return errors.New("invalid change_type: " + changeType)
		}
	}

	// Si no se especifican tipos de cambio, usar todos por defecto
	if len(r.ChangeTypes) == 0 {
		r.ChangeTypes = []string{
			"category_created", "category_updated", "category_deleted", "category_moved",
			"attribute_created", "attribute_updated", "attribute_deleted", "hierarchy_changed",
		}
	}

	// Validar que al menos una opción de sincronización esté habilitada
	if !r.SyncOptions.UpdateMappings &&
		!r.SyncOptions.CreateMissingMappings &&
		!r.SyncOptions.RemoveOrphanMappings &&
		!r.SyncOptions.UpdateAttributes {
		return errors.New("at least one sync option must be enabled")
	}

	return nil
}
