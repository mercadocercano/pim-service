//go:build integration

package quickstart_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- DTOs para deserializar respuestas ---

type businessTypeResponse struct {
	ID          string                 `json:"id"`
	Code        string                 `json:"code"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Icon        string                 `json:"icon"`
	Color       string                 `json:"color"`
	IsActive    bool                   `json:"is_active"`
	SortOrder   int                    `json:"sort_order"`
	Metadata    map[string]interface{} `json:"metadata"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

type listBusinessTypesResponse struct {
	Items      []businessTypeResponse `json:"items"`
	TotalCount int                    `json:"total_count"`
	Page       int                    `json:"page"`
	PageSize   int                    `json:"page_size"`
	TotalPages int                    `json:"total_pages"`
}

type templateResponse struct {
	ID             string `json:"id"`
	BusinessTypeID string `json:"business_type_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Version        string `json:"version"`
	Region         string `json:"region"`
	IsActive       bool   `json:"is_active"`
	IsDefault      bool   `json:"is_default"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type templateWrapperResponse struct {
	Template templateResponse `json:"template"`
}

type listTemplatesResponse struct {
	Items      []templateResponse `json:"items"`
	TotalCount int                `json:"total_count"`
}

type templateAnalyticsResponse struct {
	TemplateID     string  `json:"template_id"`
	TenantsUsed    int     `json:"tenants_used"`
	LastActivated  *string `json:"last_activated"`
	CompletionRate float64 `json:"completion_rate"`
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

	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		require.NoError(t, err)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(bodyBytes))
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

func uniqueCode(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixNano())
}

// createTestBusinessType crea un business type con código único y retorna el response
func createTestBusinessType(t *testing.T, srv *testServer, baseName string) businessTypeResponse {
	t.Helper()

	code := uniqueCode(baseName)
	payload := map[string]interface{}{
		"code":        code,
		"name":        baseName + " Nombre",
		"description": "Descripción de test",
		"icon":        "shopping-bag",
		"color":       "#4F46E5",
	}

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/business-types", payload, adminHeader())
	require.Equal(t, http.StatusCreated, resp.StatusCode, "fallo al crear business type: %s", baseName)

	var bt businessTypeResponse
	decodeJSON(t, resp, &bt)
	return bt
}

// createTestTemplate crea un template vinculado a un business type y retorna el response
func createTestTemplate(t *testing.T, srv *testServer, businessTypeID, name string) templateResponse {
	t.Helper()

	payload := map[string]interface{}{
		"business_type_id": businessTypeID,
		"name":             name,
		"description":      "Template de test",
		"version":          "1.0.0",
		"region":           "AR",
	}

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/business-type-templates", payload, adminHeader())
	require.Equal(t, http.StatusCreated, resp.StatusCode, "fallo al crear template: %s", name)

	var wrapper templateWrapperResponse
	decodeJSON(t, resp, &wrapper)
	return wrapper.Template
}

// =============================================================================
// Tests de Business Types
// =============================================================================

// TestCreateBusinessType_ValidPayload_Returns201 verifica la creación exitosa
func TestCreateBusinessType_ValidPayload_Returns201(t *testing.T) {
	srv := newTestServer(t)

	code := uniqueCode("ferreteria")
	payload := map[string]interface{}{
		"code":        code,
		"name":        "Ferretería",
		"description": "Tienda de herramientas",
		"icon":        "wrench",
		"color":       "#FF6B35",
	}

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/business-types", payload, adminHeader())
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var bt businessTypeResponse
	decodeJSON(t, resp, &bt)
	assert.NotEmpty(t, bt.ID)
	assert.Equal(t, code, bt.Code)
	assert.Equal(t, "Ferretería", bt.Name)
	assert.True(t, bt.IsActive, "nuevo business type debe estar activo por defecto")
}

// TestCreateBusinessType_DuplicateCode_ReturnsConflict verifica el rechazo de códigos duplicados
func TestCreateBusinessType_DuplicateCode_ReturnsConflict(t *testing.T) {
	srv := newTestServer(t)

	bt := createTestBusinessType(t, srv, "duplicado")

	payload := map[string]interface{}{
		"code": bt.Code,
		"name": "Otro nombre con mismo código",
	}

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/business-types", payload, adminHeader())
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
	resp.Body.Close()
}

// TestGetBusinessType_ExistingID_Returns200 verifica la obtención por ID
func TestGetBusinessType_ExistingID_Returns200(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "farmacia")

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-types/"+bt.ID, nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var got businessTypeResponse
	decodeJSON(t, resp, &got)
	assert.Equal(t, bt.ID, got.ID)
	assert.Equal(t, bt.Code, got.Code)
}

