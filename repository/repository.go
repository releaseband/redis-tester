package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/releaseband/redis-tester/hooks/results"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	redis *redis.ClusterClient
}

func NewRepository(client *redis.ClusterClient) Repository {
	return Repository{
		redis: client,
	}
}

func addWrite(client *redis.Client) {
	results.AddWrite(client.String())
}

func addRead(client *redis.Client) {
	results.AddRead(client.String())
}

func (r Repository) Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	master, err := r.redis.MasterForKey(ctx, key)
	if err != nil {
		return fmt.Errorf("get MasterForKey failed: %w", err)
	}

	if err := master.Set(ctx, key, v, expiration).Err(); err != nil {
		return fmt.Errorf("set for %s failed: %w", key, err)
	}

	addWrite(master)

	return nil
}

func (r Repository) Get(ctx context.Context, key string) (string, error) {
	slave, err := r.redis.SlaveForKey(ctx, key)
	if err != nil {
		return "", fmt.Errorf("get SlaveForKey failed: %w", err)
	}

	addRead(slave)

	return slave.Get(ctx, key).Result()
}
func (r Repository) Del(ctx context.Context, key string) error {
	return r.redis.Del(ctx, key).Err()
}

func (r Repository) RPush(ctx context.Context, key string, val ...interface{}) error {
	return r.redis.RPush(ctx, key, val).Err()
}

func (r Repository) LTrim(ctx context.Context, key string, start, stop int64) error {
	return r.redis.LTrim(ctx, key, start, stop).Err()
}

func (r Repository) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.redis.LRange(ctx, key, start, stop).Result()
}
