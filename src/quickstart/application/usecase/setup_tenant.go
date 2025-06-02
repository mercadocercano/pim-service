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
	quickstartService        *service.QuickstartService
	historyRepo              port.TenantQuickstartHistoryRepository
	categoryService          port.CategoryService
	attributeService         port.AttributeService
	categoryAttributeService port.CategoryAttributeService
	variantService           port.VariantService
	productService           port.ProductService
}

// NewSetupTenantUseCase crea una nueva instancia del caso de uso
func NewSetupTenantUseCase(
	quickstartService *service.QuickstartService,
	historyRepo port.TenantQuickstartHistoryRepository,
	categoryService port.CategoryService,
	attributeService port.AttributeService,
	categoryAttributeService port.CategoryAttributeService,
	variantService port.VariantService,
	productService port.ProductService,
) *SetupTenantUseCase {
	return &SetupTenantUseCase{
		quickstartService:        quickstartService,
		historyRepo:              historyRepo,
		categoryService:          categoryService,
		attributeService:         attributeService,
		categoryAttributeService: categoryAttributeService,
		variantService:           variantService,
		productService:           productService,
	}
}

// Execute ejecuta el caso de uso para configurar un tenant
func (uc *SetupTenantUseCase) Execute(ctx context.Context, tenantID string, setupDataMap map[string]interface{}) (*entity.TenantQuickstartHistory, error) {
	// Verificación temporal: si no hay repositorio, crear una respuesta mock exitosa
	if uc.historyRepo == nil {
		// Implementación temporal sin persistencia
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

		// Ejecutar la configuración real aunque no tengamos historial
		if err := uc.executeSetup(ctx, tenantID, setupData); err != nil {
			return nil, fmt.Errorf("error ejecutando configuración: %w", err)
		}

		// Crear una respuesta mock de historial exitoso
		history, err := entity.NewTenantQuickstartHistory(tenantID, setupData.BusinessType, setupDataJSON)
		if err != nil {
			return nil, fmt.Errorf("error creando historial de configuración: %w", err)
		}

		// Marcar como completado sin persistir
		history.MarkAsCompleted()

		// Log para debugging
		fmt.Printf("Setup completado para tenant %s con business type %s\n", tenantID, setupData.BusinessType)

		return history, nil
	}

	// Código original para cuando las dependencias estén implementadas
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

		// Si tenemos categoryService implementado, usarlo
		if uc.categoryService != nil {
			if err := uc.categoryService.CreateFromTemplate(ctx, tenantID, categoriesData); err != nil {
				return fmt.Errorf("error creando categorías: %w", err)
			}
		} else {
			fmt.Printf("CategoryService no implementado, omitiendo creación de categorías\n")
		}
	}

	// 2. Crear atributos seleccionados
	if setupData.HasAttributes() {
		attributesData, err := uc.quickstartService.GetAttributesByBusinessType(ctx, setupData.BusinessType)
		if err != nil {
			return fmt.Errorf("error obteniendo atributos: %w", err)
		}

		// Si tenemos attributeService implementado, usarlo
		if uc.attributeService != nil {
			if err := uc.attributeService.CreateFromTemplate(ctx, tenantID, attributesData); err != nil {
				return fmt.Errorf("error creando atributos: %w", err)
			}
		} else {
			fmt.Printf("AttributeService no implementado, omitiendo creación de atributos\n")
		}
	}

	// 2.5. Crear relaciones categoría-atributo si tenemos tanto categorías como atributos
	if setupData.HasCategories() && setupData.HasAttributes() {
		// Si tenemos categoryAttributeService implementado, usarlo
		if uc.categoryAttributeService != nil {
			if err := uc.categoryAttributeService.CreateFromTemplate(ctx, tenantID, nil); err != nil {
				return fmt.Errorf("error creando relaciones categoría-atributo: %w", err)
			}
		} else {
			fmt.Printf("CategoryAttributeService no implementado, omitiendo creación de relaciones categoría-atributo\n")
		}
	}

	// 3. Crear variantes seleccionadas
	if setupData.HasVariants() {
		variantsData, err := uc.quickstartService.GetVariantsByBusinessType(ctx, setupData.BusinessType)
		if err != nil {
			return fmt.Errorf("error obteniendo variantes: %w", err)
		}

		// Si tenemos variantService implementado, usarlo
		if uc.variantService != nil {
			if err := uc.variantService.CreateFromTemplate(ctx, tenantID, variantsData); err != nil {
				return fmt.Errorf("error creando variantes: %w", err)
			}
		} else {
			fmt.Printf("VariantService no implementado, omitiendo creación de variantes\n")
		}
	}

	// 4. Crear productos básicos seleccionados
	if setupData.HasProducts() {
		productsData, err := uc.quickstartService.GetProductsByBusinessType(ctx, setupData.BusinessType)
		if err != nil {
			return fmt.Errorf("error obteniendo productos: %w", err)
		}

		// Si tenemos productService implementado, usarlo
		if uc.productService != nil {
			if err := uc.productService.CreateFromTemplate(ctx, tenantID, productsData); err != nil {
				return fmt.Errorf("error creando productos: %w", err)
			}
		} else {
			fmt.Printf("ProductService no implementado, omitiendo creación de productos\n")
		}
	}

	return nil
}
