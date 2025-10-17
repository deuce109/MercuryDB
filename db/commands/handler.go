package commands

import (
	coreIo "io"
	"log/slog"
	"path/filepath"
	"strings"
	"sync"

	"github.com/deuce109/mercurydb/v2/common/geometry"
	"github.com/deuce109/mercurydb/v2/db/config"
	"github.com/deuce109/mercurydb/v2/db/io"
	"github.com/deuce109/mercurydb/v2/db/models"
	"github.com/deuce109/mercurydb/v2/db/serialization"
)

type CommandHandler struct {
	directoryScanner *io.DirectoryScanner
}

func newCommandHandler() (*CommandHandler, error) {
	fileScanner, err := io.NewDirectoryScanner(filepath.Join(config.GetUserDbPath(), "geo"))

	if err != nil {
		return nil, err
	}

	return &CommandHandler{
		directoryScanner: fileScanner,
	}, nil
}

func DefaultCommandHandler() (*CommandHandler, error) {
	return newCommandHandler()
}

func (h *CommandHandler) HandleMessage(message string) (coreIo.Reader, error) {
	command, args, err := ParseCommand(message)

	if err != nil {
		return strings.NewReader(err.Error()), nil
	}

	switch command {
	case PING:
		return h.ping(), nil
	case CONTAINS:
		return h.getDataMatching(h.contains, args)
	case INSERT:
		return h.insert(args)
	case INTERSECTS:
		return h.getDataMatching(h.intersects, args)
	case WITHIN:
		return h.getDataMatching(h.within, args)
	default:
		return nil, nil
	}
}

func (h *CommandHandler) findFilesMatching(geo *geometry.Geometry, matcher func(*io.MercuryFile, *geometry.Geometry, chan *io.MercuryFile, *sync.WaitGroup)) []*io.MercuryFile {
	files := h.directoryScanner.FindMercuryFiles()

	results := make(chan *io.MercuryFile)
	wg := &sync.WaitGroup{}

	for _, f := range files {
		wg.Add(1)
		go matcher(f, geo, results, wg)
	}

	wg.Wait()
	close(results)

	data := make([]*io.MercuryFile, 0, len(files))

	for r := range results {
		data = append(data, r)
	}

	return data
}

func (h *CommandHandler) getDataMatching(fileFinder func(*geometry.Geometry) []*io.MercuryFile, args []string) (coreIo.Reader, error) {
	geo := geometry.FromWKT(args[0])
	files := fileFinder(geo)
	return h.getDataFromFiles(files, args[1:])
}

func (h *CommandHandler) getDataFromFiles(files []*io.MercuryFile, metadataNames []string) (coreIo.Reader, error) {

	results := make(chan *models.GeoData)

	wg := &sync.WaitGroup{}

	for _, f := range files {
		wg.Add(1)
		go h.getDataFromFile(f, metadataNames, results, wg)
	}

	wg.Wait()
	close(results)

	resultList := make([]*models.GeoData, 0, len(results))
	for r := range results {
		resultList = append(resultList, r)
	}

	return models.SerializeGeoDataList(resultList)
}

func (h *CommandHandler) getDataFromFile(f *io.MercuryFile, metadataNames []string, results chan *models.GeoData, wg *sync.WaitGroup) {
	defer wg.Done()

	err := f.Open()
	if err != nil {
		slog.Error("Error opening file", "path", f.Path, "error", err)
		return
	}

	defer f.Close()

	data, err := f.ReadGeoData(metadataNames)
	if err != nil {
		results <- data
	}

}

func (h *CommandHandler) insert(args []string) (coreIo.Reader, error) {
	geo := geometry.FromWKT(args[0])

	metaData := make(models.Metadata)

	for i := 1; i < len(args); {

		name := args[i]

		value := serialization.DeserializeJsonStringToMap(args[i+1])

		metaData[name] = value

		i += 2
	}

	geoData := &models.GeoData{
		Geometry: geo,
		Metadata: metaData,
	}

	result, err := insert(geoData)

	if err != nil {
		slog.Error("Error inserting geodata", "error", err)
		return strings.NewReader("Error insterting geodata"), err
	}

	return strings.NewReader(result), nil

}
