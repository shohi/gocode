package align_test

import (
	"log"
	"testing"
	"unsafe"
)

func TestChanAlign(t *testing.T) {
	ch := make(chan struct{})
	log.Printf("channel - size: %v, align %v", unsafe.Sizeof(ch), unsafe.Alignof(ch))
}
