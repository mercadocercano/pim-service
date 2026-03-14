package exception

import "errors"

// Errores de validación
var (
	ErrMarketplaceSummaryInvalidName = errors.New("nombre de MarketplaceSummary inválido")
	ErrMarketplaceSummaryNameRequired = errors.New("nombre de MarketplaceSummary es requerido")
)

// Errores de negocio
var (
	ErrMarketplaceSummaryNotFound      = errors.New("MarketplaceSummary no encontrado")
	ErrMarketplaceSummaryAlreadyExists = errors.New("MarketplaceSummary ya existe")
	ErrAttributeStatsNotFound          = errors.New("AttributeStats no encontrado")
	ErrBrandStatsNotFound              = errors.New("BrandStats no encontrado")
	ErrCategoryStatsNotFound           = errors.New("CategoryStats no encontrado")
	ErrProductStatsNotFound            = errors.New("ProductStats no encontrado")
)

// Errores de persistencia
var (
	ErrMarketplaceSummaryCreateFailed = errors.New("error al crear MarketplaceSummary")
	ErrMarketplaceSummaryUpdateFailed = errors.New("error al actualizar MarketplaceSummary")
	ErrMarketplaceSummaryDeleteFailed = errors.New("error al eliminar MarketplaceSummary")
	ErrAttributeStatsCreateFailed     = errors.New("error al crear AttributeStats")
	ErrAttributeStatsUpdateFailed     = errors.New("error al actualizar AttributeStats")
	ErrAttributeStatsDeleteFailed     = errors.New("error al eliminar AttributeStats")
	ErrBrandStatsCreateFailed         = errors.New("error al crear BrandStats")
	ErrBrandStatsUpdateFailed         = errors.New("error al actualizar BrandStats")
	ErrBrandStatsDeleteFailed         = errors.New("error al eliminar BrandStats")
	ErrCategoryStatsCreateFailed      = errors.New("error al crear CategoryStats")
	ErrCategoryStatsUpdateFailed      = errors.New("error al actualizar CategoryStats")
	ErrCategoryStatsDeleteFailed      = errors.New("error al eliminar CategoryStats")
	ErrProductStatsCreateFailed       = errors.New("error al crear ProductStats")
	ErrProductStatsUpdateFailed       = errors.New("error al actualizar ProductStats")
	ErrProductStatsDeleteFailed       = errors.New("error al eliminar ProductStats")
)
