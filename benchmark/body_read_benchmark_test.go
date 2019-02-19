package benchmark

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var probBuf [1]byte

func readWithIOUtil(resp *http.Response) []byte {
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}

	return data
}

func readWithBodyRead(resp *http.Response) []byte {
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.Body == nil || resp.ContentLength <= 0 {
		return nil
	}

	data := make([]byte, resp.ContentLength)

	_, err := io.ReadFull(resp.Body, data)
	if err != nil {
		return nil
	}

	if err == nil {
		_, err = io.ReadFull(resp.Body, probBuf[0:])
		if err == nil {
			return nil
		}
	}

	return data
}

func BenchmarkRead(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(resp *http.Response) []byte
	}{
		{"ReadWithIOUtil", readWithIOUtil},
		{"ReadWithBody", readWithBodyRead},
	}

	var resp http.Response

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			// r := rand.New(rand.NewSource(1))
			// size := r.Int63n(1024*1024*2) + 1
			var size int64 = 1024 * 1024 * 2
			for k := 0; k < b.N; k++ {
				resp.Body = ioutil.NopCloser(bytes.NewReader(make([]byte, size)))
				resp.ContentLength = size

				bm.fn(&resp)
			}
		})
	}
}

func TestReadBody(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		fn   func(resp *http.Response) []byte
	}{
		{"ReadWithIOUtil", readWithIOUtil},
		{"ReadWithBody", readWithBodyRead},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var resp http.Response
			const size = 1024 * 1024 * 2
			resp.Body = ioutil.NopCloser(bytes.NewReader(make([]byte, size)))
			resp.ContentLength = size

			data := test.fn(&resp)
			assert.NotNil(data)
		})
	}
}
