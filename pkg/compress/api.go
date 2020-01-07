package compress

import "errors"

const CompressedMaxSize = 1 << 10

var (
	ErrNotCompressible    = errors.New("data is not compressible")
	ErrCompressedTooLarge = errors.New("data is too large to be decompressed")
)

type Algo interface {
	Compress([]byte) ([]byte, error)
	Uncompress([]byte) ([]byte, error)
}
