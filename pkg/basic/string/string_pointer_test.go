package string_test

import (
	"log"
	"testing"
)

func TestStringPointerConvert(t *testing.T) {
	var strptr *string
	var str string
	str = "hello world"

	// Must be initialized before using
	strptr = &str
	*strptr = "world"

	log.Println(*strptr)
	log.Println(str)
}

func TestStringPointerAddress(t *testing.T) {
	ids := []string{"hello", "world"}
	for _, id := range ids {
		// use new variable to avoid same address issue
		dd := id
		log.Printf("id ==> content - [%v], address - [%v]", dd, &dd)

		ddd := &id
		log.Printf("id ==> content - [%v], address - [%v]", *ddd, ddd)
	}
}
