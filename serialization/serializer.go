package serialization

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"

	"github.com/deuce109/mercurydb/v2/geometry"
)

func gobEncodeObject(obj interface{}) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(obj)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func compress(data []byte) ([]byte, int, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write(data)
	if err != nil {
		return nil, -1, err
	}
	err = gz.Close()
	if err != nil {
		return nil, -1, err
	}

	compressed := buf.Bytes()
	return compressed, len(compressed), nil
}

func SerializeDataChunk(data interface{}) ([]byte, int, error) {
	gobbedData, err := gobEncodeObject(data)
	if err != nil {
		return nil, -1, err
	}
	return compress(gobbedData)
}

func SerializeGeometry(geo geometry.Geometry) ([]byte, int, error) {
	return SerializeDataChunk(geo)
}
