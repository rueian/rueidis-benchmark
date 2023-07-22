#!/bin/sh

go test -bench="Benchmark/.*/Set" -benchmem -benchtime=100000x -count=10 -lib=rueidis . | tee set-rueidis.txt
go test -bench="Benchmark/.*/Set" -benchmem -benchtime=100000x -count=10 -lib=goredis . | tee set-goredis.txt

go test -bench="Benchmark/.*/Get" -benchmem -benchtime=100000x -count=10 -lib=rueidis . | tee get-rueidis.txt
go test -bench="Benchmark/.*/Get" -benchmem -benchtime=100000x -count=10 -lib=rueicsc . | tee get-rueidiscsc.txt
go test -bench="Benchmark/.*/Get" -benchmem -benchtime=100000x -count=10 -lib=goredis . | tee get-goredis.txt
