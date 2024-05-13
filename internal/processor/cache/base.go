package cache

import (
	"context"
	"encoding/json"
	"github.com/doxanocap/pkg/errs"
	"push-api/internal/interfaces"
	"time"
)

type Cache struct {
	provider interfaces.ICacheProvider
}

func NewCacheProcessor(provider interfaces.ICacheProvider) *Cache {
	return &Cache{
		provider: provider,
	}
}

func (c *Cache) Set(ctx context.Context, key string, value []byte) error {
	err := c.provider.Set(ctx, key, value)
	if err != nil {
		return errs.Wrap("cache.processor.Set", err)
	}
	return nil
}

func (c *Cache) SetJSON(ctx context.Context, key string, value interface{}) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return errs.Wrap("marshal", err)
	}

	err = c.provider.Set(ctx, key, raw)
	if err != nil {
		return errs.Wrap("cache.processor.SetJSON", err)
	}
	return nil
}

func (c *Cache) SetJSONWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return errs.Wrap("marshal", err)
	}

	err = c.provider.SetWithTTL(ctx, key, raw, ttl)
	if err != nil {
		return errs.Wrap("cache.processor.SetJSONWithTTL", err)
	}
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	raw, err := c.provider.Get(ctx, key)
	if err != nil {
		return nil, errs.Wrap("cache.processor.Get", err)
	}
	return raw, nil
}

func (c *Cache) GetJSON(ctx context.Context, key string, v interface{}) error {
	raw, err := c.provider.Get(ctx, key)
	if err != nil {
		return errs.Wrap("cache.processor.GetJSON", err)
	}

	err = json.Unmarshal(raw, v)
	if err != nil {
		return errs.Wrap("unmarshal", err)
	}
	return nil
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	err := c.provider.Delete(ctx, key)
	if err != nil {
		return errs.Wrap("cache.processor.Delete", err)
	}
	return nil
}

func (c *Cache) FlushAll(ctx context.Context) error {
	err := c.provider.FlushAll(ctx)
	if err != nil {
		return errs.Wrap("cache.processor.FlushAll", err)
	}
	return nil
}
