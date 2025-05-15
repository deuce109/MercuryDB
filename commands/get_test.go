package commands

import (
	"testing"

	"github.com/deuce109/mercurydb/v2/config"
	"github.com/deuce109/mercurydb/v2/directories"
	"github.com/deuce109/mercurydb/v2/geometry"
)

func TestMain(m *testing.M) {
	config.Load()

	geom := &geometry.Geometry{
		Coordinates: []geometry.Coordinates{
			{X:5, Y:7},
			{X:5, Y:4},
			{X:7, Y: 7},
		},
	}

	directories.WriteToDirectory(geom,config.Config.DataPath)
	m.Run()
}

func TestGetWithWithinFilter(t *testing.T) {
	wkt := "((30 10, 40 40, 20 40, 10 20, 30 10))"

	_, err := Get("WITHIN", wkt)

	if err != nil {
		t.Fatalf("Expected Get with Within filter to not fail, failed with error: %v", err)
	}
}

func TestGetWithIntersectsFilter(t *testing.T) {
	wkt := "((30 10, 40 40, 20 40, 10 20, 30 10))"

	_, err := Get("INTERSECTS", wkt)

	if err != nil {
		t.Fatalf("Expected Get with Within filter to not fail, failed with error: %v", err)
	}
}

func TestDefaultCase(t *testing.T) {
	_, err := Get("default", "")

	if err == nil {
		t.Fatal("Expected error to not be nil")
	}

}