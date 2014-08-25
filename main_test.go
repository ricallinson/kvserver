package main

import (
	. "github.com/ricallinson/simplebdd"
	"strconv"
	"testing"
	"time"
    "github.com/ricallinson/mapr"
)

func TestServer(t *testing.T) {

	// Create a Server for all tests to run against.
	go mapr.NewServer("9090")
	// Wait while the server starts.
	time.Sleep(1 * time.Second)

	Describe("Put() Get()", func() {
		It("should return [true]", func() {
			// Create a client.
			c := mapr.NewClient("127.0.0.1:9090")
			// Put a Value
			AssertEqual(c.Put("a", "b"), true)
			// Get a Value
			AssertEqual(c.Get("a"), "b")
		})
	})

	Report(t)
}

func BenchmarkServer(b *testing.B) {
	c := mapr.NewClient("127.0.0.1:9090")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		c.Put(key, key)
		c.Get(key)
	}
}
