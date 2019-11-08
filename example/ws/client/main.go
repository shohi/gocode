package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// refer, https://github.com/gorilla/websocket/blob/master/examples/echo/client.go
var wsURL string

func main() {
	flag.StringVar(&wsURL, "ws", "ws://localhost:8080/ws", `Websocket server address, e.g. "ws://localhost:8080/ws"`)

	flag.Parse()

	conn, resp, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()
	log.Printf("===> response: [%v]", resp)

	doneCh := make(chan struct{})

	go readLoop(conn, doneCh)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	writeLoop(conn, doneCh, sigCh)
}

func readLoop(c *websocket.Conn, doneCh chan struct{}) {
	defer close(doneCh)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}

}

func writeLoop(c *websocket.Conn, doneCh chan struct{}, interruptCh chan os.Signal) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-doneCh:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interruptCh:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-doneCh:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
