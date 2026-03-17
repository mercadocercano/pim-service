package service

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	uuidRegex    *regexp.Regexp
	emailRegex   *regexp.Regexp
	urlRegex     *regexp.Regexp
	dateRegex    *regexp.Regexp
	priceRegex   *regexp.Regexp
	skuRegex     *regexp.Regexp

	commonMappings map[string]string
}

func NewCSVAnalyzerService() *CSVAnalyzerService {
	return &CSVAnalyzerService{
		uuidRegex:  regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`),
		emailRegex: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
		urlRegex:   regexp.MustCompile(`^https?://[^\s]+$`),
		dateRegex:  regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(T\d{2}:\d{2}:\d{2})?`),
		priceRegex: regexp.MustCompile(`^\$?\s*[\d.,]+$`),
		skuRegex:   regexp.MustCompile(`^[A-Za-z0-9][-A-Za-z0-9_.]{1,18}[A-Za-z0-9]$`),
		commonMappings: buildCommonMappings(),
	}
}

func buildCommonMappings() map[string]string {
	return map[string]string{
		// === Nombre / Producto ===
		"producto":         "name",
		"nombre producto":  "name",
		"nombre_producto":  "name",
		"nombre":           "name",
		"articulo":         "name",
		"detalle":          "name",
		"item":             "name",
		"product":          "name",
		"product name":     "name",
		"product_name":     "name",
		"name":             "name",
		"title":            "name",

		// === SKU / Código ===
		"codigo":           "sku",
		"codigo producto":  "sku",
		"codigo_producto":  "sku",
		"cod":              "sku",
		"cod. barras":      "sku",
		"cod barras":       "sku",
		"codigo barras":    "sku",
		"ean":              "sku",
		"upc":              "sku",
		"barcode":          "sku",
		"code":             "sku",
		"product code":     "sku",
		"product_code":     "sku",
		"sku":              "sku",
		"variant sku":      "sku",

		// === Precio ===
		"precio":           "price",
		"precio unitario":  "price",
		"precio_unitario":  "price",
		"precio unit":      "price",
		"precio_unit":      "price",
		"p. venta":         "price",
		"p.venta":          "price",
		"p venta":          "price",
		"precio venta":     "price",
		"precio_venta":     "price",
		"imp venta":        "price",
		"imp_venta":        "price",
		"importe":          "price",
		"price":            "price",
		"unit price":       "price",
		"unit_price":       "price",
		"regular price":    "price",
		"variant price":    "price",
		"sale price":       "price",

		// === Descripción ===
		"descripcion":      "description",
		"description":      "description",
		"short description": "description",
		"body":             "description",
		"body (html)":      "description",

		// === Categoría ===
		"categoria":        "category_name",
		"rubro":            "category_name",
		"tipo":             "category_name",
		"category":         "category_name",
		"category name":    "category_name",
		"category_name":    "category_name",
		"categories":       "category_name",
		"type":             "category_name",

		// === Marca ===
		"marca":            "brand_name",
		"fabricante":       "brand_name",
		"brand":            "brand_name",
		"brand name":       "brand_name",
		"brand_name":       "brand_name",
		"vendor":           "brand_name",
		"manufacturer":     "brand_name",

		// === Stock ===
		"inventario":       "stock",
		"existencias":      "stock",
		"cantidad":         "stock",
		"stock":            "stock",
		"stock actual":     "stock",
		"stock_actual":     "stock",
		"stock act":        "stock",
		"stock_act":        "stock",
		"exist":            "stock",
		"inventory":        "stock",
		"quantity":         "stock",
		"qty":              "stock",
		"variant inventory qty": "stock",
	}
}

// detectDelimiter reads the first few lines and scores candidate delimiters.
// It picks the one that produces the most consistent (and >1) column count.
func (s *CSVAnalyzerService) detectDelimiter(data []byte) rune {
	candidates := []rune{',', ';', '|', '\t'}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	var sampleLines []string
	for scanner.Scan() && len(sampleLines) < 10 {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			sampleLines = append(sampleLines, line)
		}
	}

	if len(sampleLines) == 0 {
		return ','
	}

	// Handle || as a special case: pre-replace with a single rare char
	pipeDoubleSample := false
	if strings.Contains(sampleLines[0], "||") {
		pipeDoubleSample = true
	}

	type score struct {
		delim      rune
		avgCols    float64
		consistent bool
	}

	best := score{delim: ',', avgCols: 0, consistent: false}

	for _, delim := range candidates {
		lines := sampleLines
		if delim == '|' && pipeDoubleSample {
			// If || is present, split by || instead of |
			tmpLines := make([]string, len(sampleLines))
			for i, l := range sampleLines {
				tmpLines[i] = strings.ReplaceAll(l, "||", "\x00")
			}
			lines = tmpLines
			delim = '\x00'
		}

		counts := make([]int, 0, len(lines))
		for _, line := range lines {
			parts := strings.FieldsFunc(line, func(r rune) bool { return r == delim })
			counts = append(counts, len(parts))
		}

		if len(counts) == 0 {
			continue
		}

		headerCols := counts[0]
		if headerCols <= 1 {
			continue
		}

		consistent := true
		for _, c := range counts[1:] {
			if c != headerCols {
				consistent = false
				break
			}
		}

		avg := float64(headerCols)

		actualDelim := delim
		if delim == '\x00' {
			actualDelim = '|' // Mark as pipe; we'll handle || normalization in AnalyzeCSV
		}

		if (consistent && !best.consistent) ||
			(consistent == best.consistent && avg > best.avgCols) {
			best = score{delim: actualDelim, avgCols: avg, consistent: consistent}
		}
	}

	return best.delim
}

// AnalyzeCSV analiza un archivo CSV y detecta tipos de datos.
// Auto-detects the delimiter (, ; | || \t).
func (s *CSVAnalyzerService) AnalyzeCSV(reader io.Reader, maxRows int) (*CSVAnalysis, error) {
	rawData, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading CSV data: %w", err)
	}

	detectedDelim := s.detectDelimiter(rawData)

	// For ||, normalize to a single | before parsing
	effectiveData := rawData
	if detectedDelim == '|' && bytes.Contains(rawData, []byte("||")) {
		effectiveData = bytes.ReplaceAll(rawData, []byte("||"), []byte("|"))
	}

	csvReader := csv.NewReader(bytes.NewReader(effectiveData))
	csvReader.Comma = detectedDelim
	csvReader.TrimLeadingSpace = true
	csvReader.LazyQuotes = true

	headers, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading headers: %w", err)
	}

	// Trim possible BOM and whitespace from headers
	for i, h := range headers {
		headers[i] = strings.TrimSpace(strings.TrimPrefix(h, "\xef\xbb\xbf"))
	}

	analysis := &CSVAnalysis{
		Headers:       headers,
		ColumnTypes:   make(map[string]DataType),
		ColumnSamples: make(map[string][]string),
		RowCount:      0,
		Delimiter:     detectedDelim,
	}

	typeCounters := make(map[int]map[DataType]int)
	for i := range headers {
		typeCounters[i] = make(map[DataType]int)
	}

	rowsRead := 0
	for rowsRead < maxRows {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		rowsRead++

		for i, value := range row {
			if i >= len(headers) {
				continue
			}

			dataType := s.detectDataType(value)
			typeCounters[i][dataType]++

			header := headers[i]
			if len(analysis.ColumnSamples[header]) < 5 {
				analysis.ColumnSamples[header] = append(analysis.ColumnSamples[header], value)
			}
		}
	}

	analysis.RowCount = rowsRead

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

// removeAccents strips common Spanish/Portuguese diacritics
func removeAccents(s string) string {
	replacer := strings.NewReplacer(
		"á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u",
		"Á", "A", "É", "E", "Í", "I", "Ó", "O", "Ú", "U",
		"ñ", "n", "Ñ", "N", "ü", "u", "Ü", "U",
	)
	return replacer.Replace(s)
}

// normalizeColumnName normalizes a column name for matching
func normalizeColumnName(name string) string {
	n := strings.ToLower(strings.TrimSpace(name))
	n = removeAccents(n)
	n = strings.ReplaceAll(n, "_", " ")
	n = strings.ReplaceAll(n, "-", " ")
	// Collapse multiple spaces
	for strings.Contains(n, "  ") {
		n = strings.ReplaceAll(n, "  ", " ")
	}
	return n
}

// SuggestMapping suggests a field mapping based on column header name
func (s *CSVAnalyzerService) SuggestMapping(columnName string) string {
	normalized := normalizeColumnName(columnName)

	// Exact match (both sides normalized)
	for key, value := range s.commonMappings {
		if normalizeColumnName(key) == normalized {
			return value
		}
	}

	// Exact substring: the normalized name contains a known key (only for keys >=4 chars to avoid false positives)
	for key, value := range s.commonMappings {
		nk := normalizeColumnName(key)
		if len(nk) >= 4 && normalized == nk {
			return value
		}
	}

	return ""
}

// SuggestMappingByContent analyzes sample values to infer the field type when header-based mapping fails.
// It receives samples for the target column and all columns for cross-referencing cardinality.
func (s *CSVAnalyzerService) SuggestMappingByContent(samples []string, allColumnSamples map[string][]string) string {
	if len(samples) == 0 {
		return ""
	}

	if s.looksLikePrice(samples) {
		return "price"
	}
	if s.looksLikeSKU(samples) {
		return "sku"
	}
	if s.looksLikeStock(samples) {
		return "stock"
	}
	if s.looksLikeDescription(samples) {
		return "description"
	}
	if s.looksLikeLowCardinality(samples) {
		return "category_name"
	}
	if s.looksLikeName(samples) {
		return "name"
	}

	return ""
}

func (s *CSVAnalyzerService) looksLikePrice(samples []string) bool {
	numericCount := 0
	hasDecimal := false
	for _, v := range samples {
		v = strings.TrimSpace(v)
		cleaned := strings.ReplaceAll(strings.ReplaceAll(strings.TrimPrefix(v, "$"), ".", ""), ",", ".")
		cleaned = strings.TrimSpace(cleaned)
		if f, err := strconv.ParseFloat(cleaned, 64); err == nil {
			numericCount++
			if f != float64(int64(f)) || strings.Contains(v, ".") || strings.Contains(v, ",") {
				hasDecimal = true
			}
			if f < 0 || f > 100000000 {
				return false
			}
		}
	}
	return numericCount == len(samples) && len(samples) > 0 && (hasDecimal || strings.Contains(samples[0], "$"))
}

func (s *CSVAnalyzerService) looksLikeSKU(samples []string) bool {
	skuLike := 0
	uniqueValues := make(map[string]bool)
	for _, v := range samples {
		v = strings.TrimSpace(v)
		uniqueValues[v] = true
		if len(v) >= 3 && len(v) <= 20 {
			hasLetter := false
			hasDigit := false
			for _, r := range v {
				if unicode.IsLetter(r) {
					hasLetter = true
				}
				if unicode.IsDigit(r) {
					hasDigit = true
				}
			}
			if (hasLetter && hasDigit) || s.skuRegex.MatchString(v) {
				skuLike++
			} else if hasDigit && len(v) >= 7 && len(v) <= 14 {
				// EAN/UPC barcodes
				skuLike++
			}
		}
	}
	uniqueRatio := float64(len(uniqueValues)) / float64(len(samples))
	return skuLike > len(samples)/2 && uniqueRatio > 0.8
}

func (s *CSVAnalyzerService) looksLikeStock(samples []string) bool {
	intCount := 0
	hasValueAboveOne := false
	for _, v := range samples {
		v = strings.TrimSpace(v)
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			if n >= 0 && n <= 99999 {
				intCount++
				if n > 1 {
					hasValueAboveOne = true
				}
			}
		}
	}
	// Require at least one value > 1 to distinguish from boolean columns (0/1)
	return intCount == len(samples) && len(samples) > 0 && hasValueAboveOne
}

func (s *CSVAnalyzerService) looksLikeDescription(samples []string) bool {
	totalLen := 0
	for _, v := range samples {
		totalLen += len(v)
	}
	avgLen := float64(totalLen) / float64(len(samples))
	return avgLen > 50
}

func (s *CSVAnalyzerService) looksLikeLowCardinality(samples []string) bool {
	if len(samples) < 5 {
		return false
	}
	unique := make(map[string]bool)
	for _, v := range samples {
		v = strings.TrimSpace(v)
		// Skip empty values, very short values, and booleans
		if len(v) < 3 {
			return false
		}
		lv := strings.ToLower(v)
		if lv == "true" || lv == "false" || lv == "yes" || lv == "no" {
			return false
		}
		if _, err := strconv.ParseFloat(v, 64); err == nil {
			return false
		}
		unique[v] = true
	}
	uniqueRatio := float64(len(unique)) / float64(len(samples))
	return uniqueRatio < 0.4 && len(unique) >= 2
}

func (s *CSVAnalyzerService) looksLikeName(samples []string) bool {
	textCount := 0
	totalLen := 0
	unique := make(map[string]bool)
	for _, v := range samples {
		v = strings.TrimSpace(v)
		totalLen += len(v)
		unique[v] = true
		if len(v) >= 5 && len(v) <= 150 {
			if _, err := strconv.ParseFloat(v, 64); err != nil {
				hasSpace := strings.Contains(v, " ")
				if hasSpace {
					textCount++
				}
			}
		}
	}
	avgLen := float64(totalLen) / float64(len(samples))
	uniqueRatio := float64(len(unique)) / float64(len(samples))
	// Product names: medium-length text, mostly unique, contain spaces
	return textCount > len(samples)/2 && avgLen >= 8 && avgLen <= 80 && uniqueRatio > 0.7
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
	Delimiter     rune
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