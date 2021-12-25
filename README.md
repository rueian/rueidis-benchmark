# Rueidis Benchmark

## Benchmark comparison with go-redis v8.11.4

Rueidis has higher throughput than go-redis v8.11.4 across 1, 8, and 64 parallelism settings.

In some case, it is even able to achieve ~9x throughput over go-redis on local benchmark. (BenchmarkSingleClient/RueidisSet-parallel(64)-key(16)-value(64)-10)

```shell
# run redis-server 6.2.6 at 127.0.0.1:6379
▶ ./redis-server
▶ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: rueidis-benchmark
BenchmarkSingleClient/RueidisSet-parallel(1)-key(16)-value(64)-10         	  401235	      2990 ns/op	     108 B/op	       3 allocs/op
BenchmarkSingleClient/GoRedisSet-parallel(1)-key(16)-value(64)-10         	  245170	      5158 ns/op	     264 B/op	       7 allocs/op
BenchmarkSingleClient/RueidisSet-parallel(1)-key(16)-value(256)-10        	  364747	      3021 ns/op	     111 B/op	       4 allocs/op
BenchmarkSingleClient/GoRedisSet-parallel(1)-key(16)-value(256)-10        	  236398	      4923 ns/op	     264 B/op	       7 allocs/op
BenchmarkSingleClient/RueidisSet-parallel(1)-key(16)-value(1024)-10       	  369381	      3071 ns/op	     112 B/op	       4 allocs/op
BenchmarkSingleClient/GoRedisSet-parallel(1)-key(16)-value(1024)-10       	  221430	      5117 ns/op	     264 B/op	       7 allocs/op
BenchmarkSingleClient/RueidisSet-parallel(8)-key(16)-value(64)-10         	 1404769	       837.9 ns/op	     108 B/op	       3 allocs/op
BenchmarkSingleClient/GoRedisSet-parallel(8)-key(16)-value(64)-10         	  228680	      5124 ns/op	     267 B/op	       7 allocs/op
BenchmarkSingleClient/RueidisSet-parallel(8)-key(16)-value(256)-10        	 1392675	       843.4 ns/op	     111 B/op	       4 allocs/op
BenchmarkSingleClient/GoRedisSet-parallel(8)-key(16)-value(256)-10        	  266722	      5059 ns/op	     266 B/op	       7 allocs/op
BenchmarkSingleClient/RueidisSet-parallel(8)-key(16)-value(1024)-10       	 1059328	      1084 ns/op	     112 B/op	       4 allocs/op
BenchmarkSingleClient/GoRedisSet-parallel(8)-key(16)-value(1024)-10       	  243488	      5769 ns/op	     266 B/op	       7 allocs/op
BenchmarkSingleClient/RueidisSet-parallel(64)-key(16)-value(64)-10        	 1877408	       592.0 ns/op	     108 B/op	       3 allocs/op
BenchmarkSingleClient/GoRedisSet-parallel(64)-key(16)-value(64)-10        	  207786	      4983 ns/op	     268 B/op	       7 allocs/op
BenchmarkSingleClient/RueidisSet-parallel(64)-key(16)-value(256)-10       	 1611327	       752.9 ns/op	     111 B/op	       4 allocs/op
BenchmarkSingleClient/GoRedisSet-parallel(64)-key(16)-value(256)-10       	  248240	      5168 ns/op	     267 B/op	       7 allocs/op
BenchmarkSingleClient/RueidisSet-parallel(64)-key(16)-value(1024)-10      	  873829	      1258 ns/op	     112 B/op	       4 allocs/op
BenchmarkSingleClient/GoRedisSet-parallel(64)-key(16)-value(1024)-10      	  224452	      5483 ns/op	     267 B/op	       7 allocs/op
PASS
ok  	rueidis-benchmark	31.005s
```