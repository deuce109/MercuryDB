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

func FromWKT(coordinateString string) *Polygon {
	pointList, err := getPoints(coordinateString)
	if err != nil {
		fmt.Printf("Error processing WKT points%s", err)
		return nil
	}

	return &Polygon{
		Coordinates: pointList,
	}
}

func getPoints(pointString string) ([]*Point, error) {

	pointsSplit := strings.Split(strings.TrimPrefix(strings.TrimSuffix(pointString, ")"), "("), ", ")

	points := make([]*Point, len(pointsSplit))

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

		pt := &Point{
			X: x,
			Y: y,
		}

		points = append(points, pt)
	}

	return points, nil

}
