package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"saas-mt-pim-service/src/schema_validation/domain/entity"
	"saas-mt-pim-service/src/schema_validation/domain/value_object"
)

func TestNewSchemaValidation_CreatesWithDefaults(t *testing.T) {
	// Act
	validation := entity.NewSchemaValidation("tenant-123", "products.csv")

	// Assert
	assert.NotNil(t, validation)
	assert.Equal(t, "tenant-123", validation.TenantID)
	assert.Equal(t, "products.csv", validation.FileName)
	assert.False(t, validation.IsValid)
	assert.False(t, validation.CanImport)
	assert.Empty(t, validation.Columns)
	assert.Empty(t, validation.Recommendations)
	assert.Empty(t, validation.SuggestedMappings)
	assert.False(t, validation.IsExpired())
}

func TestSchemaValidation_AddColumn(t *testing.T) {
	// Arrange
	validation := entity.NewSchemaValidation("tenant-123", "test.csv")
	col := entity.NewColumnValidation("nombre", 0)

	// Act
	validation.AddColumn(col)

	// Assert
	assert.Len(t, validation.Columns, 1)
	assert.Contains(t, validation.Columns, "nombre")
}

func TestSchemaValidation_AddRecommendation(t *testing.T) {
	// Arrange
	validation := entity.NewSchemaValidation("tenant-123", "test.csv")

	// Act
	validation.AddRecommendation("Agregar columna de precio")
	validation.AddRecommendation("Revisar formato de SKU")

	// Assert
	assert.Len(t, validation.Recommendations, 2)
	assert.Contains(t, validation.Recommendations, "Agregar columna de precio")
}

func TestSchemaValidation_AddSuggestedMapping(t *testing.T) {
	// Arrange
	validation := entity.NewSchemaValidation("tenant-123", "test.csv")

	// Act
	validation.AddSuggestedMapping("nombre_producto", "name")
	validation.AddSuggestedMapping("precio_unitario", "price")

	// Assert
	assert.Len(t, validation.SuggestedMappings, 2)
	assert.Equal(t, "name", validation.SuggestedMappings["nombre_producto"])
	assert.Equal(t, "price", validation.SuggestedMappings["precio_unitario"])
}

func TestSchemaValidation_CalculateSummary_AllValid(t *testing.T) {
	// Arrange
	validation := entity.NewSchemaValidation("tenant-123", "test.csv")
	col := entity.NewColumnValidation("nombre", 0)
	col.MapTo("name")
	col.SetRequired(true)
	validation.AddColumn(col)

	preview := &entity.TablePreview{
		Headers: []entity.HeaderInfo{{Name: "nombre", Index: 0}},
		Rows: []entity.RowPreview{
			{RowNumber: 1, Cells: []entity.CellValidation{{Status: "valid"}}, RowStatus: "valid"},
			{RowNumber: 2, Cells: []entity.CellValidation{{Status: "valid"}}, RowStatus: "valid"},
		},
	}
	validation.SetTablePreview(preview)

	// Act
	validation.CalculateSummary()

	// Assert
	assert.NotNil(t, validation.Summary)
	assert.Equal(t, 2, validation.Summary.TotalRows)
	assert.Equal(t, 2, validation.Summary.ValidRows)
	assert.Equal(t, 0, validation.Summary.RowsWithErrors)
	assert.Equal(t, float64(100), validation.Summary.EstimatedSuccessRate)
	assert.True(t, validation.IsValid)
	assert.True(t, validation.CanImport)
}

func TestSchemaValidation_ApplyMappings(t *testing.T) {
	// Arrange
	validation := entity.NewSchemaValidation("tenant-123", "test.csv")
	col1 := entity.NewColumnValidation("col_a", 0)
	col2 := entity.NewColumnValidation("col_b", 1)
	validation.AddColumn(col1)
	validation.AddColumn(col2)

	// Act
	validation.ApplyMappings(map[string]string{
		"col_a": "name",
		"col_b": "price",
	})

	// Assert
	assert.Equal(t, "name", validation.Columns["col_a"].MappedTo)
	assert.Equal(t, "price", validation.Columns["col_b"].MappedTo)
}

func TestNewColumnValidation_Defaults(t *testing.T) {
	// Act
	col := entity.NewColumnValidation("test_col", 5)

	// Assert
	assert.Equal(t, "test_col", col.Name)
	assert.Equal(t, 5, col.Index)
	assert.Equal(t, value_object.ValidationStatusInfo, col.Status)
	assert.Empty(t, col.Issues)
	assert.Empty(t, col.SampleValues)
}

func TestColumnValidation_SetRequired(t *testing.T) {
	col := entity.NewColumnValidation("name", 0)
	col.SetRequired(true)
	assert.True(t, col.Required)
}

func TestColumnValidation_MapTo(t *testing.T) {
	col := entity.NewColumnValidation("nombre", 0)
	col.MapTo("name")
	assert.Equal(t, "name", col.MappedTo)
}

func TestColumnValidation_AddIssue(t *testing.T) {
	col := entity.NewColumnValidation("precio", 0)
	col.Status = value_object.ValidationStatusValid
	col.AddIssue("Tipo incompatible")

	assert.Len(t, col.Issues, 1)
	assert.Equal(t, value_object.ValidationStatusWarning, col.Status)
}

func TestColumnValidation_AddSampleValue_LimitedTo5(t *testing.T) {
	col := entity.NewColumnValidation("nombre", 0)
	for i := 0; i < 10; i++ {
		col.AddSampleValue("sample")
	}
	assert.Len(t, col.SampleValues, 5)
}
