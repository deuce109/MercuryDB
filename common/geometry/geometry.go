package geometry

import "math"

type Geometry struct {
	Coordinates []Coordinates
	Name        string
}

func (g *Geometry) IsPoint() bool {
	return len(g.Coordinates) > 1
}

func (g *Geometry) IsPolygon() bool {
	return len(g.Coordinates) > 3
}

func (g *Geometry) IsLine() bool {
	return len(g.Coordinates) == 2
}

func (g *Geometry) raycast(pt Coordinates) bool {

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

func (g *Geometry) getBoundingBox() (minX, minY, maxX, maxY float64) {
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

func (g *Geometry) containsCoordinates(pt Coordinates) bool {
	minX, minY, maxX, maxY := g.getBoundingBox()
	if pt.X < minX || pt.X > maxX || pt.Y < minY || pt.Y > maxY {
		return false
	}

	return g.raycast(pt)
}

func (g *Geometry) Contains(other *Geometry) bool {

	if g == nil || other == nil {
		return false
	}

	for _, pt := range other.Coordinates {
		if !g.containsCoordinates(pt) {
			return false
		}
	}
	return true
}

func (g *Geometry) Intersects(other *Geometry) bool {
	if g == nil || other == nil {
		return false
	}
	for _, pt := range other.Coordinates {
		if g.containsCoordinates(pt) {
			return true
		}
	}
	for _, pt := range g.Coordinates {
		if other.containsCoordinates(pt) {
			return true
		}
	}
	return false
}

func (g *Geometry) Center() Coordinates {
	if g.IsPoint() {
		return g.Coordinates[0]
	}
	var x, y float64
	for _, pt := range g.Coordinates {
		x += pt.X
		y += pt.Y
	}
	return Coordinates{
		X: x / float64(len(g.Coordinates)),
		Y: y / float64(len(g.Coordinates)),
	}
}

func (g *Geometry) Area() float64 {
	if g.IsPoint() {
		return 0.0
	}
	area := 0.0
	var nextPoint Coordinates
	for i, pt := range g.Coordinates {

		switch i {
		case len(g.Coordinates) - 1:
			nextPoint = g.Coordinates[0]
		default:
			nextPoint = g.Coordinates[i+1]
		}

		area += (nextPoint.X * pt.Y) - (pt.X * nextPoint.Y)
	}
	return math.Abs(area / 2.0)
}

func (g *Geometry) Perimeter() float64 {

	if g.IsPoint() {
		return 0.0
	}

	perimiter := 0.0
	var nextPoint Coordinates
	for i, pt := range g.Coordinates {

		switch i {
		case len(g.Coordinates) - 1:
			nextPoint = g.Coordinates[0]
		default:
			nextPoint = g.Coordinates[i+1]
		}

		perimiter += nextPoint.DistanceTo(pt)

	}

	return perimiter
}
