//go:build integration

// Package brands_test contiene tests de integración para los endpoints de marketplace-brands.
// Usa TestContainers para levantar un PostgreSQL real y verificar el comportamiento completo
// del stack HTTP → Handler → Repository → DB.
//
// Ejecutar con: go test -tags=integration ./integration_test/brands/... -v
package brands_test

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

type marketplaceBrandResponse struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Slug               string   `json:"slug"`
	NormalizedName     string   `json:"normalized_name"`
	Description        string   `json:"description"`
	LogoURL            string   `json:"logo_url"`
	Website            string   `json:"website"`
	VerificationStatus string   `json:"verification_status"`
	QualityScore       float64  `json:"quality_score"`
	ProductCount       int      `json:"product_count"`
	IsActive           bool     `json:"is_active"`
	CategoryTags       []string `json:"category_tags"`
	Aliases            []string `json:"aliases"`
	BackgroundColor    string   `json:"background_color"`
	TextColor          string   `json:"text_color"`
	Typography         string   `json:"typography"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
}

type paginationInfo struct {
	Offset     int  `json:"offset"`
	Limit      int  `json:"limit"`
	Total      int  `json:"total"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
	TotalPages int  `json:"total_pages"`
}

// listBrandsResponse refleja la forma real devuelta por GetAllMarketplaceBrands:
// {"brands": [...], "pagination": {...}}
type listBrandsResponse struct {
	Brands     []marketplaceBrandResponse `json:"brands"`
	Pagination paginationInfo             `json:"pagination"`
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

func createBrandPayload(name string) map[string]interface{} {
	return map[string]interface{}{
		"name":        name,
		"description": "Marca creada en test de integración",
	}
}

// --- Tests ---

// TestGetAllMarketplaceBrands_Empty verifica que la lista vacía devuelve shape correcta.
func TestGetAllMarketplaceBrands_Empty(t *testing.T) {
	srv := newTestServer(t)

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/marketplace-brands", nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result listBrandsResponse
	decodeJSON(t, resp, &result)

	assert.NotNil(t, result.Brands, "campo 'brands' debe existir aunque esté vacío")
	assert.Equal(t, 0, result.Pagination.Total, "total debe ser 0 con tabla vacía")
	assert.Equal(t, 0, result.Pagination.Offset)
}

// TestGetAllMarketplaceBrands_WithPagination verifica paginación con marcas creadas.
func TestGetAllMarketplaceBrands_WithPagination(t *testing.T) {
	srv := newTestServer(t)

	// Crear 3 marcas
	names := []string{"Marca Alpha", "Marca Beta", "Marca Gamma"}
	for _, name := range names {
		resp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands",
			createBrandPayload(name), adminHeader())
		require.Equal(t, http.StatusCreated, resp.StatusCode, "crear marca %q debe retornar 201", name)
		resp.Body.Close()
	}

	// Listar con page_size=2
	resp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/marketplace-brands?page=1&page_size=2",
		nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result listBrandsResponse
	decodeJSON(t, resp, &result)

	assert.Len(t, result.Brands, 2, "page_size=2 debe retornar 2 items")
	assert.Equal(t, 3, result.Pagination.Total, "total debe ser 3")
	assert.Equal(t, 2, result.Pagination.TotalPages, "total_pages debe ser 2")
	assert.True(t, result.Pagination.HasNext, "debe haber página siguiente")
	assert.False(t, result.Pagination.HasPrev, "no debe haber página anterior en página 1")
}

// TestCreateMarketplaceBrand_ReturnsCreated verifica creación con status 201 y campos correctos.
func TestCreateMarketplaceBrand_ReturnsCreated(t *testing.T) {
	srv := newTestServer(t)

	payload := createBrandPayload("Nike Test")
	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands", payload, adminHeader())
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var brand marketplaceBrandResponse
	decodeJSON(t, resp, &brand)

	assert.NotEmpty(t, brand.ID, "id debe estar presente")
	assert.Equal(t, "Nike Test", brand.Name)
	assert.Equal(t, "unverified", brand.VerificationStatus,
		"verification_status por defecto debe ser 'unverified'")
	assert.Equal(t, 0, brand.ProductCount)
	assert.True(t, brand.IsActive, "is_active debe ser true por defecto")
	assert.NotEmpty(t, brand.CreatedAt)
	assert.NotEmpty(t, brand.UpdatedAt)
}

