package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	BaseURL string `toml:"base_url"`
	Token   string `toml:"token"`
}

var defaultConfig = Config{
	BaseURL: "",
	Token:   "",
}

func initDefaultConfig() error {
	configDir := getConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}
	data, err := toml.Marshal(defaultConfig)
	if err != nil {
		return err
	}
	configPath := getConfigPath()
	return os.WriteFile(configPath, data, 0644)
}

func getConfigDir() string {
	homePath, _ := os.UserHomeDir()
	configPath := filepath.Join(homePath, ".config", "op-cli")
	return configPath
}

func getConfigPath() string {
	configPath := filepath.Join(getConfigDir(), "config.toml")
	return configPath
}

func LoadConfig() (*Config, error) {
	configPath := getConfigPath()
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		if err := initDefaultConfig(); err != nil {
			return nil, fmt.Errorf("config init error: %w", err)
		}
		fmt.Println("config init successfully: ", configPath)
		return &defaultConfig, nil
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("config read error: %w", err)
	}
	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("config unmarshal error: %w", err)
	}
	return &config, nil
}

func SaveConfig(config *Config) error {
	configPath := getConfigPath()
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		configDir := getConfigDir()
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return err
		}
	}
	data, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
