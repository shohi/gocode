package slice_test

import (
	"fmt"
	"testing"
)

type Pair struct {
	Key   string
	Value interface{}
}

type Pairs []Pair

func (p *Pairs) Add(key string, value interface{}) {
	*p = append(*p, Pair{key, value})
}

func TestSlice_Modify(t *testing.T) {
	var tmp [4]Pair

	var a Pairs = tmp[:0]
	a.Add("k1", "v1")
	a.Add("k2", "v2")

	b := a
	b.Add("k3", "v3")

	fmt.Printf("a: %p ==> %+v\n", a, a)
	fmt.Printf("b: %p ==> %+v\n", b, b)
}
