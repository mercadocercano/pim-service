package usecase

import (
	"context"
	"fmt"
	"io"
	"strings"

	"saas-mt-pim-service/src/schema_validation/domain/entity"
	"saas-mt-pim-service/src/schema_validation/domain/service"
	"saas-mt-pim-service/src/schema_validation/domain/value_object"
)

// ProductSchemaDefinition define el schema esperado para productos
type ProductSchemaDefinition struct {
	RequiredColumns map[string]ColumnDefinition
	OptionalColumns map[string]ColumnDefinition
}

// ColumnDefinition define las características de una columna
type ColumnDefinition struct {
	Name         string
	Type         string
	Required     bool
	Validations  []string
}

// ValidateCSVSchemaUseCase valida el schema de un archivo CSV
type ValidateCSVSchemaUseCase struct {
	csvAnalyzer    *service.CSVAnalyzerService
	schemaCache    SchemaValidationCache
	productSchema  ProductSchemaDefinition
}

// SchemaValidationCache define la interfaz para cache de validaciones
type SchemaValidationCache interface {
	Get(ctx context.Context, id string) (*entity.SchemaValidation, error)
	Set(ctx context.Context, validation *entity.SchemaValidation) error
}

// NewValidateCSVSchemaUseCase crea una nueva instancia del caso de uso
func NewValidateCSVSchemaUseCase(
	csvAnalyzer *service.CSVAnalyzerService,
	schemaCache SchemaValidationCache,
) *ValidateCSVSchemaUseCase {
	return &ValidateCSVSchemaUseCase{
		csvAnalyzer: csvAnalyzer,
		schemaCache: schemaCache,
		productSchema: ProductSchemaDefinition{
			RequiredColumns: map[string]ColumnDefinition{
				"name": {Name: "name", Type: "string", Required: true},
				"sku":  {Name: "sku", Type: "string", Required: true},
				"price": {Name: "price", Type: "number", Required: true},
			},
			OptionalColumns: map[string]ColumnDefinition{
				"description":    {Name: "description", Type: "string", Required: false},
				"category_id":    {Name: "category_id", Type: "uuid", Required: false},
				"category_name":  {Name: "category_name", Type: "string", Required: false},
				"brand_id":       {Name: "brand_id", Type: "uuid", Required: false},
				"brand_name":     {Name: "brand_name", Type: "string", Required: false},
				"stock":          {Name: "stock", Type: "integer", Required: false},
				"barcode":        {Name: "barcode", Type: "string", Required: false},
				"weight":         {Name: "weight", Type: "number", Required: false},
				"dimensions":     {Name: "dimensions", Type: "string", Required: false},
			},
		},
	}
}

// Execute ejecuta la validación del schema CSV
func (uc *ValidateCSVSchemaUseCase) Execute(
	ctx context.Context, 
	reader io.Reader, 
	tenantID string,
	fileName string,
	maxRows int,
) (*entity.SchemaValidation, error) {
	
	if maxRows <= 0 {
		maxRows = 10 // Default
	}
	
	// Crear nueva validación
	validation := entity.NewSchemaValidation(tenantID, fileName)
	
	// Analizar CSV
	analysis, err := uc.csvAnalyzer.AnalyzeCSV(reader, maxRows)
	if err != nil {
		return nil, fmt.Errorf("error analyzing CSV: %w", err)
	}
	
	// Validar columnas
	uc.validateColumns(validation, analysis)
	
	// Generar preview de tabla
	preview := uc.generateTablePreview(analysis, validation)
	validation.SetTablePreview(preview)
	
	// Generar recomendaciones
	uc.generateRecommendations(validation)
	
	// Calcular resumen
	validation.CalculateSummary()
	
	// Guardar en cache
	if err := uc.schemaCache.Set(ctx, validation); err != nil {
		// Log error pero no fallar
		fmt.Printf("Error caching validation: %v\n", err)
	}
	
	return validation, nil
}

