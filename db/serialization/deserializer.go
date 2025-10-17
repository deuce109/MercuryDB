package serialization

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"io"
	"log/slog"

	"github.com/deuce109/mercurydb/v2/common/geometry"
)

func gobDecodeObject(r io.Reader, result interface{}) error {
	return gob.NewDecoder(r).Decode(result)
}

func decompress(r io.Reader) (io.Reader, error) {

	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	return &buf, err
}

func DeserializeDataChunk(r io.Reader, result interface{}) error {

	decompressed, err := decompress(r)

	if err != nil {
		return err
	}
	return gobDecodeObject(decompressed, result)
}

func DeserializeGeometry(r io.Reader, geom *geometry.Geometry) error {
	return DeserializeDataChunk(r, geom)
}

func DeserializeJsonStringToMap(value string) map[string]interface{} {

	result := make(map[string]interface{})

	err := json.Unmarshal([]byte(value), &result)

	if err != nil {
		slog.Error("Error deserializing JSON string to map", "error", err, "value", value)
		return nil
	}

	return result
}
