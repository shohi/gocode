package map_test

import (
	"log"
	"testing"
)

func TestMap_QueryByEmpty(t *testing.T) {
	m := make(map[string]string, 0)
	m["hello"] = "world"

	val, ok := m[""]
	log.Printf("entry for empty key: [%v], exist: [%v]", val, ok)
}
