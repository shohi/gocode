package slice_test

import (
	"log"
	"testing"
)

func TestInterfaceSlice(t *testing.T) {
	var ss []interface{}
	var str []byte = []byte("hello")
	ss = append(ss, 10)
	ss = append(ss, str)
	log.Printf("ss content: [%v]", ss)
}
