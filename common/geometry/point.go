package geometry

import (
	"fmt"
	"log/slog"
	"math"
)

type Point struct {
	X float64
	Y float64
	Z float64
}

func (p *Point) Equals(other Geometry) bool {

	pt, err := ToPoint(other)

	if err != nil {
		slog.Error("Point equals check failed", "error", err)
		return false
	}

	return p.X == pt.X && p.Y == pt.Y && p.Z == pt.Z

}

func (p *Point) GetCentroid() *Point {
	return p
}

func (p *Point) GetType() GeometryType {
	return PointType
}

func (p *Point) Contains(other Geometry) (bool, error) {
	switch other.GetType() {
	case PointType:
		if pt, ok := other.(*Point); ok {
			return p.Equals(pt), nil
		}
	}
	return false, nil
}

func (p *Point) Intersects(other Geometry) (bool, error) {
	return p.Contains(other)
}

func (p *Point) DistanceTo(other Point) (float64, error) {

	if p == nil || &other == nil {
		return 0.0, fmt.Errorf("One set of coordinates are nil")
	}

	return math.Sqrt(math.Pow(other.X-p.X, 2) + math.Pow(other.Y-p.Y, 2) + math.Pow(other.Z-p.Z, 2)), nil
}


func allPointsEqual(mainPoints, otherPoints []*Point) bool {
	if len(mainPoints) != len(otherPoints) {
		return false
	}

	for i := 0; i < len(mainPoints); i++ {
		if !mainPoints[i].Equals(otherPoints[i]) {
			return false
		}
	}

	return true
}