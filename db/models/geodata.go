package models

import (
	"io"

	"github.com/deuce109/mercurydb/v2/common/geometry"
	"github.com/deuce109/mercurydb/v2/db/serialization"
)

type Metadata = map[string]interface{}

type GeoData struct {
	Geometry *geometry.Geometry
	Metadata Metadata
	Id       string
}

func SerializeGeoDataList(data []*GeoData) (io.Reader, error) {
	return serialization.SerializeToJson(data)
}
