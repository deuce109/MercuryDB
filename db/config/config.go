package config

import (
	"log/slog"

	"github.com/deuce109/envparse/v2"
)

var config *dBConfig = &dBConfig{}

var defaultConfig *dBConfig = &dBConfig{
	Port:           12000,
	InterfaceIP:    "0.0.0.0",
	UnixSocketPath: "/opt/mercury/db.sock",
}

type dBConfig struct {
	Port           uint16 `envparse:"MERCURY_PORT"`
	InterfaceIP    string `envparse:"MERCURY_INTERFACE_IP"`
	UnixSocketPath string `envparse:"MERCURY_SOCKET_PATH"`
	DbPath         string `envparse:"MERCURY_USER_DB_PATH"`
	DataPath       string `envparse:"MERCURY_DATA_PATH"`
}

func Load() {
	slog.Debug("Loading configuration...")
	err := envparse.Parse(config)
	if err != nil {
		slog.Error("Error loading configuration, using defaults", "error", err)
		config = defaultConfig
	}

	slog.Debug("Got configuration",
		"interface", GetInterfaceIP(),
		"port", GetPort(),
		"unix_socket_path", GetUnixSocketPath())
}

func GetPort() uint16 {
	return config.Port
}

func GetInterfaceIP() string {
	return config.InterfaceIP
}

func GetUnixSocketPath() string {
	return config.UnixSocketPath
}

func GetUserDbPath() string {
	return config.DbPath
}

func GetDataPath() string {
	return config.DataPath
}
