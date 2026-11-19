package serialization

import (
	"testing"

	"github.com/deuce109/mercurydb/v2/common/geometry"
)

// geometryTestCase pairs a geometry input with the validation it should pass.
type geometryTestCase struct {
	name     string
	geometry geometry.Geometry
}
var cases []*geometryTestCase = []*geometryTestCase{
	{
		name: "Circle serailization",
		geometry: &geometry.Circle{
			Center: &geometry.Point{
				X: 0, Y: 5.505,
			},
			Radius: 20.0000000007,
		},
	},
	{
		name: "Point serialization",
		geometry: &geometry.Point{
			X: 12.25,
			Y: -8.5,
		},
	},
	{
		name: "Line serialization",
		geometry: &geometry.Line{
			 Coordinates: []*geometry.Point{ 
				{X: -10, Y: 2.5}, 
				{X: 30.75, Y: -40},
			},
		},
	},
	{
		name: "Polygon serialization",
		geometry: &geometry.Polygon{
			Coordinates: []*geometry.Point{
				{X: 0, Y: 0},
				{X: 10, Y: 0},
				{X: 10, Y: 10},
				{X: 0, Y: 10},
			},
		},
	},
}

func TestJsonSerializationLifecycle(t *testing.T) {
	var data map[string]string = map[string]string{
		"test": "test",
	}

	serialized, err := SerializeToJson(data)

	if err != nil {
		t.Fatal(err)
	}

	got, err := DesierializeJsonReaderToMap(serialized)

	if err != nil {
		t.Fatal(err)
	} else if got["test"] != data["test"] {
		t.Fatalf("Deserialized data does not match, expected %s, got %s", data["test"], got["test"])
	}
}

func TestGeometrySerializationLifecycle(t *testing.T) {
	t.Helper()
	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			serializeAndDeserializeGeometry(t, tc.geometry)
		})
	}
}

func serializeAndDeserializeGeometry(t *testing.T, input geometry.Geometry) {
	r, err := SerializeGeometry(input)

	if err != nil {
		t.Fatal(err)
	}

	output, err := DeserializeGeometry(r)

	if err != nil {
		t.Fatal(err)
	}

	if match := input.Equals(output); !match {
		t.Fatal("Values do not match")
	}
}