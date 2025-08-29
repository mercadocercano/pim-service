package adapters

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"saas-mt-pim-service/src/shared/domain/port"
)

// BaseCSVFileImporter implementación base para importar archivos CSV
type BaseCSVFileImporter[T any] struct {
	// delimiter delimitador de campos (por defecto ',')
	delimiter rune
	
	// hasHeader indica si el CSV tiene encabezados
	hasHeader bool
	
	// requiredColumns columnas requeridas en el CSV
	requiredColumns []string
}

// NewBaseCSVFileImporter crea una nueva instancia del importador base
func NewBaseCSVFileImporter[T any](delimiter rune, hasHeader bool, requiredColumns []string) *BaseCSVFileImporter[T] {
	if delimiter == 0 {
		delimiter = ','
	}
	return &BaseCSVFileImporter[T]{
		delimiter:       delimiter,
		hasHeader:       hasHeader,
		requiredColumns: requiredColumns,
	}
}

// Import procesa el archivo CSV y devuelve el resultado
func (b *BaseCSVFileImporter[T]) Import(ctx context.Context, reader io.Reader, tenantID string, parser RowParser[T]) (*port.ImportResult[T], error) {
	result := port.NewImportResult[T]()
	
	// Crear lector CSV
	csvReader := csv.NewReader(reader)
	csvReader.Comma = b.delimiter
	csvReader.TrimLeadingSpace = true
	
	// Leer todas las filas
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error al leer archivo CSV: %w", err)
	}
	
	if len(records) == 0 {
		return result, nil
	}
	
	// Procesar encabezados si existen
	var headers []string
	startRow := 0
	
	if b.hasHeader {
		headers = records[0]
		startRow = 1
		
		// Validar columnas requeridas
		if err := b.validateRequiredColumns(headers); err != nil {
			return nil, err
		}
	}
	
	// Procesar cada fila
	for i := startRow; i < len(records); i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			// Continuar procesando
		}
		
		row := records[i]
		rowNumber := i + 1 // 1-indexed para mejor legibilidad
		
		// Crear mapa de datos si hay headers
		rowData := b.createRowDataMap(headers, row)
		
		// Validar y parsear la fila
		item, validationErrors := parser.ParseRow(row, headers, rowData, tenantID)
		
		if len(validationErrors) > 0 {
			result.AddError(rowNumber, rowData, validationErrors)
		} else if item != nil {
			result.AddSuccess(*item)
		}
	}
	
	result.TotalRows = len(records) - startRow
	return result, nil
}

// validateRequiredColumns valida que todas las columnas requeridas estén presentes
func (b *BaseCSVFileImporter[T]) validateRequiredColumns(headers []string) error {
	headerMap := make(map[string]bool)
	for _, h := range headers {
		headerMap[strings.TrimSpace(strings.ToLower(h))] = true
	}
	
	missingColumns := []string{}
	for _, required := range b.requiredColumns {
		if !headerMap[strings.ToLower(required)] {
			missingColumns = append(missingColumns, required)
		}
	}
	
	if len(missingColumns) > 0 {
		return fmt.Errorf("columnas requeridas faltantes: %s", strings.Join(missingColumns, ", "))
	}
	
	return nil
}

// createRowDataMap crea un mapa de columna a valor para una fila
func (b *BaseCSVFileImporter[T]) createRowDataMap(headers []string, row []string) map[string]string {
	data := make(map[string]string)
	
	if len(headers) == 0 {
		// Si no hay headers, usar índices como claves
		for i, value := range row {
			data[fmt.Sprintf("column_%d", i)] = strings.TrimSpace(value)
		}
	} else {
		// Mapear headers a valores
		for i, header := range headers {
			if i < len(row) {
				data[strings.TrimSpace(header)] = strings.TrimSpace(row[i])
			} else {
				data[strings.TrimSpace(header)] = ""
			}
		}
	}
	
	return data
}

// RowParser define la interfaz para parsear una fila del CSV
type RowParser[T any] interface {
	// ParseRow convierte una fila en una entidad T
	// Retorna la entidad y una lista de errores de validación
	ParseRow(row []string, headers []string, rowData map[string]string, tenantID string) (*T, []string)
}

// CSVFileImporter interfaz concreta que deben implementar los importadores específicos
type CSVFileImporter[T any] interface {
	port.FileImporter[T]
	RowParser[T]
}