package usecase

import (
	"context"
	"saas-mt-pim-service/src/product/global_catalog/domain/entity"
	"saas-mt-pim-service/src/product/global_catalog/domain/exception"
	"saas-mt-pim-service/src/product/global_catalog/domain/port"
	"saas-mt-pim-service/src/product/global_catalog/domain/value_object"
)

// UpdateGlobalProductByIDRequest contiene los datos para actualizar un producto global
type UpdateGlobalProductByIDRequest struct {
	ID           string                 `json:"id" validate:"required"`
	EAN          *string                `json:"ean,omitempty"`
	Name         *string                `json:"name,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Brand        *string                `json:"brand,omitempty"`
	Category     *string                `json:"category,omitempty"`
	Price        *float64               `json:"price,omitempty"`
	ImageURL     *string                `json:"image_url,omitempty"`
	ImageURLs    []string               `json:"image_urls,omitempty"`
	Source       *string                `json:"source,omitempty"`
	SourceURL    *string                `json:"source_url,omitempty"`
	Reliability  *float64               `json:"reliability,omitempty"`
	BusinessType *string                `json:"business_type,omitempty"`
	Tags         []string               `json:"tags,omitempty"`
	IsVerified   *bool                  `json:"is_verified,omitempty"`
	IsActive     *bool                  `json:"is_active,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// UpdateGlobalProductByIDResponse contiene la respuesta del caso de uso
type UpdateGlobalProductByIDResponse struct {
	ID           string                 `json:"id"`
	EAN          *string                `json:"ean,omitempty"`
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
}

// UpdateGlobalProductByID implementa el caso de uso de actualizar un producto global por ID
type UpdateGlobalProductByID struct {
	globalProductRepository port.GlobalProductRepository
}

// NewUpdateGlobalProductByID crea una nueva instancia del caso de uso
func NewUpdateGlobalProductByID(globalProductRepository port.GlobalProductRepository) *UpdateGlobalProductByID {
	return &UpdateGlobalProductByID{
		globalProductRepository: globalProductRepository,
	}
}

// Execute ejecuta el caso de uso
func (uc *UpdateGlobalProductByID) Execute(ctx context.Context, request UpdateGlobalProductByIDRequest) (*UpdateGlobalProductByIDResponse, error) {
	// Validar ID
	if request.ID == "" {
		return nil, exception.NewValidationError("ID del producto es obligatorio", nil)
	}

	// Buscar el producto existente
	existingProduct, err := uc.globalProductRepository.FindByID(request.ID)
	if err != nil {
		return nil, exception.NewInternalError("Error al buscar el producto", err)
	}
	if existingProduct == nil {
		return nil, exception.NewGlobalProductNotFoundByID(request.ID)
	}

	// Verificar si se está cambiando el EAN y si ya existe otro producto con ese EAN
	currentEAN := ""
	if existingProduct.EAN() != nil {
		currentEAN = existingProduct.EAN().Value()
	}
	if request.EAN != nil && *request.EAN != currentEAN {
		// Validar nuevo EAN
		_, err := value_object.NewEAN13(*request.EAN)
		if err != nil {
			return nil, exception.NewValidationError("EAN inválido", err)
		}

		// Verificar que no exista otro producto con el mismo EAN
		conflictProduct, err := uc.globalProductRepository.FindByEAN(*request.EAN)
		if err != nil {
			return nil, exception.NewInternalError("Error al verificar EAN existente", err)
		}
		if conflictProduct != nil && conflictProduct.IDString() != request.ID {
			return nil, exception.NewConflictError("Ya existe otro producto con el EAN " + *request.EAN)
		}
	}

	// Preparar datos para actualización usando UpdateFromScraping
	name := existingProduct.Name()
	if request.Name != nil {
		if len(*request.Name) < 3 || len(*request.Name) > 500 {
			return nil, exception.NewValidationError("El nombre debe tener entre 3 y 500 caracteres", nil)
		}
		name = *request.Name
	}

	description := existingProduct.Description()
	if request.Description != nil {
		description = request.Description
	}

	brand := existingProduct.Brand()
	if request.Brand != nil {
		brand = request.Brand
	}

	category := existingProduct.Category()
	if request.Category != nil {
		category = request.Category
	}

	price := existingProduct.Price()
	if request.Price != nil {
		price = request.Price
	}

	imageURL := existingProduct.ImageURL()
	if request.ImageURL != nil {
		imageURL = request.ImageURL
	}

	// Preparar fuente actualizada
	currentSource := existingProduct.Source()
	source := currentSource.Source()
	if request.Source != nil {
		source = *request.Source
	}

	var sourceURL *string
	if request.SourceURL != nil {
		sourceURL = request.SourceURL
	} else {
		sourceURL = currentSource.SourceURL()
	}

	reliability := currentSource.Reliability()
	if request.Reliability != nil {
		reliability = *request.Reliability
	}

	var newProductSource *value_object.ProductSource
	if sourceURL != nil {
		newProductSource, err = value_object.NewScrapingSource(source, *sourceURL, reliability)
	} else {
		newProductSource, err = value_object.NewProductSource(source, nil, nil, reliability)
	}
	if err != nil {
		return nil, exception.NewValidationError("Fuente del producto inválida", err)
	}

	// Actualizar usando UpdateFromScraping
	err = existingProduct.UpdateFromScraping(name, description, brand, category, price, imageURL, newProductSource)
	if err != nil {
		return nil, exception.NewValidationError("Error al actualizar producto", err)
	}

	// Si se cambió el EAN, necesitamos crear un nuevo producto con el EAN actualizado
	// ya que UpdateFromScraping no permite cambiar el EAN
	if request.EAN != nil && *request.EAN != currentEAN {
		newEAN, _ := value_object.NewEAN13(*request.EAN)

		// Crear nuevo producto con el EAN actualizado
		newProduct, err := entity.NewGlobalProductFromRepository(
			existingProduct.ID(),
			newEAN,
			name,
			description,
			brand,
			category,
			price,
			imageURL,
			existingProduct.ImageURLs(),
			newProductSource,
			existingProduct.QualityScore(),
			existingProduct.IsVerified(),
			existingProduct.IsActive(),
			existingProduct.BusinessType(),
			existingProduct.Tags(),
			existingProduct.Metadata(),
			existingProduct.CreatedAt(),
			existingProduct.UpdatedAt(),
			existingProduct.LastScrapedAt(),
		)
		if err != nil {
			return nil, exception.NewInternalError("Error al crear producto con nuevo EAN", err)
		}
		existingProduct = newProduct
	}

	// Actualizar business type si se proporciona
	if request.BusinessType != nil {
		existingProduct.SetBusinessType(*request.BusinessType)
	}

	// Actualizar tags si se proporcionan
	if request.Tags != nil {
		// Remover todos los tags existentes
		for _, tag := range existingProduct.Tags() {
			existingProduct.RemoveTag(tag)
		}
		// Agregar los nuevos tags
		for _, tag := range request.Tags {
			existingProduct.AddTag(tag)
		}
	}

	// Actualizar estados si se proporcionan
	if request.IsVerified != nil {
		if *request.IsVerified {
			existingProduct.Verify()
		} else {
			existingProduct.Unverify()
		}
	}

	if request.IsActive != nil {
		if *request.IsActive {
			existingProduct.Activate()
		} else {
			existingProduct.Deactivate()
		}
	}

	// Actualizar metadata si se proporciona
	if request.Metadata != nil {
		for key, value := range request.Metadata {
			existingProduct.SetMetadata(key, value)
		}
	}

	// Actualizar image_urls si se proporcionan
	if request.ImageURLs != nil {
		existingProduct.SetImageURLs(request.ImageURLs)
	}

	// Guardar cambios en el repositorio
	updatedProduct, err := uc.globalProductRepository.Update(existingProduct)
	if err != nil {
		return nil, exception.NewInternalError("Error al actualizar producto", err)
	}

	// Mapear a respuesta
	response := &UpdateGlobalProductByIDResponse{
		ID:           updatedProduct.IDString(),
		Name:         updatedProduct.Name(),
		Description:  updatedProduct.Description(),
		Brand:        updatedProduct.Brand(),
		Category:     updatedProduct.Category(),
		Price:        updatedProduct.Price(),
		ImageURL:     updatedProduct.ImageURL(),
		ImageURLs:    updatedProduct.ImageURLs(),
		QualityScore: updatedProduct.QualityScore().Value(),
		Source:       updatedProduct.Source().Source(),
		SourceURL:    updatedProduct.Source().SourceURL(),
		IsVerified:   updatedProduct.IsVerified(),
		IsActive:     updatedProduct.IsActive(),
		BusinessType: updatedProduct.BusinessType(),
		Tags:         updatedProduct.Tags(),
		Metadata:     updatedProduct.Metadata(),
		CreatedAt:    updatedProduct.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    updatedProduct.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}

	// Agregar EAN solo si existe
	if updatedProduct.EAN() != nil {
		eanValue := updatedProduct.EAN().Value()
		response.EAN = &eanValue
	}

	return response, nil
}
