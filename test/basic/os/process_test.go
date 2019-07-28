package os_test

import (
	"log"
	"os"
	"testing"
)

func TestGetPid(t *testing.T) {

	log.Printf("pid: %v", os.Getpid())
}
