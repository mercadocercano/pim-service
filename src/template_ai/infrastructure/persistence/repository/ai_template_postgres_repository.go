package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"github.com/lib/pq"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/exception"
)

// AITemplatePostgresRepository implements the AITemplateRepository interface using PostgreSQL
type AITemplatePostgresRepository struct {
	db *sql.DB
}

// NewAITemplatePostgresRepository creates a new instance of the repository
func NewAITemplatePostgresRepository(db *sql.DB) *AITemplatePostgresRepository {
	return &AITemplatePostgresRepository{db: db}
}

// Save saves a new AI template
func (r *AITemplatePostgresRepository) Save(ctx context.Context, template *entity.AITemplate) error {
	query := `
		INSERT INTO business_type_templates (
			id, business_type_id, tenant_id, name, description, 
			generated_by, ai_generation_params, performance_metrics,
			last_ai_update, product_selection_rules, category_distribution,
			brand_preferences, regional_adaptations, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`

	aiGenParams, _ := json.Marshal(template.AIGenerationParams)
	perfMetrics, _ := json.Marshal(template.PerformanceMetrics)
	prodRules, _ := json.Marshal(template.ProductSelectionRules)
	catDist, _ := json.Marshal(template.CategoryDistribution)
	brandPrefs, _ := json.Marshal(template.BrandPreferences)
	regionalAdapt, _ := json.Marshal(template.RegionalAdaptations)

	_, err := r.db.ExecContext(ctx, query,
		template.ID,
		template.BusinessTypeID,
		template.TenantID,
		template.Name,
		template.Description,
		template.GeneratedBy,
		aiGenParams,
		perfMetrics,
		template.LastAIUpdate,
		prodRules,
		catDist,
		brandPrefs,
		regionalAdapt,
		template.CreatedAt,
		template.UpdatedAt,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return exception.ErrTemplateAlreadyExists
		}
		return fmt.Errorf("failed to save template: %w", err)
	}

	return nil
}

// Update updates an existing AI template
func (r *AITemplatePostgresRepository) Update(ctx context.Context, template *entity.AITemplate) error {
	query := `
		UPDATE business_type_templates SET
			name = $2, description = $3, generated_by = $4,
			ai_generation_params = $5, performance_metrics = $6,
			last_ai_update = $7, product_selection_rules = $8,
			category_distribution = $9, brand_preferences = $10,
			regional_adaptations = $11, updated_at = $12
		WHERE id = $1
	`

	aiGenParams, _ := json.Marshal(template.AIGenerationParams)
	perfMetrics, _ := json.Marshal(template.PerformanceMetrics)
	prodRules, _ := json.Marshal(template.ProductSelectionRules)
	catDist, _ := json.Marshal(template.CategoryDistribution)
	brandPrefs, _ := json.Marshal(template.BrandPreferences)
	regionalAdapt, _ := json.Marshal(template.RegionalAdaptations)

	result, err := r.db.ExecContext(ctx, query,
		template.ID,
		template.Name,
		template.Description,
		template.GeneratedBy,
		aiGenParams,
		perfMetrics,
		template.LastAIUpdate,
		prodRules,
		catDist,
		brandPrefs,
		regionalAdapt,
		template.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update template: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrTemplateNotFound
	}

	return nil
}

