package net_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListenWithoutAccept(t *testing.T) {
	assert := assert.New(t)
	ln, err := net.Listen("tcp", "127.0.0.1:0")

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				panic("err")
			}
			fmt.Printf("new connections arrival, address: %v\n", conn.RemoteAddr())
			time.Sleep(50 * time.Millisecond)
			_ = conn.Close()
		}
	}()
	// ln.(*net.TCPListener).Accept()
	assert.Nil(err)

	// FIXME: connection also can be established even through listner does not accept
	conn, err := net.DialTimeout("tcp", ln.Addr().String(), 20*time.Millisecond)
	fmt.Println("dial done")

	assert.Nil(err)
	n, err := conn.Write([]byte("Hello"))
	conn.SetReadDeadline(time.Now().Add(50 * time.Millisecond))
	buf := make([]byte, 1024)
	rn, rerr := conn.Read(buf)

	fmt.Printf("conn write, conn: %+v, remote: %v, listener: %v, count: %v, err: %v\n",
		conn, conn.RemoteAddr().String(), ln.Addr().String(),
		n, err)
	fmt.Printf("conn read, conn: %+v, count: %v, err: %v\n", conn, rn, rerr)
}
