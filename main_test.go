package rueidis_benchmark

import (
	"fmt"
	"strings"
	"testing"

	"github.com/valyala/fastrand"
)

type Benchmark struct {
	Parallelism   int
	Keys          []string
	Val           string
	TargetBuilder TargetBuilder
}

type Target struct {
	Close func()
	Do    func(keys []string, value string) error
}

type TargetBuilder struct {
	Name string
	Make func(bench Benchmark) (Target, error)
}

func gen(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(fastrand.Uint32n(26) + 'a'))
	}
	return sb.String()
}

func gens(s, n int) (gs []string) {
	gs = make([]string, n)
	for i := range gs {
		gs[i] = gen(s)
	}
	return gs
}

func compose(parallelisms, keySizes, valSizes []int, numKeys int, builders []TargetBuilder) []Benchmark {
	benchmarks := make([]Benchmark, 0, len(parallelisms)*len(keySizes)*len(valSizes)*len(builders))

	for _, p := range parallelisms {
		for _, k := range keySizes {
			keys := gens(k, numKeys)
			for _, v := range valSizes {
				val := gen(v)
				for _, builder := range builders {
					benchmarks = append(benchmarks, Benchmark{
						Parallelism:   p,
						Keys:          keys,
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
		b.Run(fmt.Sprintf("%s-parallelism(%d)-key(%d)-value(%d)", bench.TargetBuilder.Name, bench.Parallelism, len(bench.Keys[0]), len(bench.Val)), func(b *testing.B) {
			target, err := bench.TargetBuilder.Make(bench)
			if err != nil {
				b.Fatalf("%s setup fail: %v", bench.TargetBuilder.Name, err)
			}
			b.SetParallelism(bench.Parallelism)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if err := target.Do(bench.Keys, bench.Val); err != nil {
						b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
					}
				}
			})
			b.StopTimer()
			target.Close()
		})
	}
}
