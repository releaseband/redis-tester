package hooks

import (
	"context"

	"github.com/releaseband/redis-tester/hooks/results"

	"github.com/releaseband/redis-tester/hooks/params"
)

func OnConnect(ctx context.Context) {
	name, ok := params.GetName(ctx)
	if ok {
		results.ConnectionsCount.Add(name)
	}
}
