package io_test

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr

	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		var buf bytes.Buffer
		wg.Done()
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			// Log the stdout line to my event logger
			buf.Write([]byte("@@" + line + "\n"))
		}
		out <- buf.String()
		reader.Close()
	}()

	wg.Wait()
	f()
	// close the writer to trigger reader side to EOF
	writer.Close()
	return <-out
}

func TestIO_captureSTDOUT(t *testing.T) {
	log.Println("hello")

	re := captureOutput(func() {
		fmt.Println("test fmt 1")
		os.Stderr.WriteString("test stderr 1\n")
		fmt.Println("test fmt 2")
		fmt.Println("test fmt 3")
		log.Println("test log 1")
	})
	fmt.Printf(re)
	fmt.Println("exit")
}
