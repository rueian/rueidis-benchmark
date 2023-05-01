package rueidis_benchmark

import (
	"context"
	"runtime"
	"testing"

	redis "github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"
	"github.com/valyala/fastrand"
)

func BenchmarkClusterClientSet(b *testing.B) {
	var (
		ncpu         = runtime.NumCPU()
		ctx          = context.Background()
		addresses    = []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"}
		parallelisms = []int{4, 16, 96}
		keySizes     = []int{16}
		numKeys      = 1024
		valSizes     = []int{64, 256, 1024}
		builders     = []TargetBuilder{
			{
				Name: "Rueidis",
				Make: func(bench Benchmark) (Target, error) {
					client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: addresses})
					if err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.B().Flushall().Build()).Error(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(keys []string, value string) error {
							return client.Do(ctx, client.B().Set().Key(keys[fastrand.Uint32n(uint32(len(keys)))]).Value(value).Build()).Error()
						},
					}, nil
				},
			},
			{
				Name: "GoRedis",
				Make: func(bench Benchmark) (Target, error) {
					client := redis.NewClusterClient(&redis.ClusterOptions{Addrs: addresses, PoolSize: parallelisms[len(parallelisms)-1] * ncpu})
					if err := client.FlushAll(ctx).Err(); err != nil {
						return Target{}, err
					}
					return Target{
						Close: func() { client.Close() },
						Do: func(keys []string, value string) error {
							return client.Set(ctx, keys[fastrand.Uint32n(uint32(len(keys)))], value, 0).Err()
						},
					}, nil
				},
			},
		}
	)

	RunBenchmark(b, compose(parallelisms, keySizes, valSizes, numKeys, builders))
}
