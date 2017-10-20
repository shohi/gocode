package main

import "fmt"

type A interface {
	Hello() string
}

type SA struct {
	A
}

type IA struct {
	word string
}

func (a IA) Hello() string {
	return a.word
}

func main() {
	aa := SA{IA{"world"}}
	var bb A
	bb = &aa

	fmt.Println(aa)
	fmt.Println(bb)

}
