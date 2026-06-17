package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"time"

	businessTypeUsecase "saas-mt-pim-service/src/businesstype/application/usecase"
	"saas-mt-pim-service/src/quickstart/domain/entity"
	"saas-mt-pim-service/src/quickstart/domain/port"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// SimpleWizardHandler maneja los endpoints del wizard de forma simplificada
// usando solo los business types desde la BD, sin template ni tenant setup
type SimpleWizardHandler struct {
	listBusinessTypesUseCase *businessTypeUsecase.ListBusinessTypesUseCase
	historyRepo              port.QuickstartHistoryRepository
	db                       *sql.DB
}

// NewSimpleWizardHandler crea un nuevo handler simplificado del wizard
func NewSimpleWizardHandler(
	listBusinessTypesUseCase *businessTypeUsecase.ListBusinessTypesUseCase,
	db *sql.DB,
	historyRepo port.QuickstartHistoryRepository,
) *SimpleWizardHandler {
	return &SimpleWizardHandler{
		listBusinessTypesUseCase: listBusinessTypesUseCase,
		db:                       db,
		historyRepo:              historyRepo,
	}
}

// RegisterRoutes registra las rutas del wizard simplificado
func (h *SimpleWizardHandler) RegisterRoutes(router *gin.RouterGroup) {
	wizard := router.Group("/wizard")
	{
		wizard.GET("/status", h.GetWizardStatus)
		wizard.POST("/start", h.StartWizard)
		wizard.PUT("/step", h.UpdateWizardStep)
		wizard.GET("/template/:businessTypeId", h.GetTemplateData)
		wizard.GET("/template/:businessTypeId/:section", h.GetTemplateSectionData)
		wizard.POST("/complete", h.CompleteWizard)
		wizard.DELETE("/reset", h.ResetQuickstart) // TEMPORAL - BORRAR DESPUÉS DE PRUEBAS
	}
}

// GetWizardStatus obtiene el estado actual del wizard
func (h *SimpleWizardHandler) GetWizardStatus(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "X-Tenant-ID header is required",
		})
		return
	}

	// Buscar wizard activo en la base de datos
	history, err := h.historyRepo.FindActiveByTenantID(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al buscar wizard activo",
		})
		return
	}

	// Si no hay wizard activo, verificar si ya se completó uno anteriormente
	if history == nil {
		// Verificar si existe un wizard completado
		var completedCount int
		err = h.db.QueryRow(`
			SELECT COUNT(*) FROM tenant_quickstart_history 
			WHERE tenant_id = $1 AND setup_completed = true
		`, tenantID).Scan(&completedCount)

		if err == nil && completedCount > 0 {
			// Ya completó el wizard anteriormente
			c.JSON(http.StatusOK, gin.H{
				"wizard_id":       nil,
				"tenant_id":       tenantID,
				"setup_data":      gin.H{},
				"setup_completed": true,
				"message":         "El wizard ya fue completado anteriormente",
			})
			return
		}

		// No hay wizard activo ni completado
		c.JSON(http.StatusOK, gin.H{
			"wizard_id": nil,
			"tenant_id": tenantID,
			"setup_data": gin.H{
				"step":            "not_started",
				"completed_steps": []string{},
			},
			"setup_completed": false,
		})
		return
	}

	// Parsear setup_data
	var setupData map[string]interface{}
	if err := json.Unmarshal(history.SetupData, &setupData); err != nil {
		setupData = map[string]interface{}{
			"step":            "unknown",
			"completed_steps": []string{},
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"wizard_id":        history.ID,
		"tenant_id":        history.TenantID,
		"business_type_id": history.BusinessTypeID,
		"setup_data":       setupData,
		"setup_completed":  history.SetupCompleted,
		"created_at":       history.CreatedAt,
		"updated_at":       history.UpdatedAt,
	})
}

