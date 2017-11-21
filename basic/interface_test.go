package basic

import (
	"log"
	"testing"
	"unsafe"
)

func TestCompare(t *testing.T) {
	var s struct{}
	log.Println(unsafe.Sizeof(s))
	var i interface{}
	log.Println(unsafe.Sizeof(i))
	var b bool
	log.Println(unsafe.Sizeof(b))
}
