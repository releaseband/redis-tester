package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/releaseband/redis-tester/hooks/results"

	"github.com/go-redis/redis/v8"
	"github.com/releaseband/redis-tester/options"
	"github.com/releaseband/redis-tester/repository"
	"github.com/releaseband/redis-tester/tester"
)

const (
	configsFilePath = "./.configs.yaml"
)

func Run() error {
	ts, err := makeTesterSettings(configsFilePath)
	if err != nil {
		return err
	}

	var client redis.Cmdable
	if ts.UseCluster {
		client = redis.NewClusterClient(options.ClusterOptions(ts.ClusterSettings()))
	} else {
		client = redis.NewClient(options.RedisOptions(ts.RedisSettings()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis ping failed: %w", err)
	}

	tester.NewRedisTester(repository.NewRepository(client), ts.Test).Run()
	connections, err := results.GetConnectionsInfo()
	if err != nil {
		return fmt.Errorf("GetConnectionInfo failed: %w", err)
	}

	fmt.Println(connections)

	return nil
}
