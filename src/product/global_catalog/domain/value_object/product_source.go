package value_object

import (
	"errors"
	"strings"
	"time"
)

// ProductSource representa la fuente de un producto en el catálogo global
type ProductSource struct {
	source      string
	sourceURL   *string
	scrapedAt   *time.Time
	reliability float64
}

// Fuentes predefinidas
const (
	SourceManual    = "manual"
	SourceDisco     = "disco"
	SourceCarrefour = "carrefour"
	SourceFravega   = "fravega"
	SourceCoto      = "coto"
	SourceJumbo     = "jumbo"
	SourceAPI       = "api"
	// Nuevas fuentes agregadas
	SourcePigmento             = "scraper_pigmento"
	SourceFarmashop            = "scraper_farmashop"
	SourceJuleriaque           = "scraper_juleriaque"
	SourceAlmacenNatural       = "scraper_almacen_natural"
	SourcePolirubroOnline      = "scraper_polirubro_online"
	SourceWatchesWorld         = "scraper_watches_world"
	SourceRelojeriaSuiza       = "scraper_relojeria_suiza"
	SourceLubricentroOnline    = "scraper_lubricentro_online"
	SourceLubricantesOnline    = "scraper_lubricantes_online"
	SourceFerreteriaGeneralPaz = "scraper_ferreteria_general_paz"
	SourceCasaDelTornillo      = "scraper_casa_del_tornillo"
	SourceRelojeriaKronos      = "scraper_relojeria_kronos"
)

// NewProductSource crea un nuevo ProductSource
func NewProductSource(source string, sourceURL *string, scrapedAt *time.Time, reliability float64) (*ProductSource, error) {
	if source == "" {
		return nil, errors.New("la fuente del producto es obligatoria")
	}

	// Validar que la fuente esté en la lista de fuentes válidas
	if !isValidSource(source) {
		return nil, errors.New("fuente de producto no válida")
	}

	// Validar reliability
	if reliability < 0 || reliability > 1 {
		return nil, errors.New("la confiabilidad debe estar entre 0 y 1")
	}

	return &ProductSource{
		source:      strings.ToLower(source),
		sourceURL:   sourceURL,
		scrapedAt:   scrapedAt,
		reliability: reliability,
	}, nil
}

// NewManualSource crea una fuente manual con alta confiabilidad
func NewManualSource() (*ProductSource, error) {
	return NewProductSource(SourceManual, nil, nil, 1.0)
}

// NewScrapingSource crea una fuente de scraping con parámetros específicos
func NewScrapingSource(source string, sourceURL string, reliability float64) (*ProductSource, error) {
	now := time.Now()
	return NewProductSource(source, &sourceURL, &now, reliability)
}

// NewAPISource crea una fuente de API externa
func NewAPISource(sourceURL string, reliability float64) (*ProductSource, error) {
	return NewProductSource(SourceAPI, &sourceURL, nil, reliability)
}

// Source retorna la fuente del producto
func (ps *ProductSource) Source() string {
	return ps.source
}

// SourceURL retorna la URL de origen si existe
func (ps *ProductSource) SourceURL() *string {
	return ps.sourceURL
}

// ScrapedAt retorna cuándo fue scrapeado el producto
func (ps *ProductSource) ScrapedAt() *time.Time {
	return ps.scrapedAt
}

// Reliability retorna la confiabilidad de la fuente (0-1)
func (ps *ProductSource) Reliability() float64 {
	return ps.reliability
}

// IsManual indica si el producto fue agregado manualmente
func (ps *ProductSource) IsManual() bool {
	return ps.source == SourceManual
}

