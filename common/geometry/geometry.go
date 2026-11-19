package geometry

type GeometryType string

const (
	PolygonType GeometryType = "POLYGON"
	PointType   GeometryType = "POINT"
	LineType    GeometryType = "LINE"
	CircleType  GeometryType = "CIRCLE"
)

const MinTypeStringLength = 4
const MaxTypeStringLength = 7

func GeometryTypes() []string {
	return []string{
		string(PolygonType),
		string(PointType),
		string(LineType),
		string(CircleType),
	}
}

type Geometry interface {
	Contains(other Geometry) (bool, error)
	Intersects(other Geometry) (bool, error)
	GetType() GeometryType
	GetCentroid() *Point
	Equals(other Geometry) bool
}