// FindByID finds a template by ID
func (r *AITemplatePostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.AITemplate, error) {
	query := `
		SELECT id, business_type_id, tenant_id, name, description,
			generated_by, ai_generation_params, performance_metrics,
			last_ai_update, product_selection_rules, category_distribution,
			brand_preferences, regional_adaptations, created_at, updated_at
		FROM business_type_templates
		WHERE id = $1
	`

	template := &entity.AITemplate{}
	var aiGenParams, perfMetrics, prodRules, catDist, brandPrefs, regionalAdapt []byte

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&template.ID,
		&template.BusinessTypeID,
		&template.TenantID,
		&template.Name,
		&template.Description,
		&template.GeneratedBy,
		&aiGenParams,
		&perfMetrics,
		&template.LastAIUpdate,
		&prodRules,
		&catDist,
		&brandPrefs,
		&regionalAdapt,
		&template.CreatedAt,
		&template.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, exception.ErrTemplateNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find template: %w", err)
	}

	// Unmarshal JSON fields
	json.Unmarshal(aiGenParams, &template.AIGenerationParams)
	json.Unmarshal(perfMetrics, &template.PerformanceMetrics)
	json.Unmarshal(prodRules, &template.ProductSelectionRules)
	json.Unmarshal(catDist, &template.CategoryDistribution)
	json.Unmarshal(brandPrefs, &template.BrandPreferences)
	json.Unmarshal(regionalAdapt, &template.RegionalAdaptations)

	return template, nil
}

// FindByBusinessTypeAndTenant finds templates by business type and optional tenant
func (r *AITemplatePostgresRepository) FindByBusinessTypeAndTenant(
	ctx context.Context,
	businessTypeID uuid.UUID,
	tenantID *uuid.UUID,
) ([]*entity.AITemplate, error) {
	var query string
	var args []interface{}

	if tenantID != nil {
		query = `
			SELECT id, business_type_id, tenant_id, name, description,
				generated_by, ai_generation_params, performance_metrics,
				last_ai_update, product_selection_rules, category_distribution,
				brand_preferences, regional_adaptations, created_at, updated_at
			FROM business_type_templates
			WHERE business_type_id = $1 AND (tenant_id = $2 OR tenant_id IS NULL)
			ORDER BY tenant_id NULLS LAST, created_at DESC
		`
		args = []interface{}{businessTypeID, tenantID}
	} else {
		query = `
			SELECT id, business_type_id, tenant_id, name, description,
				generated_by, ai_generation_params, performance_metrics,
				last_ai_update, product_selection_rules, category_distribution,
				brand_preferences, regional_adaptations, created_at, updated_at
			FROM business_type_templates
			WHERE business_type_id = $1 AND tenant_id IS NULL
			ORDER BY created_at DESC
		`
		args = []interface{}{businessTypeID}
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query templates: %w", err)
	}
	defer rows.Close()

	var templates []*entity.AITemplate
	for rows.Next() {
		template := &entity.AITemplate{}
		var aiGenParams, perfMetrics, prodRules, catDist, brandPrefs, regionalAdapt []byte

		err := rows.Scan(
			&template.ID,
			&template.BusinessTypeID,
			&template.TenantID,
			&template.Name,
			&template.Description,
			&template.GeneratedBy,
			&aiGenParams,
			&perfMetrics,
			&template.LastAIUpdate,
			&prodRules,
			&catDist,
			&brandPrefs,
			&regionalAdapt,
			&template.CreatedAt,
			&template.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan template: %w", err)
		}

		// Unmarshal JSON fields
		json.Unmarshal(aiGenParams, &template.AIGenerationParams)
		json.Unmarshal(perfMetrics, &template.PerformanceMetrics)
		json.Unmarshal(prodRules, &template.ProductSelectionRules)
		json.Unmarshal(catDist, &template.CategoryDistribution)
		json.Unmarshal(brandPrefs, &template.BrandPreferences)
		json.Unmarshal(regionalAdapt, &template.RegionalAdaptations)

		templates = append(templates, template)
	}

	return templates, nil
}

