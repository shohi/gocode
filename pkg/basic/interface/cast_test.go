package interface_test

import (
	"log"
	"testing"
	"time"
)

func TestIncompatibleCast(t *testing.T) {

	var info interface{}
	info = "hello"
	d, err := info.(time.Duration)

	log.Printf("duration: %v, err: %v", d, err)
}

type MyReader interface {
	reader()
}

type myReader struct{}

func (m *myReader) reader() {
	log.Printf("Hello")
}

func (m *myReader) name() {
	log.Print("yyReader")
}

func TestCastToInterface(t *testing.T) {

	fn := func() interface{} {
		return &myReader{}
	}
	info := fn()

	r, ok := info.(MyReader)

	log.Printf("%v, ok: %v", r, ok)

	r.reader()

}
