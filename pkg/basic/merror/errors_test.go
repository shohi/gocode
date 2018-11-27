package merror_test

import (
	"log"
	"testing"

	"github.com/pkg/errors"
)

func TestErrorsFormat(t *testing.T) {
	err := errors.New("error1")
	err2 := errors.Wrap(err, "error2")

	log.Printf("[%v]", err2)
}

func TestTypeAssert(t *testing.T) {
	fn := func() (interface{}, error) {
		return nil, nil
	}

	res, _ := fn()

	data := res.([]byte)
	log.Printf("%v", data)
}

const (
	a1 int = 10 + iota

	a2

	a3
)

func TestIOTA(t *testing.T) {
	log.Printf("value ==> <%v, %v, %v>", a1, a2, a3)
}
