package basic

import (
	"fmt"
	"os"
	"testing"
)

func hostname() {
	fmt.Println(os.Hostname())
}

func exit() {
	fmt.Println("going to exit")
	os.Exit(-21)
}

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("exit")
	}()
	exit()
	//
}
