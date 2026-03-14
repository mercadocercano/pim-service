package controller

import (
	"net/http"
	"strconv"

	"saas-mt-pim-service/src/brand/application/request"
	"saas-mt-pim-service/src/brand/application/response"
	"saas-mt-pim-service/src/brand/application/usecase"
	"saas-mt-pim-service/src/brand/domain/entity"
	"saas-mt-pim-service/src/brand/domain/port"
	cr "github.com/mercadocercano/criteria"

	"github.com/gin-gonic/gin"
)

// MarketplaceBrandHandler maneja las peticiones HTTP para marcas marketplace
type MarketplaceBrandHandler struct {
	repository    port.MarketplacebrandRepository
	updateUseCase *usecase.UpdateMarketplaceBrandUseCase
}

// NewMarketplaceBrandHandler crea una nueva instancia del manejador de marcas marketplace
func NewMarketplaceBrandHandler(repository port.MarketplacebrandRepository) *MarketplaceBrandHandler {
	return &MarketplaceBrandHandler{
		repository:    repository,
		updateUseCase: usecase.NewUpdateMarketplaceBrandUseCase(repository),
	}
}

// RegisterRoutes registra las rutas del API para marcas marketplace
func (h *MarketplaceBrandHandler) RegisterRoutes(router *gin.RouterGroup) {
	marketplace := router.Group("/marketplace-brands")
	{
		// Rutas para administradores (gestión de marcas globales)
		marketplace.GET("", h.GetAllMarketplaceBrands)
		marketplace.POST("", h.CreateMarketplaceBrand)
		marketplace.GET("/:id", h.GetMarketplaceBrandByID)
		marketplace.PUT("/:id", h.UpdateMarketplaceBrand)
		marketplace.DELETE("/:id", h.DeleteMarketplaceBrand)
		marketplace.PUT("/:id/verify", h.VerifyMarketplaceBrand)
		marketplace.PUT("/:id/unverify", h.UnverifyMarketplaceBrand)
	}
}

