package basic

import (
	"log"
	"strconv"
	"testing"
)

type state int

const (
	open state = iota
	halfopen
	closed
)

func TestMap1(t *testing.T) {
	m := make(map[string]bool)
	m["hello"] = true
	m["world"] = false

	for v := range m {
		log.Println(v)
	}
}

func TestMap2(t *testing.T) {
	m := make(map[string]int, 2)
	log.Println(len(m))

	for v := range m {
		log.Println(v)
	}
}

func TestMapNoInitialize(t *testing.T) {
	var m map[string]int

	log.Println(m)

	// map must be intialized before use
	m = make(map[string]int)
	for k := 0; k < 10; k++ {
		m[strconv.Itoa(k)] = k
	}

	for k, v := range m {
		log.Println(k, v)
	}
}

func TestEnum(t *testing.T) {
	log.Printf("%v, %T\n", open, open)
	log.Printf("%v, %T\n", halfopen, halfopen)
	log.Println(open == 0)
}
