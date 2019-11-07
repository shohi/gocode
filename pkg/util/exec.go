package util

import (
	"bytes"
	"os/exec"

	"github.com/atrox/homedir"
)

func expandDir(dir string) (string, error) {
	return homedir.Expand(dir)
}

func expandDirs(dirs []string) ([]string, error) {
	var ret = make([]string, len(dirs))

	for k, v := range dirs {
		val, err := expandDir(v)
		if err != nil {
			return nil, err
		}

		ret[k] = val
	}

	return ret, nil

}

// RunShellCommand runs shell command and get its output into byte.Buffer.
// refer, https://jamiethompson.me/posts/Unit-Testing-Exec-Command-In-Golang/
func RunShellCommand(name string, args ...string) (*bytes.Buffer, error) {
	cmdName, err := expandDir(name)
	if err != nil {
		return nil, err
	}
	cmdArgs, err := expandDirs(args)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(cmdName, cmdArgs...)

	// Set up byte buffers to read stdout
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return &outb, nil
}
