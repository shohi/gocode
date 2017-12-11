package basic

import (
	"log"
	"testing"
)

func TestInterfaceForNil(t *testing.T) {

	tFunc := func(val interface{}) {
		log.Println(val == nil)
		log.Println(val)
	}

	tFunc(nil)
	tFunc("hello")
	tFunc(12)
	tFunc(false)
}

func TestInterfaceSwitchWithBreakReturn(t *testing.T) {
	tFunc := func(val interface{}) {
		switch t := val.(type) {
		case int:
			log.Println("int ==> ", t)
			return
		case string:
			log.Println("string ==> ", t)
			return
		case int32:
			log.Println("int32 ==> ", t)
			break
		}

		log.Println("tFunc end")
	}

	tFunc(32)
}
