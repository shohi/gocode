package os_test

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"
)

func TestHostname(t *testing.T) {
	log.Println(os.Hostname())
}

func TestHostIP(t *testing.T) {
	net.InterfaceAddrs()
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}
}

func TestHostIP_Nonloop(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Printf("ip ===> %v\n", ipnet.IP)
			}
		}
	}
}
