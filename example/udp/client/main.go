package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")

	buf := make([]byte, 2048)
	_, err = bufio.NewReader(conn).Read(buf)
	if err == nil {
		fmt.Printf("%s\n", buf)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
}