// TestGetBusinessType_UnknownID_Returns404 verifica que un ID inexistente retorna 404
func TestGetBusinessType_UnknownID_Returns404(t *testing.T) {
	srv := newTestServer(t)

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-types/00000000-0000-0000-0000-000000000000", nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestListBusinessTypes_AfterCreate_ContainsCreated verifica que el BT aparece en el listado
func TestListBusinessTypes_AfterCreate_ContainsCreated(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "panaderia")

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-types", nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var listed listBusinessTypesResponse
	decodeJSON(t, resp, &listed)

	found := false
	for _, item := range listed.Items {
		if item.ID == bt.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "business type creado debe aparecer en el listado")
}

// TestListBusinessTypes_FilterByIsActive_ReturnsCorrectSubset verifica el filtro por is_active
func TestListBusinessTypes_FilterByIsActive_ReturnsCorrectSubset(t *testing.T) {
	srv := newTestServer(t)

	// Crear un BT activo y uno que luego desactivamos
	active := createTestBusinessType(t, srv, "activo")
	toDeactivate := createTestBusinessType(t, srv, "inactivo")

	// Desactivar el segundo
	deactivateResp := doRequest(t, http.MethodPatch, baseURL(srv)+"/business-types/"+toDeactivate.ID+"/deactivate", nil, adminHeader())
	require.Equal(t, http.StatusOK, deactivateResp.StatusCode)
	deactivateResp.Body.Close()

	// Filtrar solo activos
	activeURL := baseURL(srv) + "/business-types?is_active=true"
	resp := doRequest(t, http.MethodGet, activeURL, nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var listed listBusinessTypesResponse
	decodeJSON(t, resp, &listed)

	activeFound := false
	for _, item := range listed.Items {
		assert.True(t, item.IsActive, "solo deben aparecer BTs activos")
		if item.ID == active.ID {
			activeFound = true
		}
	}
	assert.True(t, activeFound, "el BT activo debe estar en la lista filtrada")
}

// TestUpdateBusinessType_ValidPayload_Returns200 verifica la actualización
func TestUpdateBusinessType_ValidPayload_Returns200(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "tecnologia")

	payload := map[string]interface{}{
		"name":        "Tecnología Actualizada",
		"description": "Descripción actualizada",
		"icon":        "laptop",
		"color":       "#3B82F6",
	}

	resp := doRequest(t, http.MethodPut, baseURL(srv)+"/business-types/"+bt.ID, payload, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var updated businessTypeResponse
	decodeJSON(t, resp, &updated)
	assert.Equal(t, "Tecnología Actualizada", updated.Name)
}

// TestDeleteBusinessType_NotFound_Returns404 verifica delete de ID inexistente
func TestDeleteBusinessType_NotFound_Returns404(t *testing.T) {
	srv := newTestServer(t)

	resp := doRequest(t, http.MethodDelete, baseURL(srv)+"/business-types/00000000-0000-0000-0000-000000000000", nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestActivateDeactivate_BusinessType_ChangesStatus verifica el toggle de estado
func TestActivateDeactivate_BusinessType_ChangesStatus(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "salud")

	// Desactivar
	deactivateURL := baseURL(srv) + "/business-types/" + bt.ID + "/deactivate"
	deactResp := doRequest(t, http.MethodPatch, deactivateURL, nil, adminHeader())
	assert.Equal(t, http.StatusOK, deactResp.StatusCode)
	var deactivated businessTypeResponse
	decodeJSON(t, deactResp, &deactivated)
	assert.False(t, deactivated.IsActive, "business type debe estar inactivo")

	// Activar
	activateURL := baseURL(srv) + "/business-types/" + bt.ID + "/activate"
	actResp := doRequest(t, http.MethodPatch, activateURL, nil, adminHeader())
	assert.Equal(t, http.StatusOK, actResp.StatusCode)
	var activated businessTypeResponse
	decodeJSON(t, actResp, &activated)
	assert.True(t, activated.IsActive, "business type debe estar activo nuevamente")
}

// TestCRUD_BusinessType_FullLifecycle verifica create → get → update → delete → 404
func TestCRUD_BusinessType_FullLifecycle(t *testing.T) {
	srv := newTestServer(t)

	// Create
	bt := createTestBusinessType(t, srv, "ciclo")
	assert.NotEmpty(t, bt.ID)

	// Get
	getResp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-types/"+bt.ID, nil, adminHeader())
	assert.Equal(t, http.StatusOK, getResp.StatusCode)
	getResp.Body.Close()

	// Update
	updatePayload := map[string]interface{}{
		"name": "Ciclo Actualizado",
	}
	updResp := doRequest(t, http.MethodPut, baseURL(srv)+"/business-types/"+bt.ID, updatePayload, adminHeader())
	assert.Equal(t, http.StatusOK, updResp.StatusCode)
	var updated businessTypeResponse
	decodeJSON(t, updResp, &updated)
	assert.Equal(t, "Ciclo Actualizado", updated.Name)

	// Delete
	delResp := doRequest(t, http.MethodDelete, baseURL(srv)+"/business-types/"+bt.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNoContent, delResp.StatusCode)
	delResp.Body.Close()

	// 404 after delete
	notFoundResp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-types/"+bt.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, notFoundResp.StatusCode)
	notFoundResp.Body.Close()
}

// =============================================================================
// Tests de Templates
// =============================================================================

// TestCreateTemplate_ValidPayload_Returns201 verifica la creación de template
func TestCreateTemplate_ValidPayload_Returns201(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "restaurante")

	payload := map[string]interface{}{
		"business_type_id": bt.ID,
		"name":             "Template Básico",
		"description":      "Template para restaurantes medianos",
		"version":          "1.0.0",
		"region":           "AR",
	}

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/business-type-templates", payload, adminHeader())
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var wrapper templateWrapperResponse
	decodeJSON(t, resp, &wrapper)
	assert.NotEmpty(t, wrapper.Template.ID)
	assert.Equal(t, bt.ID, wrapper.Template.BusinessTypeID)
	assert.Equal(t, "Template Básico", wrapper.Template.Name)
}

// TestCreateTemplate_MissingBusinessTypeID_Returns400 verifica validación de business_type_id
func TestCreateTemplate_MissingBusinessTypeID_Returns400(t *testing.T) {
	srv := newTestServer(t)

	payload := map[string]interface{}{
		"name": "Template sin BT",
	}

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/business-type-templates", payload, adminHeader())
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp.Body.Close()
}

// TestCreateTemplate_WithNonExistentBusinessType_ReturnsBadRequest verifica validación de FK
func TestCreateTemplate_WithNonExistentBusinessType_ReturnsBadRequest(t *testing.T) {
	srv := newTestServer(t)

	payload := map[string]interface{}{
		"business_type_id": "00000000-0000-0000-0000-000000000000",
		"name":             "Template BT inexistente",
	}

	resp := doRequest(t, http.MethodPost, baseURL(srv)+"/business-type-templates", payload, adminHeader())
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp.Body.Close()
}

// TestGetTemplate_ExistingID_Returns200 verifica la obtención de template por ID
func TestGetTemplate_ExistingID_Returns200(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "kiosco")
	tpl := createTestTemplate(t, srv, bt.ID, "Kiosco Básico")

	resp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-type-templates/"+tpl.ID, nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var wrapper templateWrapperResponse
	decodeJSON(t, resp, &wrapper)
	assert.Equal(t, tpl.ID, wrapper.Template.ID)
	assert.Equal(t, "Kiosco Básico", wrapper.Template.Name)
}

