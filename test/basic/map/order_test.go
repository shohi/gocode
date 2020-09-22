package map_test

import (
	"fmt"
	"testing"
)

func TestMap_Order(t *testing.T) {
	m := map[string]int{
		"k1": 1,
		"k2": 2,
	}

	// first
	fmt.Println("first time ===> ")
	for k, v := range m {
		fmt.Printf("%v-%v\n", k, v)
	}

	// second
	fmt.Println("second time ===> ")
	for k, v := range m {
		fmt.Printf("%v-%v\n", k, v)
	}

}
