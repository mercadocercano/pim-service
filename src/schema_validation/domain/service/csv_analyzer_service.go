package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// DataType representa el tipo de dato detectado
type DataType string

const (
	DataTypeString   DataType = "string"
	DataTypeInteger  DataType = "integer"
	DataTypeFloat    DataType = "float"
	DataTypeBoolean  DataType = "boolean"
	DataTypeDate     DataType = "date"
	DataTypeUUID     DataType = "uuid"
	DataTypeEmail    DataType = "email"
	DataTypeURL      DataType = "url"
	DataTypeUnknown  DataType = "unknown"
)

// CSVAnalyzerService analiza archivos CSV y detecta tipos de datos
type CSVAnalyzerService struct {
	// Expresiones regulares para detección de tipos
	uuidRegex  *regexp.Regexp
	emailRegex *regexp.Regexp
	urlRegex   *regexp.Regexp
	dateRegex  *regexp.Regexp
	
	// Mapeos comunes de nombres de columnas
	commonMappings map[string]string
}

// NewCSVAnalyzerService crea una nueva instancia del servicio
func NewCSVAnalyzerService() *CSVAnalyzerService {
	return &CSVAnalyzerService{
		uuidRegex:  regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`),
		emailRegex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
		urlRegex:   regexp.MustCompile(`^https?://[^\s]+$`),
		dateRegex:  regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(T\d{2}:\d{2}:\d{2})?`),
		commonMappings: map[string]string{
			// Español
			"producto":          "name",
			"nombre_producto":   "name",
			"nombre":           "name",
			"codigo":           "sku",
			"código":           "sku",
			"codigo_producto":  "sku",
			"precio":           "price",
			"precio_unitario":  "price",
			"precio_unit":      "price",
			"descripcion":      "description",
			"descripción":      "description",
			"categoria":        "category_name",
			"categoría":        "category_name",
			"marca":            "brand_name",
			"inventario":       "stock",
			"existencias":      "stock",
			"cantidad":         "stock",
			
			// Inglés
			"product":          "name",
			"product_name":     "name",
			"item":             "name",
			"code":             "sku",
			"product_code":     "sku",
			"unit_price":       "price",
			"category":         "category_name",
			"brand":            "brand_name",
			"inventory":        "stock",
			"quantity":         "stock",
			"qty":              "stock",
		},
	}
}

// AnalyzeCSV analiza un archivo CSV y detecta tipos de datos
func (s *CSVAnalyzerService) AnalyzeCSV(reader io.Reader, maxRows int) (*CSVAnalysis, error) {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true
	csvReader.LazyQuotes = true
	
	// Leer headers
	headers, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading headers: %w", err)
	}
	
	analysis := &CSVAnalysis{
		Headers:       headers,
		ColumnTypes:   make(map[string]DataType),
		ColumnSamples: make(map[string][]string),
		RowCount:      0,
	}
	
	// Inicializar contadores por columna
	typeCounters := make(map[int]map[DataType]int)
	for i := range headers {
		typeCounters[i] = make(map[DataType]int)
	}
	
	// Leer filas hasta maxRows
	rowsRead := 0
	for rowsRead < maxRows {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// Continuar con errores de formato
			continue
		}
		
		rowsRead++
		
		// Analizar cada celda
		for i, value := range row {
			if i >= len(headers) {
				continue
			}
			
			// Detectar tipo de dato
			dataType := s.detectDataType(value)
			typeCounters[i][dataType]++
			
			// Guardar muestra
			header := headers[i]
			if len(analysis.ColumnSamples[header]) < 5 {
				analysis.ColumnSamples[header] = append(analysis.ColumnSamples[header], value)
			}
		}
	}
	
	analysis.RowCount = rowsRead
	
	// Determinar tipo predominante por columna
	for i, header := range headers {
		analysis.ColumnTypes[header] = s.getMostFrequentType(typeCounters[i])
	}
	
	return analysis, nil
}

// detectDataType detecta el tipo de dato de un valor
func (s *CSVAnalyzerService) detectDataType(value string) DataType {
	value = strings.TrimSpace(value)
	
	// Vacío
	if value == "" {
		return DataTypeUnknown
	}
	
	// UUID
	if s.uuidRegex.MatchString(value) {
		return DataTypeUUID
	}
	
	// Email
	if s.emailRegex.MatchString(value) {
		return DataTypeEmail
	}
	
	// URL
	if s.urlRegex.MatchString(value) {
		return DataTypeURL
	}
	
	// Boolean
	lowerValue := strings.ToLower(value)
	if lowerValue == "true" || lowerValue == "false" || 
	   lowerValue == "yes" || lowerValue == "no" ||
	   lowerValue == "1" || lowerValue == "0" {
		return DataTypeBoolean
	}
	
	// Integer
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return DataTypeInteger
	}
	
	// Float
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return DataTypeFloat
	}
	
	// Date
	if s.dateRegex.MatchString(value) {
		if _, err := time.Parse("2006-01-02", value[:10]); err == nil {
			return DataTypeDate
		}
	}
	
	// Default to string
	return DataTypeString
}

// getMostFrequentType obtiene el tipo más frecuente
func (s *CSVAnalyzerService) getMostFrequentType(typeCounts map[DataType]int) DataType {
	var mostFrequent DataType = DataTypeString
	maxCount := 0
	
	for dataType, count := range typeCounts {
		if dataType != DataTypeUnknown && count > maxCount {
			mostFrequent = dataType
			maxCount = count
		}
	}
	
	return mostFrequent
}

// SuggestMapping sugiere mapeos basados en nombres de columnas
func (s *CSVAnalyzerService) SuggestMapping(columnName string) string {
	// Normalizar nombre
	normalized := strings.ToLower(strings.TrimSpace(columnName))
	normalized = strings.ReplaceAll(normalized, "_", " ")
	normalized = strings.ReplaceAll(normalized, "-", " ")
	
	// Buscar en mapeos comunes
	if mapping, exists := s.commonMappings[normalized]; exists {
		return mapping
	}
	
	// Buscar coincidencias parciales
	for key, value := range s.commonMappings {
		if strings.Contains(normalized, key) || strings.Contains(key, normalized) {
			return value
		}
	}
	
	return ""
}

// ValidateValue valida un valor según el tipo esperado
func (s *CSVAnalyzerService) ValidateValue(value string, expectedType string) (bool, string) {
	value = strings.TrimSpace(value)
	
	switch expectedType {
	case "string":
		if value == "" {
			return false, "valor vacío"
		}
		return true, ""
		
	case "integer":
		if _, err := strconv.ParseInt(value, 10, 64); err != nil {
			return false, "no es un número entero válido"
		}
		return true, ""
		
	case "float", "number":
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			return false, "no es un número válido"
		}
		return true, ""
		
	case "uuid":
		if !s.uuidRegex.MatchString(value) {
			return false, "no es un UUID válido"
		}
		return true, ""
		
	case "email":
		if !s.emailRegex.MatchString(value) {
			return false, "no es un email válido"
		}
		return true, ""
		
	case "boolean":
		lowerValue := strings.ToLower(value)
		validBools := []string{"true", "false", "yes", "no", "1", "0"}
		for _, valid := range validBools {
			if lowerValue == valid {
				return true, ""
			}
		}
		return false, "no es un valor booleano válido"
		
	default:
		return true, ""
	}
}

// CSVAnalysis contiene el resultado del análisis
type CSVAnalysis struct {
	Headers       []string
	ColumnTypes   map[string]DataType
	ColumnSamples map[string][]string
	RowCount      int
}

// GetExpectedType mapea tipos de datos detectados a tipos esperados del schema
func (s *CSVAnalyzerService) GetExpectedType(detectedType DataType) string {
	switch detectedType {
	case DataTypeInteger, DataTypeFloat:
		return "number"
	case DataTypeUUID:
		return "uuid"
	case DataTypeEmail:
		return "email"
	case DataTypeDate:
		return "date"
	case DataTypeBoolean:
		return "boolean"
	default:
		return "string"
	}
}