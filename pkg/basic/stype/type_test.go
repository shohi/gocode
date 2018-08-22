package stype

import (
	"log"
	"reflect"
	"testing"
)

type MyType struct{}

func TestType(t *testing.T) {
	var a interface{}
	b := MyType{}
	a = b

	typ := reflect.TypeOf(a)

	// var bb interface{}
	// panic: nil pointer
	// reflect.TypeOf(bb).Size()
	log.Printf("type: %v, name: %v, pkg: %v, size: %v", typ, typ.Name(), typ.PkgPath(), typ.Size())
	log.Printf("kind: %v", typ.Kind())
}

func myfn(a int, b string) {
	log.Printf("string: %v", b)
}

func TestMyFn(t *testing.T) {
	myfn(10, "string")
}
