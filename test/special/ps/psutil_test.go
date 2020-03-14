package ps_test

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/shirou/gopsutil/process"
)

func TestPsUtil(t *testing.T) {
	lst, _ := process.Processes()
	for _, p := range lst {
		cli, err := p.Cmdline()
		if err != nil {
			log.Println(err)
			continue
		}
		if strings.Contains(cli, "ssh") {
			fmt.Printf("process: [%v], pid: %v\n", cli, p.Pid)
		}
	}
}
