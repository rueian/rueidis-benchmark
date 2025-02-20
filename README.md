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
cpu: Apple M1 Pro
                                            │ set-1-goredis.txt │           set-2-glide.txt            │          set-3-rueidis.txt           │
                                            │      sec/op       │    sec/op     vs base                │    sec/op     vs base                │
/OneNode/Set-parall(1)-key(16)-val(64)-10          10.969µ ± 1%    8.601µ ± 1%  -21.59% (p=0.000 n=10)   5.563µ ±  1%  -49.29% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10         11.122µ ± 0%    8.793µ ± 0%  -20.94% (p=0.000 n=10)   5.752µ ±  1%  -48.28% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10           8.050µ ± 2%    4.708µ ± 1%  -41.52% (p=0.000 n=10)   1.364µ ±  1%  -83.06% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10          8.323µ ± 1%    4.897µ ± 1%  -41.17% (p=0.000 n=10)   1.392µ ±  1%  -83.28% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10         7997.0n ± 1%   4734.5n ± 1%  -40.80% (p=0.000 n=10)   557.8n ±  2%  -93.02% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10        8272.0n ± 2%   4966.5n ± 2%  -39.96% (p=0.000 n=10)   620.7n ± 14%  -92.50% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10          10.970µ ± 1%    8.640µ ± 0%  -21.24% (p=0.000 n=10)   5.677µ ±  0%  -48.25% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10         11.137µ ± 1%    8.823µ ± 0%  -20.77% (p=0.000 n=10)   5.844µ ±  1%  -47.53% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10           8.500µ ± 1%    4.723µ ± 1%  -44.44% (p=0.000 n=10)   1.389µ ±  1%  -83.66% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10          8.698µ ± 2%    4.944µ ± 1%  -43.15% (p=0.000 n=10)   1.427µ ±  2%  -83.59% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10         8499.5n ± 1%   4752.0n ± 1%  -44.09% (p=0.000 n=10)   594.8n ±  4%  -93.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10        8702.5n ± 1%   4963.5n ± 1%  -42.96% (p=0.000 n=10)   677.0n ± 12%  -92.22% (p=0.000 n=10)
geomean                                             9.187µ         5.884µ       -35.95%                  1.694µ        -81.56%

                                            │ set-1-goredis.txt │           set-2-glide.txt            │            set-3-rueidis.txt            │
                                            │       B/op        │     B/op      vs base                │    B/op     vs base                     │
/OneNode/Set-parall(1)-key(16)-val(64)-10            273.0 ± 0%     156.0 ± 0%  -42.86% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10           273.0 ± 0%     156.0 ± 0%  -42.86% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10            281.0 ± 0%     156.0 ± 0%  -44.48% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10           281.0 ± 0%     156.0 ± 0%  -44.48% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10         286.000 ± 0%   157.000 ± 0%  -45.10% (p=0.000 n=10)   1.000 ± 0%   -99.65% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10        286.000 ± 0%   157.000 ± 0%  -45.10% (p=0.000 n=10)   1.000 ± 0%   -99.65% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10            273.0 ± 0%     156.0 ± 0%  -42.86% (p=0.000 n=10)     0.0 ±  ?  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10           273.0 ± 0%     156.0 ± 0%  -42.86% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10            278.0 ± 0%     156.0 ± 0%  -43.88% (p=0.000 n=10)     0.0 ± 0%  -100.00% (p=0.000 n=10)
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

Comparing to the go-redis baseline, rueidis can achieve about **92%**, **76%** and **87%** reduction on `sec/op`, `B/op` and `allocs/op` respectively.

With the client-side caching, rueidis can achieve about **97%**, **99%** and **100%** reduction on `sec/op`, `B/op` and `allocs/op` respectively in cache hit scenario.

