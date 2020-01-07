package compress

import (
	"github.com/pierrec/lz4"
)

type Lz4Algo struct{}

func (Lz4Algo) Compress(data []byte) ([]byte, error) {
	buf := make([]byte, len(data))
	ht := make([]int, 64<<10) // buffer for the compression table

	n, err := lz4.CompressBlock(data, buf, ht)

	if err != nil {
		return nil, err
	}

	if n >= len(data) {
		return nil, ErrNotCompressible
	}

	if n == 0 {
		// if input data is too small, LZ4 will not compress it.
		return data, nil
	}

	ret := buf[:n]
	if len(ret) > CompressedMaxSize {
		return nil, ErrCompressedTooLarge

	}

	return ret, nil
}

func (Lz4Algo) Uncompress(data []byte) ([]byte, error) {
	out := make([]byte, CompressedMaxSize)
	n, err := lz4.UncompressBlock(data, out)

	if err == lz4.ErrInvalidSourceShortBuffer {
		return data, nil
	}

	if err != nil {
		return nil, err
	}

	return out[:n], nil
}
