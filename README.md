# Rueidis Benchmark
This is a benchmark repository for https://github.com/redis/rueidis

## Benchmark comparison with go-redis v9

Rueidis has higher throughput than go-redis v9 across 1, 8, and 64 parallelism settings.

## Run the benchmark

```shell
# prepare redis servers
▶ ./redis-server --save "" --appendonly no
▶ ./redis-server --port 7001 --save "" --appendonly no --cluster-enabled yes --cluster-config-file 7001.conf
▶ ./redis-server --port 7002 --save "" --appendonly no --cluster-enabled yes --cluster-config-file 7002.conf
▶ ./redis-server --port 7003 --save "" --appendonly no --cluster-enabled yes --cluster-config-file 7003.conf
▶ ./redis-cli --cluster create 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 --cluster-yes
# run the benchmark
▶ ./bench.sh
# compare the results
▶ benchstat get*
▶ benchstat set*
```

## Redis SET

Comparing to the go-redis baseline, rueidis can achieve about **88%**, **99%** and **100%** reduction on `sec/op`, `B/op` and `allocs/op` respectively.

```shell
▶ benchstat set*
goos: darwin
goarch: arm64
pkg: rueidis-benchmark
                                             │ set-goredis.txt │           set-rueidis.txt           │
                                             │     sec/op      │   sec/op     vs base                │
/OneNode/Set-parall(1)-key(16)-val(64)-10         5.536µ ± 14%   3.344µ ± 2%  -39.60% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10        5.455µ ±  4%   3.377µ ± 1%  -38.09% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(1024)-10       5.702µ ±  9%   3.431µ ± 0%  -39.83% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10         6.284µ ±  8%   1.067µ ± 1%  -83.02% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10        6.186µ ± 12%   1.123µ ± 4%  -81.84% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(1024)-10       5.841µ ± 10%   1.220µ ± 1%  -79.11% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10       6368.5n ± 10%   727.2n ± 1%  -88.58% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10      6623.0n ± 13%   890.1n ± 2%  -86.56% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(1024)-10     6619.5n ± 17%   953.4n ± 1%  -85.60% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10         5.486µ ±  6%   3.399µ ± 1%  -38.05% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10        5.676µ ±  8%   3.415µ ± 0%  -39.83% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(1024)-10       5.686µ ±  4%   3.465µ ± 1%  -39.06% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10         6.387µ ±  7%   1.133µ ± 2%  -82.26% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10        6.048µ ±  8%   1.183µ ± 1%  -80.44% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(1024)-10       6.220µ ± 12%   1.282µ ± 1%  -79.39% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10       6326.5n ±  9%   782.9n ± 6%  -87.63% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10      7064.5n ± 14%   984.2n ± 3%  -86.07% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(1024)-10      6.741µ ±  8%   1.026µ ± 1%  -84.79% (p=0.000 n=10)
geomean                                           6.108µ         1.521µ       -75.09%

                                             │ set-goredis.txt │              set-rueidis.txt              │
                                             │      B/op       │     B/op      vs base                     │
/OneNode/Set-parall(1)-key(16)-val(64)-10           257.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10          257.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(1024)-10         257.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10           264.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10          264.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(1024)-10         264.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10        269.000 ± 0%   1.000 ±   0%   -99.63% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10       269.000 ± 0%   1.000 ±   0%   -99.63% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(1024)-10      269.000 ± 0%   1.000 ±   0%   -99.63% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10           256.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10          256.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(1024)-10         256.5 ± 0%     0.0 ±    ?  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10           261.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10          262.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(1024)-10         262.0 ± 0%     0.0 ±   0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10        264.000 ± 0%   1.000 ±   0%   -99.62% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10       264.000 ± 0%   1.000 ± 100%   -99.62% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(1024)-10      264.000 ± 0%   1.000 ±   0%   -99.62% (p=0.000 n=10)
geomean                                             261.9                      ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

                                             │ set-goredis.txt │             set-rueidis.txt             │
                                             │    allocs/op    │ allocs/op   vs base                     │
/OneNode/Set-parall(1)-key(16)-val(64)-10           9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10          9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(1024)-10         9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10           9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10          9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(1024)-10         9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10          9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10         9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(1024)-10        9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10           9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10          9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(1024)-10         9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10           9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10          9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(1024)-10         9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10          9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10         9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(1024)-10        9.000 ± 0%   0.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                                             9.000                    ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

```

## Redis GET

Comparing to the go-redis baseline, rueidis can achieve about **92%**, **76%** and **87%** reduction on `sec/op`, `B/op` and `allocs/op` respectively.

With the client-side caching, rueidis can achieve about **97%**, **99%** and **100%** reduction on `sec/op`, `B/op` and `allocs/op` respectively in cache hit scenario.

