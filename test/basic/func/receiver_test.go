package func_test

import (
	"fmt"
	"testing"
)

type MyFunc func() string

func (m MyFunc) String() string {
	if m == nil {
		return "nil"
	}

	return m()
}

func TestFunc_funcReceiver(t *testing.T) {
	var m MyFunc
	fmt.Printf("====> %v\n", m.String())
}
