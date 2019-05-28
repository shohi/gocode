package print

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

type Level int

func TestTypeOfVariable(t *testing.T) {
	l := Level(0)

	log.Printf("%T", l)
}

func TestSprint(t *testing.T) {
	var v = 123
	log.Printf("%v", fmt.Sprint(v))
	log.Printf("%v", strconv.FormatUint(uint64(v), 10))
}
