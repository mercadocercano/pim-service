package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	"saas-mt-pim-service/src/category/application/response"
	"saas-mt-pim-service/src/category/domain/port"
)

// MarketplaceProductPostgresRepository implementa queries cross-tenant de productos para el marketplace
type MarketplaceProductPostgresRepository struct {
	db *sql.DB
}

// NewMarketplaceProductPostgresRepository crea una nueva instancia del repositorio
func NewMarketplaceProductPostgresRepository(db *sql.DB) port.MarketplaceProductRepository {
	return &MarketplaceProductPostgresRepository{db: db}
}

// FindProductsByStoreType busca productos de tenants con un business_type específico
func (r *MarketplaceProductPostgresRepository) FindProductsByStoreType(
	ctx context.Context,
	storeTypeCode string,
	page, pageSize int,
) ([]*response.MarketplaceProductResponse, int, error) {
	offset := (page - 1) * pageSize

	// Contar total
	countQuery := `
		SELECT COUNT(DISTINCT p.id)
		FROM products p
		JOIN tenant_business_type_setup tbts ON tbts.tenant_id = p.tenant_id
		JOIN business_types bt ON bt.id = tbts.business_type_id
		WHERE bt.code = $1 AND p.status = 'active'
	`
	var total int
	if err := r.db.QueryRowContext(ctx, countQuery, storeTypeCode).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("error contando productos por store type: %w", err)
	}

	// Buscar productos con variante default
	query := `
		SELECT
			p.id, p.tenant_id, p.name, p.description,
			p.category_name, p.brand_name, p.image_url,
			bt.code, bt.name, bt.icon, bt.color,
			pv.id, pv.name, pv.sku, pv.price, pv.stock
		FROM products p
		JOIN tenant_business_type_setup tbts ON tbts.tenant_id = p.tenant_id
		JOIN business_types bt ON bt.id = tbts.business_type_id
		LEFT JOIN product_variants pv ON pv.product_id = p.id
			AND pv.is_default = true
			AND pv.status != 'deleted'
		WHERE bt.code = $1 AND p.status = 'active'
		ORDER BY p.created_at DESC
		LIMIT $2 OFFSET $3
	`

	return r.scanProducts(ctx, query, storeTypeCode, pageSize, offset, total)
}

// FindAllProducts busca todos los productos activos cross-tenant
func (r *MarketplaceProductPostgresRepository) FindAllProducts(
	ctx context.Context,
	search string,
	page, pageSize int,
) ([]*response.MarketplaceProductResponse, int, error) {
	offset := (page - 1) * pageSize

	var total int
	var countQuery string
	var dataQuery string
	var countArgs []interface{}
	var dataArgs []interface{}

	if search != "" {
		searchPattern := "%" + search + "%"

		countQuery = `
			SELECT COUNT(DISTINCT p.id)
			FROM products p
			WHERE p.status = 'active' AND (p.name ILIKE $1 OR p.category_name ILIKE $1 OR p.brand_name ILIKE $1)
		`
		countArgs = []interface{}{searchPattern}

		dataQuery = `
			SELECT
				p.id, p.tenant_id, p.name, p.description,
				p.category_name, p.brand_name, p.image_url,
				bt.code, bt.name, bt.icon, bt.color,
				pv.id, pv.name, pv.sku, pv.price, pv.stock
			FROM products p
			LEFT JOIN tenant_business_type_setup tbts ON tbts.tenant_id = p.tenant_id
			LEFT JOIN business_types bt ON bt.id = tbts.business_type_id
			LEFT JOIN product_variants pv ON pv.product_id = p.id
				AND pv.is_default = true
				AND pv.status != 'deleted'
			WHERE p.status = 'active' AND (p.name ILIKE $1 OR p.category_name ILIKE $1 OR p.brand_name ILIKE $1)
			ORDER BY p.created_at DESC
			LIMIT $2 OFFSET $3
		`
		dataArgs = []interface{}{searchPattern, pageSize, offset}
	} else {
		countQuery = `SELECT COUNT(*) FROM products WHERE status = 'active'`
		countArgs = nil

		dataQuery = `
			SELECT
				p.id, p.tenant_id, p.name, p.description,
				p.category_name, p.brand_name, p.image_url,
				bt.code, bt.name, bt.icon, bt.color,
				pv.id, pv.name, pv.sku, pv.price, pv.stock
			FROM products p
			LEFT JOIN tenant_business_type_setup tbts ON tbts.tenant_id = p.tenant_id
			LEFT JOIN business_types bt ON bt.id = tbts.business_type_id
			LEFT JOIN product_variants pv ON pv.product_id = p.id
				AND pv.is_default = true
				AND pv.status != 'deleted'
			WHERE p.status = 'active'
			ORDER BY p.created_at DESC
			LIMIT $1 OFFSET $2
		`
		dataArgs = []interface{}{pageSize, offset}
	}

	if err := r.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("error contando productos: %w", err)
	}

	rows, err := r.db.QueryContext(ctx, dataQuery, dataArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("error buscando productos: %w", err)
	}
	defer rows.Close()

	products := make([]*response.MarketplaceProductResponse, 0)
	for rows.Next() {
		p, err := r.scanProductRow(rows)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	return products, total, nil
}

