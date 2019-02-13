package benchmark

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

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
	resp.Body.Read(data)

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
			// r = rand.New(rand.NewSource(99))
			for k := 0; k < b.N; k++ {
				const size = 1024 * 1024 * 2
				resp.Body = ioutil.NopCloser(bytes.NewReader(make([]byte, size)))
				resp.ContentLength = size

				bm.fn(&resp)
			}
		})
	}

}
