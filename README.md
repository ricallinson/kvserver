# KV Server

Simple key value server written in Go.

## Build

	go build

## Run

	kvserver -port 9090

## Debuging

	http://localhost:9090/debug/rpc

## Testing

First you'll need to __get__ the testing library. You only have to do this once.

    go get github.com/ricallinson/mapr
	go get github.com/ricallinson/simplebdd

### Run Tests

	go test

### Run Benchmark

	go test -bench .
