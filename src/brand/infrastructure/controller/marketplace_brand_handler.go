package controller

import (
	httpresp "github.com/hornosg/go-shared/infrastructure/response"
	"net/http"
	"strconv"

	cr "github.com/hornosg/go-shared/criteria"
	"saas-mt-pim-service/src/brand/application/request"
	"saas-mt-pim-service/src/brand/application/response"
	"saas-mt-pim-service/src/brand/application/usecase"
	"saas-mt-pim-service/src/brand/domain/entity"
	"saas-mt-pim-service/src/brand/domain/port"

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
		httpresp.JSON(c, http.StatusForbidden, "Solo administradores pueden ver marcas marketplace")
		return
	}

	// Campos permitidos para marketplace brands
	allowedFields := []string{"id", "name", "slug", "status", "verification_status", "is_active", "quality_score", "created_at", "updated_at"}

	criteriaBuilder := cr.NewCriteriaBuilder().FromURLValues(c.Request.URL.Query())

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

	if pageSize > 1000 {
		pageSize = 1000
	}

	criteriaBuilder.SetPagination(page, pageSize)

	if search := c.Query("search"); search != "" {
		criteriaBuilder.AddLikeFilter("name", search)
	}

	if isVerifiedStr := c.Query("is_verified"); isVerifiedStr != "" {
		if isVerified, err := strconv.ParseBool(isVerifiedStr); err == nil && isVerified {
			criteriaBuilder.AddEqualFilter("verification_status", "verified")
		}
	}

	if c.Query("sort_by") == "" {
		criteriaBuilder.SetOrder("name", "ASC")
	}

	builtCriteria := cr.Sanitize(criteriaBuilder.Build(), allowedFields)

	// Usar el repository para obtener datos
	brands, err := h.repository.SearchByCriteria(c.Request.Context(), builtCriteria)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Contar total
	total, err := h.repository.CountByCriteria(c.Request.Context(), builtCriteria)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, err.Error())
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
		httpresp.JSON(c, http.StatusForbidden, "Solo administradores pueden crear marcas marketplace")
		return
	}

	// Estructura para recibir los datos de creación
	var createData struct {
		Name         string   `json:"name" binding:"required"`
		Description  string   `json:"description"`
		LogoURL      string   `json:"logo_url" binding:"omitempty,url"`
		Website      string   `json:"website" binding:"omitempty,url"`
		Aliases      []string `json:"aliases"`
		CategoryTags []string `json:"category_tags"`

		BackgroundColor string `json:"background_color"`
		TextColor       string `json:"text_color"`
		Typography      string `json:"typography"`
	}

	if err := c.ShouldBindJSON(&createData); err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "Error en el formato de la petición: "+err.Error())
		return
	}

	// Validar que el nombre no esté vacío
	if createData.Name == "" {
		httpresp.JSON(c, http.StatusBadRequest, "El nombre de la marca es requerido")
		return
	}

	// Crear la entidad Marketplacebrand usando el constructor
	brand, err := entity.NewMarketplacebrand(createData.Name)
	if err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "Error al crear la marca: "+err.Error())
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

	// Identidad visual (valida hex/typography; "" = fallback del design system)
	if err := brand.SetVisualIdentity(entity.VisualIdentityParams{
		BackgroundColor: createData.BackgroundColor,
		TextColor:       createData.TextColor,
		Typography:      createData.Typography,
	}); err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "Error en identidad visual: "+err.Error())
		return
	}

	// Guardar en el repositorio
	if err := h.repository.Create(c.Request.Context(), brand); err != nil {
		// Verificar si es un error de duplicado
		if err.Error() == "marca ya existe" || err.Error() == "duplicate key" {
			httpresp.JSON(c, http.StatusConflict, "Una marca con ese nombre ya existe")
			return
		}
		httpresp.JSON(c, http.StatusInternalServerError, "Error al crear la marca: "+err.Error())
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
		httpresp.JSON(c, http.StatusForbidden, "Solo administradores pueden ver marcas marketplace")
		return
	}

	// Obtener ID de la marca
	brandID := c.Param("id")
	if brandID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "ID de marca es requerido")
		return
	}

	// Buscar la marca por ID
	brand, err := h.repository.FindByID(c.Request.Context(), brandID)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, "Error buscando la marca: "+err.Error())
		return
	}

	if brand == nil {
		httpresp.JSON(c, http.StatusNotFound, "Marca no encontrada")
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
		httpresp.JSON(c, http.StatusForbidden, "Solo administradores pueden actualizar marcas marketplace")
		return
	}

	// Obtener ID de la marca
	brandID := c.Param("id")
	if brandID == "" {
		httpresp.JSON(c, http.StatusBadRequest, "ID de marca es requerido")
		return
	}

	// Estructura para recibir los datos de actualización
	var updateData struct {
		Name         string   `json:"name"`
		Description  string   `json:"description"`
		LogoURL      string   `json:"logo_url" binding:"omitempty,url"`
		Website      string   `json:"website" binding:"omitempty,url"`
		Aliases      []string `json:"aliases"`
		CategoryTags []string `json:"category_tags"`
		Sources      []string `json:"sources"`
		QualityScore float64  `json:"quality_score"`
		IsActive     bool     `json:"is_active"`

		BackgroundColor string `json:"background_color"`
		TextColor       string `json:"text_color"`
		Typography      string `json:"typography"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		httpresp.JSON(c, http.StatusBadRequest, "Error en el formato de la petición: "+err.Error())
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

		BackgroundColor: updateData.BackgroundColor,
		TextColor:       updateData.TextColor,
		Typography:      updateData.Typography,
	}

	// Ejecutar el use case
	if err := h.updateUseCase.Execute(c.Request.Context(), updateRequest); err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, "Error actualizando la marca: "+err.Error())
		return
	}

	// Obtener la marca actualizada para devolverla
	updatedBrand, err := h.repository.FindByID(c.Request.Context(), brandID)
	if err != nil {
		httpresp.JSON(c, http.StatusInternalServerError, "Error obteniendo marca actualizada: "+err.Error())
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
		httpresp.JSON(c, http.StatusForbidden, "Solo administradores pueden eliminar marcas marketplace")
		return
	}

	_ = c.Param("id") // TODO: usar id cuando se implemente
	// TODO: Implementar eliminación
	httpresp.JSON(c, http.StatusNotImplemented, "Eliminar marca marketplace - pendiente implementación completa")
}

// VerifyMarketplaceBrand maneja la solicitud para verificar una marca marketplace
func (h *MarketplaceBrandHandler) VerifyMarketplaceBrand(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		httpresp.JSON(c, http.StatusForbidden, "Solo administradores pueden verificar marcas marketplace")
		return
	}

	_ = c.Param("id") // TODO: usar id cuando se implemente
	// TODO: Implementar verificación
	httpresp.JSON(c, http.StatusNotImplemented, "Verificar marca marketplace - pendiente implementación completa")
}

// UnverifyMarketplaceBrand maneja la solicitud para desverificar una marca marketplace
func (h *MarketplaceBrandHandler) UnverifyMarketplaceBrand(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		httpresp.JSON(c, http.StatusForbidden, "Solo administradores pueden desverificar marcas marketplace")
		return
	}

	_ = c.Param("id") // TODO: usar id cuando se implemente
	// TODO: Implementar desverificación
	httpresp.JSON(c, http.StatusNotImplemented, "Desverificar marca marketplace - pendiente implementación completa")
}
