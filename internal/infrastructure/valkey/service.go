package valkey

import (
	"context"
	"encoding/json"
	"time"

	"github.com/valkey-io/valkey-go"

	"github.com/ixmael/99minutos/internal/core/ports"
)

type valkeycache struct {
	client valkey.Client
	logger ports.Logger
}

func NewValkeyService(valkeyURL string, logger ports.Logger) (ports.CacheService, error) {
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{
			valkeyURL,
		},
	})
	if err != nil {
		return nil, err
	}

	valkeyCache := valkeycache{
		client: client,
		logger: logger,
	}

	return &valkeyCache, nil
}

func (c *valkeycache) Set(ctx context.Context, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	setCommand := c.client.B().Set().Key(key).Value(valkey.BinaryString(data)).Ex(24 * time.Hour).Build()
	return c.client.Do(ctx, setCommand).Error()
}

func (c *valkeycache) Get(ctx context.Context, key string) (any, error) {
	getCommand := c.client.B().Get().Key(key).Build()
	resp, err := c.client.Do(ctx, getCommand).ToString()
	if err != nil && !valkey.IsValkeyNil(err) {
		return nil, err
	}

	return resp, nil
}

func (c *valkeycache) Delete(ctx context.Context, key string) error {
	return nil
}
