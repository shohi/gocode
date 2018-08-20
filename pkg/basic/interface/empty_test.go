package interface_test

import (
	"log"
	"testing"
)

// Empty interface test
type S struct {
	i int
}

func (s *S) Get() int {
	return s.i
}

func (s *S) Put(v int) {
	s.i = v
}

type I interface {
	Get() int
	Put(int)
}

func g(i interface{}) int {
	return i.(I).Get()
}

func TestEmptyConversion(t *testing.T) {
	var s S
	log.Println(g(&s))

	// interface conversion failed. will panic
	// log.Println(g(s))
}
