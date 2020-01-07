package compress

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
)

type GzipAlgo struct{}

func (GzipAlgo) Compress(data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	w := gzip.NewWriter(buf)

	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	// NOTE: Not known why `w.Close` can't be defered.
	// Otherwise, gzip will not work.
	if err := w.Close(); err != nil {
		return nil, err
	}

	ret := buf.Bytes()
	if len(ret) > CompressedMaxSize {
		return nil, ErrCompressedTooLarge
	}

	return ret, nil
}

func (GzipAlgo) Uncompress(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	r, err := gzip.NewReader(ioutil.NopCloser(buf))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, r)

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
