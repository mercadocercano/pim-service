package response

import (
	"saas-mt-pim-service/src/schema_validation/domain/entity"
	"time"
)

// SchemaValidationResponse representa la respuesta de validación de schema
type SchemaValidationResponse struct {
	ID                string                            `json:"id"`
	IsValid           bool                              `json:"is_valid"`
	CanImport         bool                              `json:"can_import"`
	Columns           map[string]ColumnValidationDTO    `json:"columns"`
	TablePreview      *TablePreviewDTO                  `json:"table_preview"`
	Summary           *ValidationSummaryDTO             `json:"summary"`
	Recommendations   []string                          `json:"recommendations"`
	SuggestedMappings map[string]string                 `json:"suggested_mappings"`
	SourceFormat      string                            `json:"source_format,omitempty"`
	SheetName         string                            `json:"sheet_name,omitempty"`
	DetectedDelimiter string                            `json:"detected_delimiter,omitempty"`
	DeducedCategories map[int]string                    `json:"deduced_categories,omitempty"`
	CreatedAt         time.Time                         `json:"created_at"`
	ExpiresAt         time.Time                         `json:"expires_at"`
}

// ColumnValidationDTO representa la validación de una columna
type ColumnValidationDTO struct {
	Name            string                `json:"name"`
	Index           int                   `json:"index"`
	Status          string                `json:"status"`
	TypeExpected    string                `json:"type_expected"`
	TypeDetected    string                `json:"type_detected"`
	Required        bool                  `json:"required"`
	MappedTo        string                `json:"mapped_to,omitempty"`
	Statistics      ColumnStatisticsDTO   `json:"statistics"`
	Issues          []string              `json:"issues"`
	SampleValues    []string              `json:"sample_values"`
	InvalidExamples []InvalidExampleDTO   `json:"invalid_examples,omitempty"`
}

// ColumnStatisticsDTO contiene estadísticas de una columna
type ColumnStatisticsDTO struct {
	ValidCount   int     `json:"valid_count"`
	InvalidCount int     `json:"invalid_count"`
	NullCount    int     `json:"null_count"`
	UniqueCount  int     `json:"unique_count"`
	TotalCount   int     `json:"total_count"`
	ValidRate    float64 `json:"valid_rate"`
}

// InvalidExampleDTO representa un ejemplo de valor inválido
type InvalidExampleDTO struct {
	Row    int    `json:"row"`
	Value  string `json:"value"`
	Reason string `json:"reason"`
}

// TablePreviewDTO representa la vista previa de la tabla
type TablePreviewDTO struct {
	Headers []HeaderInfoDTO  `json:"headers"`
	Rows    []RowPreviewDTO  `json:"rows"`
}

// HeaderInfoDTO contiene información de un encabezado
type HeaderInfoDTO struct {
	Name     string `json:"name"`
	Index    int    `json:"index"`
	Status   string `json:"status"`
	MappedTo string `json:"mapped_to,omitempty"`
}

// RowPreviewDTO representa una fila de preview
type RowPreviewDTO struct {
	RowNumber int                  `json:"row_number"`
	Cells     []CellValidationDTO  `json:"cells"`
	RowStatus string               `json:"row_status"`
}

// CellValidationDTO representa la validación de una celda
type CellValidationDTO struct {
	Value          string `json:"value"`
	Status         string `json:"status"`
	Message        string `json:"message,omitempty"`
	CorrectedValue string `json:"corrected_value,omitempty"`
	ColumnIndex    int    `json:"column_index"`
}

// ValidationSummaryDTO contiene el resumen de la validación
type ValidationSummaryDTO struct {
	TotalRows            int     `json:"total_rows"`
	ValidRows            int     `json:"valid_rows"`
	RowsWithErrors       int     `json:"rows_with_errors"`
	RowsWithWarnings     int     `json:"rows_with_warnings"`
	EstimatedSuccessRate float64 `json:"estimated_success_rate"`
	TotalColumns         int     `json:"total_columns"`
	MappedColumns        int     `json:"mapped_columns"`
	RequiredColumns      int     `json:"required_columns"`
}

