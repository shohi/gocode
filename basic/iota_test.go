package basic

import (
	"fmt"
	"testing"
)

type Code int

const (
	CodeNormal Code = iota + 900
	CodeErr
)

func TestIota(t *testing.T) {
	fmt.Printf("%v ==> %T", CodeErr, CodeErr)
}
