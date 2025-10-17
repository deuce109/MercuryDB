package main

import (
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"github.com/deuce109/mercurydb/v2/db/config"
	"github.com/deuce109/mercurydb/v2/db/networking"
)

func main() {

	flag.String("log_level", "info", "Set the logging level (debug, info, warn, error)")
	flag.Parse()

	setLoggerLevel()

	config.Load()

	slog.Info("MercuryDB starting...")

	networking.RegisterNetworkListener(config.GetInterfaceIP(), config.GetPort())

	if runtime.GOOS != "windows" {
		networking.RegisterUnixSocketListener(config.GetUnixSocketPath())
	}

	networking.StartListeners()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer close(c)
}

func close(c chan os.Signal) {
	for sig := range c {

		networking.CloseListeners()

		slog.Info("Received signal, shutting down...", "signal", sig)
		os.Exit(0)
	}
}

func setLoggerLevel() {

	var level slog.Level

	levelFlag := flag.Lookup("log_level")

	if levelFlag == nil || levelFlag.Value == nil {
		level = slog.LevelInfo
	}

	switch strings.ToLower(levelFlag.Value.String()) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	slog.SetLogLoggerLevel(level)
}
