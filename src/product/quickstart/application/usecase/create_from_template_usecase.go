package usecase

import (
	"context"
	"errors"
	"fmt"

	tenantEntity "saas-mt-pim-service/src/product/tenant/domain/entity"
	tenantPort "saas-mt-pim-service/src/product/tenant/domain/port"
	tenantValueObject "saas-mt-pim-service/src/product/tenant/domain/value_object"
)

// CreateFromTemplateUseCase caso de uso para crear productos desde template del catálogo global
type CreateFromTemplateUseCase struct {
	tenantProductRepo tenantPort.ProductRepository
}

// NewCreateFromTemplateUseCase crea una nueva instancia del caso de uso
func NewCreateFromTemplateUseCase(
	tenantProductRepo tenantPort.ProductRepository,
) *CreateFromTemplateUseCase {
	return &CreateFromTemplateUseCase{
		tenantProductRepo: tenantProductRepo,
	}
}

// CreateFromTemplateRequest request para crear producto desde template
type CreateFromTemplateRequest struct {
	TenantID      string                `json:"tenant_id"`
	TemplateName  string                `json:"template_name"`
	InitialStatus string                `json:"initial_status,omitempty"` // "draft" por defecto
	CustomName    *string               `json:"custom_name,omitempty"`
	CustomSKU     *string               `json:"custom_sku,omitempty"`
	ImportOptions *ProductImportOptions `json:"import_options,omitempty"`
}

// ProductImportOptions opciones adicionales para la importación
type ProductImportOptions struct {
	AutoActivate   bool `json:"auto_activate"`   // Activar automáticamente si cumple validaciones
	SkipValidation bool `json:"skip_validation"` // Saltar validaciones de negocio
}

// CreateFromTemplateResponse response del caso de uso
type CreateFromTemplateResponse struct {
	ProductID        string   `json:"product_id"`
	Status           string   `json:"status"`
	RequiredActions  []string `json:"required_actions,omitempty"`
	ValidationErrors []string `json:"validation_errors,omitempty"`
	WarningMessages  []string `json:"warning_messages,omitempty"`
}

// Execute ejecuta el caso de uso
func (uc *CreateFromTemplateUseCase) Execute(ctx context.Context, request CreateFromTemplateRequest) (*CreateFromTemplateResponse, error) {
	// Validar request
	if err := uc.validateRequest(request); err != nil {
		return nil, err
	}

	// Determinar estado inicial
	initialStatus := request.InitialStatus
	if initialStatus == "" {
		initialStatus = tenantValueObject.ProductStatusDraftValue
	}

	// Crear producto base desde template (simulado por ahora)
	product, err := uc.createProductFromTemplate(request, initialStatus)
	if err != nil {
		return nil, fmt.Errorf("error al crear producto desde template: %w", err)
	}

	// Guardar producto
	if err := uc.tenantProductRepo.Save(ctx, product); err != nil {
		return nil, fmt.Errorf("error al guardar producto: %w", err)
	}

	// Intentar auto-activación si está configurada
	finalStatus := product.Status().Value()
	validationErrors := []string{}
	requiredActions := []string{}

	if request.ImportOptions != nil && request.ImportOptions.AutoActivate {
		if err := product.ActivateWithValidation(); err != nil {
			validationErrors = append(validationErrors, err.Error())
			requiredActions = uc.getRequiredActionsForActivation(product)
		} else {
			finalStatus = tenantValueObject.ProductStatusActiveValue
			// Actualizar en base de datos
			if err := uc.tenantProductRepo.Save(ctx, product); err != nil {
				return nil, fmt.Errorf("error al actualizar producto activado: %w", err)
			}
		}
	} else {
		// Obtener acciones requeridas según el estado actual
		requiredActions = uc.getRequiredActionsByStatus(product)
	}

	return &CreateFromTemplateResponse{
		ProductID:        product.IDString(),
		Status:           finalStatus,
		RequiredActions:  requiredActions,
		ValidationErrors: validationErrors,
		WarningMessages:  []string{},
	}, nil
}

