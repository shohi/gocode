package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Inner struct {
	Key   string
	Value string
}

type Outer struct {
	Inner
}

func TestJson_Embed(t *testing.T) {
	var v = &Outer{}
	v.Key = "hello"
	v.Value = "world"

	b, _ := json.MarshalIndent(v, "", "  ")

	fmt.Printf("====> content: %v\n", string(b))

	var vv = Outer{}
	_ = json.Unmarshal(b, &vv)
	fmt.Printf("====> new content: %+v\n", vv)
}
