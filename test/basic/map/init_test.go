package map_test

import (
	"log"
	"testing"
)

type event struct {
	code    int
	message string
}

func TestMap_InitStruct(t *testing.T) {
	m := map[int]event{
		// NOTE: `event` type is unnessesary when setting the value
		1: {message: "hello"},
	}
	log.Printf("event: %+v", m)
}

func TestMap_ZeroCap(t *testing.T) {
	m := make(map[string]string, 0)

	m["hello"] = "world"

	log.Printf("map: %v", m)
}

func TestMap_DumpEmpty(t *testing.T) {
	m := make(map[string]string, 0)

	log.Printf("map: %v", m)
}
