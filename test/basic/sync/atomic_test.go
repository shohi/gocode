package sync_test

import (
	"fmt"
	"sync/atomic"
	"testing"
)

type MyStruct struct{}

func TestAtomic_Value(t *testing.T) {
	var v atomic.Value
	var a *MyStruct = nil
	v.Store(a)

	aa := v.Load().(*MyStruct)

	fmt.Println(aa)
}

func TestAtomic_Int(t *testing.T) {

}
