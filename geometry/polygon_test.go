package geometry

import "testing"

var polygon = Polygon{
	Coordinates: []Coordinates{
		{X: 0.0, Y: 0.0},
		{X: 4.0, Y: 0.0},
		{X: 4.0, Y: 3.0},
	},
}

func TestPolygonArea(t *testing.T) {
	result := polygon.Area()
	if result != 6 {
		t.Fatalf("Expected area 6, got %f", result)
	}
}

func TestPolygonPerimeter(t *testing.T) {
	result := polygon.Perimeter()
	if result != 12 {
		t.Fatalf("Expected perimeter 12, got %f", result)
	}
}

func TestPolygonCenter(t *testing.T) {
	result := polygon.Center()
	expected := Coordinates{X: 8.0 / 3.0, Y: 1.0}
	if result != expected {
		t.Fatalf("Expected center (%f, %f), got (%f, %f)", expected.X, expected.Y, result.X, result.Y)
	}
}

func TestPolygonContainsPoint(t *testing.T) {

	pointInside := Coordinates{X: 2, Y: 1}
	pointOutside := Coordinates{X: 5, Y: 1}

	if !polygon.ContainsCoordinates(pointInside) {
		t.Fatalf("Expected polygon to contain point (%f, %f)", pointInside.X, pointInside.Y)
	}
	if polygon.ContainsCoordinates(pointOutside) {
		t.Fatalf("Expected polygon to not contain point (%f, %f)", pointOutside.X, pointOutside.Y)
	}
}

func TestPolygonContainsPointNearEdge(t *testing.T) {
	pointInside := Coordinates{X: 3.999999, Y: 0.999999}
	pointOutside := Coordinates{X: 4.000001, Y: -0.000001}
	if !polygon.ContainsCoordinates(pointInside) {
		t.Errorf("Expected polygon to contain point (%f, %f)", pointInside.X, pointInside.Y)
	}
	if polygon.ContainsCoordinates(pointOutside) {
		t.Errorf("Expected polygon to not contain point (%f, %f)", pointOutside.X, pointOutside.Y)
	}
}

func TestPolygonContainsPolygon(t *testing.T) {
	containedPolygon := Polygon{
		Coordinates: []Coordinates{
			{X: 1.0, Y: 1.0},
			{X: 2.0, Y: 1.0},
			{X: 2.0, Y: 2.0},
			{X: 1.0, Y: 2.0},
		},
	}
	if !polygon.ContainsPolygon(containedPolygon) {
		t.Fatalf("Expected polygon to contain contained polygon")
	}

	if containedPolygon.ContainsPolygon(polygon) {
		t.Fatalf("Expected contained polygon to not contain polygon")
	}
}

func TestPolygonIntersectsPolygon(t *testing.T) {
	intersectingPolygon := Polygon{
		Coordinates: []Coordinates{
			{X: 2.0, Y: 2.0},
			{X: 5.0, Y: 2.0},
			{X: 5.0, Y: 5.0},
			{X: 2.0, Y: 5.0},
		},
	}
	if !polygon.Intersects(intersectingPolygon) {
		t.Fatalf("Expected polygons to intersect")
	}
}

func TestPolygonIntersectsEdge(t *testing.T) {
	intersectingPolygon := Polygon{
		Coordinates: []Coordinates{
			{X: -1, Y: 2},
			{X: 3, Y: 2},
			{X: 3, Y: 6},
		},
	}
	if !intersectingPolygon.Intersects(polygon) {
		t.Fatalf("Expected polygons to intersect")
	}
}

func TestPolygonDoesNotIntersect(t *testing.T) {
	nonIntersectingPolygon := Polygon{
		Coordinates: []Coordinates{
			{X: 5.0, Y: 5.0},
			{X: 6.0, Y: 5.0},
			{X: 6.0, Y: 6.0},
			{X: 5.0, Y: 6.0},
		},
	}
	if polygon.Intersects(nonIntersectingPolygon) {
		t.Fatalf("Expected polygons to not intersect")
	}
}
