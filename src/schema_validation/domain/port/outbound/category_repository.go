package outbound

import "context"

// CategoryRepository provides access to the tenant's existing categories.
type CategoryRepository interface {
	GetCategoryNames(ctx context.Context, tenantID string) ([]string, error)
}
