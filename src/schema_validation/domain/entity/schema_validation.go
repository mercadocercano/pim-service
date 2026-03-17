package entity

import (
	"time"

	"github.com/google/uuid"
)

// SchemaValidation representa el resultado completo de validación de un archivo CSV
type SchemaValidation struct {
	// ID es el identificador único de la validación
	ID uuid.UUID `json:"id"`
	
	// TenantID es el ID del tenant
	TenantID string `json:"tenant_id"`
	
	// FileName es el nombre del archivo validado
	FileName string `json:"file_name"`
	
	// IsValid indica si el schema es válido
	IsValid bool `json:"is_valid"`
	
	// CanImport indica si se puede importar con advertencias
	CanImport bool `json:"can_import"`
	
	// Columns contiene la validación de cada columna
	Columns map[string]*ColumnValidation `json:"columns"`
	
	// TablePreview contiene la vista previa de la tabla
	TablePreview *TablePreview `json:"table_preview"`
	
	// Summary contiene el resumen de la validación
	Summary *ValidationSummary `json:"summary"`
	
	// Recommendations son las recomendaciones generadas
	Recommendations []string `json:"recommendations"`
	
	// SuggestedMappings son los mapeos sugeridos automáticamente
	SuggestedMappings map[string]string `json:"suggested_mappings"`

	// SourceFormat indica el formato de origen (csv, json, excel)
	SourceFormat string `json:"source_format,omitempty"`

	// SheetName es el nombre de la hoja de Excel usada (solo para Excel)
	SheetName string `json:"sheet_name,omitempty"`

	// DetectedDelimiter es el delimitador detectado en el CSV (ej: , ; | tab)
	DetectedDelimiter string `json:"detected_delimiter,omitempty"`

	// DeducedCategories maps row index → deduced category name (only when no category column mapped)
	DeducedCategories map[int]string `json:"deduced_categories,omitempty"`
	
	// CreatedAt es la fecha de creación
	CreatedAt time.Time `json:"created_at"`
	
	// ExpiresAt es la fecha de expiración (para cache)
	ExpiresAt time.Time `json:"expires_at"`
}

// TablePreview representa la vista previa de la tabla
type TablePreview struct {
	// Headers son los encabezados con su información
	Headers []HeaderInfo `json:"headers"`
	
	// Rows son las filas de preview
	Rows []RowPreview `json:"rows"`
}

// HeaderInfo contiene información de un encabezado
type HeaderInfo struct {
	Name     string `json:"name"`
	Index    int    `json:"index"`
	Status   string `json:"status"`
	MappedTo string `json:"mapped_to,omitempty"`
}

// RowPreview representa una fila de preview
type RowPreview struct {
	RowNumber int               `json:"row_number"`
	Cells     []CellValidation  `json:"cells"`
	RowStatus string           `json:"row_status"`
}

// ValidationSummary contiene el resumen de la validación
type ValidationSummary struct {
	TotalRows            int     `json:"total_rows"`
	ValidRows            int     `json:"valid_rows"`
	RowsWithErrors       int     `json:"rows_with_errors"`
	RowsWithWarnings     int     `json:"rows_with_warnings"`
	EstimatedSuccessRate float64 `json:"estimated_success_rate"`
	TotalColumns         int     `json:"total_columns"`
	MappedColumns        int     `json:"mapped_columns"`
	RequiredColumns      int     `json:"required_columns"`
}

// NewSchemaValidation crea una nueva validación de schema
func NewSchemaValidation(tenantID, fileName string) *SchemaValidation {
	now := time.Now()
	return &SchemaValidation{
		ID:                uuid.New(),
		TenantID:          tenantID,
		FileName:          fileName,
		IsValid:           false,
		CanImport:         false,
		Columns:           make(map[string]*ColumnValidation),
		Recommendations:   make([]string, 0),
		SuggestedMappings: make(map[string]string),
		CreatedAt:         now,
		ExpiresAt:         now.Add(30 * time.Minute), // TTL de 30 minutos
	}
}

// AddColumn agrega una validación de columna
func (sv *SchemaValidation) AddColumn(column *ColumnValidation) {
	sv.Columns[column.Name] = column
}

// SetTablePreview establece la vista previa
func (sv *SchemaValidation) SetTablePreview(preview *TablePreview) {
	sv.TablePreview = preview
}

// AddRecommendation agrega una recomendación
func (sv *SchemaValidation) AddRecommendation(recommendation string) {
	sv.Recommendations = append(sv.Recommendations, recommendation)
}

// AddSuggestedMapping agrega un mapeo sugerido
func (sv *SchemaValidation) AddSuggestedMapping(csvColumn, schemaField string) {
	sv.SuggestedMappings[csvColumn] = schemaField
}

// CalculateSummary calcula el resumen de la validación
func (sv *SchemaValidation) CalculateSummary() {
	summary := &ValidationSummary{
		TotalColumns: len(sv.Columns),
	}
	
	// Contar columnas mapeadas y requeridas
	for _, col := range sv.Columns {
		if col.MappedTo != "" {
			summary.MappedColumns++
		}
		if col.Required {
			summary.RequiredColumns++
		}
	}
	
	// Analizar filas si hay preview
	if sv.TablePreview != nil && len(sv.TablePreview.Rows) > 0 {
		summary.TotalRows = len(sv.TablePreview.Rows)
		
		for _, row := range sv.TablePreview.Rows {
			hasError := false
			hasWarning := false
			
			for _, cell := range row.Cells {
				switch cell.Status {
				case "error":
					hasError = true
				case "warning":
					hasWarning = true
				}
			}
			
			if hasError {
				summary.RowsWithErrors++
			} else if hasWarning {
				summary.RowsWithWarnings++
			} else {
				summary.ValidRows++
			}
		}
		
		// Calcular tasa de éxito estimada
		if summary.TotalRows > 0 {
			summary.EstimatedSuccessRate = float64(summary.ValidRows) / float64(summary.TotalRows) * 100
		}
	}
	
	sv.Summary = summary
	
	// Actualizar IsValid y CanImport
	sv.IsValid = summary.RowsWithErrors == 0
	sv.CanImport = summary.EstimatedSuccessRate >= 50 // Permitir importación si al menos 50% es válido
}

// IsExpired verifica si la validación ha expirado
func (sv *SchemaValidation) IsExpired() bool {
	return time.Now().After(sv.ExpiresAt)
}

// ApplyMappings aplica nuevos mapeos a las columnas
func (sv *SchemaValidation) ApplyMappings(mappings map[string]string) {
	for csvColumn, schemaField := range mappings {
		if column, exists := sv.Columns[csvColumn]; exists {
			column.MapTo(schemaField)
		}
	}
	// Recalcular resumen después de aplicar mapeos
	sv.CalculateSummary()
}