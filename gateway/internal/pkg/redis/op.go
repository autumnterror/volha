package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/internal/utils/format"
	"gateway/internal/views"
	"github.com/redis/go-redis/v9"
	"time"
)

const cacheKey = "dictionaries:all"

var CashedCategoriesId []string

func (c *Client) GetDictionaries() (*views.Dictionaries, error) {
	const (
		op = "redis.GetAllDictionaries"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ds, err := c.Rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var cached views.Dictionaries
		if err := json.Unmarshal([]byte(ds), &cached); err != nil {
			return nil, format.Error(op, err)
		}
		return &cached, nil
	}
	if err == redis.Nil {
		return nil, format.Error(op, fmt.Errorf("no cache"))
	}

	return nil, format.Error(op, err)
}
func (c *Client) SetDictionaries(d *views.Dictionaries) error {
	const (
		op = "redis.SetDictionaries"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	bytes, err := json.Marshal(d)
	if err != nil {
		return format.Error(op, fmt.Errorf("SetDictionariesToCache: failed to marshal: %w", err))
	}

	if err := c.Rdb.Set(ctx, cacheKey, bytes, 1*time.Hour).Err(); err != nil {
		return format.Error(op, err)
	}

	return nil
}

func (c *Client) CleanDictionaries() error {
	const (
		op = "redis.CleanDictionaries"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := c.Rdb.Del(ctx, cacheKey).Err(); err != nil {
		return format.Error(op, err)
	}
	for _, id := range CashedCategoriesId {
		if err := c.Rdb.Del(ctx, cacheKey+id).Err(); err != nil {
			return format.Error(op, err)
		}
	}
	CashedCategoriesId = []string{}
	return nil
}

func (c *Client) GetDictionariesByCategory(id string) (*views.Dictionaries, error) {
	const (
		op = "redis.GetDictionariesByCategory"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ds, err := c.Rdb.Get(ctx, cacheKey+id).Result()
	if err == nil {
		var cached views.Dictionaries
		if err := json.Unmarshal([]byte(ds), &cached); err != nil {
			return nil, format.Error(op, err)
		}
		return &cached, nil
	}
	if err == redis.Nil {
		return nil, format.Error(op, fmt.Errorf("no cache"))
	}

	return nil, format.Error(op, err)
}
func (c *Client) SetDictionariesByCategory(id string, d *views.Dictionaries) error {
	const (
		op = "redis.SetDictionariesByCategory"
	)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	bytes, err := json.Marshal(d)
	if err != nil {
		return format.Error(op, fmt.Errorf("SetDictionariesToCache: failed to marshal: %w", err))
	}

	if err := c.Rdb.Set(ctx, cacheKey+id, bytes, 1*time.Hour).Err(); err != nil {
		return format.Error(op, err)
	}

	CashedCategoriesId = append(CashedCategoriesId, id)

	return nil
}
