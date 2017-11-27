package basic

import (
	"fmt"
	"log"
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

func TestDeferWithMultipleAssignment(t *testing.T) {
	aa := 10
	defer func() {
		log.Println("aa ==> ", aa)
	}()

	aa = 12
}

func TestDeferWithArgument(t *testing.T) {
	aa := 10
	defer func(k int) {
		log.Println("argument ==> ", k)
	}(aa)

	aa = 12
}
