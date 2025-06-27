package conf

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"log/slog"
)

type ConfigInterface interface {
	validate() error
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env:"SEVER_HOST"`
	Port     string `yaml:"port" env:"SERVER_PORT" default:"3306"`
	User     string `yaml:"user" env:"DATABASE_USER"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD"`
	Name     string `yaml:"name" env:"DATABASE_NAME"`
}

func (c DatabaseConfig) validate() error {
	if c.Host == "" {
		return fmt.Errorf("database host is required")
	}
	if c.Port == "" {
		return fmt.Errorf("database port is required")
	}
	if c.User == "" {
		return fmt.Errorf("database user is required")
	}
	if c.Password == "" {
		return fmt.Errorf("database password is required")
	}
	if c.Name == "" {
		return fmt.Errorf("database name is required")
	}
	return nil
}

type ServerConfig struct {
	Host string `yaml:"host" env:"SERVER_HOST" default:"${SERVER_HOST | localhost}"`
	Port string `yaml:"port" env:"SERVER_PORT" default:"${SERVER_PORT | 50051}"`
}

func (c ServerConfig) validate() error {
	if c.Host == "" {
		return fmt.Errorf("server host is required")
	}
	if c.Port == "" {
		return fmt.Errorf("server port is required")
	}
	return nil
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func (c Config) validate() error {
	err := c.Server.validate()
	if err != nil {
		slog.Error("Error loading server config", "error", err)
		return err
	}
	//err = c.DatabaseConfig.validate()
	if err != nil {
		slog.Error("Error loading database config", "error", err)
		return err
	}
	return nil
}

func Load(configFiles ...string) (*Config, error) {
	cfg := config.NewWithOptions("journalful",
		//config.ParseDefault, // Parse default values from struct tags
		config.ParseEnv, // Parse environment variables in config values
	)
	//cfg.Options().DecoderConfig.TagName = "yaml"
	cfg.AddDriver(yaml.Driver)

	err := cfg.LoadExists(configFiles...)
	if err != nil {
		slog.Warn("failed to load config files")
		return nil, err
	}

	//load env values
	envConversionMap := map[string]string{
		"SERVER_HOST": "server.host",
		"SERVER_PORT": "server.port",

		"DB_HOST":     "database.host",
		"DB_PORT":     "database.port",
		"DB_USER":     "database.user",
		"DB_PASSWORD": "database.PASSWORD",
		"DB_NAME":     "database.name",
	}
	cfg.LoadOSEnvs(envConversionMap)

	var appCfg Config

	err = cfg.Decode(&appCfg)
	if err != nil {
		slog.Error("failed to decode config")
		return nil, err
	}
	err = appCfg.validate()
	if err != nil {
		slog.Error("failed to validate config")
		return nil, err
	}
	return &appCfg, nil
}
