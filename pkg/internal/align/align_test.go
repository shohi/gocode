package align_test

import (
	"log"
	"testing"
	"unsafe"
)

func TestAlign_Structure(t *testing.T) {
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

	type D struct {
		charF byte
		boolF bool
	}

	a := A{}
	b := B{}
	c := C{}
	d := D{}
	e := []byte{}

	log.Printf("A: (%v, %v)", unsafe.Sizeof(a), unsafe.Alignof(a))
	log.Printf("B: (%v, %v)", unsafe.Sizeof(b), unsafe.Alignof(b))
	log.Printf("C: (%v, %v)", unsafe.Sizeof(c), unsafe.Alignof(c))
	log.Printf("D: (%v, %v)", unsafe.Sizeof(d), unsafe.Alignof(d))
	log.Printf("E: (%v, %v)", unsafe.Sizeof(e), unsafe.Alignof(e))
}

func check(b, b2 bool, v *float32, v2 float32, v3 int16) {
	p1 := uintptr(unsafe.Pointer(&b))
	p2 := uintptr(unsafe.Pointer(&b2))
	p3 := uintptr(unsafe.Pointer(&v))
	p4 := uintptr(unsafe.Pointer(&v2))
	p5 := uintptr(unsafe.Pointer(&v3))
	log.Printf("Displacement: \n%v\n%v\n%v\n%v\n%v", p1, p2, p3, p4, p5)
}

// https://www.geeksforgeeks.org/structure-member-alignment-padding-and-data-packing/
func TestAlign_Argument(t *testing.T) {
	val := float32(1)
	check(true, false, &val, 1, 1)
}
