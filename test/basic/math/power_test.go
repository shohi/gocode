package math_test

import (
	"log"
	"math"
	"testing"
)

func TestPower(t *testing.T) {
	// not power
	log.Printf("2^4: %v", 2^4)

	// this is math power
	log.Printf("2^4: %v", math.Pow(2, 4))
}
