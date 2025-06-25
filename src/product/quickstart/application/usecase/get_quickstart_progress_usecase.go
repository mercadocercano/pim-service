package usecase

import (
	"context"
	"fmt"

	tenantPort "pim/src/product/tenant/domain/port"
)

// GetQuickstartProgressUseCase caso de uso para obtener progreso del quickstart
type GetQuickstartProgressUseCase struct {
	tenantProductRepo tenantPort.ProductRepository
	// TODO: Agregar repositorio del tenant/business_type cuando esté disponible
}

// NewGetQuickstartProgressUseCase crea una nueva instancia del caso de uso
func NewGetQuickstartProgressUseCase(
	tenantProductRepo tenantPort.ProductRepository,
) *GetQuickstartProgressUseCase {
	return &GetQuickstartProgressUseCase{
		tenantProductRepo: tenantProductRepo,
	}
}

// GetQuickstartProgressRequest request para obtener progreso
type GetQuickstartProgressRequest struct {
	TenantID string `json:"tenant_id"`
}

// GetQuickstartProgressResponse response del progreso del quickstart
type GetQuickstartProgressResponse struct {
	TenantID             string          `json:"tenant_id"`
	QuickstartSteps      QuickstartSteps `json:"quickstart_steps"`
	CompletionPercentage int             `json:"completion_percentage"`
	NextStep             string          `json:"next_step"`
	Recommendations      []string        `json:"recommendations"`
	ProductsStats        ProductsStats   `json:"products_stats"`
}

// QuickstartSteps estado de cada paso del quickstart
type QuickstartSteps struct {
	BusinessTypeSelected bool `json:"business_type_selected"`
	CategoriesImported   bool `json:"categories_imported"`
	ProductsImported     bool `json:"products_imported"`
	ProductsConfigured   bool `json:"products_configured"`
	ProductsActivated    bool `json:"products_activated"`
	SetupCompleted       bool `json:"setup_completed"`
}

// ProductsStats estadísticas de productos del tenant
type ProductsStats struct {
	TotalProducts    int `json:"total_products"`
	DraftProducts    int `json:"draft_products"`
	PendingProducts  int `json:"pending_products"`
	ActiveProducts   int `json:"active_products"`
	InactiveProducts int `json:"inactive_products"`
}

// Execute ejecuta el caso de uso para obtener progreso del quickstart
func (uc *GetQuickstartProgressUseCase) Execute(
	ctx context.Context,
	request GetQuickstartProgressRequest,
) (*GetQuickstartProgressResponse, error) {
	// Validar request
	if request.TenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}

	// TODO: Obtener estadísticas reales cuando esté implementado el repositorio
	// Por ahora simulamos las estadísticas
	productsStats := uc.getProductsStats(ctx, request.TenantID)

	// Evaluar progreso basado en estadísticas
	steps := uc.evaluateQuickstartSteps(productsStats)

	// Calcular porcentaje de completitud
	completionPercentage := uc.calculateCompletionPercentage(steps)

	// Determinar siguiente paso
	nextStep := uc.determineNextStep(steps)

	// Generar recomendaciones
	recommendations := uc.generateRecommendations(steps, productsStats)

	return &GetQuickstartProgressResponse{
		TenantID:             request.TenantID,
		QuickstartSteps:      steps,
		CompletionPercentage: completionPercentage,
		NextStep:             nextStep,
		Recommendations:      recommendations,
		ProductsStats:        productsStats,
	}, nil
}

// getProductsStats obtiene estadísticas de productos del tenant
func (uc *GetQuickstartProgressUseCase) getProductsStats(ctx context.Context, tenantID string) ProductsStats {
	// TODO: Implementar consultas reales al repositorio
	// Por ahora retornamos datos simulados
	return ProductsStats{
		TotalProducts:    5,
		DraftProducts:    2,
		PendingProducts:  1,
		ActiveProducts:   2,
		InactiveProducts: 0,
	}
}

// evaluateQuickstartSteps evalúa qué pasos del quickstart están completados
func (uc *GetQuickstartProgressUseCase) evaluateQuickstartSteps(stats ProductsStats) QuickstartSteps {
	return QuickstartSteps{
		BusinessTypeSelected: true, // TODO: Verificar desde repositorio de tenant
		CategoriesImported:   true, // TODO: Verificar desde repositorio de categorías
		ProductsImported:     stats.TotalProducts > 0,
		ProductsConfigured:   stats.PendingProducts > 0 || stats.ActiveProducts > 0,
		ProductsActivated:    stats.ActiveProducts > 0,
		SetupCompleted:       stats.ActiveProducts >= 3, // Criterio: al menos 3 productos activos
	}
}

// calculateCompletionPercentage calcula el porcentaje de completitud del quickstart
func (uc *GetQuickstartProgressUseCase) calculateCompletionPercentage(steps QuickstartSteps) int {
	totalSteps := 6
	completedSteps := 0

	if steps.BusinessTypeSelected {
		completedSteps++
	}
	if steps.CategoriesImported {
		completedSteps++
	}
	if steps.ProductsImported {
		completedSteps++
	}
	if steps.ProductsConfigured {
		completedSteps++
	}
	if steps.ProductsActivated {
		completedSteps++
	}
	if steps.SetupCompleted {
		completedSteps++
	}

	return (completedSteps * 100) / totalSteps
}

// determineNextStep determina cuál es el siguiente paso a realizar
func (uc *GetQuickstartProgressUseCase) determineNextStep(steps QuickstartSteps) string {
	if !steps.BusinessTypeSelected {
		return "select_business_type"
	}
	if !steps.CategoriesImported {
		return "import_categories"
	}
	if !steps.ProductsImported {
		return "import_products"
	}
	if !steps.ProductsConfigured {
		return "configure_products"
	}
	if !steps.ProductsActivated {
		return "activate_products"
	}
	if !steps.SetupCompleted {
		return "complete_setup"
	}
	return "setup_completed"
}

// generateRecommendations genera recomendaciones basadas en el estado actual
func (uc *GetQuickstartProgressUseCase) generateRecommendations(steps QuickstartSteps, stats ProductsStats) []string {
	var recommendations []string

	if !steps.ProductsImported {
		recommendations = append(recommendations, "Importa productos desde el catálogo global para comenzar a vender")
	}

	if stats.DraftProducts > 0 {
		recommendations = append(recommendations, fmt.Sprintf("Tienes %d productos en borrador. Configúralos para activarlos", stats.DraftProducts))
	}

	if stats.PendingProducts > 0 {
		recommendations = append(recommendations, fmt.Sprintf("Tienes %d productos pendientes. Agrega precios y stock para activarlos", stats.PendingProducts))
	}

	if stats.ActiveProducts < 3 {
		recommendations = append(recommendations, "Activa al menos 3 productos para completar la configuración inicial")
	}

	if stats.ActiveProducts > 0 && !steps.SetupCompleted {
		recommendations = append(recommendations, "¡Excelente! Ya tienes productos activos. Completa la configuración para finalizar el quickstart")
	}

	if len(recommendations) == 0 {
		recommendations = append(recommendations, "¡Felicitaciones! Has completado la configuración inicial de tu tienda")
	}

	return recommendations
}
