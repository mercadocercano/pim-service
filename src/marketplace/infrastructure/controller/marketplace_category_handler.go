package controller

import (
	"net/http"

	"pim/src/marketplace/application/request"
	"pim/src/marketplace/application/usecase"

	"github.com/gin-gonic/gin"
)

// MarketplaceCategoryHandler maneja las peticiones HTTP para categorías marketplace
type MarketplaceCategoryHandler struct {
	createMarketplaceCategoryUseCase   *usecase.CreateMarketplaceCategoryUseCase
	getAllMarketplaceCategoriesUseCase *usecase.GetAllMarketplaceCategoriesUseCase
	updateMarketplaceCategoryUseCase   *usecase.UpdateMarketplaceCategoryUseCase
	getTenantTaxonomyUseCase           *usecase.GetTenantTaxonomyUseCase
	validateCategoryHierarchyUseCase   *usecase.ValidateCategoryHierarchyUseCase
	syncMarketplaceChangesUseCase      *usecase.SyncMarketplaceChangesUseCase
}

// NewMarketplaceCategoryHandler crea una nueva instancia del manejador de categorías marketplace
func NewMarketplaceCategoryHandler(
	createMarketplaceCategoryUseCase *usecase.CreateMarketplaceCategoryUseCase,
	getAllMarketplaceCategoriesUseCase *usecase.GetAllMarketplaceCategoriesUseCase,
	updateMarketplaceCategoryUseCase *usecase.UpdateMarketplaceCategoryUseCase,
	getTenantTaxonomyUseCase *usecase.GetTenantTaxonomyUseCase,
	validateCategoryHierarchyUseCase *usecase.ValidateCategoryHierarchyUseCase,
	syncMarketplaceChangesUseCase *usecase.SyncMarketplaceChangesUseCase,
) *MarketplaceCategoryHandler {
	return &MarketplaceCategoryHandler{
		createMarketplaceCategoryUseCase:   createMarketplaceCategoryUseCase,
		getAllMarketplaceCategoriesUseCase: getAllMarketplaceCategoriesUseCase,
		updateMarketplaceCategoryUseCase:   updateMarketplaceCategoryUseCase,
		getTenantTaxonomyUseCase:           getTenantTaxonomyUseCase,
		validateCategoryHierarchyUseCase:   validateCategoryHierarchyUseCase,
		syncMarketplaceChangesUseCase:      syncMarketplaceChangesUseCase,
	}
}

// RegisterRoutes registra las rutas del API para categorías marketplace
func (h *MarketplaceCategoryHandler) RegisterRoutes(router *gin.RouterGroup) {
	marketplace := router.Group("/marketplace")
	{
		// Rutas para administradores (crear y obtener categorías globales)
		marketplace.GET("/categories", h.GetAllMarketplaceCategories)
		marketplace.POST("/categories", h.CreateMarketplaceCategory)
		marketplace.PUT("/categories/:id", h.UpdateMarketplaceCategory)
		marketplace.POST("/categories/validate-hierarchy", h.ValidateCategoryHierarchy)
		marketplace.POST("/sync-changes", h.SyncMarketplaceChanges)

		// Rutas para tenants (obtener taxonomía personalizada)
		marketplace.GET("/taxonomy", h.GetTenantTaxonomy)
	}
}

// CreateMarketplaceCategory maneja la solicitud para crear una nueva categoría marketplace
// Solo disponible para administradores del marketplace
func (h *MarketplaceCategoryHandler) CreateMarketplaceCategory(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden crear categorías marketplace"})
		return
	}

	var req request.CreateMarketplaceCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	category, err := h.createMarketplaceCategoryUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GetAllMarketplaceCategories maneja la solicitud para obtener todas las categorías marketplace
// Solo disponible para administradores del marketplace
func (h *MarketplaceCategoryHandler) GetAllMarketplaceCategories(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden ver categorías marketplace"})
		return
	}

	categories, err := h.getAllMarketplaceCategoriesUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// UpdateMarketplaceCategory maneja la solicitud para actualizar una categoría marketplace
// Solo disponible para administradores del marketplace
func (h *MarketplaceCategoryHandler) UpdateMarketplaceCategory(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden actualizar categorías marketplace"})
		return
	}

	// Obtener el ID de la categoría de la URL
	categoryID := c.Param("id")
	if categoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de categoría es requerido"})
		return
	}

	var req request.UpdateMarketplaceCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	category, err := h.updateMarketplaceCategoryUseCase.Execute(c.Request.Context(), categoryID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetTenantTaxonomy maneja la solicitud para obtener la taxonomía personalizada de un tenant
func (h *MarketplaceCategoryHandler) GetTenantTaxonomy(c *gin.Context) {
	// Obtener el tenantID del header
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "el header X-Tenant-ID es obligatorio"})
		return
	}

	// Parámetros opcionales de filtrado
	includeInactive := c.Query("include_inactive") == "true"
	includeCustomAttributes := c.Query("include_custom_attributes") == "true"
	includeMarketplaceData := c.Query("include_marketplace_data") == "true"
	format := c.Query("format")
	if format == "" {
		format = "tree"
	}

	req := request.GetTenantTaxonomyRequest{
		TenantID:                tenantID,
		IncludeInactive:         includeInactive,
		IncludeCustomAttributes: includeCustomAttributes,
		IncludeMarketplaceData:  includeMarketplaceData,
		Format:                  format,
	}

	taxonomy, err := h.getTenantTaxonomyUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, taxonomy)
}

// ValidateCategoryHierarchy maneja la solicitud para validar una jerarquía de categorías
func (h *MarketplaceCategoryHandler) ValidateCategoryHierarchy(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden validar jerarquías"})
		return
	}

	var req request.ValidateCategoryHierarchyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	validation, err := h.validateCategoryHierarchyUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, validation)
}

// SyncMarketplaceChanges maneja la solicitud para sincronizar cambios del marketplace
func (h *MarketplaceCategoryHandler) SyncMarketplaceChanges(c *gin.Context) {
	// Validar que el usuario tenga permisos de administrador
	userRole := c.GetHeader("X-User-Role")
	if userRole != "marketplace_admin" && userRole != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo administradores pueden sincronizar cambios"})
		return
	}

	var req request.SyncMarketplaceChangesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el formato de la petición: " + err.Error()})
		return
	}

	result, err := h.syncMarketplaceChangesUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
