package store

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisRepository(client *redis.Client) *redisRepository {
	return &redisRepository{client: client}
}

func (r *redisRepository) Set(ctx context.Context, payload Item) (Item, error) {
	payloadBytes, _ := json.Marshal(payload)

	key := "items." + payload.ID
	err := r.client.Set(ctx, key, string(payloadBytes), 24*time.Hour).Err()

	return payload, err
}

func (r *redisRepository) Get(ctx context.Context, id string) (Item, error) {
	key := "items." + id
	result, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return Item{}, err
	}

	var item Item
	err = json.Unmarshal(result, &item)

	return item, err
}

func (r *redisRepository) Clear(ctx context.Context, id string) error {
	key := "items." + id

	return r.client.Del(ctx, key).Err()
}
