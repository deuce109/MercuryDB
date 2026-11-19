package serialization

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strings"

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

func DeserializeGeometry(r io.Reader) (geometry.Geometry, error) {

	var geom geometry.Geometry
	var geomString string

	geomTypeBytes := make([]byte, geometry.MinTypeStringLength)

	_, err := r.Read(geomTypeBytes)

	if err != nil {
		return nil, err
	}

	geomString = string(geomTypeBytes)


	for i := geometry.MinTypeStringLength; i <= geometry.MaxTypeStringLength; i ++ {
		

		switch geomString {
		case string(geometry.LineType):
			geom = &geometry.Line{}
		case string(geometry.PointType):
			geom = &geometry.Point{}
		case string(geometry.CircleType):
			geom = &geometry.Circle{}
		case string(geometry.PolygonType):
			geom = &geometry.Polygon{}
		}

		if geom != nil {
			break
		}

		extraByte := make([]byte, 1)

		r.Read(extraByte)

		extraString := string(extraByte)

		geomString += extraString
	}

	if geom == nil {

		errMsg := fmt.Sprintf("Unexpected geometry type. Found: %s", geomString)
		errMsg = strings.ReplaceAll(errMsg, "\x00", "")

		return nil, fmt.Errorf("%s", errMsg)
	}

	err = DeserializeDataChunk(r, geom)

	return geom, err
}

func DesierializeJsonReaderToMap(r io.Reader) (map[string]interface{}, error) {
	var buffer bytes.Buffer

	_, err := io.Copy(&buffer, r)
	if err != nil {
		return nil, err
	}

	jsonString := buffer.String()

	return DeserializeJsonStringToMap(jsonString), nil
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
