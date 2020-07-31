package flaw_test

import "testing"

func BenchmarkAssignmentIndirect(b *testing.B) {
	type X struct {
		p *int
	}
	for i := 0; i < b.N; i++ {
		var i1 int
		x1 := &X{
			p: &i1, // GOOD: i1 does not escape
		}
		_ = x1

		var i2 int
		x2 := &X{}
		x2.p = &i2 // BAD: Cause of i2 escape
	}
}

func BenchmarkLiteralFunctions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var y1 int
		foo(&y1, 42) // GOOD: y1 does not escape

		var y2 int
		func(p *int, x int) {
			*p = x
		}(&y2, 42) // BAD: Cause of y2 escape

		var y3 int
		p := foo
		p(&y3, 42) // BAD: Cause of y3 escape
	}
}

func foo(p *int, x int) {
	*p = x
}
