package type_test

import (
	"fmt"
	"testing"
)

func TestType_Switch(t *testing.T) {
	do := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Twice %v is %v\n", v, v*2)
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}

	do(21)
	do("hello")
	do(true)
}

type IGet interface {
	Get() string
}

type ISet interface {
	Set(v string)
}

type MyStruct struct {
	V string
}

func (m *MyStruct) Get() string {
	return m.V
}

func (m *MyStruct) Set(v string) {
	m.V = v
}

func TestType_Interface_Switch(t *testing.T) {
	var m interface{} = &MyStruct{V: "hello"}

	switch m.(type) {
	case IGet:
		v := m.(IGet).Get()
		fmt.Printf("get: %v\n", v)
	case ISet:
		m.(ISet).Set("world")
		fmt.Printf("set: %v\n", m)
	default:
		fmt.Printf("not ISet/IGet: %v\n", m)
	}

}
