package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello h2c")
}

// net/http包默认可以采用http2进行服务，在没有进行https的服务上开启H2，
// 需要修改ListenAndServer的默认h2服务
// refer, https://www.jianshu.com/p/ff16b0308e7c

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)

	// http2.Server.ServeConn()
	s2 := &http2.Server{
		IdleTimeout: 1 * time.Minute,
	}

	l, _ := net.Listen("tcp", ":8972")
	defer l.Close()

	srv := &http.Server{
		Addr:    ":8972",
		Handler: mux,
	}

	log.Printf("Serving on http://0.0.0.0:8972")
	for {
		rwc, err := l.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		// Use h2 to serve connection
		go s2.ServeConn(rwc, &http2.ServeConnOpts{BaseConfig: srv})

	}
}
