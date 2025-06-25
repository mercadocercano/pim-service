package usecase

import (
	"pim/src/product/global_catalog/domain/entity"
	"pim/src/product/global_catalog/domain/exception"
	"pim/src/product/global_catalog/domain/port"
	"pim/src/product/global_catalog/domain/value_object"
)

// CreateGlobalProductRequest contiene los datos para crear un producto global
type CreateGlobalProductRequest struct {
	EAN          string                 `json:"ean" validate:"required,len=13"`
	Name         string                 `json:"name" validate:"required,min=3,max=500"`
	Description  *string                `json:"description,omitempty"`
	Brand        *string                `json:"brand,omitempty"`
	Category     *string                `json:"category,omitempty"`
	Price        *float64               `json:"price,omitempty"`
	ImageURL     *string                `json:"image_url,omitempty"`
	Source       string                 `json:"source" validate:"required"`
	SourceURL    *string                `json:"source_url,omitempty"`
	Reliability  *float64               `json:"reliability,omitempty"`
	BusinessType *string                `json:"business_type,omitempty"`
	Tags         []string               `json:"tags,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// CreateGlobalProductResponse contiene la respuesta del caso de uso
type CreateGlobalProductResponse struct {
	ID           string                 `json:"id"`
	EAN          string                 `json:"ean"`
	Name         string                 `json:"name"`
	Description  *string                `json:"description"`
	Brand        *string                `json:"brand"`
	Category     *string                `json:"category"`
	Price        *float64               `json:"price"`
	ImageURL     *string                `json:"image_url"`
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
}

// CreateGlobalProduct implementa el caso de uso de crear un producto global
type CreateGlobalProduct struct {
	globalProductRepository port.GlobalProductRepository
}

// NewCreateGlobalProduct crea una nueva instancia del caso de uso
func NewCreateGlobalProduct(globalProductRepository port.GlobalProductRepository) *CreateGlobalProduct {
	return &CreateGlobalProduct{
		globalProductRepository: globalProductRepository,
	}
}

// Execute ejecuta el caso de uso
func (uc *CreateGlobalProduct) Execute(request CreateGlobalProductRequest) (*CreateGlobalProductResponse, error) {
	// Validar y crear EAN
	ean, err := value_object.NewEAN13(request.EAN)
	if err != nil {
		return nil, exception.NewValidationError("EAN inválido", err)
	}

	// Verificar que no exista un producto con el mismo EAN
	existingProduct, err := uc.globalProductRepository.FindByEAN(ean.Value())
	if err != nil {
		return nil, exception.NewInternalError("Error al verificar EAN existente", err)
	}
	if existingProduct != nil {
		return nil, exception.NewConflictError("Ya existe un producto con el EAN " + request.EAN)
	}

	// Crear ProductSource
	reliability := 0.5 // Valor por defecto
	if request.Reliability != nil {
		reliability = *request.Reliability
	}

	var productSource *value_object.ProductSource
	if request.SourceURL != nil {
		productSource, err = value_object.NewScrapingSource(request.Source, *request.SourceURL, reliability)
	} else {
		productSource, err = value_object.NewProductSource(request.Source, nil, nil, reliability)
	}
	if err != nil {
		return nil, exception.NewValidationError("Fuente del producto inválida", err)
	}

	// Crear el producto global
	globalProduct, err := entity.NewGlobalProduct(
		ean,
		request.Name,
		request.Description,
		request.Brand,
		request.Category,
		request.Price,
		request.ImageURL,
		productSource,
	)
	if err != nil {
		return nil, exception.NewValidationError("Error al crear producto", err)
	}

	// Agregar business type si se proporciona
	if request.BusinessType != nil {
		globalProduct.SetBusinessType(*request.BusinessType)
	}

	// Agregar tags si se proporcionan
	if request.Tags != nil {
		for _, tag := range request.Tags {
			globalProduct.AddTag(tag)
		}
	}

	// Agregar metadata si se proporciona
	if request.Metadata != nil {
		for key, value := range request.Metadata {
			globalProduct.SetMetadata(key, value)
		}
	}

	// Guardar en el repositorio
	savedProduct, err := uc.globalProductRepository.Save(globalProduct)
	if err != nil {
		return nil, exception.NewInternalError("Error al guardar producto", err)
	}

	// Mapear a respuesta
	response := &CreateGlobalProductResponse{
		ID:           savedProduct.IDString(),
		EAN:          savedProduct.EAN().Value(),
		Name:         savedProduct.Name(),
		Description:  savedProduct.Description(),
		Brand:        savedProduct.Brand(),
		Category:     savedProduct.Category(),
		Price:        savedProduct.Price(),
		ImageURL:     savedProduct.ImageURL(),
		QualityScore: savedProduct.QualityScore().Value(),
		Source:       savedProduct.Source().Source(),
		SourceURL:    savedProduct.Source().SourceURL(),
		IsVerified:   savedProduct.IsVerified(),
		IsActive:     savedProduct.IsActive(),
		BusinessType: savedProduct.BusinessType(),
		Tags:         savedProduct.Tags(),
		Metadata:     savedProduct.Metadata(),
		CreatedAt:    savedProduct.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    savedProduct.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}

	return response, nil
}
