package basic

import (
	"log"
	"testing"
)

func TestForRange(t *testing.T) {

	m := map[string]int{
		"NewYork":  1,
		"San Jose": 2,
		"Seattle":  3,
	}

	for range m {
		log.Println("hello")
	}
}

func TestForBreak(t *testing.T) {
loop:
	for i := 0; i < 10; i++ {
		select {
		default:
			log.Println(i)
			break loop
		}
	}
}

func TestForContinue(t *testing.T) {
	for i := 0; i < 10; i++ {
		select {
		default:
			log.Println(i)
			continue
		}
	}
}
