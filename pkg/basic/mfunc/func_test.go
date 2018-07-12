package mfunc

import (
	"log"
	"testing"
)

func TestFuncEqual(t *testing.T) {
	aa := func() {}

	bb := aa
	log.Printf("func equal: %v, %v", bb, aa)
}
