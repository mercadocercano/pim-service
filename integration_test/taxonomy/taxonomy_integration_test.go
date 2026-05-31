//go:build integration

// Package taxonomy_test contiene tests de integración para los endpoints de marketplace-categories.
// Usa TestContainers para levantar un PostgreSQL real y verificar el comportamiento completo
// del stack HTTP → Handler → Repository → DB.
//
// Ejecutar con: go test -tags=integration ./integration_test/taxonomy/... -v
package taxonomy_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- Tipos de respuesta que refleja la forma real del JSON ---

type marketplaceCategoryResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
	ParentID    *string `json:"parent_id"`
	Level       int     `json:"level"`
	IsActive    bool    `json:"is_active"`
	SortOrder   int     `json:"sort_order"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type categoryPaginationInfo struct {
	Offset     int  `json:"offset"`
	Limit      int  `json:"limit"`
	Total      int  `json:"total"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
	TotalPages int  `json:"total_pages"`
}

type listCategoriesResponse struct {
	Categories []marketplaceCategoryResponse `json:"categories"`
	Pagination categoryPaginationInfo        `json:"pagination"`
}

type singleCategoryResponse struct {
	Category marketplaceCategoryResponse `json:"category"`
}

type categoryTreeNode struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Slug     string             `json:"slug"`
	Level    int                `json:"level"`
	ParentID *string            `json:"parent_id"`
	Children []categoryTreeNode `json:"children"`
}

type treeCategoriesResponse struct {
	Categories []categoryTreeNode `json:"categories"`
}

// --- Helpers de peticiones HTTP ---

func adminHeader() http.Header {
	h := http.Header{}
	h.Set("X-User-Role", "marketplace_admin")
	h.Set("Content-Type", "application/json")
	return h
}

