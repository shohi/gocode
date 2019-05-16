package struct_test

import (
	"fmt"
	"testing"
)

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("%v", o.Quantity)
}

func TestStruct_Format(t *testing.T) {
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	// NOTE: orange did not implement String() method
	fmt.Println(orange)
}
