package strings_test

import (
	"log"
	"strings"
	"testing"
)

func TestStringRepeat(t *testing.T) {
	str := "na"
	log.Println("ba" + strings.Repeat(str, 2))
}
