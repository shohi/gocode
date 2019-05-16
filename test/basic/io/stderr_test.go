package io_test

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// used to check output of splunk metrics as they are dumped to Stderr.
// Caller MUST close the sigCh to recover Stderr after use.
func captureStdErrToBuffer() (sigCh chan struct{}, buf *bytes.Buffer) {
	sigCh = make(chan struct{}, 1)

	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stderr := os.Stderr
	os.Stderr = writer

	buf = &bytes.Buffer{}
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()
			buf.Write([]byte(line))
		}
	}()

	// recover Stderr
	go func() {
		<-sigCh
		os.Stderr = stderr
		writer.Close()
	}()

	return
}

func TestCaptureStdErrToBuffer(t *testing.T) {
	assert := assert.New(t)

	sigCh, buf := captureStdErrToBuffer()

	l := log.New(os.Stderr, "", 0)
	l.Println("msg")

	// wait for log collected
	time.Sleep(50 * time.Millisecond)

	assert.Equal(buf.String(), "msg")

	// release Stderr capture
	close(sigCh)

	time.Sleep(50 * time.Millisecond)
	l.Println("hello")

	assert.NotContains(buf.String(), "hello")
}
