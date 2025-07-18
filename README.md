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
                                            │ set-1-goredis.txt │           set-2-glide.txt            │          set-3-rueidis.txt           │
                                            │      sec/op       │    sec/op     vs base                │    sec/op     vs base                │
/OneNode/Set-parall(0)-key(16)-val(64)-10           21.54µ ± 0%    34.20µ ± 0%  +58.77% (p=0.000 n=10)   20.69µ ±  1%   -3.91% (p=0.000 n=10)
/OneNode/Set-parall(0)-key(16)-val(256)-10          22.37µ ± 1%    34.86µ ± 1%  +55.89% (p=0.000 n=10)   20.96µ ±  0%   -6.29% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(64)-10           9.965µ ± 1%    8.310µ ± 1%  -16.61% (p=0.000 n=10)   5.165µ ±  0%  -48.17% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10         10.098µ ± 0%    8.435µ ± 1%  -16.46% (p=0.000 n=10)   5.324µ ±  0%  -47.28% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10           7.335µ ± 2%    4.244µ ± 1%  -42.14% (p=0.000 n=10)   1.257µ ±  0%  -82.87% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10          7.504µ ± 2%    4.430µ ± 1%  -40.96% (p=0.000 n=10)   1.292µ ±  1%  -82.79% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10         7357.0n ± 2%   4271.0n ± 2%  -41.95% (p=0.000 n=10)   492.2n ±  1%  -93.31% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10        7529.5n ± 2%   4469.0n ± 2%  -40.65% (p=0.000 n=10)   624.2n ± 18%  -91.71% (p=0.000 n=10)
/Cluster/Set-parall(0)-key(16)-val(64)-10           22.49µ ± 1%    37.51µ ± 0%  +66.77% (p=0.000 n=10)   20.64µ ±  1%   -8.21% (p=0.000 n=10)
/Cluster/Set-parall(0)-key(16)-val(256)-10          22.33µ ± 0%    38.14µ ± 0%  +70.80% (p=0.000 n=10)   21.10µ ±  2%   -5.52% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10          10.036µ ± 1%    9.320µ ± 0%   -7.13% (p=0.000 n=10)   5.262µ ±  1%  -47.57% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10         10.149µ ± 1%    9.529µ ± 1%   -6.11% (p=0.000 n=10)   5.410µ ±  1%  -46.69% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10           7.479µ ± 3%    5.461µ ± 1%  -26.98% (p=0.000 n=10)   1.300µ ±  1%  -82.62% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10          7.712µ ± 2%    5.675µ ± 1%  -26.41% (p=0.000 n=10)   1.331µ ±  1%  -82.75% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10         7735.5n ± 2%   5087.5n ± 1%  -34.23% (p=0.000 n=10)   536.4n ±  3%  -93.07% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10        7604.5n ± 4%   5296.5n ± 1%  -30.35% (p=0.000 n=10)   673.4n ±  1%  -91.15% (p=0.000 n=10)
geomean                                             10.61µ         9.309µ       -12.23%                  3.013µ        -71.60%

                                            │ set-1-goredis.txt │           set-2-glide.txt            │             set-3-rueidis.txt              │
                                            │       B/op        │     B/op      vs base                │     B/op       vs base                     │
/OneNode/Set-parall(0)-key(16)-val(64)-10            266.0 ± 0%     176.0 ± 0%  -33.83% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(0)-key(16)-val(256)-10           266.0 ± 0%     176.0 ± 0%  -33.83% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(64)-10            267.0 ± 0%     176.0 ± 0%  -34.08% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10           267.0 ± 0%     176.0 ± 0%  -34.08% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10            275.0 ± 0%     176.0 ± 0%  -36.00% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10           275.0 ± 0%     176.0 ± 0%  -36.00% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10         280.000 ± 0%   179.000 ± 1%  -36.07% (p=0.000 n=10)   1.000 ±    0%   -99.64% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10        280.000 ± 0%   179.000 ± 0%  -36.07% (p=0.000 n=10)   1.000 ±    0%   -99.64% (p=0.000 n=10)
/Cluster/Set-parall(0)-key(16)-val(64)-10            266.0 ± 0%     176.0 ± 0%  -33.83% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(0)-key(16)-val(256)-10           266.0 ± 0%     176.0 ± 0%  -33.83% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10            267.0 ± 0%     176.0 ± 0%  -34.08% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10           267.0 ± 0%     176.0 ± 0%  -34.08% (p=0.000 n=10)     0.0 ±     ?  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10            272.0 ± 0%     176.0 ± 0%  -35.29% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10           272.0 ± 0%     176.0 ± 0%  -35.29% (p=0.000 n=10)     0.0 ±    0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10         274.000 ± 0%   179.000 ± 1%  -34.67% (p=0.000 n=10)   1.000 ± 1400%   -99.64% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10        274.000 ± 0%   179.000 ± 1%  -34.67% (p=0.000 n=10)   1.000 ±    0%   -99.64% (p=0.000 n=10)
geomean                                              270.8          176.7       -34.74%                                 ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

                                            │ set-1-goredis.txt │          set-2-glide.txt           │            set-3-rueidis.txt            │
                                            │     allocs/op     │ allocs/op   vs base                │ allocs/op   vs base                     │
