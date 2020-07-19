package defer_test

import (
	"fmt"
	"testing"
)

func TestDefer_When(t *testing.T) {
	fn := func() (r int) {
		t := 5
		defer func() {
			t = t + 5
		}()

		return t
	}

	fmt.Println(fn())
}

func TestDefer_Order(t *testing.T) {
	f1 := func() {
		a := 1
		defer func() {
			fmt.Println("fn1 defer")
		}()
		_ = a
	}

	f2 := func() {
		a := 2
		defer func() {
			fmt.Println("fn2 defer")
		}()
		_ = a
	}

	f3 := func() {
		a := 3
		defer func() {
			fmt.Println("fn3 defer")
		}()
		_ = a
	}

	f1()
	f2()
	f3()
}
