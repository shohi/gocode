package interface_test

import (
	"log"
	"testing"
)

type adder interface {
	add(i int) int
}

type Int int

func (i Int) add(j int) int {
	return int(i) + j
}

func TestInterface_Value(t *testing.T) {
	var myInt Int = 10
	var adder1 = adder(myInt)
	myInt = 20
	log.Printf("use interface by value: %v", adder1.add(10))
	log.Printf("use type by value: %v", myInt.add(10))

	var myInt2 Int = 10
	var adder2 = adder(&myInt2)
	myInt2 = 20
	log.Printf("use interface by ptr: %v", adder2.add(10))
	log.Printf("use type by ptr: %v", myInt2.add(10))
}
