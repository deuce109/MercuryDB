package commands

import (
	"sync"

	"github.com/deuce109/mercurydb/v2/common/geometry"
	"github.com/deuce109/mercurydb/v2/db/io"
)

const INTERSECTS string = "INTERSECTS"

func fileContainsIntersectingGeo(file *io.MercuryFile, geo *geometry.Geometry, results chan *io.MercuryFile, wg *sync.WaitGroup) {
	if file.GeometryIntersects(geo) {
		results <- file
	}
	wg.Done()
}

func (h *CommandHandler) intersects(geo *geometry.Geometry) []*io.MercuryFile {
	return h.findFilesMatching(geo, fileContainsIntersectingGeo)
}
