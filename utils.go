package main

import (
	"os"
	"path/filepath"
)

func getFullPath(path string) (string, error) {
	if path[0] == '/' {
		return path, nil
	}
	if path[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return path, err
		}
		return filepath.Join(home, path[1:]), nil
	}
	workDir, err := os.Getwd()
	if err != nil {
		return path, err
	}
	return filepath.Join(workDir, path), nil
}
