package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal_Map(t *testing.T) {
	type MyStruct struct {
		Key string
	}

	var m = map[MyStruct]int{
		MyStruct{Key: "k1"}: 1,
		MyStruct{Key: "k2"}: 2,
	}

	data, err := json.Marshal(m)
	fmt.Printf("content: [%v], err: %v\n",
		string(data), err)

}

func TestMarshal_ArrayAndSlice(t *testing.T) {
	var arr [4]byte
	var ns []byte
	var es []byte = make([]byte, 0, 4)

	arrBytes, _ := json.Marshal(arr)
	nsBytes, _ := json.Marshal(ns)
	esBytes, _ := json.Marshal(es)

	fmt.Printf("array: %q, nil slice: %q, empty slice: %q\n",
		string(arrBytes),
		string(nsBytes),
		string(esBytes))
	// Output:
	// array: "[0,0,0,0]", nil slice: "null", empty slice: "\"\""
}
