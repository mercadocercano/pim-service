package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/category/application/request"
	"saas-mt-pim-service/src/category/application/response"
	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/port"
)

// MapTenantCategoryUseCase maneja el mapeo de categorías tenant a categorías marketplace
type MapTenantCategoryUseCase struct {
	categoryRepo      port.MarketplaceCategoryRepository
	tenantMappingRepo port.TenantCategoryMappingRepository
}

// NewMapTenantCategoryUseCase crea una nueva instancia del caso de uso
func NewMapTenantCategoryUseCase(
	categoryRepo port.MarketplaceCategoryRepository,
	tenantMappingRepo port.TenantCategoryMappingRepository,
) *MapTenantCategoryUseCase {
	return &MapTenantCategoryUseCase{
		categoryRepo:      categoryRepo,
		tenantMappingRepo: tenantMappingRepo,
	}
}

// Execute ejecuta el caso de uso de mapeo de categoría tenant
func (uc *MapTenantCategoryUseCase) Execute(
	ctx context.Context,
	req *request.MapTenantCategoryRequest,
	tenantID string,
) (*response.TenantCategoryMappingResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Verificar que la categoría marketplace existe
	marketplaceCategory, err := uc.categoryRepo.GetByID(ctx, req.MarketplaceCategoryID)
	if err != nil {
		return nil, fmt.Errorf("marketplace category not found: %w", err)
	}

	// Verificar si ya existe un mapeo para este tenant y categoría
	existingMapping, err := uc.tenantMappingRepo.GetByTenantAndMarketplaceCategory(
		ctx, tenantID, req.MarketplaceCategoryID,
	)
	if err == nil && existingMapping != nil {
		return nil, fmt.Errorf("mapping already exists for tenant %s and marketplace category %s",
			tenantID, req.MarketplaceCategoryID)
	}

	// Crear el mapeo usando el constructor de la entidad
	customName := &req.CustomName
	mapping, err := entity.NewTenantCategoryMapping(
		tenantID,
		req.CategoryID,
		req.MarketplaceCategoryID,
		customName,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tenant category mapping: %w", err)
	}

	// Guardar el mapeo
	if err := uc.tenantMappingRepo.Save(ctx, mapping); err != nil {
		return nil, fmt.Errorf("failed to save tenant category mapping: %w", err)
	}

	// Convertir a respuesta
	return &response.TenantCategoryMappingResponse{
		ID:                      mapping.ID,
		TenantID:                mapping.TenantID,
		MarketplaceCategoryID:   mapping.MarketplaceCategoryID,
		MarketplaceCategoryName: marketplaceCategory.Name,
		CustomName:              req.CustomName,
		IsActive:                true, // Por defecto activo
		CreatedAt:               mapping.CreatedAt,
		UpdatedAt:               mapping.UpdatedAt,
	}, nil
}
