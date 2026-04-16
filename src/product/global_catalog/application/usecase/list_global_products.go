package usecase

import (
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
)

// ListGlobalProductsRequest contiene los parámetros para listar productos
type ListGlobalProductsRequest struct {
	// Paginación
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`

	// Filtros específicos
	BusinessType    *string `json:"business_type,omitempty"`
	Source          *string `json:"source,omitempty"`
	MinQuality      *int    `json:"min_quality,omitempty"`
	MaxQuality      *int    `json:"max_quality,omitempty"`
	OnlyActive      bool    `json:"only_active,omitempty"`
	OnlyVerified    bool    `json:"only_verified,omitempty"`
	OnlyArgentine   bool    `json:"only_argentine,omitempty"`
	OnlyHighQuality bool    `json:"only_high_quality,omitempty"`

	// Búsquedas de texto
	SearchName     *string  `json:"search_name,omitempty"`
	SearchBrand    *string  `json:"search_brand,omitempty"`
	SearchCategory *string  `json:"search_category,omitempty"`
	SearchTags     []string `json:"search_tags,omitempty"`
}

// GlobalProductSummary contiene un resumen del producto para listados
type GlobalProductSummary struct {
	ID                 string   `json:"id"`
	EAN                string   `json:"ean"`
	Name               string   `json:"name"`
	Brand              *string  `json:"brand"`
	Category           *string  `json:"category"`
	Price              *float64 `json:"price"`
	ImageURL           *string  `json:"image_url"`
	QualityScore       int      `json:"quality_score"`
	QualityLevel       string   `json:"quality_level"`
	Source             string   `json:"source"`
	SourceDisplayName  string   `json:"source_display_name"`
	IsVerified         bool     `json:"is_verified"`
	IsActive           bool     `json:"is_active"`
	BusinessType       *string  `json:"business_type"`
	IsArgentineProduct bool     `json:"is_argentine_product"`
	Tags               []string `json:"tags"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
}

// ListGlobalProductsResponse contiene la respuesta paginada
type ListGlobalProductsResponse struct {
	Products   []*GlobalProductSummary `json:"products"`
	Pagination *PaginationInfo         `json:"pagination"`
	Summary    *ListSummary            `json:"summary"`
}