// StartWizard inicia el proceso del wizard
func (h *SimpleWizardHandler) StartWizard(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "X-Tenant-ID header is required",
		})
		return
	}

	var req struct {
		BusinessTypeID string `json:"business_type_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Verificar si ya existe un wizard activo
	existingHistory, err := h.historyRepo.FindActiveByTenantID(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al verificar wizard existente",
		})
		return
	}

	if existingHistory != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error":     "Ya existe un wizard activo para este tenant",
			"wizard_id": existingHistory.ID,
			"message":   "Debe completar o cancelar el wizard actual antes de iniciar uno nuevo",
		})
		return
	}

	// Verificar si ya existe un wizard completado para este tenant
	var completedCount int
	err = h.db.QueryRow(`
		SELECT COUNT(*) FROM tenant_quickstart_history 
		WHERE tenant_id = $1 AND setup_completed = true
	`, tenantID).Scan(&completedCount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al verificar historial de wizards",
		})
		return
	}

	if completedCount > 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"error":      "Este tenant ya ha completado el proceso de onboarding",
			"message":    "El wizard de configuración inicial solo puede ejecutarse una vez",
			"suggestion": "Use las opciones de administración para modificar su configuración",
		})
		return
	}

	// Crear nuevo wizard
	setupData := map[string]interface{}{
		"step":             "business_type_selected",
		"completed_steps":  []string{"business_type_selected"},
		"business_type_id": req.BusinessTypeID,
		"current_step":     1,
	}

	setupDataJSON, _ := json.Marshal(setupData)

	history := &entity.TenantQuickstartHistory{
		ID:             uuid.New().String(),
		TenantID:       tenantID,
		BusinessTypeID: req.BusinessTypeID,
		SetupCompleted: false,
		SetupData:      setupDataJSON,
	}

	// Guardar en la base de datos
	if err := h.historyRepo.Create(c.Request.Context(), history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al crear wizard",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wizard_id":        history.ID,
		"tenant_id":        history.TenantID,
		"business_type_id": history.BusinessTypeID,
		"setup_data":       setupData,
		"setup_completed":  history.SetupCompleted,
		"created_at":       history.CreatedAt,
	})
}

// UpdateWizardStep actualiza el paso actual del wizard
func (h *SimpleWizardHandler) UpdateWizardStep(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "X-Tenant-ID header is required",
		})
		return
	}

	var req struct {
		Step int                    `json:"step" binding:"required"`
		Data map[string]interface{} `json:"data" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + err.Error(),
		})
		return
	}

	// Buscar wizard activo
	history, err := h.historyRepo.FindActiveByTenantID(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al buscar wizard activo",
		})
		return
	}

	if history == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No hay wizard activo para este tenant",
		})
		return
	}

	// Actualizar setup_data
	var setupData map[string]interface{}
	if err := json.Unmarshal(history.SetupData, &setupData); err != nil {
		setupData = make(map[string]interface{})
	}

	// Actualizar datos del paso
	stepKey := getStepKey(req.Step)
	setupData[stepKey] = req.Data
	setupData["current_step"] = req.Step

	// Actualizar completed_steps
	completedSteps, ok := setupData["completed_steps"].([]interface{})
	if !ok {
		completedSteps = []interface{}{}
	}

	// Agregar el paso actual si no está
	stepCompleted := false
	for _, s := range completedSteps {
		if s == stepKey {
			stepCompleted = true
			break
		}
	}
	if !stepCompleted {
		completedSteps = append(completedSteps, stepKey)
		setupData["completed_steps"] = completedSteps
	}

	// Serializar y actualizar
	setupDataJSON, _ := json.Marshal(setupData)
	history.SetupData = setupDataJSON

	if err := h.historyRepo.Update(c.Request.Context(), history); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al actualizar wizard",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wizard_id":       history.ID,
		"tenant_id":       history.TenantID,
		"setup_data":      setupData,
		"setup_completed": history.SetupCompleted,
		"updated_at":      history.UpdatedAt,
	})
}

// getStepKey devuelve la clave para identificar cada paso
func getStepKey(step int) string {
	switch step {
	case 1:
		return "business_type_selected"
	case 2:
		return "categories_attributes_selected"
	case 3:
		return "brands_products_selected"
	default:
		return "unknown_step"
	}
}

