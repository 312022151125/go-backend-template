package conf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type Server struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
}

type Log struct {
	Level string `json:"level" mapstructure:"level"`
}

type Database struct {
	Type string `json:"type" mapstructure:"type"`
	Path string `json:"path" mapstructure:"path"`
}

type Config struct {
	Server   Server   `mapstructure:"server"`
	Log      Log      `mapstructure:"logging"`
	Database Database `mapstructure:"database"`
}

var AppConfig Config

func Load(path string) error {
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath("data")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix(APP_NAME)

	setDefaults()

	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		missing := false
		var notFound viper.ConfigFileNotFoundError
		if errors.As(err, &notFound) {
			missing = true
		} else if errors.Is(err, os.ErrNotExist) {
			missing = true
		}

		if !missing {
			return fmt.Errorf("error reading config file: %w", err)
		}

		outPath := path
		if outPath == "" {
			outPath = "data/config.json"
		}

		log.Infof("Config file not found, creating default at %s", outPath)
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}
		if err := viper.SafeWriteConfigAs(outPath); err != nil {
			return fmt.Errorf("failed to create default config: %w", err)
		}
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("unable to decode config into struct: %w", err)
	}
	if level, err := log.ParseLevel(AppConfig.Log.Level); err == nil {
		log.SetLevel(level)
	} else {
		return fmt.Errorf("invalid log level: %w", err)
	}
	return nil
}

func setDefaults() {
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.type", "sqlite")
	viper.SetDefault("database.path", "data/data.db")
	viper.SetDefault("logging.level", "info")
}
