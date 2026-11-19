package config

import (
	"log/slog"
	"path"
	"runtime"

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
	DataPath       string `envparse:"MERCURY_DATA_PATH"`
	UserDatabaseConnectionString string `envparse:"MERCURY_USER_DB_CONNECTION_STRING`
}

func buildDefaultConfig() *dBConfig {
	dataPath := path.Join("mercury", "data")
	if runtime.GOOS == "windows" {
		path.Join("C:", dataPath)
	}

	defaultConfig.DataPath = dataPath
	return defaultConfig
}

func Load() {
	slog.Debug("Loading configuration...")
	err := envparse.Parse(config)
	if err != nil {
		slog.Error("Error loading configuration, using defaults", "error", err)
		config = buildDefaultConfig()
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

func GetDataPath() string {
	return config.DataPath
}

func GetUserDBConnectionString() string {
	return config.UserDatabaseConnectionString
}