// validateRequest valida el request
func (uc *CreateFromTemplateUseCase) validateRequest(request CreateFromTemplateRequest) error {
	if request.TenantID == "" {
		return errors.New("tenant_id es obligatorio")
	}

	if request.TemplateName == "" {
		return errors.New("template_name es obligatorio")
	}

	// Validar estado inicial si se proporciona
	if request.InitialStatus != "" {
		if _, err := tenantValueObject.NewProductStatus(request.InitialStatus); err != nil {
			return fmt.Errorf("estado inicial inválido: %w", err)
		}
	}

	return nil
}

// createProductFromTemplate crea el producto base desde el template (simulado)
func (uc *CreateFromTemplateUseCase) createProductFromTemplate(
	request CreateFromTemplateRequest,
	initialStatus string,
) (*tenantEntity.Product, error) {
	// Determinar nombre del producto
	productName := request.TemplateName
	if request.CustomName != nil && *request.CustomName != "" {
		productName = *request.CustomName
	}

	// Crear SKU si se proporciona
	var productSKU *tenantValueObject.ProductSKU
	if request.CustomSKU != nil && *request.CustomSKU != "" {
		sku, err := tenantValueObject.NewProductSKU(*request.CustomSKU)
		if err != nil {
			return nil, fmt.Errorf("SKU inválido: %w", err)
		}
		productSKU = sku
	}

	// Por ahora crear producto sin referencias de categoría/marca
	// TODO: Integrar con global catalog cuando esté disponible
	product, err := tenantEntity.NewProduct(
		request.TenantID,
		productName,
		nil, // Descripción por ahora vacía
		productSKU,
		nil, // Categoría por ahora vacía
		nil, // Marca por ahora vacía
	)
	if err != nil {
		return nil, err
	}

	// Establecer estado inicial
	if err := product.TransitionToStatus(initialStatus); err != nil {
		return nil, fmt.Errorf("error al establecer estado inicial: %w", err)
	}

	return product, nil
}

// getRequiredActionsForActivation obtiene las acciones requeridas para activar el producto
func (uc *CreateFromTemplateUseCase) getRequiredActionsForActivation(product *tenantEntity.Product) []string {
	actions := []string{}

	// Validar nombre
	if product.Name() == "" {
		actions = append(actions, "Configurar nombre del producto")
	}

	// Validar variantes
	if !product.HasVariants() {
		actions = append(actions, "Agregar al menos una variante")
	}

	// Validar precios (esto dependería de cómo manejes los precios)
	actions = append(actions, "Configurar precios para las variantes")

	// Validar stock (esto dependería de integración con stock service)
	actions = append(actions, "Configurar stock inicial")

	return actions
}

// getRequiredActionsByStatus obtiene las acciones requeridas según el estado actual
func (uc *CreateFromTemplateUseCase) getRequiredActionsByStatus(product *tenantEntity.Product) []string {
	actions := []string{}

	switch product.Status().Value() {
	case tenantValueObject.ProductStatusDraftValue:
		actions = append(actions, "Revisar y configurar información básica del producto")
		actions = append(actions, "Configurar variantes si es necesario")
		actions = append(actions, "Cambiar estado a 'pending' cuando esté listo")

	case tenantValueObject.ProductStatusPendingValue:
		actions = append(actions, "Configurar precios para todas las variantes")
		actions = append(actions, "Configurar stock inicial")
		actions = append(actions, "Activar producto cuando esté completamente configurado")

	case tenantValueObject.ProductStatusActiveValue:
		actions = append(actions, "El producto está listo para vender")

	case tenantValueObject.ProductStatusInactiveValue:
		actions = append(actions, "Reactivar producto si desea venderlo nuevamente")
	}

	return actions
}
