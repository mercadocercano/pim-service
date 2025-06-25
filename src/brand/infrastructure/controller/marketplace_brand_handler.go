package controller

import (
	"net/http"
	"strconv"

	"pim/src/brand/application/request"
	"pim/src/brand/application/response"
	"pim/src/brand/application/usecase"
	"pim/src/brand/domain/port"
	"pim/src/shared/domain/criteria"

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
	criteriaBuilder := criteria.NewCriteriaBuilder().FromURLValues(c.Request.URL.Query())

	// Paginación por defecto si no se especifica
	if c.Query("page") == "" {
		criteriaBuilder.SetPagination(1, 20)
	}

	// Filtros
	if search := c.Query("search"); search != "" {
		criteriaBuilder.AddLikeFilter("name", search)
	}

	if isVerifiedStr := c.Query("is_verified"); isVerifiedStr != "" {
		if isVerified, err := strconv.ParseBool(isVerifiedStr); err == nil && isVerified {
			criteriaBuilder.AddEqualFilter("verification_status", "verified")
		}
	}

	// Ordenamiento por defecto: calidad primero, luego nombre
	criteriaBuilder.SetOrder("quality_score", "DESC")

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

	// TODO: Implementar binding de request y casos de uso
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Crear marca marketplace - pendiente implementación completa"})
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
