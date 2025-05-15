package geometry

import (
	"math"
)

func (p *Geometry) Center() Coordinates {
	var x, y float64
	for _, pt := range p.Coordinates {
		x += pt.X
		y += pt.Y
	}
	return Coordinates{
		X: x / float64(len(p.Coordinates)),
		Y: y / float64(len(p.Coordinates)),
	}
}

func (p *Geometry) Area() float64 {
	area := 0.0
	var nextPoint Coordinates
	for i, pt := range p.Coordinates {

		switch i {
		case len(p.Coordinates) - 1:
			nextPoint = p.Coordinates[0]
		default:
			nextPoint = p.Coordinates[i+1]
		}

		area += (nextPoint.X * pt.Y) - (pt.X * nextPoint.Y)
	}
	return math.Abs(area / 2.0)
}

func (p *Geometry) Perimeter() float64 {
	perimiter := 0.0
	var nextPoint Coordinates
	for i, pt := range p.Coordinates {

		switch i {
		case len(p.Coordinates) - 1:
			nextPoint = p.Coordinates[0]
		default:
			nextPoint = p.Coordinates[i+1]
		}

		perimiter += nextPoint.DistanceTo(pt)

	}

	return perimiter
}

func (p *Geometry) getBoundingBox() (minX, minY, maxX, maxY float64) {
	minX = p.Coordinates[0].X
	minY = p.Coordinates[0].Y
	maxX = p.Coordinates[0].X
	maxY = p.Coordinates[0].Y

	for _, point := range p.Coordinates {

		minX = math.Min(minX, point.X)
		minY = math.Min(minY, point.Y)
		maxX = math.Max(maxX, point.X)
		maxY = math.Max(maxY, point.Y)
	}

	return minX, minY, maxX, maxY
}

func (p *Geometry) raycast(pt Coordinates) bool {

	intersections := 0
	numVerticies := len(p.Coordinates)

	x, y := pt.X, pt.Y

	for i := range p.Coordinates {
		x1, y1 := p.Coordinates[i].X, p.Coordinates[i].Y
		x2, y2 := p.Coordinates[(i+1)%numVerticies].X, p.Coordinates[(i+1)%numVerticies].Y

		if (y1 <= y && y < y2) && (x-x1)*(y2-y1) < (x2-x1)*(y-y1) {
			intersections += 1
		}
	}
	return intersections%2 == 1
}

func (p *Geometry) containsCoordinates(pt Coordinates) bool {
	minX, minY, maxX, maxY := p.getBoundingBox()
	if pt.X < minX || pt.X > maxX || pt.Y < minY || pt.Y > maxY {
		return false
	}

	return p.raycast(pt)
}

func (p *Geometry) Intersects(other *Geometry) bool {
	for _, pt := range other.Coordinates {
		if p.containsCoordinates(pt) {
			return true
		}
	}

	for _, pt := range p.Coordinates {
		if other.containsCoordinates(pt) {
			return true
		}
	}
	return false
}

func (p *Geometry) containsPolygon(other *Geometry) bool {
	for _, pt := range other.Coordinates {
		if !p.containsCoordinates(pt) {
			return false
		}
	}
	return true
}

func (p *Geometry) Contains(other *Geometry) bool {
	if len(other.Coordinates) == 1 {
		return p.containsCoordinates(other.Coordinates[0])
	} else {
		return p.containsPolygon(other)
	}
}