// GetTemplateData obtiene los datos del template completo
func (h *SimpleWizardHandler) GetTemplateData(c *gin.Context) {
	businessTypeID := c.Param("businessTypeId")
	if businessTypeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "business type ID is required",
		})
		return
	}

	// Datos de ejemplo basados en el tipo de negocio
	templateData := h.getTemplateDataByBusinessType(businessTypeID)

	c.JSON(http.StatusOK, gin.H{
		"template_data": templateData,
	})
}

// getTemplateDataByBusinessType devuelve datos reales desde la tabla business_type_templates
func (h *SimpleWizardHandler) getTemplateDataByBusinessType(businessTypeID string) gin.H {
	// Query para obtener el template por defecto del business type
	query := `
		SELECT categories, attributes, products, brands
		FROM business_type_templates
		WHERE business_type_id = $1 AND is_default = true AND is_active = true
		LIMIT 1
	`

	var categoriesJSON, attributesJSON, productsJSON, brandsJSON []byte
	err := h.db.QueryRow(query, businessTypeID).Scan(&categoriesJSON, &attributesJSON, &productsJSON, &brandsJSON)

	if err != nil {
		// Si no hay datos, devolver datos de ejemplo
		return h.getDefaultTemplateData(businessTypeID)
	}

	// Parsear los JSONs
	var categories []interface{}
	var attributes []interface{}
	var products []interface{}
	var brands []interface{}

	json.Unmarshal(categoriesJSON, &categories)
	json.Unmarshal(attributesJSON, &attributes)
	json.Unmarshal(productsJSON, &products)
	json.Unmarshal(brandsJSON, &brands)

	// Procesar categorías para agregar IDs si faltan
	processedCategories := make([]gin.H, 0)
	for i, cat := range categories {
		if catMap, ok := cat.(map[string]interface{}); ok {
			// Agregar ID si no existe
			if _, hasID := catMap["id"]; !hasID || catMap["id"] == "" {
				catMap["id"] = i + 1
			}
			processedCategories = append(processedCategories, catMap)
		}
	}

	// Procesar atributos
	processedAttributes := make([]gin.H, 0)
	for i, attr := range attributes {
		if attrMap, ok := attr.(map[string]interface{}); ok {
			if _, hasID := attrMap["id"]; !hasID || attrMap["id"] == "" {
				attrMap["id"] = i + 1
			}
			processedAttributes = append(processedAttributes, attrMap)
		}
	}

	// Procesar marcas
	processedBrands := make([]gin.H, 0)
	for i, brand := range brands {
		if brandMap, ok := brand.(map[string]interface{}); ok {
			if _, hasID := brandMap["id"]; !hasID || brandMap["id"] == "" {
				brandMap["id"] = i + 1
			}
			processedBrands = append(processedBrands, brandMap)
		}
	}

	// Procesar productos
	processedProducts := make([]gin.H, 0)
	for i, prod := range products {
		if prodMap, ok := prod.(map[string]interface{}); ok {
			if _, hasID := prodMap["id"]; !hasID || prodMap["id"] == "" {
				prodMap["id"] = i + 1
			}
			processedProducts = append(processedProducts, prodMap)
		}
	}

	return gin.H{
		"categories": processedCategories,
		"attributes": processedAttributes,
		"products":   processedProducts,
		"brands":     processedBrands,
		"metadata": gin.H{
			"business_type_id": businessTypeID,
			"total_categories": len(processedCategories),
			"total_products":   len(processedProducts),
			"total_brands":     len(processedBrands),
			"total_attributes": len(processedAttributes),
		},
	}
}

