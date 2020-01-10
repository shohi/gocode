package util

import (
	"fmt"
	"reflect"
	"testing"
)

func emptyFunction() {}

func TestGetFuncName_UseRuntime(t *testing.T) {
	hello := func() {}
	fmt.Println(GetFuncName(hello))
	// output: github.com/shohi/gocode/pkg/util.TestGetFuncName.func1

	fmt.Println(GetFuncBasename(hello))
	// output: func1

	fmt.Println(GetFuncName(emptyFunction))
	// output: github.com/shohi/gocode/pkg/util.emptyFunction

	fmt.Println(GetFuncBasename(emptyFunction))
	// output: emptyFunction
}

// NOTE: can't get function name from its type Name method.
func TestGetFuncName_UseType(t *testing.T) {
	vType := reflect.TypeOf(emptyFunction)
	fmt.Println(vType.Name())
	// Output: <nothing>
}
