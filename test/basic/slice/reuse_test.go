package slice_test

import (
	"log"
	"testing"
)

func TestSliceReuse(t *testing.T) {
	aa := []string{"hello", "world"}

	bb := aa[:0]
	// NOTE: will change slice `aa` as well
	bb = append(bb, "space")

	log.Printf("original: %v, sub: %v", aa, bb)
}
