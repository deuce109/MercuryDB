package geometry

import "log/slog"

type Circle struct {
	Center *Point
	Radius float64
}

func (c *Circle) GetType() GeometryType {
	return CircleType
}

func (c *Circle) Contains(other Geometry) (bool, error) {
	switch other.GetType() {
	case PointType:
		return c.containsPoint(other)
	case CircleType:
		return c.containsCircle(other)
	case LineType:
		return c.containsLine(other)
	case PolygonType:
		return c.containsPolygon(other)
	}

	return false, nil
}

func (c *Circle) Intersects(other Geometry) (bool, error) {
	switch other.GetType() {
	case PointType:
		return c.containsPoint(other)
	case CircleType:
		return c.intersectsCircle(other)
	case LineType:
		return c.intersectsLine(other)
	case PolygonType:
		return c.intersectsPolygon(other)
	}
	return false, nil
}

func (c *Circle) GetCentroid() *Point {
	return c.Center
}

func (c *Circle) Equals(other Geometry) bool {

	otherCir, err := ToCircle(other)
	if err != nil {
		slog.Error("Equals failed with error", "error", err)
		return false
	}

	return c.Radius == otherCir.Radius && c.Center.Equals(otherCir.Center)
}

func (c *Circle) containsPoint(other Geometry) (bool, error) {
	if pt, ok := other.(*Point); ok {
		distance, err := c.Center.DistanceTo(*pt)
		return distance <= c.Radius, err
	}
	return false, nil
}

func (c *Circle) containsCircle(other Geometry) (bool, error) {
	if otherCircle, ok := other.(*Circle); ok {
		distance, err := c.Center.DistanceTo(*otherCircle.Center)
		radiusDif := c.Radius - otherCircle.Radius
		return distance <= radiusDif, err
	}
	return false, nil
}

func (c *Circle) containsPolygon(other Geometry) (bool, error) {
	if poly, ok := other.(*Polygon); ok {
		return c.containsAllPoints(poly.Coordinates)
	}
	return false, nil
}

func (c *Circle) containsLine(other Geometry) (bool, error) {
	if line, ok := other.(*Line); ok {
		return c.containsAllPoints(line.Coordinates)
	}
	return false, nil
}

func (c *Circle) containsAllPoints(points []*Point) (bool, error) {
	for _, pt := range points {
		containsPt, err := c.containsPoint(pt)
		if err != nil || !containsPt {
			return false, err
		}
	}
	return true, nil
}

func (c *Circle) intersectsCircle(other Geometry) (bool, error) {
	if otherCircle, ok := other.(*Circle); ok {
		distance, err := c.Center.DistanceTo(*otherCircle.Center)
		radiusSum := c.Radius + otherCircle.Radius
		return distance <= radiusSum, err
	}
	return false, nil
}

func (c *Circle) intersectsLine(other Geometry) (bool, error) {
	if line, ok := other.(*Line); ok {
		return c.containsAnyPoint(line.Coordinates)
	}
	return false, nil
}

func (c *Circle) intersectsPolygon(other Geometry) (bool, error) {
	if poly, ok := other.(*Polygon); ok {
		return c.containsAnyPoint(poly.Coordinates)
	}
	return false, nil
}

func (c *Circle) containsAnyPoint(points []*Point) (bool, error) {
	for _, pt := range points {
		containsPt, err := c.containsPoint(pt)
		if containsPt {
			return true, err
		}
	}
	return false, nil
}
