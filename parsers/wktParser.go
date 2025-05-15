package parsers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/deuce109/mercurydb/v2/geometry"
)

const POLY_PREFIX = "POLYGON"
const POINT_PREFIX = "POINT"

func GetSupportedWktTypes() []string {
	return []string{POLY_PREFIX, POINT_PREFIX}
}

func ParseWktCoordinates(coordinateString string) *geometry.Geometry {
	pointList, err := getPoints(coordinateString)
	if err != nil {
		fmt.Printf("Error processing WKT points%s", err)
		return nil
	}

	return &geometry.Geometry{
		Coordinates: pointList,
	}
}

func getPoints(pointString string) ([]geometry.Coordinates, error) {

	pointsSplit := strings.Split(strings.TrimPrefix(strings.TrimSuffix(pointString, ")"), "("), ", ")

	points := make([]geometry.Coordinates, len(pointsSplit))

	for _, p := range pointsSplit {
		coords := strings.Split(strings.TrimPrefix(strings.TrimSuffix(p, ")"), "("), " ")

		x, err := strconv.ParseFloat(coords[0], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(coords[1], 64)
		if err != nil {
			return nil, err
		}

		pt := geometry.Coordinates{
			X: x,
			Y: y,
		}

		points = append(points, pt)
	}

	return points, nil

}
