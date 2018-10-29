package syntax_test

import (
	"log"
	"testing"
)

func TestAssignment(t *testing.T) {
	a := 20
	if 20 > 10 {
		a, b := 10, 20
		log.Println(a, b)
	}

	log.Println(a)

}

func TestReassignment(t *testing.T) {

	a := 20
	a, b := 15, 20

	if b > a {
		log.Println(a, b)
	}

	log.Println(a, b)
}

func TestSwap(t *testing.T) {
	aa := []int{1, 2, 3}
	log.Printf("slice: [%v]", aa)

	// swap
	aa[0], aa[2] = aa[2], aa[0]
	log.Printf("slice: [%v]", aa)
}
