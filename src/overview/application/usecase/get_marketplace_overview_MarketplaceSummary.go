package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"saas-mt-pim-service/src/overview/application/request"
	"saas-mt-pim-service/src/overview/application/response"

	// Importar las interfaces de repositorios
	attributePort "saas-mt-pim-service/src/attribute/domain/port"
	brandPort "saas-mt-pim-service/src/brand/domain/port"
	categoryPort "saas-mt-pim-service/src/category/domain/port"
	globalProductPort "saas-mt-pim-service/src/product/global_catalog/domain/port"
	cr "github.com/mercadocercano/criteria"
)

// GetMarketplaceOverviewUseCase implementa el caso de uso para obtener overview del marketplace
type GetMarketplaceOverviewUseCase struct {
	// Repositorios con interfaces reales
	marketplaceCategoryRepo  categoryPort.MarketplaceCategoryRepository
	marketplaceBrandRepo     brandPort.MarketplacebrandRepository
	marketplaceAttributeRepo attributePort.MarketplaceAttributeRepository
	globalProductRepo        globalProductPort.GlobalProductRepository
}

// NewGetMarketplaceOverviewUseCase crea una nueva instancia del caso de uso
func NewGetMarketplaceOverviewUseCase(
	marketplaceCategoryRepo categoryPort.MarketplaceCategoryRepository,
	marketplaceBrandRepo brandPort.MarketplacebrandRepository,
	marketplaceAttributeRepo attributePort.MarketplaceAttributeRepository,
	globalProductRepo globalProductPort.GlobalProductRepository,
) *GetMarketplaceOverviewUseCase {
	return &GetMarketplaceOverviewUseCase{
		marketplaceCategoryRepo:  marketplaceCategoryRepo,
		marketplaceBrandRepo:     marketplaceBrandRepo,
		marketplaceAttributeRepo: marketplaceAttributeRepo,
		globalProductRepo:        globalProductRepo,
	}
}

// Execute ejecuta el caso de uso para obtener overview del marketplace
func (uc *GetMarketplaceOverviewUseCase) Execute(
	ctx context.Context,
	req *request.GetMarketplaceOverviewRequest,
) (*response.GetMarketplaceOverviewResponse, error) {
	// Validar request
	if err := uc.validateRequest(req); err != nil {
		return nil, fmt.Errorf("validación fallida: %w", err)
	}

	startTime := time.Now()
	resp := &response.GetMarketplaceOverviewResponse{
		Success:     true,
		GeneratedAt: startTime,
		Data:        make(map[string]interface{}),
		Metadata: response.OverviewMetadata{
			ParallelExecution: req.Parallel,
			Sections:          req.Sections,
		},
	}

	// Ejecutar estrategias según las secciones solicitadas
	if req.Parallel {
		err := uc.executeParallel(ctx, req, resp)
		if err != nil {
			return nil, err
		}
	} else {
		err := uc.executeSequential(ctx, req, resp)
		if err != nil {
			return nil, err
		}
	}

	resp.Metadata.TotalTime = time.Since(startTime)
	return resp, nil
}

// validateRequest valida los parámetros del request
func (uc *GetMarketplaceOverviewUseCase) validateRequest(req *request.GetMarketplaceOverviewRequest) error {
	if len(req.Sections) == 0 {
		return fmt.Errorf("debe especificarse al menos una sección")
	}

	validSections := map[string]bool{
		"dashboard":      true,
		"taxonomy":       true,
		"brands":         true,
		"global-catalog": true,
		"attributes":     true,
		"curation":       true,
	}

	for _, section := range req.Sections {
		if !validSections[section] {
			return fmt.Errorf("sección inválida: %s", section)
		}
	}

	if req.TimeRangeDays < 1 || req.TimeRangeDays > 365 {
		req.TimeRangeDays = 7 // Default
	}

	if req.Limit < 1 || req.Limit > 1000 {
		req.Limit = 10 // Default
	}

	return nil
}

// executeParallel ejecuta las consultas en paralelo
func (uc *GetMarketplaceOverviewUseCase) executeParallel(
	ctx context.Context,
	req *request.GetMarketplaceOverviewRequest,
	resp *response.GetMarketplaceOverviewResponse,
) error {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	errors := make([]error, 0)

	for _, section := range req.Sections {
		wg.Add(1)
		go func(sectionName string) {
			defer wg.Done()

			data, err := uc.executeSection(ctx, sectionName, req)

			mutex.Lock()
			defer mutex.Unlock()

			if err != nil {
				errors = append(errors, fmt.Errorf("error en sección %s: %w", sectionName, err))
			} else {
				resp.Data[sectionName] = data
			}
		}(section)
	}

	wg.Wait()

	if len(errors) > 0 {
		return fmt.Errorf("errores en ejecución paralela: %v", errors)
	}

	return nil
}

