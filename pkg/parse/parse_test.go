package parse

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLabels(t *testing.T) {

	cases := []struct {
		name   string
		input  string
		expNum int
	}{
		{"normal", "k1:v1, k2:v2", 2},
		{"only-key", "k1, k2", 2},
		{"has-only-key", "k1:v1, k2, k3:v3", 3},
		{"only-value", ":v1, :v3", 0},
		{"empty", "", 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert := assert.New(t)
			labels := ParseLabels(c.input)

			log.Printf("labels: %v", labels)

			assert.Equal(c.expNum, len(labels))
		})
	}

}
