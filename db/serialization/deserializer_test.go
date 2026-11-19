package serialization

import (
	"bytes"
	"errors"
	"testing"
	"testing/iotest"

	"github.com/deuce109/mercurydb/v2/common/geometry"
)

func TestDeserialze(t *testing.T) {
	// Example struct to serialize

	example := &geometry.Circle{
		Center: &geometry.Point{
			X: 5,
			Y: 6,
		},
		Radius: 0.4,
	}

	// Serialize the example struct
	data, err := SerializeGeometry(example)
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}
	geo, err := DeserializeGeometry(data)

	if err != nil {
		t.Fatalf("Failed to deserialize: %v", err)
	}

	

	// Check if the deserialized data matches the original
	if geo.GetType() != "CIRCLE" {
		t.Fatalf("Incorrect geometry type got: \"%s\" expected \"CIRCLE\"", geo.GetType())
	}

	if circle, ok := geo.(*geometry.Circle); ok {
		if circle.Center.X != 5.0 || circle.Center.Y != 6.0 || circle.Radius != 0.4 {
			t.Fatalf("Deserialized data is incorrect expeted X: 5, Y: 6, Radius 0.4 got X: %f, Y: %f, Radius: %f", circle.Center.X, circle.Center.Y, circle.Radius)
		}
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

func TestBadGeometryType(t *testing.T) {
	var buffer bytes.Buffer

	_, err := buffer.WriteString("BADSTRING")

	if err != nil {
		t.Fatal(err)
	}

	_, err = DeserializeGeometry(&buffer)

	if err.Error() != "Unexpected geometry type. Found: BADSTRIN"{
		t.Fatal(err)
	}
}

func TestDeserializeBadGeometryReader(t *testing.T) {
	expectedErr := errors.New("simulated read error")
	
	// Fails immediately on the first Read()
	failingReader := iotest.ErrReader(expectedErr)

	_, err := DeserializeGeometry(failingReader)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}

func TestDeserializeBadJsonString(t *testing.T) {

	badString := "{]"

	output := DeserializeJsonStringToMap(badString)

	if output != nil {
		t.Fatal("How did that serialize properly")
	}
}

func TestDeserializeBadJsonReader(t *testing.T) {
	expectedErr := errors.New("simulated read error")
	
	// Fails immediately on the first Read()
	failingReader := iotest.ErrReader(expectedErr)

	_, err := DesierializeJsonReaderToMap(failingReader)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}
