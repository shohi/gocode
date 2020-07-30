package args_test

import (
	"log"
	"testing"
	"time"
)

type options struct {
	manualAck bool
	timeout   time.Duration
}

type Option func(opts *options) error

func WithTimeout(d time.Duration) Option {
	
	return func(opts *options) error {
		opts.timeout = d
		return nil
	}
}

func runWithOptions(option ...Option) {
	opts := &options{}

	for _, o := range option {
		o(opts)
	}

	log.Printf("options: %v", opts)
}

func TestArgs_VariableArguments(t *testing.T) {
	// without options
	runWithOptions()

	// with options
	runWithOptions(WithTimeout(10 * time.Millisecond))
}
