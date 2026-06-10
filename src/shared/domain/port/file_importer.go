package port

import (
	"context"
	"io"
)

// FileImporter define la interfaz para importar archivos y convertirlos en entidades
type FileImporter[T any] interface {
	// Import procesa un archivo y devuelve el resultado de la importación
	Import(ctx context.Context, reader io.Reader, tenantID string) (*ImportResult[T], error)
}

// ImportResult contiene el resultado de una importación
type ImportResult[T any] struct {
	// TotalRows número total de filas procesadas
	TotalRows int `json:"total_rows"`

	// SuccessfulImports número de importaciones exitosas
	SuccessfulImports int `json:"successful_imports"`

	// FailedImports número de importaciones fallidas
	FailedImports int `json:"failed_imports"`

	// ImportedItems elementos importados exitosamente
	ImportedItems []T `json:"imported_items"`

	// Errors errores encontrados durante la importación
	Errors []ImportError `json:"errors"`
}

// ImportError representa un error durante la importación de una fila específica
type ImportError struct {
	// Row número de fila donde ocurrió el error (1-indexed)
	Row int `json:"row"`

	// Data datos de la fila que falló
	Data map[string]string `json:"data"`

	// Errors lista de errores encontrados en esta fila
	Errors []string `json:"errors"`
}

// NewImportResult crea un nuevo resultado de importación vacío
func NewImportResult[T any]() *ImportResult[T] {
	return &ImportResult[T]{
		ImportedItems: make([]T, 0),
		Errors:        make([]ImportError, 0),
	}
}

// AddSuccess agrega un elemento importado exitosamente
func (r *ImportResult[T]) AddSuccess(item T) {
	r.SuccessfulImports++
	r.ImportedItems = append(r.ImportedItems, item)
}

// AddError agrega un error de importación
func (r *ImportResult[T]) AddError(row int, data map[string]string, errors []string) {
	r.FailedImports++
	r.Errors = append(r.Errors, ImportError{
		Row:    row,
		Data:   data,
		Errors: errors,
	})
}

// IsSuccess retorna true si la importación fue completamente exitosa
func (r *ImportResult[T]) IsSuccess() bool {
	return r.FailedImports == 0
}

// HasErrors retorna true si hubo errores durante la importación
func (r *ImportResult[T]) HasErrors() bool {
	return r.FailedImports > 0
}
