package net_test

import (
	"fmt"
	"net"
	"testing"
)

func TestSplit(t *testing.T) {
	hostport := "localhost"
	h, p, err := net.SplitHostPort(hostport)

	fmt.Printf("===> host: %v, port: %v, err: %v\n",
		h, p, err)
}
