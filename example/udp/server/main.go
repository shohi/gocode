package main

import (
	"log"
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	log.Printf("[server] listen: 127.0.0.1:3000")

	for {
		buf := make([]byte, 2048)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		log.Printf("[server] recieved: %s\n", buf[:n])
		go serve(pc, addr, buf[:n])
	}
}

// just echo response
func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	pc.WriteTo(buf, addr)
}
