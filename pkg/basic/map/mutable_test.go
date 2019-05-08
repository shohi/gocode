package map_test

import (
	"log"
	"testing"
)

// mutability propogation
func TestMutability(t *testing.T) {
	m := map[string][]string{
		"hello": []string{"ny", "la"},
	}

	mm := m
	mm["hello"] = append(mm["hello"], "bo")

	log.Printf("====> map values: %v", mm)
}