// IsScraped indica si el producto fue obtenido via scraping
func (ps *ProductSource) IsScraped() bool {
	scraped := []string{
		SourceDisco, SourceCarrefour, SourceFravega, SourceCoto, SourceJumbo,
		// Nuevas fuentes de scraping
		SourcePigmento, SourceFarmashop, SourceJuleriaque, SourceAlmacenNatural,
		SourcePolirubroOnline, SourceWatchesWorld, SourceRelojeriaSuiza,
		SourceLubricentroOnline, SourceLubricantesOnline, SourceFerreteriaGeneralPaz,
		SourceCasaDelTornillo, SourceRelojeriaKronos,
	}
	for _, s := range scraped {
		if ps.source == s {
			return true
		}
	}
	return false
}

// IsFromAPI indica si el producto proviene de una API externa
func (ps *ProductSource) IsFromAPI() bool {
	return ps.source == SourceAPI
}

// IsHighReliability indica si la fuente tiene alta confiabilidad (>= 0.7)
func (ps *ProductSource) IsHighReliability() bool {
	return ps.reliability >= 0.7
}

// NeedsUpdate indica si el producto necesita actualización (para fuentes scrapeadas)
func (ps *ProductSource) NeedsUpdate(maxAge time.Duration) bool {
	if !ps.IsScraped() || ps.scrapedAt == nil {
		return false
	}

	return time.Since(*ps.scrapedAt) > maxAge
}

// GetSourceDisplayName retorna un nombre legible de la fuente
func (ps *ProductSource) GetSourceDisplayName() string {
	switch ps.source {
	case SourceManual:
		return "Manual"
	case SourceDisco:
		return "Disco Argentina"
	case SourceCarrefour:
		return "Carrefour Argentina"
	case SourceFravega:
		return "Fravega"
	case SourceCoto:
		return "Coto Digital"
	case SourceJumbo:
		return "Jumbo Argentina"
	case SourceAPI:
		return "API Externa"
	// Nuevas fuentes agregadas
	case SourcePigmento:
		return "Pigmento Perfumería"
	case SourceFarmashop:
		return "Farmashop"
	case SourceJuleriaque:
		return "Juleriaque Belleza"
	case SourceAlmacenNatural:
		return "Almacén Natural"
	case SourcePolirubroOnline:
		return "Polirubro Online"
	case SourceWatchesWorld:
		return "Watches World"
	case SourceRelojeriaSuiza:
		return "Relojería Suiza"
	case SourceLubricentroOnline:
		return "Lubricentro Online"
	case SourceLubricantesOnline:
		return "Lubricantes Online"
	case SourceFerreteriaGeneralPaz:
		return "Ferretería General Paz"
	case SourceCasaDelTornillo:
		return "La Casa del Tornillo"
	case SourceRelojeriaKronos:
		return "Relojería Kronos"
	default:
		return "Desconocida"
	}
}

// isValidSource valida que la fuente esté en la lista de fuentes permitidas
func isValidSource(source string) bool {
	validSources := []string{
		SourceManual, SourceDisco, SourceCarrefour, SourceFravega,
		SourceCoto, SourceJumbo, SourceAPI,
		// Nuevas fuentes agregadas
		SourcePigmento, SourceFarmashop, SourceJuleriaque, SourceAlmacenNatural,
		SourcePolirubroOnline, SourceWatchesWorld, SourceRelojeriaSuiza,
		SourceLubricentroOnline, SourceLubricantesOnline, SourceFerreteriaGeneralPaz,
		SourceCasaDelTornillo, SourceRelojeriaKronos,
	}

	source = strings.ToLower(source)
	for _, validSource := range validSources {
		if source == validSource {
			return true
		}
	}
	return false
}

// Equals compara dos ProductSource
func (ps *ProductSource) Equals(other *ProductSource) bool {
	if other == nil {
		return false
	}
	return ps.source == other.source &&
		((ps.sourceURL == nil && other.sourceURL == nil) ||
			(ps.sourceURL != nil && other.sourceURL != nil && *ps.sourceURL == *other.sourceURL))
}

// String implementa la interfaz Stringer
func (ps *ProductSource) String() string {
	return ps.GetSourceDisplayName()
}
