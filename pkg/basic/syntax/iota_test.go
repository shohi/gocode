package syntax_test

import (
	"log"
	"testing"
)

const (
	a1 int = 10 + iota
	a2
	a3
)

func TestIOTA(t *testing.T) {
	log.Printf("value ==> <%v, %v, %v>", a1, a2, a3)
}

type Code int

const (
	CodeNormal Code = iota + 900
	CodeErr

	CodeTimeout
)

func TestIota2(t *testing.T) {
	log.Printf("%v ==> %T", CodeErr, CodeErr)
	log.Printf("%v ==> %T", CodeTimeout, CodeTimeout)
}