// validateColumns valida las columnas del CSV contra el schema esperado
func (uc *ValidateCSVSchemaUseCase) validateColumns(
	validation *entity.SchemaValidation,
	analysis *service.CSVAnalysis,
) {
	// Mapear columnas encontradas
	foundColumns := make(map[string]bool)
	
	// Validar cada columna del CSV
	for i, csvColumn := range analysis.Headers {
		columnValidation := entity.NewColumnValidation(csvColumn, i)
		
		// Detectar tipo de dato
		detectedType := analysis.ColumnTypes[csvColumn]
		columnValidation.SetDetectedType(string(detectedType))
		
		// Agregar valores de muestra
		if samples, exists := analysis.ColumnSamples[csvColumn]; exists {
			for _, sample := range samples {
				columnValidation.AddSampleValue(sample)
			}
		}
		
		// Buscar mapeo automático
		suggestedMapping := uc.csvAnalyzer.SuggestMapping(csvColumn)
		
		if suggestedMapping != "" {
			// Verificar si es columna requerida u opcional
			if colDef, isRequired := uc.productSchema.RequiredColumns[suggestedMapping]; isRequired {
				columnValidation.SetRequired(true)
				columnValidation.SetExpectedType(colDef.Type)
				columnValidation.MapTo(suggestedMapping)
				foundColumns[suggestedMapping] = true
				
				// Validar tipo
				expectedType := uc.csvAnalyzer.GetExpectedType(detectedType)
				if expectedType != colDef.Type && colDef.Type != "string" {
					columnValidation.AddIssue(fmt.Sprintf("Tipo detectado '%s' no coincide con esperado '%s'", expectedType, colDef.Type))
					columnValidation.Status = value_object.ValidationStatusWarning
				} else {
					columnValidation.Status = value_object.ValidationStatusValid
				}
			} else if colDef, isOptional := uc.productSchema.OptionalColumns[suggestedMapping]; isOptional {
				columnValidation.SetRequired(false)
				columnValidation.SetExpectedType(colDef.Type)
				columnValidation.MapTo(suggestedMapping)
				foundColumns[suggestedMapping] = true
				columnValidation.Status = value_object.ValidationStatusValid
			}
			
			validation.AddSuggestedMapping(csvColumn, suggestedMapping)
		} else {
			// Columna no reconocida
			columnValidation.AddIssue("Columna no reconocida en el schema")
			columnValidation.Status = value_object.ValidationStatusInfo
		}
		
		validation.AddColumn(columnValidation)
	}
	
	// Verificar columnas requeridas faltantes
	for reqColumn, colDef := range uc.productSchema.RequiredColumns {
		if !foundColumns[reqColumn] {
			// Agregar columna faltante
			missingCol := entity.NewColumnValidation(reqColumn, -1)
			missingCol.SetRequired(true)
			missingCol.SetExpectedType(colDef.Type)
			missingCol.Status = value_object.ValidationStatusError
			missingCol.AddIssue("Columna requerida no encontrada en el CSV")
			
			validation.AddColumn(missingCol)
			validation.AddRecommendation(fmt.Sprintf("Agregar columna requerida '%s'", reqColumn))
		}
	}
}

