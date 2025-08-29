package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofrs/uuid/v5"
	"saas-mt-pim-service/src/template_ai/domain/entity"
	"saas-mt-pim-service/src/template_ai/domain/exception"
	"saas-mt-pim-service/src/template_ai/domain/port"
)

// FeedbackPostgresRepository implements port.FeedbackRepository
type FeedbackPostgresRepository struct {
	db *sql.DB
}

// NewFeedbackPostgresRepository creates a new instance
func NewFeedbackPostgresRepository(db *sql.DB) port.FeedbackRepository {
	return &FeedbackPostgresRepository{
		db: db,
	}
}

// Save saves product feedback
func (r *FeedbackPostgresRepository) Save(ctx context.Context, feedback *entity.ProductFeedback) error {
	query := `
		INSERT INTO ai_product_feedback 
		(id, tenant_id, template_id, global_product_id, action, 
		 original_quantity, new_quantity, replacement_product_id, 
		 feedback_reason, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
		return exception.NewInternalError("failed to save feedback: " + err.Error())
	}

	return nil
}

// GetByTemplate retrieves feedback for a template
func (r *FeedbackPostgresRepository) GetByTemplate(ctx context.Context, templateID uuid.UUID) ([]*entity.ProductFeedback, error) {
	return r.GetFeedbackByTemplate(ctx, templateID)
}

// GetByTenant retrieves feedback for a tenant within a date range
func (r *FeedbackPostgresRepository) GetByTenant(ctx context.Context, tenantID uuid.UUID, startDate, endDate time.Time) ([]*entity.ProductFeedback, error) {
	query := `
		SELECT id, tenant_id, template_id, global_product_id, action,
			   original_quantity, new_quantity, replacement_product_id,
			   feedback_reason, created_at
		FROM ai_product_feedback
		WHERE tenant_id = $1 AND created_at BETWEEN $2 AND $3
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID, startDate, endDate)
	if err != nil {
		return nil, exception.NewInternalError("failed to query feedback: " + err.Error())
	}
	defer rows.Close()

	var feedbackList []*entity.ProductFeedback
	for rows.Next() {
		feedback := &entity.ProductFeedback{}
		var replacementProductID sql.NullString

		err := rows.Scan(
			&feedback.ID,
			&feedback.TenantID,
			&feedback.TemplateID,
			&feedback.GlobalProductID,
			&feedback.Action,
			&feedback.OriginalQuantity,
			&feedback.NewQuantity,
			&replacementProductID,
			&feedback.FeedbackReason,
			&feedback.CreatedAt,
		)
		if err != nil {
			return nil, exception.NewInternalError("failed to scan feedback: " + err.Error())
		}

		if replacementProductID.Valid {
			id, _ := uuid.FromString(replacementProductID.String)
			feedback.ReplacementProductID = &id
		}

		feedbackList = append(feedbackList, feedback)
	}

	return feedbackList, nil
}

// GetFeedbackByTemplate retrieves feedback for a template
func (r *FeedbackPostgresRepository) GetFeedbackByTemplate(ctx context.Context, templateID uuid.UUID) ([]*entity.ProductFeedback, error) {
	query := `
		SELECT id, tenant_id, template_id, global_product_id, action,
			   original_quantity, new_quantity, replacement_product_id,
			   feedback_reason, created_at
		FROM ai_product_feedback
		WHERE template_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, templateID)
	if err != nil {
		return nil, exception.NewInternalError("failed to query feedback: " + err.Error())
	}
	defer rows.Close()

	var feedbackList []*entity.ProductFeedback
	for rows.Next() {
		feedback := &entity.ProductFeedback{}
		var replacementProductID sql.NullString

		err := rows.Scan(
			&feedback.ID,
			&feedback.TenantID,
			&feedback.TemplateID,
			&feedback.GlobalProductID,
			&feedback.Action,
			&feedback.OriginalQuantity,
			&feedback.NewQuantity,
			&replacementProductID,
			&feedback.FeedbackReason,
			&feedback.CreatedAt,
		)
		if err != nil {
			return nil, exception.NewInternalError("failed to scan feedback: " + err.Error())
		}

		if replacementProductID.Valid {
			id, _ := uuid.FromString(replacementProductID.String)
			feedback.ReplacementProductID = &id
		}

		feedbackList = append(feedbackList, feedback)
	}

	return feedbackList, nil
}

// GetFeedbackByProduct retrieves feedback for a specific product
func (r *FeedbackPostgresRepository) GetFeedbackByProduct(ctx context.Context, productID uuid.UUID) ([]*entity.ProductFeedback, error) {
	query := `
		SELECT id, tenant_id, template_id, global_product_id, action,
			   original_quantity, new_quantity, replacement_product_id,
			   feedback_reason, created_at
		FROM ai_product_feedback
		WHERE global_product_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, productID)
	if err != nil {
		return nil, exception.NewInternalError("failed to query feedback: " + err.Error())
	}
	defer rows.Close()

	var feedbackList []*entity.ProductFeedback
	for rows.Next() {
		feedback := &entity.ProductFeedback{}
		var replacementProductID sql.NullString

		err := rows.Scan(
			&feedback.ID,
			&feedback.TenantID,
			&feedback.TemplateID,
			&feedback.GlobalProductID,
			&feedback.Action,
			&feedback.OriginalQuantity,
			&feedback.NewQuantity,
			&replacementProductID,
			&feedback.FeedbackReason,
			&feedback.CreatedAt,
		)
		if err != nil {
			return nil, exception.NewInternalError("failed to scan feedback: " + err.Error())
		}

		if replacementProductID.Valid {
			id, _ := uuid.FromString(replacementProductID.String)
			feedback.ReplacementProductID = &id
		}

		feedbackList = append(feedbackList, feedback)
	}

	return feedbackList, nil
}

