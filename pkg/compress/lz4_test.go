package compress

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompression_LZ4(t *testing.T) {
	assert := assert.New(t)

	dataset := []struct {
		Name  string
		Value string
	}{
		{"hello", "world"},
		{"hello", "world"},
	}

	// input := []byte(strings.Repeat("helloworld", 90))
	input, _ := json.Marshal(&dataset)
	algo := Lz4Algo{}
	data, err := algo.Compress(input)
	assert.Nil(err)

	log.Printf("LZ4 =====> %v, %v, %v", len(input), len(data), string(data))

	raw, err := algo.Uncompress(data)
	assert.Nil(err)
	assert.Equal(string(input), string(raw))
}
