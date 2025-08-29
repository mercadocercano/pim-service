package entity

import (
	"errors"
	"strings"
	"time"

	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"

	"github.com/google/uuid"
)

// GlobalProduct representa un producto en el catálogo global argentino
type GlobalProduct struct {
	id            uuid.UUID
	ean           *value_object.EAN13
	name          string
	description   *string
	brand         *string
	category      *string
	price         *float64
	imageURL      *string
	imageURLs     []string
	source        *value_object.ProductSource
	qualityScore  *value_object.QualityScore
	isVerified    bool
	isActive      bool
	businessType  *string // tipo de negocio recomendado
	tags          []string
	metadata      map[string]interface{}
	createdAt     time.Time
	updatedAt     time.Time
	lastScrapedAt *time.Time
}

// NewGlobalProduct crea un nuevo producto del catálogo global
func NewGlobalProduct(
	ean *value_object.EAN13,
	name string,
	description *string,
	brand *string,
	category *string,
	price *float64,
	imageURL *string,
	source *value_object.ProductSource,
) (*GlobalProduct, error) {
	if ean == nil {
		return nil, errors.New("EAN-13 es obligatorio")
	}

	if name == "" {
		return nil, errors.New("el nombre del producto es obligatorio")
	}

	if source == nil {
		return nil, errors.New("la fuente del producto es obligatoria")
	}

	id := uuid.New()

	// Calcular quality score basado en los datos disponibles
	qualityScore := calculateInitialQualityScore(name, description, brand, category, price, imageURL, ean)

	globalProduct := &GlobalProduct{
		id:            id,
		ean:           ean,
		name:          strings.TrimSpace(name),
		description:   description,
		brand:         brand,
		category:      category,
		price:         price,
		imageURL:      imageURL,
		imageURLs:     make([]string, 0),
		source:        source,
		qualityScore:  qualityScore,
		isVerified:    false,
		isActive:      true,
		businessType:  nil,
		tags:          make([]string, 0),
		metadata:      make(map[string]interface{}),
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
		lastScrapedAt: source.ScrapedAt(),
	}

	// Si hay imagen principal, agregarla a la lista
	if imageURL != nil && *imageURL != "" {
		globalProduct.imageURLs = append(globalProduct.imageURLs, *imageURL)
	}

	return globalProduct, nil
}

// NewGlobalProductFromRepository crea una instancia desde datos de repositorio
func NewGlobalProductFromRepository(
	id uuid.UUID,
	ean *value_object.EAN13,
	name string,
	description *string,
	brand *string,
	category *string,
	price *float64,
	imageURL *string,
	imageURLs []string,
	source *value_object.ProductSource,
	qualityScore *value_object.QualityScore,
	isVerified bool,
	isActive bool,
	businessType *string,
	tags []string,
	metadata map[string]interface{},
	createdAt time.Time,
	updatedAt time.Time,
	lastScrapedAt *time.Time,
) (*GlobalProduct, error) {
	// EAN es opcional para productos del catálogo global
	// ya que algunos productos scrapeados no tienen EAN disponible

	if name == "" {
		return nil, errors.New("el nombre del producto es obligatorio")
	}

	if source == nil {
		return nil, errors.New("la fuente del producto es obligatoria")
	}

	if qualityScore == nil {
		qualityScore = calculateInitialQualityScore(name, description, brand, category, price, imageURL, ean)
	}

	return &GlobalProduct{
		id:            id,
		ean:           ean,
		name:          name,
		description:   description,
		brand:         brand,
		category:      category,
		price:         price,
		imageURL:      imageURL,
		imageURLs:     imageURLs,
		source:        source,
		qualityScore:  qualityScore,
		isVerified:    isVerified,
		isActive:      isActive,
		businessType:  businessType,
		tags:          tags,
		metadata:      metadata,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		lastScrapedAt: lastScrapedAt,
	}, nil
}

// Getters
func (gp *GlobalProduct) ID() uuid.UUID                            { return gp.id }
func (gp *GlobalProduct) IDString() string                         { return gp.id.String() }
func (gp *GlobalProduct) EAN() *value_object.EAN13                 { return gp.ean }
func (gp *GlobalProduct) Name() string                             { return gp.name }
func (gp *GlobalProduct) Description() *string                     { return gp.description }
func (gp *GlobalProduct) Brand() *string                           { return gp.brand }
func (gp *GlobalProduct) Category() *string                        { return gp.category }
func (gp *GlobalProduct) Price() *float64                          { return gp.price }
func (gp *GlobalProduct) ImageURL() *string                        { return gp.imageURL }
func (gp *GlobalProduct) ImageURLs() []string                      { return gp.imageURLs }
func (gp *GlobalProduct) Source() *value_object.ProductSource      { return gp.source }
func (gp *GlobalProduct) QualityScore() *value_object.QualityScore { return gp.qualityScore }
func (gp *GlobalProduct) IsVerified() bool                         { return gp.isVerified }
func (gp *GlobalProduct) IsActive() bool                           { return gp.isActive }
func (gp *GlobalProduct) BusinessType() *string                    { return gp.businessType }
func (gp *GlobalProduct) Tags() []string                           { return gp.tags }
func (gp *GlobalProduct) Metadata() map[string]interface{}         { return gp.metadata }
func (gp *GlobalProduct) CreatedAt() time.Time                     { return gp.createdAt }
func (gp *GlobalProduct) UpdatedAt() time.Time                     { return gp.updatedAt }
func (gp *GlobalProduct) LastScrapedAt() *time.Time                { return gp.lastScrapedAt }

