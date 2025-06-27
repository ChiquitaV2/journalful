package main

import (
	"flag"
	"github.com/chiquitav2/journalful/internal/app"
	"github.com/chiquitav2/journalful/pkg/conf"
	"log/slog"
	"os"
)

var (
	configPath = flag.String("config", "config.yaml", "path to config file")
)

func main() {
	// parse flags  to get the configuration file path

	flag.Parse()
	cfg, err := conf.Load(getConfigPath())
	slog.Info("Loaded configuration", "cfg", cfg)
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		return
	}
	slog.Info("Configuration loaded successfully, server host: " + cfg.Server.Host + ", port: " + cfg.Server.Port)
	slog.Info("Configuration loaded successfully, db host: " + cfg.Database.Host + ", port: " + cfg.Database.Port)

	// Initialize the application with the loaded configuration

	appSvr := app.NewApp(cfg)
	err = appSvr.Init()
	if err != nil {
		slog.Error("Failed to initialize appSvr", "error", err)
		return
	}
	appSvr.Start()

}

// get config path from flags or environment variables and test that the file exists
func getConfigPath() string {

	envPath := os.Getenv("CONFIG_PATH")
	if envPath != "" && checkConfigFileExists(envPath) {
		slog.Info("Using config path from environment variable", "path", envPath)
		return envPath
	}

	if *configPath != "" && checkConfigFileExists(*configPath) {
		slog.Info("Using config path from flag", "path", *configPath)
		return *configPath
	}

	defaultPath := "testdata/config.yml"
	slog.Info("Using default config path", "path", defaultPath)
	return defaultPath
}

func checkConfigFileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		slog.Error("Config file does not exist", "path", path)
		return false
	}
	if err != nil {
		slog.Error("Error checking config file", "path", path, "error", err)
		return false
	}
	slog.Info("Config file exists", "path", path)
	return true
}
