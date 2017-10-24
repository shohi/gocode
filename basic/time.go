package main

import (
	"fmt"
	"strings"
	"time"
)

func printSleep() {
	aa := time.Now()
	time.Sleep(1 * time.Second)
	fmt.Println(aa)
	bb := time.Since(aa)
	bb = bb - (bb % time.Second)
	fmt.Println(bb)
}

func printDuration() {
	d := 1000 * time.Second
	fmt.Println(strings.ToUpper(d.String()))
}

func main() {
	printDuration()
}
