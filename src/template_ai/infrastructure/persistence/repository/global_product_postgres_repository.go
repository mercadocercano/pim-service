package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
)

// GlobalProductPostgresRepository implements the GlobalProductRepository interface
type GlobalProductPostgresRepository struct {
	db *sql.DB
}

// NewGlobalProductPostgresRepository creates a new instance of the repository
func NewGlobalProductPostgresRepository(db *sql.DB) *GlobalProductPostgresRepository {
	return &GlobalProductPostgresRepository{db: db}
}

// FindByIDs finds global products by their IDs
func (r *GlobalProductPostgresRepository) FindByIDs(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]interface{}, error) {
	if len(ids) == 0 {
		return make(map[uuid.UUID]interface{}), nil
	}

	// This is a simplified implementation
	// In a real implementation, this would query the global_products table
	// from the global catalog schema
	
	query := `
		SELECT id, name, description, ean, category_id, brand_id, 
			average_price, created_at, updated_at
		FROM global_products
		WHERE id = ANY($1)
	`

	// Convert UUIDs to string array for PostgreSQL
	idStrings := make([]string, len(ids))
	for i, id := range ids {
		idStrings[i] = id.String()
	}

	rows, err := r.db.QueryContext(ctx, query, idStrings)
	if err != nil {
		return nil, fmt.Errorf("failed to query global products: %w", err)
	}
	defer rows.Close()

	result := make(map[uuid.UUID]interface{})
	for rows.Next() {
		var product struct {
			ID           uuid.UUID
			Name         string
			Description  sql.NullString
			EAN          sql.NullString
			CategoryID   sql.NullString
			BrandID      sql.NullString
			AveragePrice sql.NullFloat64
		}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.EAN,
			&product.CategoryID,
			&product.BrandID,
			&product.AveragePrice,
		)
		if err != nil {
			continue
		}

		// Convert to generic interface for flexibility
		productMap := map[string]interface{}{
			"id":          product.ID,
			"name":        product.Name,
			"description": product.Description.String,
			"ean":         product.EAN.String,
			"category_id": product.CategoryID.String,
			"brand_id":    product.BrandID.String,
		}

		if product.AveragePrice.Valid {
			productMap["average_price"] = product.AveragePrice.Float64
		}

		result[product.ID] = productMap
	}

	return result, nil
}

// SearchByCategory searches global products by category
func (r *GlobalProductPostgresRepository) SearchByCategory(ctx context.Context, categoryID string, limit int) ([]interface{}, error) {
	query := `
		SELECT id, name, description, ean, category_id, brand_id, 
			average_price, quality_score
		FROM global_products
		WHERE category_id = $1
		ORDER BY quality_score DESC, average_price ASC
		LIMIT $2
	`

	rows, err := r.db.QueryContext(ctx, query, categoryID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to search products by category: %w", err)
	}
	defer rows.Close()

	var results []interface{}
	for rows.Next() {
		var product struct {
			ID           uuid.UUID
			Name         string
			Description  sql.NullString
			EAN          sql.NullString
			CategoryID   string
			BrandID      sql.NullString
			AveragePrice sql.NullFloat64
			QualityScore sql.NullFloat64
		}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.EAN,
			&product.CategoryID,
			&product.BrandID,
			&product.AveragePrice,
			&product.QualityScore,
		)
		if err != nil {
			continue
		}

		productMap := map[string]interface{}{
			"id":          product.ID,
			"name":        product.Name,
			"description": product.Description.String,
			"ean":         product.EAN.String,
			"category_id": product.CategoryID,
			"brand_id":    product.BrandID.String,
		}

		if product.AveragePrice.Valid {
			productMap["average_price"] = product.AveragePrice.Float64
		}
		if product.QualityScore.Valid {
			productMap["quality_score"] = product.QualityScore.Float64
		}

		results = append(results, productMap)
	}

	return results, nil
}

