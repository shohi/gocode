package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
	}()
	buff := make([]byte, 10000)
	stackSize := runtime.Stack(buff, true)
	fmt.Println(string(buff[0:stackSize]))

}
