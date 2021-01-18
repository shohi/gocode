package reflect_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestReflect_Law(t *testing.T) {
	fmt.Printf("%v\n", reflect.TypeOf(10))
	fmt.Printf("%v\n", reflect.TypeOf(reflect.TypeOf(10)))

	var i int = 10
	rv := reflect.ValueOf(&i)
	// rv.Elem().SetInt(100)
	rv.Elem().Set(reflect.ValueOf(100))

	fmt.Printf("%v\n", i)
}

func TestMarshal_Panic(t *testing.T) {
	var i int
	var ch chan string
	_, _ = i, ch

	var arr = []interface{}{
		unsafe.Pointer(&i),
		json.Marshal,
		ch,
	}

	data, err := json.Marshal(arr)
	fmt.Printf("%s, err: %v\n", string(data), err)
}
