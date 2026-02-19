package ports

import "context"

// QueueService interface defines the methods for interacting with a queue service.
type QueueService interface {
	RegisterQueue(queue string) error
	Publish(ctx context.Context, queue string, message any) error
	Consume(ctx context.Context, queue string, handler func([]byte) error) error
	Stop()
}
