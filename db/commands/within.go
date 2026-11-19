package commands

import (
	"log/slog"
	"sync"

	"github.com/deuce109/mercurydb/v2/common/geometry"
	"github.com/deuce109/mercurydb/v2/db/io"
)

const WITHIN string = "WITHIN"

func fileGeoWithin(file *io.MercuryFile, geo geometry.Geometry, results chan *io.MercuryFile, wg *sync.WaitGroup) {
	if contains, err := geo.Contains(file.GetGeometry()); contains {
		results <- file
	} else if err != nil {
		slog.Error("Error check if geometry contains another", "error", err)
	}
	wg.Done()
}

func (h *CommandHandler) within(geo geometry.Geometry) []*io.MercuryFile {
	return h.findFilesMatching(geo, fileGeoWithin)
}
