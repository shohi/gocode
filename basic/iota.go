package main

import "fmt"

type Code int

const (
	CodeNormal Code = iota + 900
	CodeErr
)

func main() {
	fmt.Printf("%v ==> %T", CodeErr, New(CodeErr))
}