// GetFilteredProducts retrieves global products based on filters
func (r *GlobalProductPostgresRepository) GetFilteredProducts(ctx context.Context, filters map[string]interface{}) ([]*entity.GlobalProduct, error) {
	query := `
		SELECT id, name, category_id, category_name, brand_id, brand_name, 
			average_price as price, is_bulk, is_perishable, tags, attributes
		FROM global_products
		WHERE 1=1
	`
	
	args := []interface{}{}
	argCount := 0
	
	// Build dynamic query based on filters
	if businessType, ok := filters["business_type"].(string); ok && businessType != "" {
		argCount++
		query += fmt.Sprintf(" AND business_type = $%d", argCount)
		args = append(args, businessType)
	}
	
	if isActive, ok := filters["is_active"].(bool); ok {
		argCount++
		query += fmt.Sprintf(" AND is_active = $%d", argCount)
		args = append(args, isActive)
	}
	
	if categories, ok := filters["categories"].([]string); ok && len(categories) > 0 {
		argCount++
		query += fmt.Sprintf(" AND category_id = ANY($%d)", argCount)
		args = append(args, categories)
	}
	
	if excludeBrands, ok := filters["exclude_brands"].([]string); ok && len(excludeBrands) > 0 {
		argCount++
		query += fmt.Sprintf(" AND brand_id != ALL($%d)", argCount)
		args = append(args, excludeBrands)
	}
	
	query += " ORDER BY quality_score DESC, created_at DESC"
	
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get filtered products: %w", err)
	}
	defer rows.Close()
	
	var products []*entity.GlobalProduct
	for rows.Next() {
		var product entity.GlobalProduct
		var tags sql.NullString
		var attributes sql.NullString
		
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.CategoryID,
			&product.CategoryName,
			&product.BrandID,
			&product.BrandName,
			&product.Price,
			&product.IsBulk,
			&product.IsPerishable,
			&tags,
			&attributes,
		)
		if err != nil {
			continue
		}
		
		// Parse JSON fields
		if tags.Valid {
			// In a real implementation, we would parse JSON array
			product.Tags = []string{}
		}
		if attributes.Valid {
			// In a real implementation, we would parse JSON object
			product.Attributes = make(map[string]interface{})
		}
		
		products = append(products, &product)
	}
	
	return products, nil
}

// SearchByBrand searches global products by brand
func (r *GlobalProductPostgresRepository) SearchByBrand(ctx context.Context, brandID string, limit int) ([]interface{}, error) {
	query := `
		SELECT id, name, description, ean, category_id, brand_id, 
			average_price, quality_score
		FROM global_products
		WHERE brand_id = $1
		ORDER BY quality_score DESC, average_price ASC
		LIMIT $2
	`

	rows, err := r.db.QueryContext(ctx, query, brandID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to search products by brand: %w", err)
	}
	defer rows.Close()

	var results []interface{}
	for rows.Next() {
		var product struct {
			ID           uuid.UUID
			Name         string
			Description  sql.NullString
			EAN          sql.NullString
			CategoryID   sql.NullString
			BrandID      string
			AveragePrice sql.NullFloat64
			QualityScore sql.NullFloat64
		}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.EAN,
			&product.CategoryID,
			&product.BrandID,
			&product.AveragePrice,
			&product.QualityScore,
		)
		if err != nil {
			continue
		}

		productMap := map[string]interface{}{
			"id":          product.ID,
			"name":        product.Name,
			"description": product.Description.String,
			"ean":         product.EAN.String,
			"category_id": product.CategoryID.String,
			"brand_id":    product.BrandID,
		}

		if product.AveragePrice.Valid {
			productMap["average_price"] = product.AveragePrice.Float64
		}
		if product.QualityScore.Valid {
			productMap["quality_score"] = product.QualityScore.Float64
		}

		results = append(results, productMap)
	}

	return results, nil
}
