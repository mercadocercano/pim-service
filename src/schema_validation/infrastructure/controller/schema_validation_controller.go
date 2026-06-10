package controller

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"saas-mt-pim-service/src/schema_validation/application/request"
	"saas-mt-pim-service/src/schema_validation/application/response"
	"saas-mt-pim-service/src/schema_validation/application/usecase"
	"saas-mt-pim-service/src/schema_validation/domain/entity"
)

// SchemaValidationController maneja las peticiones HTTP para validación de schema
type SchemaValidationController struct {
	validateCSVSchemaUseCase    *usecase.ValidateCSVSchemaUseCase
	validateJSONSchemaUseCase   *usecase.ValidateJSONSchemaUseCase
	validateExcelSchemaUseCase  *usecase.ValidateExcelSchemaUseCase
	importFromValidationUseCase *usecase.ImportFromValidationUseCase
}

// NewSchemaValidationController crea una nueva instancia del controller
func NewSchemaValidationController(
	validateCSVSchemaUseCase *usecase.ValidateCSVSchemaUseCase,
	validateJSONSchemaUseCase *usecase.ValidateJSONSchemaUseCase,
	validateExcelSchemaUseCase *usecase.ValidateExcelSchemaUseCase,
	importFromValidationUseCase *usecase.ImportFromValidationUseCase,
) *SchemaValidationController {
	return &SchemaValidationController{
		validateCSVSchemaUseCase:    validateCSVSchemaUseCase,
		validateJSONSchemaUseCase:   validateJSONSchemaUseCase,
		validateExcelSchemaUseCase:  validateExcelSchemaUseCase,
		importFromValidationUseCase: importFromValidationUseCase,
	}
}

// ValidateSchema godoc
// @Summary Validar schema de archivo CSV o JSON
// @Description Valida el schema de un archivo CSV o JSON para importación de productos con visualización tipo spreadsheet
// @Tags schema-validation
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Archivo CSV o JSON"
// @Param use_ai formData bool false "Usar AI para análisis avanzado"
// @Param max_rows formData int false "Número máximo de filas a analizar" default(10)
// @Success 200 {object} response.SchemaValidationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/validate-schema [post]
// @Security BearerAuth
func (c *SchemaValidationController) ValidateSchema(ctx *gin.Context) {
	// Obtener tenant ID
	tenantID := ctx.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Obtener archivo
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener el archivo", "details": err.Error()})
		return
	}
	defer file.Close()

	// Determinar tipo de archivo
	filename := strings.ToLower(header.Filename)
	isCSV := strings.HasSuffix(filename, ".csv")
	isJSON := strings.HasSuffix(filename, ".json")
	isExcel := strings.HasSuffix(filename, ".xlsx") || strings.HasSuffix(filename, ".xls")

	if !isCSV && !isJSON && !isExcel {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato no soportado. Use CSV, Excel (.xlsx) o JSON."})
		return
	}

	// Obtener parámetros opcionales
	maxRows := 10
	if maxRowsStr := ctx.PostForm("max_rows"); maxRowsStr != "" {
		if parsed, err := parseInt(maxRowsStr); err == nil && parsed > 0 {
			maxRows = parsed
		}
	}

	// Leer archivo en buffer para poder analizarlo múltiples veces
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer archivo", "details": err.Error()})
		return
	}

	// Crear reader desde el buffer
	reader := bytes.NewReader(buf.Bytes())

	// Ejecutar validación según tipo de archivo
	var validation *entity.SchemaValidation

	if isExcel {
		validation, err = c.validateExcelSchemaUseCase.Execute(
			ctx.Request.Context(),
			reader,
			tenantID,
			header.Filename,
			maxRows,
		)
	} else if isCSV {
		validation, err = c.validateCSVSchemaUseCase.Execute(
			ctx.Request.Context(),
			reader,
			tenantID,
			header.Filename,
			maxRows,
		)
	} else {
		validation, err = c.validateJSONSchemaUseCase.Execute(
			ctx.Request.Context(),
			reader,
			tenantID,
			header.Filename,
			maxRows,
		)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al validar schema", "details": err.Error()})
		return
	}

	// Convertir a response DTO
	responseDTO := response.NewSchemaValidationResponse(validation)

	ctx.JSON(http.StatusOK, responseDTO)
}

// ApplyMapping godoc
// @Summary Aplicar mapeo de columnas
// @Description Aplica un mapeo manual de columnas y re-valida el schema
// @Tags schema-validation
// @Accept json
// @Produce json
// @Param body body request.ApplyMappingRequest true "Mapeos a aplicar"
// @Success 200 {object} response.SchemaValidationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/apply-mapping [post]
// @Security BearerAuth
func (c *SchemaValidationController) ApplyMapping(ctx *gin.Context) {
	var req request.ApplyMappingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos", "details": err.Error()})
		return
	}

	// Aplicar mapeos (funciona tanto para CSV como JSON)
	validation, err := c.validateCSVSchemaUseCase.ApplyMapping(
		ctx.Request.Context(),
		req.ValidationID,
		req.Mappings,
	)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Validación no encontrada o expirada"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al aplicar mapeos", "details": err.Error()})
		}
		return
	}

	// Convertir a response DTO
	responseDTO := response.NewSchemaValidationResponse(validation)

	ctx.JSON(http.StatusOK, responseDTO)
}

