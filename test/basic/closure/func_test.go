package closure_test

import (
	"fmt"
	"testing"
)

// https://github.com/polaris1119/golangweekly/blob/master/docs/issue-073.md
func TestClosure_Func(t *testing.T) {
	var x int
	inc := func() int {
		x++
		return x
	}
	// Out: 1 2
	fmt.Println(func() (a, b int) {
		return inc(), inc()
	}())
}