func doRequest(t *testing.T, method, url string, body interface{}, headers http.Header) *http.Response {
	t.Helper()

	var bodyReader *bytes.Reader
	if body != nil {
		b, err := json.Marshal(body)
		require.NoError(t, err)
		bodyReader = bytes.NewReader(b)
	} else {
		bodyReader = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	require.NoError(t, err)

	for k, vals := range headers {
		for _, v := range vals {
			req.Header.Add(k, v)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	return resp
}

func decodeJSON(t *testing.T, resp *http.Response, target interface{}) {
	t.Helper()
	defer resp.Body.Close()
	require.NoError(t, json.NewDecoder(resp.Body).Decode(target))
}

func createCategoryPayload(name, slug string, parentID *string) map[string]interface{} {
	p := map[string]interface{}{
		"name":        name,
		"slug":        slug,
		"description": "Categoría creada en test de integración",
		"is_active":   true,
	}
	if parentID != nil {
		p["parent_id"] = *parentID
	}
	return p
}

func createCategory(t *testing.T, srv *testServer, name, slug string, parentID *string) marketplaceCategoryResponse {
	t.Helper()
	resp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/marketplace/categories",
		createCategoryPayload(name, slug, parentID),
		adminHeader(),
	)
	require.Equal(t, http.StatusCreated, resp.StatusCode, "crear categoría %q debe retornar 201", name)

	var created marketplaceCategoryResponse
	decodeJSON(t, resp, &created)
	require.NotEmpty(t, created.ID, "id debe estar presente")
	return created
}

// --- Tests ---

// TestCreateMarketplaceCategory_ReturnsCreated verifica creación con status 201 y campos correctos.
func TestCreateMarketplaceCategory_ReturnsCreated(t *testing.T) {
	srv := newTestServer(t)

	payload := createCategoryPayload("Electrónica", "electronica", nil)
	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace/categories", payload, adminHeader())
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var category marketplaceCategoryResponse
	decodeJSON(t, resp, &category)

	assert.NotEmpty(t, category.ID, "id debe estar presente")
	assert.Equal(t, "Electrónica", category.Name)
	assert.Equal(t, "electronica", category.Slug)
	assert.True(t, category.IsActive, "is_active debe ser true por defecto")
	assert.NotEmpty(t, category.CreatedAt)
	assert.NotEmpty(t, category.UpdatedAt)
}

// TestCreateMarketplaceCategory_WithParentID_ReturnsCreated verifica creación de subcategoría.
func TestCreateMarketplaceCategory_WithParentID_ReturnsCreated(t *testing.T) {
	srv := newTestServer(t)

	parent := createCategory(t, srv, "Electrónica", "electronica", nil)

	payload := createCategoryPayload("Celulares", "celulares", &parent.ID)
	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace/categories", payload, adminHeader())
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var child marketplaceCategoryResponse
	decodeJSON(t, resp, &child)

	assert.NotEmpty(t, child.ID)
	assert.Equal(t, "Celulares", child.Name)
	require.NotNil(t, child.ParentID, "parent_id debe estar presente")
	assert.Equal(t, parent.ID, *child.ParentID)
}

// TestGetAllMarketplaceCategories_ReturnsList verifica que se devuelve una lista paginada.
func TestGetAllMarketplaceCategories_ReturnsList(t *testing.T) {
	srv := newTestServer(t)

	names := []struct{ name, slug string }{
		{"Electrónica", "electronica"},
		{"Ropa", "ropa"},
		{"Hogar", "hogar"},
	}
	for _, n := range names {
		createCategory(t, srv, n.name, n.slug, nil)
	}

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/marketplace/categories", nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result listCategoriesResponse
	decodeJSON(t, resp, &result)

	assert.NotNil(t, result.Categories, "campo 'categories' debe existir")
	assert.GreaterOrEqual(t, result.Pagination.Total, 3, "total debe ser al menos 3")
}

// TestGetMarketplaceCategoriesTree_ReturnsTreeWithChildren verifica árbol jerárquico.
func TestGetMarketplaceCategoriesTree_ReturnsTreeWithChildren(t *testing.T) {
	srv := newTestServer(t)

	parent := createCategory(t, srv, "Tecnología", "tecnologia", nil)
	createCategory(t, srv, "Laptops", "laptops", &parent.ID)
	createCategory(t, srv, "Tablets", "tablets", &parent.ID)

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/marketplace/categories/tree", nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result treeCategoriesResponse
	decodeJSON(t, resp, &result)

	assert.NotEmpty(t, result.Categories, "árbol no debe estar vacío")

	var techNode *categoryTreeNode
	for i := range result.Categories {
		if result.Categories[i].ID == parent.ID {
			techNode = &result.Categories[i]
			break
		}
	}
	require.NotNil(t, techNode, "nodo raíz 'Tecnología' debe estar en el árbol")
	assert.Len(t, techNode.Children, 2, "debe tener 2 hijos")
}

// TestGetMarketplaceCategoryByID_Found verifica que se puede obtener una categoría por ID.
func TestGetMarketplaceCategoryByID_Found(t *testing.T) {
	srv := newTestServer(t)

	created := createCategory(t, srv, "Deportes", "deportes", nil)

	resp := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/marketplace/categories/%s", baseURL(srv), created.ID),
		nil, adminHeader(),
	)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result singleCategoryResponse
	decodeJSON(t, resp, &result)

	assert.Equal(t, created.ID, result.Category.ID)
	assert.Equal(t, "Deportes", result.Category.Name)
	assert.Equal(t, "deportes", result.Category.Slug)
}

// TestGetMarketplaceCategoryByID_NotFound verifica 404 para ID inexistente.
func TestGetMarketplaceCategoryByID_NotFound(t *testing.T) {
	srv := newTestServer(t)

	resp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/marketplace/categories/00000000-0000-0000-0000-000000000000",
		nil, adminHeader(),
	)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestUpdateMarketplaceCategory_UpdatesName verifica que PUT actualiza el nombre.
func TestUpdateMarketplaceCategory_UpdatesName(t *testing.T) {
	srv := newTestServer(t)

	created := createCategory(t, srv, "Juguetes Original", "juguetes-original", nil)

	updatePayload := map[string]interface{}{
		"name":        "Juguetes Actualizado",
		"slug":        "juguetes-actualizado",
		"description": "Descripción actualizada",
		"is_active":   true,
	}
	resp := doRequest(t, http.MethodPut,
		fmt.Sprintf("%s/marketplace/categories/%s", baseURL(srv), created.ID),
		updatePayload, adminHeader(),
	)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var updated marketplaceCategoryResponse
	decodeJSON(t, resp, &updated)

	assert.Equal(t, "Juguetes Actualizado", updated.Name)
	assert.Equal(t, "juguetes-actualizado", updated.Slug)
}

// TestDeleteMarketplaceCategory_LeafCategory_ReturnsNoContent verifica eliminación de categoría hoja.
func TestDeleteMarketplaceCategory_LeafCategory_ReturnsNoContent(t *testing.T) {
	srv := newTestServer(t)

	created := createCategory(t, srv, "Categoría Para Eliminar", "cat-para-eliminar", nil)

	resp := doRequest(t, http.MethodDelete,
		fmt.Sprintf("%s/marketplace/categories/%s", baseURL(srv), created.ID),
		nil, adminHeader(),
	)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	resp.Body.Close()

	// Verificar que ya no existe
	getResp := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/marketplace/categories/%s", baseURL(srv), created.ID),
		nil, adminHeader(),
	)
	assert.Equal(t, http.StatusNotFound, getResp.StatusCode)
	getResp.Body.Close()
}

