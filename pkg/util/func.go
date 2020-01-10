package util

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
func GetFuncName(f interface{}) string {
	if !isFunc(f) {
		msg := fmt.Sprintf("<not a function: %v>", reflect.TypeOf(f).Kind())
		panic(msg)
	}

	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// https://wycd.net/posts/2014-07-02-logging-function-names-in-go.html
func GetFuncBasename(f interface{}) string {
	fullname := GetFuncName(f)

	dotName := filepath.Ext(fullname)
	fnName := strings.TrimLeft(dotName, ".")

	return fnName
}
