package go113_test

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

// refer, https://blog.golang.org/go1.13-errors
func TestErrors_Wrap(t *testing.T) {

	errBadStuff := errors.New("something happened")
	err := fmt.Errorf("some context '%s': %w", "parse", errBadStuff)

	log.Printf("%v, unwrapped: %v", err, errors.Unwrap(err))
	// TODO: add errors.As and errors.Is
}
