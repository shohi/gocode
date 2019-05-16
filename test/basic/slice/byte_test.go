package slice_test

import (
	"log"
	"testing"
)

func TestInitWithoutEnoughContent(t *testing.T) {
	var ss = [100]byte{77, 83, 71, 32}
	log.Printf("bytes: %v, string: %v", ss, string(ss[:]))
}
