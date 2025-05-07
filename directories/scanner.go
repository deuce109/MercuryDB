package directories

import (
	"os"
	"path/filepath"
)

func GetFiles(dirPath string, recursive bool) ([]string, error) {

	path, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	return getFilesFromPath(path, recursive)

}

func getFilesFromPath(path string, recursive bool) ([]string, error) {
	files := make([]string, 0)

	subPaths, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	for _, file := range subPaths {
		if file.IsDir() && recursive {
			subDirFiles, err := getFilesFromPath(filepath.Join(path, file.Name()), recursive)
			if err != nil {
				return nil, err
			}
			files = append(files, subDirFiles...)
		} else {
			files = append(files, filepath.Join(path, file.Name()))
		}
	}

	return files, nil
}
