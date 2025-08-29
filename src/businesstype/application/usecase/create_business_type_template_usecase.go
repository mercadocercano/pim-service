package usecase

import (
	"context"
	"fmt"

	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/businesstype/domain/port"
)

// CreateBusinessTypeTemplateUseCase maneja la creación de templates de business type
type CreateBusinessTypeTemplateUseCase struct {
	templateRepo    port.BusinessTypeTemplateRepository
	businessTypeRepo port.BusinessTypeRepository
}

// NewCreateBusinessTypeTemplateUseCase crea una nueva instancia del caso de uso
func NewCreateBusinessTypeTemplateUseCase(
	templateRepo port.BusinessTypeTemplateRepository,
	businessTypeRepo port.BusinessTypeRepository,
) *CreateBusinessTypeTemplateUseCase {
	return &CreateBusinessTypeTemplateUseCase{
		templateRepo:     templateRepo,
		businessTypeRepo: businessTypeRepo,
	}
}

// CreateTemplateRequest representa la request para crear un template
type CreateTemplateRequest struct {
	BusinessTypeID string                           `json:"business_type_id"`
	Name           string                           `json:"name"`
	Description    string                           `json:"description"`
	Version        string                           `json:"version"`
	Region         string                           `json:"region"`
	Categories     []entity.CategoryTemplate        `json:"categories"`
	Attributes     []entity.AttributeTemplate       `json:"attributes"`
	Products       []entity.ProductTemplate         `json:"products"`
	Brands         []string                         `json:"brands"`
	IsDefault      bool                             `json:"is_default"`
	Metadata       map[string]interface{}           `json:"metadata"`
}

// Execute ejecuta el caso de uso para crear un template
func (uc *CreateBusinessTypeTemplateUseCase) Execute(ctx context.Context, req CreateTemplateRequest) (*entity.BusinessTypeTemplate, error) {
	// Validar que el business type existe
	businessType, err := uc.businessTypeRepo.FindByID(ctx, req.BusinessTypeID)
	if err != nil {
		return nil, fmt.Errorf("error validating business type: %w", err)
	}
	if businessType == nil {
		return nil, fmt.Errorf("business type with id %s not found", req.BusinessTypeID)
	}

	// Crear el template
	template, err := entity.NewBusinessTypeTemplate(
		req.BusinessTypeID,
		req.Name,
		req.Description,
		req.Version,
		req.Region,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating template entity: %w", err)
	}

	// Asignar datos adicionales
	template.Categories = req.Categories
	template.Attributes = req.Attributes
	template.Products = req.Products
	template.Brands = req.Brands
	template.Metadata = req.Metadata

	if req.IsDefault {
		template.SetAsDefault()
	}

	// Si se marca como default, verificar que no haya otro default para el mismo business type y región
	if template.IsDefault {
		existingDefault, err := uc.templateRepo.FindDefault(ctx, req.BusinessTypeID, req.Region)
		if err != nil {
			return nil, fmt.Errorf("error checking existing default template: %w", err)
		}
		if existingDefault != nil {
			return nil, fmt.Errorf("already exists a default template for business type %s in region %s", req.BusinessTypeID, req.Region)
		}
	}

	// Guardar en la base de datos
	if err := uc.templateRepo.Create(ctx, template); err != nil {
		return nil, fmt.Errorf("error saving template: %w", err)
	}

	return template, nil
}