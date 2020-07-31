package map_test

import (
	"fmt"
	"testing"
)

func TestMap_Nil_Traverse(t *testing.T) {
	var m map[string]interface{}

	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

}
