package commands

import (
	"github.com/deuce109/mercurydb/v2/config"
	"github.com/deuce109/mercurydb/v2/directories"
	"github.com/deuce109/mercurydb/v2/parsers"
)

func Put(wkt string) {
	geom := parsers.ParseWktCoordinates(wkt)
	directories.WriteToDirectory(geom, config.Config.DataPath)
}