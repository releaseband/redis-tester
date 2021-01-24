package options

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func makeNodes(addresses []string) []redis.ClusterNode {
	nodes := make([]redis.ClusterNode, len(addresses))
	for i, addr := range addresses {
		nodes[i] = redis.ClusterNode{
			Addr: addr,
		}
	}

	return nodes
}

func clusterSlots(ts ManualCluster) func(ctx context.Context) ([]redis.ClusterSlot, error) {
	if !ts.Use || len(ts.Slots) == 0 {
		return nil
	}

	return func(ctx context.Context) ([]redis.ClusterSlot, error) {
		slots := make([]redis.ClusterSlot, len(ts.Slots))
		for i, cfg := range ts.Slots {
			slots[i] = redis.ClusterSlot{
				Start: cfg.Start,
				End:   cfg.End,
				Nodes: makeNodes(ts.Slots[i].Addresses),
			}
		}

		return slots, nil
	}
}
