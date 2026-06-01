package usecase

import (
	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

// BulkImportRow representa una fila del payload de importación masiva
type BulkImportRow struct {
	EAN          string   `json:"ean"`
	Name         string   `json:"name"`
	Description  *string  `json:"description,omitempty"`
	Brand        *string  `json:"brand,omitempty"`
	Category     *string  `json:"category,omitempty"`
	Price        *float64 `json:"price,omitempty"`
	ImageURL     *string  `json:"image_url,omitempty"`
	BusinessType *string  `json:"business_type,omitempty"`
}

// BulkImportError describe un error en una fila específica
type BulkImportError struct {
	Row     int    `json:"row"`
	Message string `json:"message"`
}

// BulkImportRequest contiene la lista de filas a importar
type BulkImportRequest struct {
	Rows   []BulkImportRow `json:"rows"`
	Source string          `json:"source"`
}

// BulkImportResponse resume el resultado de la importación
type BulkImportResponse struct {
	ImportedCount int               `json:"imported_count"`
	FailedCount   int               `json:"failed_count"`
	Errors        []BulkImportError `json:"errors"`
}

// BulkImportGlobalProducts importa productos en lote, acumulando errores por fila
type BulkImportGlobalProducts struct {
	createGlobalProduct *CreateGlobalProduct
	repo                port.GlobalProductRepository
}

// NewBulkImportGlobalProducts crea una nueva instancia del caso de uso
func NewBulkImportGlobalProducts(repo port.GlobalProductRepository) *BulkImportGlobalProducts {
	return &BulkImportGlobalProducts{
		createGlobalProduct: NewCreateGlobalProduct(repo),
		repo:                repo,
	}
}

// Execute procesa todas las filas, acumula errores y retorna el resumen
func (uc *BulkImportGlobalProducts) Execute(request BulkImportRequest) (*BulkImportResponse, error) {
	if len(request.Rows) == 0 {
		return nil, exception.NewValidationError("la lista de productos no puede estar vacía", nil)
	}

	source := request.Source
	if source == "" {
		source = "bulk_import"
	}

	response := &BulkImportResponse{
		Errors: make([]BulkImportError, 0),
	}

	for i, row := range request.Rows {
		rowNum := i + 1
		rowErr := uc.importRow(row, source)
		if rowErr != nil {
			response.FailedCount++
			response.Errors = append(response.Errors, BulkImportError{
				Row:     rowNum,
				Message: rowErr.Error(),
			})
			continue
		}
		response.ImportedCount++
	}

	return response, nil
}

func (uc *BulkImportGlobalProducts) importRow(row BulkImportRow, source string) error {
	req := CreateGlobalProductRequest{
		EAN:          row.EAN,
		Name:         row.Name,
		Description:  row.Description,
		Brand:        row.Brand,
		Category:     row.Category,
		Price:        row.Price,
		ImageURL:     row.ImageURL,
		Source:       source,
		BusinessType: row.BusinessType,
	}
	_, err := uc.createGlobalProduct.Execute(req)
	return err
}
