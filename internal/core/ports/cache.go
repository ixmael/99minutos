package ports

import "context"

// CacheService interface defines the methods for interacting with a cache service.
type CacheService interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string) (any, error)
	Delete(ctx context.Context, key string) error
}