// getDefaultTemplateData devuelve datos de ejemplo por defecto
func (h *SimpleWizardHandler) getDefaultTemplateData(businessTypeID string) gin.H {
	return gin.H{
		"categories": []gin.H{
			{"id": 1, "name": "Bebidas", "code": "beverages", "description": "Bebidas en general"},
			{"id": 2, "name": "Lácteos", "code": "dairy", "description": "Productos lácteos"},
			{"id": 3, "name": "Panadería", "code": "bakery", "description": "Pan y productos de panadería"},
			{"id": 4, "name": "Limpieza", "code": "cleaning", "description": "Productos de limpieza"},
			{"id": 5, "name": "Golosinas", "code": "candy", "description": "Dulces y golosinas"},
			{"id": 6, "name": "Conservas", "code": "canned", "description": "Productos en conserva"},
			{"id": 7, "name": "Congelados", "code": "frozen", "description": "Productos congelados"},
			{"id": 8, "name": "Snacks", "code": "snacks", "description": "Aperitivos y snacks"},
		},
		"products": []gin.H{
			{"id": 1, "name": "Coca-Cola 2.5L", "category": "Bebidas", "brand": "Coca-Cola"},
			{"id": 2, "name": "Leche Entera 1L", "category": "Lácteos", "brand": "La Serenísima"},
			{"id": 3, "name": "Pan Blanco", "category": "Panadería", "brand": "Bimbo"},
			{"id": 4, "name": "Detergente 500ml", "category": "Limpieza", "brand": "Skip"},
			{"id": 5, "name": "Caramelos Surtidos", "category": "Golosinas", "brand": "Arcor"},
			{"id": 6, "name": "Atún en Aceite", "category": "Conservas", "brand": "La Campagnola"},
			{"id": 7, "name": "Helado 1kg", "category": "Congelados", "brand": "Frigor"},
			{"id": 8, "name": "Papas Fritas", "category": "Snacks", "brand": "Lays"},
		},
		"brands": []gin.H{
			{"id": 1, "name": "Coca-Cola"},
			{"id": 2, "name": "La Serenísima"},
			{"id": 3, "name": "Bimbo"},
			{"id": 4, "name": "Skip"},
			{"id": 5, "name": "Arcor"},
			{"id": 6, "name": "La Campagnola"},
			{"id": 7, "name": "Frigor"},
			{"id": 8, "name": "Lays"},
			{"id": 9, "name": "Quilmes"},
			{"id": 10, "name": "Marolio"},
		},
		"attributes": []gin.H{
			{"id": 1, "name": "Precio", "type": "number"},
			{"id": 2, "name": "Stock", "type": "number"},
			{"id": 3, "name": "Código de barras", "type": "text"},
			{"id": 4, "name": "Fecha de vencimiento", "type": "date"},
			{"id": 5, "name": "Peso/Volumen", "type": "text"},
			{"id": 6, "name": "Unidad de medida", "type": "select"},
		},
		"metadata": gin.H{
			"business_type_id": businessTypeID,
			"total_categories": 8,
			"total_products":   8,
			"total_brands":     10,
			"total_attributes": 6,
		},
	}
}

// GetTemplateSectionData obtiene los datos de una sección específica del template
func (h *SimpleWizardHandler) GetTemplateSectionData(c *gin.Context) {
	businessTypeID := c.Param("businessTypeId")
	section := c.Param("section")

	if businessTypeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "business type ID is required",
		})
		return
	}

	// Parámetros de paginación
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	var data interface{}

	switch section {
	case "categories":
		data = []gin.H{
			{"id": 1, "name": "Electrónicos", "code": "electronics"},
			{"id": 2, "name": "Ropa", "code": "clothing"},
			{"id": 3, "name": "Hogar", "code": "home"},
		}
	case "products":
		data = []gin.H{
			{"id": 1, "name": "Smartphone", "category": "electronics"},
			{"id": 2, "name": "Laptop", "category": "electronics"},
			{"id": 3, "name": "Camiseta", "category": "clothing"},
		}
	case "brands":
		data = []string{"Samsung", "Apple", "LG", "Sony", "Nike", "Adidas"}
	case "attributes":
		data = []gin.H{
			{"id": 1, "name": "Color", "type": "text"},
			{"id": 2, "name": "Tamaño", "type": "text"},
			{"id": 3, "name": "Material", "type": "text"},
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid section. Valid sections are: categories, products, brands, attributes",
		})
		return
	}

	// Calcular longitud basado en el tipo de data
	var totalItems int
	switch d := data.(type) {
	case []gin.H:
		totalItems = len(d)
	case []string:
		totalItems = len(d)
	default:
		totalItems = 0
	}

	// Respuesta con paginación
	c.JSON(http.StatusOK, gin.H{
		"section": section,
		"data":    data,
		"pagination": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total_items": totalItems,
			"has_next":    false,
			"has_prev":    false,
		},
	})
}

