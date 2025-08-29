package usecase

import (
	"context"
	"fmt"
	"time"

	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/businesstype/domain/port"
)

// UpdateBusinessTypeTemplateUseCase maneja la actualización de templates de business type
type UpdateBusinessTypeTemplateUseCase struct {
	templateRepo port.BusinessTypeTemplateRepository
}

// NewUpdateBusinessTypeTemplateUseCase crea una nueva instancia del caso de uso
func NewUpdateBusinessTypeTemplateUseCase(templateRepo port.BusinessTypeTemplateRepository) *UpdateBusinessTypeTemplateUseCase {
	return &UpdateBusinessTypeTemplateUseCase{
		templateRepo: templateRepo,
	}
}

// UpdateTemplateRequest representa la request para actualizar un template
type UpdateTemplateRequest struct {
	ID          string                           `json:"id"`
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Version     string                           `json:"version"`
	Region      string                           `json:"region"`
	Categories  []entity.CategoryTemplate        `json:"categories"`
	Attributes  []entity.AttributeTemplate       `json:"attributes"`
	Products    []entity.ProductTemplate         `json:"products"`
	Brands      []string                         `json:"brands"`
	IsActive    bool                             `json:"is_active"`
	IsDefault   bool                             `json:"is_default"`
	Metadata    map[string]interface{}           `json:"metadata"`
}

// Execute ejecuta el caso de uso para actualizar un template
func (uc *UpdateBusinessTypeTemplateUseCase) Execute(ctx context.Context, req UpdateTemplateRequest) (*entity.BusinessTypeTemplate, error) {
	// Buscar el template existente
	existingTemplate, err := uc.templateRepo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("error finding template: %w", err)
	}
	if existingTemplate == nil {
		return nil, fmt.Errorf("template with id %s not found", req.ID)
	}

	// Si se está cambiando a default, verificar que no haya otro default
	if req.IsDefault && !existingTemplate.IsDefault {
		existingDefault, err := uc.templateRepo.FindDefault(ctx, existingTemplate.BusinessTypeID, req.Region)
		if err != nil {
			return nil, fmt.Errorf("error checking existing default template: %w", err)
		}
		if existingDefault != nil && existingDefault.ID != req.ID {
			return nil, fmt.Errorf("already exists a default template for business type %s in region %s", existingTemplate.BusinessTypeID, req.Region)
		}
	}

	// Actualizar campos del template
	existingTemplate.Name = req.Name
	existingTemplate.Description = req.Description
	existingTemplate.Version = req.Version
	existingTemplate.Region = req.Region
	existingTemplate.Categories = req.Categories
	existingTemplate.Attributes = req.Attributes
	existingTemplate.Products = req.Products
	existingTemplate.Brands = req.Brands
	existingTemplate.IsDefault = req.IsDefault
	existingTemplate.Metadata = req.Metadata
	existingTemplate.UpdatedAt = time.Now()

	// Actualizar estado activo/inactivo
	if req.IsActive {
		existingTemplate.Activate()
	} else {
		existingTemplate.Deactivate()
	}

	// Guardar cambios
	if err := uc.templateRepo.Update(ctx, existingTemplate); err != nil {
		return nil, fmt.Errorf("error updating template: %w", err)
	}

	return existingTemplate, nil
}