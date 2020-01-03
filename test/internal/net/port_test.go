package net_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// refer-1, [udp-server using net.ListenUDP](https://gist.github.com/miguelmota/01ba5131838ae31947ac9b03e57f3773)
// refer-2, [udp-server using net.ListenPacket](https://gist.github.com/miekg/d9bc045c89578f3cc66a214488e68227)

func TestTwoUPDSockets_SamePort(t *testing.T) {
	assert := assert.New(t)

	addr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 5556,
	}

	conn1, err := net.ListenUDP("udp", addr)
	assert.Nil(err)
	defer func() {
		if conn1 != nil {
			conn1.Close()
		}
	}()

	conn2, err := net.ListenUDP("udp", addr)
	defer func() {
		if conn2 != nil {
			conn2.Close()
		}
	}()

	assert.NotNil(err)
}

// refer, https://jamison.dance/05-27-2018/multiple-sockets-on-the-same-port
func TestTcpUdpSockets_SamePort(t *testing.T) {
	assert := assert.New(t)

	addr := &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 5557,
	}

	uc, err := net.ListenUDP("udp", addr)
	assert.Nil(err)
	defer uc.Close()

	tc, err := net.Listen("tcp", "127.0.0.1:5556")
	assert.Nil(err)
	defer tc.Close()

}