// GetAggregatedFeedback retrieves aggregated feedback statistics
func (r *FeedbackPostgresRepository) GetAggregatedFeedback(ctx context.Context, templateID uuid.UUID, startDate, endDate time.Time) (map[string]interface{}, error) {
	query := `
		SELECT 
			action,
			COUNT(*) as count,
			COUNT(DISTINCT tenant_id) as unique_tenants
		FROM ai_product_feedback
		WHERE template_id = $1 
		AND created_at BETWEEN $2 AND $3
		GROUP BY action
	`

	rows, err := r.db.QueryContext(ctx, query, templateID, startDate, endDate)
	if err != nil {
		return nil, exception.NewInternalError("failed to query aggregated feedback: " + err.Error())
	}
	defer rows.Close()

	actionStats := make(map[string]map[string]int)
	totalCount := 0
	uniqueTenants := make(map[uuid.UUID]bool)

	for rows.Next() {
		var action string
		var count, tenantCount int

		err := rows.Scan(&action, &count, &tenantCount)
		if err != nil {
			return nil, exception.NewInternalError("failed to scan aggregated feedback: " + err.Error())
		}

		actionStats[action] = map[string]int{
			"count":          count,
			"unique_tenants": tenantCount,
		}
		totalCount += count
	}

	// Get most modified products
	productQuery := `
		SELECT global_product_id, COUNT(*) as modification_count
		FROM ai_product_feedback
		WHERE template_id = $1 
		AND created_at BETWEEN $2 AND $3
		AND action IN ('removed', 'quantity_changed', 'replaced')
		GROUP BY global_product_id
		ORDER BY modification_count DESC
		LIMIT 10
	`

	productRows, err := r.db.QueryContext(ctx, productQuery, templateID, startDate, endDate)
	if err != nil {
		return nil, exception.NewInternalError("failed to query product modifications: " + err.Error())
	}
	defer productRows.Close()

	var topModifiedProducts []map[string]interface{}
	for productRows.Next() {
		var productID uuid.UUID
		var count int

		err := productRows.Scan(&productID, &count)
		if err != nil {
			return nil, exception.NewInternalError("failed to scan product modifications: " + err.Error())
		}

		topModifiedProducts = append(topModifiedProducts, map[string]interface{}{
			"product_id": productID.String(),
			"count":      count,
		})
	}

	result := map[string]interface{}{
		"action_breakdown":       actionStats,
		"total_feedback_count":   totalCount,
		"unique_tenants":         len(uniqueTenants),
		"top_modified_products":  topModifiedProducts,
		"period_start":           startDate,
		"period_end":             endDate,
	}

	return result, nil
}

// GetCommonModifications retrieves the most common modifications
func (r *FeedbackPostgresRepository) GetCommonModifications(ctx context.Context, tenantID uuid.UUID, limit int) ([]port.CommonModification, error) {
	query := `
		WITH product_names AS (
			SELECT DISTINCT global_product_id, 
				   FIRST_VALUE(feedback_reason) OVER (PARTITION BY global_product_id ORDER BY created_at DESC) as latest_reason
			FROM ai_product_feedback
			WHERE tenant_id = $1
		)
		SELECT 
			f.action,
			f.global_product_id,
			pn.latest_reason,
			COUNT(*) as frequency,
			COUNT(DISTINCT f.template_id) as affected_tenants
		FROM ai_product_feedback f
		JOIN product_names pn ON f.global_product_id = pn.global_product_id
		WHERE f.tenant_id = $1
		GROUP BY f.action, f.global_product_id, pn.latest_reason
		ORDER BY frequency DESC
		LIMIT $2
	`

	rows, err := r.db.QueryContext(ctx, query, tenantID, limit)
	if err != nil {
		return nil, exception.NewInternalError("failed to query common modifications: " + err.Error())
	}
	defer rows.Close()

	var modifications []port.CommonModification
	for rows.Next() {
		mod := port.CommonModification{}
		var productID uuid.UUID
		var reason sql.NullString

		err := rows.Scan(
			&mod.Action,
			&productID,
			&reason,
			&mod.Frequency,
			&mod.AffectedTenants,
		)
		if err != nil {
			return nil, exception.NewInternalError("failed to scan common modification: " + err.Error())
		}

		mod.ProductID = productID
		if reason.Valid {
			mod.Reason = reason.String
		}

		modifications = append(modifications, mod)
	}

	return modifications, nil
}