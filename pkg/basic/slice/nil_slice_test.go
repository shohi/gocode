package slice

import (
	"log"
	"testing"
	"time"
)

func TestNilSliceTraverse(t *testing.T) {
	var a []*int

	for k, v := range a {
		log.Println(k, v)
	}

	bb := []string{"hello", "world"}
	for k := range bb {
		log.Println("key: ", k)
	}

	for k, v := range bb {
		log.Println("key: ", k, ", value: ", v)
	}
}

func TestNilSliceLen(t *testing.T) {
	var a []int
	log.Printf("nil slice length: %v", len(a))

	var tt time.Time
	log.Printf("time: %v", tt)

	current := time.Now()
	log.Printf("time after: %v", current.Add(-100*time.Hour).After(tt))

}
