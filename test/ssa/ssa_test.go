package ssa_test

import "testing"

type IData interface {
	Set(int)
}

type Data struct {
	V int
}

//go:noinline
func (d *Data) Set(v int) {
	d.V = v
}

//go:noinline
func setSSA(d *Data, v int) {
	d.Set(v)
}

//go:noinline
func TestData(t *testing.T) {
	var d Data
	setSSA(&d, 10)
}
