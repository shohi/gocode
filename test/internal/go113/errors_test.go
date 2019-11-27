package go113_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/pkg/errors"
)

// refer, https://blog.golang.org/go1.13-errors
func TestErrors_Wrap(t *testing.T) {

	errBadStuff := errors.New("something happened")
	err := fmt.Errorf("some context '%s': %w", "parse", errBadStuff)

	log.Printf("%v", err)
}
