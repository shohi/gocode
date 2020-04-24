package struct_test

import (
	"fmt"
	"testing"
)

type MyConfig struct {
	Val     string
	OnEvent func()
}

func (c *MyConfig) init() {
	if c.OnEvent == nil {
		c.OnEvent = func() {
			fmt.Printf("====> val: %v\n", c.Val)
		}
	}
}

func TestStruct_Closure(t *testing.T) {
	conf := MyConfig{
		Val: "xxx",
	}
	conf.init()

	conf2 := conf
	conf2.Val = "yyy"
	conf2.init()

	conf2.OnEvent()
}
