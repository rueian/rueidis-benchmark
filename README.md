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

Comparing to the go-redis baseline, rueidis can achieve about **92%**, **99%** and **100%** reduction on `sec/op`, `B/op` and `allocs/op` respectively.

```shell
▶ benchstat set*
goos: darwin
goarch: arm64
pkg: rueidis-benchmark
cpu: Apple M1 Pro
                                            │ set-1-goredis.txt │           set-2-glide.txt            │          set-3-rueidis.txt          │
                                            │      sec/op       │    sec/op     vs base                │   sec/op     vs base                │
/OneNode/Set-parall(1)-key(16)-val(64)-10          10.993µ ± 1%    8.639µ ± 0%  -21.41% (p=0.000 n=10)   5.585µ ± 1%  -49.20% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10         11.208µ ± 1%    8.835µ ± 3%  -21.18% (p=0.000 n=10)   5.741µ ± 1%  -48.78% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10           8.071µ ± 1%    4.717µ ± 1%  -41.55% (p=0.000 n=10)   1.362µ ± 1%  -83.12% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10          8.236µ ± 1%    4.944µ ± 1%  -39.97% (p=0.000 n=10)   1.399µ ± 1%  -83.01% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10         7964.0n ± 1%   4751.5n ± 1%  -40.34% (p=0.000 n=10)   562.4n ± 3%  -92.94% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10        8103.0n ± 3%   5020.5n ± 1%  -38.04% (p=0.000 n=10)   624.3n ± 1%  -92.30% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10          11.052µ ± 1%    9.932µ ± 0%  -10.14% (p=0.000 n=10)   5.693µ ± 1%  -48.49% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10         11.165µ ± 0%   10.134µ ± 0%   -9.23% (p=0.000 n=10)   5.829µ ± 1%  -47.79% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10           8.486µ ± 1%    6.314µ ± 1%  -25.60% (p=0.000 n=10)   1.410µ ± 2%  -83.38% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10          8.674µ ± 1%    6.655µ ± 1%  -23.28% (p=0.000 n=10)   1.432µ ± 1%  -83.50% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10         8493.5n ± 1%   6156.5n ± 2%  -27.52% (p=0.000 n=10)   592.4n ± 3%  -93.03% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10        8662.5n ± 1%   6495.5n ± 2%  -25.02% (p=0.000 n=10)   674.4n ± 3%  -92.21% (p=0.000 n=10)
geomean                                             9.169µ         6.627µ       -27.72%                  1.698µ       -81.48%

                                            │ set-1-goredis.txt │           set-2-glide.txt            │            set-3-rueidis.txt            │
                                            │       B/op        │     B/op      vs base                │    B/op     vs base                     │
/OneNode/Set-parall(1)-key(16)-val(64)-10            273.0 ± 0%     156.0 ± 0%  -42.86% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10           273.0 ± 0%     156.0 ± 0%  -42.86% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10            281.0 ± 0%     156.0 ± 0%  -44.48% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10           281.0 ± 0%     156.0 ± 0%  -44.48% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10         286.000 ± 0%   157.000 ± 0%  -45.10% (p=0.000 n=10)   1.000 ± 0%   -99.65% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10        286.000 ± 0%   157.000 ± 0%  -45.10% (p=0.000 n=10)   1.000 ± 0%   -99.65% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10            273.0 ± 0%     156.0 ± 0%  -42.86% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10           273.0 ± 0%     156.0 ± 0%  -42.86% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10          278.000 ± 0%   156.000 ± 0%  -43.88% (p=0.000 n=10)   7.000 ±  ?   -97.48% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10           278.0 ± 0%     156.0 ± 0%  -43.88% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10         280.000 ± 0%   157.000 ± 0%  -43.93% (p=0.000 n=10)   1.000 ± 0%   -99.64% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10        280.000 ± 0%   157.000 ± 0%  -43.93% (p=0.000 n=10)   1.000 ± 0%   -99.64% (p=0.000 n=10)
geomean                                              278.5          156.3       -43.86%                              ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

                                            │ set-1-goredis.txt │          set-2-glide.txt           │            set-3-rueidis.txt            │
                                            │     allocs/op     │ allocs/op   vs base                │ allocs/op   vs base                     │
/OneNode/Set-parall(1)-key(16)-val(64)-10            9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10           9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10            9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10           9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10           9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10          9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10            9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10           9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10            9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10           9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10           9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10          9.000 ± 0%   6.000 ± 0%  -33.33% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                                              9.000        6.000       -33.33%                              ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

```

## Redis GET

Comparing to the go-redis baseline, rueidis can achieve about **92%**, **77%** and **87%** reduction on `sec/op`, `B/op` and `allocs/op` respectively.

With the client-side caching, rueidis can achieve about **98%**, **100%** and **100%** reduction on `sec/op`, `B/op` and `allocs/op` respectively in cache hit scenario.

