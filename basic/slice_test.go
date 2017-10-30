package basic

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var aa []int
	for i := 0; i < 5; i++ {
		aa = append(aa, 0)
	}

	fmt.Println(aa)
}
