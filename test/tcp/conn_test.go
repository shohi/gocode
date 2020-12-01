package tcp_test

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTcpServer(port int) {
	net.Listen("", address string)
	// net.Listen()
}

func TestWriteOnClosedConn(t *testing.T) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	server := httptest.NewServer(http.HandlerFunc(fn))
	defer server.Close()
	addr := strings.TrimPrefix(server.URL, "http://")

	conn, err := net.Dial("tcp", addr)
	fmt.Printf("====> establish connection - err: %v\n", err)

	err = conn.Close()
	fmt.Printf("====> close connection - err: %v\n", err)

	n, err := conn.Write([]byte("hello"))
	fmt.Printf("====> write to closed connection - err: %v, count: %v\n", err, n)
}
