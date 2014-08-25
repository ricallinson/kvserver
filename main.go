package main

import (
	"flag"
	"fmt"
)

func main() {
	var port = flag.String("port", "9090", "port to use for the server")
	flag.Parse()
	fmt.Printf("Server running on port %v\n", *port)
	NewServer(*port)
}
