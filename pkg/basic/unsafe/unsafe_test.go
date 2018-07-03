package unsafe_test

import (
	"unsafe"
	"testing"
	"log"
)


func TestUnsafePointer(t *testing.T) {
	var a = 10
	log.Println(unsafe.Pointer(&a))
}

