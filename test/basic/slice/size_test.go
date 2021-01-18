package slice_test

import (
	"fmt"
	"testing"
)

func TestSlice_Subslice(t *testing.T) {
	var a = []int{0, 1, 2, 3, 4, 5, 6}
	var n = len(a) - 1

	fmt.Printf("sub: [%v]\n", a[0:n])
}
