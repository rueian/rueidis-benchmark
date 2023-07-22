package rueidis_benchmark

import (
	"fmt"
	"strings"
	"testing"

	"github.com/valyala/fastrand"
)

type Bench struct {
	Parallelism   int
	Key           string
	Val           string
	TargetBuilder TargetBuilder
}

type Target struct {
	Close func()
	Do    func(keys string, value string) error
}

type TargetBuilder struct {
	Name string
	Make func(bench Bench) (Target, error)
}

func gen(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(byte(fastrand.Uint32n(26) + 'a'))
	}
	return sb.String()
}

func compose(parallelisms, keySizes, valSizes []int, builders []TargetBuilder) []Bench {
	benchmarks := make([]Bench, 0, len(parallelisms)*len(keySizes)*len(valSizes)*len(builders))

	for _, p := range parallelisms {
		for _, k := range keySizes {
			for _, v := range valSizes {
				for _, builder := range builders {
					benchmarks = append(benchmarks, Bench{
						Parallelism:   p,
						Key:           gen(k),
						Val:           gen(v),
						TargetBuilder: builder,
					})
				}
			}
		}
	}

	return benchmarks
}

func RunBenchmark(b *testing.B, benchmarks []Bench) {
	for _, bench := range benchmarks {
		bench := bench
		b.Run(fmt.Sprintf("%s-parall(%d)-key(%d)-val(%d)", bench.TargetBuilder.Name, bench.Parallelism, len(bench.Key), len(bench.Val)), func(b *testing.B) {
			target, err := bench.TargetBuilder.Make(bench)
			if err != nil {
				b.Fatalf("%s setup fail: %v", bench.TargetBuilder.Name, err)
			}
			if bench.Parallelism == 0 {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					if err := target.Do(bench.Key, bench.Val); err != nil {
						b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
					}
				}
			} else {
				b.SetParallelism(bench.Parallelism)
				b.ResetTimer()
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						if err := target.Do(bench.Key, bench.Val); err != nil {
							b.Errorf("%s error during benchmark: %v", bench.TargetBuilder.Name, err)
						}
					}
				})
			}
			b.StopTimer()
			target.Close()
		})
	}
}
