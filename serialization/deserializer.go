package serialization

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io"

	"github.com/deuce109/mercurydb/v2/geometry"
)

func gobDecodeObject(data []byte, result interface{}) error {
	reader := bytes.NewReader(data)
	dec := gob.NewDecoder(reader)
	return dec.Decode(result)
}

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

func DeserializeDataChunk(data []byte, result interface{}) error {

	decompressed, err := decompress(data)

	if err != nil {
		return err
	}
	return gobDecodeObject(decompressed, result)
}

func DeserializeGeometry(data []byte, geom *geometry.Geometry) error {
	return DeserializeDataChunk(data, geom)
}
