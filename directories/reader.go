package directories

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/deuce109/mercurydb/v2/config"
	"github.com/deuce109/mercurydb/v2/geometry"
	"github.com/deuce109/mercurydb/v2/serialization"
)

type ReaderResult struct {
	Geometry *geometry.Geometry
	Err error
}

type DataReader struct {
	Results chan ReaderResult
}

func NewDataReader() *DataReader {
	return &DataReader{
		Results: make(chan ReaderResult),
	}
}

func (dr *DataReader) ReadDir(path string, recursive bool) error {
	files, err := GetFiles(path, recursive)
	if err != nil {
		return err
	}

	dr.ReadFiles(files...)
	return nil
}

func (dr *DataReader) ReadFiles(files ...string) {

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go dr.readFile(file, &wg)
	}

	wg.Wait()
	close(dr.Results)
}

func (dr *DataReader) ReadSingleFile(id string) {
	path, err := filepath.Abs(filepath.Join(config.Config.DataPath, fmt.Sprintf("%s.bin", id)))
	
	if err != nil {
		dr.Results <- ReaderResult{
			Geometry: nil,
			Err: err,
		}
		return
	}
	
	var wg sync.WaitGroup

	wg.Add(1)

	go dr.readFile(path, &wg)
} 

func (dr *DataReader) readFile(file string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := os.ReadFile(file)
	if err != nil {
		dr.Results <- ReaderResult{Geometry: nil, Err: err}
		return
	}

	geo := &geometry.Geometry{}

	err = serialization.Deserialize(data, geo)

	dr.Results <- ReaderResult{Geometry: geo, Err: err}

}
