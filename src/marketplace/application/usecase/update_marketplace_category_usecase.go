package usecase

import (
	"context"
	"fmt"
	"time"

	"pim/src/marketplace/application/request"
	"pim/src/marketplace/application/response"
	"pim/src/marketplace/domain/port"
)

// UpdateMarketplaceCategoryUseCase maneja la actualización de categorías marketplace
type UpdateMarketplaceCategoryUseCase struct {
	categoryRepo port.MarketplaceCategoryRepository
}

// NewUpdateMarketplaceCategoryUseCase crea una nueva instancia del caso de uso
func NewUpdateMarketplaceCategoryUseCase(
	categoryRepo port.MarketplaceCategoryRepository,
) *UpdateMarketplaceCategoryUseCase {
	return &UpdateMarketplaceCategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

// Execute ejecuta el caso de uso de actualización de categoría marketplace
func (uc *UpdateMarketplaceCategoryUseCase) Execute(
	ctx context.Context,
	categoryID string,
	req *request.UpdateMarketplaceCategoryRequest,
) (*response.MarketplaceCategoryResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Obtener la categoría existente
	category, err := uc.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	// Verificar que el slug no esté en uso por otra categoría (si se está cambiando)
	if req.Slug != nil && *req.Slug != category.Slug {
		exists, err := uc.categoryRepo.ExistsBySlug(ctx, *req.Slug)
		if err != nil {
			return nil, fmt.Errorf("error checking slug uniqueness: %w", err)
		}
		if exists {
			return nil, fmt.Errorf("category with slug '%s' already exists", *req.Slug)
		}
	}

	// Actualizar campos si se proporcionan
	if req.Name != nil {
		category.Name = *req.Name
	}

	if req.Slug != nil {
		category.Slug = *req.Slug
	}

	if req.Description != nil {
		category.Description = *req.Description
	}

	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	if req.SortOrder != nil {
		category.UpdateSortOrder(*req.SortOrder)
	}

	// Actualizar timestamp
	category.UpdatedAt = time.Now()

	// Guardar la categoría actualizada
	if err := uc.categoryRepo.Save(ctx, category); err != nil {
		return nil, fmt.Errorf("failed to update marketplace category: %w", err)
	}

	// Convertir a respuesta
	var responseDescription *string
	if category.Description != "" {
		responseDescription = &category.Description
	}

	return &response.MarketplaceCategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: responseDescription,
		ParentID:    category.ParentID,
		Level:       category.Level,
		IsActive:    category.IsActive,
		SortOrder:   category.SortOrder,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}
