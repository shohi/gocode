package string_test

import (
	"fmt"
	"testing"
)

func TestStringUtil_Create(t *testing.T) {

	a := []byte{'a', 'b', 'c'}

	b := append([]byte(nil), ':')
	b = append(b, a...)

	fmt.Printf("===> a: %v, b:[%v]\n", string(a), string(b))

}
