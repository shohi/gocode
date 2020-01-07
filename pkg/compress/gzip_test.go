package compress

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompression_GZIP(t *testing.T) {
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
	algo := GzipAlgo{}
	data, err := algo.Compress(input)
	assert.Nil(err)

	log.Printf("GZIP =====> %v, %v, %v", len(input), len(data), string(data))

	raw, err := algo.Uncompress(data)
	assert.Nil(err)
	assert.Equal(string(input), string(raw))
}

func TestCompression_GZIP_2(t *testing.T) {
	var ll = []struct {
		Name string
		Id   int
	}{
		{"Hello", 1},
		{"World", 2},
	}

	b, err := json.Marshal(ll)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)

	_, err = gz.Write(b)
	if err != nil {
		panic(err)
	}

	err = gz.Close()
	if err != nil {
		panic(err)
	}

	out := buf.Bytes()

	//Read, without same buffer

	r, err := gzip.NewReader(ioutil.NopCloser(bytes.NewBuffer(out)))
	if err != nil {
		panic(err)
	}
	defer r.Close()
	bb, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bb))
}
