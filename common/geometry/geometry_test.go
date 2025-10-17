package geometry

import "testing"

func TestDistance(t *testing.T) {
	p1 := Coordinates{X: 0, Y: 0}
	p2 := Coordinates{X: 3, Y: 4}
	expected := 5.0
	result := p1.DistanceTo(p2)
	if result != expected {
		t.Fatalf("Expected distance %f, got %f", expected, result)
	}
}
