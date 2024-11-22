package config

import (
	"github.com/Siroshun09/serrors"
	"os"
)

func getRequiredString(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", serrors.New("env '" + key + "' is required")
	}
	return value, nil
}
