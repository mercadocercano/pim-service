package service_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"saas-mt-pim-service/src/schema_validation/domain/service"
)

func TestCSVAnalyzerService_AnalyzeCSV_CommaDelimiter(t *testing.T) {
	// Arrange
	analyzer := service.NewCSVAnalyzerService()
	csvData := "nombre,precio,sku\nProducto 1,100.50,SKU001\nProducto 2,200.00,SKU002\n"
	reader := strings.NewReader(csvData)

	// Act
	result, err := analyzer.AnalyzeCSV(reader, 10)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 3, len(result.Headers))
	assert.Contains(t, result.Headers, "nombre")
	assert.Contains(t, result.Headers, "precio")
	assert.Contains(t, result.Headers, "sku")
	assert.Equal(t, 2, result.RowCount)
	assert.Equal(t, ',', result.Delimiter)
}

func TestCSVAnalyzerService_AnalyzeCSV_SemicolonDelimiter(t *testing.T) {
	// Arrange
	analyzer := service.NewCSVAnalyzerService()
	csvData := "nombre;precio;sku\nProducto 1;100.50;SKU001\nProducto 2;200.00;SKU002\n"
	reader := strings.NewReader(csvData)

	// Act
	result, err := analyzer.AnalyzeCSV(reader, 10)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 3, len(result.Headers))
	assert.Equal(t, ';', result.Delimiter)
}

func TestCSVAnalyzerService_AnalyzeCSV_DetectsColumnTypes(t *testing.T) {
	// Arrange
	analyzer := service.NewCSVAnalyzerService()
	csvData := "nombre,precio,stock,activo\nWidget,99.99,50,true\nGadget,149.50,25,false\n"
	reader := strings.NewReader(csvData)

	// Act
	result, err := analyzer.AnalyzeCSV(reader, 10)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, service.DataTypeString, result.ColumnTypes["nombre"])
	assert.Equal(t, service.DataTypeFloat, result.ColumnTypes["precio"])
	assert.Equal(t, service.DataTypeInteger, result.ColumnTypes["stock"])
	assert.Equal(t, service.DataTypeBoolean, result.ColumnTypes["activo"])
}

func TestCSVAnalyzerService_AnalyzeCSV_CollectsSamples(t *testing.T) {
	// Arrange
	analyzer := service.NewCSVAnalyzerService()
	csvData := "nombre,precio\nProducto A,10\nProducto B,20\nProducto C,30\n"
	reader := strings.NewReader(csvData)

	// Act
	result, err := analyzer.AnalyzeCSV(reader, 10)

	// Assert
	require.NoError(t, err)
	assert.Len(t, result.ColumnSamples["nombre"], 3)
	assert.Contains(t, result.ColumnSamples["nombre"], "Producto A")
}

func TestCSVAnalyzerService_AnalyzeCSV_MaxRows(t *testing.T) {
	// Arrange
	analyzer := service.NewCSVAnalyzerService()
	csvData := "nombre\nA\nB\nC\nD\nE\n"
	reader := strings.NewReader(csvData)

	// Act
	result, err := analyzer.AnalyzeCSV(reader, 2)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, 2, result.RowCount)
}

func TestCSVAnalyzerService_AnalyzeCSV_EmptyCSV_ShouldFail(t *testing.T) {
	// Arrange
	analyzer := service.NewCSVAnalyzerService()
	reader := strings.NewReader("")

	// Act
	_, err := analyzer.AnalyzeCSV(reader, 10)

	// Assert
	assert.Error(t, err)
}

func TestCSVAnalyzerService_SuggestMapping_KnownColumns(t *testing.T) {
	// Arrange
	analyzer := service.NewCSVAnalyzerService()

	tests := []struct {
		columnName     string
		expectedMapping string
	}{
		{"nombre", "name"},
		{"producto", "name"},
		{"precio", "price"},
		{"sku", "sku"},
		{"codigo", "sku"},
		{"descripcion", "description"},
		{"marca", "brand_name"},
		{"categoria", "category_name"},
		{"stock", "stock"},
		{"inventario", "stock"},
	}

	for _, tt := range tests {
		t.Run(tt.columnName, func(t *testing.T) {
			// Act
			result := analyzer.SuggestMapping(tt.columnName)

			// Assert
			assert.Equal(t, tt.expectedMapping, result)
		})
	}
}

func TestCSVAnalyzerService_SuggestMapping_UnknownColumn(t *testing.T) {
	// Arrange
	analyzer := service.NewCSVAnalyzerService()

	// Act
	result := analyzer.SuggestMapping("campo_desconocido_xyz")

	// Assert
	assert.Empty(t, result)
}

func TestCSVAnalyzerService_ValidateValue_String(t *testing.T) {
	analyzer := service.NewCSVAnalyzerService()

	isValid, _ := analyzer.ValidateValue("Producto A", "string")
	assert.True(t, isValid)

	isValid, msg := analyzer.ValidateValue("", "string")
	assert.False(t, isValid)
	assert.Contains(t, msg, "vacío")
}

func TestCSVAnalyzerService_ValidateValue_Number(t *testing.T) {
	analyzer := service.NewCSVAnalyzerService()

	isValid, _ := analyzer.ValidateValue("99.99", "number")
	assert.True(t, isValid)

	isValid, _ = analyzer.ValidateValue("100", "number")
	assert.True(t, isValid)

	isValid, msg := analyzer.ValidateValue("no-es-numero", "number")
	assert.False(t, isValid)
	assert.Contains(t, msg, "número")
}

func TestCSVAnalyzerService_ValidateValue_Integer(t *testing.T) {
	analyzer := service.NewCSVAnalyzerService()

	isValid, _ := analyzer.ValidateValue("42", "integer")
	assert.True(t, isValid)

	isValid, msg := analyzer.ValidateValue("3.14", "integer")
	assert.False(t, isValid)
	assert.Contains(t, msg, "entero")
}

func TestCSVAnalyzerService_ValidateValue_UUID(t *testing.T) {
	analyzer := service.NewCSVAnalyzerService()

	isValid, _ := analyzer.ValidateValue("123e4567-e89b-12d3-a456-426614174000", "uuid")
	assert.True(t, isValid)

	isValid, msg := analyzer.ValidateValue("not-a-uuid", "uuid")
	assert.False(t, isValid)
	assert.Contains(t, msg, "UUID")
}

func TestCSVAnalyzerService_GetExpectedType(t *testing.T) {
	analyzer := service.NewCSVAnalyzerService()

	assert.Equal(t, "number", analyzer.GetExpectedType(service.DataTypeInteger))
	assert.Equal(t, "number", analyzer.GetExpectedType(service.DataTypeFloat))
	assert.Equal(t, "uuid", analyzer.GetExpectedType(service.DataTypeUUID))
	assert.Equal(t, "string", analyzer.GetExpectedType(service.DataTypeString))
	assert.Equal(t, "boolean", analyzer.GetExpectedType(service.DataTypeBoolean))
}
