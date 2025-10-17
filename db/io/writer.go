package io

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/deuce109/mercurydb/v2/db/models"
)

type DataWriter struct {
	wg         *sync.WaitGroup
	outputPath string
}

func NewDataWriter(outputPath string) *DataWriter {
	return &DataWriter{
		wg:         &sync.WaitGroup{},
		outputPath: outputPath,
	}
}

func (w *DataWriter) createFile(geoData *models.GeoData) {
	defer w.wg.Done()
	path, err := filepath.Abs(filepath.Join(w.outputPath, fmt.Sprintf("%s.mercury", geoData.Id)))
	if err != nil {
		fmt.Printf("Error getting file absolute path: %s", err)
	}
	err = NewMercuryFile(path).WriteGeoData(geoData)
	if err != nil {
		fmt.Printf("Error writing geoData to file: %s", err)
	}
}
func (w *DataWriter) WriteGeoDatas(geoData ...*models.GeoData) {

	for _, data := range geoData {
		w.wg.Add(1)
		go w.createFile(data)
	}

	w.wg.Wait()
}
