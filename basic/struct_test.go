package basic

import (
	"fmt"
	"log"
	"testing"
)

type A interface {
	Hello() string
}

type SA struct {
	A
}

type EmptyStruct struct{}

type IA struct {
	word string
}

func (a IA) Hello() string {
	return a.word
}

func TestStruct(t *testing.T) {
	aa := SA{IA{"world"}}
	var bb A
	bb = &aa

	fmt.Println(aa)
	fmt.Println(bb)

}

func TestEmptyStruct(t *testing.T) {
	ss := EmptyStruct{}
	fmt.Println(ss)

}

//
type BaseStruct struct {
}

func (s *BaseStruct) SayHello() {
	panic("Hello BaseStruct")
}

type DeriveStruct struct {
	BaseStruct
}

func (s *DeriveStruct) SayHello() {
	log.Println("Hello DeriveStruct")
}
func TestEmbeddedStruct(t *testing.T) {
	base := BaseStruct{}
	derive := DeriveStruct{BaseStruct: base}

	derive.SayHello()
}
