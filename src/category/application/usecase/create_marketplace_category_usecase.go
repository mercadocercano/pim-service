package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/category/application/request"
	"saas-mt-pim-service/src/category/application/response"
	"saas-mt-pim-service/src/category/domain/entity"
	"saas-mt-pim-service/src/category/domain/port"
)

// CreateMarketplaceCategoryUseCase maneja la creación de categorías marketplace
type CreateMarketplaceCategoryUseCase struct {
	categoryRepo port.MarketplaceCategoryRepository
}

// NewCreateMarketplaceCategoryUseCase crea una nueva instancia del caso de uso
func NewCreateMarketplaceCategoryUseCase(
	categoryRepo port.MarketplaceCategoryRepository,
) *CreateMarketplaceCategoryUseCase {
	return &CreateMarketplaceCategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

// Execute ejecuta el caso de uso de creación de categoría marketplace
func (uc *CreateMarketplaceCategoryUseCase) Execute(
	ctx context.Context,
	req *request.CreateMarketplaceCategoryRequest,
) (*response.MarketplaceCategoryResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Verificar que el slug no esté en uso
	exists, err := uc.categoryRepo.ExistsBySlug(ctx, req.Slug)
	if err != nil {
		return nil, fmt.Errorf("error checking slug uniqueness: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("category with slug '%s' already exists", req.Slug)
	}

	// Validar jerarquía si tiene parent
	var level int = 0
	if req.ParentID != nil {
		parentCategory, err := uc.categoryRepo.GetByID(ctx, *req.ParentID)
		if err != nil {
			return nil, fmt.Errorf("parent category not found: %w", err)
		}

		// Validar máximo 3 niveles de profundidad
		if parentCategory.Level >= 2 {
			return nil, fmt.Errorf("maximum category depth (3 levels) exceeded")
		}

		level = parentCategory.Level + 1
	}

	// Crear la entidad categoría marketplace usando el constructor
	description := ""
	if req.Description != nil {
		description = *req.Description
	}

	category, err := entity.NewMarketplaceCategory(req.Name, req.Slug, description, req.ParentID)
	if err != nil {
		return nil, fmt.Errorf("failed to create marketplace category: %w", err)
	}

	// Actualizar campos adicionales
	category.Level = level
	category.UpdateSortOrder(req.SortOrder)

	// Guardar la categoría
	if err := uc.categoryRepo.Save(ctx, category); err != nil {
		return nil, fmt.Errorf("failed to save marketplace category: %w", err)
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
