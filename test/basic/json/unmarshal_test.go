package json_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestJsonUnmarsh_Array(t *testing.T) {
	var arr []*string
	fmt.Printf("val:%v, arr=nil: %v\n",
		reflect.ValueOf(arr), arr == nil)

	var val = reflect.ValueOf(arr)
	_ = val

	var data = `["a","b","c"]`
	err := json.Unmarshal([]byte(data), &arr)
	fmt.Printf("arr: %v, err: %v\n", arr, err)
}
