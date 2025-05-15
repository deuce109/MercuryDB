package config

import "os"

type config struct {
	DataPath string
}

var Config *config

func Load() {

	if Config == nil {
		Config = &config{}
	}

	Config.DataPath = os.Getenv("MERCURY_PATH")
}