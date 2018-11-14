package flow_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitch_break(t *testing.T) {
	assert := assert.New(t)
	c := 'a'
	output := ""
	switch c {
	case 'a':
		output = "first"
	case 'b':
		output = "second"
	default:
		output = "default"
	}

	// break is implicitly added
	assert.Equal(output, "first")
	log.Printf("output: %v", output)
}

// fallthrough - will skip next case test and directly execute next case block
func TestSwitch_fallthrough(t *testing.T) {
	assert := assert.New(t)
	c := 'b'
	output := ""
	switch c {
	case 'a':
		output = "first"
	case 'b':
		output = "second"
		fallthrough
	case 'c':
		output += " fallthrough"
		fallthrough
	default:
		output += " and default"
	}

	assert.Equal(output, "second fallthrough and default")
	log.Printf("output: %v", output)
}
