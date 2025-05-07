package geometry

import "testing"

func TestPointFromCoords(t *testing.T) {
	pt := pointFromCoords(0, 0)

	if pt.Coordinates[0].X != 0 || pt.Coordinates[0].Y != 0 {
		t.Fatalf("Expected coordinates to be 0, 0 got %f, %f", pt.Coordinates[0].X, pt.Coordinates[0].Y)
	}
}
