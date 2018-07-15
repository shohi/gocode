package slice

import (
	"log"
	"testing"
)

func TestAppendWithInitCap(t *testing.T) {
	b := make([]int, 1024)
	b = append(b, 99)
	log.Println("len:", len(b), "cap:", cap(b))
}

func TestAppendWithNil(t *testing.T) {
	var s []*int
	var a *int
	s = append(s, a)
	log.Printf("slice: %v, len: %v, cap: %v", s, len(s), cap(s))

}
