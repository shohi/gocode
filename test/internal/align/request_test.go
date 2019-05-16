package align_test

import (
	"log"
	"net/http"
	"testing"
	"unsafe"
)

func TestRequestAlign(t *testing.T) {
	var r http.Request
	var r2 http.Request
	var r3 http.Request

	log.Printf("request - 1, size: %v, align %v, pointer: %p", unsafe.Sizeof(r), unsafe.Alignof(r), &r)
	log.Printf("request - 2, size: %v, align %v, pointer: %p", unsafe.Sizeof(r2), unsafe.Alignof(r2), &r2)
	log.Printf("request - 3, size: %v, align %v, pointer: %p", unsafe.Sizeof(r3), unsafe.Alignof(r3), &r3)
}
