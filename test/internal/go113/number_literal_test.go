package go113_test

import (
	"log"
	"testing"
	"time"
)

func TestNumberLiteral(t *testing.T) {
	d := time.Duration(87_910_189 * time.Nanosecond)

	log.Printf("ms: %v", d.Milliseconds())
}
