package usecase

import (
	"context"
	"fmt"
	"log"

	globalPort "saas-mt-pim-service/src/product/global_catalog/domain/port"
	tenantEntity "saas-mt-pim-service/src/product/tenant/domain/entity"
	tenantPort "saas-mt-pim-service/src/product/tenant/domain/port"
)

type ImportFromBusinessTypeUseCase struct {
	tenantProductRepo tenantPort.ProductRepository
	globalCatalogRepo globalPort.GlobalProductRepository
}

func NewImportFromBusinessTypeUseCase(
	tenantProductRepo tenantPort.ProductRepository,
	globalCatalogRepo globalPort.GlobalProductRepository,
) *ImportFromBusinessTypeUseCase {
	return &ImportFromBusinessTypeUseCase{
		tenantProductRepo: tenantProductRepo,
		globalCatalogRepo: globalCatalogRepo,
	}
}

type ImportFromBusinessTypeRequest struct {
	TenantID       string   `json:"tenant_id"`
	BusinessTypeID string   `json:"business_type_id"`
	CategoryIDs    []string `json:"category_ids,omitempty"`
	ProductIDs     []string `json:"product_ids,omitempty"`
	ImportAll      bool     `json:"import_all"`
	InitialStatus  string   `json:"initial_status,omitempty"`
}

type ImportFromBusinessTypeResponse struct {
	TenantID         string                  `json:"tenant_id"`
	BusinessTypeID   string                  `json:"business_type_id"`
	ImportedProducts []ImportedProductSummary `json:"imported_products"`
	FailedImports    []FailedImportSummary    `json:"failed_imports"`
	Summary          ImportSummary            `json:"summary"`
}

type ImportedProductSummary struct {
	ProductID    string `json:"product_id"`
	TemplateID   string `json:"template_id"`
	ProductName  string `json:"product_name"`
	Status       string `json:"status"`
	CategoryName string `json:"category_name"`
}

type FailedImportSummary struct {
	TemplateID string `json:"template_id"`
	Error      string `json:"error"`
	Reason     string `json:"reason"`
}

type ImportSummary struct {
	TotalAttempted int `json:"total_attempted"`
	TotalSuccess   int `json:"total_success"`
	TotalFailed    int `json:"total_failed"`
	SuccessRate    int `json:"success_rate_percentage"`
}

func (uc *ImportFromBusinessTypeUseCase) Execute(
	ctx context.Context,
	request ImportFromBusinessTypeRequest,
) (*ImportFromBusinessTypeResponse, error) {
	if request.TenantID == "" {
		return nil, fmt.Errorf("tenant_id es requerido")
	}
	if request.BusinessTypeID == "" {
		return nil, fmt.Errorf("business_type_id es requerido")
	}
	if request.InitialStatus == "" {
		request.InitialStatus = "active"
	}

	globals, err := uc.globalCatalogRepo.FindByBusinessType(request.BusinessTypeID, 50)
	if err != nil {
		return nil, fmt.Errorf("buscando productos del catálogo global: %w", err)
	}

	var imported []ImportedProductSummary
	var failed []FailedImportSummary

	for _, gp := range globals {
		gpID := gp.IDString()
		if len(request.ProductIDs) > 0 && !contains(request.ProductIDs, gpID) {
			continue
		}

		product, createErr := tenantEntity.NewProduct(
			request.TenantID,
			gp.Name(),
			gp.Description(),
			nil,
			nil,
			nil,
		)
		if createErr != nil {
			failed = append(failed, FailedImportSummary{
				TemplateID: gpID,
				Error:      createErr.Error(),
				Reason:     "Error creando producto tenant",
			})
			continue
		}

		if saveErr := uc.tenantProductRepo.Save(ctx, product); saveErr != nil {
			failed = append(failed, FailedImportSummary{
				TemplateID: gpID,
				Error:      saveErr.Error(),
				Reason:     "Error guardando producto",
			})
			continue
		}

		categoryName := ""
		if cat := gp.Category(); cat != nil {
			categoryName = *cat
		}

		imported = append(imported, ImportedProductSummary{
			ProductID:    product.IDString(),
			TemplateID:   gpID,
			ProductName:  gp.Name(),
			Status:       request.InitialStatus,
			CategoryName: categoryName,
		})

		log.Printf("[quickstart] imported global product %s → tenant product %s for tenant %s",
			gpID, product.IDString(), request.TenantID)
	}

	totalAttempted := len(imported) + len(failed)
	successRate := 0
	if totalAttempted > 0 {
		successRate = (len(imported) * 100) / totalAttempted
	}

	return &ImportFromBusinessTypeResponse{
		TenantID:         request.TenantID,
		BusinessTypeID:   request.BusinessTypeID,
		ImportedProducts: imported,
		FailedImports:    failed,
		Summary: ImportSummary{
			TotalAttempted: totalAttempted,
			TotalSuccess:   len(imported),
			TotalFailed:    len(failed),
			SuccessRate:    successRate,
		},
	}, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
