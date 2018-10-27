package string_test

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestStringFromNumeric(t *testing.T) {
	// not work
	// log.Println(string(10))

	aa := 10
	log.Println(strconv.Itoa(aa))
}

func TestStringFromBytes(t *testing.T) {
	// case 1 - string from nil byte slice
	var a []byte
	a = nil
	b := string(a)
	log.Printf("String from nil byte slice equal to empty: %v", b == "")

	// case 2 - string from non-nil byte slice
	bs := []byte("ABCDE")
	log.Printf("%v ==> %v\n", string(bs), bs)
}

func TestStringType(t *testing.T) {
	c := '/'
	s := "/"
	log.Println(fmt.Sprintf("%T", c))
	log.Println(fmt.Sprintf("%T", s))

	log.Println(string(c))
}
