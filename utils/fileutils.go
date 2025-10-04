package utils

import (
	"os"
	"path/filepath"
)

func IsValidPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CleanPath(path string) string {
	return filepath.Clean(path)
}
