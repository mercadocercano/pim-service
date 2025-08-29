package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// ImportOperationsTotal tracks total number of import operations
	ImportOperationsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pim_import_operations_total",
			Help: "Total number of import operations",
		},
		[]string{"tenant_id", "import_type", "status"},
	)

	// ImportRecordsProcessed tracks number of records processed
	ImportRecordsProcessed = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pim_import_records_processed_total",
			Help: "Total number of records processed during imports",
		},
		[]string{"tenant_id", "import_type", "result"},
	)

	// ImportDuration tracks duration of import operations
	ImportDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "pim_import_duration_seconds",
			Help:    "Duration of import operations in seconds",
			Buckets: []float64{0.1, 0.5, 1, 2, 5, 10, 30, 60, 120, 300},
		},
		[]string{"tenant_id", "import_type"},
	)

	// ImportFileSizeBytes tracks size of imported files
	ImportFileSizeBytes = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "pim_import_file_size_bytes",
			Help:    "Size of imported files in bytes",
			Buckets: []float64{1024, 10240, 102400, 1048576, 10485760, 52428800, 104857600}, // 1KB to 100MB
		},
		[]string{"tenant_id", "import_type"},
	)

	// ImportValidationErrors tracks validation errors during import
	ImportValidationErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pim_import_validation_errors_total",
			Help: "Total number of validation errors during import",
		},
		[]string{"tenant_id", "import_type", "error_type", "field"},
	)

	// ImportBatchSize tracks batch sizes for batch operations
	ImportBatchSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "pim_import_batch_size",
			Help:    "Size of batches in batch import operations",
			Buckets: []float64{1, 10, 50, 100, 500, 1000, 5000, 10000},
		},
		[]string{"tenant_id", "entity_type"},
	)

	// SKUValidationDuration tracks duration of SKU validation operations
	SKUValidationDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "pim_sku_validation_duration_seconds",
			Help:    "Duration of SKU validation operations in seconds",
			Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2, 5},
		},
		[]string{"tenant_id"},
	)

	// SKUValidationBatchSize tracks number of SKUs validated per request
	SKUValidationBatchSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "pim_sku_validation_batch_size",
			Help:    "Number of SKUs validated per request",
			Buckets: []float64{1, 5, 10, 50, 100, 500, 1000},
		},
		[]string{"tenant_id"},
	)
)

// RecordImportMetrics is a helper function to record common import metrics
func RecordImportMetrics(tenantID, importType string, success bool, recordsProcessed, successCount, failureCount int, duration float64) {
	status := "success"
	if !success {
		status = "failure"
	}

	// Record operation
	ImportOperationsTotal.WithLabelValues(tenantID, importType, status).Inc()

	// Record duration
	ImportDuration.WithLabelValues(tenantID, importType).Observe(duration)

	// Record processed records
	if successCount > 0 {
		ImportRecordsProcessed.WithLabelValues(tenantID, importType, "success").Add(float64(successCount))
	}
	if failureCount > 0 {
		ImportRecordsProcessed.WithLabelValues(tenantID, importType, "failure").Add(float64(failureCount))
	}
}

// RecordValidationError records a validation error metric
func RecordValidationError(tenantID, importType, errorType, field string) {
	ImportValidationErrors.WithLabelValues(tenantID, importType, errorType, field).Inc()
}

// RecordFileSize records the size of an imported file
func RecordFileSize(tenantID, importType string, sizeBytes int64) {
	ImportFileSizeBytes.WithLabelValues(tenantID, importType).Observe(float64(sizeBytes))
}

// RecordBatchOperation records metrics for batch operations
func RecordBatchOperation(tenantID, entityType string, batchSize int) {
	ImportBatchSize.WithLabelValues(tenantID, entityType).Observe(float64(batchSize))
}

// RecordSKUValidation records metrics for SKU validation operations
func RecordSKUValidation(tenantID string, skuCount int, duration float64) {
	SKUValidationBatchSize.WithLabelValues(tenantID).Observe(float64(skuCount))
	SKUValidationDuration.WithLabelValues(tenantID).Observe(duration)
}