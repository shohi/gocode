package slice

import (
	"log"
	"testing"
)

func TestNilSliceTraverse(t *testing.T) {
	var a []*int

	for k, v := range a {
		log.Println(k, v)
	}

	bb := []string{"hello", "world"}
	for k := range bb {
		log.Println("key: ", k)
	}

	for k, v := range bb {
		log.Println("key: ", k, ", value: ", v)
	}
}