// executeSequential ejecuta las consultas secuencialmente
func (uc *GetMarketplaceOverviewUseCase) executeSequential(
	ctx context.Context,
	req *request.GetMarketplaceOverviewRequest,
	resp *response.GetMarketplaceOverviewResponse,
) error {
	for _, section := range req.Sections {
		data, err := uc.executeSection(ctx, section, req)
		if err != nil {
			return fmt.Errorf("error en sección %s: %w", section, err)
		}
		resp.Data[section] = data
	}
	return nil
}

// executeSection ejecuta la estrategia específica para cada sección
func (uc *GetMarketplaceOverviewUseCase) executeSection(
	ctx context.Context,
	section string,
	req *request.GetMarketplaceOverviewRequest,
) (interface{}, error) {
	switch section {
	case "dashboard":
		return uc.getDashboardStats(ctx, req)
	case "taxonomy":
		return uc.getTaxonomyStats(ctx, req)
	case "brands":
		return uc.getBrandsStats(ctx, req)
	case "global-catalog":
		return uc.getGlobalCatalogStats(ctx, req)
	case "attributes":
		return uc.getAttributesStats(ctx, req)
	case "curation":
		return uc.getCurationStats(ctx, req)
	default:
		return nil, fmt.Errorf("sección no implementada: %s", section)
	}
}

// getDashboardStats obtiene estadísticas para el dashboard principal
func (uc *GetMarketplaceOverviewUseCase) getDashboardStats(
	ctx context.Context,
	req *request.GetMarketplaceOverviewRequest,
) (map[string]interface{}, error) {
	stats := map[string]interface{}{
		"total_categories":      0,
		"total_brands":          0,
		"total_attributes":      0,
		"total_global_products": 0,
		"active_tenants":        0,
		"recent_activity":       []interface{}{},
	}

	// Consultar datos reales desde los repositorios
	emptyCriteria := cr.Criteria{} // Sin filtros para contar todos

	// Contar categorías de marketplace
	totalCategories, err := uc.marketplaceCategoryRepo.CountByCriteria(ctx, emptyCriteria)
	if err != nil {
		totalCategories = 0 // Fallback en caso de error
	}
	stats["total_categories"] = totalCategories

	// Contar marcas de marketplace
	totalBrands, err := uc.marketplaceBrandRepo.CountByCriteria(ctx, emptyCriteria)
	if err != nil {
		totalBrands = 0 // Fallback en caso de error
	}
	stats["total_brands"] = totalBrands

	// Contar atributos de marketplace
	totalAttributes, err := uc.marketplaceAttributeRepo.CountByCriteria(ctx, emptyCriteria)
	if err != nil {
		totalAttributes = 0 // Fallback en caso de error
	}
	stats["total_attributes"] = totalAttributes

	// Contar productos globales
	totalProducts, err := uc.globalProductRepo.CountTotal()
	if err != nil {
		totalProducts = 0 // Fallback en caso de error
	}
	stats["total_global_products"] = totalProducts

	// Por ahora active_tenants queda como valor por defecto
	// TODO: Implementar cuando tengamos acceso al repositorio de tenants
	stats["active_tenants"] = 0

	return stats, nil
}

// getTaxonomyStats obtiene estadísticas de taxonomía
func (uc *GetMarketplaceOverviewUseCase) getTaxonomyStats(
	ctx context.Context,
	req *request.GetMarketplaceOverviewRequest,
) (map[string]interface{}, error) {
	stats := map[string]interface{}{
		"total_categories":     0,
		"active_categories":    0,
		"root_categories":      0,
		"avg_depth":            0,
		"most_used_categories": []interface{}{},
	}

	// TODO: Implementar consultas reales
	stats["total_categories"] = 150
	stats["active_categories"] = 145
	stats["root_categories"] = 12
	stats["avg_depth"] = 3.2

	return stats, nil
}

// getBrandsStats obtiene estadísticas de marcas
func (uc *GetMarketplaceOverviewUseCase) getBrandsStats(
	ctx context.Context,
	req *request.GetMarketplaceOverviewRequest,
) (map[string]interface{}, error) {
	stats := map[string]interface{}{
		"total_brands":    0,
		"verified_brands": 0,
		"pending_brands":  0,
		"top_brands":      []interface{}{},
		"recent_brands":   []interface{}{},
	}

	// Consultar datos reales
	emptyCriteria := cr.Criteria{}

	// Contar todas las marcas
	totalBrands, err := uc.marketplaceBrandRepo.CountByCriteria(ctx, emptyCriteria)
	if err != nil {
		totalBrands = 0
	}
	stats["total_brands"] = totalBrands

	// Contar marcas verificadas
	verifiedCriteria := cr.Criteria{
		Filters: cr.NewFilters(
			cr.NewFilter("verification_status", cr.OpEqual, "verified"),
		),
	}
	verifiedBrands, err := uc.marketplaceBrandRepo.CountByCriteria(ctx, verifiedCriteria)
	if err != nil {
		verifiedBrands = 0
	}
	stats["verified_brands"] = verifiedBrands

	// Contar marcas pendientes
	pendingCriteria := cr.Criteria{
		Filters: cr.NewFilters(
			cr.NewFilter("verification_status", cr.OpEqual, "pending"),
		),
	}
	pendingBrands, err := uc.marketplaceBrandRepo.CountByCriteria(ctx, pendingCriteria)
	if err != nil {
		pendingBrands = 0
	}
	stats["pending_brands"] = pendingBrands

	return stats, nil
}

