package basic

import (
	"fmt"
	"testing"
)

func TestDeferInLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer func(k int) {
			fmt.Println(k)
		}(i)

		if i > 5 {
			break
		}
	}
}
