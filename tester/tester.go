package tester

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/releaseband/redis-tester/hooks/params"

	"github.com/go-redis/redis/v8"

	"github.com/releaseband/redis-tester/options"

	"github.com/releaseband/redis-tester/repository"
)

const (
	setGetKey = "set_get"
	listKey   = "list_key"
)

func sgd(i int) string {
	return setGetKey + strconv.Itoa(i)
}

type RedisTester struct {
	repo   repository.Repository
	opt    options.Test
	logger *log.Logger
}

func NewRedisTester(repo repository.Repository, opt options.Test) RedisTester {
	return RedisTester{
		repo:   repo,
		opt:    opt,
		logger: log.New(os.Stdout, "test", 1),
	}
}

func (t RedisTester) test(testName string, callback func(ctx context.Context, i int) error) {
	ctx := params.CtxWithName(context.Background(), testName)

	wg := sync.WaitGroup{}
	goroutines := t.opt.Goroutines()
	count := t.opt.Count()

	t.logger.Println("test name", testName)

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(gNumber int) {
			defer wg.Done()

			for j := 0; j < count; j++ {
				if err := callback(ctx, j*gNumber); err != nil {
					t.logger.Println(fmt.Errorf("iteration %d failed: %w", j, err))
					return
				}
			}
		}(i + 1)
	}

	wg.Wait()
}

func (t RedisTester) testSet() {
	expiration := t.opt.Expiration()
	t.logger.Printf("%s parms: expiration: %s", params.Set, expiration.String())

	t.test(params.Set, func(ctx context.Context, i int) error {
		v := sgd(i)
		if err := t.repo.Set(ctx, v, v, expiration); err != nil {
			return fmt.Errorf("repo.Set failed: %w", err)
		}

		return nil
	})
}

func (t RedisTester) testGet() {
	t.test(params.Get, func(ctx context.Context, i int) error {
		v := sgd(i)
		if _, err := t.repo.Get(ctx, v); err != nil && err != redis.Nil {
			return fmt.Errorf("repo.Get failed: %w", err)
		}

		return nil
	})
}

func (t RedisTester) testDet() {
	t.test(params.Del, func(ctx context.Context, i int) error {
		v := sgd(i)
		if err := t.repo.Del(ctx, v); err != nil {
			return fmt.Errorf("repo.Del failed: %w", err)
		}

		return nil
	})
}

func (t RedisTester) testRPush() {
	t.test(params.RPush, func(ctx context.Context, i int) error {
		v := listKey + strconv.Itoa(i)
		if err := t.repo.RPush(ctx, listKey, v); err != nil {
			return fmt.Errorf("repo.RPush failed: %w", err)
		}

		return nil
	})
}

func (t RedisTester) testLTrim() {
	t.test(params.LTrim, func(ctx context.Context, i int) error {
		if err := t.repo.LTrim(ctx, listKey, 0, int64(i)); err != nil {
			return fmt.Errorf("repo.LTrim failed: %w", err)
		}

		return nil
	})
}

func (t RedisTester) testLRange() {
	t.test(params.LRange, func(ctx context.Context, i int) error {
		if _, err := t.repo.LRange(ctx, listKey, 0, int64(i)); err != nil {
			return fmt.Errorf("repo.LRange failed: %w", err)
		}

		return nil
	})
}

func (t RedisTester) Run() {
	c := t.opt.Commands

	t.logger.Printf("goroutines: %d | iterations: %d", t.opt.Goroutines(), t.opt.Iterations)

	if c.Set {
		t.testSet()
	}

	if c.Get {
		t.testGet()
	}

	if c.Del {
		t.testDet()
	}

	if c.RPush {
		t.testRPush()
	}

	if c.LRange {
		t.testLRange()
	}

	if c.LTrim {
		t.testLTrim()
	}

}