// GetAllMarketplaceBrands maneja la solicitud para obtener todas las marcas marketplace
func (h *MarketplaceBrandHandler) GetAllMarketplaceBrands(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden ver marcas marketplace"})
		return
	}

	// Construir criterios de filtrado desde query parameters
	criteriaBuilder := cr.NewCriteriaBuilder().FromURLValues(c.Request.URL.Query())

	// Configurar paginación personalizada para marcas (sin límite de 100)
	page := 1
	pageSize := 20

	if pageStr := c.Query("page"); pageStr != "" {
		if pageInt, err := strconv.Atoi(pageStr); err == nil && pageInt > 0 {
			page = pageInt
		}
	}

	if pageSizeStr := c.Query("page_size"); pageSizeStr != "" {
		if pageSizeInt, err := strconv.Atoi(pageSizeStr); err == nil && pageSizeInt > 0 {
			pageSize = pageSizeInt
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if limitInt, err := strconv.Atoi(limitStr); err == nil && limitInt > 0 {
			pageSize = limitInt
		}
	}

	// Limitar solo para seguridad (1000 marcas máximo)
	if pageSize > 1000 {
		pageSize = 1000
	}

	// Establecer paginación manualmente
	criteriaBuilder.SetPagination(page, pageSize)

	// Filtros
	if search := c.Query("search"); search != "" {
		criteriaBuilder.AddLikeFilter("name", search)
	}

	if isVerifiedStr := c.Query("is_verified"); isVerifiedStr != "" {
		if isVerified, err := strconv.ParseBool(isVerifiedStr); err == nil && isVerified {
			criteriaBuilder.AddEqualFilter("verification_status", "verified")
		}
	}

	// Ordenamiento por defecto solo si no se especificó uno
	if c.Query("sort_by") == "" {
		criteriaBuilder.SetOrder("name", "ASC") // Cambiar a orden alfabético por defecto
	}

	// Construir criteria final
	builtCriteria := criteriaBuilder.Build()

	// Usar el repository para obtener datos
	brands, err := h.repository.SearchByCriteria(c.Request.Context(), builtCriteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Contar total
	total, err := h.repository.CountByCriteria(c.Request.Context(), builtCriteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convertir a response
	brandResponses := make([]*response.MarketplaceBrandResponse, len(brands))
	for i, brand := range brands {
		brandResponses[i] = response.NewMarketplaceBrandResponse(brand)
	}

	// Obtener valores de paginación del criteria
	pagination := builtCriteria.Pagination
	offset := pagination.Offset
	limit := pagination.Limit

	// Devolver respuesta con información de paginación
	response := gin.H{
		"brands": brandResponses,
		"pagination": gin.H{
			"offset":      offset,
			"limit":       limit,
			"total":       total,
			"has_next":    offset+limit < total,
			"has_prev":    offset > 0,
			"total_pages": (total + limit - 1) / limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

// CreateMarketplaceBrand maneja la solicitud para crear una nueva marca marketplace
func (h *MarketplaceBrandHandler) CreateMarketplaceBrand(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden crear marcas marketplace"})
		return
	}

	// Estructura para recibir los datos de creación
	var createData struct {
		Name        string   `json:"name" binding:"required"`
		Description string   `json:"description"`
		LogoURL     string   `json:"logo_url"`
		Website     string   `json:"website"`
		Aliases     []string `json:"aliases"`
		CategoryTags []string `json:"category_tags"`
	}

	if err := c.ShouldBindJSON(&createData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	// Validar que el nombre no esté vacío
	if createData.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre de la marca es requerido"})
		return
	}

	// Crear la entidad Marketplacebrand usando el constructor
	brand, err := entity.NewMarketplacebrand(createData.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al crear la marca: " + err.Error()})
		return
	}

	// Actualizar campos adicionales si se proporcionaron
	if createData.Description != "" {
		brand.Description = createData.Description
	}
	// LogoURL debe ser NULL si está vacío (por restricción de BD)
	if createData.LogoURL != "" {
		brand.LogoURL = createData.LogoURL
	} else {
		brand.LogoURL = "" // Se convertirá a NULL en el repositorio si está vacío
	}
	// Website debe ser NULL si está vacío (por restricción de BD)
	if createData.Website != "" {
		brand.Website = createData.Website
	} else {
		brand.Website = "" // Se convertirá a NULL en el repositorio si está vacío
	}
	if len(createData.Aliases) > 0 {
		brand.Aliases = createData.Aliases
	}
	if len(createData.CategoryTags) > 0 {
		brand.CategoryTags = createData.CategoryTags
	}

	// Guardar en el repositorio
	if err := h.repository.Create(c.Request.Context(), brand); err != nil {
		// Verificar si es un error de duplicado
		if err.Error() == "marca ya existe" || err.Error() == "duplicate key" {
			c.JSON(http.StatusConflict, gin.H{"error": "Una marca con ese nombre ya existe"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la marca: " + err.Error()})
		return
	}

	// Convertir a response y retornar
	brandResponse := response.NewMarketplaceBrandResponse(brand)
	c.JSON(http.StatusCreated, brandResponse)
}

// GetMarketplaceBrandByID maneja la solicitud para obtener una marca marketplace por ID
func (h *MarketplaceBrandHandler) GetMarketplaceBrandByID(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden ver marcas marketplace"})
		return
	}

	// Obtener ID de la marca
	brandID := c.Param("id")
	if brandID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de marca es requerido"})
		return
	}

	// Buscar la marca por ID
	brand, err := h.repository.FindByID(c.Request.Context(), brandID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error buscando la marca: " + err.Error()})
		return
	}

	if brand == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Marca no encontrada"})
		return
	}

	// Convertir a response
	brandResponse := response.NewMarketplaceBrandResponse(brand)

	c.JSON(http.StatusOK, brandResponse)
}

// UpdateMarketplaceBrand maneja la solicitud para actualizar una marca marketplace
func (h *MarketplaceBrandHandler) UpdateMarketplaceBrand(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden actualizar marcas marketplace"})
		return
	}

	// Obtener ID de la marca
	brandID := c.Param("id")
	if brandID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de marca es requerido"})
		return
	}

	// Estructura para recibir los datos de actualización
	var updateData struct {
		Name         string   `json:"name"`
		Description  string   `json:"description"`
		LogoURL      string   `json:"logo_url"`
		Website      string   `json:"website"`
		Aliases      []string `json:"aliases"`
		CategoryTags []string `json:"category_tags"`
		Sources      []string `json:"sources"`
		QualityScore float64  `json:"quality_score"`
		IsActive     bool     `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	// Crear request para el use case
	updateRequest := &request.UpdateMarketplaceBrandRequest{
		ID:           brandID,
		Name:         updateData.Name,
		Description:  updateData.Description,
		LogoURL:      updateData.LogoURL,
		Website:      updateData.Website,
		Aliases:      updateData.Aliases,
		CategoryTags: updateData.CategoryTags,
		Sources:      updateData.Sources,
		QualityScore: updateData.QualityScore,
		IsActive:     updateData.IsActive,
	}

	// Ejecutar el use case
	if err := h.updateUseCase.Execute(c.Request.Context(), updateRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error actualizando la marca: " + err.Error()})
		return
	}

	// Obtener la marca actualizada para devolverla
	updatedBrand, err := h.repository.FindByID(c.Request.Context(), brandID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo marca actualizada: " + err.Error()})
		return
	}

	// Preparar respuesta
	brandResponse := response.NewMarketplaceBrandResponse(updatedBrand)
	c.JSON(http.StatusOK, brandResponse)
}

// DeleteMarketplaceBrand maneja la solicitud para eliminar una marca marketplace
func (h *MarketplaceBrandHandler) DeleteMarketplaceBrand(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden eliminar marcas marketplace"})
		return
	}

	_ = c.Param("id") // TODO: usar id cuando se implemente
	// TODO: Implementar eliminación
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Eliminar marca marketplace - pendiente implementación completa"})
}

// VerifyMarketplaceBrand maneja la solicitud para verificar una marca marketplace
func (h *MarketplaceBrandHandler) VerifyMarketplaceBrand(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden verificar marcas marketplace"})
		return
	}

	_ = c.Param("id") // TODO: usar id cuando se implemente
	// TODO: Implementar verificación
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Verificar marca marketplace - pendiente implementación completa"})
}

// UnverifyMarketplaceBrand maneja la solicitud para desverificar una marca marketplace
func (h *MarketplaceBrandHandler) UnverifyMarketplaceBrand(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden desverificar marcas marketplace"})
		return
	}

	_ = c.Param("id") // TODO: usar id cuando se implemente
	// TODO: Implementar desverificación
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Desverificar marca marketplace - pendiente implementación completa"})
}