// UpdateFromScraping actualiza el producto con datos obtenidos de scraping
func (gp *GlobalProduct) UpdateFromScraping(
	name string,
	description *string,
	brand *string,
	category *string,
	price *float64,
	imageURL *string,
	source *value_object.ProductSource,
) error {
	if name == "" {
		return errors.New("el nombre del producto es obligatorio")
	}

	if source == nil {
		return errors.New("la fuente del producto es obligatoria")
	}

	// Actualizar campos
	gp.name = strings.TrimSpace(name)
	gp.description = description
	gp.brand = brand
	gp.category = category
	gp.price = price
	gp.imageURL = imageURL
	gp.source = source
	gp.updatedAt = time.Now()
	gp.lastScrapedAt = source.ScrapedAt()

	// Recalcular quality score
	gp.qualityScore = calculateInitialQualityScore(name, description, brand, category, price, imageURL, gp.ean)

	return nil
}

// Verify marca el producto como verificado
func (gp *GlobalProduct) Verify() {
	gp.isVerified = true
	gp.updatedAt = time.Now()
}

// Unverify marca el producto como no verificado
func (gp *GlobalProduct) Unverify() {
	gp.isVerified = false
	gp.updatedAt = time.Now()
}

// Activate activa el producto
func (gp *GlobalProduct) Activate() {
	gp.isActive = true
	gp.updatedAt = time.Now()
}

// Deactivate desactiva el producto
func (gp *GlobalProduct) Deactivate() {
	gp.isActive = false
	gp.updatedAt = time.Now()
}

// SetBusinessType establece el tipo de negocio recomendado
func (gp *GlobalProduct) SetBusinessType(businessType string) {
	gp.businessType = &businessType
	gp.updatedAt = time.Now()
}

// AddTag agrega una etiqueta al producto
func (gp *GlobalProduct) AddTag(tag string) {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return
	}

	// Verificar que no exista ya
	for _, existingTag := range gp.tags {
		if strings.EqualFold(existingTag, tag) {
			return
		}
	}

	gp.tags = append(gp.tags, tag)
	gp.updatedAt = time.Now()
}

// RemoveTag elimina una etiqueta del producto
func (gp *GlobalProduct) RemoveTag(tag string) {
	for i, existingTag := range gp.tags {
		if strings.EqualFold(existingTag, tag) {
			gp.tags = append(gp.tags[:i], gp.tags[i+1:]...)
			gp.updatedAt = time.Now()
			break
		}
	}
}

// AddImage agrega una URL de imagen adicional
func (gp *GlobalProduct) AddImage(imageURL string) {
	if imageURL == "" {
		return
	}

	// Verificar que no exista ya
	for _, existingURL := range gp.imageURLs {
		if existingURL == imageURL {
			return
		}
	}

	gp.imageURLs = append(gp.imageURLs, imageURL)
	gp.updatedAt = time.Now()

	// Recalcular quality score
	gp.qualityScore = calculateInitialQualityScore(gp.name, gp.description, gp.brand, gp.category, gp.price, gp.imageURL, gp.ean)
}

// SetImageURLs reemplaza todas las URLs de imágenes adicionales
func (gp *GlobalProduct) SetImageURLs(imageURLs []string) {
	// Filtrar URLs vacías y duplicadas
	var validURLs []string
	seen := make(map[string]bool)

	for _, url := range imageURLs {
		if url != "" && !seen[url] {
			validURLs = append(validURLs, url)
			seen[url] = true
		}
	}

	gp.imageURLs = validURLs
	gp.updatedAt = time.Now()

	// Recalcular quality score
	gp.qualityScore = calculateInitialQualityScore(gp.name, gp.description, gp.brand, gp.category, gp.price, gp.imageURL, gp.ean)
}

// SetMetadata establece un valor en el metadata
func (gp *GlobalProduct) SetMetadata(key string, value interface{}) {
	if gp.metadata == nil {
		gp.metadata = make(map[string]interface{})
	}
	gp.metadata[key] = value
	gp.updatedAt = time.Now()
}

// GetMetadata obtiene un valor del metadata
func (gp *GlobalProduct) GetMetadata(key string) interface{} {
	if gp.metadata == nil {
		return nil
	}
	return gp.metadata[key]
}

// IsArgentineProduct verifica si es un producto argentino usando el EAN
func (gp *GlobalProduct) IsArgentineProduct() bool {
	return gp.ean.IsArgentineProduct()
}

// HasCompleteData verifica si el producto tiene datos completos
func (gp *GlobalProduct) HasCompleteData() bool {
	return gp.qualityScore.IsHighQuality()
}

// NeedsUpdate verifica si el producto necesita actualización
func (gp *GlobalProduct) NeedsUpdate(maxAge time.Duration) bool {
	return gp.source.NeedsUpdate(maxAge)
}

// calculateInitialQualityScore calcula el quality score inicial
func calculateInitialQualityScore(
	name string,
	description *string,
	brand *string,
	category *string,
	price *float64,
	imageURL *string,
	ean *value_object.EAN13,
) *value_object.QualityScore {
	metrics := value_object.QualityMetrics{
		HasName:           name != "",
		HasDescription:    description != nil && *description != "",
		HasImage:          imageURL != nil && *imageURL != "",
		HasPrice:          price != nil && *price > 0,
		HasBrand:          brand != nil && *brand != "",
		HasCategory:       category != nil && *category != "",
		HasEAN:            ean != nil,
		HasStock:          false, // No manejamos stock en catálogo global
		ImageCount:        1,     // Por ahora solo imagen principal
		DescriptionLength: 0,
	}

	if description != nil {
		metrics.DescriptionLength = len(*description)
	}

	qualityScore, _ := value_object.NewQualityScoreFromMetrics(metrics)
	return qualityScore
}
