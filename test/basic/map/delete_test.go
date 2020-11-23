package map_test

import (
	"log"
	"testing"
)

func TestMapDelete(t *testing.T) {

	m := make(map[string]int, 1024)

	log.Printf("len: %v", len(m))
}

func TestMap_DeleteNonExist(t *testing.T) {
	m := make(map[string]int, 1024)

	// If m is nil or there is no such element, delete is a no-op.
	delete(m, "non-exist")
}
