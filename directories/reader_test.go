package directories

import "testing"

func TestDataReaderReadDir(t *testing.T) {
	dir := "../testFiles"

	dr := NewDataReader()
	go dr.ReadDir(dir, true)

	for data := range dr.Results {
		if len(data) == 0 {
			t.Fatalf("Expected data, got empty result")
		} else if len(data) != 20 {
			t.Fatalf("Did not get all data, got %d bytes", len(data))
		}
	}
}
