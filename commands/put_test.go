package commands

import (
	"testing"
)

func TestPut(t *testing.T) {
	wkt := "((10 20, 5 6, 10 6))"
	data := map[string]interface{}{
		"name": "Name",
		"id": "123",
	}
	err := Put(wkt, data)
	if err != nil {
		t.Fatalf("Expected no error from Put, got: %v", err)
	}
}

func TestPutOverwrite(t *testing.T) {
	wkt := "((10 20, 5 6, 10 6))"
	data := map[string]interface{}{
		"name": "Name2",
		"id": "123",
		"new": "test",
	}
	err := Put(wkt, data)
	if err != nil {
		t.Fatalf("Expected no error from Put, got: %v", err)
	}
}