// GetCSVTemplate godoc
// @Summary Descargar plantilla CSV
// @Description Descarga una plantilla CSV con el formato correcto para importación
// @Tags schema-validation
// @Produce text/csv
// @Success 200 {file} string "Archivo CSV"
// @Router /products/csv-template [get]
func (c *SchemaValidationController) GetCSVTemplate(ctx *gin.Context) {
	// Crear CSV de ejemplo
	records := [][]string{
		{"name", "sku", "price", "description", "category_name", "brand_name", "stock", "barcode"},
		{"Laptop Dell XPS 13", "DELL-XPS-001", "1299.99", "Ultrabook profesional con pantalla táctil", "Laptops", "Dell", "25", "1234567890123"},
		{"Mouse Logitech MX Master", "LOG-MX-001", "79.99", "Mouse inalámbrico ergonómico", "Accesorios", "Logitech", "150", "1234567890124"},
		{"Teclado Mecánico RGB", "KEY-RGB-001", "149.99", "Teclado mecánico con retroiluminación RGB", "Accesorios", "Razer", "50", "1234567890125"},
	}

	// Configurar headers
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", "attachment; filename=product_import_template.csv")

	// Escribir CSV
	writer := csv.NewWriter(ctx.Writer)
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar CSV"})
			return
		}
	}
	writer.Flush()
}

// GetJSONTemplate godoc
// @Summary Descargar plantilla JSON
// @Description Descarga una plantilla JSON con el formato correcto para importación
// @Tags schema-validation
// @Produce application/json
// @Success 200 {object} []map[string]interface{} "Array de productos de ejemplo"
// @Router /products/json-template [get]
func (c *SchemaValidationController) GetJSONTemplate(ctx *gin.Context) {
	// Crear JSON de ejemplo
	products := []map[string]interface{}{
		{
			"name":          "Laptop Dell XPS 13",
			"sku":           "DELL-XPS-001",
			"price":         1299.99,
			"description":   "Ultrabook profesional con pantalla táctil",
			"category_name": "Laptops",
			"brand_name":    "Dell",
			"stock":         25,
			"barcode":       "1234567890123",
		},
		{
			"name":          "Mouse Logitech MX Master",
			"sku":           "LOG-MX-001",
			"price":         79.99,
			"description":   "Mouse inalámbrico ergonómico",
			"category_name": "Accesorios",
			"brand_name":    "Logitech",
			"stock":         150,
			"barcode":       "1234567890124",
		},
		{
			"name":          "Teclado Mecánico RGB",
			"sku":           "KEY-RGB-001",
			"price":         149.99,
			"description":   "Teclado mecánico con retroiluminación RGB",
			"category_name": "Accesorios",
			"brand_name":    "Razer",
			"stock":         50,
			"barcode":       "1234567890125",
		},
	}

	// Configurar headers
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Content-Disposition", "attachment; filename=product_import_template.json")

	// Enviar JSON
	ctx.JSON(http.StatusOK, products)
}

// ImportFromValidation godoc
// @Summary Importar productos desde una validación cacheada
// @Description Toma una validación previamente generada por validate-schema, aplica los mapeos de columnas y retorna los datos listos para importar
// @Tags schema-validation
// @Accept json
// @Produce json
// @Param body body usecase.ImportFromValidationRequest true "Request de importación"
// @Success 200 {object} usecase.ImportFromValidationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /products/import-from-validation [post]
// @Security BearerAuth
func (c *SchemaValidationController) ImportFromValidation(ctx *gin.Context) {
	tenantID := ctx.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	var req usecase.ImportFromValidationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos", "details": err.Error()})
		return
	}

	result, products, err := c.importFromValidationUseCase.Execute(ctx.Request.Context(), req, tenantID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "expired") {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result":   result,
		"products": products,
	})
}

// RegisterRoutes registra las rutas del controller
func (c *SchemaValidationController) RegisterRoutes(router *gin.RouterGroup) {
	products := router.Group("/products")
	{
		products.POST("/validate-schema", c.ValidateSchema)
		products.POST("/apply-mapping", c.ApplyMapping)
		products.POST("/import-from-validation", c.ImportFromValidation)
		products.GET("/csv-template", c.GetCSVTemplate)
		products.GET("/json-template", c.GetJSONTemplate)
	}
}

// parseInt helper para parsear strings a int
func parseInt(s string) (int, error) {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err
}
