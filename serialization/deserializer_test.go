package serialization

import "testing"

type Example struct {
	Name  string
	Value int
}

func TestDeserialze(t *testing.T) {
	// Example struct to serialize

	example := Example{Name: "Test", Value: 42}

	// Serialize the example struct
	data, err := Serialize(example)
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	// Deserialize the data back into a struct
	var deserializedExample Example
	err = Deserialize(data, &deserializedExample)
	if err != nil {
		t.Fatalf("Failed to deserialize: %v", err)
	}

	// Check if the deserialized data matches the original
	if deserializedExample.Name != example.Name || deserializedExample.Value != example.Value {
		t.Fatalf("Deserialized data does not match original: got %+v, want %+v", deserializedExample, example)
	}
}

func TestDeserializeEmptyData(t *testing.T) {
	// Attempt to deserialize empty data
	var deserializedExample Example
	err := Deserialize([]byte{}, &deserializedExample)
	if err == nil {
		t.Fatal("Expected error when deserializing empty data, got nil")
	}
}
