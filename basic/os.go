package main

import (
	"fmt"
	"os"
)

func hostname() {
	fmt.Println(os.Hostname())
}

func exit() {
	fmt.Println("going to exit")
	os.Exit(-21)
}

func main() {
	defer func() {
		fmt.Println("exit")
	}()
	exit()
	//
}