/OneNode/Set-parall(0)-key(16)-val(64)-10            8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(0)-key(16)-val(256)-10           8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(64)-10            8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(1)-key(16)-val(256)-10           8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(64)-10            8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(8)-key(16)-val(256)-10           8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(64)-10           8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Set-parall(64)-key(16)-val(256)-10          8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(0)-key(16)-val(64)-10            8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(0)-key(16)-val(256)-10           8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(64)-10            8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(1)-key(16)-val(256)-10           8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(64)-10            8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(8)-key(16)-val(256)-10           8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(64)-10           8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Set-parall(64)-key(16)-val(256)-10          8.000 ± 0%   5.000 ± 0%  -37.50% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                                              8.000        5.000       -37.50%                              ?                       ¹ ²
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
                                            │ get-1-goredis.txt │            get-2-glide.txt            │           get-3-rueidis.txt           │         get-4-rueidiscsc.txt         │
                                            │      sec/op       │    sec/op      vs base                │    sec/op      vs base                │    sec/op     vs base                │
/OneNode/Get-parall(0)-key(16)-val(64)-10         21176.0n ± 1%   33966.5n ± 1%  +60.40% (p=0.000 n=10)   19783.0n ± 2%   -6.58% (p=0.000 n=10)   170.4n ± 33%  -99.20% (p=0.000 n=10)
/OneNode/Get-parall(0)-key(16)-val(256)-10        21366.0n ± 0%   34195.5n ± 0%  +60.05% (p=0.000 n=10)   20071.5n ± 2%   -6.06% (p=0.000 n=10)   168.7n ±  1%  -99.21% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(64)-10          9889.5n ± 1%    8314.5n ± 0%  -15.93% (p=0.000 n=10)    5018.5n ± 0%  -49.25% (p=0.000 n=10)   151.1n ±  8%  -98.47% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10         9889.5n ± 0%    8431.0n ± 0%  -14.75% (p=0.000 n=10)    5091.0n ± 0%  -48.52% (p=0.000 n=10)   152.8n ±  3%  -98.46% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10          7130.5n ± 1%    4323.0n ± 0%  -39.37% (p=0.000 n=10)    1170.5n ± 1%  -83.58% (p=0.000 n=10)   151.2n ±  2%  -97.88% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10         7344.5n ± 1%    4394.5n ± 0%  -40.17% (p=0.000 n=10)    1210.5n ± 1%  -83.52% (p=0.000 n=10)   143.4n ± 14%  -98.05% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10         7209.5n ± 3%    4375.5n ± 3%  -39.31% (p=0.000 n=10)     506.3n ± 2%  -92.98% (p=0.000 n=10)   143.7n ± 14%  -98.01% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10        7381.0n ± 2%    4522.5n ± 2%  -38.73% (p=0.000 n=10)     561.9n ± 1%  -92.39% (p=0.000 n=10)   149.2n ±  7%  -97.98% (p=0.000 n=10)
/Cluster/Get-parall(0)-key(16)-val(64)-10         21808.5n ± 0%   37371.5n ± 0%  +71.36% (p=0.000 n=10)   20367.5n ± 2%   -6.61% (p=0.000 n=10)   178.1n ±  1%  -99.18% (p=0.000 n=10)
/Cluster/Get-parall(0)-key(16)-val(256)-10        22117.5n ± 0%   37584.5n ± 0%  +69.93% (p=0.000 n=10)   20628.0n ± 0%   -6.73% (p=0.000 n=10)   178.6n ±  1%  -99.19% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10          9945.0n ± 1%    9331.0n ± 0%   -6.17% (p=0.000 n=10)    5142.0n ± 0%  -48.30% (p=0.000 n=10)   170.2n ±  2%  -98.29% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10         9960.5n ± 0%    9418.5n ± 0%   -5.44% (p=0.000 n=10)    5202.0n ± 0%  -47.77% (p=0.000 n=10)   166.1n ±  2%  -98.33% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10          7430.5n ± 1%    5545.0n ± 1%  -25.38% (p=0.000 n=10)    1213.5n ± 1%  -83.67% (p=0.000 n=10)   163.2n ±  2%  -97.80% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10         7561.5n ± 2%    5598.0n ± 0%  -25.97% (p=0.000 n=10)    1240.0n ± 1%  -83.60% (p=0.000 n=10)   163.2n ±  3%  -97.84% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10         7444.0n ± 2%    5202.0n ± 2%  -30.12% (p=0.000 n=10)     530.9n ± 3%  -92.87% (p=0.000 n=10)   164.4n ±  2%  -97.79% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10        7631.5n ± 2%    5239.0n ± 1%  -31.35% (p=0.000 n=10)     557.6n ± 2%  -92.69% (p=0.000 n=10)   163.2n ±  1%  -97.86% (p=0.000 n=10)
geomean                                             10.40µ          9.312µ       -10.48%                    2.864µ       -72.47%                  160.7n        -98.46%

                                            │ get-1-goredis.txt │            get-2-glide.txt            │          get-3-rueidis.txt           │           get-4-rueidiscsc.txt           │
                                            │       B/op        │     B/op       vs base                │     B/op      vs base                │    B/op      vs base                     │
