package usecase

import (
	"pim/src/product/global_catalog/domain/entity"
	"pim/src/product/global_catalog/domain/exception"
	"pim/src/product/global_catalog/domain/port"
	"pim/src/product/global_catalog/domain/value_object"
)

// SearchByEANRequest contiene los datos para buscar por EAN
type SearchByEANRequest struct {
	EAN        string `json:"ean" validate:"required,len=13"`
	OnlyActive bool   `json:"only_active,omitempty"`
}

// SearchByEANResponse contiene la respuesta del caso de uso
type SearchByEANResponse struct {
	ID           string                 `json:"id"`
	EAN          string                 `json:"ean"`
	Name         string                 `json:"name"`
	Description  *string                `json:"description"`
	Brand        *string                `json:"brand"`
	Category     *string                `json:"category"`
	Price        *float64               `json:"price"`
	ImageURL     *string                `json:"image_url"`
	ImageURLs    []string               `json:"image_urls"`
	QualityScore int                    `json:"quality_score"`
	Source       string                 `json:"source"`
	SourceURL    *string                `json:"source_url"`
	IsVerified   bool                   `json:"is_verified"`
	IsActive     bool                   `json:"is_active"`
	BusinessType *string                `json:"business_type"`
	Tags         []string               `json:"tags"`
	Metadata     map[string]interface{} `json:"metadata"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	// Campos adicionales para respuesta de búsqueda
	IsArgentineProduct bool   `json:"is_argentine_product"`
	QualityLevel       string `json:"quality_level"`
	SourceDisplayName  string `json:"source_display_name"`
}

// SearchByEAN implementa el caso de uso de buscar un producto por EAN
type SearchByEAN struct {
	globalProductRepository port.GlobalProductRepository
}

// NewSearchByEAN crea una nueva instancia del caso de uso
func NewSearchByEAN(globalProductRepository port.GlobalProductRepository) *SearchByEAN {
	return &SearchByEAN{
		globalProductRepository: globalProductRepository,
	}
}

// Execute ejecuta el caso de uso
func (uc *SearchByEAN) Execute(request SearchByEANRequest) (*SearchByEANResponse, error) {
	// Validar EAN
	ean, err := value_object.NewEAN13(request.EAN)
	if err != nil {
		return nil, exception.NewValidationError("EAN inválido", err)
	}

	// Buscar el producto
	var globalProduct *entity.GlobalProduct
	if request.OnlyActive {
		globalProduct, err = uc.globalProductRepository.FindActiveByEAN(ean.Value())
	} else {
		globalProduct, err = uc.globalProductRepository.FindByEAN(ean.Value())
	}

	if err != nil {
		return nil, exception.NewInternalError("Error al buscar producto por EAN", err)
	}

	if globalProduct == nil {
		return nil, exception.NewGlobalProductNotFoundByEAN(request.EAN)
	}

	// Mapear a respuesta con información adicional
	response := &SearchByEANResponse{
		ID:                 globalProduct.IDString(),
		EAN:                globalProduct.EAN().Value(),
		Name:               globalProduct.Name(),
		Description:        globalProduct.Description(),
		Brand:              globalProduct.Brand(),
		Category:           globalProduct.Category(),
		Price:              globalProduct.Price(),
		ImageURL:           globalProduct.ImageURL(),
		ImageURLs:          globalProduct.ImageURLs(),
		QualityScore:       globalProduct.QualityScore().Value(),
		Source:             globalProduct.Source().Source(),
		SourceURL:          globalProduct.Source().SourceURL(),
		IsVerified:         globalProduct.IsVerified(),
		IsActive:           globalProduct.IsActive(),
		BusinessType:       globalProduct.BusinessType(),
		Tags:               globalProduct.Tags(),
		Metadata:           globalProduct.Metadata(),
		CreatedAt:          globalProduct.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:          globalProduct.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
		IsArgentineProduct: globalProduct.IsArgentineProduct(),
		QualityLevel:       globalProduct.QualityScore().Level(),
		SourceDisplayName:  globalProduct.Source().GetSourceDisplayName(),
	}

	return response, nil
}
