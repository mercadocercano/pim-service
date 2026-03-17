package service

import (
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

// ExcelAnalyzerService analiza archivos Excel (.xlsx) y produce la misma estructura CSVAnalysis
type ExcelAnalyzerService struct {
	csvAnalyzer *CSVAnalyzerService
}

// ExcelAnalysisResult extiende CSVAnalysis con metadata de Excel
type ExcelAnalysisResult struct {
	*CSVAnalysis
	SheetName string
}

func NewExcelAnalyzerService(csvAnalyzer *CSVAnalyzerService) *ExcelAnalyzerService {
	return &ExcelAnalyzerService{csvAnalyzer: csvAnalyzer}
}

// AnalyzeExcel parsea un archivo .xlsx y retorna la misma estructura que CSVAnalysis
func (s *ExcelAnalyzerService) AnalyzeExcel(reader io.Reader, maxRows int) (*ExcelAnalysisResult, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error opening Excel file: %w", err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, fmt.Errorf("no worksheets found in Excel file")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("error reading sheet '%s': %w", sheetName, err)
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("sheet '%s' is empty", sheetName)
	}

	headers := rows[0]
	if len(headers) == 0 {
		return nil, fmt.Errorf("no headers found in first row")
	}

	analysis := &CSVAnalysis{
		Headers:       headers,
		ColumnTypes:   make(map[string]DataType),
		ColumnSamples: make(map[string][]string),
		RowCount:      0,
	}

	typeCounters := make(map[int]map[DataType]int)
	for i := range headers {
		typeCounters[i] = make(map[DataType]int)
	}

	dataRows := rows[1:]
	rowsRead := 0
	for _, row := range dataRows {
		if rowsRead >= maxRows {
			break
		}
		rowsRead++

		for i, value := range row {
			if i >= len(headers) {
				continue
			}

			dataType := s.csvAnalyzer.detectDataType(value)
			typeCounters[i][dataType]++

			header := headers[i]
			if len(analysis.ColumnSamples[header]) < 5 {
				analysis.ColumnSamples[header] = append(analysis.ColumnSamples[header], value)
			}
		}
	}

	analysis.RowCount = rowsRead

	for i, header := range headers {
		analysis.ColumnTypes[header] = s.csvAnalyzer.getMostFrequentType(typeCounters[i])
	}

	return &ExcelAnalysisResult{
		CSVAnalysis: analysis,
		SheetName:   sheetName,
	}, nil
}
