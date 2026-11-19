package geometry

import (
	"errors"
	"log/slog"
	"math"
)

type Polygon struct {
	Coordinates []*Point
}

func (g *Polygon) GetType() GeometryType {
	return PolygonType
}

func (g *Polygon) Equals(other Geometry) bool {

	otherPolygon, err := ToPolygon(other)
	if err != nil {
		slog.Error("Polygon equals check failed", "error", err)
		return false
	}

	return allPointsEqual(g.Coordinates, otherPolygon.Coordinates)
}

func (g *Polygon) GetCentroid() *Point {
	return get2dCentroidFromPoints(append(g.Coordinates, g.Coordinates[0]))
}

func (g *Polygon) raycast(pt *Point) bool {

	intersections := 0
	numVerticies := len(g.Coordinates)

	x, y := pt.X, pt.Y

	for i := range g.Coordinates {
		x1, y1 := g.Coordinates[i].X, g.Coordinates[i].Y
		x2, y2 := g.Coordinates[(i+1)%numVerticies].X, g.Coordinates[(i+1)%numVerticies].Y

		if (y1 <= y && y < y2) && (x-x1)*(y2-y1) < (x2-x1)*(y-y1) {
			intersections += 1
		}
	}
	return intersections%2 == 1
}

func (g *Polygon) getBoundingBox() (minX, minY, maxX, maxY float64) {
	minX = g.Coordinates[0].X
	minY = g.Coordinates[0].Y
	maxX = g.Coordinates[0].X
	maxY = g.Coordinates[0].Y

	for _, point := range g.Coordinates {

		minX = math.Min(minX, point.X)
		minY = math.Min(minY, point.Y)
		maxX = math.Max(maxX, point.X)
		maxY = math.Max(maxY, point.Y)
	}

	return minX, minY, maxX, maxY
}

func (g *Polygon) containsPoint(pt *Point) bool {
	minX, minY, maxX, maxY := g.getBoundingBox()
	if pt.X < minX || pt.X > maxX || pt.Y < minY || pt.Y > maxY {
		return false
	}

	return g.raycast(pt)
}

func (g *Polygon) containsAllPoints(points []*Point) (bool) {
	for _, pt := range points {
		if !g.containsPoint(pt) {
			return false
		}
	}
	return true
}

func (g *Polygon) containsAnyPoint(points []*Point) (bool) {
	for _, pt := range points {
		if g.containsPoint(pt) {
			return true
		}
	}

	return false
}



func (g *Polygon) Contains(other Geometry) (bool, error) {
	if g == nil || other == nil {
		return false, errors.New("One of the polygons is nil")
	}

	switch other.GetType() {
	case PointType:
		if pt, ok := other.(*Point); ok {
			return g.containsPoint(pt), nil
		}
	case LineType:
		if line, ok := other.(*Line); ok {
			return g.containsAllPoints(line.Coordinates), nil
		}
	case PolygonType:
		if poly, ok := other.(*Polygon); ok {
			return g.containsAllPoints(poly.Coordinates), nil
		}
	
	}
return false, errors.New("Unsupported geometry type for polygon containment")
}

func (g *Polygon) Intersects(other Geometry) (bool, error) {
	if g == nil || other == nil {
		return false, errors.New("One of the polygons is nil")
	}

	switch other.GetType() {
	case PointType:
		if pt, ok := other.(*Point); ok {
			return g.containsPoint(pt), nil
		}
	case LineType:
		if line, ok := other.(*Line); ok {
			return g.containsAnyPoint(line.Coordinates), nil
		}
	case PolygonType:
		if poly, ok := other.(*Polygon); ok {
			return g.containsAnyPoint(poly.Coordinates), nil
		}

	case CircleType:
		if circle, ok := other.(*Circle); ok {
			return circle.Intersects(g)
		}
	}
	return false, errors.New("Unsupported geometry type for polygon intersection")
}
