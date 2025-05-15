package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/deuce109/mercurydb/v2/config"
	"github.com/deuce109/mercurydb/v2/directories"
	"github.com/deuce109/mercurydb/v2/geometry"
	"github.com/deuce109/mercurydb/v2/parsers"
)



func filterResults(dr *directories.DataReader, filter func(*geometry.Geometry) bool) []*geometry.Geometry {
	results := make([]*geometry.Geometry, 0)

	for result := range dr.Results {
		if result.Err != nil {
			fmt.Printf("Error decoding object: %v", result.Err)
		} else if filter(result.Geometry) {
			results = append(results, result.Geometry)
		}
	}

	return results
}

func within(poly *geometry.Geometry) []*geometry.Geometry {	

	dr := directories.NewDataReader()
	go dr.ReadDir(config.Config.DataPath, true)

	return filterResults(dr, poly.Contains)

}

func intersect(poly *geometry.Geometry) []*geometry.Geometry {
	
	dr := directories.NewDataReader()
	go dr.ReadDir(config.Config.DataPath, true)

	return filterResults(dr, poly.Intersects)
}

func Get(method string, wkt string, args... []string) ([]*geometry.Geometry, error) {
	geom := parsers.ParseWktCoordinates(wkt)

	switch (strings.ToLower(method)) {
	case "within":
		return within(geom), nil
	case "intersects":
		return intersect(geom), nil
	default:
		return nil, errors.New("get functions must have a method of 'WITHIN' or 'INTERSECTS'")
	}

} 