// TestCreateMarketplaceBrand_WithoutAdminRole_ReturnsForbidden verifica que usuarios no-admin son rechazados.
func TestCreateMarketplaceBrand_WithoutAdminRole_ReturnsForbidden(t *testing.T) {
	srv := newTestServer(t)

	h := http.Header{}
	h.Set("X-User-Role", "tenant_user")
	h.Set("Content-Type", "application/json")

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands",
		createBrandPayload("Marca Prohibida"), h)
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	resp.Body.Close()
}

// TestGetMarketplaceBrandByID_Found verifica que se puede obtener una marca por ID.
func TestGetMarketplaceBrandByID_Found(t *testing.T) {
	srv := newTestServer(t)

	// Crear marca
	createResp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands",
		createBrandPayload("Adidas Integration"), adminHeader())
	require.Equal(t, http.StatusCreated, createResp.StatusCode)

	var created marketplaceBrandResponse
	decodeJSON(t, createResp, &created)
	require.NotEmpty(t, created.ID)

	// Obtener por ID
	getResp := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/marketplace-brands/%s", baseURL(srv), created.ID),
		nil, adminHeader())
	assert.Equal(t, http.StatusOK, getResp.StatusCode)

	var fetched marketplaceBrandResponse
	decodeJSON(t, getResp, &fetched)

	assert.Equal(t, created.ID, fetched.ID)
	assert.Equal(t, "Adidas Integration", fetched.Name)
	assert.Equal(t, "unverified", fetched.VerificationStatus)
}

// TestGetMarketplaceBrandByID_NotFound verifica 404 para ID inexistente.
func TestGetMarketplaceBrandByID_NotFound(t *testing.T) {
	srv := newTestServer(t)

	resp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/marketplace-brands/00000000-0000-0000-0000-000000000000",
		nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestUpdateMarketplaceBrand_UpdatesName verifica que PUT actualiza el nombre.
func TestUpdateMarketplaceBrand_UpdatesName(t *testing.T) {
	srv := newTestServer(t)

	// Crear marca
	createResp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands",
		createBrandPayload("Puma Original"), adminHeader())
	require.Equal(t, http.StatusCreated, createResp.StatusCode)

	var created marketplaceBrandResponse
	decodeJSON(t, createResp, &created)
	require.NotEmpty(t, created.ID)

	// Actualizar
	updatePayload := map[string]interface{}{
		"name":        "Puma Updated",
		"description": "Descripción actualizada",
		"is_active":   true,
	}
	updateResp := doRequest(t, http.MethodPut,
		fmt.Sprintf("%s/marketplace-brands/%s", baseURL(srv), created.ID),
		updatePayload, adminHeader())
	assert.Equal(t, http.StatusOK, updateResp.StatusCode)

	var updated marketplaceBrandResponse
	decodeJSON(t, updateResp, &updated)

	assert.Equal(t, "Puma Updated", updated.Name)
	assert.Equal(t, "Descripción actualizada", updated.Description)
}

// TestDeleteMarketplaceBrand_ReturnsNotImplemented verifica el estado actual del endpoint DELETE.
// DELETE retorna 501 ya que la implementación está pendiente (TODO en el handler).
func TestDeleteMarketplaceBrand_ReturnsNotImplemented(t *testing.T) {
	srv := newTestServer(t)

	// Crear marca
	createResp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands",
		createBrandPayload("Marca Para Eliminar"), adminHeader())
	require.Equal(t, http.StatusCreated, createResp.StatusCode)

	var created marketplaceBrandResponse
	decodeJSON(t, createResp, &created)
	require.NotEmpty(t, created.ID)

	// DELETE — estado actual: 501 Not Implemented
	deleteResp := doRequest(t, http.MethodDelete,
		fmt.Sprintf("%s/marketplace-brands/%s", baseURL(srv), created.ID),
		nil, adminHeader())
	// El handler retorna 501 pending implementación completa
	assert.Equal(t, http.StatusNotImplemented, deleteResp.StatusCode,
		"DELETE debe retornar 501 hasta que la implementación esté completa")
	deleteResp.Body.Close()
}

