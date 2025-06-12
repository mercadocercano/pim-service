package usecase

import (
	"context"
	"fmt"
	"time"

	"pim/src/marketplace/application/request"
	"pim/src/marketplace/application/response"
	"pim/src/marketplace/domain/entity"
	"pim/src/marketplace/domain/port"
)

// GetTenantTaxonomyUseCase maneja la obtención de la taxonomía completa de un tenant
type GetTenantTaxonomyUseCase struct {
	categoryRepo        port.MarketplaceCategoryRepository
	mappingRepo         port.TenantCategoryMappingRepository
	customAttributeRepo port.TenantCustomAttributeRepository
}

// NewGetTenantTaxonomyUseCase crea una nueva instancia del caso de uso
func NewGetTenantTaxonomyUseCase(
	categoryRepo port.MarketplaceCategoryRepository,
	mappingRepo port.TenantCategoryMappingRepository,
	customAttributeRepo port.TenantCustomAttributeRepository,
) *GetTenantTaxonomyUseCase {
	return &GetTenantTaxonomyUseCase{
		categoryRepo:        categoryRepo,
		mappingRepo:         mappingRepo,
		customAttributeRepo: customAttributeRepo,
	}
}

// Execute ejecuta el caso de uso para obtener la taxonomía del tenant
func (uc *GetTenantTaxonomyUseCase) Execute(
	ctx context.Context,
	req *request.GetTenantTaxonomyRequest,
) (*response.TenantTaxonomyResponse, error) {
	// Validar la petición
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Obtener todos los mapeos del tenant
	mappings, err := uc.mappingRepo.GetByTenantID(ctx, req.TenantID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant mappings: %w", err)
	}

	// Crear mapa de mapeos por categoría para acceso rápido
	mappingsByCategory := make(map[string]*entity.TenantCategoryMapping)
	mappingsByMarketplaceCategory := make(map[string]*entity.TenantCategoryMapping)

	for _, mapping := range mappings {
		mappingsByCategory[mapping.CategoryID] = mapping
		mappingsByMarketplaceCategory[mapping.MarketplaceCategoryID] = mapping
	}

	// Obtener categorías marketplace si se necesitan
	var marketplaceCategories map[string]*entity.MarketplaceCategory
	if req.IncludeMarketplaceData {
		marketplaceCategories, err = uc.getMarketplaceCategories(ctx, mappings)
		if err != nil {
			return nil, fmt.Errorf("failed to get marketplace categories: %w", err)
		}
	}

	// Obtener atributos personalizados si se solicitan
	var customAttributes []*entity.TenantCustomAttribute
	if req.IncludeCustomAttributes {
		customAttributes, err = uc.customAttributeRepo.GetByTenantID(ctx, req.TenantID)
		if err != nil {
			return nil, fmt.Errorf("failed to get custom attributes: %w", err)
		}
	}

	// TODO: Obtener categorías del tenant desde el servicio correspondiente
	// Por ahora simulamos categorías tenant basadas en los mapeos
	tenantCategories := uc.buildTenantCategoriesFromMappings(mappings)

	// Construir nodos de taxonomía
	categoryNodes := uc.buildCategoryNodes(
		tenantCategories,
		mappingsByCategory,
		marketplaceCategories,
		customAttributes,
		req,
	)

	// Formatear según el tipo solicitado
	formattedCategories := uc.formatCategories(categoryNodes, req.Format, req.MaxDepth)

	// Construir metadatos
	metadata := uc.buildMetadata(tenantCategories, mappings, customAttributes, req)

	// Construir respuesta
	taxonomyResponse := &response.TenantTaxonomyResponse{
		TenantID:              req.TenantID,
		Format:                req.Format,
		TotalCategories:       len(tenantCategories),
		TotalMappings:         len(mappings),
		TotalCustomAttributes: len(customAttributes),
		Categories:            formattedCategories,
		Metadata:              metadata,
		GeneratedAt:           time.Now(),
	}

	// Incluir atributos personalizados en la respuesta si se solicitan
	if req.IncludeCustomAttributes {
		taxonomyResponse.CustomAttributes = uc.buildCustomAttributeInfos(customAttributes)
	}

	return taxonomyResponse, nil
}

// getMarketplaceCategories obtiene las categorías marketplace necesarias
func (uc *GetTenantTaxonomyUseCase) getMarketplaceCategories(
	ctx context.Context,
	mappings []*entity.TenantCategoryMapping,
) (map[string]*entity.MarketplaceCategory, error) {
	categories := make(map[string]*entity.MarketplaceCategory)

	for _, mapping := range mappings {
		if _, exists := categories[mapping.MarketplaceCategoryID]; !exists {
			category, err := uc.categoryRepo.GetByID(ctx, mapping.MarketplaceCategoryID)
			if err != nil {
				continue // Continuar con otras categorías si una falla
			}
			categories[mapping.MarketplaceCategoryID] = category
		}
	}

	return categories, nil
}

