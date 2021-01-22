package hooks

import (
	"context"

	"github.com/ayupov-ayaz/redis-tester/hooks/results"

	"github.com/ayupov-ayaz/redis-tester/hooks/params"
)

func OnConnect(ctx context.Context) {
	name, ok := params.GetName(ctx)
	if ok {
		results.ConnectionsCount.Add(name)
	}
}
