package basic

import (
	"fmt"
	"log"
	"testing"
)

type state int

const (
	open state = iota
	halfopen
	closed
)

func map1() {
	m := make(map[string]bool)
	m["hello"] = true
	m["world"] = false

	for v := range m {
		fmt.Println(v)
	}
}

func map2() {
	m := make(map[string]int, 2)
	log.Println(len(m))
	for v := range m {
		fmt.Println(v)
	}
}

func TestEnum(t *testing.T) {
	fmt.Printf("%v, %T\n", open, open)
	fmt.Printf("%v, %T\n", halfopen, halfopen)
	fmt.Println(open == 0)
}