// TestListTemplates_FilterByBusinessTypeID_ReturnsCorrectSubset verifica el filtro por BT
func TestListTemplates_FilterByBusinessTypeID_ReturnsCorrectSubset(t *testing.T) {
	srv := newTestServer(t)
	bt1 := createTestBusinessType(t, srv, "supermercado")
	bt2 := createTestBusinessType(t, srv, "libreria")

	tpl1 := createTestTemplate(t, srv, bt1.ID, "Super Template")
	_ = createTestTemplate(t, srv, bt2.ID, "Librería Template")

	url := baseURL(srv) + "/business-type-templates?business_type_id=" + bt1.ID + "&include_inactive=true"
	resp := doRequest(t, http.MethodGet, url, nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var listed listTemplatesResponse
	decodeJSON(t, resp, &listed)

	for _, item := range listed.Items {
		assert.Equal(t, bt1.ID, item.BusinessTypeID, "solo deben aparecer templates del BT filtrado")
	}

	found := false
	for _, item := range listed.Items {
		if item.ID == tpl1.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "el template creado debe aparecer en el listado filtrado")
}

// TestUpdateTemplate_ValidPayload_Returns200 verifica la actualización de template
func TestUpdateTemplate_ValidPayload_Returns200(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "carniceria")
	tpl := createTestTemplate(t, srv, bt.ID, "Carnicería Original")

	payload := map[string]interface{}{
		"business_type_id": bt.ID,
		"name":             "Carnicería Actualizada",
		"description":      "Descripción actualizada",
	}

	resp := doRequest(t, http.MethodPut, baseURL(srv)+"/business-type-templates/"+tpl.ID, payload, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var wrapper templateWrapperResponse
	decodeJSON(t, resp, &wrapper)
	assert.Equal(t, "Carnicería Actualizada", wrapper.Template.Name)
}

// TestDeleteTemplate_ExistingID_Returns204 verifica la eliminación de template
func TestDeleteTemplate_ExistingID_Returns204(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "heladeria")
	tpl := createTestTemplate(t, srv, bt.ID, "Heladería Template")

	resp := doRequest(t, http.MethodDelete, baseURL(srv)+"/business-type-templates/"+tpl.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	resp.Body.Close()

	// Verificar que ya no existe
	getResp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-type-templates/"+tpl.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, getResp.StatusCode)
	getResp.Body.Close()
}

// TestGetTemplateAnalytics_ExistingTemplate_Returns200 verifica el endpoint de analytics
func TestGetTemplateAnalytics_ExistingTemplate_Returns200(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "almacen")
	tpl := createTestTemplate(t, srv, bt.ID, "Almacén Template")

	analyticsURL := baseURL(srv) + "/business-type-templates/" + tpl.ID + "/analytics"
	resp := doRequest(t, http.MethodGet, analyticsURL, nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var analytics templateAnalyticsResponse
	decodeJSON(t, resp, &analytics)
	assert.Equal(t, tpl.ID, analytics.TemplateID)
	assert.GreaterOrEqual(t, analytics.TenantsUsed, 0)
	assert.GreaterOrEqual(t, analytics.CompletionRate, float64(0))
}

// TestGetTemplateAnalytics_UnknownID_Returns404 verifica analytics de template inexistente
func TestGetTemplateAnalytics_UnknownID_Returns404(t *testing.T) {
	srv := newTestServer(t)

	analyticsURL := baseURL(srv) + "/business-type-templates/00000000-0000-0000-0000-000000000000/analytics"
	resp := doRequest(t, http.MethodGet, analyticsURL, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestDuplicateTemplate_ExistingTemplate_Returns201WithCopy verifica la duplicación
func TestDuplicateTemplate_ExistingTemplate_Returns201WithCopy(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "verduleria")
	original := createTestTemplate(t, srv, bt.ID, "Verdulería Básica")

	duplicateURL := baseURL(srv) + "/business-type-templates/" + original.ID + "/duplicate"
	resp := doRequest(t, http.MethodPost, duplicateURL, nil, adminHeader())
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var wrapper templateWrapperResponse
	decodeJSON(t, resp, &wrapper)
	assert.NotEmpty(t, wrapper.Template.ID)
	assert.NotEqual(t, original.ID, wrapper.Template.ID, "la copia debe tener ID diferente")
	assert.Equal(t, "Verdulería Básica (copia)", wrapper.Template.Name)
	assert.Equal(t, original.BusinessTypeID, wrapper.Template.BusinessTypeID)
	assert.False(t, wrapper.Template.IsDefault, "la copia nunca debe ser template por defecto")
}

// TestDuplicateTemplate_WithCustomName_Returns201WithGivenName verifica duplicación con nombre custom
func TestDuplicateTemplate_WithCustomName_Returns201WithGivenName(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "zapateria")
	original := createTestTemplate(t, srv, bt.ID, "Zapatería Original")

	duplicateURL := baseURL(srv) + "/business-type-templates/" + original.ID + "/duplicate"
	payload := map[string]interface{}{
		"new_name": "Mi Copia Personalizada",
	}
	resp := doRequest(t, http.MethodPost, duplicateURL, payload, adminHeader())
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var wrapper templateWrapperResponse
	decodeJSON(t, resp, &wrapper)
	assert.Equal(t, "Mi Copia Personalizada", wrapper.Template.Name)
}

// TestDuplicateTemplate_UnknownID_Returns404 verifica duplicación de template inexistente
func TestDuplicateTemplate_UnknownID_Returns404(t *testing.T) {
	srv := newTestServer(t)

	duplicateURL := baseURL(srv) + "/business-type-templates/00000000-0000-0000-0000-000000000000/duplicate"
	resp := doRequest(t, http.MethodPost, duplicateURL, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestCRUD_Template_FullLifecycle verifica create → get → update → delete → 404
func TestCRUD_Template_FullLifecycle(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "ciclo-template")

	// Create
	tpl := createTestTemplate(t, srv, bt.ID, "Template Ciclo Completo")
	assert.NotEmpty(t, tpl.ID)

	// Get
	getResp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-type-templates/"+tpl.ID, nil, adminHeader())
	assert.Equal(t, http.StatusOK, getResp.StatusCode)
	getResp.Body.Close()

	// Update
	updatePayload := map[string]interface{}{
		"business_type_id": bt.ID,
		"name":             "Template Actualizado",
		"description":      "Nueva descripción",
	}
	updResp := doRequest(t, http.MethodPut, baseURL(srv)+"/business-type-templates/"+tpl.ID, updatePayload, adminHeader())
	assert.Equal(t, http.StatusOK, updResp.StatusCode)
	var updWrapper templateWrapperResponse
	decodeJSON(t, updResp, &updWrapper)
	assert.Equal(t, "Template Actualizado", updWrapper.Template.Name)

	// Analytics (antes de borrar)
	analyticsResp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-type-templates/"+tpl.ID+"/analytics", nil, adminHeader())
	assert.Equal(t, http.StatusOK, analyticsResp.StatusCode)
	analyticsResp.Body.Close()

	// Duplicate
	dupResp := doRequest(t, http.MethodPost, baseURL(srv)+"/business-type-templates/"+tpl.ID+"/duplicate", nil, adminHeader())
	assert.Equal(t, http.StatusCreated, dupResp.StatusCode)
	var dupWrapper templateWrapperResponse
	decodeJSON(t, dupResp, &dupWrapper)
	assert.NotEqual(t, tpl.ID, dupWrapper.Template.ID)

	// Delete original
	delResp := doRequest(t, http.MethodDelete, baseURL(srv)+"/business-type-templates/"+tpl.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNoContent, delResp.StatusCode)
	delResp.Body.Close()

	// 404 after delete
	notFoundResp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-type-templates/"+tpl.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, notFoundResp.StatusCode)
	notFoundResp.Body.Close()

	// Copia todavía existe
	dupGetResp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-type-templates/"+dupWrapper.Template.ID, nil, adminHeader())
	assert.Equal(t, http.StatusOK, dupGetResp.StatusCode)
	dupGetResp.Body.Close()
}

// TestCascadeDelete_BusinessType_DeletesTemplates verifica que al borrar un BT se borran sus templates
func TestCascadeDelete_BusinessType_DeletesTemplates(t *testing.T) {
	srv := newTestServer(t)
	bt := createTestBusinessType(t, srv, "cascade")
	tpl1 := createTestTemplate(t, srv, bt.ID, "Template Cascade 1")
	tpl2 := createTestTemplate(t, srv, bt.ID, "Template Cascade 2")

	// Eliminar el BT padre
	delResp := doRequest(t, http.MethodDelete, baseURL(srv)+"/business-types/"+bt.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNoContent, delResp.StatusCode)
	delResp.Body.Close()

	// Los templates deben haberse borrado en cascade (FK ON DELETE CASCADE en migración 014)
	tpl1Resp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-type-templates/"+tpl1.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, tpl1Resp.StatusCode)
	tpl1Resp.Body.Close()

	tpl2Resp := doRequest(t, http.MethodGet, baseURL(srv)+"/business-type-templates/"+tpl2.ID, nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, tpl2Resp.StatusCode)
	tpl2Resp.Body.Close()
}
