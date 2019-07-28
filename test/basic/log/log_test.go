package log_test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLogPrint(t *testing.T) {
	log.Print("hello\naaaa")
	log.Print("world")
}

func TestLogPrintf(t *testing.T) {
	log.Printf("test log: %s", "hello")
	log.Printf("test log: %s", "world")
}

func TestLogFlag(t *testing.T) {
	flags := log.LstdFlags | log.Lmicroseconds
	pre := fmt.Sprintf("[%d] ", os.Getpid())

	logger1 := log.New(os.Stderr, pre, flags)
	logger1.Printf("Hello")

	logger2 := log.New(os.Stderr, pre, 0)
	logger2.Printf("Hello2")
}
