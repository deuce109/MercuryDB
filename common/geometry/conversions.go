package geometry

import (
	"errors"
	"fmt"
)


func ToCircle(geo Geometry) (*Circle, error) {

	if geo.GetType() != CircleType {
		return nil, fmt.Errorf("GeometryType mismatch, expected \"%s\" got \"%s\"", CircleType, geo.GetType())
	}

	if cir, ok := geo.(*Circle); ok {
		return cir, nil
	} else {
		return nil, errors.New("Failed to cast geometry interface to Circle struct")
	}
}

func ToPoint(geo Geometry) (*Point, error) {

if geo.GetType() != PointType {
		return nil, fmt.Errorf("GeometryType mismatch, expected \"%s\" got \"%s\"", PointType, geo.GetType())
	}

	if cir, ok := geo.(*Point); ok {
		return cir, nil
	} else {
		return nil, errors.New("Failed to cast geometry interface to Point struct")
	}
}

func ToLine(geo Geometry) (*Line, error) {
	if geo.GetType() != LineType {
		return nil, fmt.Errorf("GeometryType mismatch, expected \"%s\" got \"%s\"", LineType, geo.GetType())
	}

	if line, ok := geo.(*Line); ok {
		return line, nil
	} else {
		return nil, errors.New("Failed to cast geometry interface to Line struct")
	}
}

func ToPolygon(geo Geometry) (*Polygon, error) {
	if geo.GetType() != PolygonType {
		return nil, fmt.Errorf("GeometryType mismatch, expected \"%s\" got \"%s\"", PolygonType, geo.GetType())
	}

	if polygon, ok := geo.(*Polygon); ok {
		return polygon, nil
	} else {
		return nil, errors.New("Failed to cast geometry interface to Polygon struct")
	}
}