package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Input struct {
	_ struct{} `type:"structure"`
	A string   `location:"us" type:"string"`
}

func TestReflect(t *testing.T) {
	input := &Input{A: "hello"}

	v := reflect.ValueOf(input)

	fmt.Printf("===> type: %v", v.Type())
	fmt.Printf("===> elem: %v", v.Elem())
	fmt.Printf("===> interface: %v", v.Interface())

	// print field info
	for k := 0; k < v.Elem().NumField(); k++ {
		fieldInfo := v.Elem().Type().Field(k)

		// Get field metadata
		// fmt.Println(fieldInfo)
		fmt.Println(fieldInfo.Tag.Get("type"))
		fmt.Println(fieldInfo.Name)
		fmt.Println(fieldInfo.Index)
		fmt.Println(fieldInfo.Offset)

		// Get field value
		fmt.Println(v.Elem().Field(k))
	}
}
