package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"saas-mt-pim-service/src/schema_validation/domain/entity"
)

// MemoryValidationCache implementa un cache en memoria para validaciones
type MemoryValidationCache struct {
	mu         sync.RWMutex
	store      map[string]*entity.SchemaValidation
	cleanupTTL time.Duration
}

// NewMemoryValidationCache crea una nueva instancia del cache
func NewMemoryValidationCache(cleanupInterval time.Duration) *MemoryValidationCache {
	cache := &MemoryValidationCache{
		store:      make(map[string]*entity.SchemaValidation),
		cleanupTTL: 30 * time.Minute,
	}
	
	// Iniciar limpieza periódica
	go cache.cleanupExpired(cleanupInterval)
	
	return cache
}

// Get obtiene una validación del cache
func (c *MemoryValidationCache) Get(ctx context.Context, id string) (*entity.SchemaValidation, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	validation, exists := c.store[id]
	if !exists {
		return nil, fmt.Errorf("validation not found")
	}
	
	// Verificar si ha expirado
	if validation.IsExpired() {
		return nil, fmt.Errorf("validation expired")
	}
	
	return validation, nil
}

// Set guarda una validación en el cache
func (c *MemoryValidationCache) Set(ctx context.Context, validation *entity.SchemaValidation) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.store[validation.ID.String()] = validation
	return nil
}

// Delete elimina una validación del cache
func (c *MemoryValidationCache) Delete(ctx context.Context, id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	delete(c.store, id)
	return nil
}

// cleanupExpired limpia las validaciones expiradas periódicamente
func (c *MemoryValidationCache) cleanupExpired(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		
		for id, validation := range c.store {
			if validation.IsExpired() || now.After(validation.ExpiresAt) {
				delete(c.store, id)
			}
		}
		
		c.mu.Unlock()
	}
}