/OneNode/Get-parall(0)-key(16)-val(64)-10           280.00 ± 0%     288.00 ± 0%   +2.86% (p=0.000 n=10)     64.00 ± 0%  -77.14% (p=0.000 n=10)     0.00 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(0)-key(16)-val(256)-10           488.0 ± 0%      672.0 ± 0%  +37.70% (p=0.000 n=10)     256.0 ± 0%  -47.54% (p=0.000 n=10)      0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(64)-10           281.00 ± 0%     288.00 ± 0%   +2.49% (p=0.000 n=10)     64.00 ± 0%  -77.22% (p=0.000 n=10)     0.00 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10           489.0 ± 0%      672.0 ± 0%  +37.42% (p=0.000 n=10)     256.0 ± 0%  -47.65% (p=0.000 n=10)      0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10           289.00 ± 0%     288.00 ± 0%   -0.35% (p=0.000 n=10)     64.00 ± 0%  -77.85% (p=0.000 n=10)     0.00 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10           497.0 ± 0%      673.0 ± 0%  +35.41% (p=0.000 n=10)     256.0 ± 0%  -48.49% (p=0.000 n=10)      0.0 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10        294.0000 ± 0%   291.0000 ± 0%   -1.02% (p=0.000 n=10)   65.0000 ± 0%  -77.89% (p=0.000 n=10)   0.5000 ±  ?   -99.83% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10          502.0 ± 0%      676.0 ± 0%  +34.66% (p=0.000 n=10)     257.0 ± 0%  -48.80% (p=0.000 n=10)      0.0 ±  ?  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(0)-key(16)-val(64)-10           280.00 ± 0%     288.00 ± 0%   +2.86% (p=0.000 n=10)     64.00 ± 0%  -77.14% (p=0.000 n=10)     0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(0)-key(16)-val(256)-10           488.0 ± 0%      672.0 ± 0%  +37.70% (p=0.000 n=10)     256.0 ± 0%  -47.54% (p=0.000 n=10)      0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10           281.00 ± 0%     288.00 ± 0%   +2.49% (p=0.000 n=10)     64.00 ± 0%  -77.22% (p=0.000 n=10)     0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10           489.0 ± 0%      672.0 ± 0%  +37.42% (p=0.000 n=10)     256.0 ± 0%  -47.65% (p=0.000 n=10)      0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10           286.00 ± 0%     288.00 ± 0%   +0.70% (p=0.000 n=10)     64.00 ± 0%  -77.62% (p=0.000 n=10)     0.00 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10           495.0 ± 0%      672.0 ± 0%  +35.76% (p=0.000 n=10)     256.0 ± 0%  -48.28% (p=0.000 n=10)      0.0 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10          288.00 ± 0%     290.50 ± 0%   +0.87% (p=0.000 n=10)     65.00 ± 0%  -77.43% (p=0.000 n=10)     0.00 ±  ?  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10          497.0 ± 0%      675.0 ± 0%  +35.81% (p=0.000 n=10)     257.0 ± 0%  -48.29% (p=0.000 n=10)      0.0 ±  ?  -100.00% (p=0.000 n=10)
geomean                                              374.8           440.8       +17.61%                    128.3       -65.76%                               ?                       ¹ ²
¹ summaries must be >0 to compute geomean
² ratios must be >0 to compute geomean

                                            │ get-1-goredis.txt │           get-2-glide.txt           │         get-3-rueidis.txt          │          get-4-rueidiscsc.txt           │
                                            │     allocs/op     │ allocs/op   vs base                 │ allocs/op   vs base                │ allocs/op   vs base                     │
/OneNode/Get-parall(0)-key(16)-val(64)-10            7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(0)-key(16)-val(256)-10           7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(64)-10            7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(1)-key(16)-val(256)-10           7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(64)-10            7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(8)-key(16)-val(256)-10           7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(64)-10           7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/OneNode/Get-parall(64)-key(16)-val(256)-10          7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(0)-key(16)-val(64)-10            7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(0)-key(16)-val(256)-10           7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(64)-10            7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(1)-key(16)-val(256)-10           7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(64)-10            7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(8)-key(16)-val(256)-10           7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(64)-10           7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
/Cluster/Get-parall(64)-key(16)-val(256)-10          7.000 ± 0%   7.000 ± 0%       ~ (p=1.000 n=10) ¹   1.000 ± 0%  -85.71% (p=0.000 n=10)   0.000 ± 0%  -100.00% (p=0.000 n=10)
geomean                                              7.000        7.000       +0.00%                    1.000       -85.71%                              ?                       ² ³
¹ all samples are equal
² summaries must be >0 to compute geomean
³ ratios must be >0 to compute geomean

```