// generateTablePreview genera la vista previa de la tabla
func (uc *ValidateCSVSchemaUseCase) generateTablePreview(
	analysis *service.CSVAnalysis,
	validation *entity.SchemaValidation,
) *entity.TablePreview {
	preview := &entity.TablePreview{
		Headers: make([]entity.HeaderInfo, 0),
		Rows:    make([]entity.RowPreview, 0),
	}
	
	// Generar headers
	for i, header := range analysis.Headers {
		if col, exists := validation.Columns[header]; exists {
			preview.Headers = append(preview.Headers, entity.HeaderInfo{
				Name:     header,
				Index:    i,
				Status:   string(col.Status),
				MappedTo: col.MappedTo,
			})
		}
	}
	
	// Aquí normalmente leeríamos más filas del CSV para el preview
	// Por ahora, generamos un preview básico con los samples
	
	// Simular algunas filas de preview basadas en los samples
	if len(analysis.ColumnSamples) > 0 {
		// Crear hasta 5 filas de preview
		maxPreviewRows := 5
		for rowNum := 0; rowNum < maxPreviewRows && rowNum < analysis.RowCount; rowNum++ {
			rowPreview := entity.RowPreview{
				RowNumber: rowNum + 1,
				Cells:     make([]entity.CellValidation, 0),
			}
			
			hasError := false
			hasWarning := false
			
			// Para cada columna
			for colIdx, header := range analysis.Headers {
				cell := entity.NewCellValidation("", colIdx, rowNum)
				
				// Obtener valor del sample si existe
				if samples, exists := analysis.ColumnSamples[header]; exists && len(samples) > rowNum {
					cell.Value = samples[rowNum]
					
					// Validar celda
					if col, exists := validation.Columns[header]; exists {
						if col.MappedTo != "" && col.TypeExpected != "" {
							isValid, message := uc.csvAnalyzer.ValidateValue(cell.Value, col.TypeExpected)
							if !isValid {
								cell.SetError(message)
								hasError = true
							}
						}
					}
				}
				
				rowPreview.Cells = append(rowPreview.Cells, *cell)
			}
			
			// Establecer estado de la fila
			if hasError {
				rowPreview.RowStatus = "error"
			} else if hasWarning {
				rowPreview.RowStatus = "warning"
			} else {
				rowPreview.RowStatus = "valid"
			}
			
			preview.Rows = append(preview.Rows, rowPreview)
		}
	}
	
	return preview
}

// generateRecommendations genera recomendaciones basadas en la validación
func (uc *ValidateCSVSchemaUseCase) generateRecommendations(validation *entity.SchemaValidation) {
	// Recomendaciones de mapeo
	for csvCol, suggestion := range validation.SuggestedMappings {
		if col, exists := validation.Columns[csvCol]; exists && col.MappedTo == "" {
			validation.AddRecommendation(fmt.Sprintf("Mapear columna '%s' a '%s'", csvCol, suggestion))
		}
	}
	
	// Recomendaciones de tipos de datos
	for _, col := range validation.Columns {
		if col.Status == value_object.ValidationStatusWarning && len(col.Issues) > 0 {
			for _, issue := range col.Issues {
				if strings.Contains(issue, "tipo") {
					validation.AddRecommendation(fmt.Sprintf("Revisar formato de datos en columna '%s'", col.Name))
					break
				}
			}
		}
	}
	
	// Recomendaciones generales
	missingOptional := []string{}
	for optCol := range uc.productSchema.OptionalColumns {
		found := false
		for _, col := range validation.Columns {
			if col.MappedTo == optCol {
				found = true
				break
			}
		}
		if !found {
			missingOptional = append(missingOptional, optCol)
		}
	}
	
	if len(missingOptional) > 0 {
		validation.AddRecommendation(fmt.Sprintf("Considerar agregar columnas opcionales: %s", strings.Join(missingOptional, ", ")))
	}
}

// GetValidationByID obtiene una validación del cache
func (uc *ValidateCSVSchemaUseCase) GetValidationByID(ctx context.Context, id string) (*entity.SchemaValidation, error) {
	return uc.schemaCache.Get(ctx, id)
}

// ApplyMapping aplica nuevos mapeos a una validación existente
func (uc *ValidateCSVSchemaUseCase) ApplyMapping(
	ctx context.Context,
	validationID string,
	mappings map[string]string,
) (*entity.SchemaValidation, error) {
	// Obtener validación del cache
	validation, err := uc.schemaCache.Get(ctx, validationID)
	if err != nil {
		return nil, fmt.Errorf("validation not found: %w", err)
	}
	
	// Aplicar mapeos
	validation.ApplyMappings(mappings)
	
	// Actualizar cache
	if err := uc.schemaCache.Set(ctx, validation); err != nil {
		return nil, fmt.Errorf("error updating cache: %w", err)
	}
	
	return validation, nil
}