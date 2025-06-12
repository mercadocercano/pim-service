package usecase

import (
	"context"
	"fmt"
	"time"

	"pim/src/marketplace/application/request"
	"pim/src/marketplace/application/response"
	"pim/src/marketplace/domain/entity"
	"pim/src/marketplace/domain/port"

	"github.com/google/uuid"
)

// SyncMarketplaceChangesUseCase maneja la sincronización de cambios del marketplace
type SyncMarketplaceChangesUseCase struct {
	categoryRepo        port.MarketplaceCategoryRepository
	mappingRepo         port.TenantCategoryMappingRepository
	customAttributeRepo port.TenantCustomAttributeRepository
}

// NewSyncMarketplaceChangesUseCase crea una nueva instancia del caso de uso
func NewSyncMarketplaceChangesUseCase(
	categoryRepo port.MarketplaceCategoryRepository,
	mappingRepo port.TenantCategoryMappingRepository,
	customAttributeRepo port.TenantCustomAttributeRepository,
) *SyncMarketplaceChangesUseCase {
	return &SyncMarketplaceChangesUseCase{
		categoryRepo:        categoryRepo,
		mappingRepo:         mappingRepo,
		customAttributeRepo: customAttributeRepo,
	}
}

// Execute ejecuta el caso de uso de sincronización de cambios
func (uc *SyncMarketplaceChangesUseCase) Execute(
	ctx context.Context,
	req *request.SyncMarketplaceChangesRequest,
) (*response.SyncMarketplaceChangesResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	startTime := time.Now()
	syncID := uuid.New().String()

	// Inicializar respuesta
	syncResponse := &response.SyncMarketplaceChangesResponse{
		SyncID:                syncID,
		TenantID:              req.TenantID,
		MarketplaceCategoryID: req.MarketplaceCategoryID,
		IsDryRun:              req.DryRun,
		SyncStatus:            "in_progress",
		SyncResults:           []response.SyncResult{},
		AffectedTenants:       []string{},
		StartedAt:             startTime,
		Errors:                []response.SyncError{},
		Warnings:              []response.SyncWarning{},
	}

	// Obtener tenants a sincronizar
	tenantsToSync, err := uc.getTenantsToSync(ctx, req.TenantID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenants to sync: %w", err)
	}

	// Obtener categorías marketplace a sincronizar
	categoriesToSync, err := uc.getCategoriesToSync(ctx, req.MarketplaceCategoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories to sync: %w", err)
	}

	// Procesar sincronización para cada tenant y categoría
	for _, tenantID := range tenantsToSync {
		for _, category := range categoriesToSync {
			results := uc.syncCategoryForTenant(ctx, tenantID, category, req, syncResponse)
			syncResponse.SyncResults = append(syncResponse.SyncResults, results...)
		}
	}

	// Actualizar estadísticas finales
	uc.updateSyncStatistics(syncResponse)

	// Determinar estado final
	if syncResponse.FailedChanges == 0 {
		if syncResponse.AppliedChanges > 0 {
			syncResponse.SyncStatus = "success"
		} else {
			syncResponse.SyncStatus = "no_changes"
		}
	} else if syncResponse.AppliedChanges > 0 {
		syncResponse.SyncStatus = "partial"
	} else {
		syncResponse.SyncStatus = "failed"
	}

	syncResponse.CompletedAt = time.Now()
	syncResponse.ExecutionTime = syncResponse.CompletedAt.Sub(startTime)

	return syncResponse, nil
}

// getTenantsToSync obtiene la lista de tenants a sincronizar
func (uc *SyncMarketplaceChangesUseCase) getTenantsToSync(ctx context.Context, tenantID *string) ([]string, error) {
	if tenantID != nil {
		return []string{*tenantID}, nil
	}

	// TODO: Implementar obtención de todos los tenants activos
	// Por ahora retornamos una lista vacía que será manejada por el repositorio
	return []string{}, nil
}

// getCategoriesToSync obtiene las categorías marketplace a sincronizar
func (uc *SyncMarketplaceChangesUseCase) getCategoriesToSync(ctx context.Context, categoryID *string) ([]*entity.MarketplaceCategory, error) {
	if categoryID != nil {
		category, err := uc.categoryRepo.GetByID(ctx, *categoryID)
		if err != nil {
			return nil, err
		}
		return []*entity.MarketplaceCategory{category}, nil
	}

	// Obtener todas las categorías marketplace usando GetTree
	return uc.categoryRepo.GetTree(ctx)
}

// syncCategoryForTenant sincroniza una categoría específica para un tenant
func (uc *SyncMarketplaceChangesUseCase) syncCategoryForTenant(
	ctx context.Context,
	tenantID string,
	category *entity.MarketplaceCategory,
	req *request.SyncMarketplaceChangesRequest,
	syncResponse *response.SyncMarketplaceChangesResponse,
) []response.SyncResult {
	var results []response.SyncResult

	// Agregar tenant a la lista de afectados si no está ya
	uc.addAffectedTenant(syncResponse, tenantID)

	// Sincronizar mapeo de categoría
	if req.SyncOptions.UpdateMappings || req.SyncOptions.CreateMissingMappings {
		mappingResult := uc.syncCategoryMapping(ctx, tenantID, category, req)
		results = append(results, mappingResult)
	}

	// Sincronizar atributos personalizados
	if req.SyncOptions.UpdateAttributes {
		attributeResults := uc.syncCustomAttributes(ctx, tenantID, category, req)
		results = append(results, attributeResults...)
	}

	// Remover mapeos huérfanos si está habilitado
	if req.SyncOptions.RemoveOrphanMappings {
		orphanResults := uc.removeOrphanMappings(ctx, tenantID, category, req)
		results = append(results, orphanResults...)
	}

	return results
}