// TestVerifyMarketplaceBrand_ReturnsNotImplemented verifica el estado actual del endpoint verify.
// El endpoint está declarado pero retorna 501 — se documenta para que el equipo lo implemente.
func TestVerifyMarketplaceBrand_ReturnsNotImplemented(t *testing.T) {
	srv := newTestServer(t)

	// Crear marca
	createResp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands",
		createBrandPayload("Marca Para Verificar"), adminHeader())
	require.Equal(t, http.StatusCreated, createResp.StatusCode)

	var created marketplaceBrandResponse
	decodeJSON(t, createResp, &created)
	require.NotEmpty(t, created.ID)

	// PUT /:id/verify — estado actual: 501 Not Implemented
	verifyResp := doRequest(t, http.MethodPut,
		fmt.Sprintf("%s/marketplace-brands/%s/verify", baseURL(srv), created.ID),
		nil, adminHeader())
	assert.Equal(t, http.StatusNotImplemented, verifyResp.StatusCode,
		"verify debe retornar 501 hasta implementación completa")
	verifyResp.Body.Close()
}

// TestUnverifyMarketplaceBrand_ReturnsNotImplemented verifica el estado actual del endpoint unverify.
func TestUnverifyMarketplaceBrand_ReturnsNotImplemented(t *testing.T) {
	srv := newTestServer(t)

	// Crear marca
	createResp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands",
		createBrandPayload("Marca Para Desverificar"), adminHeader())
	require.Equal(t, http.StatusCreated, createResp.StatusCode)

	var created marketplaceBrandResponse
	decodeJSON(t, createResp, &created)
	require.NotEmpty(t, created.ID)

	// PUT /:id/unverify — estado actual: 501 Not Implemented
	unverifyResp := doRequest(t, http.MethodPut,
		fmt.Sprintf("%s/marketplace-brands/%s/unverify", baseURL(srv), created.ID),
		nil, adminHeader())
	assert.Equal(t, http.StatusNotImplemented, unverifyResp.StatusCode,
		"unverify debe retornar 501 hasta implementación completa")
	unverifyResp.Body.Close()
}

// TestCreateMarketplaceBrand_DuplicateName_ReturnsConflict verifica que nombres duplicados dan 409.
func TestCreateMarketplaceBrand_DuplicateName_ReturnsConflict(t *testing.T) {
	srv := newTestServer(t)

	payload := createBrandPayload("Marca Duplicada")

	// Primera creación
	resp1 := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands", payload, adminHeader())
	require.Equal(t, http.StatusCreated, resp1.StatusCode)
	resp1.Body.Close()

	// Segunda creación con el mismo nombre — debe dar 409
	resp2 := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands", payload, adminHeader())
	assert.Equal(t, http.StatusConflict, resp2.StatusCode,
		"nombre duplicado debe retornar 409 Conflict")
	resp2.Body.Close()
}

// TestBrandResponse_FieldsShape verifica que la respuesta incluye todos los campos esperados por mc_admin.
func TestBrandResponse_FieldsShape(t *testing.T) {
	srv := newTestServer(t)

	createResp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace-brands",
		createBrandPayload("Levi's Shape Test"), adminHeader())
	require.Equal(t, http.StatusCreated, createResp.StatusCode)

	// Decodificar como map para verificar presencia exacta de keys
	var raw map[string]interface{}
	defer createResp.Body.Close()
	require.NoError(t, json.NewDecoder(createResp.Body).Decode(&raw))

	requiredFields := []string{
		"id", "name", "slug", "normalized_name",
		"verification_status", "product_count", "is_active", "quality_score",
		"created_at", "updated_at",
	}
	for _, field := range requiredFields {
		_, ok := raw[field]
		assert.True(t, ok, "campo '%s' debe estar presente en la respuesta", field)
	}
}
