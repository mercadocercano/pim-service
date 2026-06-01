package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"saas-mt-pim-service/src/businesstype/domain/entity"
	"saas-mt-pim-service/src/businesstype/domain/port"
)

// DuplicateTemplateUseCase duplica un template existente
type DuplicateTemplateUseCase struct {
	templateRepo port.BusinessTypeTemplateRepository
}

// NewDuplicateTemplateUseCase crea una nueva instancia del caso de uso
func NewDuplicateTemplateUseCase(templateRepo port.BusinessTypeTemplateRepository) *DuplicateTemplateUseCase {
	return &DuplicateTemplateUseCase{templateRepo: templateRepo}
}

// DuplicateTemplateRequest contiene los parámetros opcionales para la duplicación
type DuplicateTemplateRequest struct {
	TemplateID string `json:"template_id"`
	NewName    string `json:"new_name"`
}

// Execute duplica un template y retorna la copia
func (uc *DuplicateTemplateUseCase) Execute(ctx context.Context, req DuplicateTemplateRequest) (*entity.BusinessTypeTemplate, error) {
	if req.TemplateID == "" {
		return nil, fmt.Errorf("template_id es requerido")
	}

	original, err := uc.templateRepo.FindByID(ctx, req.TemplateID)
	if err != nil {
		return nil, fmt.Errorf("error buscando template: %w", err)
	}
	if original == nil {
		return nil, fmt.Errorf("template no encontrado")
	}

	newName := buildDuplicateName(req.NewName, original.Name)
	copy := buildTemplateCopy(original, newName)

	if err := uc.templateRepo.Create(ctx, copy); err != nil {
		return nil, fmt.Errorf("error guardando copia del template: %w", err)
	}

	return copy, nil
}

// buildDuplicateName determina el nombre de la copia
func buildDuplicateName(requestedName, originalName string) string {
	if requestedName != "" {
		return requestedName
	}
	return originalName + " (copia)"
}

// buildTemplateCopy construye la entidad copia a partir del original
func buildTemplateCopy(original *entity.BusinessTypeTemplate, newName string) *entity.BusinessTypeTemplate {
	now := time.Now()
	return &entity.BusinessTypeTemplate{
		ID:             uuid.New().String(),
		BusinessTypeID: original.BusinessTypeID,
		Name:           newName,
		Description:    original.Description,
		Version:        original.Version,
		Region:         original.Region,
		Categories:     original.Categories,
		Attributes:     original.Attributes,
		Products:       original.Products,
		Brands:         original.Brands,
		IsActive:       true,
		IsDefault:      false, // La copia nunca es default
		Metadata:       original.Metadata,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
