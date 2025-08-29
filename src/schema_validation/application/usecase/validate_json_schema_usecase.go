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

// ValidateJSONSchemaUseCase valida el schema de un archivo JSON
type ValidateJSONSchemaUseCase struct {
	jsonAnalyzer    *service.JSONAnalyzerService
	schemaCache     SchemaValidationCache
	productSchema   ProductSchemaDefinition
}

// NewValidateJSONSchemaUseCase crea una nueva instancia del caso de uso
func NewValidateJSONSchemaUseCase(
	jsonAnalyzer *service.JSONAnalyzerService,
	schemaCache SchemaValidationCache,
) *ValidateJSONSchemaUseCase {
	return &ValidateJSONSchemaUseCase{
		jsonAnalyzer: jsonAnalyzer,
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

// Execute ejecuta la validación del schema JSON
func (uc *ValidateJSONSchemaUseCase) Execute(
	ctx context.Context, 
	reader io.Reader, 
	tenantID string,
	fileName string,
	maxRecords int,
) (*entity.SchemaValidation, error) {
	
	if maxRecords <= 0 {
		maxRecords = 10 // Default
	}
	
	// Crear nueva validación
	validation := entity.NewSchemaValidation(tenantID, fileName)
	
	// Analizar JSON
	analysis, err := uc.jsonAnalyzer.AnalyzeJSON(reader, maxRecords)
	if err != nil {
		return nil, fmt.Errorf("error analyzing JSON: %w", err)
	}
	
	// Validar campos
	uc.validateFields(validation, analysis)
	
	// Generar preview de tabla
	preview := uc.generateTablePreview(analysis, validation)
	validation.SetTablePreview(preview)
	
	// Generar recomendaciones
	uc.generateRecommendations(validation, analysis)
	
	// Calcular resumen
	validation.CalculateSummary()
	
	// Guardar en cache
	if err := uc.schemaCache.Set(ctx, validation); err != nil {
		// Log error pero no fallar
		fmt.Printf("Error caching validation: %v\n", err)
	}
	
	return validation, nil
}

// validateFields valida los campos del JSON contra el schema esperado
func (uc *ValidateJSONSchemaUseCase) validateFields(
	validation *entity.SchemaValidation,
	analysis *service.JSONAnalysis,
) {
	// Mapear campos encontrados
	foundFields := make(map[string]bool)
	
	// Validar cada campo del JSON
	for i, jsonField := range analysis.Fields {
		fieldValidation := entity.NewColumnValidation(jsonField, i)
		
		// Detectar tipo de dato
		detectedType := analysis.FieldTypes[jsonField]
		fieldValidation.SetDetectedType(string(detectedType))
		
		// Agregar valores de muestra
		if samples, exists := analysis.FieldSamples[jsonField]; exists {
			for _, sample := range samples {
				fieldValidation.AddSampleValue(sample)
			}
		}
		
		// Buscar mapeo automático
		suggestedMapping := uc.jsonAnalyzer.SuggestMapping(jsonField)
		
		if suggestedMapping != "" {
			// Verificar si es campo requerido u opcional
			if colDef, isRequired := uc.productSchema.RequiredColumns[suggestedMapping]; isRequired {
				fieldValidation.SetRequired(true)
				fieldValidation.SetExpectedType(colDef.Type)
				fieldValidation.MapTo(suggestedMapping)
				foundFields[suggestedMapping] = true
				
				// Validar tipo
				expectedType := uc.jsonAnalyzer.GetExpectedType(detectedType)
				if expectedType != colDef.Type && colDef.Type != "string" {
					fieldValidation.AddIssue(fmt.Sprintf("Tipo detectado '%s' no coincide con esperado '%s'", expectedType, colDef.Type))
					fieldValidation.Status = value_object.ValidationStatusWarning
				} else {
					fieldValidation.Status = value_object.ValidationStatusValid
				}
			} else if colDef, isOptional := uc.productSchema.OptionalColumns[suggestedMapping]; isOptional {
				fieldValidation.SetRequired(false)
				fieldValidation.SetExpectedType(colDef.Type)
				fieldValidation.MapTo(suggestedMapping)
				foundFields[suggestedMapping] = true
				fieldValidation.Status = value_object.ValidationStatusValid
			}
			
			validation.AddSuggestedMapping(jsonField, suggestedMapping)
		} else {
			// Campo no reconocido
			fieldValidation.AddIssue("Campo no reconocido en el schema")
			fieldValidation.Status = value_object.ValidationStatusInfo
		}
		
		validation.AddColumn(fieldValidation)
	}
	
	// Verificar campos requeridos faltantes
	for reqField, colDef := range uc.productSchema.RequiredColumns {
		if !foundFields[reqField] {
			// Agregar campo faltante
			missingCol := entity.NewColumnValidation(reqField, -1)
			missingCol.SetRequired(true)
			missingCol.SetExpectedType(colDef.Type)
			missingCol.Status = value_object.ValidationStatusError
			missingCol.AddIssue("Campo requerido no encontrado en el JSON")
			
			validation.AddColumn(missingCol)
			validation.AddRecommendation(fmt.Sprintf("Agregar campo requerido '%s'", reqField))
		}
	}
}

// generateTablePreview genera la vista previa de la tabla
func (uc *ValidateJSONSchemaUseCase) generateTablePreview(
	analysis *service.JSONAnalysis,
	validation *entity.SchemaValidation,
) *entity.TablePreview {
	preview := &entity.TablePreview{
		Headers: make([]entity.HeaderInfo, 0),
		Rows:    make([]entity.RowPreview, 0),
	}
	
	// Generar headers
	for i, field := range analysis.Fields {
		if col, exists := validation.Columns[field]; exists {
			preview.Headers = append(preview.Headers, entity.HeaderInfo{
				Name:     field,
				Index:    i,
				Status:   string(col.Status),
				MappedTo: col.MappedTo,
			})
		}
	}
	
	// Simular algunas filas de preview basadas en los samples
	if len(analysis.FieldSamples) > 0 {
		// Crear hasta 5 filas de preview
		maxPreviewRows := 5
		if maxPreviewRows > analysis.RecordCount {
			maxPreviewRows = analysis.RecordCount
		}
		
		for rowNum := 0; rowNum < maxPreviewRows; rowNum++ {
			rowPreview := entity.RowPreview{
				RowNumber: rowNum + 1,
				Cells:     make([]entity.CellValidation, 0),
			}
			
			hasError := false
			hasWarning := false
			
			// Para cada campo
			for colIdx, field := range analysis.Fields {
				cell := entity.NewCellValidation("", colIdx, rowNum)
				
				// Obtener valor del sample si existe
				if samples, exists := analysis.FieldSamples[field]; exists && len(samples) > rowNum {
					cell.Value = samples[rowNum]
					
					// Validar celda
					if col, exists := validation.Columns[field]; exists {
						if col.MappedTo != "" && col.TypeExpected != "" {
							isValid, message := uc.jsonAnalyzer.ValidateValue(cell.Value, col.TypeExpected)
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
func (uc *ValidateJSONSchemaUseCase) generateRecommendations(validation *entity.SchemaValidation, analysis *service.JSONAnalysis) {
	// Recomendaciones de mapeo
	for jsonField, suggestion := range validation.SuggestedMappings {
		if col, exists := validation.Columns[jsonField]; exists && col.MappedTo == "" {
			validation.AddRecommendation(fmt.Sprintf("Mapear campo '%s' a '%s'", jsonField, suggestion))
		}
	}
	
	// Recomendaciones de tipos de datos
	for _, col := range validation.Columns {
		if col.Status == value_object.ValidationStatusWarning && len(col.Issues) > 0 {
			for _, issue := range col.Issues {
				if strings.Contains(issue, "tipo") {
					validation.AddRecommendation(fmt.Sprintf("Revisar formato de datos en campo '%s'", col.Name))
					break
				}
			}
		}
	}
	
	// Recomendaciones de estructura
	if analysis.Structure == "object" {
		validation.AddRecommendation("El archivo contiene un solo objeto. Para importación masiva, use un array de objetos")
	}
	
	// Recomendaciones generales
	missingOptional := []string{}
	for optField := range uc.productSchema.OptionalColumns {
		found := false
		for _, col := range validation.Columns {
			if col.MappedTo == optField {
				found = true
				break
			}
		}
		if !found {
			missingOptional = append(missingOptional, optField)
		}
	}
	
	if len(missingOptional) > 0 {
		validation.AddRecommendation(fmt.Sprintf("Considerar agregar campos opcionales: %s", strings.Join(missingOptional, ", ")))
	}
}