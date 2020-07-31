package map_test

import (
	"fmt"
	"testing"
)

type Fields map[string]int

func (fs Fields) Appends(more Fields) {
	for k, v := range more {
		fs[k] = v
	}
}

func TestMap_Expand(t *testing.T) {
	m := make(Fields, 10)
	fn := func(p string, n int) map[string]int {
		m := make(map[string]int, n)
		for k := 0; k < n; k++ {
			m[fmt.Sprintf("%v-%d", p, k)] = k
		}

		return m
	}

	fmt.Printf("%p\n", m)
	m.Appends(fn("1", 100))
	fmt.Printf("%p\n", m)
	fmt.Printf("%v\n", m)

	m.Appends(fn("1", 100000))
	fmt.Printf("%p, len: %v\n", m, len(m))
}