// CompleteWizard completa el proceso del wizard
func (h *SimpleWizardHandler) CompleteWizard(c *gin.Context) {
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "X-Tenant-ID header is required",
		})
		return
	}

	// Buscar wizard activo
	history, err := h.historyRepo.FindActiveByTenantID(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al buscar wizard activo",
		})
		return
	}

	if history == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No hay wizard activo para este tenant",
		})
		return
	}

	// Parsear setup_data para obtener las selecciones
	var setupData map[string]interface{}
	if err := json.Unmarshal(history.SetupData, &setupData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al procesar datos del wizard",
		})
		return
	}

	// Iniciar transacción
	tx, err := h.db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al iniciar transacción",
		})
		return
	}
	defer tx.Rollback()

	summary := gin.H{
		"categories_created": 0,
		"attributes_created": 0,
		"brands_created":     0,
		"products_created":   0,
	}

	// Procesar categorías y atributos seleccionados
	if catAttrData, ok := setupData["categories_attributes_selected"].(map[string]interface{}); ok {
		// Crear categorías
		if selectedCategories, ok := catAttrData["selectedCategories"].([]interface{}); ok {
			for _, cat := range selectedCategories {
				if catMap, ok := cat.(map[string]interface{}); ok {
					if err := h.createCategory(tx, tenantID, catMap); err == nil {
						summary["categories_created"] = summary["categories_created"].(int) + 1
					}
				}
			}
		}

		// Crear atributos
		if selectedAttributes, ok := catAttrData["selectedAttributes"].([]interface{}); ok {
			for _, attr := range selectedAttributes {
				if attrMap, ok := attr.(map[string]interface{}); ok {
					if err := h.createAttribute(tx, tenantID, attrMap); err == nil {
						summary["attributes_created"] = summary["attributes_created"].(int) + 1
					}
				}
			}
		}
	}

	// Procesar marcas y productos seleccionados
	if brandProdData, ok := setupData["brands_products_selected"].(map[string]interface{}); ok {
		// Crear marcas
		if selectedBrands, ok := brandProdData["selectedBrands"].([]interface{}); ok {
			for _, brand := range selectedBrands {
				if brandMap, ok := brand.(map[string]interface{}); ok {
					if err := h.createBrand(tx, tenantID, brandMap); err == nil {
						summary["brands_created"] = summary["brands_created"].(int) + 1
					}
				}
			}
		}

		// Crear productos
		if selectedProducts, ok := brandProdData["selectedProducts"].([]interface{}); ok {
			for _, prod := range selectedProducts {
				if prodMap, ok := prod.(map[string]interface{}); ok {
					if err := h.createProduct(tx, tenantID, prodMap); err == nil {
						summary["products_created"] = summary["products_created"].(int) + 1
					}
				}
			}
		}
	}

	// Commit de la transacción
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al confirmar transacción: " + err.Error(),
		})
		return
	}

	// Marcar wizard como completado después del commit exitoso.
	// Limpiar primero wizards completados anteriores del tenant (restricción UNIQUE).
	h.db.Exec(`
		DELETE FROM tenant_quickstart_history
		WHERE tenant_id = $1 AND setup_completed = true AND id != $2
	`, tenantID, history.ID)

	if err := h.historyRepo.MarkAsCompleted(c.Request.Context(), history.ID); err != nil {
		// Fallback: query directa si el repo falla
		_, directErr := h.db.Exec(`
			UPDATE tenant_quickstart_history
			SET setup_completed = true, updated_at = CURRENT_TIMESTAMP
			WHERE id = $1
		`, history.ID)

		if directErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error al marcar wizard como completado: " + err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"message":      "Wizard completado exitosamente. Los datos han sido importados.",
		"tenant_id":    tenantID,
		"wizard_id":    history.ID,
		"completed_at": time.Now().Format(time.RFC3339),
		"summary":      summary,
		"next_steps": gin.H{
			"generate_variants": "Los productos están en estado inactivo. El siguiente paso es generar las variantes con sus precios.",
			"review_products":   "Revisar y activar los productos después de configurar las variantes.",
		},
	})
}

