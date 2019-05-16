package panic_test

import (
	"log"
	"testing"
)

func TestPanic(t *testing.T) {
	a := 10
	b := 0

	defer func() {
		if err := recover(); err != nil {
			log.Printf("%v", err)
		}
	}()

	log.Printf("value: %v", a/b)
}

// ref, https://medium.com/@thedevsaddam/go-101-defer-panic-and-recover-65a40ee7dcb4

func TestPanicRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	panic("unable to run program")
}
