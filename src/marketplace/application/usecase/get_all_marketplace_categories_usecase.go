package usecase

import (
	"context"

	"pim/src/marketplace/application/response"
	"pim/src/marketplace/domain/port"
)

// GetAllMarketplaceCategoriesUseCase maneja la obtención de todas las categorías marketplace
type GetAllMarketplaceCategoriesUseCase struct {
	categoryRepository port.MarketplaceCategoryRepository
}

// NewGetAllMarketplaceCategoriesUseCase crea una nueva instancia del caso de uso
func NewGetAllMarketplaceCategoriesUseCase(
	categoryRepository port.MarketplaceCategoryRepository,
) *GetAllMarketplaceCategoriesUseCase {
	return &GetAllMarketplaceCategoriesUseCase{
		categoryRepository: categoryRepository,
	}
}

// Execute ejecuta el caso de uso para obtener todas las categorías marketplace
func (uc *GetAllMarketplaceCategoriesUseCase) Execute(ctx context.Context) ([]*response.MarketplaceCategoryResponse, error) {
	// Obtener todas las categorías del marketplace
	categories, err := uc.categoryRepository.GetTree(ctx)
	if err != nil {
		return nil, err
	}

	// Convertir a response
	responses := make([]*response.MarketplaceCategoryResponse, 0, len(categories))
	for _, category := range categories {
		var description *string
		if category.Description != "" {
			description = &category.Description
		}

		responses = append(responses, &response.MarketplaceCategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Slug:        category.Slug,
			Description: description,
			ParentID:    category.ParentID,
			Level:       category.Level,
			IsActive:    category.IsActive,
			SortOrder:   category.SortOrder,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		})
	}

	return responses, nil
}
