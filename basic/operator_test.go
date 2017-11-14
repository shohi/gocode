package basic

import (
	"fmt"
	"testing"
)

func TestBetweenOperator(t *testing.T) {
	a := 10

	if 1 < a && a < 20 {
		fmt.Println(a)
	}

}

func TestParallelAssign(t *testing.T) {
	a, b := 10, 20

	fmt.Println(a, b)
}
