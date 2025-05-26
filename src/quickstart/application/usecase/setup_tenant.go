package usecase

import (
	"context"
	"fmt"

	"pim/src/quickstart/domain/entity"
	"pim/src/quickstart/domain/exception"
	"pim/src/quickstart/domain/port"
	"pim/src/quickstart/domain/service"
	"pim/src/quickstart/domain/value_object"
)

// SetupTenantUseCase implementa el caso de uso para configurar un tenant con quickstart
type SetupTenantUseCase struct {
	quickstartService *service.QuickstartService
	historyRepo       port.TenantQuickstartHistoryRepository
	categoryService   port.CategoryService
	attributeService  port.AttributeService
	variantService    port.VariantService
	productService    port.ProductService
}

// NewSetupTenantUseCase crea una nueva instancia del caso de uso
func NewSetupTenantUseCase(
	quickstartService *service.QuickstartService,
	historyRepo port.TenantQuickstartHistoryRepository,
	categoryService port.CategoryService,
	attributeService port.AttributeService,
	variantService port.VariantService,
	productService port.ProductService,
) *SetupTenantUseCase {
	return &SetupTenantUseCase{
		quickstartService: quickstartService,
		historyRepo:       historyRepo,
		categoryService:   categoryService,
		attributeService:  attributeService,
		variantService:    variantService,
		productService:    productService,
	}
}

// Execute ejecuta el caso de uso para configurar un tenant
func (uc *SetupTenantUseCase) Execute(ctx context.Context, tenantID string, setupDataMap map[string]interface{}) (*entity.TenantQuickstartHistory, error) {
	// Verificar si ya existe una configuración completada para este tenant
	latestHistory, err := uc.historyRepo.GetLatestByTenantID(ctx, tenantID)
	if err == nil && latestHistory != nil && latestHistory.IsCompleted() {
		return nil, exception.ErrQuickstartAlreadyCompleted
	}

	// Verificar si hay una configuración en progreso
	if err == nil && latestHistory != nil && latestHistory.IsPending() {
		return nil, exception.ErrQuickstartInProgress
	}

	// Preparar y validar los datos de configuración
	setupDataJSON, err := uc.quickstartService.PrepareSetupData(setupDataMap)
	if err != nil {
		return nil, fmt.Errorf("error preparando datos de configuración: %w", err)
	}

	// Crear el objeto de valor SetupData para validaciones adicionales
	setupData, err := value_object.FromJSON(setupDataJSON)
	if err != nil {
		return nil, fmt.Errorf("error validando datos de configuración: %w", err)
	}

	// Crear el historial de configuración
	history, err := entity.NewTenantQuickstartHistory(tenantID, setupData.BusinessType, setupDataJSON)
	if err != nil {
		return nil, fmt.Errorf("error creando historial de configuración: %w", err)
	}

	// Guardar el historial como pendiente
	if err := uc.historyRepo.Create(ctx, history); err != nil {
		return nil, fmt.Errorf("error guardando historial de configuración: %w", err)
	}

	// Ejecutar la configuración
	if err := uc.executeSetup(ctx, tenantID, setupData); err != nil {
		// Marcar como fallido
		history.MarkAsFailed(err.Error())
		uc.historyRepo.Update(ctx, history)
		return history, fmt.Errorf("error ejecutando configuración: %w", err)
	}

	// Marcar como completado
	history.MarkAsCompleted()
	if err := uc.historyRepo.Update(ctx, history); err != nil {
		return history, fmt.Errorf("error actualizando historial de configuración: %w", err)
	}

	return history, nil
}

// executeSetup ejecuta la configuración del tenant paso a paso
func (uc *SetupTenantUseCase) executeSetup(ctx context.Context, tenantID string, setupData *value_object.SetupData) error {
	// 1. Crear categorías seleccionadas
	if setupData.HasCategories() {
		categoriesData, err := uc.quickstartService.GetCategoriesByBusinessType(ctx, setupData.BusinessType)
		if err != nil {
			return fmt.Errorf("error obteniendo categorías: %w", err)
		}

		if err := uc.categoryService.CreateFromTemplate(ctx, tenantID, categoriesData); err != nil {
			return fmt.Errorf("error creando categorías: %w", err)
		}
	}

	// 2. Crear atributos seleccionados
	if setupData.HasAttributes() {
		attributesData, err := uc.quickstartService.GetAttributesByBusinessType(ctx, setupData.BusinessType)
		if err != nil {
			return fmt.Errorf("error obteniendo atributos: %w", err)
		}

		if err := uc.attributeService.CreateFromTemplate(ctx, tenantID, attributesData); err != nil {
			return fmt.Errorf("error creando atributos: %w", err)
		}
	}

	// 3. Crear variantes seleccionadas
	if setupData.HasVariants() {
		variantsData, err := uc.quickstartService.GetVariantsByBusinessType(ctx, setupData.BusinessType)
		if err != nil {
			return fmt.Errorf("error obteniendo variantes: %w", err)
		}

		if err := uc.variantService.CreateFromTemplate(ctx, tenantID, variantsData); err != nil {
			return fmt.Errorf("error creando variantes: %w", err)
		}
	}

	// 4. Crear productos básicos seleccionados
	if setupData.HasProducts() {
		productsData, err := uc.quickstartService.GetProductsByBusinessType(ctx, setupData.BusinessType)
		if err != nil {
			return fmt.Errorf("error obteniendo productos: %w", err)
		}

		if err := uc.productService.CreateFromTemplate(ctx, tenantID, productsData); err != nil {
			return fmt.Errorf("error creando productos: %w", err)
		}
	}

	return nil
}
