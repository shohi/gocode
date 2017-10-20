package main

import (
	"fmt"
	"unsafe"
)

func compare() {
	var s struct{}
	fmt.Println(unsafe.Sizeof(s))
	var i interface{}
	fmt.Println(unsafe.Sizeof(i))
	var b bool
	fmt.Println(unsafe.Sizeof(b))
}

func main() {
	compare()
}
