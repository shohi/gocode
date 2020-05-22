package struct_test

import (
	"fmt"
	"testing"
)

type MyOtherStruct struct {
	ID string
}

type MyNewStruct MyOtherStruct

func TestStruct_Alias(t *testing.T) {

	a := MyOtherStruct{
		ID: "id",
	}
	/* not work
	s := MyNewStruct{
		ID: "hello"
	}
	*/

	s := MyNewStruct(a)

	fmt.Printf("struct: %+v\n", s)
}
