package request

import (
	"mime/multipart"
)

// ValidateSchemaRequest representa la petición para validar un schema CSV
type ValidateSchemaRequest struct {
	// File es el archivo CSV a validar
	File multipart.File `form:"file" binding:"required"`
	
	// FileName es el nombre del archivo
	FileName string
	
	// UseAI indica si usar AI para análisis avanzado
	UseAI bool `form:"use_ai"`
	
	// MaxRows es el número máximo de filas a analizar
	MaxRows int `form:"max_rows"`
}

// ApplyMappingRequest representa la petición para aplicar mapeos
type ApplyMappingRequest struct {
	// ValidationID es el ID de la validación previa
	ValidationID string `json:"validation_id" binding:"required"`
	
	// Mappings son los mapeos de columnas CSV a campos del schema
	Mappings map[string]string `json:"mappings" binding:"required"`
}