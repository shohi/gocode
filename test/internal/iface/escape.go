package main

type Addifier interface{ Add(a, b int32) int32 }

type Adder struct{ name string }

//go:noinline
func (adder Adder) Add(a, b int32) int32 { return a + b }

// GOOS=linux GOARCH=amd64 go tool compile -m escape.go
func main() {
	adder := Adder{name: "myAdder"}
	adder.Add(10, 32)           // doesn't escape
	Addifier(adder).Add(10, 32) // escapes
}
