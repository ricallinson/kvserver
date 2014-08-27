package main

import (
	"flag"
	"fmt"
	"github.com/ricallinson/mapr"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "bench" {
		loadBenchmarker()
		return
	}
	var port = flag.String("port", "9090", "port to use for the server")
	flag.Parse()
	fmt.Printf("Server running on port %v\n", *port)
	s := mapr.NewServer(*port)
	s.Listen()
}
