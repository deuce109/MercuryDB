package io

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/deuce109/mercurydb/v2/common/collections"
)

type DirectoryScanner struct {
	directory string
}

func NewDirectoryScanner(path string) (*DirectoryScanner, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	return &DirectoryScanner{
		directory: absPath,
	}, nil
}

func (s *DirectoryScanner) FindFilesByExt(ext string) ([]string, error) {
	return filepath.Glob(fmt.Sprintf("%s*%s", s.directory, ext))
}

func (s *DirectoryScanner) FindMercuryFiles() []*MercuryFile {
	mercuryFilePaths, err := s.FindFilesByExt(".mercury")

	if err != nil {
		slog.Error("Unable to find MercuryFiles", "error", err)
	}

	files := make([]*MercuryFile, 0, len(mercuryFilePaths))

	collections.
		NewCollection(mercuryFilePaths).FilterNonZero().
		ForEach(func(item string) {
			file := NewMercuryFile(item)
			files = append(files, file)
		})

	return files
}
