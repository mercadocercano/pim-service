package controller

import (
	"net/http"

	"pim/src/brand/application/request"
	"pim/src/brand/application/usecase"
	"pim/src/brand/infrastructure/criteria"

	"github.com/gin-gonic/gin"
)

// BrandController maneja las peticiones HTTP para marcas
type BrandController struct {
	createBrandUseCase          *usecase.CreateBrandUseCase
	getBrandByIDUseCase         *usecase.GetBrandByIDUseCase
	updateBrandUseCase          *usecase.UpdateBrandUseCase
	deleteBrandUseCase          *usecase.DeleteBrandUseCase
	listBrandsByCriteriaUseCase *usecase.ListBrandsByCriteriaUseCase
	criteriaBuilder             *criteria.BrandCriteriaBuilder
}

// NewBrandController crea una nueva instancia del controller
func NewBrandController(
	createBrandUseCase *usecase.CreateBrandUseCase,
	getBrandByIDUseCase *usecase.GetBrandByIDUseCase,
	updateBrandUseCase *usecase.UpdateBrandUseCase,
	deleteBrandUseCase *usecase.DeleteBrandUseCase,
	listBrandsByCriteriaUseCase *usecase.ListBrandsByCriteriaUseCase,
	criteriaBuilder *criteria.BrandCriteriaBuilder,
) *BrandController {
	return &BrandController{
		createBrandUseCase:          createBrandUseCase,
		getBrandByIDUseCase:         getBrandByIDUseCase,
		updateBrandUseCase:          updateBrandUseCase,
		deleteBrandUseCase:          deleteBrandUseCase,
		listBrandsByCriteriaUseCase: listBrandsByCriteriaUseCase,
		criteriaBuilder:             criteriaBuilder,
	}
}

// CreateBrand godoc
// @Summary Crear una nueva marca
// @Description Crea una nueva marca en el sistema
// @Tags brands
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "ID del tenant"
// @Param brand body request.CreateBrandRequest true "Datos de la marca"
// @Success 201 {object} response.BrandResponse
// @Failure 400 {object} gin.H
// @Failure 409 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /brands [post]
func (ctrl *BrandController) CreateBrand(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	var req request.CreateBrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand, err := ctrl.createBrandUseCase.Execute(c.Request.Context(), &req, tenantID)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, brand)
}

// GetBrand godoc
// @Summary Obtener una marca por ID
// @Description Obtiene los detalles de una marca específica
// @Tags brands
// @Produce json
// @Param X-Tenant-ID header string true "ID del tenant"
// @Param id path string true "ID de la marca"
// @Success 200 {object} response.BrandResponse
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /brands/{id} [get]
func (ctrl *BrandController) GetBrand(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	brandID := c.Param("id")
	if brandID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de marca es requerido"})
		return
	}

	brand, err := ctrl.getBrandByIDUseCase.Execute(c.Request.Context(), brandID, tenantID)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, brand)
}

// UpdateBrand godoc
// @Summary Actualizar una marca
// @Description Actualiza los datos de una marca existente
// @Tags brands
// @Accept json
// @Produce json
// @Param X-Tenant-ID header string true "ID del tenant"
// @Param id path string true "ID de la marca"
// @Param brand body request.UpdateBrandRequest true "Datos actualizados de la marca"
// @Success 200 {object} response.BrandResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 409 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /brands/{id} [put]
func (ctrl *BrandController) UpdateBrand(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	brandID := c.Param("id")
	if brandID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de marca es requerido"})
		return
	}

	var req request.UpdateBrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand, err := ctrl.updateBrandUseCase.Execute(c.Request.Context(), brandID, &req, tenantID)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, brand)
}

// DeleteBrand godoc
// @Summary Eliminar una marca
// @Description Elimina una marca del sistema (soft delete)
// @Tags brands
// @Param X-Tenant-ID header string true "ID del tenant"
// @Param id path string true "ID de la marca"
// @Success 204
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /brands/{id} [delete]
func (ctrl *BrandController) DeleteBrand(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	brandID := c.Param("id")
	if brandID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de marca es requerido"})
		return
	}

	err := ctrl.deleteBrandUseCase.Execute(c.Request.Context(), brandID, tenantID)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// ListBrands godoc
// @Summary Listar marcas con filtros y paginación
// @Description Obtiene una lista de marcas con soporte para filtros, ordenamiento y paginación
// @Tags brands
// @Produce json
// @Param X-Tenant-ID header string true "ID del tenant"
// @Param page query int false "Número de página" default(1)
// @Param page_size query int false "Tamaño de página" default(10)
// @Param sort_by query string false "Campo de ordenamiento" default("name")
// @Param sort_dir query string false "Dirección de ordenamiento" Enums(asc, desc) default("asc")
// @Param status query string false "Estado de la marca" Enums(active, inactive, deleted)
// @Param name query string false "Búsqueda por nombre"
// @Param description query string false "Búsqueda por descripción"
// @Param active query boolean false "Solo marcas activas"
// @Param include_deleted query boolean false "Incluir marcas eliminadas"
// @Param has_website query boolean false "Filtrar por presencia de sitio web"
// @Param has_logo query boolean false "Filtrar por presencia de logo"
// @Success 200 {object} response.BrandListResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /brands [get]
func (ctrl *BrandController) ListBrands(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Tenant-ID header es requerido"})
		return
	}

	// Construir criterios validados
	validCriteria := ctrl.criteriaBuilder.BuildValidated(c)

	// Ejecutar búsqueda
	result, err := ctrl.listBrandsByCriteriaUseCase.Execute(c.Request.Context(), validCriteria)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// handleError maneja los errores de manera consistente
func (ctrl *BrandController) handleError(c *gin.Context, err error) {
	switch err.Error() {
	case "marca no encontrada":
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case "ya existe una marca con ese nombre":
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case "el nombre de la marca es obligatorio":
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case "el tenant ID es obligatorio":
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case "no se puede eliminar la marca":
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
	}
}

// RegisterRoutes registra las rutas del módulo Brand
func (ctrl *BrandController) RegisterRoutes(router *gin.RouterGroup) {
	brands := router.Group("/brands")
	{
		brands.POST("", ctrl.CreateBrand)
		brands.GET("", ctrl.ListBrands)
		brands.GET("/:id", ctrl.GetBrand)
		brands.PUT("/:id", ctrl.UpdateBrand)
		brands.DELETE("/:id", ctrl.DeleteBrand)
	}
}
