package main

import (
	"io"
	"log"
	"net/http"

	"github.com/posener/h2conn"
)

func main() {
	// 在 8000 端口启动服务器
	// 确切地说，如何运行HTTP/1.1服务器。
	mux := http.NewServeMux()
	// mux.HandleFunc("/", handleBasic)
	mux.HandleFunc("/", handleAdvance)
	mux.HandleFunc("/echo", echo)
	srv := &http.Server{Addr: ":8000", Handler: mux}

	// 用TLS启动服务器，因为我们运行的是http/2，它必须是与TLS一起运行。
	// 确切地说，如何使用TLS连接运行HTTP/1.1服务器。
	log.Printf("Serving on https://0.0.0.0:8000")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func handleBasic(w http.ResponseWriter, r *http.Request) {
	// 记录请求协议
	log.Printf("Got connection: %s", r.Proto)
	// 向客户发送一条消息
	w.Write([]byte("Hello"))
}

func handleAdvance(w http.ResponseWriter, r *http.Request) {
	// Log the request protocol
	log.Printf("Got connection: %s", r.Proto)

	// Handle 2nd request, must be before push to prevent recursive calls.
	// Don't worry - Go protect us from recursive push by panicking.
	if r.URL.Path == "/2nd" {
		log.Println("Handling 2nd")
		w.Write([]byte("Hello Again!"))
		return
	}

	// Handle 1st request
	log.Println("Handling 1st")

	// Server push must be before response body is being written.
	// In order to check if the connection supports push, we should use
	// a type-assertion on the response writer.
	// If the connection does not support server push, or that the push
	// fails we just ignore it - server pushes are only here to improve
	// the performance for HTTP/2 clients.
	pusher, ok := w.(http.Pusher)
	if !ok {
		log.Println("Can't push to client")
	} else {
		err := pusher.Push("/2nd", nil)
		if err != nil {
			log.Printf("Failed push: %v", err)
		}
	}

	// Send response body
	w.Write([]byte("Hello"))
}

func echo(w http.ResponseWriter, r *http.Request) {
	// Accept returns a connection to the client  that can be used:
	//   1. Write - send data to the client
	//   2. Read - receive data from the client
	conn, err := h2conn.Accept(w, r)
	if err != nil {
		log.Printf("Failed creating connection from %s: %s", r.RemoteAddr, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Send back to the client everything that we receive
	io.Copy(conn, conn)
}
