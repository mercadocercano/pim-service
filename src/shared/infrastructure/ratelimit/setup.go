// Package ratelimit cablea el rate limiting por tenant/plan (ADR-003) en pim-service:
// construye el limiter (Redis), el provider que cachea la matriz desde iam, y el
// middleware Gin para la feature pim.bulk_import. Arranca en OBSERVE-ONLY (mide, no
// rechaza) salvo RATE_LIMIT_ENFORCE=true.
package ratelimit

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	gmw "github.com/hornosg/go-shared/infrastructure/middleware"
	grl "github.com/hornosg/go-shared/infrastructure/ratelimit"
)

const featureBulkImport = "pim.bulk_import"

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// NewBulkImportMiddleware construye el middleware de rate limiting para el import bulk.
// Debe montarse DESPUÉS de TenantValidation (depende de tenant_id/jwt_claims en contexto).
// Resiliente: si Redis/iam fallan, en observe-only o fail-open deja pasar (no rompe el import).
func NewBulkImportMiddleware() gin.HandlerFunc {
	redisAddr := env("REDIS_ADDR", "lab-redis:6379")
	iamURL := env("IAM_SERVICE_URL", "http://mc-iam-service:8080")
	apiKey := os.Getenv("S2S_API_KEY")
	enforce := os.Getenv("RATE_LIMIT_ENFORCE") == "true" // default: observe-only

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	limiter := grl.NewRedisLimiter(rdb)

	fetch := func(ctx context.Context) (grl.Matrix, error) {
		return fetchMatrixFromIAM(ctx, iamURL, apiKey)
	}
	provider := grl.NewRefreshingProvider(context.Background(), fetch, 5*time.Minute)

	log.Printf("[ratelimit] pim.bulk_import middleware listo (enforce=%v, redis=%s, iam=%s)", enforce, redisAddr, iamURL)

	return gmw.RateLimit(gmw.RateLimitConfig{
		Limiter:     limiter,
		Provider:    provider,
		Feature:     featureBulkImport,
		ObserveOnly: !enforce,
		OnLimitExceeded: func(c *gin.Context, feature string) {
			log.Printf("[ratelimit] LIMIT EXCEEDED tenant=%s feature=%s tier=%s enforce=%v",
				c.GetString("tenant_id"), feature, gmw.PlanTierFromContext(c), enforce)
		},
		OnBackendUnavailable: func(_ *gin.Context, err error) {
			log.Printf("[ratelimit] backend unavailable (fail-open): %v", err)
		},
	})
}

// fetchMatrixFromIAM lee la matriz de límites del endpoint de planes de iam (ADR-003 D4).
// Tolerante: features con regla inválida se omiten; un plan sin reglas no entra a la matriz.
func fetchMatrixFromIAM(ctx context.Context, baseURL, apiKey string) (grl.Matrix, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/api/v1/plans", nil)
	if err != nil {
		return nil, err
	}
	if apiKey != "" {
		req.Header.Set("X-API-Key", apiKey)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("iam /plans status %d", resp.StatusCode)
	}

	var body struct {
		Plans []struct {
			Type       string                `json:"type"`
			RateLimits map[string]grl.RawRule `json:"rate_limits"`
		} `json:"plans"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	m := grl.Matrix{}
	for _, p := range body.Plans {
		rules := map[string]grl.Rule{}
		for feature, raw := range p.RateLimits {
			rule, perr := grl.Parse(raw)
			if perr != nil {
				log.Printf("[ratelimit] regla inválida plan=%s feature=%s: %v", p.Type, feature, perr)
				continue
			}
			rules[feature] = rule
		}
		if len(rules) > 0 {
			m[p.Type] = rules
		}
	}
	return m, nil
}
