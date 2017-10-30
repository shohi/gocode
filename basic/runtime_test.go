package basic

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestRuntime(t *testing.T) {
	go func() {
		time.Sleep(time.Second)
	}()
	buff := make([]byte, 10000)
	stackSize := runtime.Stack(buff, true)
	fmt.Println(string(buff[0:stackSize]))

}
