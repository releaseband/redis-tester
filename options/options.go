package options

import (
	"context"
	"crypto/tls"

	"github.com/go-redis/redis/v8"
	"github.com/releaseband/redis-tester/hooks"
)

func ClusterOptions(opt ClusterSettings) *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:              opt.Addresses,
		MaxRedirects:       opt.MaxRedirects,
		ReadOnly:           opt.ReadOnly,
		RouteByLatency:     opt.RouteByLatency,
		RouteRandomly:      opt.RouteRandomly,
		Username:           opt.User,
		Password:           opt.Password,
		MaxRetries:         opt.MaxRetries,
		MinRetryBackoff:    opt.MinRetryBackoff,
		MaxRetryBackoff:    opt.MaxRetryBackoff,
		DialTimeout:        opt.Dial,
		ReadTimeout:        opt.Read,
		WriteTimeout:       opt.Write,
		PoolSize:           opt.PoolSize,
		MinIdleConns:       opt.MinIdleConns,
		MaxConnAge:         opt.MaxConnAge,
		PoolTimeout:        opt.Pool,
		IdleTimeout:        opt.Idle,
		IdleCheckFrequency: opt.IdleCheckFrequency,
		TLSConfig:          &tls.Config{InsecureSkipVerify: true},
		NewClient:          nil,
		ClusterSlots:       nil,
		Dialer:             nil,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			hooks.OnConnect(ctx)
			return nil
		},
	}
}

func RedisOptions(opt RedisSettings) *redis.Options {
	return &redis.Options{
		Network:            opt.Network,
		Addr:               opt.Address,
		Username:           opt.User,
		Password:           opt.Password,
		DB:                 0,
		MaxRetries:         opt.MaxRetries,
		MinRetryBackoff:    opt.MinRetryBackoff,
		MaxRetryBackoff:    opt.MaxRetryBackoff,
		DialTimeout:        opt.Dial,
		ReadTimeout:        opt.Read,
		WriteTimeout:       opt.Write,
		PoolSize:           opt.PoolSize,
		MinIdleConns:       opt.MinIdleConns,
		MaxConnAge:         opt.MaxConnAge,
		PoolTimeout:        opt.Pool,
		IdleTimeout:        opt.Idle,
		IdleCheckFrequency: opt.IdleCheckFrequency,
		// TLSConfig:          &tls.Config{InsecureSkipVerify: true},
		Limiter: nil,
		Dialer:  nil,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			hooks.OnConnect(ctx)
			return nil
		},
	}
}
