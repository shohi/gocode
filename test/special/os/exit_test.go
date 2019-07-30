package os_test

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func someExit() {
	log.Printf("hello")
	time.Sleep(1 * time.Second)
	os.Exit(-1)
}

func Test_ExitLog(t *testing.T) {
	assert := assert.New(t)

	if os.Getenv("BE_SOMEEXIT") == "1" {
		someExit()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=Test_ExitLog")
	cmd.Env = append(os.Environ(), "BE_SOMEEXIT=1")
	r, err := cmd.StderrPipe()
	var buf bytes.Buffer
	go func() {
		_, _ = buf.ReadFrom(r)
	}()

	defer r.Close()
	assert.Nil(err)

	err = cmd.Start()
	assert.Nil(err)

	err = cmd.Wait()
	assert.NotNil(err)

	assert.Contains(buf.String(), "hello")

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("process ran with err %v, want exit status nonzero", err)
}
