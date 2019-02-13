package flag_test

import (
	"log"
	"testing"
)

func TestArg(t *testing.T) {
	fn := func(strs ...string) []string {
		log.Printf("args len: %v", len(strs))

		sl := make([]string, len(strs))

		for k := 0; k < len(strs); k++ {
			sl[k] = strs[k]
		}

		return sl
	}

	sl := fn("1", "2", "3")

	log.Printf("values: %v", sl)
}
