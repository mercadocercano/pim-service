package response

// TenantCustomAttributesListResponse representa la respuesta de una lista de atributos personalizados de un tenant
type TenantCustomAttributesListResponse struct {
	TenantID         string                    `json:"tenant_id"`
	TotalCount       int                       `json:"total_count"`
	CustomAttributes []CustomAttributeResponse `json:"custom_attributes"`
	FilteredBy       map[string]interface{}    `json:"filtered_by,omitempty"`
}
