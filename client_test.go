package rueidis_benchmark

import (
	"context"
	"runtime"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rueian/rueidis"
)

func BenchmarkSingleClientSet(b *testing.B) {
	var (
		ncpu         = runtime.NumCPU()
		ctx          = context.Background()
		address      = "127.0.0.1:6379"
		parallelisms = []int{1, 8, 64}
		keySizes     = []int{16}
		valSizes     = []int{64, 256, 1024}
		builders     = []TargetBuilder{
			{
				Name: "Rueidis",
				Make: func(bench Benchmark) (Target, error) {
					client, err := rueidis.NewSingleClient(rueidis.SingleClientOption{Address: address})
					if err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.Cmd.Flushall().Build()).Error(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(key, value string) error {
							return client.Do(ctx, client.Cmd.Set().Key(key).Value(value).Build()).Error()
						},
					}, nil
				},
			},
			{
				Name: "GoRedis",
				Make: func(bench Benchmark) (Target, error) {
					client := redis.NewClient(&redis.Options{Addr: address, PoolSize: parallelisms[len(parallelisms)-1] * ncpu})
					if err := client.FlushAll(ctx).Err(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(key, value string) error {
							return client.Set(ctx, key, value, 0).Err()
						},
					}, nil
				},
			},
		}
	)

	RunBenchmark(b, compose(parallelisms, keySizes, valSizes, builders))
}

func BenchmarkSingleClientGet(b *testing.B) {
	var (
		ncpu         = runtime.NumCPU()
		ctx          = context.Background()
		address      = "127.0.0.1:6379"
		parallelisms = []int{1, 8, 64}
		keySizes     = []int{16}
		valSizes     = []int{64, 256, 1024}
		builders     = []TargetBuilder{
			{
				Name: "RueidisCSC",
				Make: func(bench Benchmark) (Target, error) {
					client, err := rueidis.NewSingleClient(rueidis.SingleClientOption{Address: address})
					if err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.Cmd.Flushall().Build()).Error(); err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.Cmd.Set().Key(bench.Key).Value(bench.Val).Build()).Error(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(key, value string) error {
							return client.DoCache(ctx, client.Cmd.Get().Key(key).Cache(), 10*time.Second).Error()
						},
					}, nil
				},
			},
			{
				Name: "Rueidis",
				Make: func(bench Benchmark) (Target, error) {
					client, err := rueidis.NewSingleClient(rueidis.SingleClientOption{Address: address})
					if err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.Cmd.Flushall().Build()).Error(); err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.Cmd.Set().Key(bench.Key).Value(bench.Val).Build()).Error(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(key, value string) error {
							return client.Do(ctx, client.Cmd.Get().Key(key).Build()).Error()
						},
					}, nil
				},
			},
			{
				Name: "GoRedis",
				Make: func(bench Benchmark) (Target, error) {
					client := redis.NewClient(&redis.Options{Addr: address, PoolSize: parallelisms[len(parallelisms)-1] * ncpu})
					if err := client.FlushAll(ctx).Err(); err != nil {
						return Target{}, err
					}
					if err := client.Set(ctx, bench.Key, bench.Val, 0).Err(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(key, value string) error {
							return client.Get(ctx, key).Err()
						},
					}, nil
				},
			},
		}
	)

	RunBenchmark(b, compose(parallelisms, keySizes, valSizes, builders))
}
