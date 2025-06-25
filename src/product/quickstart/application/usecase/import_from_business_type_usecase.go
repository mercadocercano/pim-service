package usecase

import (
	"context"
	"fmt"

	tenantPort "pim/src/product/tenant/domain/port"
)

// ImportFromBusinessTypeUseCase caso de uso para importar productos masivamente
type ImportFromBusinessTypeUseCase struct {
	tenantProductRepo tenantPort.ProductRepository
	// TODO: Agregar repositorio del global catalog cuando esté disponible
}

// NewImportFromBusinessTypeUseCase crea una nueva instancia del caso de uso
func NewImportFromBusinessTypeUseCase(
	tenantProductRepo tenantPort.ProductRepository,
) *ImportFromBusinessTypeUseCase {
	return &ImportFromBusinessTypeUseCase{
		tenantProductRepo: tenantProductRepo,
	}
}

// ImportFromBusinessTypeRequest request para importación masiva
type ImportFromBusinessTypeRequest struct {
	TenantID       string   `json:"tenant_id"`
	BusinessTypeID string   `json:"business_type_id"`
	CategoryIDs    []string `json:"category_ids,omitempty"`   // Filtrar por categorías específicas
	ProductIDs     []string `json:"product_ids,omitempty"`    // Importar productos específicos
	ImportAll      bool     `json:"import_all"`               // Importar todos los productos del tipo de negocio
	InitialStatus  string   `json:"initial_status,omitempty"` // Estado inicial (default: "draft")
}

// ImportFromBusinessTypeResponse response de importación masiva
type ImportFromBusinessTypeResponse struct {
	TenantID         string                   `json:"tenant_id"`
	BusinessTypeID   string                   `json:"business_type_id"`
	ImportedProducts []ImportedProductSummary `json:"imported_products"`
	FailedImports    []FailedImportSummary    `json:"failed_imports"`
	Summary          ImportSummary            `json:"summary"`
}

// ImportedProductSummary resumen de producto importado exitosamente
type ImportedProductSummary struct {
	ProductID    string `json:"product_id"`
	TemplateID   string `json:"template_id"`
	ProductName  string `json:"product_name"`
	Status       string `json:"status"`
	CategoryName string `json:"category_name"`
}

// FailedImportSummary resumen de importación fallida
type FailedImportSummary struct {
	TemplateID string `json:"template_id"`
	Error      string `json:"error"`
	Reason     string `json:"reason"`
}

// ImportSummary resumen general de la importación
type ImportSummary struct {
	TotalAttempted int `json:"total_attempted"`
	TotalSuccess   int `json:"total_success"`
	TotalFailed    int `json:"total_failed"`
	SuccessRate    int `json:"success_rate_percentage"`
}

// Execute ejecuta la importación masiva desde tipo de negocio
func (uc *ImportFromBusinessTypeUseCase) Execute(
	ctx context.Context,
	request ImportFromBusinessTypeRequest,
) (*ImportFromBusinessTypeResponse, error) {
	// Validar request
	if request.TenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}
	if request.BusinessTypeID == "" {
		return nil, fmt.Errorf("business_type_id es requerido")
	}

	// Establecer estado inicial por defecto
	if request.InitialStatus == "" {
		request.InitialStatus = "draft"
	}

	// TODO: Implementar lógica real cuando esté disponible el repositorio del global catalog
	// Por ahora simulamos la importación
	response := uc.simulateImport(request)

	return response, nil
}

// simulateImport simula la importación hasta que tengamos el global catalog real
func (uc *ImportFromBusinessTypeUseCase) simulateImport(request ImportFromBusinessTypeRequest) *ImportFromBusinessTypeResponse {
	// Simulación de productos importados
	importedProducts := []ImportedProductSummary{
		{
			ProductID:    "prod-001",
			TemplateID:   "template-iphone-15",
			ProductName:  "iPhone 15 Pro - " + request.TenantID,
			Status:       request.InitialStatus,
			CategoryName: "Smartphones",
		},
		{
			ProductID:    "prod-002",
			TemplateID:   "template-samsung-s24",
			ProductName:  "Samsung Galaxy S24 - " + request.TenantID,
			Status:       request.InitialStatus,
			CategoryName: "Smartphones",
		},
	}

	// Simulación de fallos
	failedImports := []FailedImportSummary{
		{
			TemplateID: "template-invalid",
			Error:      "Template no encontrado",
			Reason:     "El template especificado no existe en el catálogo global",
		},
	}

	totalAttempted := len(importedProducts) + len(failedImports)
	totalSuccess := len(importedProducts)
	totalFailed := len(failedImports)
	successRate := 0
	if totalAttempted > 0 {
		successRate = (totalSuccess * 100) / totalAttempted
	}

	return &ImportFromBusinessTypeResponse{
		TenantID:         request.TenantID,
		BusinessTypeID:   request.BusinessTypeID,
		ImportedProducts: importedProducts,
		FailedImports:    failedImports,
		Summary: ImportSummary{
			TotalAttempted: totalAttempted,
			TotalSuccess:   totalSuccess,
			TotalFailed:    totalFailed,
			SuccessRate:    successRate,
		},
	}
}