// FindProductByID busca un producto por ID (cross-tenant)
func (r *MarketplaceProductPostgresRepository) FindProductByID(
	ctx context.Context,
	productID string,
) (*response.MarketplaceProductResponse, error) {
	query := `
		SELECT
			p.id, p.tenant_id, p.name, p.description,
			p.category_name, p.brand_name, p.image_url,
			bt.code, bt.name, bt.icon, bt.color,
			pv.id, pv.name, pv.sku, pv.price, pv.stock
		FROM products p
		LEFT JOIN tenant_business_type_setup tbts ON tbts.tenant_id = p.tenant_id
		LEFT JOIN business_types bt ON bt.id = tbts.business_type_id
		LEFT JOIN product_variants pv ON pv.product_id = p.id
			AND pv.is_default = true
			AND pv.status != 'deleted'
		WHERE p.id = $1 AND p.status = 'active'
		LIMIT 1
	`

	rows, err := r.db.QueryContext(ctx, query, productID)
	if err != nil {
		return nil, fmt.Errorf("error buscando producto por ID: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	return r.scanProductRow(rows)
}

// FindProductsByTenantID busca productos de un tenant específico
func (r *MarketplaceProductPostgresRepository) FindProductsByTenantID(
	ctx context.Context,
	tenantID string,
	page, pageSize int,
) ([]*response.MarketplaceProductResponse, int, error) {
	offset := (page - 1) * pageSize

	var total int
	countQuery := `SELECT COUNT(*) FROM products WHERE tenant_id = $1 AND status = 'active'`
	if err := r.db.QueryRowContext(ctx, countQuery, tenantID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("error contando productos del tenant: %w", err)
	}

	query := `
		SELECT
			p.id, p.tenant_id, p.name, p.description,
			p.category_name, p.brand_name, p.image_url,
			bt.code, bt.name, bt.icon, bt.color,
			pv.id, pv.name, pv.sku, pv.price, pv.stock
		FROM products p
		LEFT JOIN tenant_business_type_setup tbts ON tbts.tenant_id = p.tenant_id
		LEFT JOIN business_types bt ON bt.id = tbts.business_type_id
		LEFT JOIN product_variants pv ON pv.product_id = p.id
			AND pv.is_default = true
			AND pv.status != 'deleted'
		WHERE p.tenant_id = $1 AND p.status = 'active'
		ORDER BY p.created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error buscando productos del tenant: %w", err)
	}
	defer rows.Close()

	products := make([]*response.MarketplaceProductResponse, 0)
	for rows.Next() {
		p, err := r.scanProductRow(rows)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	return products, total, nil
}

// GetStoreTypesWithCounts lista business_types con conteo de tiendas y productos
func (r *MarketplaceProductPostgresRepository) GetStoreTypesWithCounts(ctx context.Context) ([]*response.MarketplaceStoreTypeResponse, error) {
	// Solo muestra business_types que tienen templates (los del onboarding)
	query := `
		SELECT
			bt.code, bt.name,
			COALESCE(bt.icon, '') as icon,
			COALESCE(bt.color, '') as color,
			COUNT(DISTINCT tbts.tenant_id) as store_count,
			COUNT(DISTINCT p.id) as product_count
		FROM business_types bt
		INNER JOIN business_type_templates btt ON btt.business_type_id = bt.id AND btt.is_active = true
		LEFT JOIN tenant_business_type_setup tbts ON tbts.business_type_id = bt.id
		LEFT JOIN products p ON p.tenant_id = tbts.tenant_id AND p.status = 'active'
		WHERE bt.is_active = true
		GROUP BY bt.code, bt.name, bt.icon, bt.color, bt.sort_order
		ORDER BY bt.sort_order ASC, bt.name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo store types: %w", err)
	}
	defer rows.Close()

	storeTypes := make([]*response.MarketplaceStoreTypeResponse, 0)
	for rows.Next() {
		st := &response.MarketplaceStoreTypeResponse{}
		if err := rows.Scan(&st.Code, &st.Name, &st.Icon, &st.Color, &st.StoreCount, &st.ProductCount); err != nil {
			return nil, fmt.Errorf("error escaneando store type: %w", err)
		}
		storeTypes = append(storeTypes, st)
	}

	return storeTypes, nil
}

// scanProducts ejecuta la query y escanea los resultados
func (r *MarketplaceProductPostgresRepository) scanProducts(
	ctx context.Context,
	query string,
	storeTypeCode string,
	pageSize, offset, total int,
) ([]*response.MarketplaceProductResponse, int, error) {
	rows, err := r.db.QueryContext(ctx, query, storeTypeCode, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error buscando productos: %w", err)
	}
	defer rows.Close()

	products := make([]*response.MarketplaceProductResponse, 0)
	for rows.Next() {
		p, err := r.scanProductRow(rows)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	return products, total, nil
}

// scanProductRow escanea una fila de producto con variante
func (r *MarketplaceProductPostgresRepository) scanProductRow(rows *sql.Rows) (*response.MarketplaceProductResponse, error) {
	p := &response.MarketplaceProductResponse{}

	var description, categoryName, brandName, imageURL sql.NullString
	var btCode, btName, btIcon, btColor sql.NullString
	var variantID, variantName, variantSKU sql.NullString
	var variantPrice sql.NullFloat64
	var variantStock sql.NullInt64

	err := rows.Scan(
		&p.ID, &p.TenantID, &p.Name, &description,
		&categoryName, &brandName, &imageURL,
		&btCode, &btName, &btIcon, &btColor,
		&variantID, &variantName, &variantSKU, &variantPrice, &variantStock,
	)
	if err != nil {
		return nil, fmt.Errorf("error escaneando producto: %w", err)
	}

	if description.Valid {
		p.Description = &description.String
	}
	if categoryName.Valid {
		p.CategoryName = &categoryName.String
	}
	if brandName.Valid {
		p.BrandName = &brandName.String
	}
	if imageURL.Valid {
		p.ImageURL = &imageURL.String
	}

	if btCode.Valid {
		p.StoreType = &response.MarketplaceStoreTypeInfo{
			Code:  btCode.String,
			Name:  btName.String,
			Icon:  btIcon.String,
			Color: btColor.String,
		}
	}

	if variantID.Valid {
		var sku *string
		if variantSKU.Valid {
			sku = &variantSKU.String
		}
		p.Variant = &response.MarketplaceProductVariantInfo{
			ID:    variantID.String,
			Name:  variantName.String,
			SKU:   sku,
			Price: variantPrice.Float64,
			Stock: int(variantStock.Int64),
		}
	}

	return p, nil
}

// CalculateTotalPages calcula el total de páginas
func CalculateTotalPages(total, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	return int(math.Ceil(float64(total) / float64(pageSize)))
}