// Delete deletes a template
func (r *AITemplatePostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM business_type_templates WHERE id = $1`
	
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete template: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrTemplateNotFound
	}

	return nil
}

// SaveTemplateProduct saves a template-product association
func (r *AITemplatePostgresRepository) SaveTemplateProduct(ctx context.Context, templateProduct *entity.TemplateGlobalProduct) error {
	query := `
		INSERT INTO template_global_products (
			id, template_id, global_product_id, priority, quantity_suggestion,
			ai_reasoning, relevance_score, category_match_score, brand_match_score,
			regional_preference_score, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := r.db.ExecContext(ctx, query,
		templateProduct.ID,
		templateProduct.TemplateID,
		templateProduct.GlobalProductID,
		templateProduct.Priority,
		templateProduct.QuantitySuggestion,
		templateProduct.AIReasoning,
		templateProduct.RelevanceScore,
		templateProduct.CategoryMatchScore,
		templateProduct.BrandMatchScore,
		templateProduct.RegionalPreferenceScore,
		templateProduct.CreatedAt,
		templateProduct.UpdatedAt,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return exception.ErrDuplicateProduct
		}
		return fmt.Errorf("failed to save template product: %w", err)
	}

	return nil
}

// UpdateTemplateProduct updates a template-product association
func (r *AITemplatePostgresRepository) UpdateTemplateProduct(ctx context.Context, templateProduct *entity.TemplateGlobalProduct) error {
	query := `
		UPDATE template_global_products SET
			priority = $3, quantity_suggestion = $4, ai_reasoning = $5,
			relevance_score = $6, category_match_score = $7, brand_match_score = $8,
			regional_preference_score = $9, updated_at = $10
		WHERE template_id = $1 AND global_product_id = $2
	`

	result, err := r.db.ExecContext(ctx, query,
		templateProduct.TemplateID,
		templateProduct.GlobalProductID,
		templateProduct.Priority,
		templateProduct.QuantitySuggestion,
		templateProduct.AIReasoning,
		templateProduct.RelevanceScore,
		templateProduct.CategoryMatchScore,
		templateProduct.BrandMatchScore,
		templateProduct.RegionalPreferenceScore,
		templateProduct.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update template product: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrProductNotFound
	}

	return nil
}

