package testutils

import (
	"github.com/Siroshun09/serrors"
	"os"
	"path/filepath"
)

// GetProjectRoot searches for the project root directory containing go.mod
func GetProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", serrors.WithStackTrace(err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", serrors.WithStackTrace(os.ErrNotExist)
}
