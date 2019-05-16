package syscall_test

import (
	"log"
	"syscall"
	"testing"
)

// Test for retrieving resource limits
func TestSysCall_Rlimit(t *testing.T) {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		t.Fatalf("failed to call GETRLIMIT, err: %v", err)
	}
	log.Printf("resource limits: [%v]", rLimit)
}
