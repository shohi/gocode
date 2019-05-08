package align_test

import (
	"log"
	"testing"
	"unsafe"
)

func TestAlign(t *testing.T) {
	type A struct {
		floatF float32
		charF  byte
		boolF  bool
	}

	type B struct {
		boolF bool
		intF  int32
		charF byte
	}

	type C struct {
		intF  int32
		data  []byte
		boolF bool
	}

	a := A{}
	b := B{}
	c := C{}

	log.Printf("A: (%v, %v), B: (%v, %v), C: (%v, %v)",
		unsafe.Sizeof(a),
		unsafe.Alignof(a),
		unsafe.Sizeof(b),
		unsafe.Alignof(b),
		unsafe.Sizeof(c),
		unsafe.Alignof(c),
	)
}
