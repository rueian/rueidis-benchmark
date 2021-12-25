package rueidis_benchmark

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/rueian/rueidis"
)

type Benchmark struct {
	Parallelism   int
	Key           string
	Val           string
	TargetBuilder TargetBuilder
}

type Target struct {
	Close func()
	Do    func(key, value string) error
}

type TargetBuilder struct {
	Name string
	Make func() (Target, error)
}

func BenchmarkSingleClient(b *testing.B) {
	var (
		ctx          = context.Background()
		address      = "127.0.0.1:6379"
		parallelisms = []int{1, 8, 64}
		KeySizes     = []int{16}
		valSizes     = []int{64, 256, 1024}
		builders     = []TargetBuilder{
			{
				Name: "RueidisSet",
				Make: func() (Target, error) {
					client, err := rueidis.NewSingleClient(rueidis.SingleClientOption{Address: address})
					if err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.Cmd.Flushall().Build()).Error(); err != nil {
						return Target{}, err
					}
					if err := client.Do(ctx, client.Cmd.ConfigSet().ParameterValue().ParameterValue("save", "").Build()).Error(); err != nil {
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
				Name: "GoRedisSet",
				Make: func() (Target, error) {
					client := redis.NewClient(&redis.Options{Addr: address, PoolSize: parallelisms[len(parallelisms)-1]})
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

	benchmarks := make([]Benchmark, 0, len(parallelisms)*len(KeySizes)*len(valSizes)*len(builders))

	for _, p := range parallelisms {
		for _, k := range KeySizes {
			key := gen(k)
			for _, v := range valSizes {
				val := gen(v)
				for _, builder := range builders {
					benchmarks = append(benchmarks, Benchmark{
						Parallelism:   p,
						Key:           key,
						Val:           val,
						TargetBuilder: builder,
					})
				}
			}
		}
	}

	for _, bench := range benchmarks {
		bench := bench
		b.Run(fmt.Sprintf("%s-parallel(%d)-key(%d)-value(%d)", bench.TargetBuilder.Name, bench.Parallelism, len(bench.Key), len(bench.Val)), func(b *testing.B) {
			target, err := bench.TargetBuilder.Make()
			if err != nil {
				b.Fatalf("%s setup fail: %v", bench.TargetBuilder.Name, err)
			}
			b.SetParallelism(bench.Parallelism)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if err := target.Do(bench.Key, bench.Val); err != nil {
						b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
					}
				}
			})
			b.StopTimer()
			target.Close()
		})
	}
}

func gen(n int) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(rand.Intn(26) + 'a'))
	}
	return sb.String()
}
