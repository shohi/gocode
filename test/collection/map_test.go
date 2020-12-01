package collection_test

import (
	"fmt"
	"testing"
)

func TestSliceMap_Grow(t *testing.T) {
	m := make(map[string][]int)
	a := []int{1, 2, 3, 4}
	m["a"] = a

	for i := 0; i < 5; i++ {
		s := m["a"]
		s = append(s, i+10)
	}

	// NOTE: m will NOT change
	fmt.Printf("map: %v\nslice: %v\n", m, a)
}
