package cache

import (
	"context"
	"sync"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type InMemoryCache struct {
	data map[string]any
	lock sync.Mutex
}

func NewInMemoryCache() (ports.CacheService, error) {
	service := &InMemoryCache{
		data: make(map[string]any),
	}

	return service, nil
}

func (c *InMemoryCache) Set(ctx context.Context, key string, value any) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[key] = value

	return nil
}
func (c *InMemoryCache) Get(ctx context.Context, key string) (any, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if val, ok := c.data[key]; !ok {
		return val, nil
	}

	return nil, nil
}
func (c *InMemoryCache) Delete(ctx context.Context, key string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.data[key]; !ok {
		delete(c.data, key)
	}

	return nil
}
