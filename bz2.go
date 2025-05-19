package bz2

import (
	"bytes"
	"github.com/dsnet/compress/bzip2"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/compression/bz2", new(BZ2))
}

type BZ2 struct{}

func (*BZ2) Decompress(compressed []byte) ([]byte, error) {
	reader := bytes.NewReader(compressed)
	bz2Reader, err := bzip2.NewReader(reader, nil)
	if err != nil {
		return nil, err
	}
	defer bz2Reader.Close()

	var decompressed bytes.Buffer
	_, err = decompressed.ReadFrom(bz2Reader)
	if err != nil {
		return nil, err
	}

	return decompressed.Bytes(), nil
}