// createCategory crea una categoría en la base de datos
func (h *SimpleWizardHandler) createCategory(tx *sql.Tx, tenantID string, category map[string]interface{}) error {
	id := uuid.New().String()
	name := ""
	description := ""

	// Extraer campos del mapa
	if n, ok := category["name"].(string); ok {
		name = n
	}
	if d, ok := category["description"].(string); ok {
		description = d
	}

	// Verificar si la categoría ya existe para este tenant
	var existingID string
	err := tx.QueryRow(`
		SELECT id FROM categories 
		WHERE tenant_id = $1 AND name = $2 AND status = 'active'
		LIMIT 1
	`, tenantID, name).Scan(&existingID)

	if err == sql.ErrNoRows {
		// No existe, crearla
		query := `
			INSERT INTO categories (id, tenant_id, name, description, status, created_at, updated_at)
			VALUES ($1, $2, $3, $4, 'active', $5, $6)
		`

		now := time.Now()
		_, err = tx.Exec(query, id, tenantID, name, description, now, now)
	}

	return err
}

// createAttribute crea un atributo en la base de datos
func (h *SimpleWizardHandler) createAttribute(tx *sql.Tx, tenantID string, attribute map[string]interface{}) error {
	id := uuid.New().String()
	name := ""
	description := ""
	attrType := "text"
	required := false
	var options []string

	// Extraer campos del mapa
	if n, ok := attribute["name"].(string); ok {
		name = n
	}
	if d, ok := attribute["description"].(string); ok {
		description = d
	}
	if t, ok := attribute["type"].(string); ok {
		attrType = t
	}
	if r, ok := attribute["required"].(bool); ok {
		required = r
	}

	// Extraer opciones si existen
	if opts, ok := attribute["options"].([]interface{}); ok {
		for _, opt := range opts {
			if optStr, ok := opt.(string); ok {
				options = append(options, optStr)
			}
		}
	}

	// Verificar si el atributo ya existe para este tenant
	var existingID string
	err := tx.QueryRow(`
		SELECT id FROM attributes 
		WHERE tenant_id = $1 AND name = $2 AND status = 'active'
		LIMIT 1
	`, tenantID, name).Scan(&existingID)

	if err == sql.ErrNoRows {
		// No existe, crear nuevo atributo
		query := `
			INSERT INTO attributes (
				id, tenant_id, name, description, type, required, options,
				status, created_at, updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7, 'active', $8, $9)
		`

		now := time.Now()
		_, err = tx.Exec(query, id, tenantID, name, description, attrType, required,
			pq.Array(options), now, now)
	}

	return err
}

// createBrand crea una marca en la base de datos
func (h *SimpleWizardHandler) createBrand(tx *sql.Tx, tenantID string, brand map[string]interface{}) error {
	id := uuid.New().String()
	name := ""
	description := ""

	// Extraer campos del mapa
	if n, ok := brand["name"].(string); ok {
		name = n
	}
	if d, ok := brand["description"].(string); ok {
		description = d
	}

	// Usar tabla brands que sí tiene tenant_id
	query := `
		INSERT INTO brands (id, tenant_id, name, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, 'active', $5, $6)
		ON CONFLICT (tenant_id, name) DO NOTHING
	`

	now := time.Now()
	_, err := tx.Exec(query, id, tenantID, name, description, now, now)
	return err
}

