package basic

import (
	"log"
	"os"
	"testing"
)

func TestHostname(t *testing.T) {
	log.Println(os.Hostname())
}

func TestExit(t *testing.T) {
	log.Println("going to exit")
	os.Exit(-21)
}

func TestDefer(t *testing.T) {
	defer func() {
		log.Println("exit")
	}()
	os.Exit(-1)
}

func TestGetenv(t *testing.T) {
	log.Println(os.Getenv("GOPATH"))
}
