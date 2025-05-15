package geometry

func pointFromCoords(x float64, y float64) *Geometry {
	coords := Coordinates{
		X: x,
		Y: y,
	}
	pointCoordinates := []Coordinates{
		coords,
	}

	return &Geometry{
		Coordinates: pointCoordinates,
	}
}
