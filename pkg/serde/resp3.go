package serde

import (
	"bytes"

	"github.com/smallnest/resp3"
)

func SerializeRawRESP3(cmdName string, cmdArgs ...string) []byte {
	var buf bytes.Buffer
	w := resp3.NewWriter(&buf)
	var args = []string{cmdName}
	args = append(args, cmdArgs...)

	w.WriteCommand(args...)
	w.Flush()
	return buf.Bytes()
}

func SerializeRESP3(v *resp3.Value) []byte {
	return []byte(v.ToRESP3String())
}

func DeserializeRESP3(data []byte) (*resp3.Value, error) {
	buf := bytes.NewReader(data)

	r := resp3.NewReader(buf)
	cmd, _, err := r.ReadValue()

	return cmd, err
}
