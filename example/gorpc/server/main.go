package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/shohi/gocode/example/gorpc/core"
)

func main() {
	// Publish our Handler methods
	rpc.Register(&core.Handler{})

	l, err := net.Listen("tcp", ":6001")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	// Wait for incoming connections
	rpc.Accept(l)
}
