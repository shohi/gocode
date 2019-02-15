package func_test

import (
	"testing"
)

func TestFuncEqual(t *testing.T) {
	aa := func() {}

	bb := aa
	// Error
	// log.Printf("func equal: %v, %v", bb, aa)
	bb()
	aa()
}
