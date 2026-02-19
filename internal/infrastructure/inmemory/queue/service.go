package queue

import (
	"context"
	"sync"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type InMemoryQueue struct {
	data []any
	lock sync.Mutex
}

func NewInMemoryQueue() (ports.QueueService, error) {
	service := &InMemoryQueue{
		data: make([]any, 0),
	}

	return service, nil
}

func (c *InMemoryQueue) RegisterQueue(queue string) error {
	return nil
}

func (c *InMemoryQueue) Publish(ctx context.Context, queue string, message any) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data = append(c.data, message)

	return nil
}

func (c *InMemoryQueue) Consume(ctx context.Context, queue string, handler func([]byte) error) error {
	return nil
}

func (c *InMemoryQueue) Stop() {}
