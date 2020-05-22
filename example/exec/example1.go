package main

import (
	"fmt"
	"os/exec"
	"time"
)

// https://medium.com/@felixge/killing-a-child-process-and-all-of-its-children-in-go-54079af94773

func main() {
	cmd := exec.Command("sleep", "5")
	start := time.Now()
	time.AfterFunc(3*time.Second, func() { cmd.Process.Kill() })
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%s\n",
		cmd.Process.Pid, time.Since(start), err)
}
