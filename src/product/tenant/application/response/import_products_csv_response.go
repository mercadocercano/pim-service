package response

import (
	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/shared/domain/port"
)

// ImportProductsCSVResponse respuesta de la importación de productos desde CSV
type ImportProductsCSVResponse struct {
	Success          bool                      `json:"success"`
	Summary          ImportSummary             `json:"summary"`
	Products         []ProductResponse         `json:"imported_products,omitempty"`
	Errors           []port.ImportError        `json:"errors,omitempty"`
	ProcessingErrors []ProcessingErrorResponse `json:"processing_errors,omitempty"`
}

// ImportSummary resumen de la importación
type ImportSummary struct {
	TotalRows         int `json:"total_rows"`
	SuccessfulImports int `json:"successful_imports"`
	FailedImports     int `json:"failed_imports"`
	SavedProducts     int `json:"saved_products"`
	ProcessingErrors  int `json:"processing_errors"`
}

// ProcessingError error durante el procesamiento/guardado
type ProcessingError struct {
	Product *entity.Product `json:"product"`
	Error   string          `json:"error"`
}

// ProcessingErrorResponse error durante el procesamiento
type ProcessingErrorResponse struct {
	ProductName string `json:"product_name"`
	ProductSKU  string `json:"product_sku"`
	Error       string `json:"error"`
}

// NewImportProductsCSVResponse crea una nueva respuesta de importación
func NewImportProductsCSVResponse(
	result *port.ImportResult[entity.Product],
	savedProducts []*entity.Product,
	processingErrors []ProcessingError,
) *ImportProductsCSVResponse {
	// Mapear productos guardados a responses
	products := make([]ProductResponse, 0, len(savedProducts))
	for _, product := range savedProducts {
		productResp := ProductResponse{
			ID:        product.IDString(),
			Name:      product.Name(),
			Status:    product.Status().Value(),
			CreatedAt: product.CreatedAt(),
			UpdatedAt: product.UpdatedAt(),
		}
		// SKU opcional
		if product.HasSKU() {
			sku := product.SKU().Value()
			productResp.SKU = &sku
		}
		products = append(products, productResp)
	}

	// Mapear errores de procesamiento
	procErrors := make([]ProcessingErrorResponse, 0, len(processingErrors))
	for _, err := range processingErrors {
		procError := ProcessingErrorResponse{
			Error: err.Error,
		}
		if err.Product != nil {
			procError.ProductName = err.Product.Name()
			if err.Product.HasSKU() {
				procError.ProductSKU = err.Product.SKU().Value()
			}
		}
		procErrors = append(procErrors, procError)
	}

	return &ImportProductsCSVResponse{
		Success: result.IsSuccess() && len(processingErrors) == 0,
		Summary: ImportSummary{
			TotalRows:         result.TotalRows,
			SuccessfulImports: result.SuccessfulImports,
			FailedImports:     result.FailedImports,
			SavedProducts:     len(savedProducts),
			ProcessingErrors:  len(processingErrors),
		},
		Products:         products,
		Errors:           result.Errors,
		ProcessingErrors: procErrors,
	}
}
