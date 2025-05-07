package directories

import (
	"fmt"
	"os"
	"sync"
)

type DataReader struct {
	Results chan []byte
}

func NewDataReader() *DataReader {
	return &DataReader{
		Results: make(chan []byte),
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
		fmt.Printf("Error reading file %s: %v", file, err)
	}
	dr.Results <- data
}
