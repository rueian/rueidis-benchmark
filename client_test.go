package rueidis_benchmark

import (
	"context"
	"flag"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"
	"github.com/valkey-io/valkey-glide/go/api"
)

var lib = flag.String("lib", "", "lib to benchmark")

func Benchmark(b *testing.B) {
	bench := func(b *testing.B, address []string, parallelisms []int) {
		var (
			ctx      = context.Background()
			keySizes = []int{16}
			valSizes = []int{64, 256}
			builders = []TargetBuilder{
				{
					Name: "Set",
					Make: func(bench Bench) (Target, error) {
						switch *lib {
						case "goredis":
							client := redis.NewUniversalClient(&redis.UniversalOptions{Addrs: address})
							if err := client.FlushAll(ctx).Err(); err != nil {
								return Target{}, err
							}
							return Target{
								Close: func() { client.Close() },
								Do: func(key string, value string) error {
									return client.Set(ctx, key, value, 0).Err()
								},
							}, nil
						case "glide":
							var base api.BaseClient
							if len(address) == 1 {
								host, port, _ := net.SplitHostPort(address[0])
								portInt, _ := strconv.Atoi(port)
								config := api.NewGlideClientConfiguration().WithAddress(&api.NodeAddress{Host: host, Port: portInt})
								client, err := api.NewGlideClient(config)
								if err != nil {
									return Target{}, err
								}
								if _, err := client.CustomCommand([]string{"FLUSHALL"}); err != nil {
									return Target{}, err
								}
								base = client
							} else {
								config := api.NewGlideClusterClientConfiguration()
								for _, address := range address {
									host, port, _ := net.SplitHostPort(address)
									portInt, _ := strconv.Atoi(port)
									config = config.WithAddress(&api.NodeAddress{Host: host, Port: portInt})
								}
								client, err := api.NewGlideClusterClient(config)
								if err != nil {
									return Target{}, err
								}
								if _, err := client.CustomCommand([]string{"FLUSHALL"}); err != nil {
									return Target{}, err
								}
								base = client
							}
							return Target{
								Close: func() { base.Close() },
								Do: func(key string, value string) error {
									_, err := base.Set(key, value)
									return err
								},
							}, nil
						default:
							client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: address, PipelineMultiplex: -1})
							if err != nil {
								return Target{}, err
							}
							if err := client.Do(ctx, client.B().Flushall().Build()).Error(); err != nil {
								return Target{}, err
							}
							return Target{
								Close: func() { client.Close() },
								Do: func(key string, value string) error {
									return client.Do(ctx, client.B().Set().Key(key).Value(value).Build()).Error()
								},
							}, nil
						}
					},
				},
				{
					Name: "Get",
					Make: func(bench Bench) (Target, error) {
						switch *lib {
						case "goredis":
							client := redis.NewUniversalClient(&redis.UniversalOptions{Addrs: address})
							if err := client.FlushAll(ctx).Err(); err != nil {
								return Target{}, err
							}
							if err := client.Set(ctx, bench.Key, bench.Val, 0).Err(); err != nil {
								return Target{}, err
							}
							return Target{
								Close: func() { client.Close() },
								Do: func(key string, value string) error {
									return client.Get(ctx, key).Err()
								},
							}, nil
						case "glide":
							var base api.BaseClient
							if len(address) == 1 {
								host, port, _ := net.SplitHostPort(address[0])
								portInt, _ := strconv.Atoi(port)
								config := api.NewGlideClientConfiguration().WithAddress(&api.NodeAddress{Host: host, Port: portInt})
								client, err := api.NewGlideClient(config)
								if err != nil {
									return Target{}, err
								}
								if _, err := client.CustomCommand([]string{"FLUSHALL"}); err != nil {
									return Target{}, err
								}
								base = client
							} else {
								config := api.NewGlideClusterClientConfiguration()
								for _, address := range address {
									host, port, _ := net.SplitHostPort(address)
									portInt, _ := strconv.Atoi(port)
									config = config.WithAddress(&api.NodeAddress{Host: host, Port: portInt})
								}
								client, err := api.NewGlideClusterClient(config)
								if err != nil {
									return Target{}, err
								}
								if _, err := client.CustomCommand([]string{"FLUSHALL"}); err != nil {
									return Target{}, err
								}
								base = client
							}
							if _, err := base.Set(bench.Key, bench.Val); err != nil {
								return Target{}, err
							}
							return Target{
								Close: func() { base.Close() },
								Do: func(key string, value string) error {
									_, err := base.Get(key)
									return err
								},
							}, nil
						default:
							client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: address, PipelineMultiplex: -1})
							if err != nil {
								return Target{}, err
							}
							if err := client.Do(ctx, client.B().Flushall().Build()).Error(); err != nil {
								return Target{}, err
							}
							if err := client.Do(ctx, client.B().Set().Key(bench.Key).Value(bench.Val).Build()).Error(); err != nil {
								return Target{}, err
							}
							if *lib == "rueicsc" {
								return Target{
									Close: func() { client.Close() },
									Do: func(key string, value string) error {
										return client.DoCache(ctx, client.B().Get().Key(key).Cache(), 10*time.Second).Error()
									},
								}, nil
							} else {
								return Target{
									Close: func() { client.Close() },
									Do: func(key string, value string) error {
										return client.Do(ctx, client.B().Get().Key(key).Build()).Error()
									},
								}, nil
							}
						}
					},
				},
			}
		)
		RunBenchmark(b, compose(parallelisms, keySizes, valSizes, builders))
	}
	b.Run("OneNode", func(b *testing.B) {
		bench(b, []string{"127.0.0.1:6379"}, []int{1, 8, 64})
	})
	b.Run("Cluster", func(b *testing.B) {
		bench(b, []string{"127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003"}, []int{1, 8, 64})
	})
}
