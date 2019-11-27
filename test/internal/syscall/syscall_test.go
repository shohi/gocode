package syscall_test

import (
	"log"
	"syscall"
	"testing"
)

func TestRLimit(t *testing.T) {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)

	log.Printf("====> ulimit: %v, err: %v", rLimit, err)
	log.Printf("1<<20: %v", 1<<20)
}
