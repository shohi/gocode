package switch_test

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

func TestSwitch_default(t *testing.T) {

	switch {
	case 1 == 2:
		log.Printf("1 < 2")
	case 3 > 3:
		log.Printf("3 == 3")
	default:
		log.Printf("default")
	}
}

func TestSwitch_continue(t *testing.T) {
	val := 1
	bVal := 5
	for k := 0; k < 10; k++ {
		switch k {
		case val:
			// break flow for current `k` and go to next `k`
			continue
		case bVal:
			// NOTE: break flow for current switch, not `k`
			break
		default:
			// do nothing
		}
		log.Printf("round: %v", k)
	}
}
