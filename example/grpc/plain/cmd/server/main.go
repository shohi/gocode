package main

import (
	"flag"

	"github.com/shohi/gocode/example/grpc/plain/server"
)

var conf server.Config

func setupFlags() {
	flag.IntVar(&conf.Port, "port", 9001, "listen port")
	flag.Parse()
}

func main() {
	setupFlags()
	server.Run(conf)
}
