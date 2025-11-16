package exception

import "errors"

// Errores de validación
var (
	ErrMarketplaceSummaryInvalidName = errors.New("nombre de MarketplaceSummary inválido")
	ErrMarketplaceSummaryNameRequired = errors.New("nombre de MarketplaceSummary es requerido")
)

// Errores de negocio
var (
	ErrMarketplaceSummaryNotFound = errors.New("MarketplaceSummary no encontrado")
	ErrMarketplaceSummaryAlreadyExists = errors.New("MarketplaceSummary ya existe")
)

// Errores de persistencia
var (
	ErrMarketplaceSummaryCreateFailed = errors.New("error al crear MarketplaceSummary")
	ErrMarketplaceSummaryUpdateFailed = errors.New("error al actualizar MarketplaceSummary")
	ErrMarketplaceSummaryDeleteFailed = errors.New("error al eliminar MarketplaceSummary")
)
