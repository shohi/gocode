package exec_test

import (
	"log"
	"os/exec"
	"testing"
	"time"

	"github.com/mitchellh/go-ps"
)

func isCommandAvailable(name string) bool {
	cmd := exec.Command("command", "-v", name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func TestExec_Loop(t *testing.T) {
	if !isCommandAvailable("watch") {
		t.Skipf("command `watch` is unavaible,skip test")
	}

	// invoke command
	go func() {
		cmd := exec.Command("watch", "ls")
		err := cmd.Start()
		log.Printf("error: %v", err)
	}()

	time.Sleep(5 * time.Second)
	pl, err := ps.Processes()

	if err != nil {
		log.Printf("processes error: %v", err)
	} else {

		var pNames []string
		for _, p := range pl {
			pNames = append(pNames, p.Executable())

		}
		// NOTE: invoked process does not exit once the goroutine which created it exit
		log.Printf("all processes list: [%v]", pNames)
	}
}