// PaginationInfo contiene información de paginación
type PaginationInfo struct {
	Offset     int  `json:"offset"`
	Limit      int  `json:"limit"`
	Total      int  `json:"total"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
	TotalPages int  `json:"total_pages"`
}

// ListSummary contiene estadísticas del resultado
type ListSummary struct {
	TotalProducts       int `json:"total_products"`
	ArgentineProducts   int `json:"argentine_products"`
	VerifiedProducts    int `json:"verified_products"`
	HighQualityProducts int `json:"high_quality_products"`
	AverageQuality      int `json:"average_quality"`
}

// ListGlobalProducts implementa el caso de uso de listar productos globales
type ListGlobalProducts struct {
	globalProductRepository port.GlobalProductRepository
}

// NewListGlobalProducts crea una nueva instancia del caso de uso
func NewListGlobalProducts(globalProductRepository port.GlobalProductRepository) *ListGlobalProducts {
	return &ListGlobalProducts{
		globalProductRepository: globalProductRepository,
	}
}

// Execute ejecuta el caso de uso
func (uc *ListGlobalProducts) Execute(request ListGlobalProductsRequest) (*ListGlobalProductsResponse, error) {
	// Validar y establecer valores por defecto
	offset := request.Offset
	if offset < 0 {
		offset = 0
	}

	limit := request.Limit
	if limit <= 0 || limit > 100 {
		limit = 20 // Límite por defecto
	}

	var products []*entity.GlobalProduct
	var err error

	// Aplicar filtros específicos
	if request.BusinessType != nil {
		products, err = uc.globalProductRepository.FindByBusinessType(*request.BusinessType, limit)
	} else if request.Source != nil {
		products, err = uc.globalProductRepository.FindBySource(*request.Source, limit)
	} else if request.MinQuality != nil && request.MaxQuality != nil {
		products, err = uc.globalProductRepository.FindByQualityScoreRange(*request.MinQuality, *request.MaxQuality, limit)
	} else if request.SearchName != nil {
		products, err = uc.globalProductRepository.SearchByName(*request.SearchName, limit)
	} else if request.SearchBrand != nil {
		products, err = uc.globalProductRepository.SearchByBrand(*request.SearchBrand, limit)
	} else if request.SearchCategory != nil {
		products, err = uc.globalProductRepository.SearchByCategory(*request.SearchCategory, limit)
	} else if request.SearchTags != nil && len(request.SearchTags) > 0 {
		products, err = uc.globalProductRepository.SearchByTags(request.SearchTags, limit)
	} else if request.OnlyArgentine {
		products, err = uc.globalProductRepository.FindArgentineProducts(offset, limit)
	} else if request.OnlyHighQuality {
		products, err = uc.globalProductRepository.FindHighQualityProducts(offset, limit)
	} else if request.OnlyVerified {
		products, err = uc.globalProductRepository.FindVerified(offset, limit)
	} else if request.OnlyActive {
		products, err = uc.globalProductRepository.FindActive(offset, limit)
	} else {
		// Listado general
		products, err = uc.globalProductRepository.FindAll(offset, limit)
	}

	if err != nil {
		return nil, exception.NewInternalError("Error al listar productos", err)
	}

	// Obtener estadísticas para el summary
	totalProducts, _ := uc.globalProductRepository.CountTotal()
	argentineCount, _ := uc.globalProductRepository.CountArgentineProducts()
	highQualityCount, _ := uc.globalProductRepository.CountByQualityScore(70)

	// Calcular estadísticas de productos obtenidos
	verifiedCount := 0
	totalQualitySum := 0

	for _, product := range products {
		if product.IsVerified() {
			verifiedCount++
		}
		totalQualitySum += product.QualityScore().Value()
	}

	averageQuality := 0
	if len(products) > 0 {
		averageQuality = totalQualitySum / len(products)
	}

	// Mapear productos a summary
	productSummaries := make([]*GlobalProductSummary, len(products))
	for i, product := range products {
		productSummaries[i] = &GlobalProductSummary{
			ID:                 product.IDString(),
			EAN:                product.EANString(),
			Name:               product.Name(),
			Brand:              product.Brand(),
			Category:           product.Category(),
			Price:              product.Price(),
			ImageURL:           product.ImageURL(),
			QualityScore:       product.QualityScore().Value(),
			QualityLevel:       product.QualityScore().Level(),
			Source:             product.Source().Source(),
			SourceDisplayName:  product.Source().GetSourceDisplayName(),
			IsVerified:         product.IsVerified(),
			IsActive:           product.IsActive(),
			BusinessType:       product.BusinessType(),
			IsArgentineProduct: product.IsArgentineProduct(),
			Tags:               product.Tags(),
			CreatedAt:          product.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:          product.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	// Calcular información de paginación
	totalPages := (totalProducts + limit - 1) / limit // Redondeo hacia arriba
	hasNext := offset+limit < totalProducts
	hasPrev := offset > 0

	pagination := &PaginationInfo{
		Offset:     offset,
		Limit:      limit,
		Total:      totalProducts,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
		TotalPages: totalPages,
	}

	summary := &ListSummary{
		TotalProducts:       totalProducts,
		ArgentineProducts:   argentineCount,
		VerifiedProducts:    verifiedCount,
		HighQualityProducts: highQualityCount,
		AverageQuality:      averageQuality,
	}

	response := &ListGlobalProductsResponse{
		Products:   productSummaries,
		Pagination: pagination,
		Summary:    summary,
	}

	return response, nil
}
