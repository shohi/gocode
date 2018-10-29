package interface_test

import (
	"log"
	"testing"
)

type RunnableFunc func()

func (r RunnableFunc) AndThen(other RunnableFunc) RunnableFunc {
	return func() {
		if r != nil {
			r()
		}
		if other != nil {
			other()
		}
	}
}

func TestNilAsReceiver(t *testing.T) {
	var current = RunnableFunc(func() {
		log.Printf("Hello current")
	})
	var then = RunnableFunc(func() {
		log.Printf("Hello Then")
	})

	d := current.AndThen(then)
	d()
}
