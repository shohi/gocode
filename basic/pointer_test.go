package basic

import (
	"log"
	"testing"
)

func TestDereference(t *testing.T) {

	var strPtr *string
	str := "hello"

	// *strPtr = "hello"
	strPtr = &str
	*strPtr = str

	log.Println(strPtr)
	log.Println(*strPtr)

}
