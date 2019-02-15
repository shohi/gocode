package rand_test

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rs := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(rs)

	delay := 10 * time.Second
	var maxVal = int(delay / time.Millisecond)

	log.Printf("max Val: %v", maxVal)
	for k := 0; k < 50; k++ {
		log.Printf("rand value: %v", rd.Intn(maxVal))
	}
}