// createProduct crea un producto en la base de datos
func (h *SimpleWizardHandler) createProduct(tx *sql.Tx, tenantID string, product map[string]interface{}) error {
	id := uuid.New().String()
	name := ""
	categoryName := ""
	brandName := ""

	// Extraer campos del mapa
	if n, ok := product["name"].(string); ok {
		name = n
	}
	if c, ok := product["category"].(string); ok {
		categoryName = c
	}
	if b, ok := product["brand"].(string); ok {
		brandName = b
	}

	// Buscar IDs de categoría y marca por nombre
	var categoryID, brandID sql.NullString

	if categoryName != "" {
		err := tx.QueryRow(`
			SELECT id FROM categories 
			WHERE tenant_id = $1 AND name = $2 AND status = 'active'
			LIMIT 1
		`, tenantID, categoryName).Scan(&categoryID)
		if err == nil {
			categoryID.Valid = true
		}
	}

	if brandName != "" {
		err := tx.QueryRow(`
			SELECT id FROM brands 
			WHERE tenant_id = $1 AND name = $2 AND status = 'active'
			LIMIT 1
		`, tenantID, brandName).Scan(&brandID)
		if err == nil {
			brandID.Valid = true
		}
	}

	// Verificar si el producto ya existe por nombre para este tenant
	var existingID string
	err := tx.QueryRow(`
		SELECT id FROM products 
		WHERE tenant_id = $1 AND name = $2 AND status != 'deleted'
		LIMIT 1
	`, tenantID, name).Scan(&existingID)

	if err == sql.ErrNoRows {
		// No existe, crear nuevo producto
		// Generar SKU único
		sku := "PRD-" + uuid.New().String()[:8]

		query := `
			INSERT INTO products (
				id, tenant_id, sku, name, category_id, brand_id, 
				status, created_at, updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, 'inactive', $7, $8)
		`

		now := time.Now()
		_, err = tx.Exec(query, id, tenantID, sku, name, categoryID, brandID, now, now)
	}

	return err
}

// ResetQuickstart resetea completamente el quickstart del tenant
// ⚠️ TEMPORAL - FUNCIÓN PELIGROSA SOLO PARA PRUEBAS - BORRAR EN PRODUCCIÓN
func (h *SimpleWizardHandler) ResetQuickstart(c *gin.Context) {
	// IMPORTANTE: Esta función debe deshabilitarse en producción
	env := os.Getenv("ENVIRONMENT")
	if env == "production" || env == "prod" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Esta función no está disponible en producción",
		})
		return
	}

	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "X-Tenant-ID header is required",
		})
		return
	}

	// Iniciar transacción para operaciones atómicas
	tx, err := h.db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al iniciar transacción: " + err.Error(),
		})
		return
	}
	defer tx.Rollback()

	deletedCounts := gin.H{
		"products_deleted":       0,
		"brands_deleted":         0,
		"categories_deleted":     0,
		"attributes_deleted":     0,
		"wizard_history_deleted": 0,
	}

	// 1. Borrar productos del tenant
	result, err := tx.Exec(`DELETE FROM products WHERE tenant_id = $1`, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error eliminando productos: " + err.Error(),
		})
		return
	}
	if count, _ := result.RowsAffected(); count > 0 {
		deletedCounts["products_deleted"] = count
	}

	// 2. Borrar marcas del tenant
	result, err = tx.Exec(`DELETE FROM brands WHERE tenant_id = $1`, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error eliminando marcas: " + err.Error(),
		})
		return
	}
	if count, _ := result.RowsAffected(); count > 0 {
		deletedCounts["brands_deleted"] = count
	}

	// 3. Borrar categorías del tenant
	result, err = tx.Exec(`DELETE FROM categories WHERE tenant_id = $1`, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error eliminando categorías: " + err.Error(),
		})
		return
	}
	if count, _ := result.RowsAffected(); count > 0 {
		deletedCounts["categories_deleted"] = count
	}

	// 4. Borrar atributos del tenant
	result, err = tx.Exec(`DELETE FROM attributes WHERE tenant_id = $1`, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error eliminando atributos: " + err.Error(),
		})
		return
	}
	if count, _ := result.RowsAffected(); count > 0 {
		deletedCounts["attributes_deleted"] = count
	}

	// 5. Borrar historial del wizard
	result, err = tx.Exec(`DELETE FROM tenant_quickstart_history WHERE tenant_id = $1`, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error eliminando historial del wizard: " + err.Error(),
		})
		return
	}
	if count, _ := result.RowsAffected(); count > 0 {
		deletedCounts["wizard_history_deleted"] = count
	}

	// Commit de la transacción
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al confirmar transacción: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"message":   "⚠️ RESET COMPLETO - Todos los datos del quickstart han sido eliminados",
		"tenant_id": tenantID,
		"deleted":   deletedCounts,
		"warning":   "FUNCIÓN TEMPORAL - SERÁ ELIMINADA DESPUÉS DE LAS PRUEBAS",
	})
}
