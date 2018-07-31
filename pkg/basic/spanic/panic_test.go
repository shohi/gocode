package spanic

import (
	"log"
	"testing"
)

func TestPanic(t *testing.T) {
	a := 10
	b := 0

	defer func() {
		if err := recover(); err != nil {
			log.Printf("%v", err)
		}
	}()

	log.Printf("value: %v", a/b)
}
