package main

import "fmt"

func main() {
	var aa []int
	for i := 0; i < 5; i++ {
		aa = append(aa, 0)
	}

	fmt.Println(aa)
}