// TestDeleteMarketplaceCategory_WithChildren_ReturnsConflict verifica que no se puede eliminar con hijos.
func TestDeleteMarketplaceCategory_WithChildren_ReturnsConflict(t *testing.T) {
	srv := newTestServer(t)

	parent := createCategory(t, srv, "Padre Con Hijos", "padre-con-hijos", nil)
	createCategory(t, srv, "Hijo", "hijo-de-padre", &parent.ID)

	resp := doRequest(t, http.MethodDelete,
		fmt.Sprintf("%s/marketplace/categories/%s", baseURL(srv), parent.ID),
		nil, adminHeader(),
	)
	assert.Equal(t, http.StatusConflict, resp.StatusCode)

	var body map[string]interface{}
	decodeJSON(t, resp, &body)
	assert.Contains(t, body["error"], "subcategorías", "el mensaje debe mencionar subcategorías")
}

// TestValidateCategoryHierarchy_ValidHierarchy verifica jerarquía válida retorna 200.
func TestValidateCategoryHierarchy_ValidHierarchy(t *testing.T) {
	srv := newTestServer(t)

	root := createCategory(t, srv, "Categoría Raíz", "categoria-raiz", nil)
	child := createCategory(t, srv, "Categoría Hija", "categoria-hija", &root.ID)

	payload := map[string]interface{}{
		"category_ids": []string{root.ID, child.ID},
	}
	resp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/marketplace/categories/validate-hierarchy",
		payload, adminHeader(),
	)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
}

// TestCreateMarketplaceCategory_DuplicateSlug_ReturnsConflict verifica que slugs duplicados dan 409.
func TestCreateMarketplaceCategory_DuplicateSlug_ReturnsConflict(t *testing.T) {
	srv := newTestServer(t)

	payload := createCategoryPayload("Categoría Duplicada", "slug-duplicado", nil)

	resp1 := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace/categories", payload, adminHeader())
	require.Equal(t, http.StatusCreated, resp1.StatusCode)
	resp1.Body.Close()

	resp2 := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace/categories", payload, adminHeader())
	assert.Equal(t, http.StatusConflict, resp2.StatusCode,
		"slug duplicado debe retornar 409 Conflict")
	resp2.Body.Close()
}
