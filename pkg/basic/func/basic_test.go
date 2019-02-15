package func_test

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

func TestFunctionWithNil(t *testing.T) {
	var test func(func())
	test = func(f func()) {
		log.Printf("%T", f)
		log.Println(f == nil)
	}

	test(nil)
	test(func() { log.Println("hello") })
}
