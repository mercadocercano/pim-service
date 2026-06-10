package service

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// JSONAnalyzerService analiza archivos JSON y detecta tipos de datos
type JSONAnalyzerService struct {
	// Expresiones regulares para detección de tipos
	uuidRegex  *regexp.Regexp
	emailRegex *regexp.Regexp
	urlRegex   *regexp.Regexp
	dateRegex  *regexp.Regexp

	// Mapeos comunes de nombres de campos
	commonMappings map[string]string
}

// NewJSONAnalyzerService crea una nueva instancia del servicio
func NewJSONAnalyzerService() *JSONAnalyzerService {
	return &JSONAnalyzerService{
		uuidRegex:  regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`),
		emailRegex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
		urlRegex:   regexp.MustCompile(`^https?://[^\s]+$`),
		dateRegex:  regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(T\d{2}:\d{2}:\d{2})?`),
		commonMappings: map[string]string{
			// Español
			"producto":        "name",
			"nombre_producto": "name",
			"nombre":          "name",
			"codigo":          "sku",
			"código":          "sku",
			"codigo_producto": "sku",
			"precio":          "price",
			"precio_unitario": "price",
			"precio_unit":     "price",
			"descripcion":     "description",
			"descripción":     "description",
			"categoria":       "category_name",
			"categoría":       "category_name",
			"marca":           "brand_name",
			"inventario":      "stock",
			"existencias":     "stock",
			"cantidad":        "stock",

			// Inglés
			"product":      "name",
			"product_name": "name",
			"item":         "name",
			"code":         "sku",
			"product_code": "sku",
			"unit_price":   "price",
			"category":     "category_name",
			"brand":        "brand_name",
			"inventory":    "stock",
			"quantity":     "stock",
			"qty":          "stock",
		},
	}
}

// JSONAnalysis contiene el resultado del análisis JSON
type JSONAnalysis struct {
	IsArray      bool
	Fields       []string
	FieldTypes   map[string]DataType
	FieldSamples map[string][]string
	RecordCount  int
	Structure    string // "array" o "object"
}

// AnalyzeJSON analiza un archivo JSON y detecta tipos de datos
func (s *JSONAnalyzerService) AnalyzeJSON(reader io.Reader, maxRecords int) (*JSONAnalysis, error) {
	// Leer todo el contenido JSON
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON: %w", err)
	}

	// Intentar decodificar como array primero
	var arrayData []map[string]interface{}
	if err := json.Unmarshal(data, &arrayData); err == nil {
		return s.analyzeJSONArray(arrayData, maxRecords)
	}

	// Si no es array, intentar como objeto único
	var objectData map[string]interface{}
	if err := json.Unmarshal(data, &objectData); err == nil {
		// Convertir objeto único a array para análisis uniforme
		return s.analyzeJSONArray([]map[string]interface{}{objectData}, 1)
	}

	return nil, fmt.Errorf("invalid JSON format: must be array of objects or single object")
}

// analyzeJSONArray analiza un array de objetos JSON
func (s *JSONAnalyzerService) analyzeJSONArray(data []map[string]interface{}, maxRecords int) (*JSONAnalysis, error) {
	analysis := &JSONAnalysis{
		IsArray:      len(data) > 1,
		Fields:       []string{},
		FieldTypes:   make(map[string]DataType),
		FieldSamples: make(map[string][]string),
		RecordCount:  0,
		Structure:    "array",
	}

	if len(data) == 1 {
		analysis.Structure = "object"
	}

	// Recopilar todos los campos únicos
	fieldSet := make(map[string]bool)
	for i, record := range data {
		if i >= maxRecords {
			break
		}

		for field := range record {
			fieldSet[field] = true
		}
		analysis.RecordCount++
	}

	// Convertir set a slice
	for field := range fieldSet {
		analysis.Fields = append(analysis.Fields, field)
	}

	// Analizar tipos de datos para cada campo
	typeCounters := make(map[string]map[DataType]int)
	for _, field := range analysis.Fields {
		typeCounters[field] = make(map[DataType]int)
	}

	// Procesar registros para detectar tipos
	for i, record := range data {
		if i >= maxRecords {
			break
		}

		for field, value := range record {
			// Convertir valor a string para análisis
			strValue := s.valueToString(value)

			// Detectar tipo
			dataType := s.detectDataType(strValue)
			typeCounters[field][dataType]++

			// Guardar muestra
			if len(analysis.FieldSamples[field]) < 5 && strValue != "" {
				analysis.FieldSamples[field] = append(analysis.FieldSamples[field], strValue)
			}
		}
	}

	// Determinar tipo predominante por campo
	for field, types := range typeCounters {
		analysis.FieldTypes[field] = s.getMostFrequentType(types)
	}

	return analysis, nil
}

// valueToString convierte cualquier valor a string para análisis
func (s *JSONAnalyzerService) valueToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case float64:
		// JSON decodifica números como float64
		if v == float64(int64(v)) {
			return strconv.FormatInt(int64(v), 10)
		}
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case []interface{}, map[string]interface{}:
		// Para arrays y objetos anidados, convertir a JSON
		data, _ := json.Marshal(v)
		return string(data)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// detectDataType detecta el tipo de dato de un valor
func (s *JSONAnalyzerService) detectDataType(value string) DataType {
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
	if lowerValue == "true" || lowerValue == "false" {
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
		return DataTypeDate
	}

	// JSON anidado
	if strings.HasPrefix(value, "{") || strings.HasPrefix(value, "[") {
		return DataTypeString // Por ahora tratamos JSON anidado como string
	}

	// Default to string
	return DataTypeString
}

// getMostFrequentType obtiene el tipo más frecuente
func (s *JSONAnalyzerService) getMostFrequentType(typeCounts map[DataType]int) DataType {
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

// SuggestMapping sugiere mapeos basados en nombres de campos
func (s *JSONAnalyzerService) SuggestMapping(fieldName string) string {
	// Normalizar nombre
	normalized := strings.ToLower(strings.TrimSpace(fieldName))
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
func (s *JSONAnalyzerService) ValidateValue(value string, expectedType string) (bool, string) {
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

// GetExpectedType mapea tipos de datos detectados a tipos esperados del schema
func (s *JSONAnalyzerService) GetExpectedType(detectedType DataType) string {
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

// ConvertJSONToCSVFormat convierte datos JSON a formato compatible con CSV
func (s *JSONAnalyzerService) ConvertJSONToCSVFormat(data []map[string]interface{}) ([][]string, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("no data to convert")
	}

	// Recopilar todos los campos
	fieldSet := make(map[string]bool)
	for _, record := range data {
		for field := range record {
			fieldSet[field] = true
		}
	}

	// Crear headers
	headers := []string{}
	for field := range fieldSet {
		headers = append(headers, field)
	}

	// Crear filas
	rows := [][]string{headers}
	for _, record := range data {
		row := make([]string, len(headers))
		for i, header := range headers {
			if value, exists := record[header]; exists {
				row[i] = s.valueToString(value)
			} else {
				row[i] = ""
			}
		}
		rows = append(rows, row)
	}

	return rows, nil
}