```shell
goos: darwin
goarch: arm64
pkg: rueidis-benchmark
                                             │ get-goredis.txt │            get-rueidis.txt            │         get-rueidiscsc.txt          │
                                             │     sec/op      │    sec/op      vs base                │   sec/op     vs base                │
/OneNode/Get-parall(1)-key(16)-val(64)-10        5491.5n ± 14%   3075.5n ±  2%  -44.00% (p=0.000 n=10)   151.2n ± 4%  -97.25% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10       5766.0n ±  7%   3124.5n ±  1%  -45.81% (p=0.000 n=10)   149.0n ± 4%  -97.42% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(1024)-10      5495.5n ± 11%   3320.0n ±  1%  -39.59% (p=0.000 n=10)   151.5n ± 2%  -97.24% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10        5377.5n ± 12%    719.5n ±  2%  -86.62% (p=0.000 n=10)   150.1n ± 2%  -97.21% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10       5503.5n ± 19%    761.8n ± 15%  -86.16% (p=0.000 n=10)   150.7n ± 3%  -97.26% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(1024)-10      6178.0n ±  8%   1003.5n ±  5%  -83.76% (p=0.000 n=10)   151.2n ± 1%  -97.55% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10       5878.5n ±  9%    470.4n ±  7%  -92.00% (p=0.000 n=10)   155.2n ± 1%  -97.36% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10      6498.5n ± 13%    588.0n ± 12%  -90.95% (p=0.000 n=10)   152.2n ± 3%  -97.66% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(1024)-10     6714.0n ±  7%    834.0n ± 10%  -87.58% (p=0.000 n=10)   152.5n ± 1%  -97.73% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10        5089.0n ± 27%   3174.5n ±  3%  -37.62% (p=0.000 n=10)   165.3n ± 1%  -96.75% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10       5534.0n ± 17%   3221.0n ±  0%  -41.80% (p=0.000 n=10)   167.5n ± 2%  -96.97% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(1024)-10      5643.5n ± 19%   3356.0n ±  1%  -40.53% (p=0.000 n=10)   166.9n ± 2%  -97.04% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10        5928.0n ± 13%    762.1n ±  1%  -87.14% (p=0.000 n=10)   160.3n ± 4%  -97.30% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10       6144.5n ± 21%    796.0n ±  2%  -87.04% (p=0.000 n=10)   161.9n ± 2%  -97.36% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(1024)-10      5931.0n ±  7%    958.5n ± 10%  -83.84% (p=0.000 n=10)   162.5n ± 1%  -97.26% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10       5721.0n ± 11%    622.0n ± 34%  -89.13% (p=0.000 n=10)   168.7n ± 0%  -97.05% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10      6164.5n ± 16%    537.8n ±  5%  -91.28% (p=0.000 n=10)   167.0n ± 2%  -97.29% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(1024)-10     7035.5n ±  7%    787.0n ±  5%  -88.81% (p=0.000 n=10)   167.2n ± 2%  -97.62% (p=0.000 n=10)
geomean                                           5.875µ          1.185µ        -79.83%                  158.2n       -97.31%

                                             │ get-goredis.txt │            get-rueidis.txt            │            get-rueidiscsc.txt             │
                                             │      B/op       │     B/op       vs base                │     B/op      vs base                     │
/OneNode/Get-parall(1)-key(16)-val(64)-10          269.00 ± 0%      64.00 ± 0%  -76.21% (p=0.000 n=10)      0.00 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10          477.0 ± 0%      256.0 ± 0%  -46.33% (p=0.000 n=10)       0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(1024)-10       1.310Ki ± 0%    1.000Ki ± 0%  -23.64% (p=0.000 n=10)   0.000Ki ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10          276.00 ± 0%      64.00 ± 0%  -76.81% (p=0.000 n=10)      0.00 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10          484.0 ± 0%      256.0 ± 0%  -47.11% (p=0.000 n=10)       0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(1024)-10       1.317Ki ± 0%    1.000Ki ± 0%  -24.09% (p=0.000 n=10)   0.000Ki ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10       281.0000 ± 0%    65.0000 ± 0%  -76.87% (p=0.000 n=10)    0.5000 ±  ?   -99.82% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10       489.000 ± 0%    257.000 ± 0%  -47.44% (p=0.000 n=10)     1.000 ±  ?   -99.80% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(1024)-10     1354.000 ± 0%   1025.000 ± 0%  -24.30% (p=0.000 n=10)     1.000 ±  ?   -99.93% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10          268.00 ± 0%      64.00 ± 0%  -76.12% (p=0.000 n=10)      0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10          477.0 ± 0%      256.0 ± 0%  -46.33% (p=0.000 n=10)       0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(1024)-10       1.310Ki ± 0%    1.000Ki ± 0%  -23.64% (p=0.000 n=10)   0.000Ki ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10          273.50 ± 0%      64.00 ± 0%  -76.60% (p=0.000 n=10)      0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10          482.0 ± 0%      256.0 ± 0%  -46.89% (p=0.000 n=10)       0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(1024)-10       1.315Ki ± 0%    1.000Ki ± 0%  -23.98% (p=0.000 n=10)   0.000Ki ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10        276.000 ± 0%     65.000 ± 0%  -76.45% (p=0.000 n=10)     1.000 ±  ?   -99.64% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10       484.000 ± 0%    257.000 ± 0%  -46.90% (p=0.000 n=10)     1.000 ±  ?   -99.79% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(1024)-10      1.317Ki ± 0%    1.001Ki ± 0%  -24.02% (p=0.000 n=10)   0.000Ki ±  ?  -100.00% (p=0.000 n=10)
geomean                                             562.4           256.6       -54.37%                                ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

                                             │ get-goredis.txt │          get-rueidis.txt           │           get-rueidiscsc.txt            │
                                             │    allocs/op    │ allocs/op   vs base                │ allocs/op   vs base                     │
/OneNode/Get-parall(1)-key(16)-val(64)-10           8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10          8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(1024)-10         8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10           8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10          8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(1024)-10         8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10          8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10         8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(1024)-10        8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10           8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10          8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(1024)-10         8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10           8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10          8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(1024)-10         8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10          8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10         8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(1024)-10        8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                                             8.000        1.000       -87.50%                              ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean
```
