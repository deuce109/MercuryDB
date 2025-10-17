package serialization

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"io"

	"github.com/deuce109/mercurydb/v2/common/geometry"
)

func gobEncodeObject(obj interface{}) (io.Reader, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(obj)

	return &buf, err

}

func compress(data io.Reader) (io.Reader, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := io.Copy(gz, data)
	if err != nil {
		return nil, err
	}
	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func SerializeDataChunk(data interface{}) (io.Reader, error) {
	gobbedData, err := gobEncodeObject(data)
	if err != nil {
		return nil, err
	}
	return compress(gobbedData)
}

func SerializeGeometry(geo *geometry.Geometry) (io.Reader, error) {
	return SerializeDataChunk(geo)
}

func SerializeToJson(data interface{}) (io.Reader, error) {

	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)

	err := enc.Encode(data)

	if err != nil {
		return nil, err
	}
	return &buf, nil
}
