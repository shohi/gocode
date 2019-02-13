package interface_test

import (
	"log"
	"testing"
	"time"
)

func TestIncompatibleCast(t *testing.T) {

	var info interface{}
	info = "hello"
	d, err := info.(time.Duration)

	log.Printf("duration: %v, err: %v", d, err)
}
