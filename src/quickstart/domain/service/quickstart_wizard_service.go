package service

import (
	"context"
	"encoding/json"
	"fmt"

	businessTypePort "saas-mt-pim-service/src/businesstype/domain/port"
	"saas-mt-pim-service/src/quickstart/domain/entity"
	"saas-mt-pim-service/src/quickstart/domain/port"
)

// QuickstartWizardService maneja la lógica del wizard de configuración inicial
type QuickstartWizardService struct {
	businessTypeRepo         businessTypePort.BusinessTypeRepository
	businessTypeTemplateRepo businessTypePort.BusinessTypeTemplateRepository
	tenantSetupRepo          port.TenantQuickstartHistoryRepository
}

// NewQuickstartWizardService crea una nueva instancia del servicio
func NewQuickstartWizardService(
	businessTypeRepo businessTypePort.BusinessTypeRepository,
	businessTypeTemplateRepo businessTypePort.BusinessTypeTemplateRepository,
	tenantSetupRepo port.TenantQuickstartHistoryRepository,
) *QuickstartWizardService {
	return &QuickstartWizardService{
		businessTypeRepo:         businessTypeRepo,
		businessTypeTemplateRepo: businessTypeTemplateRepo,
		tenantSetupRepo:          tenantSetupRepo,
	}
}

// GetWizardStatus obtiene el estado actual del wizard para un tenant
func (s *QuickstartWizardService) GetWizardStatus(ctx context.Context, tenantID string) (*entity.TenantQuickstartHistory, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenant ID es requerido")
	}

	return s.tenantSetupRepo.GetLatestByTenantID(ctx, tenantID)
}

// StartWizard inicia el proceso de configuración para un tenant
func (s *QuickstartWizardService) StartWizard(ctx context.Context, tenantID, businessTypeID string) (*entity.TenantQuickstartHistory, error) {
	if tenantID == "" {
		return nil, fmt.Errorf("tenant ID es requerido")
	}
	if businessTypeID == "" {
		return nil, fmt.Errorf("business type ID es requerido")
	}

	// Verificar que el business type existe
	businessType, err := s.businessTypeRepo.FindByID(ctx, businessTypeID)
	if err != nil {
		return nil, fmt.Errorf("error verificando business type: %w", err)
	}
	if businessType == nil {
		return nil, fmt.Errorf("business type %s no encontrado", businessTypeID)
	}

	// Obtener template por defecto para el business type
	templates, err := s.businessTypeTemplateRepo.FindByBusinessTypeID(ctx, businessTypeID)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo templates: %w", err)
	}
	if len(templates) == 0 {
		return nil, fmt.Errorf("no hay templates disponibles para business type %s", businessTypeID)
	}

	// Usar el primer template como defecto (después podemos agregar lógica de región)
	defaultTemplate := templates[0]

	// Crear historial inicial
	setupData := map[string]interface{}{
		"business_type_id": businessTypeID,
		"template_id":      defaultTemplate.ID,
		"step":             "business_type_selected",
		"completed_steps":  []string{"business_type_selected"},
		"total_steps":      7,
		"wizard_version":   "1.0",
	}

	setupDataBytes, err := json.Marshal(setupData)
	if err != nil {
		return nil, fmt.Errorf("error serializando setup data: %w", err)
	}

	history, err := entity.NewTenantQuickstartHistory(tenantID, businessTypeID, string(setupDataBytes))
	if err != nil {
		return nil, fmt.Errorf("error creando historial: %w", err)
	}

	err = s.tenantSetupRepo.Create(ctx, history)
	if err != nil {
		return nil, fmt.Errorf("error guardando historial: %w", err)
	}

	return history, nil
}

