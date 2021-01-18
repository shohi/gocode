package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal_Cyclic(t *testing.T) {
	type Node struct {
		Name string
		Next *Node
	}

	root := &Node{Name: "root"}
	root.Next = root

	m, err := json.Marshal(root)
	fmt.Printf("content: %q, err: %v\n",
		string(m), err)

	// Output
	// content: "", err: json: unsupported value: encountered a cycle via *json_test.Node
}
