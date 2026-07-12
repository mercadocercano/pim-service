//go:build integration

package global_catalog_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- DTOs que reflejan la forma real del JSON ---

type globalProductResponse struct {
	ID           string                 `json:"id"`
	EAN          *string                `json:"ean"`
	Name         string                 `json:"name"`
	Description  *string                `json:"description"`
	Brand        *string                `json:"brand"`
	Category     *string                `json:"category"`
	Price        *float64               `json:"price"`
	ImageURL     *string                `json:"image_url"`
	ImageURLs    []string               `json:"image_urls"`
	Source       string                 `json:"source"`
	QualityScore int                    `json:"quality_score"`
	IsVerified   bool                   `json:"is_verified"`
	IsActive     bool                   `json:"is_active"`
	BusinessType *string                `json:"business_type"`
	Tags         []string               `json:"tags"`
	Metadata     map[string]interface{} `json:"metadata"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
}

type getProductResponse struct {
	Product globalProductResponse `json:"product"`
}

type listProductsResponse struct {
	Items      []globalProductResponse `json:"items"`
	TotalCount int                     `json:"total_count"`
	Page       int                     `json:"page"`
	PageSize   int                     `json:"page_size"`
	TotalPages int                     `json:"total_pages"`
}

type bulkImportError struct {
	Row     int    `json:"row"`
	Message string `json:"message"`
}

type bulkImportResponse struct {
	ImportedCount int               `json:"imported_count"`
	FailedCount   int               `json:"failed_count"`
	Errors        []bulkImportError `json:"errors"`
}

// --- Helpers ---

func adminHeader() http.Header {
	h := http.Header{}
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

// createProductPayload genera un payload mínimo válido para crear un producto sin EAN.
func createProductPayload(name string) map[string]interface{} {
	return map[string]interface{}{
		"name":   name,
		"source": "integration_test",
	}
}

// createProductAndGetID crea un producto y devuelve su ID.
func createProductAndGetID(t *testing.T, srv *testServer, name string) string {
	t.Helper()
	resp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products",
		createProductPayload(name), adminHeader())
	require.Equal(t, http.StatusCreated, resp.StatusCode, "crear producto %q debe retornar 201", name)

	var created globalProductResponse
	decodeJSON(t, resp, &created)
	require.NotEmpty(t, created.ID)
	return created.ID
}

// --- Tests ---

// TestGlobalCatalog_CRUDFlow verifica el flujo completo: crear → obtener → listar → actualizar → obtener → eliminar → obtener 404.
func TestGlobalCatalog_CRUDFlow(t *testing.T) {
	srv := newTestServer(t)

	// 1. Crear
	createPayload := map[string]interface{}{
		"name":          "Tornillo M6x20 Integration",
		"source":        "integration_test",
		"business_type": "ferreteria",
	}
	createResp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products", createPayload, adminHeader())
	assert.Equal(t, http.StatusCreated, createResp.StatusCode)

	var created globalProductResponse
	decodeJSON(t, createResp, &created)
	assert.NotEmpty(t, created.ID, "id debe estar presente")
	assert.Equal(t, "Tornillo M6x20 Integration", created.Name)
	assert.False(t, created.IsVerified, "is_verified por defecto debe ser false")
	assert.True(t, created.IsActive, "is_active por defecto debe ser true")

	productID := created.ID

	// 2. Obtener por ID
	getResp := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), productID),
		nil, adminHeader())
	assert.Equal(t, http.StatusOK, getResp.StatusCode)

	var getResult getProductResponse
	decodeJSON(t, getResp, &getResult)
	assert.Equal(t, productID, getResult.Product.ID)
	assert.Equal(t, "Tornillo M6x20 Integration", getResult.Product.Name)

	// 3. Listar y verificar que aparece
	listResp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/global-catalog/products?page=1&page_size=50",
		nil, adminHeader())
	assert.Equal(t, http.StatusOK, listResp.StatusCode)

	var listResult listProductsResponse
	decodeJSON(t, listResp, &listResult)
	assert.GreaterOrEqual(t, listResult.TotalCount, 1)

	foundInList := false
	for _, p := range listResult.Items {
		if p.ID == productID {
			foundInList = true
			break
		}
	}
	assert.True(t, foundInList, "el producto creado debe aparecer en el listado")

	// 4. Actualizar
	updatePayload := map[string]interface{}{
		"name":     "Tornillo M6x20 Actualizado",
		"category": "Ferretería > Tornillos",
	}
	updateResp := doRequest(t, http.MethodPut,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), productID),
		updatePayload, adminHeader())
	assert.Equal(t, http.StatusOK, updateResp.StatusCode)

	// 5. Obtener de nuevo y verificar actualización
	getResp2 := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), productID),
		nil, adminHeader())
	assert.Equal(t, http.StatusOK, getResp2.StatusCode)

	var getResult2 getProductResponse
	decodeJSON(t, getResp2, &getResult2)
	assert.Equal(t, "Tornillo M6x20 Actualizado", getResult2.Product.Name)

	// 6. Eliminar
	deleteResp := doRequest(t, http.MethodDelete,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), productID),
		nil, adminHeader())
	assert.Equal(t, http.StatusNoContent, deleteResp.StatusCode)
	deleteResp.Body.Close()

	// 7. Obtener → verificar que está eliminado (soft delete → is_active=false)
	// El endpoint retorna 200 con el producto pero is_active=false (soft delete),
	// o 404 dependiendo de la implementación. Verificamos ambos casos.
	getAfterDelete := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), productID),
		nil, adminHeader())

	if getAfterDelete.StatusCode == http.StatusOK {
		var afterDelete getProductResponse
		decodeJSON(t, getAfterDelete, &afterDelete)
		assert.False(t, afterDelete.Product.IsActive, "producto eliminado debe tener is_active=false")
	} else {
		assert.Equal(t, http.StatusNotFound, getAfterDelete.StatusCode,
			"producto eliminado debe retornar 404 o bien is_active=false")
		getAfterDelete.Body.Close()
	}
}

// TestGlobalCatalog_VerifyUnverifyFlow verifica el ciclo verificar → desverificar.
func TestGlobalCatalog_VerifyUnverifyFlow(t *testing.T) {
	srv := newTestServer(t)

	// Crear producto
	productID := createProductAndGetID(t, srv, "Producto Para Verificar")

	// Verificar que empieza no verificado
	getResp := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), productID),
		nil, adminHeader())
	require.Equal(t, http.StatusOK, getResp.StatusCode)
	var initial getProductResponse
	decodeJSON(t, getResp, &initial)
	assert.False(t, initial.Product.IsVerified, "is_verified debe ser false inicialmente")

	// Verificar
	verifyResp := doRequest(t, http.MethodPatch,
		fmt.Sprintf("%s/global-catalog/products/%s/verify", baseURL(srv), productID),
		nil, adminHeader())
	assert.Equal(t, http.StatusOK, verifyResp.StatusCode)
	verifyResp.Body.Close()

	// Obtener y verificar is_verified=true
	getResp2 := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), productID),
		nil, adminHeader())
	require.Equal(t, http.StatusOK, getResp2.StatusCode)
	var afterVerify getProductResponse
	decodeJSON(t, getResp2, &afterVerify)
	assert.True(t, afterVerify.Product.IsVerified, "is_verified debe ser true tras verify")

	// Desverificar
	unverifyResp := doRequest(t, http.MethodPatch,
		fmt.Sprintf("%s/global-catalog/products/%s/unverify", baseURL(srv), productID),
		nil, adminHeader())
	assert.Equal(t, http.StatusOK, unverifyResp.StatusCode)
	unverifyResp.Body.Close()

	// Obtener y verificar is_verified=false
	getResp3 := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), productID),
		nil, adminHeader())
	require.Equal(t, http.StatusOK, getResp3.StatusCode)
	var afterUnverify getProductResponse
	decodeJSON(t, getResp3, &afterUnverify)
	assert.False(t, afterUnverify.Product.IsVerified, "is_verified debe ser false tras unverify")
}

// TestGlobalCatalog_Verify_NotFound verifica 404 para ID inexistente.
func TestGlobalCatalog_Verify_NotFound(t *testing.T) {
	srv := newTestServer(t)

	resp := doRequest(t, http.MethodPatch,
		baseURL(srv)+"/global-catalog/products/00000000-0000-0000-0000-000000000000/verify",
		nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestGlobalCatalog_BulkImport_MixedRows verifica que bulk-import maneja filas válidas e inválidas.
// Una fila es inválida si el nombre está vacío (validación de dominio).
func TestGlobalCatalog_BulkImport_MixedRows(t *testing.T) {
	srv := newTestServer(t)

	payload := map[string]interface{}{
		"source": "bulk_test",
		"rows": []map[string]interface{}{
			{"name": "Producto Bulk 1", "business_type": "almacen"},
			{"name": "Producto Bulk 2", "business_type": "kiosco"},
			{"name": "Producto Bulk 3", "business_type": "ferreteria"},
			// Fila inválida: nombre vacío (validación de dominio)
			{"name": "", "business_type": "almacen"},
			// Fila inválida: source vacío no aplica (el use case lo hereda del request)
			// Pero podemos probar con un EAN inválido (no 13 dígitos)
			{"name": "Producto Con EAN Invalido", "ean": "123", "business_type": "kiosco"},
		},
	}

	resp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products/bulk-import",
		payload, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result bulkImportResponse
	decodeJSON(t, resp, &result)

	assert.Equal(t, 3, result.ImportedCount, "deben importarse 3 productos válidos")
	assert.Equal(t, 2, result.FailedCount, "deben fallar 2 filas inválidas")
	assert.Len(t, result.Errors, 2, "errors debe tener 2 entradas")

	// Verificar que los errores tienen número de fila correcto
	errorRows := make(map[int]bool)
	for _, e := range result.Errors {
		errorRows[e.Row] = true
		assert.NotEmpty(t, e.Message, "cada error debe tener mensaje")
	}
	assert.True(t, errorRows[4], "fila 4 (nombre vacío) debe estar en errores")
	assert.True(t, errorRows[5], "fila 5 (EAN inválido) debe estar en errores")
}

// TestGlobalCatalog_BulkImport_AllValid verifica que bulk-import con todas filas válidas retorna correctamente.
func TestGlobalCatalog_BulkImport_AllValid(t *testing.T) {
	srv := newTestServer(t)

	rows := make([]map[string]interface{}, 5)
	for i := 0; i < 5; i++ {
		rows[i] = map[string]interface{}{
			"name":   fmt.Sprintf("Producto Masivo %d", i+1),
			"source": "bulk_valid_test",
		}
	}

	payload := map[string]interface{}{
		"source": "bulk_valid_test",
		"rows":   rows,
	}

	resp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products/bulk-import",
		payload, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result bulkImportResponse
	decodeJSON(t, resp, &result)

	assert.Equal(t, 5, result.ImportedCount)
	assert.Equal(t, 0, result.FailedCount)
	assert.Empty(t, result.Errors)
}

// TestGlobalCatalog_BulkImport_EmptyRows verifica que una lista vacía retorna error 400.
func TestGlobalCatalog_BulkImport_EmptyRows(t *testing.T) {
	srv := newTestServer(t)

	payload := map[string]interface{}{
		"source": "test",
		"rows":   []interface{}{},
	}

	resp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products/bulk-import",
		payload, adminHeader())
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp.Body.Close()
}

// TestGlobalCatalog_Delete_ProtectedByTenantUse verifica que el delete retorna 409 si el producto está en uso.
func TestGlobalCatalog_Delete_ProtectedByTenantUse(t *testing.T) {
	srv := newTestServer(t)

	// Crear producto
	productID := createProductAndGetID(t, srv, "Producto En Uso Por Tenant")

	// Simular que un tenant lo tiene en su catálogo insertando en tenant_global_product_links
	_, err := srv.DB.ExecContext(context.Background(), `
		INSERT INTO tenant_global_product_links (id, tenant_id, tenant_product_id, global_product_id)
		VALUES (
			gen_random_uuid(),
			gen_random_uuid(),
			gen_random_uuid(),
			$1
		)
	`, productID)
	require.NoError(t, err, "debe poder insertar el link tenant-producto")

	// Intentar eliminar → debe retornar 409
	deleteResp := doRequest(t, http.MethodDelete,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), productID),
		nil, adminHeader())
	assert.Equal(t, http.StatusConflict, deleteResp.StatusCode,
		"DELETE de producto en uso debe retornar 409 Conflict")

	var errBody map[string]interface{}
	decodeJSON(t, deleteResp, &errBody)
	assert.Contains(t, fmt.Sprint(errBody["error"]), "en uso",
		"mensaje de error debe indicar que el producto está en uso")
}

// TestGlobalCatalog_Delete_NotFound verifica 404 para ID inexistente.
func TestGlobalCatalog_Delete_NotFound(t *testing.T) {
	srv := newTestServer(t)

	resp := doRequest(t, http.MethodDelete,
		baseURL(srv)+"/global-catalog/products/00000000-0000-0000-0000-000000000000",
		nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestGlobalCatalog_GetByID_NotFound verifica 404 para ID inexistente.
func TestGlobalCatalog_GetByID_NotFound(t *testing.T) {
	srv := newTestServer(t)

	resp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/global-catalog/products/00000000-0000-0000-0000-000000000000",
		nil, adminHeader())
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

// TestGlobalCatalog_Create_MissingName verifica validación: nombre obligatorio.
func TestGlobalCatalog_Create_MissingName(t *testing.T) {
	srv := newTestServer(t)

	payload := map[string]interface{}{
		"source": "test",
	}

	resp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products", payload, adminHeader())
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp.Body.Close()
}

// TestGlobalCatalog_Create_MissingSource verifica validación: source obligatorio.
func TestGlobalCatalog_Create_MissingSource(t *testing.T) {
	srv := newTestServer(t)

	payload := map[string]interface{}{
		"name": "Producto Sin Source",
	}

	resp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products", payload, adminHeader())
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	resp.Body.Close()
}

// TestGlobalCatalog_Create_DuplicateEAN verifica que EANs duplicados dan 409.
// Usamos un EAN distinto al de las semillas (7791234567890 ya está en la BD por la migración inicial).
func TestGlobalCatalog_Create_DuplicateEAN(t *testing.T) {
	srv := newTestServer(t)

	// EAN válido de 13 dígitos con checksum correcto, no presente en semillas de migración
	payload := map[string]interface{}{
		"name":   "Producto EAN Test",
		"ean":    "7790000000003",
		"source": "test",
	}

	resp1 := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products", payload, adminHeader())
	require.Equal(t, http.StatusCreated, resp1.StatusCode,
		"primera creación debe retornar 201")
	resp1.Body.Close()

	resp2 := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products", payload, adminHeader())
	assert.Equal(t, http.StatusConflict, resp2.StatusCode,
		"EAN duplicado debe retornar 409 Conflict")
	resp2.Body.Close()
}

// TestGlobalCatalog_List_Pagination verifica paginación básica.
func TestGlobalCatalog_List_Pagination(t *testing.T) {
	srv := newTestServer(t)

	// Crear 4 productos
	for i := 1; i <= 4; i++ {
		createProductAndGetID(t, srv, fmt.Sprintf("Producto Pagina %d", i))
	}

	// Pedir page_size=2
	resp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/global-catalog/products?page=1&page_size=2",
		nil, adminHeader())
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result listProductsResponse
	decodeJSON(t, resp, &result)

	assert.Len(t, result.Items, 2, "page_size=2 debe retornar 2 items")
	assert.GreaterOrEqual(t, result.TotalCount, 4, "total_count debe ser al menos 4")
}

// --- Bulk verify + filtros de revisión ---

type bulkVerifyResponse struct {
	Mode        string                 `json:"mode"`
	SnapshotRef *string                `json:"snapshot_ref"`
	Summary     map[string]interface{} `json:"summary"`
}

func adminHeaderWithOperator(operatorID string) http.Header {
	h := adminHeader()
	h.Set("X-Operator-Id", operatorID)
	return h
}

func setProductQualityAndImage(t *testing.T, srv *testServer, productID string, qualityScore int, imageURL *string) {
	t.Helper()
	var img interface{}
	if imageURL != nil {
		img = *imageURL
	} else {
		img = nil
	}
	_, err := srv.DB.ExecContext(context.Background(),
		"UPDATE global_products SET quality_score = $1, image_url = $2 WHERE id = $3",
		qualityScore, img, productID)
	require.NoError(t, err)
}

// TestGlobalCatalog_BulkVerify_VerifyAndUnverify verifica la verificación y desverificación en lote.
func TestGlobalCatalog_BulkVerify_VerifyAndUnverify(t *testing.T) {
	srv := newTestServer(t)

	id1 := createProductAndGetID(t, srv, "Producto Bulk Verify 1")
	id2 := createProductAndGetID(t, srv, "Producto Bulk Verify 2")
	id3 := createProductAndGetID(t, srv, "Producto Bulk Verify 3")

	// Verificar 2 de los 3
	verifyResp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products/bulk-verify",
		map[string]interface{}{
			"ids":    []string{id1, id2},
			"verify": true,
		},
		adminHeaderWithOperator("operator-bulk-1"))
	if verifyResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(verifyResp.Body)
		verifyResp.Body.Close()
		t.Fatalf("bulk verify retornó %d: %s", verifyResp.StatusCode, string(body))
	}

	var verifyResult bulkVerifyResponse
	decodeJSON(t, verifyResp, &verifyResult)
	assert.Equal(t, "verify", verifyResult.Mode)
	assert.Equal(t, 2, int(verifyResult.Summary["processed"].(float64)))
	assert.Equal(t, 0, int(verifyResult.Summary["failed"].(float64)))

	// Listar solo verificados
	listVerifiedResp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/global-catalog/products?is_verified=true",
		nil, adminHeader())
	require.Equal(t, http.StatusOK, listVerifiedResp.StatusCode)

	var verifiedList listProductsResponse
	decodeJSON(t, listVerifiedResp, &verifiedList)
	verifiedIDs := make(map[string]bool)
	for _, p := range verifiedList.Items {
		verifiedIDs[p.ID] = true
	}
	assert.True(t, verifiedIDs[id1], "id1 debe estar verificado")
	assert.True(t, verifiedIDs[id2], "id2 debe estar verificado")
	assert.False(t, verifiedIDs[id3], "id3 no debe estar verificado")

	// Desverificar los 2
	unverifyResp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products/bulk-verify",
		map[string]interface{}{
			"ids":    []string{id1, id2},
			"verify": false,
		},
		adminHeaderWithOperator("operator-bulk-1"))
	if unverifyResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(unverifyResp.Body)
		unverifyResp.Body.Close()
		t.Fatalf("bulk unverify retornó %d: %s", unverifyResp.StatusCode, string(body))
	}

	var unverifyResult bulkVerifyResponse
	decodeJSON(t, unverifyResp, &unverifyResult)
	assert.Equal(t, "unverify", unverifyResult.Mode)
	assert.Equal(t, 2, int(unverifyResult.Summary["processed"].(float64)))

	// Verificar que id1 ahora está desverificado
	getResp := doRequest(t, http.MethodGet,
		fmt.Sprintf("%s/global-catalog/products/%s", baseURL(srv), id1),
		nil, adminHeader())
	require.Equal(t, http.StatusOK, getResp.StatusCode)

	var afterUnverify getProductResponse
	decodeJSON(t, getResp, &afterUnverify)
	assert.False(t, afterUnverify.Product.IsVerified, "id1 debe quedar desverificado")
}

// TestGlobalCatalog_BulkVerify_MissingOperatorID verifica que el endpoint exige X-Operator-Id.
func TestGlobalCatalog_BulkVerify_MissingOperatorID(t *testing.T) {
	srv := newTestServer(t)

	productID := createProductAndGetID(t, srv, "Producto Sin Operador")

	resp := doRequest(t, http.MethodPost,
		baseURL(srv)+"/global-catalog/products/bulk-verify",
		map[string]interface{}{
			"ids":    []string{productID},
			"verify": true,
		},
		adminHeader())
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var body map[string]interface{}
	decodeJSON(t, resp, &body)
	err, ok := body["error"].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "MISSING_OPERATOR_ID", err["code"])
}

// TestGlobalCatalog_List_QualityScoreAndImageFilters verifica filtros por quality_score y presencia de imagen.
func TestGlobalCatalog_List_QualityScoreAndImageFilters(t *testing.T) {
	srv := newTestServer(t)

	idHigh := createProductAndGetID(t, srv, "Producto Calidad Alta")
	idMedium := createProductAndGetID(t, srv, "Producto Calidad Media")
	idLow := createProductAndGetID(t, srv, "Producto Calidad Baja")
	idNoImage := createProductAndGetID(t, srv, "Producto Sin Imagen")

	imageURL := "http://example.com/img.png"
	setProductQualityAndImage(t, srv, idHigh, 85, &imageURL)
	setProductQualityAndImage(t, srv, idMedium, 55, &imageURL)
	setProductQualityAndImage(t, srv, idLow, 25, &imageURL)
	setProductQualityAndImage(t, srv, idNoImage, 60, nil)

	// Filtro quality_score >= 70
	highResp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/global-catalog/products?min_quality=70&max_quality=100",
		nil, adminHeader())
	require.Equal(t, http.StatusOK, highResp.StatusCode)

	var highList listProductsResponse
	decodeJSON(t, highResp, &highList)
	highIDs := make(map[string]bool)
	for _, p := range highList.Items {
		highIDs[p.ID] = true
	}
	assert.True(t, highIDs[idHigh], "calidad alta debe estar en el rango alto")
	assert.False(t, highIDs[idMedium], "calidad media no debe estar en el rango alto")
	assert.False(t, highIDs[idLow], "calidad baja no debe estar en el rango alto")

	// Filtro sin imagen
	noImageResp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/global-catalog/products?has_image=false",
		nil, adminHeader())
	require.Equal(t, http.StatusOK, noImageResp.StatusCode)

	var noImageList listProductsResponse
	decodeJSON(t, noImageResp, &noImageList)
	noImageIDs := make(map[string]bool)
	for _, p := range noImageList.Items {
		noImageIDs[p.ID] = true
	}
	assert.True(t, noImageIDs[idNoImage], "producto sin imagen debe aparecer")
	assert.False(t, noImageIDs[idHigh], "producto con imagen no debe aparecer")
}

// TestGlobalCatalog_List_MassVerifiedFilter verifica el filtro de productos mass-verificados.
func TestGlobalCatalog_List_MassVerifiedFilter(t *testing.T) {
	srv := newTestServer(t)

	massVerifiedID := createProductAndGetID(t, srv, "Producto Mass Verified")
	normalID := createProductAndGetID(t, srv, "Producto Normal")

	_, err := srv.DB.ExecContext(context.Background(),
		"UPDATE global_products SET mass_verified_at = NOW() WHERE id = $1",
		massVerifiedID)
	require.NoError(t, err)

	resp := doRequest(t, http.MethodGet,
		baseURL(srv)+"/global-catalog/products?mass_verified_only=true",
		nil, adminHeader())
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var result listProductsResponse
	decodeJSON(t, resp, &result)
	resultIDs := make(map[string]bool)
	for _, p := range result.Items {
		resultIDs[p.ID] = true
	}
	assert.True(t, resultIDs[massVerifiedID], "mass verified debe aparecer")
	assert.False(t, resultIDs[normalID], "producto normal no debe aparecer")
}
