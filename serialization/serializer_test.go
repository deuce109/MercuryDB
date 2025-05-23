package serialization

import (
	"testing"

	"github.com/deuce109/mercurydb/v2/geometry"
)

func TestSerialize(t *testing.T) {

	example := geometry.Geometry{
		Coordinates: []geometry.Coordinates{},
	}

	// Serialize the example struct
	data, _, err := SerializeGeometry(example)
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	// Check if the data is not empty
	if len(data) == 0 {
		t.Fatal("Serialized data is empty")
	}
}

func TestSerializeNilValue(t *testing.T) {

	// Serialize the empty struct
	_, _, err := SerializeDataChunk(nil)
	if err == nil {
		t.Fatal("Expected error when serialziing nil struct, got nil")
	}
}