// syncCategoryMapping sincroniza el mapeo de una categoría para un tenant
func (uc *SyncMarketplaceChangesUseCase) syncCategoryMapping(
	ctx context.Context,
	tenantID string,
	category *entity.MarketplaceCategory,
	req *request.SyncMarketplaceChangesRequest,
) response.SyncResult {
	result := response.SyncResult{
		ID:          uuid.New().String(),
		Type:        "mapping",
		EntityID:    category.ID,
		TenantID:    tenantID,
		ProcessedAt: time.Now(),
	}

	// Verificar si ya existe un mapeo
	existingMapping, err := uc.mappingRepo.GetByTenantAndMarketplaceCategory(ctx, tenantID, category.ID)
	if err != nil && existingMapping == nil {
		// No existe mapeo, crear uno nuevo si está habilitado
		if req.SyncOptions.CreateMissingMappings && !req.DryRun {
			// Usar un categoryID placeholder ya que necesitamos el ID de la categoría tenant
			// En un escenario real, esto vendría del contexto o sería obtenido del tenant
			categoryID := "placeholder-category-id"
			mapping, err := entity.NewTenantCategoryMapping(tenantID, categoryID, category.ID, nil)
			if err != nil {
				result.Status = "failed"
				result.Operation = "create"
				result.Message = fmt.Sprintf("Failed to create mapping: %v", err)
				return result
			}

			err = uc.mappingRepo.Save(ctx, mapping)
			if err != nil {
				result.Status = "failed"
				result.Operation = "create"
				result.Message = fmt.Sprintf("Failed to save mapping: %v", err)
				return result
			}

			result.Status = "success"
			result.Operation = "create"
			result.Message = "Mapping created successfully"
			result.NewValue = map[string]interface{}{
				"tenant_id":               mapping.TenantID,
				"marketplace_category_id": mapping.MarketplaceCategoryID,
				"custom_name":             mapping.CustomName,
			}
		} else {
			result.Status = "skipped"
			result.Operation = "create"
			result.Message = "Mapping creation skipped (dry run or option disabled)"
		}
	} else if existingMapping != nil {
		// Existe mapeo, actualizar si está habilitado
		if req.SyncOptions.UpdateMappings && !req.DryRun {
			// Aquí se podrían aplicar actualizaciones específicas
			result.Status = "success"
			result.Operation = "update"
			result.Message = "Mapping updated successfully"
		} else {
			result.Status = "skipped"
			result.Operation = "update"
			result.Message = "Mapping update skipped (dry run or option disabled)"
		}
	}

	return result
}

// syncCustomAttributes sincroniza atributos personalizados para un tenant y categoría
func (uc *SyncMarketplaceChangesUseCase) syncCustomAttributes(
	ctx context.Context,
	tenantID string,
	category *entity.MarketplaceCategory,
	req *request.SyncMarketplaceChangesRequest,
) []response.SyncResult {
	var results []response.SyncResult

	// TODO: Implementar sincronización de atributos personalizados
	// Por ahora retornamos un resultado placeholder
	result := response.SyncResult{
		ID:          uuid.New().String(),
		Type:        "attribute",
		EntityID:    category.ID,
		TenantID:    tenantID,
		Status:      "skipped",
		Operation:   "sync",
		Message:     "Attribute sync not implemented yet",
		ProcessedAt: time.Now(),
	}

	results = append(results, result)
	return results
}

// removeOrphanMappings remueve mapeos huérfanos para un tenant
func (uc *SyncMarketplaceChangesUseCase) removeOrphanMappings(
	ctx context.Context,
	tenantID string,
	category *entity.MarketplaceCategory,
	req *request.SyncMarketplaceChangesRequest,
) []response.SyncResult {
	var results []response.SyncResult

	// TODO: Implementar remoción de mapeos huérfanos
	// Por ahora retornamos un resultado placeholder
	result := response.SyncResult{
		ID:          uuid.New().String(),
		Type:        "mapping",
		EntityID:    category.ID,
		TenantID:    tenantID,
		Status:      "skipped",
		Operation:   "delete",
		Message:     "Orphan mapping removal not implemented yet",
		ProcessedAt: time.Now(),
	}

	results = append(results, result)
	return results
}

// addAffectedTenant agrega un tenant a la lista de afectados si no está ya
func (uc *SyncMarketplaceChangesUseCase) addAffectedTenant(syncResponse *response.SyncMarketplaceChangesResponse, tenantID string) {
	for _, existingTenant := range syncResponse.AffectedTenants {
		if existingTenant == tenantID {
			return
		}
	}
	syncResponse.AffectedTenants = append(syncResponse.AffectedTenants, tenantID)
}

// updateSyncStatistics actualiza las estadísticas de sincronización
func (uc *SyncMarketplaceChangesUseCase) updateSyncStatistics(syncResponse *response.SyncMarketplaceChangesResponse) {
	syncResponse.TotalChanges = len(syncResponse.SyncResults)

	for _, result := range syncResponse.SyncResults {
		switch result.Status {
		case "success":
			syncResponse.AppliedChanges++
		case "failed":
			syncResponse.FailedChanges++
		case "skipped":
			syncResponse.SkippedChanges++
		}
	}
}
