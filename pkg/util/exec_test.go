package util

import (
	"log"
	"testing"
)

func TestExec_RunShellCommand(t *testing.T) {
	buf, err := RunShellCommand("psls", "~")
	log.Printf("err: %v, buf: %v", err, buf)
}