```shell
▶ benchstat get*
goos: darwin
goarch: arm64
pkg: rueidis-benchmark
cpu: Apple M1 Pro
                                            │ get-1-goredis.txt │           get-2-glide.txt            │          get-3-rueidis.txt           │        get-4-rueidiscsc.txt         │
                                            │      sec/op       │    sec/op     vs base                │    sec/op     vs base                │   sec/op     vs base                │
/OneNode/Get-parall(1)-key(16)-val(64)-10        11070.0n ±  0%   8227.5n ± 0%  -25.68% (p=0.000 n=10)   5453.0n ± 2%  -50.74% (p=0.000 n=10)   156.7n ± 3%  -98.58% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10       10937.0n ±  1%   8337.5n ± 0%  -23.77% (p=0.000 n=10)   5526.5n ± 1%  -49.47% (p=0.000 n=10)   153.8n ± 2%  -98.59% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10         7964.5n ±  1%   4421.0n ± 1%  -44.49% (p=0.000 n=10)   1290.5n ± 1%  -83.80% (p=0.000 n=10)   155.0n ± 2%  -98.05% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10        8194.0n ±  2%   4459.5n ± 0%  -45.58% (p=0.000 n=10)   1335.5n ± 1%  -83.70% (p=0.000 n=10)   155.8n ± 2%  -98.10% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10        7802.5n ±  5%   4475.0n ± 1%  -42.65% (p=0.000 n=10)    586.2n ± 2%  -92.49% (p=0.000 n=10)   155.8n ± 4%  -98.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10       7970.5n ±  5%   4594.5n ± 1%  -42.36% (p=0.000 n=10)    641.5n ± 2%  -91.95% (p=0.000 n=10)   157.3n ± 3%  -98.03% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10        10953.0n ±  1%   8231.5n ± 0%  -24.85% (p=0.000 n=10)   5541.5n ± 0%  -49.41% (p=0.000 n=10)   176.9n ± 3%  -98.39% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10       11046.5n ±  1%   8357.0n ± 0%  -24.35% (p=0.000 n=10)   5627.5n ± 1%  -49.06% (p=0.000 n=10)   173.5n ± 2%  -98.43% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10         7338.5n ± 15%   4439.0n ± 1%  -39.51% (p=0.000 n=10)   1325.5n ± 1%  -81.94% (p=0.000 n=10)   171.0n ± 2%  -97.67% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10        7465.0n ±  2%   4461.0n ± 1%  -40.24% (p=0.000 n=10)   1348.0n ± 2%  -81.94% (p=0.000 n=10)   168.3n ± 3%  -97.74% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10        7430.5n ±  1%   4449.5n ± 1%  -40.12% (p=0.000 n=10)    591.8n ± 1%  -92.04% (p=0.000 n=10)   170.8n ± 2%  -97.70% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10       7558.0n ±  2%   4590.0n ± 1%  -39.27% (p=0.000 n=10)    631.5n ± 2%  -91.64% (p=0.000 n=10)   172.6n ± 3%  -97.72% (p=0.000 n=10)
geomean                                            8.680µ          5.504µ       -36.59%                   1.650µ       -80.99%                  163.7n       -98.11%

                                            │ get-1-goredis.txt │           get-2-glide.txt           │         get-3-rueidis.txt          │          get-4-rueidiscsc.txt          │
                                            │       B/op        │    B/op      vs base                │    B/op     vs base                │   B/op     vs base                     │
/OneNode/Get-parall(1)-key(16)-val(64)-10           285.00 ± 0%   264.00 ± 0%   -7.37% (p=0.000 n=10)   64.00 ± 0%  -77.54% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10           493.0 ± 0%    648.0 ± 0%  +31.44% (p=0.000 n=10)   256.0 ± 0%  -48.07% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10           293.00 ± 0%   264.00 ± 0%   -9.90% (p=0.000 n=10)   64.00 ± 0%  -78.16% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10           501.0 ± 0%    648.0 ± 0%  +29.34% (p=0.000 n=10)   256.0 ± 0%  -48.90% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10          298.00 ± 0%   265.00 ± 0%  -11.07% (p=0.000 n=10)   65.00 ± 0%  -78.19% (p=0.000 n=10)   0.00 ±  ?  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10          506.0 ± 0%    650.0 ± 0%  +28.46% (p=0.000 n=10)   257.0 ± 0%  -49.21% (p=0.000 n=10)    0.0 ±  ?  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10           285.00 ± 0%   264.00 ± 0%   -7.37% (p=0.000 n=10)   64.00 ± 0%  -77.54% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10           493.0 ± 0%    648.0 ± 0%  +31.44% (p=0.000 n=10)   256.0 ± 0%  -48.07% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10           290.00 ± 0%   264.00 ± 0%   -8.97% (p=0.000 n=10)   64.00 ± 0%  -77.93% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10           499.0 ± 0%    648.0 ± 0%  +29.86% (p=0.000 n=10)   256.0 ± 0%  -48.70% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10          292.00 ± 0%   265.00 ± 0%   -9.25% (p=0.000 n=10)   65.00 ± 0%  -77.74% (p=0.000 n=10)   0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10          501.0 ± 0%    650.0 ± 0%  +29.74% (p=0.000 n=10)   257.0 ± 0%  -48.70% (p=0.000 n=10)    0.0 ± 0%  -100.00% (p=0.000 n=10)
geomean                                              380.6         414.1        +8.79%                  128.4       -66.26%                             ?                       ¹ ²
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