// NewSchemaValidationResponse crea una respuesta desde la entidad
func NewSchemaValidationResponse(validation *entity.SchemaValidation) *SchemaValidationResponse {
	response := &SchemaValidationResponse{
		ID:                validation.ID.String(),
		IsValid:           validation.IsValid,
		CanImport:         validation.CanImport,
		Columns:           make(map[string]ColumnValidationDTO),
		Recommendations:   validation.Recommendations,
		SuggestedMappings: validation.SuggestedMappings,
		SourceFormat:      validation.SourceFormat,
		SheetName:         validation.SheetName,
		DetectedDelimiter: validation.DetectedDelimiter,
		DeducedCategories: validation.DeducedCategories,
		CreatedAt:         validation.CreatedAt,
		ExpiresAt:         validation.ExpiresAt,
	}
	
	// Mapear columnas
	for name, col := range validation.Columns {
		response.Columns[name] = ColumnValidationDTO{
			Name:         col.Name,
			Index:        col.Index,
			Status:       string(col.Status),
			TypeExpected: col.TypeExpected,
			TypeDetected: col.TypeDetected,
			Required:     col.Required,
			MappedTo:     col.MappedTo,
			Statistics: ColumnStatisticsDTO{
				ValidCount:   col.Statistics.ValidCount,
				InvalidCount: col.Statistics.InvalidCount,
				NullCount:    col.Statistics.NullCount,
				UniqueCount:  col.Statistics.UniqueCount,
				TotalCount:   col.Statistics.TotalCount,
				ValidRate:    col.Statistics.ValidRate,
			},
			Issues:       col.Issues,
			SampleValues: col.SampleValues,
		}
		
		// Mapear ejemplos inválidos
		if len(col.InvalidExamples) > 0 {
			invalidExamples := make([]InvalidExampleDTO, len(col.InvalidExamples))
			for i, ex := range col.InvalidExamples {
				invalidExamples[i] = InvalidExampleDTO{
					Row:    ex.Row,
					Value:  ex.Value,
					Reason: ex.Reason,
				}
			}
			// Crear nueva estructura con ejemplos inválidos
			colDTO := response.Columns[name]
			colDTO.InvalidExamples = invalidExamples
			response.Columns[name] = colDTO
		}
	}
	
	// Mapear preview
	if validation.TablePreview != nil {
		response.TablePreview = &TablePreviewDTO{
			Headers: make([]HeaderInfoDTO, len(validation.TablePreview.Headers)),
			Rows:    make([]RowPreviewDTO, len(validation.TablePreview.Rows)),
		}
		
		// Headers
		for i, h := range validation.TablePreview.Headers {
			response.TablePreview.Headers[i] = HeaderInfoDTO{
				Name:     h.Name,
				Index:    h.Index,
				Status:   h.Status,
				MappedTo: h.MappedTo,
			}
		}
		
		// Rows
		for i, r := range validation.TablePreview.Rows {
			cells := make([]CellValidationDTO, len(r.Cells))
			for j, c := range r.Cells {
				cells[j] = CellValidationDTO{
					Value:          c.Value,
					Status:         string(c.Status),
					Message:        c.Message,
					CorrectedValue: c.CorrectedValue,
					ColumnIndex:    c.ColumnIndex,
				}
			}
			
			response.TablePreview.Rows[i] = RowPreviewDTO{
				RowNumber: r.RowNumber,
				Cells:     cells,
				RowStatus: r.RowStatus,
			}
		}
	}
	
	// Mapear resumen
	if validation.Summary != nil {
		response.Summary = &ValidationSummaryDTO{
			TotalRows:            validation.Summary.TotalRows,
			ValidRows:            validation.Summary.ValidRows,
			RowsWithErrors:       validation.Summary.RowsWithErrors,
			RowsWithWarnings:     validation.Summary.RowsWithWarnings,
			EstimatedSuccessRate: validation.Summary.EstimatedSuccessRate,
			TotalColumns:         validation.Summary.TotalColumns,
			MappedColumns:        validation.Summary.MappedColumns,
			RequiredColumns:      validation.Summary.RequiredColumns,
		}
	}
	
	return response
}