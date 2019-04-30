package map_test

import (
	"log"
	"testing"
)

type event struct {
	code    int
	message string
}

func TestInitStructMap(t *testing.T) {

	m := map[int]event{
		// NOTE: `event` type is unnessesary when setting the value
		1: {message: "hello"},
	}
	log.Printf("event: %+v", m)
}
