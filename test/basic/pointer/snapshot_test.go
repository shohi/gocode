package pointer

import (
	"log"
	"testing"
)

type options struct {
	host string
	port int
}

func TestPointer_Snapshot(t *testing.T) {
	opts := &options{
		host: "localhost",
		port: 80,
	}

	opt2 := *opts

	opt2.port = 8080

	log.Printf("options: %v, snapshot: %v", opts, opt2)

}
