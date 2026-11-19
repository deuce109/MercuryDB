package geometry

import (
	"math"
)

const epsilon = 1e-9

func get2dCentroidFromPoints(points []*Point) *Point {
	totalLength := 0.0
	centroidX := 0.0
	centroidY := 0.0

	centerPoints := make([]*Point, len(points)-1)

	for i := 0; i < len(points)-1; i++ {
		segmentLength, err := points[i].DistanceTo(*points[i+1])
		if err != nil {
			return nil
		}
		totalLength += segmentLength

		x := (points[i].X + points[i+1].X) / 2
		y := (points[i].Y + points[i+1].Y) / 2
		centerPoints[i] = &Point{X: x, Y: y}
		centroidX += x * segmentLength
		centroidY += y * segmentLength
	}
	if totalLength > 0 {
		return &Point{X: centroidX / totalLength, Y: centroidY / totalLength}
	}
	return nil
}

func LineSegmentContainsPoint(a, b, p *Point) bool {
	cross := (p.Y-a.Y)*(b.X-a.X) - (p.X-a.X)*(b.Y-a.Y)
	if math.Abs(cross) > epsilon {
		return false
	}

	// 2. Check if p is within the bounding box of segment ab
	minX := math.Min(a.X, b.X) - epsilon
	maxX := math.Max(a.X, b.X) + epsilon
	minY := math.Min(a.Y, b.Y) - epsilon
	maxY := math.Max(a.Y, b.Y) + epsilon

	return p.X >= minX && p.X <= maxX && p.Y >= minY && p.Y <= maxY
}

func LineSegmentContainsPoints(a, b *Point, points []*Point) bool {
	for _, p := range points {
		if !LineSegmentContainsPoint(a, b, p) {
			return false
		}
	}
	return true
}