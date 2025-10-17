package commands

import (
	"fmt"

	"github.com/deuce109/mercurydb/v2/db/io"
	"github.com/deuce109/mercurydb/v2/db/models"

	"github.com/google/uuid"
)

const INSERT string = "INSERT"

func insert(geo *models.GeoData) (string, error) {
	if geo.Id == "" {
		geo.Id = uuid.NewString()
	}

	file := io.NewMercuryFile(geo.Id)
	return fmt.Sprintf("Inserted file with ID '%s'", geo.Id), file.WriteGeoData(geo)
}
