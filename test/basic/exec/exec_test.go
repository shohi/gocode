package exec_test

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

var pgrep = func(processName string) ([]byte, error) {
	return exec.Command("pgrep", processName).Output()
}

func TestPgrep(t *testing.T) {
	procName := "gopls"
	output, err := pgrep(procName)

	if err != nil {
		log.Printf("content: [%v], err: %v", string(output), err)
		return
	}

	var myPid = os.Getpid()
	var pidStrs = strings.Split(string(output), "\n")
	var pids = make([]int, 0, len(pidStrs))

	for _, pidStr := range pidStrs {
		if pidStr == "" {
			continue
		}
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			log.Printf("unable to resolve pid, err: %v", err)
			return
		}
		// Ignore the current process.
		if pid == myPid {
			continue
		}
		pids = append(pids, pid)
	}

	for index, pid := range pids {
		log.Printf("index: %d, pid: %v\n", index, pid)
	}
}