// UpdateWizardStep actualiza el progreso del wizard
func (s *QuickstartWizardService) UpdateWizardStep(ctx context.Context, tenantID, currentStep string, stepData map[string]interface{}) (*entity.TenantQuickstartHistory, error) {
	// Obtener historial actual
	history, err := s.tenantSetupRepo.GetLatestByTenantID(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo historial: %w", err)
	}
	if history == nil {
		return nil, fmt.Errorf("no hay wizard iniciado para tenant %s", tenantID)
	}

	// Parsear setup data actual
	var setupData map[string]interface{}
	if err := json.Unmarshal([]byte(history.SetupData), &setupData); err != nil {
		return nil, fmt.Errorf("error parseando setup data: %w", err)
	}

	// Actualizar step actual
	setupData["step"] = currentStep

	// Agregar data del step
	stepKey := fmt.Sprintf("step_%s", currentStep)
	setupData[stepKey] = stepData

	// Actualizar completed_steps
	completedSteps, _ := setupData["completed_steps"].([]interface{})
	stepExists := false
	for _, step := range completedSteps {
		if step.(string) == currentStep {
			stepExists = true
			break
		}
	}
	if !stepExists {
		completedSteps = append(completedSteps, currentStep)
		setupData["completed_steps"] = completedSteps
	}

	// Serializar de vuelta
	setupDataBytes, err := json.Marshal(setupData)
	if err != nil {
		return nil, fmt.Errorf("error serializando setup data actualizado: %w", err)
	}

	// Actualizar historial
	history.UpdateSetupData(string(setupDataBytes))

	err = s.tenantSetupRepo.Update(ctx, history)
	if err != nil {
		return nil, fmt.Errorf("error actualizando historial: %w", err)
	}

	return history, nil
}

// CompleteWizard marca el wizard como completado
func (s *QuickstartWizardService) CompleteWizard(ctx context.Context, tenantID string, finalSelections map[string]interface{}) (*entity.TenantQuickstartHistory, error) {
	// Obtener historial actual
	history, err := s.tenantSetupRepo.GetLatestByTenantID(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo historial: %w", err)
	}
	if history == nil {
		return nil, fmt.Errorf("no hay wizard iniciado para tenant %s", tenantID)
	}

	// Parsear setup data actual
	var setupData map[string]interface{}
	if err := json.Unmarshal([]byte(history.SetupData), &setupData); err != nil {
		return nil, fmt.Errorf("error parseando setup data: %w", err)
	}

	// Marcar como completado
	setupData["step"] = "completed"
	setupData["completed"] = true
	setupData["final_selections"] = finalSelections
	setupData["completed_at"] = "now()" // Se puede mejorar con timestamp real

	// Serializar
	setupDataBytes, err := json.Marshal(setupData)
	if err != nil {
		return nil, fmt.Errorf("error serializando setup data final: %w", err)
	}

	// Actualizar historial
	history.UpdateSetupData(string(setupDataBytes))

	err = s.tenantSetupRepo.Update(ctx, history)
	if err != nil {
		return nil, fmt.Errorf("error actualizando historial final: %w", err)
	}

	return history, nil
}

// GetTemplateData obtiene los datos del template para mostrar en el wizard
func (s *QuickstartWizardService) GetTemplateData(ctx context.Context, businessTypeID, section string) (interface{}, error) {
	if businessTypeID == "" {
		return nil, fmt.Errorf("business type ID es requerido")
	}

	// Obtener templates del business type
	templates, err := s.businessTypeTemplateRepo.FindByBusinessTypeID(ctx, businessTypeID)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo templates: %w", err)
	}
	if len(templates) == 0 {
		return nil, fmt.Errorf("no hay templates para business type %s", businessTypeID)
	}

	// Usar primer template como defecto
	template := templates[0]

	// Crear template data desde la estructura de la entidad
	templateData := map[string]interface{}{
		"categories": template.Categories,
		"attributes": template.Attributes,
		"products":   template.Products,
		"brands":     template.Brands,
		"metadata":   template.Metadata,
	}

	// Devolver sección específica o todo
	if section != "" {
		if sectionData, exists := templateData[section]; exists {
			return sectionData, nil
		}
		return nil, fmt.Errorf("sección %s no encontrada en template", section)
	}

	return templateData, nil
}
