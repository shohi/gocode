package escape_test

import "testing"

// Escape Analysis

func TestEscape_Frame_2(t *testing.T) {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

	// Pass the "address of" count.
	// NOTE: no escape
	increment(&count)

	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

//go:noinline
func increment(inc *int) {

	// Increment the "value of" count that the "pointer points to". (dereferencing)
	*inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
