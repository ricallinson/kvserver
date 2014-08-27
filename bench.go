package main

import (
	"flag"
	"fmt"
	"github.com/ricallinson/mapr"
	"math/rand"
	"strconv"
	"time"
)

func randomStringMaker(size int) string {
	p := make([]byte, size)
	src := rand.New(rand.NewSource(78))
	for i := range p {
		p[i] = byte(src.Int63() & 0xff)
	}
	return string(p)
}

func executePut(m *mapr.Client, n int, id int, s string) {
	for i := 0; i < n; i++ {
		key := strconv.Itoa(id) + "_" + strconv.Itoa(i)
		m.Put(key, s)
	}
}

func executeGet(m *mapr.Client, n int, id int, s string) {
	for i := 0; i < n; i++ {
		key := strconv.Itoa(id) + "_" + strconv.Itoa(i)
		if m.Get(key) != s {
			fmt.Printf("Key [%v] was not found.\n", key)
		}
	}
}

func loadBenchmarker() {
	var p = flag.Int("p", 1000, "Size of the value to put in bytes")
	var n = flag.Int("n", 1000, "Total number of requests")
	var c = flag.Int("c", 50, "Number of parallel connections")
	flag.Parse()
	var host = flag.Arg(1)

	if host == "" {
		flag.PrintDefaults()
		return
	}

	s := randomStringMaker(*p)
	m := mapr.NewClient(host)

	startPut := time.Now()

	donePut := make(chan bool, *c)

	for i := 0; i < *c; i++ {
		go func(id int) {
			executePut(m, *n, id, s)
			donePut <- true
		}(i)
	}

	<-donePut

	elapsedPut := time.Since(startPut)
	startGet := time.Now()

	doneGet := make(chan bool, *c)

	for i := 0; i < *c; i++ {
		go func(id int) {
			executeGet(m, *n, id, s)
			doneGet <- true
		}(i)
	}

	<-doneGet

	elapsedGet := time.Since(startGet)
	total := *n * *c

	fmt.Printf("\n")
	fmt.Printf("%v keys inserted in %v\n", total, elapsedPut)
	fmt.Printf("Average time per put() %v\n", elapsedPut/time.Duration(total))
	fmt.Printf("Average inserts a second %d\n", time.Second/(elapsedPut/time.Duration(total)))
	fmt.Printf("\n")
	fmt.Printf("%v keys found in %v\n", total, elapsedGet)
	fmt.Printf("Average time per get() %v\n", elapsedGet/time.Duration(total))
	fmt.Printf("Average reads a second %d\n", time.Second/(elapsedGet/time.Duration(total)))
	fmt.Printf("\n")
}
