package compress

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pierrec/lz4"
)

func TestCompress_LZ4(t *testing.T) {
	s := "helloworld"
	data := []byte(strings.Repeat(s, 100))
	buf := make([]byte, len(data))
	ht := make([]int, 64<<10) // buffer for the compression table

	n, err := lz4.CompressBlock(data, buf, ht)

	if err != nil {
		fmt.Println(err)
	}
	if n >= len(data) {
		fmt.Printf("`%s` is not compressible", s)
	}

	fmt.Printf("original: %v, compressed: %v\n", len(data), n)

	buf = buf[:n] // compressed data

	// Allocated a very large buffer for decompression.
	out := make([]byte, len(data))
	n, err = lz4.UncompressBlock(buf, out)
	if err != nil {
		fmt.Println(err)
	}
	out = out[:n] // uncompressed data

	fmt.Println(string(out[:len(s)]))
}
