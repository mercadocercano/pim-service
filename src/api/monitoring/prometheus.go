package monitoring

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// HTTP metrics
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status", "tenant_id"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint", "tenant_id"},
	)

	// Database metrics
	databaseConnectionsActive = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "database_connections_active",
			Help: "Number of active database connections",
		},
	)

	databaseQueriesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "database_queries_total",
			Help: "Total number of database queries",
		},
		[]string{"operation", "table"},
	)

	// Business metrics - PIM specific
	productsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "products_total",
			Help: "Total number of products",
		},
		[]string{"tenant_id", "category"},
	)

	productOperationsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "product_operations_total",
			Help: "Total number of product operations",
		},
		[]string{"tenant_id", "operation"},
	)
)

func init() {
	// Register metrics
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(databaseConnectionsActive)
	prometheus.MustRegister(databaseQueriesTotal)
	prometheus.MustRegister(productsTotal)
	prometheus.MustRegister(productOperationsTotal)
}

// PrometheusMiddleware middleware para capturar métricas HTTP
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Obtener tenant_id del header o contexto
		tenantID := c.GetHeader("X-Tenant-ID")
		if tenantID == "" {
			tenantID = "unknown"
		}

		c.Next()

		duration := time.Since(start)
		status := strconv.Itoa(c.Writer.Status())

		httpRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			status,
			tenantID,
		).Inc()

		httpRequestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			tenantID,
		).Observe(duration.Seconds())
	}
}

// RecordDatabaseQuery registra métricas de consultas a la base de datos
func RecordDatabaseQuery(operation, table string) {
	databaseQueriesTotal.WithLabelValues(operation, table).Inc()
}

// SetDatabaseConnections actualiza el número de conexiones activas
func SetDatabaseConnections(count float64) {
	databaseConnectionsActive.Set(count)
}

// SetProductsCount actualiza el número de productos por tenant y categoría
func SetProductsCount(tenantID, category string, count float64) {
	productsTotal.WithLabelValues(tenantID, category).Set(count)
}

// RecordProductOperation registra operaciones de productos
func RecordProductOperation(tenantID, operation string) {
	productOperationsTotal.WithLabelValues(tenantID, operation).Inc()
}

// StartPrometheusServer inicia el servidor de métricas si está habilitado
func StartPrometheusServer() {
	enabled := os.Getenv("PROMETHEUS_ENABLED")
	if enabled != "true" {
		return
	}

	port := os.Getenv("PROMETHEUS_PORT")
	if port == "" {
		port = "2113"
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			panic("Failed to start Prometheus metrics server: " + err.Error())
		}
	}()
}
