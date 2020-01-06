package os_test

import (
	"log"
	"os"
	"testing"
)

func TestHostname(t *testing.T) {
	log.Println(os.Hostname())
}
