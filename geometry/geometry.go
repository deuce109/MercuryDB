package geometry

import "math"

type Coordinates struct {
	X float64
	Y float64
}

type Geometry struct {
	Coordinates []Coordinates
	Name string
	Id string
	Data map[string]interface{}
}

func (c *Coordinates) DistanceTo(other Coordinates) float64 {
	return math.Sqrt(math.Pow(other.X-c.X, 2) + math.Pow(other.Y-c.Y, 2))
}
