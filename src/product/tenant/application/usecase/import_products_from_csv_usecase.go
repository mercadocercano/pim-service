package usecase

import (
	"context"
	"fmt"
	"io"
	"time"

	"saas-mt-pim-service/src/product/tenant/application/response"
	"saas-mt-pim-service/src/product/tenant/domain/entity"
	"saas-mt-pim-service/src/product/tenant/domain/port"
	sharedPort "saas-mt-pim-service/src/shared/domain/port"
	"saas-mt-pim-service/src/shared/infrastructure/metrics"
)

// ImportProductsFromCSVUseCase caso de uso para importar productos desde CSV
type ImportProductsFromCSVUseCase struct {
	productRepo  port.ProductCriteriaRepository
	fileImporter sharedPort.FileImporter[entity.Product]
}

// NewImportProductsFromCSVUseCase crea una nueva instancia del caso de uso
func NewImportProductsFromCSVUseCase(
	productRepo port.ProductCriteriaRepository,
	fileImporter sharedPort.FileImporter[entity.Product],
) *ImportProductsFromCSVUseCase {
	return &ImportProductsFromCSVUseCase{
		productRepo:  productRepo,
		fileImporter: fileImporter,
	}
}

// ImportResult resultado de la importación con productos guardados
type ImportResult struct {
	*sharedPort.ImportResult[entity.Product]
	SavedProducts    []*entity.Product          `json:"saved_products"`
	ProcessingErrors []response.ProcessingError `json:"processing_errors"`
}

// Execute ejecuta la importación de productos desde CSV
func (uc *ImportProductsFromCSVUseCase) Execute(ctx context.Context, reader io.Reader, tenantID string) (*ImportResult, error) {
	startTime := time.Now()

	// Validar tenant ID
	if tenantID == "" {
		return nil, fmt.Errorf("tenant ID es requerido")
	}

	// Importar productos desde CSV
	importResult, err := uc.fileImporter.Import(ctx, reader, tenantID)
	if err != nil {
		metrics.ImportOperationsTotal.WithLabelValues(tenantID, "csv_products", "failure").Inc()
		return nil, fmt.Errorf("error al importar archivo CSV: %w", err)
	}

	// Crear resultado extendido
	result := &ImportResult{
		ImportResult:     importResult,
		SavedProducts:    make([]*entity.Product, 0),
		ProcessingErrors: make([]response.ProcessingError, 0),
	}

	// Intentar guardar cada producto importado exitosamente
	for _, product := range importResult.ImportedItems {
		// Crear copia del producto para guardar
		productToSave := product // Copia del valor

		// Intentar guardar el producto
		err := uc.productRepo.Save(ctx, &productToSave)
		if err != nil {
			// Agregar error de procesamiento
			result.ProcessingErrors = append(result.ProcessingErrors, response.ProcessingError{
				Product: &productToSave,
				Error:   fmt.Sprintf("error al guardar producto: %v", err),
			})

			// Reducir el contador de éxitos y aumentar el de fallos
			result.SuccessfulImports--
			result.FailedImports++

			// Registrar error de validación
			metrics.RecordValidationError(tenantID, "csv_products", "save_error", "product")
		} else {
			// Agregar a productos guardados exitosamente
			result.SavedProducts = append(result.SavedProducts, &productToSave)
		}
	}

	// Actualizar la lista de items importados para reflejar solo los guardados
	result.ImportedItems = make([]entity.Product, 0, len(result.SavedProducts))
	for _, saved := range result.SavedProducts {
		result.ImportedItems = append(result.ImportedItems, *saved)
	}

	// Registrar métricas
	duration := time.Since(startTime).Seconds()
	success := len(result.ProcessingErrors) == 0 && err == nil
	metrics.RecordImportMetrics(
		tenantID,
		"csv_products",
		success,
		importResult.TotalRows,
		importResult.SuccessfulImports,
		importResult.FailedImports,
		duration,
	)

	// Registrar errores específicos de importación
	for _, importErr := range importResult.Errors {
		for _, err := range importErr.Errors {
			metrics.RecordValidationError(tenantID, "csv_products", "import_error", err)
		}
	}

	return result, nil
}

// ExecuteWithVariants ejecuta la importación y crea variantes por defecto
func (uc *ImportProductsFromCSVUseCase) ExecuteWithVariants(
	ctx context.Context,
	reader io.Reader,
	tenantID string,
	variantRepo port.ProductVariantRepository,
) (*ImportResult, error) {
	// Ejecutar importación base
	result, err := uc.Execute(ctx, reader, tenantID)
	if err != nil {
		return nil, err
	}

	// Para cada producto guardado, las variantes ya deberían estar creadas
	// El constructor de Product crea automáticamente una variante por defecto
	// Por lo tanto, no necesitamos crear variantes adicionales aquí

	return result, nil
}
