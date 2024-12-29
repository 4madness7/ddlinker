package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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

func (c *Config) Validate() (map[string][]string, map[string][]error) {
	allErrs := map[string][]error{}
	allWarns := map[string][]string{}
	uniqueNames := map[string]struct{}{}
	uniquePaths := map[string]struct{}{}

	for _, dest := range c.Destinations {
		// check if name of dest is empty
		if dest.Name == "" {
			allErrs[dest.Name] = append(
				allErrs[dest.Name],
				fmt.Errorf("Destination with no name."),
			)
		}

		// check if path of dest is empty
		hasPath := true
		if dest.Path == "" {
			allErrs[dest.Name] = append(
				allErrs[dest.Name],
				fmt.Errorf("Destination with no path."),
			)
			hasPath = false
		}

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
		if hasPath && dest.Path[0] != '~' && dest.Path[0] != '/' {
			allErrs[dest.Name] = append(
				allErrs[dest.Name],
				fmt.Errorf("Path is not absolute: '%s'", dest.Path),
			)
		}

		// check if dir exists
		path := dest.Path
		if hasPath && dest.Path[0] == '~' {
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

		// check if path contains '/'
		if strings.Contains(dest.Path, "//") {
			allWarns[dest.Name] = append(
				allWarns[dest.Name],
				fmt.Sprintf("Path '%s' contains consecutive '/'. It should not affect linking, but avoid using consecutive '/'.", dest.Path),
			)
		}

        // checks if no links are provided
        if len(dest.Links) == 0 {
			allWarns[dest.Name] = append(
				allWarns[dest.Name],
				fmt.Sprintf("No links provided. It should not affect linking, but avoid destinations with empty links."),
			)
        }

		for _, link := range dest.Links {
			// check if dir/file links exist
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

			// check if dir/file contains '/'
			if strings.Contains(link, "/") {
				allWarns[dest.Name] = append(
					allWarns[dest.Name],
					fmt.Sprintf("Directory/File '%s' contains '/'. It should not affect linking, but avoid using '/' and subdirs on links.", link),
				)
			}
		}

	}
	return allWarns, allErrs
}
