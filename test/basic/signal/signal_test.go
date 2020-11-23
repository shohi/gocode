package signal

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	go func() {
		pid := syscall.Getpid()
		fmt.Printf("pid: %v\n", pid)

		time.Sleep(30 * time.Second)
		syscall.Kill(pid, syscall.SIGINT)
	}()

	log.Println("awaiting signal")
	<-done
	log.Println("exiting")
}
