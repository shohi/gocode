package function

import (
	"log"
	"testing"
)

func Hello() {

}

func TestFunction(t *testing.T) {
	h := func() {}
	log.Printf("%p ==> %p", Hello, h)
}

func TestNilFunction(t *testing.T) {
	var f func()
	if f != nil {
		f()
	} else {
		t.Logf("Nil function")
	}
}