// buildTenantCategoriesFromMappings construye categorías tenant simuladas desde los mapeos
func (uc *GetTenantTaxonomyUseCase) buildTenantCategoriesFromMappings(
	mappings []*entity.TenantCategoryMapping,
) map[string]*TenantCategory {
	categories := make(map[string]*TenantCategory)

	for _, mapping := range mappings {
		categories[mapping.CategoryID] = &TenantCategory{
			ID:        mapping.CategoryID,
			Name:      fmt.Sprintf("Category %s", mapping.CategoryID), // Placeholder
			Level:     0,                                              // TODO: Calcular nivel real
			ParentID:  nil,                                            // TODO: Obtener parent real
			IsActive:  true,
			CreatedAt: mapping.CreatedAt,
			UpdatedAt: mapping.UpdatedAt,
		}
	}

	return categories
}

// TenantCategory representa una categoría del tenant (simulada)
type TenantCategory struct {
	ID        string
	Name      string
	Level     int
	ParentID  *string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// buildCategoryNodes construye los nodos de categoría para la respuesta
func (uc *GetTenantTaxonomyUseCase) buildCategoryNodes(
	tenantCategories map[string]*TenantCategory,
	mappingsByCategory map[string]*entity.TenantCategoryMapping,
	marketplaceCategories map[string]*entity.MarketplaceCategory,
	customAttributes []*entity.TenantCustomAttribute,
	req *request.GetTenantTaxonomyRequest,
) []response.TenantCategoryNode {
	var nodes []response.TenantCategoryNode

	// Agrupar atributos personalizados por categoría
	attributesByCategory := make(map[string][]*entity.TenantCustomAttribute)
	for _, attr := range customAttributes {
		// Usar MarketplaceCategoryID como clave, o "global" si es nil
		categoryKey := "global"
		if attr.MarketplaceCategoryID != nil {
			categoryKey = *attr.MarketplaceCategoryID
		}
		attributesByCategory[categoryKey] = append(attributesByCategory[categoryKey], attr)
	}

	for _, tenantCategory := range tenantCategories {
		// Filtrar por categorías específicas si se especifican
		if len(req.CategoryIDs) > 0 && !uc.containsString(req.CategoryIDs, tenantCategory.ID) {
			continue
		}

		node := response.TenantCategoryNode{
			CategoryID: tenantCategory.ID,
			Name:       tenantCategory.Name,
			Level:      tenantCategory.Level,
			ParentID:   tenantCategory.ParentID,
			IsActive:   tenantCategory.IsActive,
			CreatedAt:  tenantCategory.CreatedAt,
			UpdatedAt:  tenantCategory.UpdatedAt,
		}

		// Agregar datos de mapeo si existe
		if mapping, exists := mappingsByCategory[tenantCategory.ID]; exists {
			node.MappingID = &mapping.ID
			node.MarketplaceCategoryID = &mapping.MarketplaceCategoryID
			node.CustomName = mapping.CustomName
			node.HasMapping = true

			// Agregar datos de categoría marketplace si se solicitan
			if req.IncludeMarketplaceData && marketplaceCategories != nil {
				if marketplaceCategory, exists := marketplaceCategories[mapping.MarketplaceCategoryID]; exists {
					node.MarketplaceData = &response.MarketplaceCategoryInfo{
						ID:          marketplaceCategory.ID,
						Name:        marketplaceCategory.Name,
						Slug:        marketplaceCategory.Slug,
						Description: &marketplaceCategory.Description,
						Level:       marketplaceCategory.Level,
						ParentID:    marketplaceCategory.ParentID,
						IsActive:    marketplaceCategory.IsActive,
						CreatedAt:   marketplaceCategory.CreatedAt,
						UpdatedAt:   marketplaceCategory.UpdatedAt,
					}
				}
			}
		}

		// Agregar atributos personalizados si se solicitan
		if req.IncludeCustomAttributes {
			if attrs, exists := attributesByCategory[tenantCategory.ID]; exists {
				node.CustomAttributes = uc.buildCustomAttributeInfosForCategory(attrs)
			}
		}

		nodes = append(nodes, node)
	}

	return nodes
}

// formatCategories formatea las categorías según el formato solicitado
func (uc *GetTenantTaxonomyUseCase) formatCategories(
	nodes []response.TenantCategoryNode,
	format string,
	maxDepth *int,
) []response.TenantCategoryNode {
	switch format {
	case "tree":
		return uc.buildTree(nodes, maxDepth)
	case "hierarchical":
		return uc.sortHierarchically(nodes)
	case "flat":
		fallthrough
	default:
		return nodes
	}
}

// buildTree construye un árbol de categorías
func (uc *GetTenantTaxonomyUseCase) buildTree(
	nodes []response.TenantCategoryNode,
	maxDepth *int,
) []response.TenantCategoryNode {
	// Crear mapa de nodos por ID
	nodeMap := make(map[string]*response.TenantCategoryNode)
	for i := range nodes {
		nodeMap[nodes[i].CategoryID] = &nodes[i]
	}

	var rootNodes []response.TenantCategoryNode

	// Construir árbol
	for i := range nodes {
		node := &nodes[i]

		// Verificar límite de profundidad
		if maxDepth != nil && node.Level >= *maxDepth {
			continue
		}

		if node.ParentID == nil {
			// Es un nodo raíz
			rootNodes = append(rootNodes, *node)
		} else {
			// Es un nodo hijo, agregarlo al parent
			if parent, exists := nodeMap[*node.ParentID]; exists {
				parent.Children = append(parent.Children, *node)
			}
		}
	}

	return rootNodes
}

// sortHierarchically ordena las categorías jerárquicamente
func (uc *GetTenantTaxonomyUseCase) sortHierarchically(
	nodes []response.TenantCategoryNode,
) []response.TenantCategoryNode {
	// TODO: Implementar ordenamiento jerárquico
	// Por ahora retornamos los nodos tal como están
	return nodes
}

// buildMetadata construye los metadatos de la taxonomía
func (uc *GetTenantTaxonomyUseCase) buildMetadata(
	tenantCategories map[string]*TenantCategory,
	mappings []*entity.TenantCategoryMapping,
	customAttributes []*entity.TenantCustomAttribute,
	req *request.GetTenantTaxonomyRequest,
) response.TaxonomyMetadata {
	maxDepth := 0
	rootCount := 0
	mappedCount := len(mappings)
	unmappedCount := len(tenantCategories) - mappedCount

	for _, category := range tenantCategories {
		if category.Level > maxDepth {
			maxDepth = category.Level
		}
		if category.ParentID == nil {
			rootCount++
		}
	}

	return response.TaxonomyMetadata{
		MaxDepth:                maxDepth,
		RootCategoriesCount:     rootCount,
		MappedCategoriesCount:   mappedCount,
		UnmappedCategoriesCount: unmappedCount,
		CustomAttributesCount:   len(customAttributes),
		IncludeOptions: response.TaxonomyIncludeOptions{
			CustomAttributes:   req.IncludeCustomAttributes,
			MarketplaceData:    req.IncludeMarketplaceData,
			InactiveCategories: req.IncludeInactive,
		},
	}
}

// buildCustomAttributeInfos construye la información de atributos personalizados
func (uc *GetTenantTaxonomyUseCase) buildCustomAttributeInfos(
	attributes []*entity.TenantCustomAttribute,
) []response.TenantCustomAttributeInfo {
	var infos []response.TenantCustomAttributeInfo

	for _, attr := range attributes {
		// Determinar categoryID basado en MarketplaceCategoryID
		categoryID := "global"
		if attr.MarketplaceCategoryID != nil {
			categoryID = *attr.MarketplaceCategoryID
		}

		info := response.TenantCustomAttributeInfo{
			ID:           attr.ID,
			CategoryID:   categoryID,
			Name:         attr.Name,
			Slug:         attr.Slug,
			Type:         attr.Type,
			IsRequired:   false, // TenantCustomAttribute no tiene este campo, usar default
			DefaultValue: nil,   // TenantCustomAttribute no tiene este campo, usar default
			CreatedAt:    attr.CreatedAt,
			UpdatedAt:    attr.UpdatedAt,
		}

		// Agregar opciones si es select o multi_select
		if attr.Type == "select" || attr.Type == "multi_select" {
			// TODO: Obtener opciones desde ValidationRules
			info.Options = []string{} // Placeholder
		}

		infos = append(infos, info)
	}

	return infos
}

// buildCustomAttributeInfosForCategory construye información de atributos para una categoría específica
func (uc *GetTenantTaxonomyUseCase) buildCustomAttributeInfosForCategory(
	attributes []*entity.TenantCustomAttribute,
) []response.TenantCustomAttributeInfo {
	return uc.buildCustomAttributeInfos(attributes)
}

// containsString verifica si un slice contiene un string específico
func (uc *GetTenantTaxonomyUseCase) containsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
