package directories

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/deuce109/mercurydb/v2/geometry"
	"github.com/deuce109/mercurydb/v2/serialization"
)

type Item struct {
	Id string
}

func WriteToDirectory(obj *geometry.Geometry, path string) error {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	data, err := serialization.Serialize(obj)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(fullPath, fmt.Sprintf("%s.bin", obj.Id)), data, 0644)
	if err != nil {
		return err
	}

	return nil
}