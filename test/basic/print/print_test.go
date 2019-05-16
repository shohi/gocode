package print

import (
	"log"
	"testing"
)

type Level int

func TestTypeOfVariable(t *testing.T) {
	l := Level(0)

	log.Printf("%T", l)
}