// FindTemplateProducts finds all products for a template
func (r *AITemplatePostgresRepository) FindTemplateProducts(ctx context.Context, templateID uuid.UUID) ([]*entity.TemplateGlobalProduct, error) {
	query := `
		SELECT id, template_id, global_product_id, priority, quantity_suggestion,
			ai_reasoning, relevance_score, category_match_score, brand_match_score,
			regional_preference_score, created_at, updated_at
		FROM template_global_products
		WHERE template_id = $1
		ORDER BY priority, relevance_score DESC
	`

	rows, err := r.db.QueryContext(ctx, query, templateID)
	if err != nil {
		return nil, fmt.Errorf("failed to query template products: %w", err)
	}
	defer rows.Close()

	var products []*entity.TemplateGlobalProduct
	for rows.Next() {
		product := &entity.TemplateGlobalProduct{}
		err := rows.Scan(
			&product.ID,
			&product.TemplateID,
			&product.GlobalProductID,
			&product.Priority,
			&product.QuantitySuggestion,
			&product.AIReasoning,
			&product.RelevanceScore,
			&product.CategoryMatchScore,
			&product.BrandMatchScore,
			&product.RegionalPreferenceScore,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan template product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// DeleteTemplateProduct deletes a template-product association
func (r *AITemplatePostgresRepository) DeleteTemplateProduct(ctx context.Context, templateID, globalProductID uuid.UUID) error {
	query := `DELETE FROM template_global_products WHERE template_id = $1 AND global_product_id = $2`
	
	result, err := r.db.ExecContext(ctx, query, templateID, globalProductID)
	if err != nil {
		return fmt.Errorf("failed to delete template product: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return exception.ErrProductNotFound
	}

	return nil
}

// SaveGenerationHistory saves AI generation history
func (r *AITemplatePostgresRepository) SaveGenerationHistory(ctx context.Context, history *entity.AIGenerationHistory) error {
	query := `
		INSERT INTO ai_template_generation_history (
			id, template_id, tenant_id, business_type_id, generation_params,
			ai_model, prompt_template, generated_content, generation_status,
			error_message, generation_time_ms, applied_at, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	genParams, _ := json.Marshal(history.GenerationParams)
	genContent, _ := json.Marshal(history.GeneratedContent)

	_, err := r.db.ExecContext(ctx, query,
		history.ID,
		history.TemplateID,
		history.TenantID,
		history.BusinessTypeID,
		genParams,
		history.AIModel,
		history.PromptTemplate,
		genContent,
		history.GenerationStatus,
		history.ErrorMessage,
		history.GenerationTimeMs,
		history.AppliedAt,
		history.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save generation history: %w", err)
	}

	return nil
}

// FindGenerationHistory finds generation history
func (r *AITemplatePostgresRepository) FindGenerationHistory(
	ctx context.Context,
	templateID *uuid.UUID,
	tenantID *uuid.UUID,
	limit int,
) ([]*entity.AIGenerationHistory, error) {
	query := `
		SELECT id, template_id, tenant_id, business_type_id, generation_params,
			ai_model, prompt_template, generated_content, generation_status,
			error_message, generation_time_ms, applied_at, created_at
		FROM ai_template_generation_history
		WHERE ($1::uuid IS NULL OR template_id = $1)
		  AND ($2::uuid IS NULL OR tenant_id = $2)
		ORDER BY created_at DESC
		LIMIT $3
	`

	rows, err := r.db.QueryContext(ctx, query, templateID, tenantID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query generation history: %w", err)
	}
	defer rows.Close()

	var histories []*entity.AIGenerationHistory
	for rows.Next() {
		history := &entity.AIGenerationHistory{}
		var genParams, genContent []byte

		err := rows.Scan(
			&history.ID,
			&history.TemplateID,
			&history.TenantID,
			&history.BusinessTypeID,
			&genParams,
			&history.AIModel,
			&history.PromptTemplate,
			&genContent,
			&history.GenerationStatus,
			&history.ErrorMessage,
			&history.GenerationTimeMs,
			&history.AppliedAt,
			&history.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan generation history: %w", err)
		}

		json.Unmarshal(genParams, &history.GenerationParams)
		json.Unmarshal(genContent, &history.GeneratedContent)

		histories = append(histories, history)
	}

	return histories, nil
}

// SavePerformanceMetric saves a performance metric
func (r *AITemplatePostgresRepository) SavePerformanceMetric(ctx context.Context, metric *entity.AIPerformanceMetric) error {
	query := `
		INSERT INTO ai_template_performance_metrics (
			id, template_id, metric_type, metric_value, metric_metadata,
			period_start, period_end, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (template_id, metric_type, period_start, period_end) 
		DO UPDATE SET metric_value = $4, metric_metadata = $5
	`

	metadata, _ := json.Marshal(metric.MetricMetadata)

	_, err := r.db.ExecContext(ctx, query,
		metric.ID,
		metric.TemplateID,
		metric.MetricType,
		metric.MetricValue,
		metadata,
		metric.PeriodStart,
		metric.PeriodEnd,
		metric.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save performance metric: %w", err)
	}

	return nil
}

// FindPerformanceMetrics finds performance metrics for a template
func (r *AITemplatePostgresRepository) FindPerformanceMetrics(
	ctx context.Context,
	templateID uuid.UUID,
	metricType string,
) ([]*entity.AIPerformanceMetric, error) {
	query := `
		SELECT id, template_id, metric_type, metric_value, metric_metadata,
			period_start, period_end, created_at
		FROM ai_template_performance_metrics
		WHERE template_id = $1 AND metric_type = $2
		ORDER BY period_start DESC
	`

	rows, err := r.db.QueryContext(ctx, query, templateID, metricType)
	if err != nil {
		return nil, fmt.Errorf("failed to query performance metrics: %w", err)
	}
	defer rows.Close()

	var metrics []*entity.AIPerformanceMetric
	for rows.Next() {
		metric := &entity.AIPerformanceMetric{}
		var metadata []byte

		err := rows.Scan(
			&metric.ID,
			&metric.TemplateID,
			&metric.MetricType,
			&metric.MetricValue,
			&metadata,
			&metric.PeriodStart,
			&metric.PeriodEnd,
			&metric.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan performance metric: %w", err)
		}

		json.Unmarshal(metadata, &metric.MetricMetadata)
		metrics = append(metrics, metric)
	}

	return metrics, nil
}

// GetAggregatedMetrics gets aggregated metrics for a template
func (r *AITemplatePostgresRepository) GetAggregatedMetrics(
	ctx context.Context,
	templateID uuid.UUID,
) (map[string]float64, error) {
	query := `
		SELECT metric_type, AVG(metric_value) as avg_value
		FROM ai_template_performance_metrics
		WHERE template_id = $1
		GROUP BY metric_type
	`

	rows, err := r.db.QueryContext(ctx, query, templateID)
	if err != nil {
		return nil, fmt.Errorf("failed to query aggregated metrics: %w", err)
	}
	defer rows.Close()

	metrics := make(map[string]float64)
	for rows.Next() {
		var metricType string
		var avgValue float64

		err := rows.Scan(&metricType, &avgValue)
		if err != nil {
			return nil, fmt.Errorf("failed to scan aggregated metric: %w", err)
		}

		metrics[metricType] = avgValue
	}

	return metrics, nil
}

// SaveProductFeedback saves product feedback
func (r *AITemplatePostgresRepository) SaveProductFeedback(ctx context.Context, feedback *entity.AIProductFeedback) error {
	query := `
		INSERT INTO ai_product_feedback (
			id, tenant_id, template_id, global_product_id, action,
			original_quantity, new_quantity, replacement_product_id,
			feedback_reason, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		feedback.ID,
		feedback.TenantID,
		feedback.TemplateID,
		feedback.GlobalProductID,
		feedback.Action,
		feedback.OriginalQuantity,
		feedback.NewQuantity,
		feedback.ReplacementProductID,
		feedback.FeedbackReason,
		feedback.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save product feedback: %w", err)
	}

	return nil
}

// FindProductFeedback finds product feedback
func (r *AITemplatePostgresRepository) FindProductFeedback(
	ctx context.Context,
	templateID *uuid.UUID,
	tenantID uuid.UUID,
) ([]*entity.AIProductFeedback, error) {
	query := `
		SELECT id, tenant_id, template_id, global_product_id, action,
			original_quantity, new_quantity, replacement_product_id,
			feedback_reason, created_at
		FROM ai_product_feedback
		WHERE ($1::uuid IS NULL OR template_id = $1)
		  AND ($2::uuid IS NULL OR tenant_id = $2)
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, templateID, tenantID)
	if err != nil {
		return nil, fmt.Errorf("failed to query product feedback: %w", err)
	}
	defer rows.Close()

	var feedbacks []*entity.AIProductFeedback
	for rows.Next() {
		feedback := &entity.AIProductFeedback{}
		err := rows.Scan(
			&feedback.ID,
			&feedback.TenantID,
			&feedback.TemplateID,
			&feedback.GlobalProductID,
			&feedback.Action,
			&feedback.OriginalQuantity,
			&feedback.NewQuantity,
			&feedback.ReplacementProductID,
			&feedback.FeedbackReason,
			&feedback.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product feedback: %w", err)
		}
		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}

// GetFeedbackSummary gets a summary of feedback for a template
func (r *AITemplatePostgresRepository) GetFeedbackSummary(
	ctx context.Context,
	templateID uuid.UUID,
) (map[string]int, error) {
	query := `
		SELECT action, COUNT(*) as count
		FROM ai_product_feedback
		WHERE template_id = $1
		GROUP BY action
	`

	rows, err := r.db.QueryContext(ctx, query, templateID)
	if err != nil {
		return nil, fmt.Errorf("failed to query feedback summary: %w", err)
	}
	defer rows.Close()

	summary := make(map[string]int)
	for rows.Next() {
		var action string
		var count int

		err := rows.Scan(&action, &count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan feedback summary: %w", err)
		}

		summary[action] = count
	}

	return summary, nil
}