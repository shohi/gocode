package main

import "fmt"

func main() {
	ch := make(chan interface{})
	go func() {
		ch <- nil
	}()

	info := <-ch
	fmt.Println(info)
}
