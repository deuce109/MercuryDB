package io

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"slices"

	"github.com/deuce109/mercurydb/v2/common/geometry"
	"github.com/deuce109/mercurydb/v2/db/models"
	"github.com/deuce109/mercurydb/v2/db/serialization"
)

const MERCURY_HEADER = "MERCURY"
const UUID_SIZE = 8

const type_error_format = "expected *bytes.Buffer but got %T"

type MercuryFile struct {
	Path string
	file *os.File
}

func NewMercuryFile(uuid string) *MercuryFile {

	absPath, err := filepath.Abs(os.Getenv("MERCURY_DATA_PATH"))

	if err != nil {
		return nil
	}

	path := fmt.Sprintf("%s/%s.mercury", absPath, uuid)

	return &MercuryFile{
		Path: path,
	}
}

func (f *MercuryFile) Open() error {
	file, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	f.file = file
	return nil
}

func (f *MercuryFile) Close() error {
	if f.file == nil {
		return nil
	}
	err := f.file.Close()
	f.file = nil
	return err
}

func (f *MercuryFile) GeometryContains(geo *geometry.Geometry) bool {
	fileGeo, err := f.readGeometry()

	if err != nil {
		slog.Error("Error reading geometry from file", "file", f.Path, "error", err)
	}

	return fileGeo.Contains(geo)
}

func (f *MercuryFile) GeometryIntersects(geo *geometry.Geometry) bool {
	fileGeo, err := f.readGeometry()

	if err != nil {
		slog.Error("Error reading geometry from file", "file", f.Path, "error", err)
	}

	return fileGeo.Intersects(geo)
}

func (f *MercuryFile) writeHeader() error {
	_, err := f.file.Write([]byte(MERCURY_HEADER))
	return err
}

func (f *MercuryFile) writeId(id string) error {
	_, err := f.file.Write([]byte(id))
	return err
}

func (f *MercuryFile) writeFrom(r io.Reader) error {

	buf, ok := r.(*bytes.Buffer)

	if !ok {
		return fmt.Errorf(type_error_format, r)
	}

	intBuf := make([]byte, 8)

	binary.BigEndian.PutUint64(intBuf, uint64(buf.Len()))

	_, err := f.file.Write(intBuf)

	if err != nil {
		return err
	}

	_, err = io.Copy(f.file, buf)
	return err
}

func (f *MercuryFile) writeMetadata(dataChunk models.Metadata) error {
	data, err := serialization.SerializeDataChunk(dataChunk)
	if err != nil {
		return err
	}

	return f.writeFrom(data)
}

func (f *MercuryFile) writeGeometry(geom *geometry.Geometry) error {
	data, err := serialization.SerializeGeometry(geom)
	if err != nil {
		return err
	}

	return f.writeFrom(data)
}

func (f *MercuryFile) verifyHeader() bool {
	headerBuffer, err := f.seekThenReadData(0, 7)

	if err != nil {
		slog.Error("Error verifying MercuryFile", "error", err, "path", f.Path)
	}

	buf, ok := headerBuffer.(*bytes.Buffer)

	if !ok {
		slog.Error("Error verifying MercuryFile: expected *bytes.Buffer", "type", fmt.Sprintf("%T", headerBuffer), "path", f.Path)
		return false
	}

	return err != nil && slices.Equal(buf.Bytes(), []byte(MERCURY_HEADER))

}

func (f *MercuryFile) readId() (string, error) {
	idBuffer, err := f.seekThenReadData(7, UUID_SIZE)

	buf, ok := idBuffer.(*bytes.Buffer)
	if !ok {
		return "", fmt.Errorf(type_error_format, idBuffer)
	}

	return buf.String(), err
}

func (f *MercuryFile) seekThenReadData(start int64, size int) (io.Reader, error) {
	_, err := f.file.Seek(start, 0)
	if err != nil {
		return nil, err
	}

	return f.readData(uint64(size))

}

func (f *MercuryFile) readData(size uint64) (io.Reader, error) {
	var buf bytes.Buffer
	buf.Grow(int(size))

	_, err := io.CopyN(&buf, f.file, int64(size))

	return &buf, err
}

func (f *MercuryFile) readUInt64() (uint64, error) {
	dataReader, err := f.readData(8)

	buf, ok := dataReader.(*bytes.Buffer)
	if !ok {
		return 0, fmt.Errorf(type_error_format, dataReader)
	}

	return binary.BigEndian.Uint64(buf.Bytes()), err
}

func (f *MercuryFile) readMetadata(names []string) (models.Metadata, error) {
	size, err := f.readUInt64()

	if err != nil {
		return nil, err
	}

	data, err := f.readData(size)

	if err != nil {
		return nil, err
	}

	var result models.Metadata

	err = serialization.DeserializeDataChunk(data, result)

	var filtered models.Metadata = make(map[string]interface{})

	if len(names) > 0 {
		for _, name := range names {
			if value, ok := result[name]; ok {
				filtered[name] = value
			}
		}
		return filtered, err
	}

	return result, err
}

func (f *MercuryFile) readGeometry() (*geometry.Geometry, error) {
	size, err := f.readUInt64()

	if err != nil {
		return nil, err
	}

	data, err := f.readData(size)

	if err != nil {
		return nil, err
	}

	var result *geometry.Geometry

	err = serialization.DeserializeGeometry(data, result)
	return result, err
}

func (f *MercuryFile) GetGeometry() *geometry.Geometry {
	geo, err := f.readGeometry()

	if err != nil {
		slog.Error("Error reading geometry", "path", f.Path, "error", err)
		return nil
	}

	return geo
}

func (f *MercuryFile) ReadGeoData(metadataNames []string) (*models.GeoData, error) {

	valid := f.verifyHeader()
	if !valid {
		return nil, fmt.Errorf("file %s is not a valid Mercury file", f.Path)
	}

	id, err := f.readId()
	if err != nil {
		return nil, err
	}

	geom, err := f.readGeometry()

	if err != nil {
		return nil, err
	}

	metadata, err := f.readMetadata(metadataNames)

	return &models.GeoData{
		Id:       id,
		Metadata: metadata,
		Geometry: geom,
	}, err

}

func (f *MercuryFile) WriteGeoData(geoData *models.GeoData) error {
	err := f.writeHeader()

	if err != nil {
		return err
	}

	err = f.writeId(geoData.Id)

	if err != nil {
		return err
	}

	err = f.writeGeometry(geoData.Geometry)

	if err != nil {
		return err
	}

	err = f.writeMetadata(geoData.Metadata)

	return err
}
