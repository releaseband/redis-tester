package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/releaseband/redis-tester/hooks/results"

	"github.com/releaseband/redis-tester/hooks/params"

	"github.com/go-redis/redis/v8"
	"github.com/releaseband/redis-tester/options"
	"github.com/releaseband/redis-tester/repository"
	"github.com/releaseband/redis-tester/tester"
)

const (
	configsFilePath = "./.configs.yaml"
)

func makeRedisClient(ts *options.TesterSettings) (redis.Cmdable, error) {
	var client redis.Cmdable
	if ts.Cluster.Use {
		cluster := redis.NewClusterClient(options.ClusterOptions(ts.ClusterSettings()))
		if ts.Cluster.Manual.Use {
			cluster.ReloadState(context.Background())
		}

		client = cluster
	} else {
		client = redis.NewClient(options.RedisOptions(ts.RedisSettings()))
	}

	ctx, cancel := context.WithTimeout(params.CtxWithName(context.Background(), params.Ping), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	return client, nil
}

func printToConsole(key, val string) {
	fmt.Println(key, val)
}

func printResults() error {
	connections, err := results.GetConnectionsInfo()
	if err != nil {
		return fmt.Errorf("GetConnectionInfo failed: %w", err)
	}

	printToConsole("connections:", connections)

	return nil
}

func Run() error {
	ts, err := makeTesterSettings(configsFilePath)
	if err != nil {
		return err
	}

	client, err := makeRedisClient(ts)
	if err != nil {
		return fmt.Errorf("makeRedisClient failed: %w", err)
	}

	tester.NewRedisTester(repository.NewRepository(client), ts.Test).Run()

	return printResults()
}
