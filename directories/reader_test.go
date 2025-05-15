package directories

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/deuce109/mercurydb/v2/geometry"
	"github.com/deuce109/mercurydb/v2/serialization"
)

func TestMain(m *testing.M) {
	m.Run()

	// defer os.RemoveAll("../testFiles")
}

func createTestFile(path string, t *testing.T) {

	testItem := &geometry.Geometry{
		Id: "123",
	}

	fullPath, err := filepath.Abs(path)

	if err != nil {
		t.Fatalf("Error finding full path: %v", err)
	}

	err = os.MkdirAll(fullPath, os.ModeDir)

	if err != nil {
		t.Fatalf("Error making directory: %v", err)
	}

	testFilePath := filepath.Join(fullPath, "testFile.bin")

	data, err := serialization.Serialize(testItem)

	if err != nil {
		t.Fatalf("Error serializing testItem: %v", err)
	}

	os.WriteFile(testFilePath, data, os.ModeAppend)

	if err != nil {
		t.Fatalf("Error writing test file: %v", err)
	}


}

func TestDataReaderReadDir(t *testing.T) {
	dir := "../testFiles"

	createTestFile(dir, t)
	dr := NewDataReader()
	go dr.ReadDir(dir, true)

	for results := range dr.Results {
		if results.Err != nil {
			t.Fatalf("Expected no error but got error %v", results.Err)
		} else if results.Geometry.Id == "" {
			t.Fatalf("Expected id `123` but got an empty id")
		}
	}
}
