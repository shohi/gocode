package string_test

import (
	"log"
	"strconv"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	s := "ABCDE"
	log.Printf("%v ==> %v\n", ([]byte)(s), s)
}

func TestStringAndBytes(t *testing.T) {
	str := "12"
	b := []byte(str)

	log.Println(str, b)

	b = []byte{0x01, 0x02}
	str = string(b)

	log.Println(str, b)
}

func TestStringFromNIL(t *testing.T) {
	var a []byte
	a = nil
	b := string(a)
	log.Println(b == "")
}

func TestStringConvert(t *testing.T) {
	aa := 10
	log.Println(strconv.Itoa(aa))
}
func TestStringFromInt(t *testing.T) {
	// not work
	log.Println(string(10))

	//
	log.Println(strconv.Itoa(10))
}

func TestStringSlice(t *testing.T) {
	ids := []string{"hello", "world"}
	for _, id := range ids {
		// use new variable to avoid same address issue
		dd := id
		log.Printf("id ==> content - [%v], address - [%v]", dd, &dd)

		ddd := &id
		log.Printf("id ==> content - [%v], address - [%v]", *ddd, ddd)
	}
}

func TestStringOutput(t *testing.T) {
	str := `"hello world"`
	log.Printf("%s\n", str)
}
