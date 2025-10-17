package geometry

import (
	"fmt"
	"strconv"
	"strings"
)

const POLY_PREFIX = "POLYGON"
const POINT_PREFIX = "POINT"

func GetSupportedWktTypes() []string {
	return []string{POLY_PREFIX, POINT_PREFIX}
}

func FromWKT(coordinateString string) *Geometry {
	pointList, err := getPoints(coordinateString)
	if err != nil {
		fmt.Printf("Error processing WKT points%s", err)
		return nil
	}

	return &Geometry{
		Coordinates: pointList,
	}
}

func getPoints(pointString string) ([]Coordinates, error) {

	pointsSplit := strings.Split(strings.TrimPrefix(strings.TrimSuffix(pointString, ")"), "("), ", ")

	points := make([]Coordinates, len(pointsSplit))

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

		pt := Coordinates{
			X: x,
			Y: y,
		}

		points = append(points, pt)
	}

	return points, nil

}
