package usecase

import (
	"context"
	"fmt"
	"time"

	"pim/src/category/application/request"
	"pim/src/category/application/response"
	"pim/src/category/domain/entity"
	"pim/src/category/domain/port"
)

// ValidateCategoryHierarchyUseCase maneja la validación de jerarquías de categorías
type ValidateCategoryHierarchyUseCase struct {
	categoryRepo port.MarketplaceCategoryRepository
}

// NewValidateCategoryHierarchyUseCase crea una nueva instancia del caso de uso
func NewValidateCategoryHierarchyUseCase(
	categoryRepo port.MarketplaceCategoryRepository,
) *ValidateCategoryHierarchyUseCase {
	return &ValidateCategoryHierarchyUseCase{
		categoryRepo: categoryRepo,
	}
}

// Execute ejecuta el caso de uso de validación de jerarquía
func (uc *ValidateCategoryHierarchyUseCase) Execute(
	ctx context.Context,
	req *request.ValidateCategoryHierarchyRequest,
) (*response.CategoryHierarchyValidationResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Obtener la categoría a validar
	category, err := uc.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	// Obtener el path actual de la categoría
	currentPath, err := uc.categoryRepo.GetCategoryPath(ctx, req.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get category path: %w", err)
	}

	// Inicializar respuesta
	validationResponse := &response.CategoryHierarchyValidationResponse{
		CategoryID:      req.CategoryID,
		CurrentParentID: category.ParentID,
		NewParentID:     req.NewParentID,
		CurrentLevel:    category.Level,
		MaxDepth:        req.MaxDepth,
		CategoryPath:    uc.buildCategoryPath(currentPath),
		ValidatedAt:     time.Now(),
	}

	var validationErrors []response.ValidationError

	// Validar el nuevo parent si se especifica
	var newLevel int
	if req.NewParentID != nil {
		newParent, err := uc.categoryRepo.GetByID(ctx, *req.NewParentID)
		if err != nil {
			validationErrors = append(validationErrors, response.ValidationError{
				Code:    "PARENT_NOT_FOUND",
				Message: "New parent category not found",
				Field:   "new_parent_id",
			})
		} else {
			newLevel = newParent.Level + 1

			// Validar que no se cree un ciclo
			if uc.wouldCreateCycle(ctx, req.CategoryID, *req.NewParentID) {
				validationErrors = append(validationErrors, response.ValidationError{
					Code:    "CIRCULAR_REFERENCE",
					Message: "Moving category would create a circular reference",
					Field:   "new_parent_id",
				})
			}

			// Validar profundidad máxima
			if newLevel > req.MaxDepth {
				validationErrors = append(validationErrors, response.ValidationError{
					Code:    "MAX_DEPTH_EXCEEDED",
					Message: fmt.Sprintf("New level %d would exceed maximum depth %d", newLevel, req.MaxDepth),
					Field:   "new_parent_id",
				})
			}
		}
	} else {
		newLevel = 0 // Mover a raíz
	}

	validationResponse.NewLevel = newLevel

	// Validar hijos si se solicita
	var affectedChildren []response.AffectedChildCategory
	if req.ValidateChildren {
		children, err := uc.getDescendants(ctx, req.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to get category descendants: %w", err)
		}

		for _, child := range children {
			levelDifference := child.Level - category.Level
			newChildLevel := newLevel + levelDifference

			affectedChild := response.AffectedChildCategory{
				ID:           child.ID,
				Name:         child.Name,
				CurrentLevel: child.Level,
				NewLevel:     newChildLevel,
				WouldExceed:  newChildLevel > req.MaxDepth,
			}

			if affectedChild.WouldExceed {
				validationErrors = append(validationErrors, response.ValidationError{
					Code:       "CHILD_DEPTH_EXCEEDED",
					Message:    fmt.Sprintf("Child category '%s' would exceed maximum depth", child.Name),
					CategoryID: child.ID,
				})
			}

			affectedChildren = append(affectedChildren, affectedChild)
		}
	}

	validationResponse.ValidationErrors = validationErrors
	validationResponse.AffectedChildren = affectedChildren
	validationResponse.IsValid = len(validationErrors) == 0

	return validationResponse, nil
}

// wouldCreateCycle verifica si mover una categoría crearía una referencia circular
func (uc *ValidateCategoryHierarchyUseCase) wouldCreateCycle(ctx context.Context, categoryID, newParentID string) bool {
	// Obtener todos los ancestros del nuevo parent
	ancestors, err := uc.getAncestors(ctx, newParentID)
	if err != nil {
		return true // En caso de error, asumir que sí crearía un ciclo
	}

	// Verificar si la categoría que se quiere mover está entre los ancestros
	for _, ancestor := range ancestors {
		if ancestor.ID == categoryID {
			return true
		}
	}

	return false
}

// getAncestors obtiene todos los ancestros de una categoría
func (uc *ValidateCategoryHierarchyUseCase) getAncestors(ctx context.Context, categoryID string) ([]*entity.MarketplaceCategory, error) {
	var ancestors []*entity.MarketplaceCategory
	currentID := categoryID

	for currentID != "" {
		category, err := uc.categoryRepo.GetByID(ctx, currentID)
		if err != nil {
			break
		}

		ancestors = append(ancestors, category)

		if category.ParentID == nil {
			break
		}
		currentID = *category.ParentID
	}

	return ancestors, nil
}

// getDescendants obtiene todos los descendientes de una categoría
func (uc *ValidateCategoryHierarchyUseCase) getDescendants(ctx context.Context, categoryID string) ([]*entity.MarketplaceCategory, error) {
	var descendants []*entity.MarketplaceCategory

	// Obtener hijos directos
	children, err := uc.categoryRepo.GetByParentID(ctx, &categoryID)
	if err != nil {
		return nil, err
	}

	// Recursivamente obtener descendientes de cada hijo
	for _, child := range children {
		descendants = append(descendants, child)

		childDescendants, err := uc.getDescendants(ctx, child.ID)
		if err != nil {
			continue // Continuar con otros hijos en caso de error
		}

		descendants = append(descendants, childDescendants...)
	}

	return descendants, nil
}

// buildCategoryPath construye el path de categorías para la respuesta
func (uc *ValidateCategoryHierarchyUseCase) buildCategoryPath(categories []*entity.MarketplaceCategory) []response.CategoryPathItem {
	var path []response.CategoryPathItem

	for _, category := range categories {
		path = append(path, response.CategoryPathItem{
			ID:    category.ID,
			Name:  category.Name,
			Level: category.Level,
		})
	}

	return path
}
