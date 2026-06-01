//go:build integration

package attributes_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- DTOs para deserializar respuestas ---

type attributeResponse struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	Slug                 string                 `json:"slug"`
	Type                 string                 `json:"type"`
	IsFilterable         bool                   `json:"is_filterable"`
	IsSearchable         bool                   `json:"is_searchable"`
	IsRequiredForListing bool                   `json:"is_required_for_listing"`
	ValidationRules      map[string]interface{} `json:"validation_rules"`
	SortOrder            int                    `json:"sort_order"`
	CreatedAt            string                 `json:"created_at"`
	UpdatedAt            string                 `json:"updated_at"`
}

type listAttributesResponse struct {
	Items      []attributeResponse `json:"items"`
	TotalCount int                 `json:"total_count"`
	Page       int                 `json:"page"`
	PageSize   int                 `json:"page_size"`
}

type attributeValueResponse struct {
	ID          string `json:"id"`
	AttributeID string `json:"attribute_id"`
	Value       string `json:"value"`
	Slug        string `json:"slug"`
	SortOrder   int    `json:"sort_order"`
	IsActive    bool   `json:"is_active"`
}

type listValuesResponse struct {
	Items []attributeValueResponse `json:"items"`
}

// --- Helpers ---

func adminHeader() http.Header {
	h := http.Header{}
	h.Set("Content-Type", "application/json")
	h.Set("X-User-Role", "super_admin")
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

	for key, values := range headers {
		for _, v := range values {
			req.Header.Set(key, v)
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

// --- Tests ---

// TestCreateAttribute_ValidPayload_Returns201 verifica la creación exitosa de un atributo
func TestCreateAttribute_ValidPayload_Returns201(t *testing.T) {
	srv := newTestServer(t)

	name := uniqueName("TestColor")
	payload := map[string]interface{}{
		"name":          name,
		"type":          "select",
		"is_filterable": true,
		"is_searchable": true,
	}

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace/attributes", payload, adminHeader())
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var created attributeResponse
	decodeJSON(t, resp, &created)
	assert.Equal(t, name, created.Name)
	assert.Equal(t, "select", created.Type)
	assert.NotEmpty(t, created.ID)
	assert.NotEmpty(t, created.Slug)
}

// TestCreateAttribute_MissingName_Returns400 verifica validación de nombre requerido
func TestCreateAttribute_MissingName_Returns400(t *testing.T) {
	srv := newTestServer(t)

	payload := map[string]interface{}{
		"type": "text",
	}

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace/attributes", payload, adminHeader())
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// TestGetAttribute_ExistingID_Returns200 verifica obtención por ID
func TestGetAttribute_ExistingID_Returns200(t *testing.T) {
	srv := newTestServer(t)
	attr := createTestAttribute(t, srv, "Material", "text")

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/marketplace/attributes/"+attr.ID, nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var got attributeResponse
	decodeJSON(t, resp, &got)
	assert.Equal(t, attr.ID, got.ID)
	assert.Equal(t, attr.Name, got.Name)
}

// TestGetAttribute_UnknownID_Returns404 verifica que un ID inexistente retorna 404
func TestGetAttribute_UnknownID_Returns404(t *testing.T) {
	srv := newTestServer(t)

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/marketplace/attributes/00000000-0000-0000-0000-000000000000", nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

// TestListAttributes_AfterCreate_ContainsCreated verifica que el atributo creado aparece en el listado
func TestListAttributes_AfterCreate_ContainsCreated(t *testing.T) {
	srv := newTestServer(t)
	attr := createTestAttribute(t, srv, "Talla", "select")

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/marketplace/attributes", nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var listed listAttributesResponse
	decodeJSON(t, resp, &listed)

	found := false
	for _, item := range listed.Items {
		if item.ID == attr.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "atributo creado debe aparecer en el listado")
}

// TestListAttributes_FilterByType_ReturnsOnlyMatchingType verifica el filtro por tipo
func TestListAttributes_FilterByType_ReturnsOnlyMatchingType(t *testing.T) {
	srv := newTestServer(t)
	createTestAttribute(t, srv, "ColorFiltro", "select")
	createTestAttribute(t, srv, "TextoFiltro", "text")

	url := baseURL(srv) + "/marketplace/attributes?type=select"
	resp := doRequest(t, http.MethodGet, url, nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var listed listAttributesResponse
	decodeJSON(t, resp, &listed)

	for _, item := range listed.Items {
		assert.Equal(t, "select", item.Type, "solo deben aparecer atributos de tipo select")
	}
	assert.GreaterOrEqual(t, len(listed.Items), 1)
}

// TestUpdateAttribute_ValidPayload_Returns200 verifica la actualización exitosa
func TestUpdateAttribute_ValidPayload_Returns200(t *testing.T) {
	srv := newTestServer(t)
	attr := createTestAttribute(t, srv, "Peso", "number")

	payload := map[string]interface{}{
		"name":          "Peso neto",
		"type":          "number",
		"is_filterable": false,
	}

	resp := doRequest(t, http.MethodPut, baseURL(srv)+"/marketplace/attributes/"+attr.ID, payload, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var updated attributeResponse
	decodeJSON(t, resp, &updated)
	assert.Equal(t, "Peso neto", updated.Name)
}

// TestDeleteAttribute_NotInUse_Returns204 verifica eliminación exitosa
func TestDeleteAttribute_NotInUse_Returns204(t *testing.T) {
	srv := newTestServer(t)
	attr := createTestAttribute(t, srv, "AtributoAEliminar", "boolean")

	resp := doRequest(t, http.MethodDelete, baseURL(srv)+"/marketplace/attributes/"+attr.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	// Verificar que ya no existe
	getRsp := doRequest(t, http.MethodGet, baseURL(srv)+"/marketplace/attributes/"+attr.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, getRsp.StatusCode)
	getRsp.Body.Close()
}

// TestDeleteAttribute_InUse_Returns409 verifica la protección 409 cuando el atributo está en uso
func TestDeleteAttribute_InUse_Returns409(t *testing.T) {
	srv := newTestServer(t)
	attr := createTestAttribute(t, srv, "ColorEnUso", "select")

	// Insertar un uso directo en la tabla variant_marketplace_attributes para simular que está en uso
	insertAttributeUsage(t, srv.DB, attr.ID)

	resp := doRequest(t, http.MethodDelete, baseURL(srv)+"/marketplace/attributes/"+attr.ID, nil, adminHeader())
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
}

// TestCRUD_Attribute_FullLifecycle verifica el ciclo completo create → get → update → get (updated) → delete → 404
func TestCRUD_Attribute_FullLifecycle(t *testing.T) {
	srv := newTestServer(t)

	// Create
	attr := createTestAttribute(t, srv, "Ciclo Completo", "text")
	assert.NotEmpty(t, attr.ID)

	// Get
	getRsp := doRequest(t, http.MethodGet, baseURL(srv)+"/marketplace/attributes/"+attr.ID, nil, adminHeader())
	assert.Equal(t, http.StatusOK, getRsp.StatusCode)
	getRsp.Body.Close()

	// Update
	updatePayload := map[string]interface{}{
		"name": "Ciclo Completo Actualizado",
		"type": "text",
	}
	updRsp := doRequest(t, http.MethodPut, baseURL(srv)+"/marketplace/attributes/"+attr.ID, updatePayload, adminHeader())
	assert.Equal(t, http.StatusOK, updRsp.StatusCode)
	var updated attributeResponse
	decodeJSON(t, updRsp, &updated)
	assert.Equal(t, "Ciclo Completo Actualizado", updated.Name)

	// Delete
	delRsp := doRequest(t, http.MethodDelete, baseURL(srv)+"/marketplace/attributes/"+attr.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNoContent, delRsp.StatusCode)
	delRsp.Body.Close()

	// 404 after delete
	notFoundRsp := doRequest(t, http.MethodGet, baseURL(srv)+"/marketplace/attributes/"+attr.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, notFoundRsp.StatusCode)
	notFoundRsp.Body.Close()
}

// TestAttributeValues_FullLifecycle verifica CRUD completo de valores
func TestAttributeValues_FullLifecycle(t *testing.T) {
	srv := newTestServer(t)
	attr := createTestAttribute(t, srv, "TallaValues", "select")

	valuesURL := fmt.Sprintf("%s/marketplace/attributes/%s/values", baseURL(srv), attr.ID)

	// Crear primer valor
	v1 := createTestValue(t, srv, attr.ID, "S", 0)
	assert.NotEmpty(t, v1.ID)
	assert.Equal(t, "S", v1.Value)

	// Crear segundo valor
	v2 := createTestValue(t, srv, attr.ID, "M", 1)
	assert.NotEmpty(t, v2.ID)

	// Listar valores → ambos presentes
	listRsp := doRequest(t, http.MethodGet, valuesURL, nil, adminHeader())
	assert.Equal(t, http.StatusOK, listRsp.StatusCode)
	var listed listValuesResponse
	decodeJSON(t, listRsp, &listed)
	assert.Len(t, listed.Items, 2)

	// Actualizar primer valor
	updatePayload := map[string]interface{}{"value": "XS", "sort_order": 0}
	updURL := fmt.Sprintf("%s/marketplace/attributes/%s/values/%s", baseURL(srv), attr.ID, v1.ID)
	updRsp := doRequest(t, http.MethodPut, updURL, updatePayload, adminHeader())
	assert.Equal(t, http.StatusOK, updRsp.StatusCode)
	var updatedVal attributeValueResponse
	decodeJSON(t, updRsp, &updatedVal)
	assert.Equal(t, "XS", updatedVal.Value)

	// Eliminar segundo valor
	delURL := fmt.Sprintf("%s/marketplace/attributes/%s/values/%s", baseURL(srv), attr.ID, v2.ID)
	delRsp := doRequest(t, http.MethodDelete, delURL, nil, adminHeader())
	assert.Equal(t, http.StatusNoContent, delRsp.StatusCode)
	delRsp.Body.Close()

	// Listar → solo queda 1
	listRsp2 := doRequest(t, http.MethodGet, valuesURL, nil, adminHeader())
	assert.Equal(t, http.StatusOK, listRsp2.StatusCode)
	var listed2 listValuesResponse
	decodeJSON(t, listRsp2, &listed2)
	assert.Len(t, listed2.Items, 1)
}

// TestCreateValue_AttributeNotFound_Returns404 verifica que crear valor en atributo inexistente devuelve 404
func TestCreateValue_AttributeNotFound_Returns404(t *testing.T) {
	srv := newTestServer(t)

	payload := map[string]interface{}{"value": "Rojo"}
	url := baseURL(srv) + "/marketplace/attributes/00000000-0000-0000-0000-000000000000/values"
	resp := doRequest(t, http.MethodPost, url, payload, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestDeleteValue_WrongAttribute_Returns404 verifica que no se puede eliminar un valor con el atributo equivocado
func TestDeleteValue_WrongAttribute_Returns404(t *testing.T) {
	srv := newTestServer(t)
	attr1 := createTestAttribute(t, srv, "AttrA", "select")
	attr2 := createTestAttribute(t, srv, "AttrB", "select")
	v := createTestValue(t, srv, attr1.ID, "Valor", 0)

	// Intentar eliminar el valor de attr1 usando el ID de attr2
	delURL := fmt.Sprintf("%s/marketplace/attributes/%s/values/%s", baseURL(srv), attr2.ID, v.ID)
	resp := doRequest(t, http.MethodDelete, delURL, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// --- Helpers de test ---

// uniqueName genera un nombre único para evitar conflictos con datos de seed
func uniqueName(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixNano())
}

// createTestAttribute crea un atributo de test con nombre único y retorna el response
func createTestAttribute(t *testing.T, srv *testServer, baseName, attrType string) attributeResponse {
	t.Helper()

	name := uniqueName(baseName)
	payload := map[string]interface{}{
		"name": name,
		"type": attrType,
	}
	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/marketplace/attributes", payload, adminHeader())
	require.Equal(t, http.StatusCreated, resp.StatusCode, "fallo al crear atributo de test: %s (body: ver logs)", name)

	var attr attributeResponse
	decodeJSON(t, resp, &attr)
	return attr
}

// createTestValue crea un valor de atributo de test y retorna el response deserializado
func createTestValue(t *testing.T, srv *testServer, attributeID, value string, sortOrder int) attributeValueResponse {
	t.Helper()

	payload := map[string]interface{}{
		"value":      value,
		"sort_order": sortOrder,
	}
	url := fmt.Sprintf("%s/marketplace/attributes/%s/values", baseURL(srv), attributeID)
	resp := doRequest(t, http.MethodPost, url, payload, adminHeader())
	require.Equal(t, http.StatusCreated, resp.StatusCode, "fallo al crear valor de test: %s", value)

	var v attributeValueResponse
	decodeJSON(t, resp, &v)
	return v
}

// insertAttributeUsage inserta fila en variant_marketplace_attributes para simular uso del atributo.
// Crea producto, variante y luego el vínculo atributo-variante respetando las FKs.
func insertAttributeUsage(t *testing.T, db *sql.DB, attributeID string) {
	t.Helper()

	tenantID := "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"

	// Crear producto
	productID := "11111111-1111-1111-1111-111111111111"
	_, err := db.Exec(`
		INSERT INTO products (id, tenant_id, name, status, created_at, updated_at)
		VALUES ($1, $2, 'test-product-for-attr', 'active', NOW(), NOW())
		ON CONFLICT (id) DO NOTHING
	`, productID, tenantID)
	require.NoError(t, err, "debe poder insertar producto de test")

	// Crear variante
	variantID := "22222222-2222-2222-2222-222222222222"
	_, err = db.Exec(`
		INSERT INTO product_variants (id, tenant_id, product_id, name, sku, status, is_default, sort_order, created_at, updated_at)
		VALUES ($1, $2, $3, 'variante-test', 'sku-attr-test-unique', 'active', true, 0, NOW(), NOW())
		ON CONFLICT (id) DO NOTHING
	`, variantID, tenantID, productID)
	require.NoError(t, err, "debe poder insertar variante de test")

	// Vincular atributo a variante
	_, err = db.Exec(`
		INSERT INTO variant_marketplace_attributes (id, variant_id, marketplace_attribute_id, value, created_at, updated_at)
		VALUES (gen_random_uuid(), $1, $2, 'rojo', NOW(), NOW())
	`, variantID, attributeID)
	require.NoError(t, err, "debe poder vincular atributo a variante")
}
