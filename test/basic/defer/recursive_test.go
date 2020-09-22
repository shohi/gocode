package defer_test

import (
	"fmt"
	"testing"
)

// output
// - f(3)
// - f(2)
// - f(1)
// - defer 1
// - defer 2
// - defer 3
// panic
func TestDefer_Recursive(t *testing.T) {
	var f func(x int)
	defer func() {
		fmt.Println("main")
	}()

	f = func(x int) {
		fmt.Printf("f(%d)\n", x+0/x)
		defer fmt.Printf("defer %d\n", x)
		f(x - 1)
	}

	f(3)
}
