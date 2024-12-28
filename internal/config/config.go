package config

import (
	"fmt"
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

func Read() (*Config, error) {
	currDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(currDir, configFileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Validate() map[string][]error {
	allErrs := map[string][]error{}
	uniqueNames := map[string]struct{}{}
	uniquePaths := map[string]struct{}{}
	for _, dest := range c.Destinations {
		// check if there are duplicate names
		_, exists := uniqueNames[dest.Name]
		if exists {
			allErrs[dest.Name] = append(
				allErrs[dest.Name],
				fmt.Errorf("Duplicate destinations with name '%s'.", dest.Name),
			)
		} else {
			uniqueNames[dest.Name] = struct{}{}
		}

		// check if there are duplicate paths
		_, exists = uniquePaths[dest.Path]
		if exists {
			allErrs[dest.Name] = append(
				allErrs[dest.Name],
				fmt.Errorf("Duplicate destinations with path '%s'.", dest.Path),
			)
		} else {
			uniquePaths[dest.Path] = struct{}{}
		}

		// check if dir is absolute path
		if dest.Path[0] != '~' && dest.Path[0] != '/' {
			allErrs[dest.Name] = append(
				allErrs[dest.Name],
				fmt.Errorf("Path is not absolute: '%s'", dest.Path),
			)
		}

		// check if dir exists
		path := dest.Path
		if dest.Path[0] == '~' {
			home, err := os.UserHomeDir()
			if err != nil {
				allErrs[dest.Name] = append(allErrs[dest.Name], err)
				continue
			}
			path = filepath.Join(home, dest.Path[1:])
		}

		_, err := os.Stat(path)
		notExists := os.IsNotExist(err)
		if notExists {
			allErrs[dest.Name] = append(
				allErrs[dest.Name],
				fmt.Errorf("Directory not found: '%s'", dest.Path),
			)
		}

		// check if dir/file links exist
		for _, link := range dest.Links {
			dir, err := os.Getwd()
			if err != nil {
				allErrs[dest.Name] = append(allErrs[dest.Name], err)
				continue
			}
			_, err = os.Stat(filepath.Join(dir, link))
			notExists := os.IsNotExist(err)
			if notExists {
				allErrs[dest.Name] = append(
					allErrs[dest.Name],
					fmt.Errorf("Directory/File not found: './%s'", link),
				)
			}
		}

	}
	return allErrs
}
