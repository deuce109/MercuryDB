package serialization

import (
	"bytes"
	"testing"

	"github.com/deuce109/mercurydb/v2/common/geometry"
)

func TestSerialize(t *testing.T) {

	example := &geometry.Geometry{
		Coordinates: []geometry.Coordinates{},
	}

	// Serialize the example struct
	data, err := SerializeGeometry(example)
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	buf, ok := data.(*bytes.Buffer)
	if !ok {
		t.Fatalf("Expected *bytes.Buffer, got %T", data)
	}

	// Check if the data is not empty
	if buf.Len() == 0 {
		t.Fatal("Serialized data is empty")
	}
}

func TestSerializeNilValue(t *testing.T) {

	// Serialize the empty struct
	_, err := SerializeDataChunk(nil)
	if err == nil {
		t.Fatal("Expected error when serialziing nil struct, got nil")
	}
}
