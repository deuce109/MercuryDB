package serialization

import (
	"bytes"
	"testing"

	"github.com/deuce109/mercurydb/v2/common/geometry"
)

func TestDeserialze(t *testing.T) {
	// Example struct to serialize

	example := geometry.Geometry{
		Coordinates: []geometry.Coordinates{
			{X: 5, Y: 4},
		},
		Name: "Test",
	}

	// Serialize the example struct
	data, err := SerializeDataChunk(example)
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	// Deserialize the data back into a struct
	var deserializedExample geometry.Geometry
	err = DeserializeGeometry(data, &deserializedExample)
	if err != nil {
		t.Fatalf("Failed to deserialize: %v", err)
	}

	// Check if the deserialized data matches the original
	if deserializedExample.Name != example.Name || deserializedExample.Coordinates == nil {
		t.Fatalf("Deserialized data does not match original: got %+v, want %+v", deserializedExample, example)
	}
}

func TestDeserializeEmptyData(t *testing.T) {
	// Attempt to deserialize empty data
	var deserializedExample interface{}

	err := DeserializeDataChunk(&bytes.Buffer{}, &deserializedExample)
	if err == nil {
		t.Fatal("Expected error when deserializing empty data, got nil")
	}
}
