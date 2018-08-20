package variable

import (
	"errors"
	"log"
	"testing"
)

func testFn() (string, error) {
	return "value", errors.New("errors")
}

func TestShadow(t *testing.T) {
	var err error
	{
		val, err := testFn()
		log.Printf("value: %v, error: %v", val, err)
	}

	log.Printf("err is shadowed, err: %v", err)
}
