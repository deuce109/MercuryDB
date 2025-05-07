package serialization

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"
)

func decompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	buf.Write(data)
	gz, err := gzip.NewReader(&buf)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	decompressed, err := io.ReadAll(gz)
	return decompressed, err
}

func Deserialize(data []byte, result interface{}) error {
	decompressed, err := decompress(data)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(decompressed)
	dec := gob.NewDecoder(reader)
	err = dec.Decode(result)

	return err
}
