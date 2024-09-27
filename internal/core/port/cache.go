package port

import (
	"context"
	"time"
)

//go:generate mockgen -source=cache.go -destination=mock/cache.go -package=mock

// CacheRepository is an interface for interacting with cache-related business logic
type CacheRepository interface {
	// Set stores the value in the cache
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	// Get retrieves the value from the cache
	Get(ctx context.Context, key string) ([]byte, error)
	// Close closes the connection to the cache server
	Close() error
}
