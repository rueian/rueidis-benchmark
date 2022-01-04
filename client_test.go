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
		numKeys      = 1
		valSizes     = []int{64, 256, 1024}
		builders     = []TargetBuilder{
			{
				Name: "Rueidis",
				Make: func(bench Benchmark) (Target, error) {
					client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{address}})
					if err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.B().Flushall().Build()).Error(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(keys []string, value string) error {
							return client.Do(ctx, client.B().Set().Key(keys[0]).Value(value).Build()).Error()
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
						Do: func(keys []string, value string) error {
							return client.Set(ctx, keys[0], value, 0).Err()
						},
					}, nil
				},
			},
		}
	)

	RunBenchmark(b, compose(parallelisms, keySizes, valSizes, numKeys, builders))
}

func BenchmarkSingleClientGet(b *testing.B) {
	var (
		ncpu         = runtime.NumCPU()
		ctx          = context.Background()
		address      = "127.0.0.1:6379"
		parallelisms = []int{1, 8, 64}
		keySizes     = []int{16}
		numKeys      = 1
		valSizes     = []int{64, 256, 1024}
		builders     = []TargetBuilder{
			{
				Name: "RueidisCSC",
				Make: func(bench Benchmark) (Target, error) {
					client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{address}})
					if err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.B().Flushall().Build()).Error(); err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.B().Set().Key(bench.Keys[0]).Value(bench.Val).Build()).Error(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(keys []string, value string) error {
							return client.DoCache(ctx, client.B().Get().Key(keys[0]).Cache(), 10*time.Second).Error()
						},
					}, nil
				},
			},
			{
				Name: "Rueidis",
				Make: func(bench Benchmark) (Target, error) {
					client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{address}})
					if err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.B().Flushall().Build()).Error(); err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.B().Set().Key(bench.Keys[0]).Value(bench.Val).Build()).Error(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(keys []string, value string) error {
							return client.Do(ctx, client.B().Get().Key(keys[0]).Build()).Error()
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
					if err := client.Set(ctx, bench.Keys[0], bench.Val, 0).Err(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(keys []string, value string) error {
							return client.Get(ctx, keys[0]).Err()
						},
					}, nil
				},
			},
		}
	)

	RunBenchmark(b, compose(parallelisms, keySizes, valSizes, numKeys, builders))
}
