package error_test

import (
	"log"
	"testing"

	"github.com/goph/emperror"
	"github.com/pkg/errors"
)

func TestEmperror(t *testing.T) {
	err := emperror.With(errors.New("test"), "hello", "world")
	err2 := emperror.With(err, "key2", "value2")
	log.Printf("context: %v", emperror.Context(err2))
}
