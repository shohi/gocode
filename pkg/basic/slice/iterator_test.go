package slice_test

import (
	"log"
	"strings"
	"testing"
)

func TestSlice_Iterator(t *testing.T) {

	aa := []string{"a", "b", "c", "d"}
	// iterate `key` only
	for k := range aa {
		log.Printf("k: [%v], v: [%v]", k, aa[k])
	}

	log.Printf(strings.Repeat("=", 60))
	// iterate `<key, value>` pair
	for k, v := range aa {
		log.Printf("k: [%v], v: [%v]", k, v)
	}

}