// getGlobalCatalogStats obtiene estadísticas del catálogo global
func (uc *GetMarketplaceOverviewUseCase) getGlobalCatalogStats(
	ctx context.Context,
	req *request.GetMarketplaceOverviewRequest,
) (map[string]interface{}, error) {
	stats := map[string]interface{}{
		"total_products":     0,
		"verified_products":  0,
		"argentine_products": 0,
		"avg_quality_score":  0,
		"top_categories":     []interface{}{},
		"recent_products":    []interface{}{},
	}

	// Consultar datos reales
	totalProducts, err := uc.globalProductRepo.CountTotal()
	if err != nil {
		totalProducts = 0
	}
	stats["total_products"] = totalProducts

	// Contar productos argentinos
	argentineProducts, err := uc.globalProductRepo.CountArgentineProducts()
	if err != nil {
		argentineProducts = 0
	}
	stats["argentine_products"] = argentineProducts

	// Contar productos de alta calidad (quality_score >= 70)
	highQualityProducts, err := uc.globalProductRepo.CountByQualityScore(70)
	if err != nil {
		highQualityProducts = 0
	}
	stats["verified_products"] = highQualityProducts

	// Por ahora avg_quality_score se mantiene como 0 hasta implementar la consulta específica
	stats["avg_quality_score"] = 0

	return stats, nil
}

// getAttributesStats obtiene estadísticas de atributos
func (uc *GetMarketplaceOverviewUseCase) getAttributesStats(
	ctx context.Context,
	req *request.GetMarketplaceOverviewRequest,
) (map[string]interface{}, error) {
	stats := map[string]interface{}{
		"total_attributes":    0,
		"active_attributes":   0,
		"required_attributes": 0,
		"by_type":             map[string]int{},
		"most_used":           []interface{}{},
	}

	// Consultar datos reales
	emptyCriteria := cr.Criteria{}

	// Contar total de atributos
	totalAttributes, err := uc.marketplaceAttributeRepo.CountByCriteria(ctx, emptyCriteria)
	if err != nil {
		totalAttributes = 0
	}
	stats["total_attributes"] = totalAttributes

	// Contar atributos requeridos para listado
	requiredCriteria := cr.Criteria{
		Filters: cr.NewFilters(
			cr.NewFilter("is_required_for_listing", cr.OpEqual, true),
		),
	}
	requiredAttributes, err := uc.marketplaceAttributeRepo.CountByCriteria(ctx, requiredCriteria)
	if err != nil {
		requiredAttributes = 0
	}
	stats["required_attributes"] = requiredAttributes

	// Por simplicidad, active_attributes será igual al total
	// TODO: Implementar si hay un campo específico para atributos activos
	stats["active_attributes"] = totalAttributes

	// TODO: Implementar conteo por tipo cuando tengamos acceso a consultas más específicas
	stats["by_type"] = map[string]int{}

	return stats, nil
}

// getCurationStats obtiene estadísticas de curación desde Catalog BFF
func (uc *GetMarketplaceOverviewUseCase) getCurationStats(
	ctx context.Context,
	req *request.GetMarketplaceOverviewRequest,
) (map[string]interface{}, error) {
	stats := map[string]interface{}{
		"pending":        0,
		"approved_today": 0,
		"rejected_today": 0,
		"total_scraped":  0,
	}

	// Llamar al Catalog BFF para obtener stats de curación
	catalogBFFURL := "http://catalog-bff-service:8085" // URL interna de Docker
	dashboardURL := fmt.Sprintf("%s/api/v1/admin/dashboard/stats", catalogBFFURL)
	
	httpReq, err := http.NewRequestWithContext(ctx, "GET", dashboardURL, nil)
	if err != nil {
		// Retornar stats vacíos en caso de error
		return stats, nil
	}
	
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		// Retornar stats vacíos si BFF no está disponible
		return stats, nil
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		// Retornar stats vacíos si hubo error
		return stats, nil
	}
	
	var bffResponse struct {
		Curation struct {
			Pending       int `json:"pending"`
			ApprovedToday int `json:"approved_today"`
			RejectedToday int `json:"rejected_today"`
			TotalScraped  int `json:"total_scraped"`
		} `json:"curation"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&bffResponse); err != nil {
		// Retornar stats vacíos si no se puede parsear
		return stats, nil
	}
	
	stats["pending"] = bffResponse.Curation.Pending
	stats["approved_today"] = bffResponse.Curation.ApprovedToday
	stats["rejected_today"] = bffResponse.Curation.RejectedToday
	stats["total_scraped"] = bffResponse.Curation.TotalScraped

	return stats, nil
}
