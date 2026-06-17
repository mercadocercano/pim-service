package usecase

import (
	"context"
	"fmt"
	"io"

	"saas-mt-pim-service/src/schema_validation/domain/entity"
	"saas-mt-pim-service/src/schema_validation/domain/service"
)

// ValidateExcelSchemaUseCase valida el schema de un archivo Excel (.xlsx)
type ValidateExcelSchemaUseCase struct {
	excelAnalyzer *service.ExcelAnalyzerService
	csvUseCase    *ValidateCSVSchemaUseCase
	schemaCache   SchemaValidationCache
}

func NewValidateExcelSchemaUseCase(
	excelAnalyzer *service.ExcelAnalyzerService,
	csvUseCase *ValidateCSVSchemaUseCase,
	schemaCache SchemaValidationCache,
) *ValidateExcelSchemaUseCase {
	return &ValidateExcelSchemaUseCase{
		excelAnalyzer: excelAnalyzer,
		csvUseCase:    csvUseCase,
		schemaCache:   schemaCache,
	}
}

func (uc *ValidateExcelSchemaUseCase) Execute(
	ctx context.Context,
	reader io.Reader,
	tenantID string,
	fileName string,
	maxRows int,
) (*entity.SchemaValidation, error) {
	if maxRows <= 0 {
		maxRows = 10
	}

	result, err := uc.excelAnalyzer.AnalyzeExcel(reader, maxRows)
	if err != nil {
		return nil, fmt.Errorf("error analyzing Excel: %w", err)
	}

	validation := entity.NewSchemaValidation(tenantID, fileName)
	validation.SourceFormat = "excel"
	validation.SheetName = result.SheetName

	uc.csvUseCase.validateColumns(validation, result.CSVAnalysis)

	preview := uc.csvUseCase.generateTablePreview(result.CSVAnalysis, validation)
	validation.SetTablePreview(preview)

	uc.csvUseCase.generateRecommendations(validation)
	validation.CalculateSummary()

	// Guardar en cache — ignorar error (no crítico)
	_ = uc.schemaCache.Set(ctx, validation)

	return validation, nil
}
