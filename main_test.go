package rueidis_benchmark

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
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
	Make func(bench Benchmark) (Target, error)
}

func gen(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(rand.Intn(26) + 'a'))
	}
	return sb.String()
}

func compose(parallelisms, keySizes, valSizes []int, builders []TargetBuilder) []Benchmark {
	benchmarks := make([]Benchmark, 0, len(parallelisms)*len(keySizes)*len(valSizes)*len(builders))

	for _, p := range parallelisms {
		for _, k := range keySizes {
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

	return benchmarks
}

func RunBenchmark(b *testing.B, benchmarks []Benchmark) {
	for _, bench := range benchmarks {
		bench := bench
		b.Run(fmt.Sprintf("%s-parallelism(%d)-key(%d)-value(%d)", bench.TargetBuilder.Name, bench.Parallelism, len(bench.Key), len(bench.Val)), func(b *testing.B) {
			target, err := bench.TargetBuilder.Make(bench)
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
