package directories

import (
	"os"
	"sync"

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
