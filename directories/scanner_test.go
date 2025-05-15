package directories

import (
	"testing"
)

func TestGetFilesInSubdirectory(t *testing.T) {
	dir := "../testFiles"

	files, err := GetFiles(dir, true)
	if err != nil {
		t.Fatalf("Failed to get files: %v", err)
	}

	if len(files) != 1 {
		t.Fatalf("Expected 1 files, got %d", len(files))

	}
}
