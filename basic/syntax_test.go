package basic

import (
	"log"
	"testing"
)

func TestSyntaxAssignment(t *testing.T) {

	var a int
	a, b := 10, 20

	log.Println(a, b)

}
