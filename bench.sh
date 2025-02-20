#!/bin/sh

go test -bench="Benchmark/.*/Set" -benchmem -benchtime=100000x -count=10 -lib=rueidis . | tee set-3-rueidis.txt
go test -bench="Benchmark/.*/Set" -benchmem -benchtime=100000x -count=10 -lib=goredis . | tee set-1-goredis.txt
go test -bench="Benchmark/.*/Set" -benchmem -benchtime=100000x -count=10 -lib=glide . | tee set-2-glide.txt

go test -bench="Benchmark/.*/Get" -benchmem -benchtime=100000x -count=10 -lib=rueidis . | tee get-3-rueidis.txt
go test -bench="Benchmark/.*/Get" -benchmem -benchtime=100000x -count=10 -lib=rueicsc . | tee get-4-rueidiscsc.txt
go test -bench="Benchmark/.*/Get" -benchmem -benchtime=100000x -count=10 -lib=goredis . | tee get-1-goredis.txt
go test -bench="Benchmark/.*/Get" -benchmem -benchtime=100000x -count=10 -lib=glide . | tee get-2-glide.txt