```shell
▶ benchstat get*
goos: darwin
goarch: arm64
pkg: rueidis-benchmark
cpu: Apple M1 Pro
                                            │ get-1-goredis.txt │            get-2-glide.txt            │          get-3-rueidis.txt           │        get-4-rueidiscsc.txt         │
                                            │      sec/op       │    sec/op      vs base                │    sec/op     vs base                │   sec/op     vs base                │
/OneNode/Get-parall(1)-key(16)-val(64)-10         11041.0n ± 1%    8612.0n ± 0%  -22.00% (p=0.000 n=10)   5465.5n ± 1%  -50.50% (p=0.000 n=10)   156.9n ± 4%  -98.58% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10        11066.5n ± 1%    8753.5n ± 1%  -20.90% (p=0.000 n=10)   5532.0n ± 1%  -50.01% (p=0.000 n=10)   154.4n ± 3%  -98.60% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10          8043.5n ± 1%    4692.5n ± 2%  -41.66% (p=0.000 n=10)   1307.0n ± 2%  -83.75% (p=0.000 n=10)   156.3n ± 2%  -98.06% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10         8211.0n ± 1%    4747.5n ± 2%  -42.18% (p=0.000 n=10)   1337.0n ± 1%  -83.72% (p=0.000 n=10)   155.1n ± 3%  -98.11% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10         7153.0n ± 2%    4792.5n ± 1%  -33.00% (p=0.000 n=10)    585.9n ± 2%  -91.81% (p=0.000 n=10)   158.0n ± 3%  -97.79% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10        7294.5n ± 1%    4905.0n ± 1%  -32.76% (p=0.000 n=10)    643.8n ± 2%  -91.17% (p=0.000 n=10)   158.1n ± 6%  -97.83% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10          9683.0n ± 0%    9895.0n ± 1%   +2.19% (p=0.000 n=10)   5578.0n ± 1%  -42.39% (p=0.000 n=10)   175.4n ± 3%  -98.19% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10         9826.0n ± 4%   10022.5n ± 1%        ~ (p=0.143 n=10)   5639.0n ± 1%  -42.61% (p=0.000 n=10)   173.0n ± 3%  -98.24% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10          8068.0n ± 2%    6334.5n ± 3%  -21.49% (p=0.000 n=10)   1324.5n ± 3%  -83.58% (p=0.000 n=10)   169.1n ± 2%  -97.90% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10         8262.5n ± 4%    6373.0n ± 1%  -22.87% (p=0.000 n=10)   1362.5n ± 2%  -83.51% (p=0.000 n=10)   167.6n ± 2%  -97.97% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10         8532.0n ± 1%    5962.5n ± 1%  -30.12% (p=0.000 n=10)    587.3n ± 2%  -93.12% (p=0.000 n=10)   169.5n ± 2%  -98.01% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10        8725.0n ± 2%    5614.0n ± 2%  -35.66% (p=0.000 n=10)    627.9n ± 1%  -92.80% (p=0.000 n=10)   174.7n ± 2%  -98.00% (p=0.000 n=10)
geomean                                             8.740µ          6.461µ       -26.08%                   1.653µ       -81.08%                  163.8n       -98.13%

                                            │ get-1-goredis.txt │           get-2-glide.txt           │         get-3-rueidis.txt          │          get-4-rueidiscsc.txt          │
                                            │       B/op        │    B/op      vs base                │    B/op     vs base                │   B/op     vs base                     │
/OneNode/Get-parall(1)-key(16)-val(64)-10           285.00 ± 0%   264.00 ± 0%   -7.37% (p=0.000 n=10)   64.00 ± 0%  -77.54% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10           493.0 ± 0%    648.0 ± 0%  +31.44% (p=0.000 n=10)   256.0 ± 0%  -48.07% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10           293.00 ± 0%   264.00 ± 0%   -9.90% (p=0.000 n=10)   64.00 ± 0%  -78.16% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10           501.0 ± 0%    648.0 ± 0%  +29.34% (p=0.000 n=10)   256.0 ± 0%  -48.90% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10          298.00 ± 0%   266.00 ± 0%  -10.74% (p=0.000 n=10)   65.00 ± 0%  -78.19% (p=0.000 n=10)   0.00 ±  ?  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10          506.0 ± 0%    651.0 ± 0%  +28.66% (p=0.000 n=10)   257.0 ± 0%  -49.21% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10           285.00 ± 0%   264.00 ± 0%   -7.37% (p=0.000 n=10)   64.00 ± 0%  -77.54% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10           493.0 ± 0%    648.0 ± 0%  +31.44% (p=0.000 n=10)   256.0 ± 0%  -48.07% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10           290.00 ± 0%   264.00 ± 0%   -8.97% (p=0.000 n=10)   64.00 ± 0%  -77.93% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10           499.0 ± 0%    648.0 ± 0%  +29.86% (p=0.000 n=10)   256.0 ± 0%  -48.70% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10          292.00 ± 0%   265.00 ± 0%   -9.25% (p=0.000 n=10)   65.00 ± 0%  -77.74% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10          501.0 ± 0%    650.0 ± 0%  +29.74% (p=0.000 n=10)   257.0 ± 0%  -48.70% (p=0.000 n=10)    0.0 ±  ?  -100.00% (p=0.000 n=10)
geomean                                              380.6         414.3        +8.83%                  128.4       -66.26%                             ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

                                            │ get-1-goredis.txt │          get-2-glide.txt           │         get-3-rueidis.txt          │          get-4-rueidiscsc.txt           │
                                            │     allocs/op     │ allocs/op   vs base                │ allocs/op   vs base                │ allocs/op   vs base                     │
/OneNode/Get-parall(1)-key(16)-val(64)-10            8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10           8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10            8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10           8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10           8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10          8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10            8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10           8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10            8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10           8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10           8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10          8.000 ± 0%   6.000 ± 0%  -25.00% (p=0.000 n=10)   1.000 ± 0%  -87.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                                              8.000        6.000       -25.00%                  1.000       -87.50%                              ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

```
