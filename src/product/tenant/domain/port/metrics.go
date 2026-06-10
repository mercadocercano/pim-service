package port

// Metric event name constants for the product/tenant bounded context.
// Convention: <domain>.<entity>.<action> — follows ADR-002 from go-shared.
const (
	MetricImportOperation  = "product.import.operation"
	MetricImportRecord     = "product.import.record"
	MetricImportDuration   = "product.import.duration_s"
	MetricImportFileSize   = "product.import.file_size_bytes"
	MetricImportValidation = "product.import.validation_error"
	MetricImportBatch      = "product.import.batch_size"
	MetricSKUValidation    = "product.sku.validation_duration_s"
	MetricSKUBatchSize     = "product.sku.validation_batch_size"
)
