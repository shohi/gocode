package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect_TypeEqual(t *testing.T) {
	type A struct{}
	type B struct{}

	var a interface{} = A{}
	var b interface{} = B{}
	var c interface{} = &A{}

	typA := reflect.TypeOf(a)
	typB := reflect.TypeOf(b)
	typC := reflect.TypeOf(c)

	fmt.Printf("type of a: %v\n", typA)
	fmt.Printf("type of b: %v\n", typB)
	fmt.Printf("type of c: %v\n", typC)

	fmt.Printf("a.Type == b.Type: %v\n", typA == typB)
	fmt.Printf("a.Type == c.Type: %v\n", typA == typC)
}

func TestReflectType_Byte(t *testing.T) {
	type MyInt int

	var a byte
	var b MyInt
	fmt.Printf("a type: %v, addr type: %v\n",
		reflect.TypeOf(a),
		reflect.TypeOf(&a),
	)

	fmt.Printf("b type: %v, addr type: %v\n",
		reflect.TypeOf(b),
		reflect.TypeOf(&b),
	)

}
