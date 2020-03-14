package ps_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mitchellh/go-ps"
)

func TestProcess(t *testing.T) {
	lst, _ := ps.Processes()

	for _, p := range lst {
		if strings.Contains(p.Executable(), "ssh") {
			fmt.Printf("process: [%v], pid: %v\n", p.Executable(), p.Pid())
		}
	}
}
