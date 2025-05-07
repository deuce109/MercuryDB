package geometry

type Point struct {
	Coordinates []Coordinates
}

func pointFromCoords(x float64, y float64) *Point {
	coords := Coordinates{
		X: x,
		Y: y,
	}
	pointCoordinates := []Coordinates{
		coords,
	}

	return &Point{
		Coordinates: pointCoordinates,
	}
}
