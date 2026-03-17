package usecase

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"saas-mt-pim-service/src/schema_validation/domain/service"
)

// ImportFromValidationRequest is the request for importing products from a cached validation
type ImportFromValidationRequest struct {
	ValidationID     string            `json:"validation_id" binding:"required"`
	ColumnMappings   map[string]string `json:"column_mappings"`
	Options          ImportOptions     `json:"options"`
}

type ImportOptions struct {
	SkipRowsWithErrors      bool `json:"skip_rows_with_errors"`
	CreateMissingCategories bool `json:"create_missing_categories"`
	CreateMissingBrands     bool `json:"create_missing_brands"`
	UpdateExistingProducts  bool `json:"update_existing_products"`
}

type ImportFromValidationResponse struct {
	ImportedCount     int                       `json:"imported_count"`
	SkippedCount      int                       `json:"skipped_count"`
	ErrorCount        int                       `json:"error_count"`
	Errors            []ImportRowError          `json:"errors,omitempty"`
	CreatedCategories []string                  `json:"created_categories,omitempty"`
	CreatedBrands     []string                  `json:"created_brands,omitempty"`
	BrandCandidates   []service.BrandCandidate  `json:"brand_candidates,omitempty"`
}

type ImportRowError struct {
	Row    int    `json:"row"`
	Reason string `json:"reason"`
}

type ImportFromValidationUseCase struct {
	schemaCache          SchemaValidationCache
	brandDeductionSvc    *service.BrandDeductionService
	categoryNamesGetter  func(ctx context.Context, tenantID string) ([]string, error)
	brandNamesGetter     func(ctx context.Context, tenantID string) ([]string, error)
}

func NewImportFromValidationUseCase(
	schemaCache SchemaValidationCache,
	brandDeductionSvc *service.BrandDeductionService,
	categoryNamesGetter func(ctx context.Context, tenantID string) ([]string, error),
	brandNamesGetter func(ctx context.Context, tenantID string) ([]string, error),
) *ImportFromValidationUseCase {
	return &ImportFromValidationUseCase{
		schemaCache:         schemaCache,
		brandDeductionSvc:   brandDeductionSvc,
		categoryNamesGetter: categoryNamesGetter,
		brandNamesGetter:    brandNamesGetter,
	}
}

// Execute extracts product data from the cached validation and returns it as structured product rows.
// The actual bulk creation is delegated to the existing /products/import endpoint by the controller.
func (uc *ImportFromValidationUseCase) Execute(
	ctx context.Context,
	req ImportFromValidationRequest,
	tenantID string,
) (*ImportFromValidationResponse, []map[string]interface{}, error) {
	validation, err := uc.schemaCache.Get(ctx, req.ValidationID)
	if err != nil {
		return nil, nil, fmt.Errorf("validation not found or expired")
	}

	if validation.IsExpired() {
		return nil, nil, fmt.Errorf("validation expired, please re-upload the file")
	}

	// Apply manual mappings if provided
	effectiveMappings := validation.SuggestedMappings
	if len(req.ColumnMappings) > 0 {
		effectiveMappings = req.ColumnMappings
	}

	// Build reverse mapping: schema_field → original_column_name
	fieldToColumn := make(map[string]string)
	for colName, field := range effectiveMappings {
		if field != "" && field != "__ignore__" {
			fieldToColumn[field] = colName
		}
	}

	// Extract rows from the table preview
	if validation.TablePreview == nil || len(validation.TablePreview.Rows) == 0 {
		return nil, nil, fmt.Errorf("no data rows in validation preview")
	}

	// Build header index map
	headerIndex := make(map[string]int)
	for _, h := range validation.TablePreview.Headers {
		headerIndex[h.Name] = h.Index
	}

	response := &ImportFromValidationResponse{
		Errors:            make([]ImportRowError, 0),
		CreatedCategories: make([]string, 0),
		CreatedBrands:     make([]string, 0),
	}

	var products []map[string]interface{}

	for _, row := range validation.TablePreview.Rows {
		cellValues := make(map[int]string)
		hasError := false
		for _, cell := range row.Cells {
			cellValues[cell.ColumnIndex] = cell.Value
			if cell.Status == "error" {
				hasError = true
			}
		}

		if hasError && req.Options.SkipRowsWithErrors {
			response.SkippedCount++
			continue
		}
		if hasError && !req.Options.SkipRowsWithErrors {
			response.ErrorCount++
			response.Errors = append(response.Errors, ImportRowError{
				Row:    row.RowNumber,
				Reason: "row contains validation errors",
			})
			continue
		}

		product := make(map[string]interface{})

		for field, colName := range fieldToColumn {
			colIdx, ok := headerIndex[colName]
			if !ok {
				continue
			}
			value, exists := cellValues[colIdx]
			if !exists || strings.TrimSpace(value) == "" {
				continue
			}

			switch field {
			case "price":
				cleaned := strings.ReplaceAll(strings.TrimPrefix(strings.TrimSpace(value), "$"), ",", ".")
				if p, err := strconv.ParseFloat(cleaned, 64); err == nil {
					product["price"] = p
				}
			case "stock":
				if s, err := strconv.Atoi(strings.TrimSpace(value)); err == nil {
					product["stock"] = s
				}
			default:
				product[field] = strings.TrimSpace(value)
			}
		}

		// Apply deduced category if no category was mapped
		if _, hasCategory := product["category_name"]; !hasCategory {
			if deduced, ok := validation.DeducedCategories[row.RowNumber-1]; ok {
				product["category_name"] = deduced
			}
		}

		// Validate minimum required fields
		name, hasName := product["name"]
		if !hasName || name == "" {
			response.ErrorCount++
			response.Errors = append(response.Errors, ImportRowError{
				Row:    row.RowNumber,
				Reason: "missing required field: name",
			})
			continue
		}

		products = append(products, product)
		response.ImportedCount++
	}

	// Run brand deduction on products that don't have a brand assigned
	if uc.brandDeductionSvc != nil {
		var productNames []string
		for _, p := range products {
			if _, hasBrand := p["brand_name"]; !hasBrand {
				if name, ok := p["name"].(string); ok {
					productNames = append(productNames, name)
				}
			}
		}
		if len(productNames) > 0 {
			var existingCategories, existingBrands []string
			if uc.categoryNamesGetter != nil {
				existingCategories, _ = uc.categoryNamesGetter(ctx, tenantID)
			}
			if uc.brandNamesGetter != nil {
				existingBrands, _ = uc.brandNamesGetter(ctx, tenantID)
			}
			response.BrandCandidates = uc.brandDeductionSvc.DeduceBrands(productNames, existingCategories, existingBrands)
		}
	}

	return response, products, nil
}
