package main

import "log"

var x = '京'

func init() { x = 2 ^ 15 }
func init() { x = 4 ^ 15 }

type T struct {
	n *int
}

func 京() int {
	ch := make(chan []T, x)
	close(ch)
	return copy(<-ch, <-ch)
}

func 都() (n *int) {
	if n := 京() >> 0.0; n < 1 {
		return &n
	}
	return n
}

func main() {
	v := 都()
	log.Println(v == nil)
	log.Printf("ptr: %v, val: %v", v, *v)
}
