package usecase_test

import (
	"context"
	"errors"
	"testing"

	"saas-mt-pim-service/src/template_ai/application/request"
	"saas-mt-pim-service/src/template_ai/application/usecase"
	"saas-mt-pim-service/src/template_ai/domain/entity"

	"github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validApplyRequest(templateID, tenantID string) *request.ApplyDynamicTemplateRequest {
	return &request.ApplyDynamicTemplateRequest{
		TemplateID: templateID,
		TenantID:   tenantID,
		ApplyMode:  "full",
	}
}

func TestApplyDynamicTemplate_HappyPath_ReturnsSuccess(t *testing.T) {
	template := newTemplate()
	templateID := template.ID
	tenantID := newUUID()

	productID := newUUID()
	products := []*entity.TemplateGlobalProduct{
		{
			ID: newUUID(), TemplateID: templateID, GlobalProductID: productID,
			Priority: 1, QuantitySuggestion: 10, RelevanceScore: 0.9,
		},
	}

	globalProducts := map[uuid.UUID]interface{}{
		productID: map[string]interface{}{"name": "Coca Cola 500ml"},
	}

	repo := &mockAITemplateRepo{
		findByIDResult:    template,
		findProductsResult: products,
	}
	globalRepo := &mockGlobalProductRepo{findByIDsResult: globalProducts}

	uc := usecase.NewApplyDynamicTemplateUseCase(repo, globalRepo, &mockDomainService{}, &mockMapper{})

	result, err := uc.Execute(context.Background(), validApplyRequest(templateID.String(), tenantID.String()))

	require.NoError(t, err)
	assert.True(t, result.Success)
	assert.Equal(t, 1, result.AppliedProducts)
	assert.Equal(t, 0, result.FailedProducts)
}

func TestApplyDynamicTemplate_TemplateNotFound_ReturnsError(t *testing.T) {
	repo := &mockAITemplateRepo{findByIDErr: errors.New("not found")}
	globalRepo := &mockGlobalProductRepo{}

	uc := usecase.NewApplyDynamicTemplateUseCase(repo, globalRepo, &mockDomainService{}, &mockMapper{})

	_, err := uc.Execute(context.Background(), validApplyRequest(newUUID().String(), newUUID().String()))

	require.Error(t, err)
}

func TestApplyDynamicTemplate_InvalidRequest_ReturnsError(t *testing.T) {
	uc := usecase.NewApplyDynamicTemplateUseCase(
		&mockAITemplateRepo{}, &mockGlobalProductRepo{}, &mockDomainService{}, &mockMapper{},
	)

	req := &request.ApplyDynamicTemplateRequest{
		TemplateID: "invalid",
		TenantID:   "invalid",
		ApplyMode:  "full",
	}

	_, err := uc.Execute(context.Background(), req)

	require.Error(t, err)
}

func TestApplyDynamicTemplate_NoProducts_ReturnsEmptyResult(t *testing.T) {
	template := newTemplate()
	repo := &mockAITemplateRepo{
		findByIDResult:     template,
		findProductsResult: []*entity.TemplateGlobalProduct{},
	}
	globalRepo := &mockGlobalProductRepo{findByIDsResult: map[uuid.UUID]interface{}{}}

	uc := usecase.NewApplyDynamicTemplateUseCase(repo, globalRepo, &mockDomainService{}, &mockMapper{})

	result, err := uc.Execute(context.Background(), validApplyRequest(template.ID.String(), newUUID().String()))

	require.NoError(t, err)
	assert.Equal(t, 0, result.AppliedProducts)
}

func TestApplyDynamicTemplate_WithExcludedProducts_SkipsThem(t *testing.T) {
	template := newTemplate()
	productID1 := newUUID()
	productID2 := newUUID()

	products := []*entity.TemplateGlobalProduct{
		{ID: newUUID(), TemplateID: template.ID, GlobalProductID: productID1, Priority: 1, QuantitySuggestion: 10, RelevanceScore: 0.9},
		{ID: newUUID(), TemplateID: template.ID, GlobalProductID: productID2, Priority: 2, QuantitySuggestion: 5, RelevanceScore: 0.7},
	}

	globalProducts := map[uuid.UUID]interface{}{
		productID1: map[string]interface{}{"name": "Product 1"},
		productID2: map[string]interface{}{"name": "Product 2"},
	}

	repo := &mockAITemplateRepo{findByIDResult: template, findProductsResult: products}
	globalRepo := &mockGlobalProductRepo{findByIDsResult: globalProducts}

	uc := usecase.NewApplyDynamicTemplateUseCase(repo, globalRepo, &mockDomainService{}, &mockMapper{})

	req := validApplyRequest(template.ID.String(), newUUID().String())
	req.ExcludeProducts = []string{productID1.String()}

	result, err := uc.Execute(context.Background(), req)

	require.NoError(t, err)
	assert.Equal(t, 1, result.AppliedProducts) // only product2 applied
}
