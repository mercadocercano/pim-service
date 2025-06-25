package response

import "time"

// MarketplaceCategoryResponse representa la respuesta de una categoría marketplace
type MarketplaceCategoryResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description *string   `json:"description"`
	ParentID    *string   `json:"parent_id"`
	Level       int       `json:"level"`
	IsActive    bool      `json:"is_active"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MarketplaceCategoryTreeResponse representa una categoría con sus hijos para navegación
type MarketplaceCategoryTreeResponse struct {
	*MarketplaceCategoryResponse
	Children []*MarketplaceCategoryTreeResponse `json:"children,omitempty"`
	Path     []string                           `json:"path"` // Breadcrumb path
}

// MarketplaceCategoryListResponse representa una lista paginada de categorías
type MarketplaceCategoryListResponse struct {
	Categories []*MarketplaceCategoryResponse `json:"categories"`
	Total      int                            `json:"total"`
	Page       int                            `json:"page"`
	PageSize   int                            `json:"page_size"`
	TotalPages int                            `json:"total_pages"`
}
