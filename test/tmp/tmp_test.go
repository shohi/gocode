package tmp_test

import (
	"log"
	"testing"
)

func TestSliceRange(t *testing.T) {
	s := []int{1, 2, 3, 4}
	log.Printf("slice - [%v]", s)

	aa := s[:1]
	log.Printf("slice - [%v]", aa)
}
