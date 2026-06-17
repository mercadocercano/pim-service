package logging_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"saas-mt-pim-service/src/pim/domain/port"
	pimlog "saas-mt-pim-service/src/pim/infrastructure/logging"

	"github.com/stretchr/testify/assert"
)

// ADR-001: cada evento produce UNA línea JSON canónica con envelope ts/level/service/event.
func parseLine(t *testing.T, b []byte) map[string]any {
	t.Helper()
	lines := bytes.Split(bytes.TrimSpace(b), []byte("\n"))
	assert.Len(t, lines, 1, "debe ser exactamente una línea por evento")
	var m map[string]any
	assert.NoError(t, json.Unmarshal(lines[0], &m))
	return m
}

func TestPIMLogger_ImportCompleted_EnvelopeAndInfoLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := pimlog.NewPIMLoggerWithWriter("pim-test", &buf)

	logger.Log(port.PIMEvent{
		Event:     "pim.import_from_global_catalog_completed",
		TenantID:  "t-abc",
		ProductID: "p-001",
	})

	line := parseLine(t, buf.Bytes())
	assert.Equal(t, "pim.import_from_global_catalog_completed", line["event"])
	assert.Equal(t, "info", line["level"])
	assert.Equal(t, "pim-test", line["service"])
	assert.NotEmpty(t, line["ts"], "ts (RFC3339 UTC) siempre presente")
	assert.Equal(t, "t-abc", line["tenant_id"])
	assert.Equal(t, "p-001", line["product_id"])
}

func TestPIMLogger_ImportFailed_ErrorLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := pimlog.NewPIMLoggerWithWriter("pim-test", &buf)

	logger.Log(port.PIMEvent{
		Event:    "pim.import_failed",
		TenantID: "t-abc",
		JobID:    "j-999",
		Reason:   "3/10 records failed",
	})

	line := parseLine(t, buf.Bytes())
	assert.Equal(t, "error", line["level"])
	assert.Equal(t, "pim.import_failed", line["event"])
	assert.Equal(t, "j-999", line["job_id"])
	assert.Equal(t, "3/10 records failed", line["reason"])
}

func TestPIMLogger_BackfillError_WarnLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := pimlog.NewPIMLoggerWithWriter("pim-test", &buf)

	logger.Log(port.PIMEvent{
		Event:     "pim.backfill_product_error",
		TenantID:  "t-xyz",
		ProductID: "p-bad",
		Reason:    "sql: no rows",
	})

	line := parseLine(t, buf.Bytes())
	assert.Equal(t, "warn", line["level"])
	assert.Equal(t, "pim.backfill_product_error", line["event"])
}

func TestPIMLogger_OmitsEmptyFields(t *testing.T) {
	var buf bytes.Buffer
	logger := pimlog.NewPIMLoggerWithWriter("pim-test", &buf)

	logger.Log(port.PIMEvent{
		Event:    "pim.template_refresh_completed",
		TenantID: "t-1",
		Count:    5,
	})

	line := parseLine(t, buf.Bytes())
	assert.Equal(t, float64(5), line["count"])
	// campos vacíos omitidos
	_, hasJobID := line["job_id"]
	assert.False(t, hasJobID, "job_id vacío debe omitirse")
	_, hasSKU := line["sku"]
	assert.False(t, hasSKU, "sku vacío debe omitirse")
	_, hasProductID := line["product_id"]
	assert.False(t, hasProductID, "product_id vacío debe omitirse")
}

func TestPIMLogger_CountZeroOmitted(t *testing.T) {
	var buf bytes.Buffer
	logger := pimlog.NewPIMLoggerWithWriter("pim-test", &buf)

	logger.Log(port.PIMEvent{
		Event:    "pim.backfill_completed",
		TenantID: "t-1",
		// Count == 0 — debe omitirse
	})

	line := parseLine(t, buf.Bytes())
	assert.Equal(t, "info", line["level"])
	_, hasCount := line["count"]
	assert.False(t, hasCount, "count=0 debe omitirse")
}
