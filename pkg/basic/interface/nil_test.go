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

type result struct {
	err  error
	code int
}

func TestNilInterface(t *testing.T) {

	fn := func() *result {
		var r *result = nil
		if 2 > 3 {
			r = &result{}
		}
		return r
	}
	var rr *result
	rr = &result{}
	rr = nil

	r := fn()
	log.Printf("is nil ==> %v", r == nil)
	log.Printf("is nil ==> %v", rr == nil)
	log.Printf("%v", r)
}
