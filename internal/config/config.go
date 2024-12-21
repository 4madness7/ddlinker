package config

import (
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
	configFileName = ".ddlinker_config.toml"
)

type Config struct {
	Destinations []destination `toml:"destinations"`
}

type destination struct {
	Name  string   `toml:"name"`
	Path  string   `toml:"path"`
	Links []string `toml:"links"`
}

func Read() (Config, error) {
	currDir, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(filepath.Join(currDir, configFileName))
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
