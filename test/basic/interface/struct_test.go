package interface_test

import (
	"fmt"
	"testing"
)

type MyIface interface {
	Value() int
}

type myIface struct{}

func (myIface) Value() int { return 0 }

type MyStruct struct {
	MyIface
	Name string
}

func TestInterface_EmbedInStruct(t *testing.T) {
	var ms MyStruct
	ms.MyIface = myIface{}

	fmt.Printf("===> Hello: %v", ms)
}
