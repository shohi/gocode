package map_test

import (
	"fmt"
	"log"
	"testing"
)

func TestInitNil(t *testing.T) {
	var m map[int]int

	log.Printf("%p", m)
	log.Printf("%p", &m)

	fmt.Println(m == nil)

	fn := func(m *map[int]int) {
		*m = make(map[int]int)
	}
	fn(&m)

	log.Printf("%p", m)
}
