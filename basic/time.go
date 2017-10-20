package main

import (
	"fmt"
	"time"
)

func main() {
	aa := time.Now()
	time.Sleep(1 * time.Second)
	fmt.Println(aa)
	bb := time.Since(aa)
	bb = bb - (bb % time.Second)
	fmt.Println(bb)

}
