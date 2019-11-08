package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"

	"github.com/gorilla/websocket"
)

// ---------------------------------------gobwas/ws------------------------------------------

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	err = ws.WriteMessage(1, []byte("Gorilla ===> Hi Client!"))
	if err != nil {
		log.Println(err)
		return
	}
	// helpful log statement to show connections
	log.Println("Websocket Client Connected")
	readerWithGorillaWs(ws)
	ws.Close()
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func readerWithGorillaWs(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Printf("Gorilla Recv===> message type: [%v], conntent: [%v]\n", messageType, string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

// ---------------------------------------gobwas/ws------------------------------------------

func gwsEndpoint(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Println(err)
		return
	}

	err = wsutil.WriteServerMessage(conn, ws.OpText, []byte("Gobwas ===> Hi Client!"))
	if err != nil {
		log.Println(err)
		return
	}

	readerWithGobwasWS(conn)
	conn.Close()
}

func readerWithGobwasWS(conn net.Conn) {
	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Printf("Gobwas Recv===> message type: [%v], conntent: [%v]\n", op, string(msg))
		err = wsutil.WriteServerMessage(conn, op, msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// --------------------------------------- main ------------------------------------------
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc("/gws", gwsEndpoint)
}

func main() {
	setupRoutes()
	fmt.Println("server listening at - http://localhost:8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
