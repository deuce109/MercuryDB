package commands

import (
	"github.com/deuce109/mercurydb/v2/config"
	"github.com/deuce109/mercurydb/v2/directories"
	"github.com/deuce109/mercurydb/v2/geometry"
	"github.com/deuce109/mercurydb/v2/parsers"
)

func Put(wkt string, data map[string]interface{}) error {
	geom := parsers.ParseWktCoordinates(wkt)
	var existing *geometry.Geometry
	if name, ok := data["name"]; ok {
		if nameAsString, ok := name.(string); ok {
			geom.Name = nameAsString
			delete(data, "name")
		}
	}

	if id, ok := data["id"]; ok {
		if idAsString, ok := id.(string); ok {
			geom.Id = idAsString
			delete(data, "id")
			dr := directories.NewDataReader()
			dr.ReadSingleFile(geom.Id)
			result := <-dr.Results
			existing = result.Geometry
		
		}
	}

	geom.Data = data


	if existing != nil {
		for key, value := range geom.Data {
			existing.Data[key] = value
		}
		return directories.WriteToDirectory(existing, config.Config.DataPath)

	}

	return directories.WriteToDirectory(geom, config.Config.DataPath)
}