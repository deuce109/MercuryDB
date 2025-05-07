package serialization

import (
	"testing"
)

func TestSerialize(t *testing.T) {

	example := Example{Name: "Test", Value: 42}

	// Serialize the example struct
	data, err := Serialize(example)
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
	_, err := Serialize(nil)
	if err == nil {
		t.Fatal("Expected error when serialziing nil struct, got nil")
	}
}
