package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func startRedisServer() *exec.Cmd {
	c := exec.Command("redis-server", "--port", "7777")
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	if err := c.Start(); err != nil {
		panic(err)
	}

	go func() {
		err := c.Wait()
		log.Printf("redis server exit, err: %v\n", err)
	}()

	log.Printf("redis server pid: %d\n", c.Process.Pid)

	return c
}

func signalCh() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	return ch
}

func main() {
	cmd := startRedisServer()

	// wait 10 second to kill
	log.Println("====> wait 10s to kill")
	time.Sleep(10 * time.Second)
	cmd.Process.Kill()

	// restart
	nc := startRedisServer()

	// wait 5 minutes

	sigCh := signalCh()
	select {
	case <-time.After(5 * time.Minute):
		nc.Process.Kill()
	case <-sigCh:
		nc.Process.Kill()
	}

	// stop process
	log.Println("=====> main exit in 10s")
	time.Sleep(10 * time.Second)
}
