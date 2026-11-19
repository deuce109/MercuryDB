package geometry

import "log/slog"

type Line struct {
	Coordinates []*Point
}

func (l *Line) GetLength() (float64, error) {

	totalLength := 0.0
	for i := 0; i < len(l.Coordinates)-1; i++ {
		segment, err := l.Coordinates[i].DistanceTo(*l.Coordinates[i+1])
		if err != nil {
			return 0, err
		}
		totalLength += segment
	}
	return totalLength, nil
}

func (l *Line) GetType() GeometryType {
	return LineType
}

func (l *Line) GetCentroid() *Point {
	return get2dCentroidFromPoints(l.Coordinates)
}

func (l *Line) Equals(other Geometry) bool {

	otherLine, err := ToLine(other)
	if err != nil {
		slog.Error("Line equals check failed", "error", err)
		return false
	}

	return allPointsEqual(l.Coordinates, otherLine.Coordinates)
}

func (l *Line) Contains(other Geometry) (bool, error) {
	switch other.GetType() {
	case PointType:
		return l.containsPoint(other)
	case LineType:
		return l.containsLine(other)
	}
	return false, nil
}

func (l *Line) Intersects(other Geometry) (bool, error) {
	switch other.GetType() {
	case PointType:
		return l.containsPoint(other)
	case LineType:
		return l.intersectsLine(other)
	}
	return false, nil
}

func (l *Line) containsPoint(other Geometry) (bool, error) {
	if pt, ok := other.(*Point); ok {
		for i := 0; i < len(l.Coordinates)-1; i++ {
			if LineSegmentContainsPoint(l.Coordinates[i], l.Coordinates[i+1], pt) {
				return true, nil
			}
		}
	}

	return false, nil
}

func (l *Line) containsLine(other Geometry) (bool, error) {

	// iterate over main lines segments
	// check if any line segments of the other line are contained within the main line segments
	// break if any segment is not contained at all

	otherIndexStore := 0

	if otherLine, ok := other.(*Line); ok {
		for mainIndex := 0; mainIndex < len(l.Coordinates)-1; mainIndex++ {
			for otherIndex := otherIndexStore; otherIndex < len(otherLine.Coordinates)-1; otherIndex++ {
				if otherIndexStore != 0 {
					if !LineSegmentContainsPoints(l.Coordinates[mainIndex], l.Coordinates[mainIndex+1], otherLine.Coordinates[otherIndex:otherIndex+1]) {
						return false, nil
					}
				}

				if LineSegmentContainsPoints(l.Coordinates[mainIndex], l.Coordinates[mainIndex+1], otherLine.Coordinates[otherIndex:otherIndex+1]) {
					if otherLine.Coordinates[otherIndex].Equals(l.Coordinates[mainIndex+1]) {
						otherIndexStore = otherIndex
						break
					}
				}
			}
		}
		if otherIndexStore != 0 {
			return true, nil
		}
	}

	return false, nil
}

func (l *Line) intersectsLine(other Geometry) (bool, error) {
	if otherLine, ok := other.(*Line); ok {
		for mainIndex := 0; mainIndex < len(l.Coordinates)-1; mainIndex++ {
			for otherIndex := 0; otherIndex < len(otherLine.Coordinates)-1; otherIndex++ {
				if linesIntersect(l.Coordinates[mainIndex], l.Coordinates[mainIndex+1], otherLine.Coordinates[otherIndex], otherLine.Coordinates[otherIndex+1]) {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func linesIntersect(a1, a2, b1, b2 *Point) bool {
	// Calculate the direction of the lines
	d1 := direction(a1, a2, b1)
	d2 := direction(a1, a2, b2)
	d3 := direction(b1, b2, a1)
	d4 := direction(b1, b2, a2)

	// Check if the lines intersect
	if d1 != d2 && d3 != d4 {
		return true
	}

	return LineSegmentContainsPoints(a1, a2, []*Point{b1, b2}) || LineSegmentContainsPoints(b1, b2, []*Point{a1, a2})
}

func direction(a1, a2, b2 *Point) float64 {
	return (a2.X-a1.X)*(b2.Y-a1.Y) -
		(a2.Y-a1.Y)*(b2.X-a1.X)
}