package runtime_test

import (
	"strings"
	"testing"
)

func TestGC(t *testing.T) {

	var x = strings.Repeat("hello", 10000)

	_ = x
}
