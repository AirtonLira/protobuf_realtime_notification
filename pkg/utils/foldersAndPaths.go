package utils

import (
	"os"
	"path/filepath"
)

func GetProjectRoot() (string, error) {
	// Assume que o projeto é iniciado a partir do diretório raiz
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); !os.IsNotExist(err) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", err
		}
		dir = parent
	}
}
